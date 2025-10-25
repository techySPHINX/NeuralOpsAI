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
  Button,
  IconButton,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogActions,
  TextField,
  Grid,
} from "@mui/material";
import { Add, Download, Visibility, Delete } from "@mui/icons-material";

interface Model {
  id: string;
  name: string;
  version: string;
  framework: string;
  accuracy: number;
  status: "production" | "staging" | "archived";
  createdAt: string;
}

const Models: React.FC = () => {
  const [models] = useState<Model[]>([
    {
      id: "1",
      name: "Customer Churn Predictor",
      version: "v1.2.3",
      framework: "TensorFlow",
      accuracy: 94.5,
      status: "production",
      createdAt: "2025-10-20",
    },
    {
      id: "2",
      name: "Recommendation Engine",
      version: "v2.1.0",
      framework: "PyTorch",
      accuracy: 91.2,
      status: "staging",
      createdAt: "2025-10-22",
    },
    {
      id: "3",
      name: "Fraud Detection",
      version: "v1.0.5",
      framework: "Scikit-learn",
      accuracy: 96.8,
      status: "production",
      createdAt: "2025-10-18",
    },
  ]);

  const [openDialog, setOpenDialog] = useState(false);

  const getStatusColor = (status: string) => {
    switch (status) {
      case "production":
        return "success";
      case "staging":
        return "warning";
      case "archived":
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
        <Typography variant="h4">Model Registry</Typography>
        <Button
          variant="contained"
          startIcon={<Add />}
          onClick={() => setOpenDialog(true)}
        >
          Register Model
        </Button>
      </Box>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell>Version</TableCell>
              <TableCell>Framework</TableCell>
              <TableCell>Accuracy</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Created</TableCell>
              <TableCell>Actions</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {models.map((model) => (
              <TableRow key={model.id}>
                <TableCell>{model.name}</TableCell>
                <TableCell>{model.version}</TableCell>
                <TableCell>{model.framework}</TableCell>
                <TableCell>{model.accuracy}%</TableCell>
                <TableCell>
                  <Chip
                    label={model.status}
                    color={getStatusColor(model.status) as any}
                    size="small"
                  />
                </TableCell>
                <TableCell>{model.createdAt}</TableCell>
                <TableCell>
                  <IconButton size="small" color="primary">
                    <Visibility />
                  </IconButton>
                  <IconButton size="small" color="info">
                    <Download />
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

      <Dialog
        open={openDialog}
        onClose={() => setOpenDialog(false)}
        maxWidth="md"
        fullWidth
      >
        <DialogTitle>Register New Model</DialogTitle>
        <DialogContent>
          <Grid container spacing={2} sx={{ mt: 1 }}>
            <Grid item xs={12}>
              <TextField fullWidth label="Model Name" />
            </Grid>
            <Grid item xs={6}>
              <TextField fullWidth label="Version" />
            </Grid>
            <Grid item xs={6}>
              <TextField fullWidth label="Framework" />
            </Grid>
            <Grid item xs={12}>
              <TextField fullWidth label="Model Path/URL" />
            </Grid>
            <Grid item xs={12}>
              <TextField fullWidth label="Description" multiline rows={3} />
            </Grid>
          </Grid>
        </DialogContent>
        <DialogActions>
          <Button onClick={() => setOpenDialog(false)}>Cancel</Button>
          <Button variant="contained" onClick={() => setOpenDialog(false)}>
            Register
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default Models;
