import React from "react";
import { Routes, Route } from "react-router-dom";
import { Box } from "@mui/material";
import Layout from "./components/Layout";
import Dashboard from "./pages/Dashboard";
import Pipelines from "./pages/Pipelines";
import PipelineBuilder from "./pages/PipelineBuilder";
import Models from "./pages/Models";
import Workflows from "./pages/Workflows";
import DataCatalog from "./pages/DataCatalog";
import AutoML from "./pages/AutoML";

const App: React.FC = () => {
  return (
    <Box sx={{ display: "flex", minHeight: "100vh" }}>
      <Layout>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/pipelines" element={<Pipelines />} />
          <Route path="/pipelines/new" element={<PipelineBuilder />} />
          <Route path="/pipelines/:id" element={<PipelineBuilder />} />
          <Route path="/models" element={<Models />} />
          <Route path="/workflows" element={<Workflows />} />
          <Route path="/data" element={<DataCatalog />} />
          <Route path="/automl" element={<AutoML />} />
        </Routes>
      </Layout>
    </Box>
  );
};

export default App;
