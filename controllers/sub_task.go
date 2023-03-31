package controllers

import (
	"end-assignment/database"
	"end-assignment/repository"
	"end-assignment/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetSubTask(c *gin.Context) {
	var (
		result gin.H
	)

	subtsk, err := repository.GetSubtask(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": subtsk,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertSubTask(c *gin.Context) {
	var subtsk structs.Subtask

	err := c.ShouldBindJSON(&subtsk)
	if err != nil {
		panic(err)
	}

	err = repository.InsertsSubtask(database.DbConnection, subtsk)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert ",
	})
}

func UpdateSubTask(c *gin.Context) {
	var subtsk structs.Subtask
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&subtsk)
	if err != nil {
		panic(err)
	}

	subtsk.SubtaskID = int64(id)

	err = repository.UpdateSubtask(database.DbConnection, subtsk)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update ",
	})
}

func DeleteSubTask(c *gin.Context) {
	var subtsk structs.Subtask
	id, err := strconv.Atoi(c.Param("id"))

	subtsk.SubtaskID = int64(id)

	err = repository.DeleteSubtask(database.DbConnection, subtsk)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete ",
	})
}
