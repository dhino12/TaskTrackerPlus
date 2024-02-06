package api

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskAPI interface {
	AddTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTaskByID(c *gin.Context)
	GetTaskList(c *gin.Context)
	GetTaskListByCategory(c *gin.Context)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskRepo service.TaskService) *taskAPI {
	return &taskAPI{taskRepo}
}

func (t *taskAPI) AddTask(c *gin.Context) {
	var newTask model.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	err := t.taskService.Store(&newTask)
	fmt.Println("======== ERR")
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "add task success"})
}

func (t *taskAPI) UpdateTask(c *gin.Context) {
	// TODO: answer here
	task := model.Task{}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Invalid task ID",
		})
		return
	}
	task.ID = taskId

	errorUpdateTask := t.taskService.Update(task.ID, &task)
	if errorUpdateTask != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Error Update!..",
		})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "update task success",
	})
}

func (t *taskAPI) DeleteTask(c *gin.Context) {
	// TODO: answer here
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Error: "Invalid task ID",
		})
		return
	}

	errorDeleteTask := t.taskService.Delete(taskId)
	if errorDeleteTask != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse {
			Error: "Error Delete!...",
		})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{
		Message: "delete task success",
	})
}

func (t *taskAPI) GetTaskByID(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid task ID"})
		return
	}

	task, err := t.taskService.GetByID(taskID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (t *taskAPI) GetTaskList(c *gin.Context) {
	// TODO: answer here
	tasks, err := t.taskService.GetList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (t *taskAPI) GetTaskListByCategory(c *gin.Context) {
	// TODO: answer here
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}
	tasks, err := t.taskService.GetTaskCategory(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Error: "Internal Server Error",
		})
		return
	}
	
	c.JSON(http.StatusOK, tasks)
}
