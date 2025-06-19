package models

type Contact struct {
	Id     int64  `gorm:"primaryKey" json:"id"`
	Nama   string `gorm:"type:varchar(50)" json:"nama"`
	Alamat string `gorm:"type:varchar(100)" json:"alamat"`
	NoHp   int64  `gorm:"type:bigint" json:"no_hp"`
}
