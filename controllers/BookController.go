package controllers

import (
	"github.com/gin-gonic/gin"
	"wahaha/base"
	models "wahaha/module/gin"
	"net/http"
	"wahaha/module"
)

func List(context *gin.Context) {
	bk := make([]module.Book, 10)
	if err := models.Model.Unscoped().Where("status = 0").Find(&bk).Error; err != nil {
		panic(err)
	}
	result := base.ReturnCode(http.StatusOK, "", bk)
	context.JSON(result.Code, &result)
}
