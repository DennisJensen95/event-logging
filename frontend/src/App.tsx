import React from 'react';
import './App.css';
import kafkaSetState from "./consumer"

type EmptyProps = {}

type AppState = {
  cpu_cores: number,
  cpu_system_usage: number,
  cpu_usage: number,
  cpu_user_usage: number
  received_first_message: boolean,
}

class App extends React.Component<EmptyProps, AppState> {
  state: AppState = {
    cpu_cores: 0,
    cpu_system_usage: 0,
    cpu_usage: 0,
    cpu_user_usage: 0,
    received_first_message: false,
  };

  componentDidMount() {
    // kafkaSetState(this.setState);
  }

  render() {
    return (
      <div className="App" >
        <header className="App-header">
          <p>
            Kafka drive website. {this.state.received_first_message ? "Received first message" : "Waiting for first message"}
          </p>
        </header>
      </div>
    );
  }
}

export default App;
