package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"onemore-service-go/models"
)

type VersionController struct {
	beego.Controller
}

func (c *VersionController) URLMapping() {
	c.Mapping("GetNew", c.GetNew)
	c.Mapping("Save", c.Save)
}

// @Title 获取最新版本
// @Description 获取最新版本
// @Param   type     		query    int8  	false        "类型：1、H5；2、PC；3、安卓；4、iOS；5、iPad"
// @Param   app_type     	query    int8  	false        "app类型：1课件编辑器，2绘本编辑器"
// @Success 0 {json} JSONStruct
// @Failure 1005 获取失败
// @router /getNew [get]
func (c *VersionController) GetNew() {
	types, _ := c.GetInt8("type", 3)
	appType, _ := c.GetInt8("app_type", 1)

	valid := validation.Validation{}
	valid.Range(types, 1, 5, "type").Message("类型只能是1~5的数字")
	valid.Range(appType, 1, 2, "app_type").Message("app类型只能是1~2的数字")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, "", valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}

	v := models.Version{Type: types, AppType: appType}
	if res, err := v.GetNew(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, res, "获取成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1005, "", "获取失败"}
	}

	c.ServeJSON()
}

// @Title 保存新版本
// @Description 保存新版本
// @Param   type     		formData    int8  		true         "类型：1、H5；2、PC；3、安卓；4、iOS；5、iPad"
// @Param   app_type     	formData    int8  		false        "app类型：1课件编辑器，2绘本编辑器"
// @Param   html_version    formData    string  	true         "版本号"
// @Param   html_url     	formData    string  	true         "url地址"
// @Success 0 {json} JSONStruct
// @Failure 1003 保存失败
// @router /save [post]
func (c *VersionController) Save() {
	types, _ := c.GetInt8("type")
	appType, _ := c.GetInt8("app_type", 1)
	htmlVersion := c.GetString("html_version")
	htmlUrl := c.GetString("html_url")

	valid := validation.Validation{}
	valid.Required(types, "type").Message("类型不能为空")
	valid.Range(types, 1, 5, "type").Message("类型只能是1~5的数字")
	valid.Range(appType, 1, 2, "app_type").Message("app类型只能是1~2的数字")
	valid.Required(htmlVersion, "html_version").Message("版本号不能为空")
	valid.Required(htmlUrl, "html_url").Message("地址不能为空")
	if valid.HasErrors() {
		c.Data["json"] = JSONStruct{"error", 1001, "", valid.Errors[0].Message}

		c.ServeJSON()
		c.StopRun()
	}

	v := models.Version{Type: types, AppType: appType, HtmlVersion: htmlVersion, HtmlUrl: htmlUrl}
	if err := v.Save(); err == nil {
		c.Data["json"] = JSONStruct{"success", 0, "", "保存成功"}
	} else {
		c.Data["json"] = JSONStruct{"error", 1003, "", "保存失败"}
	}

	c.ServeJSON()
}
