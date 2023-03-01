import React from "react";
import Navbar from "./Navbar";
import Home from "./Home";
import About from "./About";
import './Navbar.css';
import './Home.css';
import './About.css';
import { BrowserRouter as Router, Route, Switch } from "react-router-dom";
import Sign from "./sign-in";
import'./sign-in.css';

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
          <Route>
            <Sign />
          </Route>
        </Switch>
      </div>
    </div>
    </Router>
  );
}

export default App;
