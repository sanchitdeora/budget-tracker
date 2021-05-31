import React from 'react';

class Home extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
		};
	};
	render() {
		return (
			<div>
				<div>
					<h3>Served by Golang server from a single binary file</h3>
				</div>
			</div>
		);
	}
}

export default Home;