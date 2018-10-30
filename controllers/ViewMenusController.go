package controllers

import (
	"github.com/gin-gonic/gin"
	models "wahaha/module/gin"
	"wahaha/module"
	"net/http"
	"wahaha/base"
)

/**
首页顶部菜单
 */
func FindIndexMenus(context *gin.Context) {
	m := make([]module.Menus, 10)
	if err := models.Model.Unscoped().Find(&m).Error; err != nil {
		panic(err)
	}
	code := base.ReturnCode(http.StatusOK, "", m)
	context.JSON(http.StatusOK, &code)
}
