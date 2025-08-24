package main

import (
	"context"
	"fmt"

	"neuralops/api/proto/iceberg/v1"
	"neuralops/pkg/clients"
	"neuralops/pkg/logging"
)

type IcebergGRPCServer struct {
	icebergv1.UnimplementedIcebergServiceServer
	logger     *logging.Logger
	minioClient *clients.MinIOClient
	nessieClient *clients.NessieClient
}

func NewIcebergGRPCServer(logger *logging.Logger, minioClient *clients.MinIOClient, nessieClient *clients.NessieClient) *IcebergGRPCServer {
	return &IcebergGRPCServer{
		logger:     logger,
		minioClient: minioClient,
		nessieClient: nessieClient,
	}
}

func (s *IcebergGRPCServer) CreateTable(ctx context.Context, req *icebergv1.CreateTableRequest) (*icebergv1.CreateTableResponse, error) {
	s.logger.Info("Received CreateTable request", "table_name", req.TableName)

	// For simplicity, assume metadata location is in MinIO
	metadataLocation := fmt.Sprintf("s3://%s/iceberg/%s/metadata", s.minioClient.bucket, req.TableName)

	err := s.nessieClient.CreateTable(ctx, req.TableName, metadataLocation)
	if err != nil {
		s.logger.Error("failed to create table in Nessie", "error", err)
		return &icebergv1.CreateTableResponse{Success: false}, fmt.Errorf("failed to create table: %w", err)
	}

	s.logger.Info("Table created successfully in Nessie", "table_name", req.TableName)
	return &icebergv1.CreateTableResponse{Success: true}, nil
}

func (s *IcebergGRPCServer) AppendToTable(ctx context.Context, req *icebergv1.AppendToTableRequest) (*icebergv1.AppendToTableResponse, error) {
	s.logger.Info("Received AppendToTable request", "table_name", req.TableName, "data_path", req.DataPath)

	// In a real scenario, you'd read the current table state from Nessie,
	// add new data files (from req.DataPath, which is a MinIO path),
	// and then commit a new snapshot to Nessie.
	// For this stub, we'll just simulate an update to the metadata location.
	newMetadataLocation := fmt.Sprintf("s3://%s/iceberg/%s/metadata-new", s.minioClient.bucket, req.TableName)

	err := s.nessieClient.AppendToTable(ctx, req.TableName, newMetadataLocation)
	if err != nil {
		s.logger.Error("failed to append to table in Nessie", "error", err)
		return &icebergv1.AppendToTableResponse{Success: false}, fmt.Errorf("failed to append to table: %w", err)
	}

	s.logger.Info("Append operation simulated successfully in Nessie", "table_name", req.TableName)
	return &icebergv1.AppendToTableResponse{Success: true}, nil
}

func (s *IcebergGRPCServer) OptimizeTable(ctx context.Context, req *icebergv1.OptimizeTableRequest) (*icebergv1.OptimizeTableResponse, error) {
	s.logger.Info("Received OptimizeTable request", "table_name", req.TableName)
	// TODO: Implement optimize logic with Nessie and MinIO (e.g., triggering compaction)
	return &icebergv1.OptimizeTableResponse{Success: true}, nil
}
