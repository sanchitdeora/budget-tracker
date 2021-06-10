import React from 'react';
import './NavBar.scss';

class NavBar extends React.Component {
	constructor(props) {
		super(props);
		this.state = {}
	}



	render() {
		return (
			<div className="navbar-container">
				<ul className="navbar-list">
					<li><a href='/'>Home</a></li>
					<li><a href='/ping'>Ping</a></li>
					
					<li>
						{this.props.isLoggedIn ? 
						<a href='/logout' className="accountControl">Log out</a> :
						<a href='/authenticate' className="accountControl">Log in</a>}
					</li>
				</ul>
			</div>
		);
	}
};

export default NavBar