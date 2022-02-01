package db

type Gizmo struct {
	Title       string   `bson:"title"`
	Materials   string   `bson:"materials"`
	Description string   `bson:"description"`
	Resource    int      `bson:"resource"`
	Answers     []string `bson:"answers"`
}
