package mmdb

import "github.com/hashicorp/go-memdb"

type Device struct {
	ID   uint32
	Name string
}

type Point struct {
	ID     uint32
	Name   string
	Driver uint32
}

type Driver struct {
	ID   uint32
	Name string
}

func NewMemDB() (*memdb.MemDB, error) {

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"device": &memdb.TableSchema{
				Name: "device",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
			"point": &memdb.TableSchema{
				Name: "point",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
					"driver": &memdb.IndexSchema{
						Name:    "driver",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Driver"},
					},
				},
			},
			"driver": &memdb.TableSchema{
				Name: "driver",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UintFieldIndex{Field: "ID"},
					},
					"name": &memdb.IndexSchema{
						Name:    "name",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Name"},
					},
				},
			},
		},
	}
	return memdb.NewMemDB(schema)
}

