import React, { useState, useCallback } from "react";
import {
  Box,
  Paper,
  Typography,
  TextField,
  Button,
  Grid,
  Divider,
} from "@mui/material";
import { Save, PlayArrow } from "@mui/icons-material";
import ReactFlow, {
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  addEdge,
  Connection,
  Edge,
  Node,
} from "react-flow-renderer";

const initialNodes: Node[] = [
  {
    id: "1",
    type: "input",
    data: { label: "Data Source" },
    position: { x: 250, y: 5 },
  },
  {
    id: "2",
    data: { label: "Preprocessing" },
    position: { x: 250, y: 100 },
  },
  {
    id: "3",
    data: { label: "Feature Engineering" },
    position: { x: 250, y: 200 },
  },
  {
    id: "4",
    data: { label: "Model Training" },
    position: { x: 100, y: 300 },
  },
  {
    id: "5",
    data: { label: "Model Evaluation" },
    position: { x: 400, y: 300 },
  },
  {
    id: "6",
    type: "output",
    data: { label: "Model Deployment" },
    position: { x: 250, y: 400 },
  },
];

const initialEdges: Edge[] = [
  { id: "e1-2", source: "1", target: "2", animated: true },
  { id: "e2-3", source: "2", target: "3" },
  { id: "e3-4", source: "3", target: "4" },
  { id: "e3-5", source: "3", target: "5" },
  { id: "e4-6", source: "4", target: "6" },
  { id: "e5-6", source: "5", target: "6" },
];

const PipelineBuilder: React.FC = () => {
  const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);
  const [pipelineName, setPipelineName] = useState("");
  const [nlQuery, setNlQuery] = useState("");

  const onConnect = useCallback(
    (params: Connection | Edge) => setEdges((eds) => addEdge(params, eds)),
    [setEdges]
  );

  const handleNLSubmit = () => {
    // TODO: Call API to generate pipeline from natural language
    console.log("Natural language query:", nlQuery);
  };

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Pipeline Builder
      </Typography>

      <Grid container spacing={3}>
        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Natural Language Pipeline Creation
            </Typography>
            <Box display="flex" gap={2}>
              <TextField
                fullWidth
                placeholder="Describe your pipeline in natural language, e.g., 'Create a pipeline that ingests data from S3, preprocesses it, trains a random forest model, and deploys it'"
                value={nlQuery}
                onChange={(e) => setNlQuery(e.target.value)}
                multiline
                rows={2}
              />
              <Button variant="contained" onClick={handleNLSubmit}>
                Generate
              </Button>
            </Box>
          </Paper>
        </Grid>

        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Box
              display="flex"
              justifyContent="space-between"
              alignItems="center"
              mb={2}
            >
              <TextField
                label="Pipeline Name"
                value={pipelineName}
                onChange={(e) => setPipelineName(e.target.value)}
                sx={{ width: "50%" }}
              />
              <Box>
                <Button variant="outlined" startIcon={<Save />} sx={{ mr: 1 }}>
                  Save
                </Button>
                <Button variant="contained" startIcon={<PlayArrow />}>
                  Run Pipeline
                </Button>
              </Box>
            </Box>

            <Divider sx={{ my: 2 }} />

            <Box sx={{ height: 500, border: "1px solid #333" }}>
              <ReactFlow
                nodes={nodes}
                edges={edges}
                onNodesChange={onNodesChange}
                onEdgesChange={onEdgesChange}
                onConnect={onConnect}
                fitView
              >
                <MiniMap />
                <Controls />
                <Background />
              </ReactFlow>
            </Box>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default PipelineBuilder;
