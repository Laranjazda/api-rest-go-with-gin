package routes

import (
	"github.com/Laranjazda/api-rest-go-with-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosOsAlunos)

	r.GET("/lista", controllers.ListaAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.POST("/api/new", controllers.NewAluno)
	r.GET("/api/labels/:age,index,sex,language", controllers.GetLabels)

	r.Run(":2020")
}
