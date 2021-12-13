package server

import (
	"fmt"

	"github.com/dineshkumar181094/key-value-store/pkg/store"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var globalStore = make(map[string]string)

func Start(port int) {
	r := gin.Default()
	r.GET("/get", getKeyHandler)
	r.POST("/set", setKeyHandler)
	r.PUT("/update", updateKeyHandler)

	r.Run(fmt.Sprintf("0.0.0.0:%v", port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getKeyHandler(c *gin.Context) {
	key := c.Query("key")
	log.Debugf("get request for key %v", key)
	value, err := store.Get(&globalStore, key)
	if err != nil {
		log.Debugf("key %v doesn't exists", key)
		c.String(400, err.Error())
	} else {
		log.Debugf("key %v value is %v", key, value)
		c.String(200, value)
	}
}

func setKeyHandler(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	log.Debugf("setting key %v with value %v", key, value)
	_, err := store.Set(&globalStore, key, value)
	if err != nil {
		c.String(400, err.Error())
	} else {
		c.String(200, "Ok")
	}
}

func updateKeyHandler(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	log.Debugf("update request for key %v with value %v", key, value)
	_, err := store.Update(&globalStore, key, value)
	if err != nil {
		c.String(400, err.Error())
	} else {
		c.String(200, "OK")
	}
}
