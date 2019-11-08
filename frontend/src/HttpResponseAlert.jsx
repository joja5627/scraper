import React from 'react';
import { Button, Alert } from "react-bootstrap";
const fontSize10 = {
  fontSize:'10px'
};
export default class HttpResponseAlert extends React.Component {
  constructor(props) {
    super(props);
  }
  // body: (...)
  // bodyUsed: true
  // headers: Headers {}
  // ok: true
  // redirected: false
  // status: 200
  // statusText: "OK"
  // type: "cors"
  // url: "http://localhost:1000/register"

  renderBody = body =>{
    return body ? <i>{JSON.stringify(body)}</i>  : <i>empty</i>  
  }
  render() {
    const {response} = this.props
    
    return (
      <Alert
        className="padding0"
        style={fontSize10} 
        variant={this.props.varient} 
        onClose={this.props.onDelete} dismissible>
      <b>HTTP Status {response.status}</b>
      <br></br>
      <small>HTTP Status Text {response.statusText}</small>
      <br></br>
      {this.renderBody(this.props.response.body)}
      <br></br>
    </Alert>
    );
  }
}