// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKqdz = "kqdz"

// Kqdz mapped from table <kqdz>
type Kqdz struct {
	KCAAA  string  `gorm:"column:KCAAA" json:"KCAAA"`
	PKGKAA string  `gorm:"column:PKGKAA" json:"PKGKAA"`
	DSB    string  `gorm:"column:DSB" json:"DSB"`
	YSACAA string  `gorm:"column:YSACAA" json:"YSACAA"`
	YSACAB string  `gorm:"column:YSACAB" json:"YSACAB"`
	YSACC  string  `gorm:"column:YSACC" json:"YSACC"`
	GZAM   string  `gorm:"column:GZAM" json:"GZAM"`
	KCAJ   string  `gorm:"column:KCAJ" json:"KCAJ"`
	DHBIA  string  `gorm:"column:DHBIA" json:"DHBIA"`
	DHBHAB float32 `gorm:"column:DHBHAB" json:"DHBHAB"`
	PKHFE  string  `gorm:"column:PKHFE" json:"PKHFE"`
	PKHFB  string  `gorm:"column:PKHFB" json:"PKHFB"`
	DHCAA  string  `gorm:"column:DHCAA" json:"DHCAA"`
	DHC    string  `gorm:"column:DHC" json:"DHC"`
	数据     string  `gorm:"column:数据" json:"数据"`
}

// TableName Kqdz's table name
func (*Kqdz) TableName() string {
	return TableNameKqdz
}
