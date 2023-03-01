import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import Navbar from "./Navbar";
import SignIn from "./pages/SignIn";
import './Navbar.css';
import './Home.css';
import './About.css';

const root = ReactDOM.createRoot(document.querySelector("#application")!);
root.render(
  <BrowserRouter>
    <Navbar />
    <div className="content">
      <Routes>
        <Route index element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/sign-in" element={<SignIn />} />
      </Routes>
    </div>
  </BrowserRouter>
);