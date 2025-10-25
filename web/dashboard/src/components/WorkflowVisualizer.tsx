import React, { useEffect, useState } from "react";
import ReactFlow, {
  MiniMap,
  Controls,
  Background,
  Node,
  Edge,
} from "react-flow-renderer";
import { Box, Paper, Typography, Chip, Card, CardContent } from "@mui/material";
import {
  CheckCircle,
  Error,
  HourglassEmpty,
  Pending,
} from "@mui/icons-material";

interface WorkflowNode {
  id: string;
  name: string;
  type: string;
  status: "pending" | "running" | "succeeded" | "failed";
  startTime?: string;
  endTime?: string;
  duration?: number;
}

interface WorkflowVisualizerProps {
  workflowId: string;
}

const WorkflowVisualizer: React.FC<WorkflowVisualizerProps> = ({
  workflowId,
}) => {
  const [nodes, setNodes] = useState<Node[]>([]);
  const [edges, setEdges] = useState<Edge[]>([]);
  const [selectedNode, setSelectedNode] = useState<WorkflowNode | null>(null);

  useEffect(() => {
    // Fetch workflow data from API
    // This is mock data for demonstration
    const mockNodes: Node[] = [
      {
        id: "1",
        type: "input",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <CheckCircle color="success" fontSize="small" />
              <span>Data Ingestion</span>
            </Box>
          ),
        },
        position: { x: 250, y: 0 },
        style: {
          background: "#1a472a",
          color: "white",
          border: "2px solid #2e7d32",
        },
      },
      {
        id: "2",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <CheckCircle color="success" fontSize="small" />
              <span>Data Validation</span>
            </Box>
          ),
        },
        position: { x: 250, y: 100 },
        style: {
          background: "#1a472a",
          color: "white",
          border: "2px solid #2e7d32",
        },
      },
      {
        id: "3",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <HourglassEmpty color="info" fontSize="small" />
              <span>Feature Engineering</span>
            </Box>
          ),
        },
        position: { x: 250, y: 200 },
        style: {
          background: "#1a3a52",
          color: "white",
          border: "2px solid #1976d2",
        },
      },
      {
        id: "4",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <Pending color="disabled" fontSize="small" />
              <span>Model Training</span>
            </Box>
          ),
        },
        position: { x: 100, y: 300 },
        style: {
          background: "#424242",
          color: "white",
          border: "2px solid #757575",
        },
      },
      {
        id: "5",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <Pending color="disabled" fontSize="small" />
              <span>Model Evaluation</span>
            </Box>
          ),
        },
        position: { x: 400, y: 300 },
        style: {
          background: "#424242",
          color: "white",
          border: "2px solid #757575",
        },
      },
      {
        id: "6",
        type: "output",
        data: {
          label: (
            <Box display="flex" alignItems="center" gap={1}>
              <Pending color="disabled" fontSize="small" />
              <span>Model Deployment</span>
            </Box>
          ),
        },
        position: { x: 250, y: 400 },
        style: {
          background: "#424242",
          color: "white",
          border: "2px solid #757575",
        },
      },
    ];

    const mockEdges: Edge[] = [
      {
        id: "e1-2",
        source: "1",
        target: "2",
        animated: false,
        style: { stroke: "#2e7d32" },
      },
      {
        id: "e2-3",
        source: "2",
        target: "3",
        animated: true,
        style: { stroke: "#1976d2" },
      },
      { id: "e3-4", source: "3", target: "4", style: { stroke: "#757575" } },
      { id: "e3-5", source: "3", target: "5", style: { stroke: "#757575" } },
      { id: "e4-6", source: "4", target: "6", style: { stroke: "#757575" } },
      { id: "e5-6", source: "5", target: "6", style: { stroke: "#757575" } },
    ];

    setNodes(mockNodes);
    setEdges(mockEdges);
  }, [workflowId]);

  const handleNodeClick = (event: any, node: Node) => {
    // Fetch node details from API
    const mockNodeDetails: WorkflowNode = {
      id: node.id,
      name: typeof node.data.label === "string" ? node.data.label : "Node",
      type: "task",
      status:
        node.id === "1" || node.id === "2"
          ? "succeeded"
          : node.id === "3"
          ? "running"
          : "pending",
      startTime: new Date().toISOString(),
      duration: 120,
    };
    setSelectedNode(mockNodeDetails);
  };

  const getStatusIcon = (status: string) => {
    switch (status) {
      case "succeeded":
        return <CheckCircle color="success" />;
      case "failed":
        return <Error color="error" />;
      case "running":
        return <HourglassEmpty color="info" />;
      default:
        return <Pending color="disabled" />;
    }
  };

  return (
    <Box display="flex" gap={2} height="600px">
      <Box flex={1} border="1px solid #333">
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodeClick={handleNodeClick}
          fitView
        >
          <MiniMap />
          <Controls />
          <Background />
        </ReactFlow>
      </Box>

      {selectedNode && (
        <Box width="300px">
          <Card>
            <CardContent>
              <Box display="flex" alignItems="center" gap={1} mb={2}>
                {getStatusIcon(selectedNode.status)}
                <Typography variant="h6">{selectedNode.name}</Typography>
              </Box>

              <Typography variant="body2" color="textSecondary" gutterBottom>
                Status: <Chip label={selectedNode.status} size="small" />
              </Typography>

              {selectedNode.startTime && (
                <Typography variant="body2" color="textSecondary">
                  Start Time:{" "}
                  {new Date(selectedNode.startTime).toLocaleString()}
                </Typography>
              )}

              {selectedNode.duration && (
                <Typography variant="body2" color="textSecondary">
                  Duration: {selectedNode.duration}s
                </Typography>
              )}

              <Box mt={2}>
                <Typography variant="subtitle2">Logs:</Typography>
                <Paper
                  variant="outlined"
                  sx={{ p: 1, mt: 1, maxHeight: 200, overflow: "auto" }}
                >
                  <Typography variant="caption" component="pre">
                    {`[INFO] Starting ${selectedNode.name}...\n[INFO] Processing data...\n[INFO] Task completed successfully`}
                  </Typography>
                </Paper>
              </Box>
            </CardContent>
          </Card>
        </Box>
      )}
    </Box>
  );
};

export default WorkflowVisualizer;
