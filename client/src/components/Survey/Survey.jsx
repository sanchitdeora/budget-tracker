import React from 'react';
import "./Survey.scss";
import axios from 'axios';
import { Redirect } from "react-router-dom";
import { EMAIL, IS_SURVEY_COMPLETE } from '../../utils/GlobalConstants'


class Survey extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			email: sessionStorage.getItem(EMAIL),
			monthlyIncome: 0,
			savingsType: "medium",
			monthlyLimit: 0,
			isSurveyComplete: JSON.parse(sessionStorage.getItem(IS_SURVEY_COMPLETE))
		};
	};

	handleChange = (event) => {
		let value = event.target.value;
		let name = event.target.name;
		this.setState({
			[name]: value
		});
	}

	handleSubmit = () => {
		console.log('The form was submitted with the following data:');
		this.postSurveyRequest()
	}

	async postSurveyRequest() {
		let res = await axios.post("/api/quickstart", this.state);
		console.log(res);
		if (res.status === 200) {
			sessionStorage.setItem(IS_SURVEY_COMPLETE, true)
			this.setState({
				isSurveyComplete: true
			})
		}
	}


	render() {
		if (this.state.isSurveyComplete) {
			return(
				<Redirect to='/home' />
			)
		} else {
			return (
				<div className="inner-container">
					<div className="header">
						Before we Start...
					</div>
					<div className="box">
						<div className="input-group">
							<label htmlFor="monthlyIncome" className="survey-label">Monthly Income</label>
							<input
								type="number"
								name="monthlyIncome"
								className="input-box"
								value={this.state.monthlyIncome}
								onChange={this.handleChange}
							/>
						</div>

						<div className="input-group">
							<label htmlFor="savingsType" className="survey-label">Savings Plan</label>
							<select className="input-box" defaultValue={this.state.savingsType}>
								<option value="light">Light</option>
								<option value="medium">Medium</option>
								<option value="heavy">Heavy</option>
							</select>
						</div>

						<div className="input-group">
							<label htmlFor="monthlyLimit" className="survey-label">Monthly Expense Limit</label>
							<input
								type="number"
								name="monthlyLimit"
								className="input-box"
								value={this.state.monthlyLimit}
								onChange={this.handleChange}
							/>
						</div>
						<button
							type="submit"
							className="submit-btn"
							onClick={this.handleSubmit}
							>
								Finish
						</button>
					</div>
				</div>
			)
		}
	}
}

export default Survey;