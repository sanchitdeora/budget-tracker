import React from 'react';
import { BrowserRouter as Router, Redirect, Route, Switch } from "react-router-dom";
import './App.scss';

// to use fontRoboto
import "@fontsource/roboto";
// importing paper and container from core
import { Paper, Container } from "@material-ui/core";

// these are for customizing the theme
import { ThemeProvider, createMuiTheme } from "@material-ui/core/styles";
/* material shell also provide the colors we can import them like these */
import { teal, orange } from "@material-ui/core/colors";
import Authenticate from './components/Authentication/Authenticate.jsx';
import Logout from './components/Authentication/Logout.jsx';
import Home from "./components/Home.jsx";
import NavBar from './components/NavBar.jsx';
import Ping from "./components/Ping";
import Survey from './components/Survey.jsx';
import { isLogin } from './utils';
import { TOKEN } from './utils/GlobalConstants';

const theme = createMuiTheme({
  typography: {
    h1: {
     /* this will change the font size for h1, we can also do 
        it for others, */
      fontSize: "3rem",
    },
  },
  palette: {
    /* this is used to turn the background dark, but we have 
        to use Paper for this inOrder to use it. */
    type: "light",
    primary: {
     // main: colorName[hue],
     // we have to import the color first to use it
      main: teal[600],
    },
    secondary: {
      main: orange[400],
    },
  },
});

class App extends React.Component<any, any> {
	constructor(props) {
		super(props);
		const token = sessionStorage.getItem(TOKEN)
		this.state = {
			isLoggedIn: token !== null,
			token: sessionStorage.getItem(TOKEN),
		};
	};

	setLoginState = (loginState: boolean) => {
		this.setState({isLoggedIn: loginState})
		// console.log(this.state)	
	}

	setToken = (tokenVal) => {
		sessionStorage.setItem(TOKEN, tokenVal)
		this.setState({token: tokenVal})
		// console.log(this.state)	
	}

	render() {
		return (
		<ThemeProvider theme={theme}>
			<Paper style={{ height: "200vh" }}>
				<Container className="root-container">
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

							<PrivateRoute component={Home} path="/home" exact />

							<Route exact path='/survey'>
								<Survey />
							</Route>

							<Route exact path="/">
								{
								this.state.isLoggedIn ? 
								<Redirect to="/home" /> : 
								<Redirect to="/authenticate" />
								} 
							</Route>
						</Switch></div>
					</Router>
				</Container>
			</Paper>
		</ThemeProvider>
		);
	}
}

const PrivateRoute = ({component: Component, ...rest}) => {
    return (

        // Show the component only when the user is logged in
        // Otherwise, redirect the user to /authenticate page
        <Route {...rest} render={props => (
            isLogin() ?
                <Component {...props} />
            : <Redirect to="/authenticate" />
        )} />
    );
};

export default App;