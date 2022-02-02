package mongodb

import (
	"context"

	"github.com/vidhanio/gizmos-go-server/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GizmoDB struct {
	*mongo.Database
	ctx context.Context
}

func New(uri string, db string) *GizmoDB {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return &GizmoDB{
		Database: client.Database(db),
		ctx:      context.Background(),
	}
}

func (d *GizmoDB) Start() error {
	return d.Client().Connect(d.ctx)
}

func (d *GizmoDB) Stop() error {
	return d.Client().Disconnect(d.ctx)
}

func (d *GizmoDB) GetGizmos() ([]*db.Gizmo, error) {
	gizmos := []*db.Gizmo{}

	cursor, err := d.Collection("gizmos").Find(d.ctx, bson.M{}, options.Find().SetSort(bson.M{"resource": 1}))
	if err != nil {
		return nil, err
	}

	for cursor.Next(d.ctx) {
		g := &db.Gizmo{}
		err := cursor.Decode(g)
		if err != nil {
			return nil, err
		}

		gizmos = append(gizmos, g)
	}

	return gizmos, nil
}

func (d *GizmoDB) GetGizmo(resource int) (*db.Gizmo, error) {
	g := &db.Gizmo{}
	err := d.Collection("gizmos").FindOne(d.ctx, bson.M{"resource": resource}).Decode(g)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (d *GizmoDB) InsertGizmo(g *db.Gizmo) error {
	_, err := d.Collection("gizmos").InsertOne(d.ctx, g)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) UpdateGizmo(resource int, g *db.Gizmo) error {
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
