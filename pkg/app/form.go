package app

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type ValidationError struct {
	Message string
}

type ValidationErrors []*ValidationError

func (v *ValidationError) Error() string {
	return v.Message
}

func (vs ValidationErrors) Error() string {
	return strings.Join(vs.Errors(), ",")
}

func (vs ValidationErrors) Errors() []string {
	var errs []string
	for _, err := range vs {
		errs = append(errs, err.Error())
	}
	return errs
}

func BindAndValidate(c *gin.Context, v interface{}) (bool, ValidationErrors) {
	var errs ValidationErrors
	err := c.ShouldBind(v)
	if err != nil {
		errs = append(errs, &ValidationError{
			Message: err.Error(),
		})
		return false, errs
	}
	return true, nil
}
