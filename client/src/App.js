import './App.css';
import React from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Home from "./components/Home"
import Ping from "./components/Ping"
import NavBar from './components/NavBar.jsx';
import Authenticate from './components/Authenticate.jsx';

class App extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			isLoginOpen: true,
			isRegisterOpen: false
		};
	};

	render() {
		return (
			<div className="root-container">
				<NavBar />
				<div className="app-container">
					<Router>
						<Switch>
							<Route exact path='/'>
								<Home />
							</Route>
							<Route exact path='/ping'>
								<Ping />
							</Route>
							<Route exact path='/authenticate'>
								<Authenticate />
							</Route>
						</Switch>
					</Router>
				</div>
			</div>
		);
	}
}

export default App;