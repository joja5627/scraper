import React, { Component } from 'react';
import HttpResponseAlert from './HttpResponseAlert';
import _ from 'lodash';

import { CSSTransition, TransitionGroup } from 'react-transition-group';

import './HttpResponseList.css';

class HttpResponseList extends Component {
    constructor(props) {
        super(props);
      }
shouldComponentUpdate(prevProps, prevState){
    return !_.isEqual(prevProps != this.props) || !_.isEqual(prevState != this.state)
}
  

  render() {
      const {responses} = this.props
      const itemsList = responses.map((response,index) => {
      return (
        <CSSTransition key={`${index}-response-key`} timeout={500} className="move">
          <HttpResponseAlert
            response={response}
            onClickRemove={this.props.onClickRemove} />
        </CSSTransition>
      );
    });

    return (
      <div className="responses-section">
        <TransitionGroup className="responses-section__list">
          {itemsList}
        </TransitionGroup>
      </div>
    );
  }
}

export default HttpResponseList;
