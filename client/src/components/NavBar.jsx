import React from 'react';
import './NavBar.scss';
import { Link } from 'react-router-dom'
// import { Navbar, Nav, NavDropdown, LinkC } from 'react-bootstrap';
import AppBar from '@material-ui/core/AppBar';
import Grid from '@material-ui/core/Grid';
import Toolbar from '@material-ui/core/Toolbar';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import Menu from '@material-ui/core/Menu';
import IconButton from '@material-ui/core/IconButton';
import AccountCircleOutlinedIcon from '@material-ui/icons/AccountCircleOutlined';

class NavBar extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			anchorEl: null
		}
	}

	handleClick = (event) => {
    	this.setState({anchorEl: event.currentTarget});
  	};

  	handleClose = () => {
		this.setState({anchorEl: null});
  	};

	render() {
		return (
			<div className="navbar-container">
				<AppBar position="static">
					<Toolbar>
						<Grid container className="grid-container" spacing={0}>
							<Grid item xs={12}>
								<Button color="inherit" className="navbar-item" component={Link} to={'/home'}>Home</Button>
								<Button color="inherit" className="navbar-item" component={Link} to={'/ping'}>Ping</Button>
								{
									this.props.isLoggedIn ?
									<div className="account-container">
										<IconButton 
											
											className="account-button"
											aria-controls="simple-menu" 
											aria-haspopup="true"
											onClick={this.handleClick.bind(this)}
										>
											<AccountCircleOutlinedIcon style={{ color: 'white', }} className="accountControl"/>
											<div className="icon-label" style={{ color: 'white' }}>{sessionStorage.getItem("username")}</div>
										</IconButton>
										<Menu
											id="simple-menu"
											anchorEl={this.state.anchorEl}
											getContentAnchorEl={null}
											anchorOrigin={{vertical: 'bottom', horizontal: 'left'}}	
											keepMounted
											open={Boolean(this.state.anchorEl)}
											onClose={this.handleClose.bind(this)}
										>
											<MenuItem 
												className="account-menu" 
												component={Link} to={'/myaccount'}
												onClick={this.handleClose.bind(this)}>
														My Account
											</MenuItem>
											<MenuItem 
												className="account-menu" 
												component={Link} to={'/logout'}
												onClick={this.handleClose.bind(this)}>
													Log out
											</MenuItem>
										</Menu>
									</div>
									:
									<div className="account-container">
										<Button color="inherit" className="account-menu" component={Link} to={'/auth'}>Log in</Button>
									</div>
								}
							</Grid>
						</Grid>
					</Toolbar>
				</AppBar>
			</div>
		);
	}
};


export default NavBar