package auth

import (
	"errors"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAuthUseCase_Login(t *testing.T) {
	type fields struct {
		accountRepo  *account.MockIAccountRepository
		tokenManager *MockITokenManager
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		wantToken    string
		wantErr      bool
	}{
		{
			name: "success login",
			args: args{
				username: "test",
				password: "testtesttest",
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().FirstByUsername("test").Return(&entity.Account{
					ID:             1,
					Username:       "test",
					HashedPassword: "$2a$10$CX8FIbhwXvQ7C3X2dE.TnexaiGfvjANy4r07Lqje4t1.QTC86keOy",
				}, nil)
				f.tokenManager.EXPECT().
					GenerateToken("test", mock.Anything, mock.Anything).
					Return("token", nil)
			},
			wantToken: "token",
		},
		{
			name: "error login (unregistered account)",
			args: args{
				username: "test",
				password: "testtesttest",
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().FirstByUsername("test").Return(nil, errors.New("error account not found"))
			},
			wantErr: true,
		},
		{
			name: "error login (error generate token)",
			args: args{
				username: "test",
				password: "testtesttest",
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().FirstByUsername("test").Return(&entity.Account{
					ID:             1,
					Username:       "test",
					HashedPassword: "$2a$10$CX8FIbhwXvQ7C3X2dE.TnexaiGfvjANy4r07Lqje4t1.QTC86keOy",
				}, nil)
				f.tokenManager.EXPECT().
					GenerateToken("test", mock.Anything, mock.Anything).
					Return("", errors.New("error generate token"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.accountRepo = account.NewMockIAccountRepository(t)
			tt.fields.tokenManager = NewMockITokenManager(t)
			uc := &AuthUseCase{
				accountRepo:  tt.fields.accountRepo,
				tokenManager: tt.fields.tokenManager,
			}
			tt.prepareMocks(&tt.fields)
			gotToken, gotExp, err := uc.Login(tt.args.username, tt.args.password)

			assert.Equal(t, tt.wantErr, err != nil, err)
			assert.Equal(t, tt.wantToken, gotToken)
			assert.NotNil(t, gotExp)
		})
	}
}

func TestAuthUseCase_ValidateToken(t *testing.T) {
	type fields struct {
		tokenManager *MockITokenManager
	}
	type args struct {
		token string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         any
		wantErr      bool
	}{
		{
			name: "success validate token",
			args: args{token: "token"},
			prepareMocks: func(f *fields) {
				f.tokenManager.EXPECT().ValidateToken("token").Return(&security.JWTClaims{}, nil)
			},
			want: &security.JWTClaims{},
		},
		{
			name: "error validate token",
			args: args{token: ""},
			prepareMocks: func(f *fields) {
				f.tokenManager.EXPECT().ValidateToken("").Return(nil, errors.New("error generate token"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.tokenManager = NewMockITokenManager(t)
			uc := AuthUseCase{
				tokenManager: tt.fields.tokenManager,
			}

			tt.prepareMocks(&tt.fields)

			got, err := uc.ValidateToken(tt.args.token)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
