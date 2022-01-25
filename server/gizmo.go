package server

type Gizmo struct {
	Title       string   `bson:"title" json:"title"`
	Materials   string   `bson:"materials" json:"materials"`
	Description string   `bson:"description" json:"description"`
	Resource    int      `bson:"resource" json:"resource"`
	Answers     []string `bson:"answers" json:"answers"`
}

type Gizmos []*Gizmo

func (gs Gizmos) Len() int {
	return len(gs)
}

func (gs Gizmos) Less(i, j int) bool {
	return gs[i].Resource < gs[j].Resource
}

func (gs Gizmos) Swap(i, j int) {
	gs[i], gs[j] = gs[j], gs[i]
}
