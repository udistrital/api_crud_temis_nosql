package controllers

import (
	"api_crud_temis_nosql/db"
	"api_crud_temis_nosql/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	_ "gopkg.in/mgo.v2"
)

// Operaciones Crud Ipc
type IpcController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Ipc
// @Failure 403 :objectId is empty
// @router / [get]
func (j *IpcController) GetAll() {
	session, _ := db.GetSession()
	obs := models.GetAllIpcs(session)
	fmt.Println(obs)
	j.Data["json"] = &obs
	j.ServeJSON()
}

// @Title Borrar Ipc
// @Description Borrar Ipc
// @Param	objectId		path 	string	true		"El ObjectId del objeto que se quiere borrar"
// @Success 200 {string} ok
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (j *IpcController) Delete() {
	session, _ := db.GetSession()
	objectId := j.Ctx.Input.Param(":objectId")
	result, _ := models.DeleteIpcById(session, objectId)
	j.Data["json"] = result
	j.ServeJSON()
}

// @Title Crear Ipc
// @Description Crear Ipc
// @Param	body		body 	models.Ipc	true		"Body para la creacion de Ipc"
// @Success 200 {int} Ipc.Id
// @Failure 403 body is empty
// @router / [post]
func (j *IpcController) Post() {
	var ipc models.Ipc
	json.Unmarshal(j.Ctx.Input.RequestBody, &ipc)
	fmt.Println(ipc)
	session, _ := db.GetSession()
	models.InsertIpc(session, ipc)
	j.Data["json"] = "insert success!"
	j.ServeJSON()
}

// @Title Update
// @Description update the Ipc
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (j *IpcController) Put() {
	objectId := j.Ctx.Input.Param(":objectId")

	var ipc models.Ipc
	json.Unmarshal(j.Ctx.Input.RequestBody, &ipc)
	session, _ := db.GetSession()

	err := models.UpdateIpc(session, ipc, objectId)
	if err != nil {
		j.Data["json"] = err.Error()
	} else {
		j.Data["json"] = "update success!"
	}
	j.ServeJSON()
}
