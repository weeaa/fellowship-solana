import React, { useState, useEffect } from 'react';
import { Button, TextField, Card, CardContent, CardHeader, Typography, Grid } from '@mui/material';
import { Connection, PublicKey, clusterApiUrl } from '@solana/web3.js';
import { encodeURL, createQR, findReference, validateTransfer } from '@solana/pay';
import BigNumber from 'bignumber.js';

const MERCHANT_WALLET = new PublicKey('EcrHvqa5Vh4NhR3bitRZVrdcUGr1Z3o6bXHz7xgBU2FB');
const NETWORK = clusterApiUrl('devnet');

const SolanaPOS = () => {
  const [products, setProducts] = useState([]);
  const [newProduct, setNewProduct] = useState({ name: '', price: '' });
  const [cart, setCart] = useState([]);
  const [paymentURL, setPaymentURL] = useState(null);
  const [paymentConfirmed, setPaymentConfirmed] = useState(false);
  const [qrCode, setQrCode] = useState(null);

  const connection = new Connection(NETWORK);

  const addProduct = () => {
    if (newProduct.name && newProduct.price) {
      setProducts([...products, { ...newProduct, id: Date.now() }]);
      setNewProduct({ name: '', price: '' });
    }
  };

  const addToCart = (product) => {
    setCart([...cart, product]);
  };

  const removeFromCart = (index) => {
    setCart(cart.filter((_, i) => i !== index));
  };

  const checkout = async () => {
    const total = cart.reduce((sum, item) => sum + parseFloat(item.price), 0);
    const totalInSol = new BigNumber(total);

    const paymentLink = encodeURL({
      recipient: MERCHANT_WALLET,
      amount: totalInSol,
      label: 'Solana POS Checkout',
      message: 'Payment for POS items',
    });

    setPaymentURL(paymentLink);

    const qr = createQR(paymentLink, 256);
    const qrCodeContainer = document.getElementById('qr-code');
    qrCodeContainer.innerHTML = '';
    qr.append(qrCodeContainer);

    await monitorPayment(totalInSol);
  };

  const monitorPayment = async (amount) => {
    const reference = new PublicKey(MERCHANT_WALLET);
    let isConfirmed = false;

    while (!isConfirmed) {
      try {
        const signature = await findReference(connection, reference);
        await validateTransfer(connection, signature, {
          recipient: MERCHANT_WALLET,
          amount,
        });
        isConfirmed = true;
        setPaymentConfirmed(true);
        setCart([]);
      } catch (error) {
        await new Promise((resolve) => setTimeout(resolve, 2000));
      }
    }
  };

  return (
      <div style={{ padding: '30px', maxWidth: '900px', margin: 'auto' }}>
        <Typography variant="h3" align="center" gutterBottom style={{ marginBottom: '30px', fontWeight: 'bold' }}>
          Solana POS System
        </Typography>

        <Grid container spacing={4}>
          <Grid item xs={12} md={6}>
            <Card variant="outlined" style={{ marginBottom: '20px', boxShadow: '0 4px 12px rgba(0,0,0,0.1)' }}>
              <CardHeader title="Add New Product" style={{ backgroundColor: '#f5f5f5', padding: '16px' }} />
              <CardContent>
                <TextField
                    fullWidth
                    label="Product Name"
                    variant="outlined"
                    value={newProduct.name}
                    onChange={(e) => setNewProduct({ ...newProduct, name: e.target.value })}
                    margin="normal"
                />
                <TextField
                    fullWidth
                    label="Price (in SOL)"
                    variant="outlined"
                    type="number"
                    value={newProduct.price}
                    onChange={(e) => setNewProduct({ ...newProduct, price: e.target.value })}
                    margin="normal"
                />
                <Button
                    variant="contained"
                    color="primary"
                    onClick={addProduct}
                    style={{ marginTop: '15px', padding: '12px 20px', borderRadius: '8px' }}
                >
                  Add Product
                </Button>
              </CardContent>
            </Card>
          </Grid>

          <Grid item xs={12} md={6}>
            <Card variant="outlined" style={{ marginBottom: '20px', boxShadow: '0 4px 12px rgba(0,0,0,0.1)' }}>
              <CardHeader title="Products" style={{ backgroundColor: '#f5f5f5', padding: '16px' }} />
              <CardContent>
                {products.length > 0 ? (
                    products.map((product) => (
                        <div key={product.id} style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '10px' }}>
                          <Typography>{product.name} - ◎{product.price}</Typography>
                          <Button
                              variant="contained"
                              onClick={() => addToCart(product)}
                              style={{ padding: '6px 12px', borderRadius: '8px' }}
                          >
                            Add to Cart
                          </Button>
                        </div>
                    ))
                ) : (
                    <Typography variant="body1" color="textSecondary">No products added.</Typography>
                )}
              </CardContent>
            </Card>
          </Grid>
        </Grid>

        <Card variant="outlined" style={{ marginBottom: '20px', boxShadow: '0 4px 12px rgba(0,0,0,0.1)' }}>
          <CardHeader title="Cart" style={{ backgroundColor: '#f5f5f5', padding: '16px' }} />
          <CardContent>
            {cart.length > 0 ? (
                cart.map((item, index) => (
                    <div key={index} style={{ display: 'flex', justifyContent: 'space-between', marginBottom: '10px' }}>
                      <Typography>{item.name} - ◎{item.price}</Typography>
                      <Button
                          variant="outlined"
                          color="error"
                          onClick={() => removeFromCart(index)}
                          style={{ padding: '6px 12px', borderRadius: '8px' }}
                      >
                        Remove
                      </Button>
                    </div>
                ))
            ) : (
                <Typography variant="body1" color="textSecondary">Your cart is empty.</Typography>
            )}
            <Typography variant="h6" style={{ fontWeight: 'bold', marginTop: '20px' }}>
              Total: ◎{cart.reduce((sum, item) => sum + parseFloat(item.price), 0).toFixed(2)}
            </Typography>
            <Button
                variant="contained"
                color="success"
                onClick={checkout}
                style={{ marginTop: '20px', padding: '12px 20px', borderRadius: '8px' }}
                disabled={cart.length === 0}
            >
              Checkout with Solana Pay
            </Button>
          </CardContent>
        </Card>

        {paymentURL && (
            <div style={{ textAlign: 'center', margin: '20px 0' }}>
              <Typography variant="h5" style={{ marginBottom: '10px' }}>Scan to Pay with Solana</Typography>
              <div id="qr-code" style={{ margin: '0 auto', padding: '10px', border: '1px solid #ddd', width: 'fit-content' }}></div> {/* QR code will be generated here */}
            </div>
        )}

        {paymentConfirmed && (
            <Card variant="outlined" style={{ marginTop: '20px', boxShadow: '0 4px 12px rgba(0,0,0,0.1)' }}>
              <CardHeader title="Payment Confirmation" style={{ backgroundColor: '#e0ffe0', padding: '16px' }} />
              <CardContent>
                <Typography style={{ color: '#388e3c', fontWeight: 'bold' }}>Payment successful! Thank you for your purchase.</Typography>
              </CardContent>
            </Card>
        )}
      </div>
  );
};

export default SolanaPOS;
