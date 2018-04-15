import React from 'react';
import ReactDOM from 'react-dom';
import Header from './components/Header';
import Footer from './components/Footer';
import Router from './components/Router';
import App from './components/App';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Header />, document.getElementById('app-header'));
ReactDOM.render(<App />, document.getElementById('app-main'));
ReactDOM.render(<Footer />, document.getElementById('app-footer'));
registerServiceWorker();
