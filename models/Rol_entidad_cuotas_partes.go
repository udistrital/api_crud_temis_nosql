package models

import (
	"api_crud_temis_nosql/db"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const RolEntidadCuotasPartesCollection = "rolentidadcuotaspartes"

type RolEntidadCuotasPartes struct {
	Id       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Nombre string        `json:"nombre"`
	Descripcion    string       `json:"descripcion"`
}

func UpdateRolEntidadCuotasPartes(session *mgo.Session, j RolEntidadCuotasPartes, id string) error {
	c := db.Cursor(session, RolEntidadCuotasPartesCollection)
	defer session.Close()
	// Update
	err := c.Update(bson.M{"_id": bson.ObjectIdHex(id)}, &j)
	if err != nil {
		panic(err)
	}
	return err

}

func InsertRolEntidadCuotasPartes(session *mgo.Session, j RolEntidadCuotasPartes) {
	c := db.Cursor(session, RolEntidadCuotasPartesCollection)
	defer session.Close()
	c.Insert(j)

}

func GetAllRolEntidadCuotasPartes(session *mgo.Session) []RolEntidadCuotasPartes {
	c := db.Cursor(session, RolEntidadCuotasPartesCollection)
	defer session.Close()
	fmt.Println("Getting all rolentidadcuotaspartes")
	var rolentidadcuotaspartes []RolEntidadCuotasPartes
	err := c.Find(bson.M{}).All(&rolentidadcuotaspartes)
	if err != nil {
		fmt.Println(err)
	}
	return rolentidadcuotaspartes
}

/*func GetRolEntidadCuotasPartesByName(session *mgo.Session, name string) ([]RolEntidadCuotasPartes, error) {
	c := db.Cursor(session, RolEntidadCuotasPartesCollection)
	defer session.Close()
	var rolentidadcuotaspartess []RolEntidadCuotasPartes
	err := c.Find(bson.M{"nombre": name}).All(&rolentidadcuotaspartess)
	return rolentidadcuotaspartess, err
}*/

func DeleteRolEntidadCuotasPartesById(session *mgo.Session, id string) (string, error) {
	c := db.Cursor(session, RolEntidadCuotasPartesCollection)
	defer session.Close()
	err := c.RemoveId(bson.ObjectIdHex(id))
	return "ok", err
}
