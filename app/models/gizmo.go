package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Gizmo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Materials   string             `bson:"materials" json:"materials"`
	Description string             `bson:"description" json:"description"`
	Resource    int                `bson:"resource" json:"resource"`
	Answers     []string           `bson:"answers" json:"answers"`
}

func NewGizmo(title string, materials string, description string, resource int, answers []string) *Gizmo {
	return &Gizmo{
		Title:       title,
		Materials:   materials,
		Description: description,
		Resource:    resource,
		Answers:     answers,
	}
}
