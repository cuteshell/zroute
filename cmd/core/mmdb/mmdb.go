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
		records, err := ReadTable(db, table.Name)
		if err != nil {
			log.Error("database error:", err)
			continue
		}
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

func ReadTable(db *gorm.DB, name string) (records []mmdb.MemDBRecord, err error) {
	rows, err := db.DB().Query("select * from " + name)
	if err != nil {
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		return
	}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err = rows.Scan(columnPointers...); err != nil {
			return
		}
		record := make(mmdb.MemDBRecord)
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			record[colName] = *val
		}
		records = append(records, record)
	}
	return
}
