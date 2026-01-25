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

func getIntFromQuery(paramName string, c *gin.Context) (int64, error) {
	paramValue := c.Query(paramName)
	slog.Info("Starting GetTeamWithPlayers", "teamIdStr", paramValue)

	if paramValue == "" {
		return 0, fmt.Errorf("%w: parameter '%s' is empty", ErrParamEmpty, paramName)
	}

	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: parameter '%s' value '%s' is not a valid integer.", ErrParamParse, paramName, paramValue)
	}

	return intValue, nil
}
