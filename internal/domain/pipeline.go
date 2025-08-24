package domain

// PipelinePlan represents a plan for a data pipeline.
type PipelinePlan struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Tasks       []Task `json:"tasks"`
}

// Task represents a single task in a pipeline.
type Task struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"` // e.g., "ingest", "transform", "load"
	DependsOn   []string          `json:"depends_on"`
	Config      map[string]string `json:"config"`
}
