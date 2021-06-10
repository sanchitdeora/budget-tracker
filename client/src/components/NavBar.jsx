import React from 'react';
import './NavBar.scss';

class NavBar extends React.Component {
	render() {
		return (
			<div className="navbar-container">
				<ul className="navbar-list">
					<li><a href='/'>Home</a></li>
					<li><a href='/ping'>Ping</a></li>
					{localStorage.getItem("isLoggedIn") 
					? 
					<li><a href='/logout'>Log out</a></li>
					:
					<li><a href='/authenticate'>Authenticate</a></li>}
				</ul>
			</div>
		);
	}
};

export default NavBar