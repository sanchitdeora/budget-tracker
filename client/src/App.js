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
import Account from "./components/Account.jsx";
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
			isLoggedIn: Boolean(token),
			token: token,
			userInfo: {
				userID: "",
				email: "",
				name: "",
				surveyComplete: false,
			}
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

	setUser = (userBody) => {
		const user = {
			userID: userBody.userId,
			name: userBody.name,
			email: userBody.email,
			surveyComplete: userBody.surveyComplete,
		}
		this.setState({userInfo: user})
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
								
								<Route exact path='/auth'>
									<Authenticate setLoginState={this.setLoginState} setToken={this.setToken} setUser={this.setUser} />
								</Route>
								
								<Route exact path='/logout'>
									<Logout setLoginState={this.setLoginState} setToken={this.setToken} />
								</Route>

								<PrivateRoute component={Home} props={this.state} path="/home" exact />
								<PrivateRoute component={Account} props={this.state.userInfo} path="/myaccount" exact />

								<Route exact path='/survey'>
								{
									(this.state.isLoggedIn && !this.state.userInfo.SurveyComplete) ? 
									<Survey /> :
									<Redirect to="/home" />								
								}
								</Route>

								<Route exact path="/">
								{
									this.state.isLoggedIn ? 
									<Redirect to="/home" /> : 
									<Redirect to="/auth" />
								}
								</Route>
							</Switch>
						</div>
					</Router>
				</Container>
			</Paper>
		</ThemeProvider>
		);
	}
}

const PrivateRoute = ({component: Component, props: Props, ...rest}) => {
    return (

        // Show the component only when the user is logged in
        // Otherwise, redirect the user to /auth page
        <Route {...rest} render={props => (
            isLogin() ?
                <Component {...Props} />
            : <Redirect to="/auth" />
        )} />
    );
};

export default App;