import React from 'react';
import Login from './Login.jsx';
import Register from './Register';
import './Authenticate.css';

class Authenticate extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			isLoginOpen: true,
			isRegisterOpen: false
		};
	};

	showRegister() {
		this.setState({ isRegisterOpen: true, isLoginOpen: false });
	}

	showLogin() {
		this.setState({ isRegisterOpen: false, isLoginOpen: true });
	}

	render() {
		return (
			<div className="box-container">
				<div className="box-controller">
					<div
						className={"controller " + (this.state.isLoginOpen
							? "selected-controller"
							: "")}
						onClick={this.showLogin.bind(this)}>
						Login
					</div>
					<div
						className={"controller " + (this.state.isRegisterOpen
							? "selected-controller"
							: "")}
						onClick={this.showRegister.bind(this)}>
						Register
       				</div>
				</div>
				<div className="box-container">
					{this.state.isLoginOpen && <Login />}
					{this.state.isRegisterOpen && <Register />}
				</div>
			</div>
		);
	}
}

export default Authenticate;