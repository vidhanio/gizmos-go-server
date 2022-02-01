package db

type GizmoDB interface {
	Start() error
	Stop() error
	GetGizmos() ([]*Gizmo, error)
	GetGizmo(resource int) (*Gizmo, error)
	InsertGizmo(g *Gizmo) error
	UpdateGizmo(resource int, g *Gizmo) error
	DeleteGizmo(resource int) error
}
