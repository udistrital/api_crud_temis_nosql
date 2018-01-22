package controllers

import (
	"api_crud_temis_nosql/db"
	"api_crud_temis_nosql/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud RolEntidadCuotasPartes
type RolEntidadCuotasPartesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.RolEntidadCuotasPartes
// @Failure 403 :objectId is empty
// @router / [get]
func (j *RolEntidadCuotasPartesController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllRolEntidadCuotasPartes(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Borrar RolEntidadCuotasPartes
// @Description Borrar RolEntidadCuotasPartes
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *RolEntidadCuotasPartesController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteRolEntidadCuotasPartesById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear RolEntidadCuotasPartes
// @Description Crear RolEntidadCuotasPartes
// @Param	body		body 	models.RolEntidadCuotasPartes	true		"Body para la creacion de RolEntidadCuotasPartes"
// @Success 200 {int} RolEntidadCuotasPartes.Id
// @Failure 403 body is empty
// @router / [post]
func (j *RolEntidadCuotasPartesController) Post() {
	var rolentidadcuotaspartes models.RolEntidadCuotasPartes
	json.Unmarshal(j.Ctx.Input.RequestBody, &rolentidadcuotaspartes)
	fmt.Println(rolentidadcuotaspartes)
	session, _ := db.GetSession()
	models.InsertRolEntidadCuotasPartes(session, rolentidadcuotaspartes)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the RolEntidadCuotasPartes
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *RolEntidadCuotasPartesController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var rolentidadcuotaspartes models.RolEntidadCuotasPartes
	json.Unmarshal(j.Ctx.Input.RequestBody, &rolentidadcuotaspartes)
	session, _ := db.GetSession()

	err := models.UpdateRolEntidadCuotasPartes(session, rolentidadcuotaspartes, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
