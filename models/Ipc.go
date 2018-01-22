package models

import (
	"api_crud_temis_nosql/db"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const IpcCollection = "ipc"

type Ipc struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Vigencia string        `json:"vigencia"`
	Valor    float32       `json:"valor"`
}

func UpdateIpc(session *mgo.Session, j Ipc, id string) error {
	c := db.Cursor(session, IpcCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertIpc(session *mgo.Session, j Ipc) {
	c := db.Cursor(session, IpcCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllIpcs(session *mgo.Session) []Ipc {
	c := db.Cursor(session, IpcCollection)
	defer session.Close()
	fmt.Println("Getting all ipcs")
	var ipcs []Ipc
	err := c.Find(bson.M{}).All(&ipcs)
	if err != nil {
		fmt.Println(err)
	}
	return ipcs
}

/*func GetIpcByName(session *mgo.Session, name string) ([]Ipc, error) {
	c := db.Cursor(session, IpcCollection)
	defer session.Close()
	var ipcs []Ipc
	err := c.Find(bson.M{"nombre": name}).All(&ipcs)
	return ipcs, err
}*/

func DeleteIpcById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, IpcCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
