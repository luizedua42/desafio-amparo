package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	models "desafio-amparo/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter sets up a router for testing
// 
// Returns:
// A router for testing
func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/tasks", GetTasks)
	router.DELETE("/tasks/:id", DelTask)
	router.PUT("/tasks", UpdateTask)
	return router
}

// TestGetTasks tests the GetTasks function
// 
// Params:
// t: testing object
func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(response), 1)
}

// TestGetTasksWithLimit tests the GetTasks function with a limit
//
// Params:
// t: testing object
func TestGetTasksWithLimit(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks?limit=3", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []models.Task
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(response))
}

// TestGetTasksWithInvalidLimit tests the GetTasks function with an invalid limit
//
// Params:
// t: testing object
func TestGetTasksWithInvalidLimit(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks?limit=5", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Limite máximo de 3 tarefas", response["message"])
}

// TestUpdateTask tests the UpdateTask function
//
// Params:
// t: testing object
func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	task := models.Task{
		ID:       2,
		Assignee: "New Assignee",
		Notes:    "New Notes",
	}

	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Tarefa atualizada com sucesso", response["message"])
}

// TestUpdateTaskInvalidID tests the UpdateTask function with an invalid ID
//
// Params:
// t: testing object
func TestUpdateTaskInvalidAssignee(t *testing.T) {
	router := setupRouter()
	
	task := models.Task{
		ID:       2,
		Assignee: "Nome muito longo para ser aceito como responsável",
	}

	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Responsável deve ter no máximo 15 caracteres", response["message"])
}

// TestUpdateTaskInvalidNotes tests the UpdateTask function with invalid notes
//
// Params:
// t: testing object
func TestUpdateTaskInvalidNotes(t *testing.T) {
	router := setupRouter()
	
	task := models.Task{
		ID:    2,
		Notes: "Test invalid notes -> greater than 350 char - Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam vehicula fermentum bibendum. Sed commodo, dolor et rhoncus ultrices, lacus ex fringilla metus, nec aliquet purus justo in magna. Vivamus lobortis pulvinar accumsan. Phasellus et sollicitudin felis. Etiam volutpat arcu porttitor, eleifend ante sed aliquam.",
	}
	
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	assert.Equal(t, http.StatusBadRequest, w.Code)
	
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Anotações devem ter no máximo 350 caracteres", response["message"])
}

// TestUpdateTaskInvalidRequest tests the UpdateTask function with an invalid request
//
// Params:
// t: testing object
func TestDelTask(t *testing.T) {
	router := setupRouter()

	// Teste para deletar uma tarefa existente
	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "Tarefa arquivada com sucesso", response["message"])

	// Teste para deletar uma tarefa inexistente
	req, _ = http.NewRequest("DELETE", "/tasks/100", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
