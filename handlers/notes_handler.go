package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/souviks72/notes-app-api/dbiface"
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

	note.DateCreated = time.Now().String()
	note.DateModified = time.Now().String()

	_, err = h.Coll.InsertOne(context.Background(), note)
	if err != nil {
		fmt.Printf("Unable to insert note to db %+v", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Unable to insert note to db"})
	}
	return c.JSON(http.StatusOK, note)
}
