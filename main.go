package main

import (
	"github.com/Laranjazda/api-rest-go-with-gin/models"
	"github.com/Laranjazda/api-rest-go-with-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Nome: "Carlos Roberto", CPF: "01551974010", RG: "7089838085"},
		{Nome: "Maria Heloisa", CPF: "25455478410", RG: "5587452145"},
	}
	routes.HandleRequests()
}
