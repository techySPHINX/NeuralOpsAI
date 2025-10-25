import React, { useState } from "react";
import {
  Box,
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
  Dialog,
  DialogTitle,
  DialogContent,
} from "@mui/material";
import { Visibility } from "@mui/icons-material";
import WorkflowVisualizer from "../components/WorkflowVisualizer";

interface WorkflowRun {
  id: string;
  name: string;
  status: "running" | "succeeded" | "failed" | "pending";
  startTime: string;
  duration?: string;
}

const Workflows: React.FC = () => {
  const [workflows] = useState<WorkflowRun[]>([
    {
      id: "1",
      name: "ML Training Pipeline",
      status: "running",
      startTime: "10:30 AM",
      duration: "15m",
    },
    {
      id: "2",
      name: "Data Ingestion",
      status: "succeeded",
      startTime: "9:00 AM",
      duration: "5m",
    },
    {
      id: "3",
      name: "Model Deployment",
      status: "succeeded",
      startTime: "8:15 AM",
      duration: "3m",
    },
    {
      id: "4",
      name: "Feature Processing",
      status: "failed",
      startTime: "7:45 AM",
      duration: "2m",
    },
  ]);

  const [selectedWorkflow, setSelectedWorkflow] = useState<string | null>(null);

  const getStatusColor = (status: string) => {
    switch (status) {
      case "running":
        return "info";
      case "succeeded":
        return "success";
      case "failed":
        return "error";
      case "pending":
        return "default";
      default:
        return "default";
    }
  };

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Workflow Execution Monitor
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Workflow Name</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Start Time</TableCell>
              <TableCell>Duration</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {workflows.map((workflow) => (
              <TableRow key={workflow.id}>
                <TableCell>{workflow.name}</TableCell>
                <TableCell>
                  <Chip
                    label={workflow.status}
                    color={getStatusColor(workflow.status) as any}
                    size="small"
                  />
                </TableCell>
                <TableCell>{workflow.startTime}</TableCell>
                <TableCell>{workflow.duration || "-"}</TableCell>
                <TableCell>
                  <IconButton
                    size="small"
                    color="primary"
                    onClick={() => setSelectedWorkflow(workflow.id)}
                  >
                    <Visibility />
                  </IconButton>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>

      <Dialog
        open={!!selectedWorkflow}
        onClose={() => setSelectedWorkflow(null)}
        maxWidth="xl"
        fullWidth
      >
        <DialogTitle>Workflow Visualization</DialogTitle>
        <DialogContent>
          {selectedWorkflow && (
            <WorkflowVisualizer workflowId={selectedWorkflow} />
          )}
        </DialogContent>
      </Dialog>
    </Box>
  );
};

export default Workflows;
