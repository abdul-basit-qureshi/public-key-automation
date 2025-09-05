package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// func ReadFileData(filepath string) ([]byte, error) {
// 	return os.ReadFile(filepath)
// }

var i, j = 0, 0

// curl http://localhost:1122
func hello(c echo.Context) error {
	i = i + 1
	var currentTime = time.Now()
	fmt.Printf("\n%d Received %s request from %s at %s\n\n", i, c.Request().Method, c.RealIP(), currentTime.Format(time.ANSIC))
	return c.String(http.StatusOK, "Welcome to Automation Server\n")
}

// curl -X POST http://localhost:1122/key
func key(c echo.Context) error {
	j = j + 1
	var currentTime = time.Now()
	fmt.Printf("\n%d\nReceived %s request from %s at %s\n", j, c.Request().Method, c.RealIP(), currentTime.Format(time.ANSIC))
	fmt.Print("Response: ", c.Response())
	// fileContent, err := ReadFileData("/root/practice/repositories/freeswitch-golang-http-server/fs-go/dp.xml")
	// if err != nil {
	// 	fmt.Println("Error reading file: ", err)
	// 	return c.String(http.StatusInternalServerError, "Error while reading file")
	// }

	// fmt.Print("Response:\n", updatedXML)

	return c.String(http.StatusOK, "Welcome to the Public Endpoint\n")
}

func main() {
	e := echo.New()
	e.GET("/", hello)
	e.POST("/public", key)

	e.Logger.Fatal(e.Start("0.0.0.0:1122"))
}
