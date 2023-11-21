package http_response

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RestErr interface {
	Status() int
	Body() RestErrBody
}

type RestErrBody struct {
	ErrStatus int    `json:"code,omitempty"`
	ErrDetail string `json:"detail"`
}

func (e RestErrBody) Body() RestErrBody {
	return e
}

func (e RestErrBody) Status() int {
	return e.ErrStatus
}

func NewRestError(status int, detail string) RestErr {
	return RestErrBody{
		ErrStatus: status,
		ErrDetail: detail,
	}
}

func ParseError(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, "no records found in the database")
	case strings.Contains(strings.ToLower(err.Error()), "not found"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "duplicate key value violates unique constraint"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "already created"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "violates foreign key constraint"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "duplicated id"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "json: expected"):
		return NewRestError(http.StatusNotFound, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "invalid UUID"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "invalid syntax"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "field validation"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "invalid parameter"):
		return NewRestError(http.StatusBadRequest, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "unauthorized"):
		return NewRestError(http.StatusUnauthorized, err.Error())
	case strings.Contains(strings.ToLower(err.Error()), "unprocessable entity"):
		return NewRestError(http.StatusBadRequest, err.Error())
	default:
		return NewRestError(http.StatusInternalServerError, err.Error())
	}
}

func ErrorCtxResponse(ctx *fiber.Ctx, err error) error {
	restErr := ParseError(err)
	return ctx.Status(restErr.Status()).JSON(map[string]interface{}{
		"data":   []string{},
		"errors": restErr.Body(),
	})
}

func CtxResponse(ctx *fiber.Ctx, code int, data, pg interface{}) error {
	if code == 0 {
		code = fiber.StatusOK
	}
	return ctx.Status(code).JSON(map[string]interface{}{
		"data":       data,
		"errors":     nil,
		"pagination": pg,
	})
}
