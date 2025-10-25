import React, { useState } from "react";
import {
  Box,
  Paper,
  Typography,
  Grid,
  Card,
  CardContent,
  Button,
  TextField,
  Select,
  MenuItem,
  FormControl,
  InputLabel,
  Divider,
} from "@mui/material";
import { PlayArrow } from "@mui/icons-material";

const AutoML: React.FC = () => {
  const [dataset, setDataset] = useState("");
  const [targetColumn, setTargetColumn] = useState("");
  const [taskType, setTaskType] = useState("classification");

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        AutoML Experiment
      </Typography>

      <Grid container spacing={3}>
        <Grid item xs={12} md={4}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Configuration
            </Typography>

            <TextField
              fullWidth
              label="Dataset Path"
              value={dataset}
              onChange={(e) => setDataset(e.target.value)}
              sx={{ mb: 2 }}
            />

            <FormControl fullWidth sx={{ mb: 2 }}>
              <InputLabel>Task Type</InputLabel>
              <Select
                value={taskType}
                onChange={(e) => setTaskType(e.target.value)}
                label="Task Type"
              >
                <MenuItem value="classification">Classification</MenuItem>
                <MenuItem value="regression">Regression</MenuItem>
                <MenuItem value="clustering">Clustering</MenuItem>
              </Select>
            </FormControl>

            <TextField
              fullWidth
              label="Target Column"
              value={targetColumn}
              onChange={(e) => setTargetColumn(e.target.value)}
              sx={{ mb: 2 }}
            />

            <Button
              fullWidth
              variant="contained"
              startIcon={<PlayArrow />}
              size="large"
            >
              Start AutoML
            </Button>
          </Paper>
        </Grid>

        <Grid item xs={12} md={8}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Experiment Results
            </Typography>
            <Divider sx={{ mb: 3 }} />

            <Grid container spacing={2}>
              {[
                { model: "Random Forest", accuracy: 94.2, f1: 0.93 },
                { model: "XGBoost", accuracy: 95.8, f1: 0.95 },
                { model: "LightGBM", accuracy: 94.9, f1: 0.94 },
                { model: "Neural Network", accuracy: 93.5, f1: 0.92 },
              ].map((result, idx) => (
                <Grid item xs={12} sm={6} key={idx}>
                  <Card variant="outlined">
                    <CardContent>
                      <Typography variant="h6">{result.model}</Typography>
                      <Typography color="textSecondary">
                        Accuracy: {result.accuracy}%
                      </Typography>
                      <Typography color="textSecondary">
                        F1 Score: {result.f1}
                      </Typography>
                      <Button size="small" sx={{ mt: 1 }}>
                        View Details
                      </Button>
                    </CardContent>
                  </Card>
                </Grid>
              ))}
            </Grid>
          </Paper>
        </Grid>

        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Hyperparameter Tuning Progress
            </Typography>
            <Typography variant="body2" color="textSecondary">
              Best configuration found after 45/100 iterations
            </Typography>
            <Box
              sx={{
                mt: 2,
                p: 2,
                bgcolor: "background.default",
                borderRadius: 1,
              }}
            >
              <pre style={{ margin: 0 }}>
                {JSON.stringify(
                  {
                    learning_rate: 0.01,
                    max_depth: 8,
                    n_estimators: 200,
                    subsample: 0.8,
                  },
                  null,
                  2
                )}
              </pre>
            </Box>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default AutoML;
