package customer

import (
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/entity"
	"github.com/mfajri11/Backend_1-Muhamad_Fajri-Mini_Project_2/repository/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCustomerUseCase_Create(t *testing.T) {
	type fields struct {
		customerRepo *customer.MockICustomerRepository
	}
	type args struct {
		customer CustomerParams
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		prepareMock func(f *fields)
		want        entity.Customer
		wantErr     bool
	}{
		{
			name: "success create customer",
			args: args{customer: CustomerParams{
				FirstName: "test",
				LastName:  "only",
				Email:     "test@example.com",
				Avatar:    "https://test/avatar1.png",
			}},
			prepareMock: func(f *fields) {
				f.customerRepo.EXPECT().Create(mock.Anything).RunAndReturn(func(e *entity.Customer) error {
					e.ID = 1
					return nil
				})
			},
			want: entity.Customer{
				ID:        1,
				FirstName: "test",
				LastName:  "only",
				Email:     "test@example.com",
				Avatar:    "https://test/avatar1.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.customerRepo = customer.NewMockICustomerRepository(t)
			tt.prepareMock(&tt.fields)
			uc := &CustomerUseCase{
				customerRepo: tt.fields.customerRepo,
			}

			got, err := uc.Create(tt.args.customer)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerUseCase_Update(t *testing.T) {
	type fields struct {
		customerRepo *customer.MockICustomerRepository
	}
	type args struct {
		customer CustomerUpdateParams
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         *entity.Customer
		wantErr      bool
	}{
		{
			name: "success update customer",
			args: args{customer: CustomerUpdateParams{
				ID:        1,
				FirstName: "updated",
				LastName:  "once again updated",
				Email:     "updated.again@example.com",
			}},
			prepareMocks: func(f *fields) {
				f.customerRepo.EXPECT().Update(mock.Anything).Return(&entity.Customer{
					ID:        1,
					FirstName: "updated",
					LastName:  "once again updated",
					Email:     "updated.again@example.com",
					Avatar:    "https://test/avatar1.png",
				}, nil)
			},
			want: &entity.Customer{
				ID:        1,
				FirstName: "updated",
				LastName:  "once again updated",
				Email:     "updated.again@example.com",
				Avatar:    "https://test/avatar1.png",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.customerRepo = customer.NewMockICustomerRepository(t)
			uc := &CustomerUseCase{
				customerRepo: tt.fields.customerRepo,
			}
			tt.prepareMocks(&tt.fields)

			got, err := uc.Update(tt.args.customer)

			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCustomerUseCase_Delete(t *testing.T) {
	type fields struct {
		customerRepo *customer.MockICustomerRepository
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		wantErr      bool
	}{
		{
			name: "success delete customer",
			args: args{id: 1},
			prepareMocks: func(f *fields) {
				f.customerRepo.EXPECT().Delete(uint(1)).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.customerRepo = customer.NewMockICustomerRepository(t)
			uc := &CustomerUseCase{
				customerRepo: tt.fields.customerRepo,
			}
			tt.prepareMocks(&tt.fields)

			err := uc.Delete(tt.args.id)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestCustomerUseCase_Search(t *testing.T) {
	type fields struct {
		customerRepo *customer.MockICustomerRepository
	}
	type args struct {
		page  int
		name  string
		email string
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		prepareMocks func(f *fields)
		want         []*entity.Customer
		wantErr      bool
	}{
		{
			name: "success search customer by email",
			args: args{
				page:  1,
				email: "test@example.com",
			},
			prepareMocks: func(f *fields) {
				f.customerRepo.
					EXPECT().
					Search(1, "", "test@example.com").
					Return([]*entity.Customer{
						{
							FirstName: "test",
							Email:     "test@example.com",
						},
					}, nil)
			},
			want: []*entity.Customer{
				{
					FirstName: "test",
					Email:     "test@example.com",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.customerRepo = customer.NewMockICustomerRepository(t)
			uc := &CustomerUseCase{
				customerRepo: tt.fields.customerRepo,
			}
			tt.prepareMocks(&tt.fields)

			got, err := uc.Search(tt.args.page, tt.args.name, tt.args.email)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.want, got)
		})
	}
}
