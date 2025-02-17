package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	e "github.com/prolgrammer/BM_package/errors"
	"net/http"
)

func (m middleware) HandleErrors(c *gin.Context) {
	if len(c.Errors) > 0 {
		err := c.Errors.Last()

		fmt.Println("an error has happened")
		if errors.Is(err, e.ErrDataBindError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, e.ErrRegistrationsError) {
			c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, e.ErrAuthRequiredError) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		//repositories
		if errors.Is(err, e.ErrEntityNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
			return
		}

		//useCases
		if errors.Is(err, e.ErrEntityAlreadyExists) {
			c.AbortWithStatusJSON(http.StatusConflict, err.Error())
			return
		}

		if errors.Is(err, e.ErrPasswordMismatch) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		fmt.Println("Unexpected error")
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
