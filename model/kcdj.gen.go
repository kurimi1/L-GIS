// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKcdj = "kcdj"

// Kcdj mapped from table <kcdj>
type Kcdj struct {
	KCAAA  string `gorm:"column:KCAAA" json:"KCAAA"`
	KCC    string `gorm:"column:KCC" json:"KCC"`
	JJDAJ  string `gorm:"column:JJDAJ" json:"JJDAJ"`
	JJGLA  string `gorm:"column:JJGLA" json:"JJGLA"`
	DWAAC  string `gorm:"column:DWAAC" json:"DWAAC"`
	DWAAD  string `gorm:"column:DWAAD" json:"DWAAD"`
	KCBA   string `gorm:"column:KCBA" json:"KCBA"`
	KCCA   string `gorm:"column:KCCA" json:"KCCA"`
	KCCB   string `gorm:"column:KCCB" json:"KCCB"`
	PKGKB  string `gorm:"column:PKGKB" json:"PKGKB"`
	KCAOC  string `gorm:"column:KCAOC" json:"KCAOC"`
	KCAOAE string `gorm:"column:KCAOAE" json:"KCAOAE"`
	KCAOAF string `gorm:"column:KCAOAF" json:"KCAOAF"`
	PKD    string `gorm:"column:PKD" json:"PKD"`
	JJDCBF string `gorm:"column:JJDCBF" json:"JJDCBF"`
	数据     string `gorm:"column:数据" json:"数据"`
}

// TableName Kcdj's table name
func (*Kcdj) TableName() string {
	return TableNameKcdj
}
