import React from 'react';
import { Redirect } from "react-router-dom";


class Logout extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};
		localStorage.setItem("isLoggedIn", false)
		console.log(localStorage.getItem("isLoggedIn"))
	};

	render() {
		if (localStorage.getItem("isLoggedIn")){
			return(
				<Redirect to='/' />
			)
		} else {
		return (
				<div className="inner-container">
					ERROR
				</div>
			);
		}
	}
}

export default Logout;