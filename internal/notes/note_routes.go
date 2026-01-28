package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoutes(r *gin.Engine, db *mongo.Database) {
	repo := NewRepo(db)
	h := NewHandler(repo)

	notesGroup := r.Group("/notes")
	{
		notesGroup.POST("", h.CreateNote)
		notesGroup.GET("", h.ListNotes)
		notesGroup.GET("/:id", h.GetNoteByID)
		notesGroup.PUT("/:id", h.UpdateNoteByID)
		notesGroup.DELETE("/:id", h.DeleteNoteByID)
	}
}
