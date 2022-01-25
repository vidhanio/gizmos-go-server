package server

import (
	"context"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type GizmoDB struct {
	*mongo.Database
	ctx context.Context
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

func NewDB(db *mongo.Database) *GizmoDB {
	return &GizmoDB{
		Database: db,
		ctx:      context.Background(),
	}
}

func (d *GizmoDB) GetGizmos() (Gizmos, error) {
	gizmos := Gizmos{}

	cursor, err := d.Collection("gizmos").Find(d.ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(d.ctx) {
		g := &Gizmo{}
		err := cursor.Decode(g)
		if err != nil {
			return nil, err
		}

		gizmos = append(gizmos, g)
	}

	sort.Sort(gizmos)

	return gizmos, nil
}

func (d *GizmoDB) GetGizmo(resource int) (*Gizmo, error) {
	g := &Gizmo{}
	err := d.Collection("gizmos").FindOne(d.ctx, bson.M{"resource": resource}).Decode(g)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", g)

	return g, nil
}

func (d *GizmoDB) InsertGizmo(g *Gizmo) error {
	_, err := d.Collection("gizmos").InsertOne(d.ctx, g)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) UpdateGizmo(resource int, g *Gizmo) error {
	result, err := d.Collection("gizmos").UpdateOne(d.ctx, bson.M{"resource": resource}, bson.M{"$set": g})
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

func (d *GizmoDB) DeleteGizmo(resource int) error {
	result, err := d.Collection("gizmos").DeleteOne(d.ctx, bson.M{"resource": resource})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
