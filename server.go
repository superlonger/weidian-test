// todo.go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

type MsgRecv struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message"`
}

func main() {
	// setup http log file
	logHttpFile, err := os.OpenFile("log/http.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		panic(err)
	}
	defer logHttpFile.Close()
	// create a new instance of Echo
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{${time_rfc3339}: "ip":"${remote_ip}","host":"${host}",` +
			`"method":"${method}","uri":"${uri}","status":${status},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
		Output: logHttpFile,
	}))
	e.Static("/", "public")
	e.POST("/", recvMsg)
	e.Logger.Fatal(e.Start(":8000"))
}
func recvMsg(c echo.Context) error {
	var msgRecv MsgRecv
	content := c.FormValue("content")
	fmt.Println(content)
	err := json.Unmarshal([]byte(content), &msgRecv)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, `{"status":"success"}`)
}
