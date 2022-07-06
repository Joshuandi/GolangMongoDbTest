package controller

import (
	"GolangMongoDbTest/config"
	"GolangMongoDbTest/model"
	"GolangMongoDbTest/response"
	"GolangMongoDbTest/util"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

func InsertUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user model.User
	defer cancel()

	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	pass, errHash := util.GenerateHashPassword(user.Password)
	if errHash != nil {
		fmt.Println("Error Hash : " + errHash.Error())
		return errHash
	}
	user.Password = pass

	newRegister := model.User{
		Id:       primitive.NewObjectID(),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}

	result, err := userCollection.InsertOne(ctx, newRegister)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Register",
			Data:    &echo.Map{"data": result}})
}

func GetUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var user model.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	return c.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Get User",
			Data:    &echo.Map{"data": user}})
}

func UpdateUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	var user model.User
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(id)
	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	update := bson.M{"username": user.Username, "email": user.Email}
	result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//getUpdated user details
	var updatedUser model.User
	if result.MatchedCount == 1 {
		err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				response.UserResponse{
					Status:  http.StatusInternalServerError,
					Message: "error",
					Data:    &echo.Map{"data": err.Error()}})
		}
	}

	return c.JSON(http.StatusCreated,
		response.UserResponse{
			Status:  http.StatusCreated,
			Message: "Successfully Updated",
			Data:    &echo.Map{"data": updatedUser}})
}

func DeleteUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("id")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	result, err := userCollection.DeleteOne(ctx, bson.M{"id": objId})
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	if result.DeletedCount < 1 {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": "User with specified ID not Found"}})
	}
	return c.JSON(http.StatusOK,
		response.UserResponse{
			Status:  http.StatusOK,
			Message: "Success",
			Data:    &echo.Map{"data": "User successfully deleted!"}})
}

func GetAllUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user []model.User
	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			response.UserResponse{
				Status:  http.StatusBadRequest,
				Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from db
	defer results.Close(ctx)
	for results.Next(ctx) {
		var oneUser model.User
		if err = results.Decode(&oneUser); err != nil {
			return c.JSON(http.StatusBadRequest,
				response.UserResponse{
					Status:  http.StatusBadRequest,
					Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
		user = append(user, oneUser)
	}
	return c.JSON(http.StatusOK, response.UserResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    &echo.Map{"data": user}})
}
