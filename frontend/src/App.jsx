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
import HttpResponseList from './HttpResponseList';
import { CSSTransition, TransitionGroup } from 'react-transition-group';
import _ from 'lodash';
import Queue from './Queue';

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

class Listings extends React.Component {
  shouldComponentUpdate(prevProps, prevState) {
    return (
      !_.isEqual(prevProps != this.props) || !_.isEqual(prevState != this.state)
    );
  }

  render() {
    var listings = this.props.listings.map((listing, index) => {
      return (
        <CSSTransition
          key={`${index}-listinglkey`}
          timeout={500}
          className="move"
        >
          <Listing
            className="margin10 width50"
            onClickRemove={this.props.onClickRemove}
            onClickEmail={this.props.onClickEmail}
            listing={listing}
          />
        </CSSTransition>
      );
    });
    return (
      <TransitionGroup className="listing-section__list">
        <Container className="padding40 listing-section" textAlign="left">
          {listings}
        </Container>
      </TransitionGroup>
    );
  }
}

class Listing extends React.Component {
  render() {
    let { listing, onClickRemove, onClickEmail } = this.props;
    return (
      <p className="listing">
        <button
          onClick={() => onClickRemove(listing)}
          className="ui icon button"
        >
          <i className="times icon"></i>
        </button>
        <button
          onClick={() => onClickEmail(listing)}
          className="ui icon button"
        >
          <i className="envelope icon"></i>
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
      responses: [],
      eventTarget: null,
      webSocket: null,
      linkViewLimit: 300,
      linkCount: 0,
      emailAll: false,
      selectedListing: null,
      sendEmailStatus: null,
      emailResponses: [],
      sendingEmail: false,
      emailQueue: new Queue()
    };

    this.clListingScrapePercentage = 0;
    this.contactInfoScrapePercentage = 0;
  }
  shouldComponentUpdate(prevProps, prevState) {
    return (
      !_.isEqual(prevProps != this.props) || !_.isEqual(prevState != this.state)
    );
  }
componentDidUpdate() {
    if (this.state.emailQueue.size() > 0 && !this.state.sendingEmail) {
      const { emailQueue } = this.state;
      let listing = emailQueue.pop();

      this.setState({ emailQueue: emailQueue, sendingEmail: true });
      console.log(listing)
      fetch('http://localhost:8080/sendEmail', {
        method: 'POST', // or 'PUT'
        body: JSON.stringify(listing), // data can be `string` or {object}!
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json'
        }
      }).then((res) => res.json())
      .then((json) => {
          console.log(json)
        let {emailResponses} = this.state
        emailResponses.concat({body:json.body,status:json.status,statusText:json.statusText})
        this.setState({
          sendingEmail: false,
          emailResponses:emailResponses.concat({body:json.body,status:json.status,statusText:json.statusText})

        });
      });
    }
  }

  onEvent = event => {

    if (event.messageType === 'state') {
       
    }
    else if(event.messageType === 'percentComplete'){
    //   console.log(event.payload)
        this.setState({ clListingScrapePercentage: Math.round(event.payload) });
    }
    else if (event.messageType === 'listing') {
       
      let newListing = JSON.parse(event.payload);
      if (!_.find(this.state.listings, { url: newListing.url })) {
        if (this.state.emailAll) {
          this.sendEmail(newListing);
        } else {
          let newListings = this.state.listings;
          newListings.push(newListing);
          this.setState({
            listings: newListings,
            linkCount: newListings.length
          });
        }
      }
    }
  };

  closeSocket = () => {
    if (this.state.socket) {
      this.state.socket.send('close server');
      this.state.socket.close();
      this.setState({
        socket: null
      });
    }
  };
  restartSocket = () => {
    this.closeSocket();
    this.getLinks();
  };
  getLinks = () => {
    fetch('http://localhost:8080/sendEmail', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json'
        }
      }).then(console.log)
    // let socket = new WebSocket('ws://localhost:8080/scrape');

    // socket.onopen = () => {
    //   this.setState({
    //     socket: socket
    //   });
    //   console.log('open');
    // };
    // socket.onmessage = e => {
    //     this.onEvent(JSON.parse(e.data));
    // };
    // socket.onclose = () => {
    //   console.log('closing');

    //   socket.close();
    //   console.log('closed');
    //   this.setState({
    //     socket: null
    //   });
    // };
  };

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
  onClickRemove = clickedListing => {
    let { listings } = this.state;
    _.remove(listings, function(listing) {
      return listing.id === clickedListing.id;
    });
    this.setState({ listings: listings });
  };
  onClickRemoveEmailResponse = emailResponse => {
    let { responses } = this.state;
    _.remove(responses, function(response) {
      return emailResponse.id === response.id;
    });
    this.setState({ emailResponses: responses });
  };
  addEmailQueue = clickedListing => {
    let { emailQueue } = this.state;
    emailQueue.push(clickedListing);
    this.setState({
      emailQueue: emailQueue
    });
    this.onClickRemove(clickedListing);
  };

  render() {
    const {emailResponses} = this.state
    return (
      <div className="App">
        <Container textAlign="center">
          <h1 className="ui header fontWeight100">Craigslist Web Scraper</h1>
          <ProgressBar
            uploadState={'uploading'}
            percentUploaded={this.state.clListingScrapePercentage}
          />

          <Grid columns={3} divided>
            <Grid.Row>
              <Grid.Column width={4}>
               
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
                        </Button.Group>
                      </Card.Meta>
                      <Card.Description className="p5">
                        <Button.Group>
                          <Button onClick={this.closeSocket} icon>
                            <Icon name="clock" />
                          </Button>

                          <Button
                            active={this.state.emailAll}
                            onClick={() =>
                              this.setState({ emailAll: !this.state.emailAll })
                            }
                            icon
                          >
                            <Icon name="envelope" />
                          </Button>
                        </Button.Group>
                        <input
                          className="m-t-10"
                          type="text"
                          onChange={this.onInputChange}
                          name="filter-field"
                          placeholder="filter results"
                        />
                        {/* // <p className="margin0 padding0" align="left">
      //   <Icon name="close icon" onClick={this.onClickRemove}></Icon>
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
              <i>{`link count: ${this.state.linkCount}`}</i>
                {/* {this.state.listings && (
                  <Listings
                    onClickEmail={listing => this.addEmailQueue(listing)}
                    onClickRemove={listing => this.onClickRemove(listing)}
                    listings={this.filterList(this.state.listings)}
                  />
                )} */}
              </Grid.Column>
              <Grid.Column width={6}>
                {this.state.emailResponses.length > 0 && (
                  <HttpResponseList
                    onClickRemove={listing =>
                      this.onClickRemoveEmailResponse(listing)
                    }
                    responses={this.state.emailResponses}
                  />
                )}
              </Grid.Column>
            </Grid.Row>
          </Grid>
        </Container>
      </div>
    );
  }
}

export default App;
