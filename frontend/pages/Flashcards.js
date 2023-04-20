import React, {useState} from "react";
import { Typography, Button, IconButton, AppBar, Card, Box, CardActions, CardActionArea, CardContent, CardMedia, CssBaseline, Grid, Toolbar, Container } from '@mui/material'
import FlashcardList from "../FlashcardList";

function Flashcards ({props}) {
    const [flashcards, setFlashcards] = useState(props)
    return (
        <div className="home">
            <CssBaseline/>
            <FlashcardList flashcards={flashcards} />
        </div>
    );
}
export default Flashcards
