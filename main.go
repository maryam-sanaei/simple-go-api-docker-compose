package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maryam-sanaei/simple-go-api-docker-compose/models"
)

func main() {

	router := gin.Default()

	router.GET("/", getPersons)
	router.POST("/addperson", addPerson)

	router.Run(":8080")
}

// Get all persons
// curl http://localhost:8080

func getPersons(c *gin.Context) {

	persons := models.GetPersons()

	if persons == nil || len(persons) == 0 {

		c.AbortWithStatus(http.StatusNotFound)

	} else {

		c.IndentedJSON(http.StatusOK, persons)

	}
}

// Add new person
/*
curl http://localhost:8080/addperson \
--include     --header "Content-Type: application/json" \
--request "POST" --data '{"firstname": "myfirstname","lastname": "mylastname"}'
*/

func addPerson(c *gin.Context) {

	var person models.Person

	if err := c.BindJSON(&person); err != nil {

		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		models.AddPerson(person)
		c.IndentedJSON(http.StatusCreated, person)
	}

}
