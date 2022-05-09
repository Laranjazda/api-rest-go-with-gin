package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/Laranjazda/api-rest-go-with-gin/database"
	"github.com/Laranjazda/api-rest-go-with-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosOsAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)

}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "Olá " + nome + ", tudo bem?",
	})

}

func NewAluno(c *gin.Context) {
	var aluno models.Aluno
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	answer, err := database.AlunosCollection().InsertOne(ctx, aluno)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, gin.H{"answer": answer})
	// fmt.Println("Novo aluno salvo.\n", answer.InsertedID, "\n", aluno)
}

func ListaAlunos(c *gin.Context) {
	var alunos []models.Aluno
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	answer, err := database.AlunosCollection().Find(ctx, &alunos)
	if err != nil {
		panic(err.Error())
	}

	c.JSON(200, answer)

}
