import React from 'react';
import axios from 'axios';


class Register extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            firstname: "",
            lastname: "",
            email: "",
            password: "",
            dateOfBirth: "",
            phoneNumber: "",
        };
    };

    handleChange(event) {
        let value = event.target.value;
        let name = event.target.name;

        this.setState({
            [name]: value
        });
    }

    submitRegister(event) {
        console.log('The register form was submitted with the following data:');
        
        axios.post("/api/register", this.state)
            .then(res => {
                // console.log(res);
                // console.log(res.data);
            })
            .catch(error => {
                // console.log(error.response)
            });
        this.props.showLogin()
    }

    render() {
        return (
            <div className="inner-container">
                <div className="header">
                    Register
                </div>

                <div className="box">

                    <div className="input-group">
                        <label htmlFor="firstname">First Name</label>
                        <input
                            type="text"
                            name="firstname"
                            className="input-box"
                            placeholder="First name"
                            onChange={this.handleChange.bind(this)}
                        />
                    </div>

                    <div className="input-group">
                        <label htmlFor="lastname">Last Name</label>
                        <input
                            type="text"
                            name="lastname"
                            className="input-box"
                            placeholder="Last Name"
                            onChange={this.handleChange.bind(this)}
                        />
                    </div>

                    <div className="input-group">
                        <label htmlFor="email">Email</label>
                        <input type="email" name="email" className="input-box" placeholder="Email"
                            onChange={this.handleChange.bind(this)} />
                    </div>

                    <div className="input-group">
                        <label htmlFor="password">Password</label>
                        <input
                            type="password"
                            name="password"
                            className="input-box"
                            placeholder="Password"
                            onChange={this.handleChange.bind(this)}
                        />
                    </div>

                    <div className="input-group">
                        <label htmlFor="dateOfBirth">Date of Birth</label>
                        <input
                            type="date"
                            name="dateOfBirth"
                            className="input-box"
                            placeholder="Date of Birth"
                            onChange={this.handleChange.bind(this)}
                        />
                    </div>

                    <div className="input-group">
                        <label htmlFor="phoneNumber">Phone Number</label>
                        <input
                            type="text"
                            name="phoneNumber"
                            className="input-box"
                            placeholder="Phone Number"
                            onChange={this.handleChange.bind(this)}
                        />
                    </div>

                    <button
                        type="button"
                        className="submit-btn"
                        onClick={this.submitRegister.bind(this)}
                    >
                        Register
                    </button>
                </div>
            </div>
        );
    }
}

export default Register;