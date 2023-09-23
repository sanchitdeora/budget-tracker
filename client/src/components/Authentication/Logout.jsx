import React from 'react';
import { Navigate } from "react-router-dom";


class Logout extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
        };
        this.props.setLoginState(false)
        this.props.setToken(null)
        sessionStorage.clear()
        console.log("log props", this.props)
    };

    render() {
        return(
            <Navigate to='/' />
        )
    }
}

export default Logout;