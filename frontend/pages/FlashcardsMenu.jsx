import React from "react";
import { Typography, Button, IconButton, AppBar, Card, Box, CardActions, CardActionArea, CardContent, CardMedia, CssBaseline, Grid, Toolbar, Container } from '@mui/material'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import StyleIcon from '@mui/icons-material/Style'; //https://mui.com/material-ui/material-icons/?query=Cards

export default function FlashcardsMenu() {
  return (
    <div className="home">
      <CssBaseline/>
      <Container maxWidth="md" sx={{
        minHeight: '60vh'
      }}>
        <Grid container spacing={1}>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardActionArea id="cop4600Button" component={Link} to="/flashcards/cop4600">
                <CardContent>
                  <Typography gutterBottom component="div">
                    COP4600
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardActionArea id="phy2053Button" component={Link} to="/flashcards/phy2053">
                <CardContent>
                  <Typography gutterBottom component="div">
                    PHY2053
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        </Grid>
      </Container>
    </div>
  );
}