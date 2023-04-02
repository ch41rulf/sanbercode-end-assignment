package controllers

import (
	"end-assignment/database"
	"end-assignment/repository"
	"end-assignment/structs"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetUsers(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUsers(c *gin.Context) {
	var users structs.Users

	err := c.ShouldBindJSON(&users)
	if err != nil {
		panic(err)
	}

	userId, err := repository.InsertsUsers(database.DbConnection, users)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert with user_id " + strconv.FormatInt(userId, 10),
	})
}

func UpdateUsers(c *gin.Context) {
	var users structs.Users
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&users)
	if err != nil {
		panic(err)
	}

	users.UserId = int64(id)

	err = repository.UpdateUsers(database.DbConnection, users)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update ",
	})
}

func DeleteUsers(c *gin.Context) {
	var users structs.Users
	id, err := strconv.Atoi(c.Param("id"))

	users.UserId = int64(id)

	err = repository.DeleteUsers(database.DbConnection, users)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete ",
	})
}
