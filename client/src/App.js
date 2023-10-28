import './App.scss';
import React from 'react';
import { TOKEN, ACTIVE_PAGE } from './utils/GlobalConstants'
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Home from "./components/Home/Home.jsx"
import Ping from "./components/Ping/Ping"
import NavBar from './components/NavBar/NavBar.jsx';
import Logout from './components/Authentication/Logout.jsx';
import Authenticate from './components/Authentication/Authenticate.jsx';
import Survey from './components/Survey/Survey.jsx';
import Transactions from './components/Transactions/Transactions';
import Bills from './components/Bills/Bills';
import BudgetCards from './components/Budgets/BudgetCards';
import GoalCards from './components/Goals/GoalCards';
import { menuItems } from './utils/menuItems';


// const TOKEN = "token"

class App extends React.Component {
    constructor(props) {
        super(props);
        const token = sessionStorage.getItem(TOKEN)
        this.state = {
            isLoggedIn: token !== null,
            token: sessionStorage.getItem(TOKEN),
            navBarActive: null,
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

    setNavBarActive = (item) => {
        // console.log("setting menu item to active: ", item)
        this.setState({ navBarActive: item })
    }

    getNavBarActive = () => {
        // console.log("getting menu item to active: ", this.state.navBarActive)
       return this.state.navBarActive;
    }

    render() {
        return (
            <div className='app-container'>
                <Router>
                    <NavBar isLoggedIn={this.state.isLoggedIn} getActive={this.getNavBarActive} setActive={this.setNavBarActive} />
                        <div className="inner-container">
                        <Routes>
                            {/* <PrivateRoute component={Home} path="/home" exact /> */}

                            <Route path="/ping" element={<Ping />} />
                            <Route exact path='/authenticate' element={<Authenticate setLoginState={this.setLoginState} setToken={this.setToken} />} />
                            <Route exact path='/logout' element={<Logout setLoginState={this.setLoginState} setToken={this.setToken} />} />

                            <Route exact path='/transactions' element={<Transactions setNavBarActive={this.setNavBarActive} />} />
                            <Route exact path='/bills' element={<Bills setNavBarActive={this.setNavBarActive} />} />
                            <Route exact path='/budgets' element={<BudgetCards setNavBarActive={this.setNavBarActive} />} />
                            <Route exact path='/goals' element={<GoalCards setNavBarActive={this.setNavBarActive} />} />

                            <Route exact path="/home" element={<Home setNavBarActive={this.setNavBarActive} />} />
                            <Route exact path="/" element={<Home setNavBarActive={this.setNavBarActive} />} />
                                {/* {
                                this.state.isLoggedIn ? 
                                <Navigate to="/home" /> : 
                                <Navigate to="/authenticate" />
                                }  
                            </Route> */}

                            {/* <Route exact path='/survey' element={<Survey />}> */}
                                {/* <Survey /> */}
                            {/* </Route> */}
                        </Routes>
                    </div>
                </Router>
            </div>
        );
    }
}

// const PrivateRoute = ({component: Component, ...rest}) => {
//     return (

//         // Show the component only when the user is logged in
//         // Otherwise, navigate the user to /authenticate page
//         <Route {...rest} render={props => (
//             (sessionStorage.getItem(TOKEN) !== null) ?
//                 <Component {...props} />
//             : <Navigate to="/authenticate" />
//         )} />
//     );
// };

//https://huemint.com/website-2/#palette=050505-2dfb86-1f80ff-fe1616

export default App;