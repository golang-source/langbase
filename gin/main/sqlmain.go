package main

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"net/http"
)

type Person struct {
	Id   int
	Name string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func FetchSingleUser(c *gin.Context) {
	id := c.Param("id")

	db, err := sql.Open("mysql", "root:wangshubo@/test?charset=utf8")
	checkErr(err)

	defer db.Close()

	err = db.Ping()
	checkErr(err)

	var (
		person Person
		result gin.H
	)
	row := db.QueryRow("select id, name from user_info where id = ?;", id)
	err = row.Scan(&person.Id, &person.Name)
	if err != nil {
		// If no results send null
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}
