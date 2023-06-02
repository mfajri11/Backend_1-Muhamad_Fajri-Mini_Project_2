package entity

import (
	"database/sql/driver"
	"strconv"
	"time"
)

// gorm prefer convention over configuration https://gorm.io/docs/models.html#Conventions

type VerifiedType bool
type ActivatedType = VerifiedType

func (vt *VerifiedType) Scan(value interface{}) error {
	var err error
	if vt == nil {
		*vt = VerifiedType(false)
		return nil
	}
	sv, err := driver.String.ConvertValue(value)
	if err != nil {
		return err
	}
	vs := sv.([]byte)
	bv, err := strconv.ParseBool(string(vs))
	if err != nil {
		return err
	}
	*vt = VerifiedType(bv)
	return nil
}

func (v VerifiedType) Value() (driver.Value, error) {
	return strconv.FormatBool(bool(v)), nil
}

type Account struct {
	ID        uint
	Username  string
	Password  string
	RoleID    uint
	Role      Role `gorm:"constraint:OnDelete:CASCADE"`
	Verified  VerifiedType
	Activated ActivatedType
	CreatedAt time.Time
	UpdatedAt time.Time
}
