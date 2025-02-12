package controller

import (
	"log"

	"github.com/Kelniit/Halu/config"
	"github.com/Kelniit/Halu/entities"
	"github.com/Kelniit/Halu/utilities"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	// Get All Users
	database, err := config.TableSetup()
	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	var UserResult []entities.UserEntity

	result := database.Find(&UserResult)
	if result.Error != nil {
		utilities.FailMess(c, 500, "Fail to Get Sample Data", result.Error.Error())
		return
	}

	c.JSON(200, UserResult)
}

func GetUserID(c *gin.Context) {
	// Get User
	UID := c.Param("UID")

	if UID == "" {
		utilities.FailMess(c, 400, "User ID is Missing !")
		return
	}

	database, err := config.TableSetup()
	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	var SingleUser entities.UserEntity

	GetFailUser := database.First(&SingleUser, UID).Error

	if GetFailUser != nil {
		if GetFailUser.Error() == "record not found" {
			utilities.FailMess(c, 400, "User Unavailable !")
		}
		return
	}

	c.JSON(200, SingleUser)
}

func MoreUsers(c *gin.Context) {
	// More Users
	var users []entities.UserEntity

	// Bind JSON Request Body to Users Slice
	if json_error := c.ShouldBindBodyWithJSON(&users); json_error != nil {
		utilities.FailMess(c, 400, "Fail to Bind JSON", json_error.Error())
		return
	}

	database, err := config.TableSetup()
	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	// Ensure Users to Insert
	if len(users) == 0 {
		utilities.FailMess(c, 400, "Empty Users !")
		return
	}

	if create_error := database.CreateInBatches(users, len(users)).Error; create_error != nil {
		utilities.FailMess(c, 500, "Fail to Create Multiple Sample !")
	}

	c.JSON(200, gin.H{"result": "Users Inserted Successfully !"})
}

func DeleteUser(c *gin.Context) {
	// Delete User
	UID := c.Param("UID")

	if UID == "" {
		utilities.FailMess(c, 400, "User ID is Missing !")
		return
	}

	database, err := config.TableSetup()
	if err != nil {
		log.Fatalf("Fail to Connect to Database : %v", err)
		utilities.FailMess(c, 500, "Database Connection Fail !")
		return
	}

	var SingleUser entities.UserEntity

	GetFailUser := database.First(&SingleUser, UID).Error

	if GetFailUser != nil {
		if GetFailUser.Error() == "record not found" {
			utilities.FailMess(c, 400, "User Unavailable !")
		}
		return
	}

	DeleteUserFail := database.Delete(&entities.UserEntity{}, UID).Error

	if DeleteUserFail != nil {
		utilities.FailMess(c, 500, "Error on Deleting Sample", GetFailUser.Error())
		return
	}

	c.JSON(200, gin.H{"result": "User Deleted !"})
}
