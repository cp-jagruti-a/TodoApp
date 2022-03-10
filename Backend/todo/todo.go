package todo

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type TodoList struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type TodolistRepository struct {
	Db *sqlx.DB
}

func New(db *sqlx.DB) *TodolistRepository {
	return &TodolistRepository{Db: db}
}
func (repository *TodolistRepository) Add(c *gin.Context) {
	input := TodoList{}

	err := c.ShouldBindWith(&input, binding.JSON)

	if err != nil {
		log.Error(err)
		c.Abort()
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data"})
		return
	}
	res, err := repository.Db.Exec(`INSERT INTO todolist(title,description,created_at,updated_at) VALUES(?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP())`, input.Title,input.Description)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	todoId, err := res.LastInsertId()
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	createdTodo, err := repository.getTodoById(int(todoId))
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createdTodo)

}
func (repository *TodolistRepository) Update(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = repository.getTodoById(int(todoId))
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	input := TodoList{}

	err = c.ShouldBindWith(&input, binding.JSON)
	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repository.Db.MustExec(`UPDATE todolist SET title=?, description=? ,updated_at = UTC_TIMESTAMP() WHERE id = ?`, input.Title, input.Description, todoId)

	updatedTodo, err := repository.getTodoById(todoId)

	if err != nil {
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, updatedTodo)
}
func (repository *TodolistRepository) Delete(c *gin.Context) {

	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	_, err = repository.getTodoById(int(todoId))
	if err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Error(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	repository.Db.MustExec(`DELETE FROM todolist WHERE todolist.id = ?`, todoId)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully", "todoId": todoId})
}
func (repository *TodolistRepository) Get(c *gin.Context) {

	todolist := []TodoList{}

	err := repository.Db.Select(&todolist, `SELECT id,title,description,created_at,updated_at FROM todolist`)

	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, todolist)
}

func (repository *TodolistRepository) GetTodoById(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todolist := TodoList{}

	err = repository.Db.Get(&todolist, `SELECT id,title,description,created_at,updated_at FROM todolist
										WHERE id = ?`, todoId)

	if err != nil {
		log.Warn(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, todolist)
}
func (repository *TodolistRepository) getTodoById(id int) (*TodoList, error) {
	todolist := TodoList{}

	err := repository.Db.Get(&todolist, `SELECT id,title,description,created_at,updated_at FROM todolist
										WHERE id = ?`, id)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &todolist, nil
}
