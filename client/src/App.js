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
		const token = sessionStorage.getItem("token")
		this.state = {
			isLoggedIn: token !== null,
			token: sessionStorage.getItem("token"),
		};
	};

	setLoginState = (loginState) => {
		this.setState({isLoggedIn: loginState})
		// console.log(this.state)	
	}

	getLoginState = () => {
		return this.state.isLoggedIn
	}

	setToken = (tokenVal) => {
		sessionStorage.setItem("token", tokenVal)
		this.setState({token: tokenVal})
		// console.log(this.state)	
	}

	getToken = () => {
		return this.state.token
	}

	render() {
		return (
			<div className="root-container">
					<Router>
						<NavBar isLoggedIn={this.state.isLoggedIn} />
						<div className="app-container">
						<Switch>
							<Route exact path='/ping'>
								<Ping />
							</Route>
							<Route exact path='/authenticate'>
								<Authenticate setLoginState={this.setLoginState} setToken={this.setToken} />
							</Route>
							<Route exact path='/logout'>
								<Logout setLoginState={this.setLoginState} setToken={this.setToken} />
							</Route>
							<PrivateRoute component={Home} function={this.getLoginState} path="/home" exact />

							<Route exact path="/">
  								{
								  this.state.isLoggedIn ? <Redirect to="/home" /> : <Redirect to="/authenticate" />
								} 
							</Route>
						</Switch></div>
					</Router>
				
			</div>
		);
	}
}

const PrivateRoute = ({component: Component, function: getLoginState, ...rest}) => {
    return (

        // Show the component only when the user is logged in
        // Otherwise, redirect the user to /signin page
        <Route {...rest} render={props => (
            getLoginState() ?
                <Component {...props} />
            : <Redirect to="/authenticate" />
        )} />
    );
};



export default App;