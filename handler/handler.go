package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func ErrRouter(c *gin.Context) {
	c.String(http.StatusBadRequest, "url err")
}

func NIP05(c *gin.Context) {
	name2pubkey := map[string]string{
		"Gordon": "19ff8c27b0ba49a9d28703a44b439f1f23d9ccd97fbfa84b02a127bbefd13fec",
       		// 可以在这里添加更多的账号，为你的朋友提供验证
      		// "<name1>":"pubkey1",
      		// "<name2>":"pubkey2",
	}
	user := c.Query("name")
	fmt.Println("nip05 verify request", user, name2pubkey[user])
	if v, ok := name2pubkey[user]; ok {
		resp := NIP05Resp{Names: map[string]string{}}
		resp.Names[user] = v
		c.JSON(http.StatusOK, resp)
		return
	}
	c.Status(http.StatusNotFound)
	return
}

type NIP05Resp struct {
	Names map[string]string `json:"names"`
}

func Cors(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Next()
}
