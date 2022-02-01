package server

import "github.com/vidhanio/gizmos-go-server/db"

type Gizmo struct {
	Title       string   `json:"title"`
	Materials   string   `json:"materials"`
	Description string   `json:"description"`
	Resource    int      `json:"resource"`
	Answers     []string `json:"answers"`
}

func NewGizmoFromDBGizmo(dbGizmo *db.Gizmo) *Gizmo {
	return &Gizmo{
		Title:       dbGizmo.Title,
		Materials:   dbGizmo.Materials,
		Description: dbGizmo.Description,
		Resource:    dbGizmo.Resource,
		Answers:     dbGizmo.Answers,
	}
}

func NewGizmosFromDBGizmos(dbGizmos []*db.Gizmo) []*Gizmo {
	gizmos := make([]*Gizmo, len(dbGizmos))

	for i, dbGizmo := range dbGizmos {
		gizmos[i] = NewGizmoFromDBGizmo(dbGizmo)
	}

	return gizmos
}

func NewDBGizmoFromGizmo(gizmo *Gizmo) *db.Gizmo {
	return &db.Gizmo{
		Title:       gizmo.Title,
		Materials:   gizmo.Materials,
		Description: gizmo.Description,
		Resource:    gizmo.Resource,
		Answers:     gizmo.Answers,
	}
}

func NewDBGizmosFromGizmos(gizmos []*Gizmo) []*db.Gizmo {
	dbGizmos := make([]*db.Gizmo, len(gizmos))

	for i, gizmo := range gizmos {
		dbGizmos[i] = NewDBGizmoFromGizmo(gizmo)
	}

	return dbGizmos
}
