/*package main

import (
    "net/http"




    "github.com/labstack/echo/v4"
)

type Message struct {
    Text string `json:"text"`
}

func main() {
    e := echo.New()

    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
    })


    e.POST("/post", func(c echo.Context) error {
        // Parse JSON request body
        var message Message
        if err := c.Bind(&message); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
        }

        return c.JSON(http.StatusOK, map[string]string{"received": message.Text})
    })

    e.Logger.Fatal(e.Start(":1323"))
}
*/
/*
package main

import (

	"github.com/labstack/echo/v4"
	"net/http"
)
type Message struct{
    Text string `json :"text"`
}

func main()  {
    e:= echo.New()

    e.GET("/" ,func(c echo.Context) error{
        return c.JSON(http.StatusOK , map[string]string{"message":"hello rehan"})
    })


    e.POST("/post" , func(c echo.Context) error {
        var message Message
        if err :=c.Bind(&message); err!= nil{
            return c.JSON(http.StatusBadRequest ,map[string]string{"error":"invaled values or parameters"})
        }
        return c.JSON(http.StatusOK , map[string]string{"received":message.Text})

    })
    e.Logger.Fatal(e.Start(":1330"))
}
*/

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/word-count/cmd"
	"net/http"
	
)

type Message struct {
	FilePath string `json:"filepath"`
	Routines int    `json:"routines"`
}




func main() {
	e := echo.New()

	
	e.POST("/post", postData)
	// e.GET("/get", getData)

	
	e.Logger.Fatal(e.Start(":8080"))
}

func postData(c echo.Context) error {
	
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}
    

    totalCounts,routines,timetaken:=cmd.ProcessFile(message.FilePath, message.Routines)
    timeTakenString := timetaken.String()
    return c.JSON(http.StatusOK, map[string]interface{}{"counts": totalCounts, "routines": routines, "timetaken": timeTakenString})

	
}


