package core

type MemTable struct {
	ID          int `gorm:"primary_key"`
	Name        string
	description string
	Columns     []MemColumn `gorm:"foreignkey:TableID"`
}

type MemColumn struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Description string
	TableID     int
	Type        string
}
