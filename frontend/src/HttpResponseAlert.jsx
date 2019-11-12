import React from 'react';
import { Button, Alert } from "react-bootstrap";
const fontSize10 = {
  fontSize:'10px'
};
export default class HttpResponseAlert extends React.Component {
  constructor(props) {
    super(props);
  }

  setInterval = () => {
    this.interval = setInterval(() => 
            this.props.onClickRemove(this.props.response), 10000);
  }
  inRange  = (x, min, max) => {
    return ((x-min)*(x-max) <= 0);
  }
  renderHttpResponseVarient = status => {
    if (this.inRange(status,200,299)) {
      return  "success"
    } else if (this.inRange(status,500,599)) {
      return "danger"
    } else if(this.inRange(status,400,499)){
      return "warning"
    }
  }
  render() {
    const {response,onClickRemove} = this.props
    let variant = this.renderHttpResponseVarient(response.status)
    this.setInterval()
    return (
      <Alert
        className="padding0 text-left"
        style={fontSize10} 
        variant={variant} 
        onClick={() => 
            onClickRemove(response)} dismissible>
      <b>HTTP Status {response.status}</b>
      <br></br>
      <small>HTTP Status Text {response.statusText}</small>
      <br></br>
      
       {Object.keys(response).map((key,index) => {
           return <span key={`${index}-http-res`}><i>{key} </i><br></br><i>{response[key]} </i></span>
       })}
    </Alert>
    );
  }
}