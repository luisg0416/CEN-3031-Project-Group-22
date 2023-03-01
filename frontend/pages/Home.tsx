import React from "react";
import { Typography, Button, IconButton, AppBar, Card, Box, CardActions, CardActionArea, CardContent, CardMedia, CssBaseline, Grid, Toolbar, Container } from '@mui/material'
import StyleIcon from '@mui/icons-material/Style'; //https://mui.com/material-ui/material-icons/?query=Cards

export default function Home() {
  return (
    <div className="home">
      <CssBaseline/>
      <Container maxWidth="md" sx={{
        minHeight: '70vh',
        marginTop: 2
      }}>
        <Grid container spacing={4}>
          <Grid item>
            <Card sx={{ minWidth: 250 }}>
              <CardActionArea>
                <CardContent>
                  <CardMedia
                    component="img"
                    height="100"
                    width="50"
                    image="flashcards.png"
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
              <CardActionArea>
                <CardContent>
                  <CardMedia
                    component="img"
                    height="100"
                    width="50"
                    image="games.png"
                    alt="flashcards"
                    sx={{ padding: "2em 2em 2em 2em", objectFit: "contain" }}
                  />
                  <Typography gutterBottom component="div">
                    Games
                  </Typography>
                </CardContent>
              </CardActionArea>
            </Card>
          </Grid>
        </Grid>
      </Container>
      <Box
        component="footer"
        sx={{
          py: 3,
          px: 2,
          mt: 'auto',
          backgroundColor: (theme) =>
            theme.palette.mode === 'light'
              ? theme.palette.grey[200]
              : theme.palette.grey[800],
        }}
      >
        <Container maxWidth="sm">
          <Typography variant="body1">
            Footer
          </Typography>
        </Container>
      </Box>
    </div>
  );
}