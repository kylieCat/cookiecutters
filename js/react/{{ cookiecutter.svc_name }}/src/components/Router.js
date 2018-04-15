import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';
import App from './App';

class Router extends Component {
  render() {
    return (
      <Router>
        <Route exact path="/" component={App} />
      </Router>
    );
  }
}

export default Router;
