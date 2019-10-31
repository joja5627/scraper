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
    var items = this.props.items.map((item, index) => {
      return (
        <TodoListItem
          key={index}
          item={item}
          index={index}
          removeItem={this.props.removeItem}
        />
      );
    });
    return <span>{items}</span>;
  }
}

class TodoListItem extends React.Component {
  constructor(props) {
    super(props);
    this.onClickClose = this.onClickClose.bind(this);
  }
  onClickClose() {
    var index = parseInt(this.props.index);
    this.props.removeItem(index);
  }

  render() {
    return (
      <p className="margin0 padding0" align="left">
        <Icon name="close icon" onClick={this.onClickClose}></Icon>
        <a href={`${this.props.item}`}> {this.props.item}</a>
      </p>
    );
  }
}

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      links: [],
      eventTarget: null,
      webSocket: null,
      linkViewLimit: 300,
      linkCount: 0
    };
    this.buildLinkEl = this.buildLinkEl.bind(this);
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
    } else if (event.messageType === 'listingURLs') {
      if (!this.state.links.includes(event.payload)) {
        let newLinks = this.state.links;
        newLinks.push(event.payload);
        this.setState({ links: newLinks, linkCount: newLinks.length });
      }
    }
  };

  closeSocket = () => {
    if (this.state.socket) {
      console.log('closing');
      this.state.socket.close();
      console.log('closed');
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

  buildLinkEl(links) {
    var items = links.map((item, index) => {
      console.log(this.removeItem);
      return (
        <TodoListItem
          key={index}
          item={item}
          index={index}
          removeItem={this.removeItem}
        />
      );
    });
    return <ul className="list-group"> {items} </ul>;
  }
  onInputChange = event => {
    this.setState({ eventTarget: event.target.value.toLowerCase() });
  };

  filterFunction = link => {
    return link.toLowerCase().search(this.state.eventTarget) !== -1;
  };

  filterList = links => {
    if (this.state.eventTarget) {
      return links.filter(link => this.filterFunction(link));
    } else {
      return links;
    }
  };
  limitResults = links => {
      return links.
  }

  render() {
    const { links } = this.state;

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

          <Grid columns={2} divided>
            <Grid.Row>
              <Grid.Column width={3}>
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
                        </Button.Group>
                      </Card.Meta>
                      <Card.Description>
                        <div className="ui action input">
                          <input type="text" placeholder="Search...">
                            <button className="ui button">Search</button>
                          </input>
                        </div>
                      </Card.Description>
                    </Card.Content>
                    <Card.Content extra>
                      <input
                        type="text"
                        onChange={this.onInputChange}
                        name="filter-field"
                        placeholder="filter results"
                      />
                    </Card.Content>
                  </Form>
                </Card>
              </Grid.Column>

              <Grid.Column width={12}>
                {links && (
                  <TodoList
                    items={this.filterList(this.state.links)}
                    removeItem={this.removeItem}
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
