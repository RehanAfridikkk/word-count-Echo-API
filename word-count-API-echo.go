
package word

import (
	"github.com/labstack/echo/v4"
	"github.com/RehanAfridikkk/word-count-Echo-API/cmd"
	"net/http"
	
)

type Message struct {
	FilePath string `json:"filepath"`
	Routines int    `json:"routines"`
}



func word () {
	e := echo.New()

	
	e.POST("/post", postData)
	

	
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


