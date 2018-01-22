package controllers

import (
	"api_crud_temis_nosql/db"
	"api_crud_temis_nosql/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud EntidadCuotasPartes
type EntidadCuotasPartesController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.EntidadCuotasPartes
// @Failure 403 :objectId is empty
// @router / [get]
func (j *EntidadCuotasPartesController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllEntidadCuotasPartes(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Borrar EntidadCuotasPartes
// @Description Borrar EntidadCuotasPartes
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *EntidadCuotasPartesController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteEntidadCuotasPartesById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear EntidadCuotasPartes
// @Description Crear EntidadCuotasPartes
// @Param	body		body 	models.EntidadCuotasPartes	true		"Body para la creacion de EntidadCuotasPartes"
// @Success 200 {int} EntidadCuotasPartes.Id
// @Failure 403 body is empty
// @router / [post]
func (j *EntidadCuotasPartesController) Post() {
	var entidadcuotaspartes models.EntidadCuotasPartes
	json.Unmarshal(j.Ctx.Input.RequestBody, &entidadcuotaspartes)
	fmt.Println(entidadcuotaspartes)
	session, _ := db.GetSession()
	models.InsertEntidadCuotasPartes(session, entidadcuotaspartes)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the EntidadCuotasPartes
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *EntidadCuotasPartesController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var entidadcuotaspartes models.EntidadCuotasPartes
	json.Unmarshal(j.Ctx.Input.RequestBody, &entidadcuotaspartes)
	session, _ := db.GetSession()

	err := models.UpdateEntidadCuotasPartes(session, entidadcuotaspartes, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
