package controllers

import (
	"end-assignment/database"
	"end-assignment/repository"
	"end-assignment/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetMainTask(c *gin.Context) {
	var (
		result gin.H
	)

	maintsk, err := repository.GetMainTask(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": maintsk,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertMainTask(c *gin.Context) {
	var maintsk structs.MainTask

	err := c.ShouldBindJSON(&maintsk)
	if err != nil {
		panic(err)
	}

	err = repository.InsertsMainTask(database.DbConnection, maintsk)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert ",
	})
}

func UpdateMainTask(c *gin.Context) {
	var maintsk structs.MainTask
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&maintsk)
	if err != nil {
		panic(err)
	}

	maintsk.TaskId = int64(id)

	err = repository.UpdateMainTask(database.DbConnection, maintsk)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update ",
	})
}

func DeleteMainTask(c *gin.Context) {
	var maintsk structs.MainTask
	id, err := strconv.Atoi(c.Param("id"))

	maintsk.TaskId = int64(id)

	err = repository.DeleteMaintask(database.DbConnection, maintsk)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete ",
	})
}
