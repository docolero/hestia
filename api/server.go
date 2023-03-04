package api

import (
	"log"
	"net/http"

	"github.com/docolero/hestia/internal/model"
	"github.com/gin-gonic/gin"
)

type Server struct {
	listenAddr string
	engine     *gin.Engine
	todolists  map[string]model.TodoList
}

func NewServer(listenAddr string, e *gin.Engine) *Server {
	return &Server{
		listenAddr: listenAddr,
		engine:     e,
		todolists:  make(map[string]model.TodoList),
	}
}

func (s *Server) Start() error {

	todoListGroup := s.engine.Group("/todolist")
	{
		todoListGroup.GET("", s.handleGetAllTodoLists)
		todoListGroup.GET("/:id", s.handleGetTodoListById)
		todoListGroup.POST("", s.handlePostTodoList)
	}

	s.engine.GET("/live", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	s.engine.GET("/ready", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return s.engine.Run(s.listenAddr)
}

func (s *Server) handlePostTodoList(c *gin.Context) {
	log.Println("called handlePostTodoList")
	var todolist model.TodoList
	err := c.ShouldBind(&todolist)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	s.todolists[todolist.Name] = todolist
	c.Status(http.StatusCreated)
}

func (s *Server) handleGetAllTodoLists(c *gin.Context) {
	log.Println("called handleGetAllTodoLists")
	c.IndentedJSON(http.StatusOK, model.TodoList{})
}

func (s *Server) handleGetTodoListById(c *gin.Context) {
	id := c.Param("id")
	log.Println("called handleGetTodoListById with id", id)
	if tl, exists := s.todolists[id]; exists {
		c.IndentedJSON(http.StatusOK, tl)
	} else {
		c.Status(http.StatusNotFound)
	}
}
