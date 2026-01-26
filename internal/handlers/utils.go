package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrParamEmpty = errors.New("query parameter is empty")
	ErrParamParse = errors.New("failed to parse query parameter")
)

func getIntQueryParam(paramName string, c *gin.Context) (int64, error) {
	stringValue := c.Query(paramName)
	slog.Info("Starting GetTeamWithPlayers", "teamIdStr", stringValue)

	if stringValue == "" {
		return 0, fmt.Errorf("%w: parameter '%s' is empty", ErrParamEmpty, paramName)
	}

	value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: parameter '%s' value '%s' is not a valid integer.", ErrParamParse, paramName, stringValue)
	}

	return value, nil
}

func getIntPathParam(paramName string, c *gin.Context) (int64, error) {
	stringValue := c.Param(paramName)

	if stringValue == "" {
		return 0, fmt.Errorf("%w: parameter '%s' is empty", ErrParamEmpty, paramName)
	}
	value, err := strconv.ParseInt(stringValue, 10, 64)

	if err != nil {
		return 0, fmt.Errorf("%w: parameter '%s' value '%s' is not a valid integer.",
			ErrParamParse, paramName, stringValue)
	}

	return value, nil
}
