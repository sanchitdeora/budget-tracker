import React from 'react';
import axios from 'axios';
import { Redirect } from "react-router-dom";


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
			this.setState({
				isLoggedIn: true,
				token: res.data.token
			});
		}
		sessionStorage.setItem("token", res.data.token)
		console.log(this.state);
	}

	submitLogin(event) {
		console.log('The form was submitted with the following data:');
		this.postLoginRequest()

		this.props.setLoginState(true)
		localStorage.setItem("isLoggedIn", true)

		console.log("checking....", localStorage.getItem("token"), localStorage.getItem("isLoggedIn"))
		if (this.state.isLoggedIn) {
			console.log("state:", this.state)
		}
	}

	render() {
		if (localStorage.getItem("isLoggedIn")){
			return(
				<Redirect to='/home' />
			)
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
							type="button"
							className="submit-btn"
							onClick={this.submitLogin.bind(this)}>Login
						</button>
					</div>
				</div>
			);
		}
	}
}

export default Login;