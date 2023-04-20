import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import Navbar from "./Navbar";
import SignIn from "./pages/SignIn";
import FlashcardsMenu from "./pages/FlashcardsMenu";
import Flashcards from "./pages/Flashcards";
import QuizzesMenu from "./pages/QuizzesMenu";
import CreateAccount from "./pages/CreateAccount"
import Quiz from "./pages/Quiz";
import './Navbar.css';
import './Home.css';
import './About.css';
import './Quiz.css';

const root = ReactDOM.createRoot(document.querySelector("#application")!);
root.render(
  <BrowserRouter>
    <Navbar />
    <div className="content">
      <Routes>
        <Route index element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/sign-in" element={<SignIn />} />
        <Route path="/sign-in/create-acc" element={<CreateAccount />} />
        <Route path="/flashcards" element={<FlashcardsMenu/>} />
        <Route path="/flashcards/cop4600" element={<Flashcards/>} />
        <Route path="/flashcards/phy2053" element={<Flashcards/>} />
        <Route path="/quizzes" element={<QuizzesMenu/>} />
        <Route path="/quizzes/cop4600" element={<Quiz/>} />
        <Route path="/quizzes/phy2053" element={<Quiz/>} />
      </Routes>
    </div>
  </BrowserRouter>
);