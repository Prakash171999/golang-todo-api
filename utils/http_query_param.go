package utils

import (
	"boilerplate-api/models"
	"github.com/gin-gonic/gin"
)

func BuildTodoQueryParams(c *gin.Context) models.TodoQueryParams {
	todoListQueryParams := models.TodoQueryParams{}
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		//fmt.Println("QUERY", query, queryValue, value, len(value), value[len(value)-1])

		switch key {
		case "start_date":
			todoListQueryParams.StartDate = queryValue
			break
		case "due_date":
			todoListQueryParams.DueDate = queryValue
			break
		case "created_at":
			todoListQueryParams.CreatedAt = queryValue
			break
		case "priority":
			todoListQueryParams.Priority = queryValue
			break
		case "status":
			todoListQueryParams.Status = queryValue
			break
		case "category":
			todoListQueryParams.Category = queryValue
			break
		}
	}

	return todoListQueryParams
}
