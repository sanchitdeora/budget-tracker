import React from 'react';
import './NavBar.scss';
import avatar from '../../img/avatar.png'
import { signout } from '../../utils/Icons'
import { Link } from 'react-router-dom'
import { menuItems } from '../../utils/menuItems'

class NavBar extends React.Component {
    constructor(props) {
        super(props);
        console.log("getting navbar props", this.props)
        this.state = {
            active: this.props.getActive(),
        }
        console.log("getting navbar state", this.state)

    }

    setActive = (item) => {
        console.log("setting menu item to active: ", item)
        this.setState({ active: item })
    }

    render() {
        return (
            <div className='navbar'>
                <div className="user-con">
                    <img src={avatar} alt="" />
                    <div className="text">
                        <h2>Sanchit</h2>
                        <p>$0</p>
                    </div>
                </div>
                <ul className="menu-items">
                    {menuItems.map((item) => {
                        return <li
                            key={item.id}
                            onClick={this.props.setActive.bind(this, item)}
                            className={this.props.getActive() === item ? 'active': ''}
                        >
                            <Link to={item.link}>
                                {item.icon}
                                <span>{item.title}</span>
                            </Link>
                        </li>
                    })}
                </ul>
                <div className="bottom-nav">
                    <li>
                        {signout} Sign Out
                    </li>
                </div>
            </div>
        );
    }
};

export default NavBar