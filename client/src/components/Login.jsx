import React from 'react';
import axios from 'axios';


class Login extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			email: "",
			password: "",
		};
	};

	handleChange(event) {
		let value = event.target.value;
		let name = event.target.name;

		this.setState({
			[name]: value
		});
	}

	submitLogin(event) {
		console.log('The form was submitted with the following data:');
		console.log("state:", this.state)

		axios.post("/api/login", this.state)
			.then(res => {
				console.log(res);
				console.log(res.data);
			})
			.catch(error => {
				console.log(error.response)
			});

	}

	render() {
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

export default Login;