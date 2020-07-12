package main

import (
	"net"
	"zroute.io/core/mmdb"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	pb "zroute.io/common/proto/gen/core/mmdbpb"
	"zroute.io/model/core"
)

func main() {
	db, err := gorm.Open("mysql", "root:123456@/model?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error("failed to connect database. err:", err)
		return
	}
	defer db.Close()
	db.SingularTable(true)
	db.LogMode(true)

	var tables []core.MemTable
	db.Find(&tables)
	for _, table := range tables {
		var columns []core.MemColumn
		db.Where("table_id = ?", table.ID).Find(&columns)
		//db.Model(&table).Related(&columns)
		mmdb.MMDB.CreateMemTable(table, columns)

		//add data to table
		var records []mmdb.MemDBRecord
		db.Table(table.Name).Scan(records)
		log.Debug("table records:", records)
		mmdb.MMDB.AddRecords(table.Name, records)
	}

	//start mmdb grpc service
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMemDBServer(s, &mmdb.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
