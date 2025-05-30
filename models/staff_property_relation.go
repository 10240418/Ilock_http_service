package models

import "time"

// StaffDeviceRelation 表示物业员工和设备之间的多对多关系
type StaffDeviceRelation struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	StaffID   uint      `gorm:"not null" json:"staff_id"`     // 物业员工ID
	DeviceID  uint      `gorm:"not null" json:"device_id"`    // 设备ID
	Role      string    `gorm:"type:varchar(50)" json:"role"` // 如：manager, maintainer, etc.
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联
	Staff  *PropertyStaff `gorm:"foreignKey:StaffID" json:"staff,omitempty"`
	Device *Device        `gorm:"foreignKey:DeviceID" json:"device,omitempty"`
}
