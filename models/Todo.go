package models



type Todo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}