import React, { useState } from "react";
import {
  Box,
  Button,
  Paper,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Chip,
  IconButton,
} from "@mui/material";
import { Add, PlayArrow, Visibility, Delete } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";

interface Pipeline {
  id: string;
  name: string;
  status: "active" | "inactive" | "running";
  lastRun: string;
  success: number;
}

const Pipelines: React.FC = () => {
  const navigate = useNavigate();
  const [pipelines] = useState<Pipeline[]>([
    {
      id: "1",
      name: "Data Ingestion Pipeline",
      status: "active",
      lastRun: "2 hours ago",
      success: 98,
    },
    {
      id: "2",
      name: "Model Training Pipeline",
      status: "running",
      lastRun: "Running",
      success: 95,
    },
    {
      id: "3",
      name: "Feature Engineering",
      status: "active",
      lastRun: "1 day ago",
      success: 99,
    },
    {
      id: "4",
      name: "Batch Prediction",
      status: "inactive",
      lastRun: "3 days ago",
      success: 92,
    },
  ]);

  const getStatusColor = (status: string) => {
    switch (status) {
      case "active":
        return "success";
      case "running":
        return "info";
      case "inactive":
        return "default";
      default:
        return "default";
    }
  };

  return (
    <Box>
      <Box
        display="flex"
        justifyContent="space-between"
        alignItems="center"
        mb={3}
      >
        <Typography variant="h4">Pipelines</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => navigate("/pipelines/new")}
        >
          Create Pipeline
        </Button>
      </Box>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Last Run</TableCell>
              <TableCell>Success Rate</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {pipelines.map((pipeline) => (
              <TableRow key={pipeline.id}>
                <TableCell>{pipeline.name}</TableCell>
                <TableCell>
                  <Chip
                    label={pipeline.status}
                    color={getStatusColor(pipeline.status) as any}
                    size="small"
                  />
                </TableCell>
                <TableCell>{pipeline.lastRun}</TableCell>
                <TableCell>{pipeline.success}%</TableCell>
                <TableCell>
                  <IconButton size="small" color="primary">
                    <PlayArrow />
                  </IconButton>
                  <IconButton
                    size="small"
                    onClick={() => navigate(`/pipelines/${pipeline.id}`)}
                  >
                    <Visibility />
                  </IconButton>
                  <IconButton size="small" color="error">
                    <Delete />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default Pipelines;
