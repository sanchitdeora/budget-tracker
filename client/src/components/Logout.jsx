import React from 'react';
import { Redirect } from "react-router-dom";


class Logout extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};
		localStorage.clear()
		console.log(localStorage)
	};

	render() {
		return(
			<Redirect to='/' />
		)
	}
}

export default Logout;