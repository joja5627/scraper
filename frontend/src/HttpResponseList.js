import React, { Component } from 'react';
import HttpResponseAlert from './HttpResponseAlert';

import { CSSTransition, TransitionGroup } from 'react-transition-group';

import './HttpResponseList.css';

class HttpResponseList extends Component {
  constructor(props) {
    super(props);
    this.state = {
      responses: []
    };
}

  
  componentDidUpdate(prevProps, prevState) {
    if(prevProps.response !== this.props.response ){
      let response = this.props.response
      if(response){
        this.setState((prevState) => {
          return { 
            responses: prevState.responses.concat(response) 
          };
        });
      }
     
    }
   
  }
  inRange  = (x, min, max) => {
    return ((x-min)*(x-max) <= 0);
  }

  renderHttpResponseVarient = status => {
    if (this.inRange(status,200,200)) {
      return  "success"
    } else if (this.inRange(status,500,599)) {
      return "danger"
    } else if(this.inRange(status,400,499)){
      return "warning"
    }
  }

  onDeleteHandler = itemKey => {
    let newStateItems = { ...this.state.responses };

    delete newStateItems[itemKey];

    this.setState({
      responses: {
        ...newStateItems
      }
    });
  };

  render() {
    const itemsList = Object.keys(this.state.responses).map(itemKey => {
      let status = this.state.responses[itemKey].status
      let varient = this.renderHttpResponseVarient(status)
      let response = this.state.responses[itemKey]
      return (
        <CSSTransition key={itemKey} timeout={500} classNames="move">
          <HttpResponseAlert
            varient={varient}
            response={response}
            onDelete={() => {
              this.onDeleteHandler(itemKey);
            }} />
        </CSSTransition>
      );
    });

    return (
      <div className="responses-section">
        {/* <Toolbar onAddHandler={this.addItemHandler} /> */}
        <TransitionGroup className="responses-section__list">
          {itemsList}
        </TransitionGroup>
      </div>
    );
  }
}

export default HttpResponseList;
