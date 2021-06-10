import './App.scss';
import React from 'react';
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
		};
		localStorage.clear()
	};

	setLoginState = (childData) => {
		this.setState({isLoggedIn: childData})
		console.log(this.state)	
	}

	render() {

		return (
			<div className="root-container">
				<NavBar />
				<div className="app-container">
					<Router>
						<Switch>
							<Route exact path='/ping'>
								<Ping />
							</Route>
							<Route exact path='/authenticate'>
								<Authenticate setLoginState={this.setLoginState} />
							</Route>
							<Route exact path='/logout'>
								<Logout />
							</Route>
							<Route exact path='/home'>
								<Home />
							</Route>
							<Route exact path="/">
  								{
								  localStorage.getItem("isLoggedIn") ?  <Redirect to="/home" /> : <Redirect to="/authenticate" />
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