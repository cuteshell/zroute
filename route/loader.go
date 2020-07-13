package route

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	pb "zroute.io/common/proto/gen/core/mmdbpb"
	"zroute.io/core/mmdb"
	"zroute.io/route/channel"
	"zroute.io/route/pnt"
	"zroute.io/route/rtu"
	"zroute.io/route/transport"
	_ "zroute.io/utils/log"
)

type Config struct {
}

type Project struct {
	Transport []transport.Transport
	Channels  map[string]channel.Channel
	Rtus      map[string]rtu.Rtu
	Pnts      map[string]pnt.Pnt
}

func LoadConfig() (*Config, error) {
	config := new(Config)

	return config, nil
}

func LoadProject() (*Project, error) {
	project := new(Project)

	client, err := mmdb.CreateMemDBClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	//Load daqchannel from memdb
	if err := project.LoadChannel(client); err != nil {
		return nil, err
	}
	//Load daqrtu from memdb
	if err := project.LoadRtu(client); err != nil {
		return nil, err
	}
	//Load daqpnt from memdb
	if err := project.LoadPnt(client); err != nil {
		return nil, err
	}

	return project, nil
}

func (p *Project) LoadChannel(client mmdb.Client) (err error) {
	records, err := client.GetRecords(client.Ctx, &pb.TableName{Name: "daqchannel"})
	if err != nil {
		return
	}
	for _, record := range records.GetRecords() {
		var rec mmdb.MemDBRecord
		data := record.GetData()
		if err := bson.Unmarshal(data, &rec); err != nil {
			return err
		}
		channel := channel.Channel{ID: rec["id"].(string), Name:rec["name"].(string), Address:rec["address"].(string)}
		p.Channels[rec["id"].(string)] = channel
	}
	return
}

func (p *Project) LoadRtu(client mmdb.Client) (err error) {
	records, err := client.GetRecords(client.Ctx, &pb.TableName{Name: "daqrtu"})
	if err != nil {
		return
	}
	for _, record := range records.GetRecords() {
		var rec mmdb.MemDBRecord
		data := record.GetData()
		if err := bson.Unmarshal(data, &rec); err != nil {
			return err
		}
		rtu := rtu.Rtu{ID: rec["id"].(string), Name: rec["name"].(string)}
		p.Rtus[rec["id"].(string)] = rtu
	}
	return
}

func (p *Project) LoadPnt(client mmdb.Client) (err error){
	records, err := client.GetRecords(client.Ctx, &pb.TableName{Name: "daqpnt"})
	if err != nil {
		return
	}
	for _, record := range records.GetRecords() {
		var rec mmdb.MemDBRecord
		data := record.GetData()
		if err := bson.Unmarshal(data, &rec); err != nil {
			return err
		}
		pnt := pnt.Pnt{ID: rec["id"].(string), Name: rec["name"].(string)}
		p.Pnts[rec["id"].(string)] = pnt
	}
	return
}
