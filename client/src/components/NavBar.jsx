import React from 'react';
import './NavBar.scss';
import PropTypes from 'prop-types';
import { Router, Link } from 'react-router-dom'

class NavBar extends React.Component {
	constructor(props) {
		super(props);
		this.state = {}
	}

	render() {
		return (
			<div className="navbar-container">
					<ul className="navbar-list">
						<li><Link to='/home'>Home</Link></li>
						<li><Link to='/ping'>Ping</Link></li>
						
						<li>
							{this.props.isLoggedIn ? 
							<Link to='/logout' className="accountControl">Log out</Link> :
							<Link to='/authenticate' className="accountControl">Log in</Link>}
						</li>
					</ul>
			</div>
		);
	}
};


export default NavBar