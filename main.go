package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TxtxReq struct {
	Text string       `json:"text"`
	Op   []*Operation `json:"op"`
}
type Operation struct {
	Type     string `json:"type"`
	OpString string `json:"command"`
}
type OpType uint

const (
	OpMatch OpType = iota
	OpDelete
	OpAppend
	OpPrepend
	OpReplace
	OpLine
)

func main() {
	fmt.Println("hello")

	r := gin.Default()

	r.Use(CORSMiddleware())
	r.POST("/op", txtx)
	r.GET("/greet/:name", greet)

	r.Run(":1212")

}

func txtx(c *gin.Context) {
	var reqData, _ = io.ReadAll(c.Request.Body)
	var txReq TxtxReq
	json.Unmarshal(reqData, &txReq)
	out := txReq.Text
	for _, op := range txReq.Op {
		switch op.Type {
		case "replace":
			opArg := strings.Split(op.OpString, "/")
			out = replace(out, opArg[0], opArg[1])
		case "delete":
			out = replace(out, op.OpString, "")
		default:
			c.String(http.StatusBadRequest, fmt.Sprint("Invalid operation type ", op.Type))
		}

	}

	c.String(http.StatusOK, out)

}

func greet(c *gin.Context) {
	greeting := fmt.Sprintf("Hey dipshit, your name is %s? What a bonehead name!\n", c.Param("name"))
	c.String(http.StatusOK, greeting)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
