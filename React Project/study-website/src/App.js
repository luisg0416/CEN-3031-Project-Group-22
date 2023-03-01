import React from "react";
import Navbar from "./Navbar";
import Home from "./Home";
import About from "./About";
import './Navbar.css';
import './Home.css';
import './About.ss'
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";

function App() {
  return (
    <Router>
    <div className='App'>
      <Navbar />
      <div classname="content">
        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
          <Route exact path="/about">
            <About />
          </Route>
        </Switch>
      </div>
    </div>
    </Router>
  );
}

export default App;
