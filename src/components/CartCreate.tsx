import React, { useEffect, useState } from "react";
import Divider from "@mui/material/Divider";
import Snackbar from "@mui/material/Snackbar";
import Container from "@mui/material/Container";
import Alert from '@mui/material/Alert';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Typography from "@mui/material/Typography";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import InputLabel from '@mui/material/InputLabel';
import Button from "@mui/material/Button";


import { EmployeeInterface } from "../interfaces/IEmployee";
import { MemberInterface } from "../interfaces/IMember";
import { ProductInterface } from "../interfaces/IProduct";
import { CartInterface } from "../interfaces/ICart";
import { OrderInterface } from "../interfaces/IOrder";

import {
  GetMembers,
  GetProducts,
  GetCarts,
  Carts,
  Orders,
  GetOrders,
} from "../services/HttpClientService";

import { parse } from "path";
import { GridPanelWrapper } from "@mui/x-data-grid";

function CartCreate(){
  const [employee, setEmployee] = useState<EmployeeInterface>();
  const [members, setMembers] = useState<MemberInterface[]>([]);
  const [products, setProducts] = useState<ProductInterface[]>([]);
  const [carts, setCarts] = useState<CartInterface>({});
  const [orders, setOrders] = useState<OrderInterface>();

  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);
  
  const apiUrl = "http://localhost:8080";
  const requestOptions = {
      method: "GET",
      headers: {
          //Authorization: `Bearer ${localStorage.getItem("token")}`,
          "Content-Type": "application/json"
      },
  };

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof CartCreate;
    const { value } = event.target;
    setCarts({ ...carts, [id]: value });
  };

  const handleClose = (
    event?: React.SyntheticEvent | Event,
    reason?: string
  ) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };
        
  const handleChange = (event: SelectChangeEvent) => {
    const name = event.target.name as keyof typeof carts;
    setCarts({
      ...carts,
      [name]: event.target.value,
    });

    // const name1 = event.target.name as keyof typeof orders;
    // setOrders({
    //   ...orders,
    //   [name]: event.target.value,
    // });
  };

  // const getMembers = async () => {
  //   let res = await GetMembers();
  //   if (res) {
  //       setMembers(res);
  //   }
  // };
  const getMembers = async () => {
    fetch(`${apiUrl}/members`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setMembers(res.data);
            }
            else { console.log("NO DATA") }
        });
};

  // const getProducts = async () => {
  //   let res = await GetProducts();
  //   if (res) {
  //     setProducts(res);
  //   }
  // };
  const getProducts = async () => {
    fetch(`${apiUrl}/products`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data)
                setProducts(res.data);
            }
            else { console.log("NO DATA") }
        });
  };


//   const getPlaylist = async () => {
//     let res = await GetPlaylistByUID();
//     watchVideo.PlaylistID = res.ID;
//     if (res) {
//       setPlaylists(res);
//     }
//   };
const getCarts = async () => {
  fetch(`${apiUrl}/carts`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
          if (res.data) {
              console.log(res.data)
              setCarts(res.data);
          }
          else { console.log("NO DATA") }
      });
};

