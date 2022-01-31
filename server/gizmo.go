package server

import "github.com/vidhanio/gizmos-go-server/database"

type Gizmo struct {
	Title       string   `json:"title"`
	Materials   string   `json:"materials"`
	Description string   `json:"description"`
	Resource    int      `json:"resource"`
	Answers     []string `json:"answers"`
}

func NewGizmoFromDBGizmo(dbGizmo *database.Gizmo) *Gizmo {
	return &Gizmo{
		Title:       dbGizmo.Title,
		Materials:   dbGizmo.Materials,
		Description: dbGizmo.Description,
		Resource:    dbGizmo.Resource,
		Answers:     dbGizmo.Answers,
	}
}

func NewGizmosFromDBGizmos(dbGizmos []*database.Gizmo) []*Gizmo {
	gizmos := make([]*Gizmo, len(dbGizmos))

	for i, dbGizmo := range dbGizmos {
		gizmos[i] = NewGizmoFromDBGizmo(dbGizmo)
	}

	return gizmos
}

func NewDBGizmoFromGizmo(gizmo *Gizmo) *database.Gizmo {
	return &database.Gizmo{
		Title:       gizmo.Title,
		Materials:   gizmo.Materials,
		Description: gizmo.Description,
		Resource:    gizmo.Resource,
		Answers:     gizmo.Answers,
	}
}

func NewDBGizmosFromGizmos(gizmos []*Gizmo) []*database.Gizmo {
	dbGizmos := make([]*database.Gizmo, len(gizmos))

	for i, gizmo := range gizmos {
		dbGizmos[i] = NewDBGizmoFromGizmo(gizmo)
	}

	return dbGizmos
}
