package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oshie15/go-course.git/db"
)

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}


// return id of the task created on success
func SaveTask(ctx *gin.Context){

	var payload PostTaskPayload


	
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":"Unable to read the body"})
		return
	}


	var id int
	
	query := `Insert into tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	err := db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg":err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})		

}