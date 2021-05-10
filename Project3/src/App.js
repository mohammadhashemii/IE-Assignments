import React from "react";

import Keypad from "./components/Keypad";
import Screen from "./components/Screen";
import "./App.css";
import { createPortal } from "react-dom";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      screenText: 0,
      firstOperand: null,
      operator: "",
      isOperatorSelected: false,
      isOperationDone: false,
      memory: null
    };
  }
  handlePressDigit = (digit) => {
    if (this.state.screenText === 0 || this.state.isOperatorSelected || this.state.isOperatorDone) {
      this.setState({screenText: digit, isOperatorDone: false, isOperatorSelected: false});
    }
    else {
      this.setState({screenText: this.state.screenText + digit.toString()});
    }
  };
  handlePressOperator = (operator) => {
    if(this.state.firstOperand === null || this.state.isOperationDone) {
      this.setState({firstOperand: this.state.screenText, operator: operator, isOperatorSelected: true});
    } else {
      const res = eval(this.state.firstOperand.toString() + this.state.operator + this.state.screenText.toString());
      this.setState({screenText: res, firstOperand: res, isOperatorSelected: true, operator: operator});
    } 
  };
  handlePressAC = () => {
    this.setState({screenText: 0, firstOperand: null, isOperatorSelected: false, isOperationDone: false})
  };
  handlePressDot = () => {
    if (this.state.isOperationDone) {
      this.setState({screenText: 0 + "."});
    } else if (!this.state.screenText.toString().includes(".")) {
      this.setState({screenText: this.state.screenText + "."});
    }
    
  };
  handlePressNegator = () => {
    this.setState({screenText: (-1) * this.state.screenText});
  };
  handlePressResult = () => {
    if (this.state.firstOperand !== null) {
      const res = eval(this.state.firstOperand.toString() + this.state.operator + this.state.screenText.toString());
      this.setState({screenText: res, firstOperand: res, isOperationDone: true});
    }
  };
  handlePressMC = () => {
    // Clear the memory
    this.setState({memory: null});
  };
  handlePressMR = () => {
    // Set the display number equal to the number has been stored in the memory
    if (this.state.memory !== null) {
      this.setState({screenText: this.state.memory});
    }
    
  };
  handlePressM_Add = () => {
    // Sum the memory value with the number on the screen
    if (this.state.memory === null) {
      this.setState({memory: this.state.screenText});
    } else {
      this.setState({memory: this.state.memory + this.state.screenText});
    }
  };
  handlePressM_Subtract = () => {
    // Subtract the memory value from the number on the screen
    if (this.state.memory === null) {
      this.setState({memory: this.state.screenText});
    } else {
      this.setState({memory: this.state.memory - this.state.screenText});
    }
  };
  handlePressMS = () => {
    // Store the number on the display into the memory
    this.setState({memory: this.state.screenText});
    
  };

  render() {
    return (
      <div>
        <Screen text={this.state.screenText} />
        <Keypad
          onPressDigit={this.handlePressDigit}
          onPressOperator={this.handlePressOperator}
          onPressAC={this.handlePressAC}
          onPressDot={this.handlePressDot}
          onPressNegator={this.handlePressNegator}
          onPressResult={this.handlePressResult}
          onPressMC={this.handlePressMC}
          onPressMR={this.handlePressMR}
          onPressM_Add={this.handlePressM_Add}
          onPressM_Subtract={this.handlePressM_Subtract}
          onPressMS={this.handlePressMS}
        />
      </div>
    );
  }
}

export default App;
