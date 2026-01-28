package notes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) CreateNote(ctx *gin.Context) {
	var Req CreateNoteRequest
	if err := ctx.ShouldBindJSON(&Req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}
	now := time.Now().UTC()
	note := Note{
		ID:        primitive.NewObjectID(),
		Title:     Req.Title,
		Content:   Req.Content,
		Pinned:    Req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}
	created, err := h.repo.Create(ctx.Request.Context(), note)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "fauled to create note here",
		})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (h *Handler) ListNotes(ctx *gin.Context) {
	notes, err := h.repo.List(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "fauled to get all note here",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})
}
