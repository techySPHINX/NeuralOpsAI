package domain

import (
	"time"
)

// DataVersion represents a version of a dataset in the data lake
type DataVersion struct {
	ID            string            `json:"id"`
	DatasetName   string            `json:"dataset_name"`
	Version       string            `json:"version"`
	SnapshotID    int64             `json:"snapshot_id"`
	CommitMessage string            `json:"commit_message"`
	CommittedBy   string            `json:"committed_by"`
	CommittedAt   time.Time         `json:"committed_at"`
	Schema        map[string]string `json:"schema"`
	Metrics       DataMetrics       `json:"metrics"`
	ParentVersion string            `json:"parent_version,omitempty"`
	Tags          []string          `json:"tags"`
}

// DataMetrics contains statistics about the dataset
type DataMetrics struct {
	RowCount      int64   `json:"row_count"`
	ColumnCount   int     `json:"column_count"`
	SizeBytes     int64   `json:"size_bytes"`
	PartitionCount int    `json:"partition_count"`
	FileCount     int     `json:"file_count"`
}

// DataLineage tracks the lineage of data transformations
type DataLineage struct {
	ID              string              `json:"id"`
	DatasetName     string              `json:"dataset_name"`
	Version         string              `json:"version"`
	Sources         []DataSource        `json:"sources"`
	Transformations []Transformation    `json:"transformations"`
	Pipeline        PipelineReference   `json:"pipeline"`
	CreatedAt       time.Time           `json:"created_at"`
}

// DataSource represents the source of data
type DataSource struct {
	Type        string            `json:"type"` // e.g., "iceberg", "s3", "database"
	Location    string            `json:"location"`
	Version     string            `json:"version,omitempty"`
	Metadata    map[string]string `json:"metadata"`
}

// Transformation represents a data transformation operation
type Transformation struct {
	Name        string            `json:"name"`
	Type        string            `json:"type"` // e.g., "filter", "join", "aggregate"
	Description string            `json:"description"`
	Parameters  map[string]string `json:"parameters"`
	Order       int               `json:"order"`
}

// PipelineReference links data to the pipeline that created it
type PipelineReference struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	RunID   string `json:"run_id"`
}

// DataCatalogEntry represents a dataset in the catalog
type DataCatalogEntry struct {
	Name            string            `json:"name"`
	Description     string            `json:"description"`
	CurrentVersion  string            `json:"current_version"`
	VersionCount    int               `json:"version_count"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
	Owner           string            `json:"owner"`
	Tags            []string          `json:"tags"`
	Schema          map<string]string `json:"schema"`
	Location        string            `json:"location"`
	Format          string            `json:"format"` // e.g., "iceberg", "parquet", "delta"
}
