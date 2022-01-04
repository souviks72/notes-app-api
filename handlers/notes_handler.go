package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/souviks72/notes-app-api/dbiface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

type Notes struct {
	Title        string `json:"title" bson:"title" validate:"required,max=10,min=3"`
	Body         string `json:"body" bson:"body" validate:"required,max=100,min=3"`
	DateCreated  string `json:"date_created,omitempty" bson:"date_created,omitempty"`
	DateModified string `json:"date_modified,omitempty" bson:"date_modified,omitempty"`
}

type NotesHandler struct {
	Coll dbiface.CollectionAPI
}

func (h NotesHandler) CreateNote(c echo.Context) error {
	var note Notes
	err := c.Bind(&note)
	if err != nil {
		fmt.Printf("Unable to bind request body %+v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Unable to bind request body"})
	}

	err = v.Struct(note)
	if err != nil {
		fmt.Printf("Invalid request body %+v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
	}

	note.DateCreated = time.Now().Local().String()
	note.DateModified = time.Now().Local().String()

	_, err = h.Coll.InsertOne(context.Background(), note)
	if err != nil {
		fmt.Printf("Unable to insert note to db %+v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Unable to insert note to db"})
	}
	return c.JSON(http.StatusOK, note)
}

func (h *NotesHandler) GetAllNotes(c echo.Context) error {
	var note Notes
	var notes []Notes
	findCursor, err := h.Coll.Find(context.Background(), bson.M{})
	if err != nil {
		fmt.Printf("Error fetching results %+v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching results"})
	}

	for findCursor.Next(context.Background()) {
		err = findCursor.Decode(&note)
		if err != nil {
			fmt.Printf("Error decoding cursor results %+v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error decoding cursor results"})
		}
		notes = append(notes, note)
	}

	return c.JSON(http.StatusOK, notes)
}

func (h *NotesHandler) GetNote(c echo.Context) error {
	var note Notes

	id := c.Param("id")
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Printf("Unable to convert id to hex %+v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to convert id to hex"})
	}
	queryFilter := bson.M{"_id": hexId}
	err = h.Coll.FindOne(context.Background(), queryFilter).Decode(&note)
	if err != nil {
		fmt.Printf("Error fetching results %+v\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching results"})
	}

	return c.JSON(http.StatusOK, note)
}

func (h *NotesHandler) DeleteNote(c echo.Context) error {
	var note Notes

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Printf("Unable to convert id to hex %+v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Unable to convert id to hex"})
	}

	queryFilter := bson.M{"_id": id}
	res := h.Coll.FindOneAndDelete(context.Background(), queryFilter)
	err = res.Decode(&note)
	if err != nil {
		fmt.Printf("Error decoding result %+v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error decoding result"})
	}

	return c.JSON(http.StatusOK, note)
}
