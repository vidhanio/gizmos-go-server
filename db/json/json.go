package json

import (
	"encoding/json"
	"io"
	"os"
	"sort"

	"github.com/vidhanio/gizmos-go-server/db"
)

type GizmoDB struct {
	*os.File
}

func New(filename string) *GizmoDB {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return &GizmoDB{
		File: file,
	}
}

func (d *GizmoDB) Start() error {
	return nil
}

func (d *GizmoDB) Stop() error {
	return d.Close()
}

func (d *GizmoDB) GetGizmos() ([]*db.Gizmo, error) {
	gizmos := []*db.Gizmo{}

	contents, err := io.ReadAll(d)
	if err != nil {
		return nil, err
	}

	_, err = d.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &gizmos)
	if err != nil {
		return nil, err
	}

	sort.Sort(gizmoSorter(gizmos))

	return gizmos, nil
}

func (d *GizmoDB) GetGizmo(resource int) (*db.Gizmo, error) {
	gizmos, err := d.GetGizmos()
	if err != nil {
		return nil, err
	}

	for _, g := range gizmos {
		if g.Resource == resource {
			return g, nil
		}
	}

	return nil, nil
}

func (d *GizmoDB) InsertGizmo(g *db.Gizmo) error {
	gizmos, err := d.GetGizmos()
	if err != nil {
		return err
	}

	gizmos = append(gizmos, g)

	contents, err := json.Marshal(gizmos)
	if err != nil {
		return err
	}

	_, err = d.Write(contents)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) UpdateGizmo(resource int, g *db.Gizmo) error {
	gizmos, err := d.GetGizmos()
	if err != nil {
		return err
	}

	for i, gizmo := range gizmos {
		if gizmo.Resource == resource {
			gizmos[i] = g
			break
		}
	}

	contents, err := json.Marshal(gizmos)
	if err != nil {
		return err
	}

	_, err = d.Write(contents)
	if err != nil {
		return err
	}

	return nil
}

func (d *GizmoDB) DeleteGizmo(resource int) error {
	gizmos, err := d.GetGizmos()
	if err != nil {
		return err
	}

	for i, gizmo := range gizmos {
		if gizmo.Resource == resource {
			gizmos = append(gizmos[:i], gizmos[i+1:]...)
			break
		}
	}

	contents, err := json.Marshal(gizmos)
	if err != nil {
		return err
	}

	_, err = d.Write(contents)
	if err != nil {
		return err
	}

	return nil
}
