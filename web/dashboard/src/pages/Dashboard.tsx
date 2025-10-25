import React from "react";
import {
  Grid,
  Paper,
  Typography,
  Box,
  Card,
  CardContent,
  LinearProgress,
} from "@mui/material";
import {
  TrendingUp,
  AccountTree,
  CheckCircle,
  Error,
} from "@mui/icons-material";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

const Dashboard: React.FC = () => {
  const stats = [
    {
      title: "Active Pipelines",
      value: "12",
      icon: <AccountTree />,
      color: "#3f51b5",
    },
    {
      title: "Successful Runs",
      value: "156",
      icon: <CheckCircle />,
      color: "#4caf50",
    },
    { title: "Failed Runs", value: "8", icon: <Error />, color: "#f44336" },
    {
      title: "Models Deployed",
      value: "23",
      icon: <TrendingUp />,
      color: "#ff9800",
    },
  ];

  const chartData = [
    { name: "Mon", success: 24, failed: 2 },
    { name: "Tue", success: 28, failed: 1 },
    { name: "Wed", success: 32, failed: 3 },
    { name: "Thu", success: 29, failed: 2 },
    { name: "Fri", success: 35, failed: 1 },
    { name: "Sat", success: 20, failed: 0 },
    { name: "Sun", success: 18, failed: 1 },
  ];

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Dashboard
      </Typography>

      <Grid container spacing={3}>
        {stats.map((stat) => (
          <Grid item xs={12} sm={6} md={3} key={stat.title}>
            <Card sx={{ bgcolor: stat.color, color: "white" }}>
              <CardContent>
                <Box
                  display="flex"
                  alignItems="center"
                  justifyContent="space-between"
                >
                  <Box>
                    <Typography variant="h4">{stat.value}</Typography>
                    <Typography variant="body2">{stat.title}</Typography>
                  </Box>
                  <Box sx={{ fontSize: 48, opacity: 0.7 }}>{stat.icon}</Box>
                </Box>
              </CardContent>
            </Card>
          </Grid>
        ))}

        <Grid item xs={12}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Pipeline Execution Trends
            </Typography>
            <ResponsiveContainer width="100%" height={300}>
              <LineChart data={chartData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Line
                  type="monotone"
                  dataKey="success"
                  stroke="#4caf50"
                  name="Successful"
                />
                <Line
                  type="monotone"
                  dataKey="failed"
                  stroke="#f44336"
                  name="Failed"
                />
              </LineChart>
            </ResponsiveContainer>
          </Paper>
        </Grid>

        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              Recent Pipelines
            </Typography>
            {[1, 2, 3].map((i) => (
              <Box key={i} mb={2}>
                <Typography variant="body1">Pipeline {i}</Typography>
                <LinearProgress
                  variant="determinate"
                  value={Math.random() * 100}
                />
              </Box>
            ))}
          </Paper>
        </Grid>

        <Grid item xs={12} md={6}>
          <Paper sx={{ p: 3 }}>
            <Typography variant="h6" gutterBottom>
              System Health
            </Typography>
            <Box mb={2}>
              <Typography variant="body2">CPU Usage</Typography>
              <LinearProgress
                variant="determinate"
                value={45}
                color="success"
              />
            </Box>
            <Box mb={2}>
              <Typography variant="body2">Memory Usage</Typography>
              <LinearProgress
                variant="determinate"
                value={67}
                color="warning"
              />
            </Box>
            <Box>
              <Typography variant="body2">Storage Usage</Typography>
              <LinearProgress variant="determinate" value={32} color="info" />
            </Box>
          </Paper>
        </Grid>
      </Grid>
    </Box>
  );
};

export default Dashboard;
