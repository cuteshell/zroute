package mmdb

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"zroute.io/model/core"
)

type MemDBRecord map[string]interface{}

type MemDBTable struct {
	Meta map[string]interface{}
	Data map[string]MemDBRecord
}

type MemDB map[string]MemDBTable

var MMDB = make(MemDB)

func (m MemDB) CreateMemTable(table core.MemTable, columns []core.MemColumn) {
	meta := make(map[string]interface{})
	for _, column := range columns {
		meta[column.Name] = column.Type
	}

	m[table.Name] = MemDBTable{Meta: meta, Data: make(map[string]MemDBRecord)}
}

func (m MemDB) AddRecords(tableName string, records []MemDBRecord) {
	for _, record := range records {
		if value, ok := record["ID"]; ok {
			if id, ok := value.(string); ok {
				m[tableName].Data[id] = record
				continue
			}
			log.Warn("record ID type is not string. record:", record)
			continue
		}
		log.Warn("reocord ID field was not found. record:", record)
	}
}

func (m MemDB) DelRecords(tableName string, records []MemDBRecord) {
	for _, record := range records {
		if value, ok := record["ID"]; ok {
			if id, ok := value.(string); ok {
				delete(m[tableName].Data, id)
				break
			}
		}
	}
}

func (m MemDB) GetRecords(tableName string) (records []MemDBRecord) {
	for _, record := range m[tableName].Data {
		records = append(records, record)
	}
	return records
}

func (m MemDB) GetRecordByID(tableName string, id string) (Record MemDBRecord) {
	if record, ok := m[tableName].Data[id]; ok {
		return record
	}
	return nil
}

func (m MemDB) UpdateRecordField(tableName string, field string, id string, value interface{}) {
	if _, ok := m[tableName].Data[id][field]; !ok {
		log.Warn("record was not found. id:")
	}
	m[tableName].Data[id][field] = value
}
