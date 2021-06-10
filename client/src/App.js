import './App.scss';
import React, { useState } from 'react';
import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import Home from "./components/Home"
import Ping from "./components/Ping"
import NavBar from './components/NavBar.jsx';
import Logout from './components/Logout.jsx';
import Authenticate from './components/Authenticate.jsx';

class App extends React.Component<any, any> {
	constructor(props) {
		super(props);
		this.state = {
			isLoggedIn: false,
			token: "",
		};
	};

	setLoginState = (loginState) => {
		this.setState({isLoggedIn: loginState})
		console.log(this.state)	
	}

	setToken = (tokenVal) => {
		this.setState({token: tokenVal})
		console.log(this.state)	
	}

	render() {
		return (
			<div className="root-container">
				<NavBar isLoggedIn={this.state.isLoggedIn} />
				<div className="app-container">
					<Router>
						<Switch>
							<Route exact path='/ping'>
								<Ping />
							</Route>
							<Route exact path='/authenticate'>
								<Authenticate setLoginState={this.setLoginState} setToken={this.setToken} />
							</Route>
							<Route exact path='/logout'>
								<Logout />
							</Route>
							<Route exact path='/home'>
								<Home />
							</Route>
							<Route exact path="/">
  								{
								  this.state.isLoggedIn ? <Redirect to="/home" /> : <Redirect to="/authenticate" />
								} 
							</Route>
						</Switch>
					</Router>
				</div>
			</div>
		);
	}
}

export default App;