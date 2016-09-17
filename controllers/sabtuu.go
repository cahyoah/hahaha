package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

// operations for Sabtuu
type SabtuuController struct {
	beego.Controller
}

func (c *SabtuuController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Create
// @Description create Sabtuu
// @Param	body		body 	models.Sabtuu	true		"body for Sabtuu content"
// @Success 201 {object} models.Sabtuu
// @Failure 403 body is empty
// @router / [post]
func (c *SabtuuController) Post() {

}

// @Title GetOne
// @Description get Sabtuu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Sabtuu
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SabtuuController) GetOne() {

}

// @Title GetAll
// @Description get Sabtuu
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Sabtuu
// @Failure 403
// @router / [get]
func (c *SabtuuController) GetAll() {
	str, err := httplib.Get("http://httpbin.org/user-agent").String()
	
	c.Ctx.Output.SetStatus(200)
	c.Ctx.Output.Body([]byte(str))

	if err != nil {

	}
}

// @Title Update
// @Description update the Sabtuu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Sabtuu	true		"body for Sabtuu content"
// @Success 200 {object} models.Sabtuu
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SabtuuController) Put() {

}

// @Title Delete
// @Description delete the Sabtuu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SabtuuController) Delete() {

}
