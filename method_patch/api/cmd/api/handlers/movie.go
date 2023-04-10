package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/domain"
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/movies"
	"github.com/gin-gonic/gin"
)

func NewContorllerService(sv movies.Service) *ControllerMovie  {
	return &ControllerMovie{sv: sv}
}

type ControllerMovie struct{
	sv movies.Service
}

func (ct *ControllerMovie) GetId() gin.HandlerFunc   {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
			return
		}
		mv, err := ct.sv.GetId(id)
		if err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusNotFound, gin.H{"message":"movie not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message":"succes", "data": mv})
	}
}

func (ct *ControllerMovie) Create() gin.HandlerFunc  {
	println("///////////////////////////here past")
	type Request struct{
		Title string `json:"title" binding:"required "`
		Year int `json:"year" binding:"required"`
		Genre string `json:"genre"  binding:"required"`
		Rating *float64 `json:"rating" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		fmt.Println(ctx.Request.Body)
		
		if err := ctx.ShouldBindJSON(&request); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invlaid request "})
		}
		fmt.Println(request)
		mv := &domain.Movie{
			Title:request.Title,
			Year : request.Year,
			Genre : request.Genre,
			Rating : *request.Rating,
		}

		err := ct.sv.Create(mv)
		if err != nil{
			if errors.Is(err, movies.ErrServiceInvalid){
				ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message":"invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message":"internal error"})
		}

		ctx.JSON(http.StatusCreated,gin.H{"message":"success", "data":mv})
	}
}

func (ct *ControllerMovie) Update() gin.HandlerFunc  {
	type Request struct{
		Title string `json:"title" binding:"requiered"`
		Year int `json:"year bindinf:"required"`
		Genre string `json:"genre" binding:"required`
		Rating float64 `json:"rating" bindin:"required"`
	}
	return func (ctx *gin.Context)  {
		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}

		var request Request
		if err := ctx.ShouldBindJSON(&request); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		//process
		mv := &domain.Movie{
			Title : request.Title,
			Year: request.Year,
			Genre: request.Genre,
			Rating: request.Rating,
		}
		if err := ct.sv.Update(id, mv); err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusNotFound, gin.H{"message": "moive not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid){
				ctx.JSON(http.StatusUnprocessableEntity,gin.H{"message": "invalid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message":"internal error"})
		}

		//response
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})
	}
}

func (ct *ControllerMovie) UpdateGenre() gin.HandlerFunc  {
	type request struct{
		Genre string `json:"genre" binding:"required"`
	}

	return func(ctx *gin.Context) {
		//request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invlaid request"})
		}

		//process
		mv, err := ct.sv.GetId(id)
		if err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusNotFound, gin.H{"message":"moive not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid){
				ctx.JSON(http.StatusUnprocessableEntity,gin.H{"message": "inlvaid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message":"intenal error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": mv})

	}
}

func (ct *ControllerMovie) UpdatePartial () gin.HandlerFunc  {
	return func(ctx *gin.Context) {
		id,err := strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message":"invlaid request"})
			return
		}
		mv, err := ct.sv.GetId(id)
		if err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": "moive not found"})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"message":"invlaid request"})
		}
		if err := json.NewDecoder(ctx.Request.Body).Decode(&mv); err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
			return
		}
		mv.Id =id
		
		//process
		if err := ct.sv.Update(id, mv); err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusNotFound,gin.H{"message":"moive not found"})
				return
			}
			if errors.Is(err, movies.ErrServiceInvalid){
				ctx.JSON(http.StatusUnprocessableEntity,gin.H{"messsage": "invlaid movie"})
				return
			}
			ctx.JSON(http.StatusInternalServerError,gin.H{"message": "internal error"})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{"message":"success", "data": mv})
	}
}

func (ct *ControllerMovie) Delete ()  gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err:= strconv.Atoi(ctx.Param("id"))
		if err != nil{
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
			return
		}
		if err := ct.sv.Delete(id); err != nil{
			if errors.Is(err, movies.ErrServiceNotFound){
				ctx.JSON(http.StatusNotFound,gin.H{"message":"movie not found"})
			}
			ctx.JSON(http.StatusInternalServerError,gin.H{"message":"internal error"})
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}