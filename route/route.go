package route

import (
	"context"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"zroute.io/core/mmdb"

	"github.com/coreos/etcd/store"
	pb "zroute.io/common/proto/gen/core/mmdbpb"
	"zroute.io/model"
	"zroute.io/route/channel"
	_ "zroute.io/utils/log"
)

type Route struct {
	ctx     context.Context
	running bool
	mmdb    mmdb.Client
	store   store.Store
	model   *model.Model
	chans   []*channel.Channel
	rtype   uint8
}

func New(ctx context.Context, project *Project) *Route {
	r := &Route{}
	r.model = model.New()

	var err error
	r.mmdb, err = mmdb.CreateMemDBClient()
	if err != nil {
		log.Fatal("create MemDBClient:", err)
	}
	//Get daqchannel record from mmdb
	records, err := r.mmdb.GetRecords(r.mmdb.Ctx, &pb.TableName{Name: "daqchannel"})
	if err != nil {
		log.Fatal("GetRecords:", err)
	}

	for _, record := range records.GetRecords()  {
		var rec mmdb.MemDBRecord
		data := record.GetData()
		if err := bson.Unmarshal(data, &rec); err != nil {
			log.Error(err)
			continue
		}
		channel := channel.Channel{Name:rec["Name"].(string),Address:rec["Address"].(string)}
		r.chans = append(r.chans, &channel)
	}
	return r
}

/*
	Ctx   context.Context
	Trans transport.Transport
	Rtus  map[int]rtu.Rtu // key is rtu_id
	Sendc chan []byte
*/
func (r *Route) Start() error {
	for _, c := range r.chans {
		c.Start()
	}
	return nil
}

func (r *Route) Close() error {
	r.model.Close()
	for _, c := range r.chans {
		c.Close()
	}
	return nil
}
