package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleHttpErrors(err error, c *gin.Context) {
	var paramsErr *ApplicationError

	switch {
	case errors.As(err, &paramsErr):
		go LogErrors(paramsErr.ErrorDescription)

		c.Header("Content-Type", "application/json")

		if paramsErr.Code == CodeErrors["INVALID_PARAMS"] {
			splitStringError := strings.Split(paramsErr.ErrorDescription, "|")
			b := new(strings.Builder)
			json.NewEncoder(b).Encode(splitStringError)
			response := `{"error":"` + paramsErr.Code + `","message":"` + paramsErr.Error() + `","fields":` + b.String() + `}`

			c.Data(paramsErr.HttpError, "application/json", []byte(response))
			return
		}

		c.Data(paramsErr.HttpError, "application/json", []byte(`{"error":"`+paramsErr.Code+`", "message":"`+err.Error()+`"}`))
		return
	default:
		go LogErrors(err.Error())

		c.Header("Content-Type", "application/json")

		c.Data(http.StatusBadGateway, "application/json", []byte(`{"error":"BAD_GATEWAY", "message":"`+err.Error()+`"}`))
		return
	}
}

func LogErrors(errMessage string) {
	fmt.Println(errMessage)
}
