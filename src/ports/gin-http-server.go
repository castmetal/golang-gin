package ports

import (
	"log"
	"net/http"
	"time"

	_config "golang-gin/src/config"
	_controllers "golang-gin/src/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateGinServer() {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	router.GET("/healthCheck", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"alive": true,
		})
	})

	_controllers.SetRestControllers(router)

	router.Use(cors.New(config))

	if _config.SystemParams.ENV != "staging" {
		_ = router.SetTrustedProxies([]string{})
		router.RemoteIPHeaders = []string{"X-Forwarded-For"}

		gin.SetMode(gin.ReleaseMode)
	}

	s := &http.Server{
		Addr:           ":" + _config.SystemParams.PORT,
		Handler:        router,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.SetFlags(0)
	log.Println("\n\nServer listen at port :8000")

	log.Fatal(s.ListenAndServe())

}
