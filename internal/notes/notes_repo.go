package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)
	if err != nil {
		return Note{}, fmt.Errorf("insert note failed: %w", err)
	}
	return note, nil
}

func (r *Repo) List(ctx context.Context) ([]Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{}
	cursor, err := r.coll.Find(opCtx, filter)
	if err != nil {
		return nil, fmt.Errorf("find note failed: %w", err)
	}

	defer cursor.Close(opCtx)
	var notes []Note

	if err := cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("Decode note failed: %w", err)
	}

	return notes, nil

}

func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	var note Note
	err := r.coll.FindOne(opCtx, filter, options.FindOne()).Decode(&note)
	if err != nil {
		return Note{}, fmt.Errorf("find note failed: %w", err)
	}

	return note, nil
}

func (r *Repo) UpdateByID(ctx context.Context, id primitive.ObjectID, update UpdateNoteRequest) (Note, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	updateData := bson.M{
		"$set": bson.M{
			"title":     update.Title,
			"content":   update.Content,
			"pinned":    update.Pinned,
			"updatedAt": time.Now().UTC(),
		},
	}
	after := options.After
	opts := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	var updatedNote Note
	err := r.coll.FindOneAndUpdate(opCtx, filter, updateData, &opts).Decode(&updatedNote)
	if err != nil {
		return Note{}, fmt.Errorf("update note failed: %w", err)
	}
	return updatedNote, nil

}

func (r *Repo) DeleteByID(ctx context.Context, id primitive.ObjectID) (bool, error) {
	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": id}
	res, err := r.coll.DeleteOne(opCtx, filter)
	if err != nil {
		return false, fmt.Errorf("delete note failed: %w", err)
	}
	if res.DeletedCount == 0 {
		return false, nil
	}
	return true, nil
}
