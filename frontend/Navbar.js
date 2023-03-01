import React from "react"
import { Link } from "react-router-dom";

const Navbar = () => {
    return (
        <nav className="navbar">
            <h1>Gator Pal</h1>
            <img src="gator.png" alt="gator" />
            <div className="links">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/sign-in">Sign-in</Link>
            </div>
        </nav>
    );
  }
   
  export default Navbar;
