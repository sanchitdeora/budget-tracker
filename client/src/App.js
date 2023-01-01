import './App.scss';
import React from 'react';
import { TOKEN } from './utils/GlobalConstants'
import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";
import Home from "./components/Home/Home.jsx"
import Ping from "./components/Ping/Ping"
import NavBar from './components/NavBar/NavBar.jsx';
import Logout from './components/Authentication/Logout.jsx';
import Authenticate from './components/Authentication/Authenticate.jsx';
import Survey from './components/Survey/Survey.jsx';
import Transactions from './components/Transactions/Transactions';
import Bills from './components/Bills/Bills';
import Budgets from './components/Budgets/Budgets';
import Goals from './components/Goals/Goals';

// const TOKEN = "token"

class App extends React.Component {
    constructor(props) {
        super(props);
        const token = sessionStorage.getItem(TOKEN)
        this.state = {
            isLoggedIn: token !== null,
            token: sessionStorage.getItem(TOKEN),
        };
    };

    setLoginState = (loginState) => {
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
            <div className="root-container">
                    <Router>
                        <NavBar isLoggedIn={this.state.isLoggedIn} />
                        <div className="app-container">
                        <Switch>
                            
                            <PrivateRoute component={Home} path="/home" exact />

                            <Route exact path='/ping'>
                                <Ping />
                            </Route>
                            
                            <Route exact path='/authenticate'>
                                <Authenticate setLoginState={this.setLoginState} setToken={this.setToken} />
                            </Route>
                            
                            <Route exact path='/logout'>
                                <Logout setLoginState={this.setLoginState} setToken={this.setToken} />
                            </Route>

                            <Route exact path='/survey'>
                                <Survey />
                            </Route>

                            <Route exact path='/transactions'>
                                <Transactions />
                            </Route>

                            <Route exact path='/bills'>
                                <Bills />
                            </Route>

                            <Route exact path='/budgets'>
                                <Budgets />
                            </Route>

                            <Route exact path='/goals'>
                                <Goals />
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
                
            </div>
        );
    }
}

const PrivateRoute = ({component: Component, ...rest}) => {
    return (

        // Show the component only when the user is logged in
        // Otherwise, redirect the user to /authenticate page
        <Route {...rest} render={props => (
            (sessionStorage.getItem(TOKEN) !== null) ?
                <Component {...props} />
            : <Redirect to="/authenticate" />
        )} />
    );
};

export default App;