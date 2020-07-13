package mmdb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2/bson"
	"time"
	pb "zroute.io/common/proto/gen/core/mmdbpb"
	_ "zroute.io/utils/log"
)

type Server struct {
	pb.UnimplementedMemDBServer
}

func convertToPBRecords(records []MemDBRecord) (*pb.Records, error) {
	var recs []*pb.Record
	for _, record := range records {
		id, ok := record["ID"].(string)
		if !ok {
			log.Warn("record(%v) id is not string", record)
			continue
		}
		if data, err := bson.Marshal(record); err != nil {
			recs = append(recs, &pb.Record{Id: id, Data: data})
			return nil, err
		}
	}
	return &pb.Records{Records: recs}, nil
}

func (s *Server) GetRecords(ctx context.Context, in *pb.TableName) (*pb.Records, error) {
	log.Debugf("Get table:%s records\n", in.GetName())
	records := MMDB.GetRecords(in.GetName())

	return convertToPBRecords(records)
}

func (s *Server) GetRecordsByField(ctx context.Context, in *pb.RequestCondition) (*pb.Records, error) {
	log.Debugf("Get Table:%s records by %s = %v\n", in.GetTableName(), in.GetFieldName(), in.GetFieldValue())
	records := MMDB.GetRecordsByField(in.GetTableName(), in.GetFieldName(), in.GetFieldValue())

	return convertToPBRecords(records)
}

type Client struct {
	pb.MemDBClient
	Ctx context.Context
	conn *grpc.ClientConn
	cancel context.CancelFunc
}

func CreateMemDBClient() (client Client, err error){
	address := "localhost:1234"
	client.conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	client.MemDBClient = pb.NewMemDBClient(client.conn)
	client.Ctx, client.cancel = context.WithTimeout(context.Background(), time.Second)
	return client, err
}

func (c *Client) Close() {
	c.conn.Close()
	c.cancel()
}
