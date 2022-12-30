import React from 'react';
import axios from 'axios';
import { Redirect } from "react-router-dom";
import { TOKEN, EMAIL, IS_SURVEY_COMPLETE } from '../../utils/GlobalConstants'

class Login extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            email: "",
            password: "",
            isLoggedIn: false,
            token: "",
        };
    };

    handleChange(event) {
        let value = event.target.value;
        let name = event.target.name;
        this.setState({
            [name]: value
        });
    }

    async postLoginRequest() {
        let res = await axios.post("/api/login", this.state);
        console.log(res);
        if (res.status === 200) {
            sessionStorage.setItem(IS_SURVEY_COMPLETE, res.data.isSurveyComplete)
            sessionStorage.setItem(EMAIL, this.state.email)
            this.setState({
                isLoggedIn: true,
                token: res.data.token
            });
            this.props.setLoginState(true)
            this.props.setToken(res.data.token)
        }
    }

    submitLogin() {
        console.log('The form was submitted with the following data:');
        this.postLoginRequest()
    }

    handleShowRegister = () => {
        this.props.showRegister()
    }

    render() {
        if (sessionStorage.getItem(TOKEN) != null) {
            if (sessionStorage.getItem(IS_SURVEY_COMPLETE) === "true") {
                return(
                    <Redirect to='/home' />
                )
            } else {
                return(
                    <Redirect to='/survey' />
                )
            }
            
        } else {
            return (
                <div className="inner-container">
                    <div className="header">
                        Login
                    </div>
                    <div className="box">
                        <div className="input-group">
                            <label htmlFor="email" className="login-label">Email Address</label>
                            <input
                                type="email"
                                name="email"
                                className="input-box"
                                placeholder="Email Address"
                                onChange={this.handleChange.bind(this)}
                            />
                        </div>

                        <div className="input-group">
                            <label htmlFor="password">Password</label>
                            <input
                                type="password"
                                name="password"
                                className="input-box"
                                placeholder="Password"
                                onChange={this.handleChange.bind(this)}
                            />
                        </div>

                        <button
                            type="submit"
                            className="submit-btn"
                            onClick={this.submitLogin.bind(this)}>Login
                        </button>

                        <button id="showRegister" onClick={this.handleShowRegister}> Don't have an account yet? Join now!
                        </button>
                    </div>
                </div>
            );
        }
    }
}

export default Login;