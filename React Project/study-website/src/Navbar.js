import { Link } from "react-router-dom";

const Navbar = () => {
    return (
      <nav className="navbar">
        <h1>Gator Pal</h1>
        <img src="gator.png" alt="gator" />
        <ul>
          <li><Link to="/">Home</Link></li>
          <li><Link to="/about">About</Link></li>
          <li><Link to="/sign-in">Sign-In</Link></li>
        </ul>
          
        
      </nav>
    );
  }
   
  export default Navbar;
