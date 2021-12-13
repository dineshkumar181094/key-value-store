package server

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRegisterUnauthenticated(t *testing.T) {
	rPath := "/get"
	router := gin.Default()
	router.GET(rPath, GET)
	w := httptest.NewRequest()
}
