package account

import (
	"context"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/account"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/utils/security"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func ctxJWTClaimWithSuperAdminRole() context.Context {
	return context.WithValue(context.Background(), "Authorization", &security.JWTClaims{Role: "super admin"})
}

func TestAccountUseCase_Create(t *testing.T) {
	type fields struct {
		accountRepo *account.MockIAccountRepository
	}
	type args struct {
		ctx           context.Context
		accountParams AccountParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		want         entity.Account
		wantErr      bool
		prepareMocks func(f *fields)
	}{
		{
			name: "success create account",
			args: args{
				ctx: ctxJWTClaimWithSuperAdminRole(),
				accountParams: AccountParams{
					UserName: "test",
					Password: "test123",
					RoleName: "test admin",
				},
			},
			want: entity.Account{
				ID:             1,
				Username:       "test",
				HashedPassword: "random string",
				RoleID:         1,
				Role: entity.Role{
					ID:   1,
					Name: "test admin",
				},
				RegisterApprovalID: 1,
				RegisterApproval: entity.RegisterApproval{
					ID:           1,
					AdminID:      1,
					SuperAdminID: 1,
					Status:       "approved",
				},
				Verified:  true,
				Activated: true,
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().Create(mock.Anything).RunAndReturn(func(e *entity.Account) error {
					e.ID = 1
					e.HashedPassword = "random string"
					e.RoleID = 1
					e.Role.ID = 1
					e.RegisterApprovalID = 1
					e.RegisterApproval.ID = 1
					e.RegisterApproval.AdminID = 1
					e.RegisterApproval.SuperAdminID = 1
					e.RegisterApproval.Status = "approved"
					e.Activated = true
					e.Verified = true
					return nil
				})
			},
		},
	}
	for _, tt := range tests {
		tt.fields.accountRepo = account.NewMockIAccountRepository(t)
		t.Run(tt.name, func(t *testing.T) {
			uc := &AccountUseCase{
				accountRepo: tt.fields.accountRepo,
			}

			tt.prepareMocks(&tt.fields)

			got, err := uc.Create(tt.args.ctx, tt.args.accountParams)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)

		})
	}
}

func TestAccountUseCase_Update(t *testing.T) {
	type fields struct {
		accountRepo *account.MockIAccountRepository
	}
	type args struct {
		ctx           context.Context
		accountParams AccountUpdateParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         *entity.Account
		wantErr      bool
	}{
		{
			name: "success update account",
			args: args{
				ctx: ctxJWTClaimWithSuperAdminRole(),
				accountParams: AccountUpdateParams{
					ID:       1,
					UserName: "updated",
				},
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().Update(mock.Anything).Return(&entity.Account{
					ID:                 1,
					Username:           "updated",
					RoleID:             2,
					Role:               entity.Role{ID: 2, Name: "test admin"},
					RegisterApprovalID: 2,
					RegisterApproval:   entity.RegisterApproval{ID: 2, AdminID: 2, SuperAdminID: 1, Status: "approved"},
					Verified:           true,
					Activated:          true,
				}, nil).Times(1)
			},
			want: &entity.Account{
				ID:                 1,
				Username:           "updated",
				RoleID:             2,
				Role:               entity.Role{ID: 2, Name: "test admin"},
				RegisterApprovalID: 2,
				RegisterApproval:   entity.RegisterApproval{ID: 2, AdminID: 2, SuperAdminID: 1, Status: "approved"},
				Verified:           true,
				Activated:          true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.accountRepo = account.NewMockIAccountRepository(t)
			tt.prepareMocks(&tt.fields)
			uc := &AccountUseCase{
				accountRepo: tt.fields.accountRepo,
			}

			got, err := uc.Update(tt.args.ctx, tt.args.accountParams)

			assert.Equal(t, tt.wantErr, err != nil, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccountUseCase_FindByUsername(t *testing.T) {
	type fields struct {
		accountRepo *account.MockIAccountRepository
	}
	type args struct {
		page     int
		username string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         *entity.Account
		wantErr      bool
	}{
		{
			name: "success find account by username",
			args: args{
				page:     1,
				username: "test",
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().FindByUsername(1, "test").Return(
					&entity.Account{
						ID:                 1,
						Username:           "test",
						RoleID:             5,
						Role:               entity.Role{ID: 5, Name: "test admin"},
						RegisterApprovalID: 5,
						RegisterApproval:   entity.RegisterApproval{ID: 5, AdminID: 5, SuperAdminID: 1, Status: "rejected"},
						Verified:           false,
						Activated:          false,
					}, nil)
			},
			want: &entity.Account{
				ID:                 1,
				Username:           "test",
				RoleID:             5,
				Role:               entity.Role{ID: 5, Name: "test admin"},
				RegisterApprovalID: 5,
				RegisterApproval:   entity.RegisterApproval{ID: 5, AdminID: 5, SuperAdminID: 1, Status: "rejected"},
				Verified:           false,
				Activated:          false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.accountRepo = account.NewMockIAccountRepository(t)
			uc := &AccountUseCase{
				accountRepo: tt.fields.accountRepo,
			}
			tt.prepareMocks(&tt.fields)
			got, err := uc.FindByUsername(tt.args.page, tt.args.username)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAccountUseCase_UpdateActivatedAccount(t *testing.T) {
	type fields struct {
		accountRepo *account.MockIAccountRepository
	}
	type args struct {
		ctx       context.Context
		id        uint
		activated bool
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		wantErr      bool
	}{
		{
			name: "success update activated account",
			args: args{
				ctx:       ctxJWTClaimWithSuperAdminRole(),
				id:        uint(5),
				activated: true,
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().UpdateActivatedAccount(uint(5), true).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.accountRepo = account.NewMockIAccountRepository(t)
			uc := &AccountUseCase{
				accountRepo: tt.fields.accountRepo,
			}
			tt.prepareMocks(&tt.fields)
			err := uc.UpdateActivatedAccount(tt.args.ctx, tt.args.id, tt.args.activated)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestAccountUseCase_Delete(t *testing.T) {
	type fields struct {
		accountRepo *account.MockIAccountRepository
	}
	type args struct {
		ctx context.Context
		id  uint
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		wantErr      bool
	}{
		{
			name: "success delete account",
			args: args{
				ctx: ctxJWTClaimWithSuperAdminRole(),
				id:  uint(5),
			},
			prepareMocks: func(f *fields) {
				f.accountRepo.EXPECT().Delete(uint(5)).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.accountRepo = account.NewMockIAccountRepository(t)
			uc := &AccountUseCase{
				accountRepo: tt.fields.accountRepo,
			}
			tt.prepareMocks(&tt.fields)
			err := uc.Delete(tt.args.ctx, tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
