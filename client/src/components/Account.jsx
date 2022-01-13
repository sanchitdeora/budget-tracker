import React from 'react';
import axios from 'axios';

class Account extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};

		this.getUserDetails()
	};

	async getUserDetails() {
		let url = "/api/user/"+this.props.userID;
		console.log(url)
		let res = await axios.get(url);
		console.log(res);
		if (res.status === 200) {
			this.setState({
				isLoggedIn: true,
				token: res.data.token,
			});
			this.props.setLoginState(true)
			this.props.setToken(res.data.token)
			this.props.setUser(res.data)
		}
	}

	render() {
		return (
			<div>
				<div>
					Hello My Account!
				</div>
			</div>
		);
	}
}

export default Account;