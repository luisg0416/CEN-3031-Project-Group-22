import React from "react"
import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav className="navbar">
        <h1>Gator Pal</h1>
        <img src="gator.png" alt="gator" />
        <ul>
          <li><Link id="home"  to="/">Home</Link></li>
          <li><Link id="about" to="/about">About</Link></li>
          <li><Link id="sign-in" to="/sign-in">Sign-In</Link></li>
        </ul>
      </nav>
    );
  }
   
  export default Navbar;
