package controllers

import (
	"encoding/json"
	"golang-webapi/models"
	"log"
	"os"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var (
	mongoConnStr = os.Getenv("MONGO_CONN_STR")
)

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)

	session, err := mgo.Dial(mongoConnStr)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("test").C("users")
	err = c.Insert(&user)
	if err != nil {
		log.Fatal(err)
		u.Data["json"] = err
	} else {
		u.Data["json"] = user
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {

	session, err := mgo.Dial(mongoConnStr)
	if err != nil {
		panic(err)
	}
	c := session.DB("test").C("users")
	var users []models.User
	queryError := c.Find(bson.M{}).All(&users)
	if queryError != nil {
		u.Data["json"] = queryError
	} else {
		u.Data["json"] = users
	}
	session.Close()
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		session, err := mgo.Dial(mongoConnStr)
		if err != nil {
			panic(err)
		}
		c := session.DB("test").C("users")
		var user models.User
		err = c.Find(bson.M{"id": uid}).One(&user)
		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = user
		}
		session.Close()
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	uid := u.GetString(":uid")
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if uid != "" {

		session, err := mgo.Dial(mongoConnStr)
		if err != nil {
			panic(err)
		}
		c := session.DB("test").C("users")
		err = c.Update(bson.M{"id": uid}, &user)

		if err != nil {
			u.Data["json"] = err
		} else {
			u.Data["json"] = user
		}
		session.Close()
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")

	if uid != "" {
		session, err := mgo.Dial(mongoConnStr)
		if err != nil {
			panic(err)
		}
		c := session.DB("test").C("users")
		err = c.Remove(bson.M{"id": uid})

		if err != nil {
			u.Data["json"] = "User not removed!"
		} else {
			u.Data["json"] = "User removed!"
		}
		session.Close()
	} else {
		u.Data["json"] = "User not removed!"
	}
	u.ServeJSON()
}
