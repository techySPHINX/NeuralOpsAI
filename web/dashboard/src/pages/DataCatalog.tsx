import React from "react";
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
} from "@mui/material";

const DataCatalog: React.FC = () => {
  const datasets = [
    {
      name: "customer_data",
      version: "v2.1",
      rows: 1000000,
      columns: 45,
      lastUpdated: "2025-10-25",
    },
    {
      name: "transactions",
      version: "v1.8",
      rows: 5000000,
      columns: 32,
      lastUpdated: "2025-10-24",
    },
    {
      name: "product_catalog",
      version: "v3.0",
      rows: 50000,
      columns: 28,
      lastUpdated: "2025-10-26",
    },
  ];

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Data Catalog
      </Typography>

      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Dataset Name</TableCell>
              <TableCell>Version</TableCell>
              <TableCell>Rows</TableCell>
              <TableCell>Columns</TableCell>
              <TableCell>Last Updated</TableCell>
              <TableCell>Status</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {datasets.map((dataset) => (
              <TableRow key={dataset.name}>
                <TableCell>{dataset.name}</TableCell>
                <TableCell>
                  <Chip label={dataset.version} size="small" />
                </TableCell>
                <TableCell>{dataset.rows.toLocaleString()}</TableCell>
                <TableCell>{dataset.columns}</TableCell>
                <TableCell>{dataset.lastUpdated}</TableCell>
                <TableCell>
                  <Chip label="Active" color="success" size="small" />
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default DataCatalog;
