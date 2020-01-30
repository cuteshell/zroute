package model

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "zroute.io/utils/log"
)

type Link struct {
	ID        uint `gorm:"primary_key"`
	TransType string
	IP        string
	Port      uint16
	ChannelID uint
}

type Channel struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	Desc     string
	LinkList []Link
	RtuList  []Rtu
}

type Parameter struct {
	ID    uint `gorm:"primary_key"`
	Name  string
	Value string
	RtuID uint
}

type Pnt struct {
	ID       uint `gorm:"primary_key"`
	Type     uint8
	IDBegin  uint32
	Num      uint32
	Function int32
	Address  int32
	RtuID    uint32
}

type Rtu struct {
	ID          uint `gorm:"primary_key"`
	Name        string
	Desc        string
	ProtoName   string
	Type        uint32
	DataAddress int32
	Dimaxnum    uint32
	Aimaxnum    uint32
	Pimaxnum    uint32
	Domaxnum    uint32
	Aomaxnum    uint32
	ChannelID   uint
	Parameter   []Parameter
	PntList     []Pnt
	Channel     Channel
}

type RouteTable struct {
	ID      uint `gorm:"primary_key"`
	SrcPath string
	DstPath string
	Type    uint32
}

type Model struct {
	*gorm.DB
}

func New() *Model {
	m := &Model{}
	dataDir := "data"
	os.MkdirAll(dataDir, 0755)
	db, err := gorm.Open("sqlite3", dataDir+"/model.db")
	if err != nil {
		log.Error("failed to connect database. err:", err)
		return nil
	}
	db.SingularTable(true)
	m.DB = db
	err = m.AutoMigrate(&Link{}, &Channel{}, &Pnt{}, &Parameter{}, &Rtu{}, &RouteTable{}).Error
	if err != nil {
		log.Error("failed to migrate. err:", err)
		m.Close()
		return nil
	}

	return m
}

func (m *Model) Close() {
	if m == nil || m.DB == nil {
		log.Warn("m.DB is nil")
		return
	}
	err := m.DB.Close()
	if err != nil {
		log.Error("Cann't close the model database, err:", err)
	}
}

func (m *Model) GetAllChannels(channels *Channel) *gorm.DB {
	m.Find(&channels)
	return m.DB
}

func (m *Model) GetAllRouteTables(routes *RouteTable) *gorm.DB {
	m.Find(&routes)
	return m.DB
}
