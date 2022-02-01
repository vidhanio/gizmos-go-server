package sql

import (
	"context"

	"github.com/jackc/pgx"
	"github.com/vidhanio/gizmos-go-server/db"
)

type GizmoDB struct {
	*pgx.Conn
	ctx context.Context
}

func New(conn *pgx.Conn) *GizmoDB {
	return &GizmoDB{
		Conn: conn,
		ctx:  context.Background(),
	}
}

func (d *GizmoDB) Start() error {
	return nil
}

func (d *GizmoDB) Stop() error {
	return d.Conn.Close()
}

func (d *GizmoDB) GetGizmos() ([]*db.Gizmo, error) {
	gizmos := []*db.Gizmo{}

	rows, err := d.Query("SELECT title, materials, description, resource, answers FROM gizmos ORDER BY resource ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		g := &db.Gizmo{}
		err := rows.Scan(&g.Title, &g.Materials, &g.Description, &g.Resource, &g.Answers)
		if err != nil {
			return nil, err
		}

		gizmos = append(gizmos, g)
	}

	return gizmos, nil
}

func (d *GizmoDB) GetGizmo(resource int) (*db.Gizmo, error) {
	g := &db.Gizmo{}
	err := d.QueryRow("SELECT title, materials, description, resource, answers FROM gizmos WHERE resource = $1", resource).Scan(&g.Title, &g.Materials, &g.Description, &g.Resource, &g.Answers)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (d *GizmoDB) InsertGizmo(g *db.Gizmo) error {
	_, err := d.Exec("INSERT INTO gizmos (title, materials, description, resource, answers) VALUES ($1, $2, $3, $4, $5)", g.Title, g.Materials, g.Description, g.Resource, g.Answers)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) UpdateGizmo(resource int, g *db.Gizmo) error {
	_, err := d.Exec("UPDATE gizmos SET title = $1, materials = $2, description = $3, answers = $4, resource = $5 WHERE resource = $6", g.Title, g.Materials, g.Description, g.Answers, g.Resource, resource)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) DeleteGizmo(resource int) error {
	_, err := d.Exec("DELETE FROM gizmos WHERE resource = $1", resource)
	if err != nil {
		return err
	}

	return nil
}
