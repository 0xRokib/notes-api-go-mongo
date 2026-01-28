package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
			"error": "failed to create note here",
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

func (h *Handler) GetNoteByID(ctx *gin.Context) {
	id := ctx.Param("id")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return
	}
	note, err := h.repo.GetByID(ctx.Request.Context(), objId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "note not found with this given ID",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch the note",
		})
		return
	}
	ctx.JSON(http.StatusOK, note)
}
