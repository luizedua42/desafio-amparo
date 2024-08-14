package handlers

import (
	"net/http"
	"time"
	gin "github.com/gin-gonic/gin"
)

/* Tasks
1	Decidir sobre doação de órgãos e tecidos
2	Decidir entre sepultamento e cremação
3	Obter declaração de óbito
4	Ligar para o plano funeral ou contratar funerária
5	Comunicar à família e amigos
6	Obter a certidão de óbito
7	Planejar o funeral
8	Solicitar licença para se ausentar do trabalho
9	Planejar a missa de sétimo dia
*/

var tasks = []models.Task{
	{ID: 1, Title: "Decidir sobre doação de órgãos e tecidos", Description: "Descrição da Tarefa", , Assignee: "Família", Notes: "Decidir sobre doação de órgãos e tecidos", Status: "Pendente"},
}