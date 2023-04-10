package main

import (
	
	"net/http"

	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/cmd/api/handlers"
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/domain"
	"github.com/IsaiasgomzR/bootcamp-golang-web.git/method_patch/internal/movies"
	"github.com/gin-gonic/gin"
)

func main()  {
	db:=[]*domain.Movie{}
	rp:=movies.NewRepositoryLocal(db,0)
	s:= movies.NewService(rp)
	ct:= handlers.NewContorllerService(s)

	sv:= gin.Default()

	sv.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK,"pong")
	})

	movies:= sv.Group("/movies")
	{
		movies.GET("/:id", ct.GetId())
		movies.POST("", ct.Create())
		
		movies.PUT("/:id", ct.Update())
		movies.PATCH("/:id/genre", ct.UpdateGenre())
		movies.PATCH("/:id", ct.UpdatePartial())
		movies.DELETE("/:id", ct.Delete())
	}

	if err:= sv.Run(); err != nil{
		panic(err)
	}
	
}