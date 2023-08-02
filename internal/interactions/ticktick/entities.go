package ticktick

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID        string   `json:"id"`
	ProjectID string   `json:"projectId"`
	Title     string   `json:"title"`
	Tags      []string `json:"tags"`
}
