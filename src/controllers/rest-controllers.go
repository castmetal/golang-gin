package controllers

import (
	"github.com/gin-gonic/gin"

	_rest_v1 "golang-gin/src/controllers/rest/v1"
)

func SetRestControllers(routerEngine *gin.Engine) {
	_rest_v1.SetUserControllers(routerEngine)
}
