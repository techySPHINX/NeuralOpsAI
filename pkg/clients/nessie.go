package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NessieClient struct {
	endpoint string
	client   *http.Client
}

func NewNessieClient(endpoint string) *NessieClient {
	return &NessieClient{
		endpoint: endpoint,
		client:   &http.Client{},
	}
}

// Nessie API models (simplified for now)
type NessieContentKey struct {
	Elements []string `json:"elements"`
}

type NessieIcebergTable struct {
	MetadataLocation string `json:"metadataLocation"`
	SnapshotID       int64  `json:"snapshotId"`
	SchemaID         int64  `json:"schemaId"`
	SpecId           int64  `json:"specId"`
	SortOrderID      int64  `json:"sortOrderId"`
}

type NessieContent struct {
	Type        string             `json:"type"`
	Id          string             `json:"id"`
	IcebergTable NessieIcebergTable `json:"icebergTable"`
}

type NessieCommit struct {
	Message string `json:"message"`
}

type NessieOperation struct {
	Type    string           `json:"type"` // e.g., "CREATE_CONTENT", "SET_CONTENT"
	Key     NessieContentKey `json:"key"`
	Content NessieContent    `json:"content"`
}

type NessieCommitRequest struct {
	Commit      NessieCommit      `json:"commit"`
	Operations  []NessieOperation `json:"operations"`
	Hash        string            `json:"hash"`
	BranchName  string            `json:"branchName"`
	ExpectedHash string           `json:"expectedHash"`
}

// CreateTable creates a new Iceberg table in Nessie.
func (c *NessieClient) CreateTable(ctx context.Context, tableName, metadataLocation string) error {
	key := NessieContentKey{Elements: []string{tableName}}
	content := NessieContent{
		Type: "ICEBERG_TABLE",
		IcebergTable: NessieIcebergTable{
			MetadataLocation: metadataLocation,
		},
	}
	operation := NessieOperation{
		Type:    "CREATE_CONTENT",
		Key:     key,
		Content: content,
	}

	commitReq := NessieCommitRequest{
		Commit:      NessieCommit{Message: fmt.Sprintf("Create table %s", tableName)},
		Operations:  []NessieOperation{operation},
		BranchName:  "main", // Default branch
		Hash:        "0",    // For initial commit
		ExpectedHash: "0",
	}

	reqBody, err := json.Marshal(commitReq)
	if err != nil {
		return fmt.Errorf("failed to marshal Nessie commit request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/trees/main/commits", c.endpoint), bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create Nessie request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send Nessie request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Nessie API returned non-200 status: %d", resp.StatusCode)
	}

	return nil
}

// AppendToTable appends data to an existing Iceberg table in Nessie.
func (c *NessieClient) AppendToTable(ctx context.Context, tableName, newMetadataLocation string) error {
	// This is a simplified append. In a real scenario, you'd read the current
	// table state, add new data files, and then commit a new snapshot.
	// For now, we'll just update the metadata location.

	key := NessieContentKey{Elements: []string{tableName}}
	content := NessieContent{
		Type: "ICEBERG_TABLE",
		IcebergTable: NessieIcebergTable{
			MetadataLocation: newMetadataLocation,
		},
	}
	operation := NessieOperation{
		Type:    "SET_CONTENT", // Use SET_CONTENT for updates
		Key:     key,
		Content: content,
	}

	// You'd need to fetch the current hash of the branch to make a valid commit
	// For simplicity, we'll use a dummy hash for now.
	commitReq := NessieCommitRequest{
		Commit:      NessieCommit{Message: fmt.Sprintf("Append to table %s", tableName)},
		Operations:  []NessieOperation{operation},
		BranchName:  "main",
		Hash:        "dummy-hash", // This needs to be the actual current hash
		ExpectedHash: "dummy-hash",
	}

	reqBody, err := json.Marshal(commitReq)
	if err != nil {
		return fmt.Errorf("failed to marshal Nessie commit request: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/trees/main/commits", c.endpoint), bytes.NewBuffer(reqBody))
	if err != nil {
		return fmt.Errorf("failed to create Nessie request: %w", err)
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return fmt.Errorf("failed to send Nessie request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Nessie API returned non-200 status: %d", resp.StatusCode)
	}

	return nil
}
