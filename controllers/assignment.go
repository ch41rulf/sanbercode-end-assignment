package controllers

import (
	"end-assignment/database"
	"end-assignment/repository"
	"end-assignment/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAssgnmnt(c *gin.Context) {
	var (
		result gin.H
	)

	assgnmnt, err := repository.GetAssgnmnt(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": assgnmnt,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertAssgnmnt(c *gin.Context) {
	var assgnmnt structs.Assignment

	err := c.ShouldBindJSON(&assgnmnt)
	if err != nil {
		panic(err)
	}

	err = repository.InsertsAssgnmnt(database.DbConnection, assgnmnt)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert ",
	})
}

func UpdateAssgnmnt(c *gin.Context) {
	var assgnmnt structs.Assignment
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&assgnmnt)
	if err != nil {
		panic(err)
	}

	assgnmnt.AssignmentID = int64(id)

	err = repository.UpdateAssgnmnt(database.DbConnection, assgnmnt)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update ",
	})
}

func DeleteAssgnmnt(c *gin.Context) {
	var assgnmnt structs.Assignment
	id, err := strconv.Atoi(c.Param("id"))

	assgnmnt.AssignmentID = int64(id)

	err = repository.DeleteAssgnmnt(database.DbConnection, assgnmnt)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete ",
	})
}
