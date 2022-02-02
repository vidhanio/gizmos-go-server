package json

import "github.com/vidhanio/gizmos-go-server/db"

type gizmoSorter []*db.Gizmo

func (g gizmoSorter) Len() int {
	return len(g)
}

func (g gizmoSorter) Less(i, j int) bool {
	return g[i].Resource < g[j].Resource
}

func (g gizmoSorter) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
