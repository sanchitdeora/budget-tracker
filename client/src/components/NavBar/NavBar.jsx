import React from 'react';
import './NavBar.scss';
import { Link } from 'react-router-dom'

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
						<li><Link to='/transactions'>Transactions</Link></li>
						<li><Link to='/bills'>Bills</Link></li>
						<li><Link to='/budgets'>Budgets</Link></li>
						
						{
							this.props.isLoggedIn ? 
							<div>
								{/* <li><Link to='/logout' className="accountControl">Account</Link></li> */}
								<li><Link to='/logout' className="accountControl">Log out</Link></li>
							</div> :
							<li><Link to='/authenticate' className="accountControl">Log in</Link></li>
						}
					</ul>
			</div>
		);
	}
};


export default NavBar