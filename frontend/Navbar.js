import React from "react"
import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav className="navbar">
            <h1>Gator Pal</h1>
            <img src="gator.png" alt="gator" />
            <div className="links">
                <Link id="home" to="/">Home</Link>
                <Link id="about" to="/about">About</Link>
                <Link id="sign-in" to="/sign-in">Sign-in</Link>
            </div>
        </nav>
    );
  }
   
  export default Navbar;
