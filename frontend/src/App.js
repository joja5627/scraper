import React, { Component } from 'react';
import './App.css';
import {
  Button,
  Container,
  Grid,
  Card,
  Progress,
  Form,
  Icon,
  Input,
  Select,
  TextArea
} from 'semantic-ui-react';
import HttpResponseList from './HttpResponseList'

import _ from "lodash"

const ProgressBar = ({ uploadState, percentUploaded }) =>
  uploadState === 'uploading' && (
    <Progress
      percent={percentUploaded}
      progress
      indicating
      size="medium"
      inverted
    />
  );

const buttonStyle = {
  margin: '10px'
};
const headerMarginStyle = {
  marginBottom: '100px'
};
const progressBarStyle = {
  marginLeft: '100px',
  marginRight: '100px'
};

class TodoList extends React.Component {
  render() {
    var items = this.props.items.map((listing, index) => {
      return (
        <TodoListItem
          className="margin10 width50"
          onListingClick={this.props.onListingClick}
          key={index}
          listing={listing}
          index={index}
          removeItem={this.props.removeItem}
        />
      );
    });
    return <Container clasName="padding40" textAlign="left">{items}</Container>;
  }
}

class TodoListItem extends React.Component {
  constructor(props) {
    super(props);
    this.onClickClose = this.onClickClose.bind(this);
    this.onClickEmail = this.onClickEmail.bind(this);

  }
 
  onClickEmail() {
    const { index,listing } = this.props
    var parsedIndex = parseInt(index);
    this.props.removeItem(parsedIndex);
    this.props.onListingClick(listing)
  }
  onClickClose() {
    var index = parseInt(this.props.index);
    this.props.removeItem(index);
  }
  //      <a href={`${this.props.item}`}> </a>
  //      <div onClick={this.props.onClick}>{this.props.item}</div>

  render() {
    const { listing } = this.props
    return (
      <p className="listing">

        <button onClick={this.onClickClose} class="ui icon button">
          <i class="times icon"></i>
        </button>
        <button onClick={this.onClickEmail} class="ui icon button">
          <i class="envelope icon"></i>
        </button>
        <span className="text-overflow"> {listing.title}</span>
      </p>


    );
  }
}

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      listings: [],
      eventTarget: null,
      webSocket: null,
      linkViewLimit: 300,
      linkCount: 0,
      emailAll: false,
      selectedListing: null,
      sendEmailResponse: null
    };

    this.removeItem = this.removeItem.bind(this);
    this.clListingScrapePercentage = 0;
    this.contactInfoScrapePercentage = 0;
  }
  removeItem(itemIndex) {
    let newLinks = this.state.links;
    newLinks.splice(itemIndex, 1);
    this.setState({ links: newLinks });
  }
  onEvent = event => {
    if (event.messageType === 'listingPercentComplete') {
      this.setState({ clListingScrapePercentage: Math.round(event.payload) });
    }
    else if (event.messageType === 'listing') {
      let newListing = JSON.parse(event.payload)
      if (!_.find(this.state.listings, { url: newListing.url })) {
        if(this.state.emailAll){
          this.sendEmail(newListing)
        }else{
          let newListings = this.state.listings;
          newListings.push(newListing);
          this.setState({ links: newListings, linkCount: newListings.length });

        }
      }

    }
  };

  closeSocket = () => {
    if (this.state.socket) {
      this.state.socket.send("close server");
      this.state.socket.close();
      this.setState({
        socket: null,
        clListingScrapePercentage: 0
      });
    }
  };
  restartSocket = () => {
    this.closeSocket();
    this.getLinks();
  };
  getLinks = () => {
    let socket = new WebSocket('ws://localhost:8080/scrape');

    socket.onopen = () => {
      this.setState({
        socket: socket
      });
      console.log('open');
    };
    socket.onmessage = e => {
      this.onEvent(JSON.parse(e.data));
    };
    socket.onclose = () => {
      console.log('closing');

      socket.close();
      console.log('closed');
      this.setState({
        socket: null
      });
    };
  };
  sendEmail = (listing) => {
    
    fetch("http://localhost:8080/sendEmail", {
      method: 'POST', // or 'PUT'
      body: JSON.stringify(listing), // data can be `string` or {object}!
      headers: {
        'Content-Type': 'application/json',
        "Accept": "application/json"
      }
    }).then(res => {
      this.setState({ sendEmailResponse: res })
    })

  }

  onInputChange = event => {
    this.setState({ eventTarget: event.target.value.toLowerCase() });
  };

  filterFunction = link => {
    return link.toLowerCase().search(this.state.eventTarget) !== -1;
  };

  filterList = listing => {
    if (this.state.eventTarget) {
      return listing.filter(link => this.filterFunction(link));
    } else {
      return listing;
    }
  };


  render() {
    const { sendEmailResponse } = this.state
    return (
      <div className="App">
        <Container textAlign="center">
          <h1 className="ui header fontWeight100">Craigslist Web Scraper</h1>
          {this.state.clListingScrapePercentage && (
            <ProgressBar
              uploadState={'uploading'}
              percentUploaded={this.state.clListingScrapePercentage}
            />
          )}

          <Grid columns={3} divided>
            <Grid.Row>
              <Grid.Column width={4}>
                <text>{`link count: ${this.state.linkCount}`}</text>
                <Card>
                  <Form>
                    <Card.Content>
                      <Card.Meta>
                        <Button.Group>
                          <Button onClick={this.getLinks} icon>
                            <Icon name="play" />
                          </Button>
                          <Button onClick={this.closeSocket} icon>
                            <Icon name="stop" />
                          </Button>
                          <Button onClick={this.restartSocket} icon>
                            <Icon name="redo" />
                          </Button>
                          <Button active={this.state.emailAll} onClick={() => this.setState({ emailAll: !this.state.emailAll })} icon>
                            <Icon name="envelope" />
                          </Button>
                        </Button.Group>
                      </Card.Meta>
                      <Card.Description  >
                        <input
                          className="m-t-10"
                          type="text"
                          onChange={this.onInputChange}
                          name="filter-field"
                          placeholder="filter results"
                        />
                        {/* // <p className="margin0 padding0" align="left">
      //   <Icon name="close icon" onClick={this.onClickClose}></Icon>
      //   <a href={`${this.props.item}`}> {this.props.item}</a>
      // </p> */}
                        {/* <div className="ui action input">
                          <input type="text" placeholder="Search...">
                            <button className="ui button">Search</button>
                          </input>
                        </div> */}
                      </Card.Description>
                    </Card.Content>

                  </Form>
                </Card>
              </Grid.Column>

              <Grid.Column width={6}>
                {this.state.listings && (
                  <TodoList
                    onListingClick={(listing) => this.sendEmail(listing)}
                    items={this.filterList(this.state.listings)}
                    removeItem={this.removeItem}
                  />
                )}
              </Grid.Column>
              <Grid.Column width={6}>
                <HttpResponseList response={sendEmailResponse} />
              </Grid.Column>
            </Grid.Row>
          </Grid>
        </Container>
      </div>
    );
  }
}

export default App;
