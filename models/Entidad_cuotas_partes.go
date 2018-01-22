package models

import (
	"api_crud_temis_nosql/db"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const EntidadCuotasPartesCollection = "entidadcuotaspartes"

type EntidadCuotasPartes struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Id_entidad int32        `json:"idEntidad"`
	Rol    RolEntidadCuotasPartes `json:"rolEntidad"`
	Sucesor bson.ObjectId `json:"sucesor"`
}

func UpdateEntidadCuotasPartes(session *mgo.Session, j EntidadCuotasPartes, id string) error {
	c := db.Cursor(session, EntidadCuotasPartesCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertEntidadCuotasPartes(session *mgo.Session, j EntidadCuotasPartes) {
	c := db.Cursor(session, EntidadCuotasPartesCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllEntidadCuotasPartes(session *mgo.Session) []EntidadCuotasPartes {
	c := db.Cursor(session, EntidadCuotasPartesCollection)
	defer session.Close()
	fmt.Println("Getting all entidadcuotaspartes")
	var entidadcuotaspartes []EntidadCuotasPartes
	err := c.Find(bson.M{}).All(&entidadcuotaspartes)
	if err != nil {
		fmt.Println(err)
	}
	return entidadcuotaspartes
}

/*func GetEntidadCuotasPartesByName(session *mgo.Session, name string) ([]EntidadCuotasPartes, error) {
	c := db.Cursor(session, EntidadCuotasPartesCollection)
	defer session.Close()
	var entidadcuotaspartess []EntidadCuotasPartes
	err := c.Find(bson.M{"nombre": name}).All(&entidadcuotaspartess)
	return entidadcuotaspartess, err
}*/

func DeleteEntidadCuotasPartesById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, EntidadCuotasPartesCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
