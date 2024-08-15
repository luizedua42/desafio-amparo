package handlers

import (
	models "desafio-amparo/models"
	net "net/url"
	"net/http"
	strconv "strconv"
	time "time"
	// fmt "fmt"
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
	{
		ID:          1,
		Title:       "Decidir sobre doação de órgãos e tecidos",
		Description: "Decidir sobre doação de órgãos e tecidos",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          2,
		Title:       "Decidir entre sepultamento e cremação",
		Description: "Decidir entre sepultamento e cremação",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          3,
		Title:       "Obter declaração de óbito",
		Description: "Obter declaração de óbito",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          4,
		Title:       "Ligar para o plano funeral ou contratar funerária",
		Description: "Ligar para o plano funeral ou contratar funerária",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          5,
		Title:       "Comunicar à família e amigos",
		Description: "Comunicar à família e amigos",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          6,
		Title:       "Obter a certidão de óbito",
		Description: "Obter a certidão de óbito",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          7,
		Title:       "Planejar o funeral",
		Description: "Planejar o funeral",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          8,
		Title:       "Solicitar licença para se ausentar do trabalho",
		Description: "Solicitar licença para se ausentar do trabalho",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
	{
		ID:          9,
		Title:       "Planejar a missa de sétimo dia",
		Description: "Planejar a missa de sétimo dia",
		DueDate:     time.Now().Format("2006-01-02"),
		Assignee:    "Nome do Responsável",
		Notes:       "Anotações sobre a tarefa",
		Status:      "notStarted",
	},
}

func GetTasks(c *gin.Context) {
	urlStr := c.Request.URL.String()
	myUrl, _ := net.Parse(urlStr)
	params, _ := net.ParseQuery(myUrl.RawQuery)
	if params["limit"] != nil {
		limit, _ := strconv.Atoi(params["limit"][0])
		if limit == 3 && len(tasks) > 3 {
			c.IndentedJSON(http.StatusOK, tasks[:limit])
			return
		}
		if limit > 3 {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Limite máximo de 3 tarefas"})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func DelTask(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID inválido"})
		return
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Tarefa arquivada com sucesso"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tarefa não encontrada"})
}

func UpdateTask(c *gin.Context) {


	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Requisição inválid"})
		return
	}

	for i, t := range tasks {
		if t.ID == task.ID {
			if task.DueDate != "" {
				tasks[i].DueDate = task.DueDate
			}
			if task.Assignee != "" {
				if len(task.Assignee) > 15 {
					c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Responsável deve ter no máximo 15 caracteres"})
					return
				}
				tasks[i].Assignee = task.Assignee
			}
			if task.Notes != "" {
				if len(task.Notes) > 350 {
					c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Anotações devem ter no máximo 350 caracteres"})
					return
				}
				tasks[i].Notes = task.Notes
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Tarefa atualizada com sucesso"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Tarefa não encontrada"})
}