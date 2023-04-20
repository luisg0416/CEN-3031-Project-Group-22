import React from "react";
import { Typography, Button, IconButton, AppBar, Card, Box, CardActions, CardActionArea, CardContent, CardMedia, CssBaseline, Grid, Toolbar, Container } from '@mui/material'
import { BrowserRouter as Router, Route, Link } from "react-router-dom";
import StyleIcon from '@mui/icons-material/Style'; //https://mui.com/material-ui/material-icons/?query=Cards

export default function Home() {
  return (
    <div className="home">
      <CssBaseline/>
      <Container maxWidth="md" sx={{
        minHeight: '60vh'
      }}>
        <Grid container spacing={1}>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardActionArea id="flashcardsButton" component={Link} to="/flashcards">
                <CardContent>
                  <CardMedia
                    component="img"
                    height="100"
                    width="50"
                    image="matching.png"
                    alt="flashcards"
                    sx={{ padding: "2em 2em 2em 2em", objectFit: "contain" }}
                  />
                  <Typography gutterBottom component="div">
                    Flash Cards
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardActionArea id="quizesButton" component={Link} to="/quizes">
                <CardContent>
                  <CardMedia
                    component="img"
                    height="100"
                    width="50"
                    image="quizes.png"
                    alt="flashcards"
                    sx={{ padding: "2em 2em 2em 2em", objectFit: "contain" }}
                  />
                  <Typography gutterBottom component="div">
                    Quizes
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        </Grid>
      </Container>
      <Box component="footer" sx={{
          py: 3,
          px: 2,
          position: 'fixed',
    bottom: 0,
    width: '100%',
          mt: 'auto',
          backgroundColor: (theme) =>
            theme.palette.mode === 'light'
              ? theme.palette.grey[200]
              : theme.palette.grey[800],
        }}>
        <Container maxWidth="sm">
          <Typography variant="body1" color="initial">
            Copyright ©2023. LLC Limited
          </Typography>
        </Container>
      </Box>
    </div>
  );
}