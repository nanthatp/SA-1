import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { OrderInterface } from "../interfaces/IOrder";
import { GetOrders } from "../services/HttpClientService";

function Orders() {
  const [orders, setOrders] = useState<OrderInterface[]>([]);

  useEffect(() => {
    getOrders();
  }, []);

  const getOrders = async () => {
    let res = await GetOrders();
    if (res) {
      setOrders(res);
    } 
  };

  const columns: GridColDef[] = [
    { field: "ID", headerName: "ID", width: 50 },
    {
      field: "Product",
      headerName: "Product",
      width: 150,
      valueFormatter: (params) => params.value.Product_name,
    },
    {
      field: "Lek",
      headerName: "Lek",
      width: 250,
      valueFormatter: (params) => params.value.Qty,
    },
    {
      field: "Cart",
      headerName: "Cart",
      width: 150,
      valueFormatter: (params) => params.value.ID,
    },
  ];

  return (
    <div>
      <Container maxWidth="md">
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Order Information
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/cart/create"
              variant="contained"
              color="primary"
            >
              CREATE ORDER
            </Button>
          </Box>
        </Box>
        <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
          <DataGrid
            rows={orders}
            getRowId={(row) => row.ID}
            columns={columns}
            pageSize={5}
            rowsPerPageOptions={[5]}
          />
        </div>
      </Container>
    </div>
  );
}

export default Orders;