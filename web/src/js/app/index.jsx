import $ from 'jquery';
import React from 'react';

export class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            time:   new Date(),
            builds: [],
        };
    }

    componentDidMount() {
        this.timer = setInterval(this.tick.bind(this), 1e3);

        $.ajax({
            url: this.props.url,
            dataType: 'json',
            cache: false,
        })
        .done((data) => {
            this.setState({ builds: data });
        }.bind(this))
        .fail((xhr, status, err) => {
            console.error(this.props.url, status, err.toString());
        }.bind(this));
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }

    render() {
        var builds = this.state.builds.map((build) => {
            return (
                <li>{build.commitId}</li>
            );
        });

        return (
            <div>
                <div>Delta Test {this.state.time.toString()} aaaaa</div>
                <ul>{builds}</ul>
            </div>
        );
    }

    tick() {
        console.log('tick');
        this.setState({
            time: new Date(),
        });
    }
}

React.render(<App url="/api/builds" />, document.querySelector('#root'));
