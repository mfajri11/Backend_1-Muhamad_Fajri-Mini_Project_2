package entity

type RegisterApproval struct {
	ID           uint
	AdminID      uint
	SuperAdminID uint
	//Account      `gorm:"foreignKey:AdminID,references:ID"`
	Status string
}
