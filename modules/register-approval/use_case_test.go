package register_approval

import (
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	register_approval "github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/register-approval"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterApprovalUseCase_FindAll(t *testing.T) {
	type fields struct {
		registerApprovalRepo *register_approval.MockIRegisterApprovalRepository
	}
	type args struct {
		page int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         []*entity.RegisterApproval
		wantErr      bool
	}{
		{
			name: "success find all pending approval",
			args: args{page: 1},
			prepareMocks: func(f *fields) {
				f.registerApprovalRepo.EXPECT().FindAll(1).Return([]*entity.RegisterApproval{
					{
						ID:           6,
						AdminID:      6,
						SuperAdminID: 1,
						Status:       "pending",
					},
					{
						ID:           7,
						AdminID:      7,
						SuperAdminID: 1,
						Status:       "pending",
					},
				}, nil)
			},
			want: []*entity.RegisterApproval{
				{
					ID:           6,
					AdminID:      6,
					SuperAdminID: 1,
					Status:       "pending",
				},
				{
					ID:           7,
					AdminID:      7,
					SuperAdminID: 1,
					Status:       "pending",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.registerApprovalRepo = register_approval.NewMockIRegisterApprovalRepository(t)
			uc := &RegisterApprovalUseCase{
				registerApprovalRepo: tt.fields.registerApprovalRepo,
			}

			tt.prepareMocks(&tt.fields)
			got, err := uc.FindAll(tt.args.page)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRegisterApprovalUseCase_UpdateApprovalStatus(t *testing.T) {
	type fields struct {
		registerApprovalRepo *register_approval.MockIRegisterApprovalRepository
	}
	type args struct {
		id  uint
		val string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		wantErr      bool
	}{
		{
			name: "success update approval status",
			args: args{
				id:  uint(1),
				val: "approved",
			},
			prepareMocks: func(f *fields) {
				f.registerApprovalRepo.EXPECT().UpdateApprovalStatus(uint(1), "approved").Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.registerApprovalRepo = register_approval.NewMockIRegisterApprovalRepository(t)
			uc := &RegisterApprovalUseCase{
				registerApprovalRepo: tt.fields.registerApprovalRepo,
			}
			tt.prepareMocks(&tt.fields)
			err := uc.UpdateApprovalStatus(tt.args.id, tt.args.val)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