const getOrders = async () => {
  fetch(`${apiUrl}/orders`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
          if (res.data) {
              console.log(res.data)
              setOrders(res.data);
          }
          else { console.log("NO DATA") }
      });
};
  // const getCarts = async () => {
  //   let res = await GetCarts();
  //   if (res) {
  //     setCarts(res);
  //   }
  // };

  // const getOrders = async () => {
  //   let res = await GetOrders();
  //   if (res) {
  //     setOrders(res);
  //   }
  // };
  
  
 

  useEffect(() => {
    getMembers();
    getProducts();
    getCarts();
    getOrders();
  }, []);

  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  async function submit() {
    let data = {
      Telephone: carts.Telephone?? "",
      ProducID:convertType(orders?.ProductID),
      Quantity:typeof orders?.Quantity ==="string"?parseInt(orders.Quantity):0,
    };
    console.log(data)

    let res = await Carts(data);
    if (res) {
      setSuccess(true);
    } else {
      setError(true);
    }
  }
  
  

  return (
    <Container maxWidth="md">
      <Snackbar
        open={success}
        autoHideDuration={3000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="success">
            Success
        </Alert>
      </Snackbar>
      <Snackbar
        open={error}
        autoHideDuration={6000}
        onClose={handleClose}
        anchorOrigin={{ vertical: "top", horizontal: "center" }}
      >
        <Alert onClose={handleClose} severity="error">
          Failed
        </Alert>
      </Snackbar>
      <Paper>
        <Box
          display="flex"
          sx={{
            marginTop: 2,
          }}
        >
          <Box sx={{ paddingX: 2, paddingY: 1 }}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              Shopping Cart
            </Typography>
          </Box>
        </Box>
        <Divider />
        <Grid container spacing={3} sx={{ padding: 2 }}>
          <Grid item xs={12}>
            <FormControl fullWidth >
            <InputLabel id="InputLabeltel">Telephone number</InputLabel>      
                <Select
                  required
                  labelId="labelIdtel"
                  id="idtel"
                  label="Telephone Number"
                  native
                  value={carts.MemberID + ""}
                  onChange={handleChange}
                  inputProps={{
                    name: "MemberID",
                  }}                
                >
                    <option aria-label="None" value=""></option>
                    {members.map((item: MemberInterface) => (
                        <option value={item.ID} key={item.ID}>
                        {item.Telephone}
                        </option>
                    ))}
                </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl fullWidth variant="outlined" sx={{ paddingY: 2 }}>
            <InputLabel id="InputLabeltel" sx={{ paddingY: 2 }}>Product</InputLabel>      
                <Select
                  required
                  labelId="labelIdProduct"
                  id="idProduct"
                  label="Product"
                  native
                  value={orders?.ProductID + ""}
                  onChange={handleChange}
                  inputProps={{
                    name: "ProductID",
                  }}                
                >
                  <option aria-label="None" value=""></option>
                  {products.map((item: ProductInterface) => (
                    <option value={item.ID} key={item.ID}>
                    {item.Product_name}
                    </option>
                   ))}
                </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6} >
            <FormControl fullWidth variant="outlined" >
            <TextField
                margin="normal"
                required
                fullWidth
                id="Lek"
                variant="outlined"
                type="number"
                size="medium"  
                value={orders?.Quantity || ""}
                onChange={handleInputChange}
                label="Qty"
              />
            </FormControl>
          </Grid>
          <Grid item xs={12}>
            {/* <Button
              component={RouterLink}
              to="/watch_videos"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button> */}
            <Button sx={{ padding: 2, marginX:42}}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              Add Order
            </Button>
          </Grid>
          <Grid item xs={12}>
            <FormControl fullWidth variant="outlined">
              <InputLabel id="demo-simple-select-label">Employee</InputLabel>      
              <Select
                native
                labelId="labelIdemployee"
                id="idemployee"
                label="Employee"                
                value={carts.EmployeeID + ""}
                onChange={handleChange}
                disabled
                inputProps={{
                  name: "EmployeeID",
                }}
              >
                <option aria-label="None" value=""></option>
                <option value={employee?.ID} key={employee?.ID}>
                  {employee?.FirstName} {employee?.LastName}
                </option>    
              </Select>
            </FormControl>
          </Grid>
          {/* <Grid item xs={12}>
            <p>Employee</p>
            <FormControl fullWidth variant="outlined">
              <TextField
                id="EmployeeID"
                variant="outlined"
                type="number"
                size="medium"
                placeholder="Please select a Employee."
                value={product.EmployeeID || ""}
                onChange={handleInputChange}
              />
            </FormControl>
          </Grid> */}
          <Grid item xs={12}>
            {/* <Button
              component={RouterLink}
              to="/cart/create"
              variant="contained"
              color="inherit"
            >
              กลับ
            </Button> */}
            <Button
              style={{ float: "right" }}
              onClick={submit}
              variant="contained"
              color="primary"
            >
              Check Out
            </Button>
          </Grid>
        </Grid>
      </Paper>
    </Container>
  );
}

export default CartCreate;

