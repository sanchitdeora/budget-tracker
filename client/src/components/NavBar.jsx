import React from 'react';
import './NavBar.css';

class NavBar extends React.Component {
	render() {
		return (
			<div className="navbar-container">
				<ul>
					<li><a href='/'>Home</a></li>
					<li><a href='/ping'>Ping</a></li>
					<li><a href='/authenticate'>Authenticate</a></li>
				</ul>
			</div>
		);
	}
};

export default NavBar