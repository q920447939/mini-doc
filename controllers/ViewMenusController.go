package controllers

import (
	"github.com/gin-gonic/gin"
	"wahaha/models"
	"wahaha/module"
	"net/http"
	"wahaha/base"
)

/**
首页顶部菜单
 */
func FindIndexMenus(context *gin.Context) {
	 result := base.BaseReturnJson{}
	m := make([]module.Menus, 10)
	if err := models.Model.Unscoped().Find(&m).Error; err != nil {
		panic(err)
	}
	result.Code = http.StatusOK
	result.Data = m
	code := base.ReturnCode(http.StatusOK, "", m)
	context.JSON(http.StatusOK, &code)
}
