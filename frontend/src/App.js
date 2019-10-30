import React, { Component } from "react";
import "./App.css";
import { Button, Container, Grid, Card, Progress, Form,Icon, Input, Select, TextArea } from 'semantic-ui-react'


const ProgressBar = ({ uploadState, percentUploaded }) =>
  uploadState === "uploading" && (
    <Progress
      percent={percentUploaded}
      progress
      indicating
      size="medium"
      inverted
    />
  );


const buttonStyle = {
  margin: "10px"
};
const headerMarginStyle = {
  marginBottom: "100px"
};
const progressBarStyle = {
  marginLeft: "100px",
  marginRight: "100px"
}

class TodoList extends React.Component {
  render() {
    var items = this.props.items.map((item, index) => {
      return (
        <TodoListItem key={index} item={item} index={index} removeItem={this.props.removeItem} />
      );
    });
    return (
      <ul className="list-group"> {items} </ul>
    );
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
      <li className="list-group-item ">
        <div>
          <span className="glyphicon glyphicon-ok icon" aria-hidden="true" onClick={this.onClickDone}>
            <a href={`${this.props.item}`}> {this.props.item}</a>
          </span>

          <Button onClick={this.onClickClose}>&times;</Button>
        </div>
      </li>
    );
  }
}

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      links: [],
      eventTarget: null
    };
    this.buildLinkEl = this.buildLinkEl.bind(this)
    this.removeItem = this.removeItem.bind(this);
    this.clListingScrapePercentage = 0
    this.contactInfoScrapePercentage = 0


  }
  removeItem(itemIndex) {
    let newLinks = this.state.links
    newLinks.splice(itemIndex, 1);
    this.setState({ links: newLinks });
  }
  onEvent = (event) => {
    if (event.messageType === "listingPercentComplete") {
      this.setState({ clListingScrapePercentage: Math.round(event.payload) })
    } else if (event.messageType === "listingURLs") {
      this.setState({ links: this.state.links.concat(event.payload)})
    }
  }
  getLinks = () => {
    const socket = new WebSocket('ws://localhost:8080/scrape')
    socket.onopen = () => {
      console.log("open")
    }
    socket.onmessage = e => {

      this.onEvent(JSON.parse(e.data))
    }
    socket.onclose = () => {
      console.log("closed")
    }


    // let ws = 
    // 



    // fetch("http://localhost:8080/scrape")
    //   .then(res => {
    //     return res.json()
    //   })
    //   .then(res => {
    //     var flatLinks = JSON.parse(res.listingUrl).flat();
    //     var uniqueLinks = Array.from(new Set(flatLinks));
    //     this.setState({ links: uniqueLinks })
    //   })
    //   .catch(err => err);
  }

  buildLinkEl(links) {
    var items = links.map((item, index) => {
      console.log(this.removeItem)
      return (
        <TodoListItem key={index} item={item} index={index} removeItem={this.removeItem} />
      );
    });
    return (
      <ul className="list-group"> {items} </ul>
    );
  }
  onInputChange = (event) => {
    this.setState({ eventTarget: event.target.value.toLowerCase() })
  }

  filterFunction = (link) => {
    return link.toLowerCase().search(
      this.state.eventTarget) !== -1;
  };


  filterList = (links) => {
    if (this.state.eventTarget) {
      return links.filter(link => this.filterFunction(link));
    } else {
      return links
    }

  }




  render() {
    const { links } = this.state

    return (
      <div className="App">
        <Container textAlign='center'>
          <h1 className="ui header fontWeight100">Craigslist Web Scraper</h1>
          {this.state.clListingScrapePercentage && <ProgressBar uploadState={"uploading"} percentUploaded={this.state.clListingScrapePercentage} />}

          <Grid columns={3} divided>
            <Grid.Row>
              <Grid.Column>
                <Card>
                  <Form>
                    <Card.Content>
                      <Card.Header>CL input parameters</Card.Header>
                      <Card.Meta>
                        <Form.Group widths='equal'>
                        </Form.Group>
                      </Card.Meta>
                      <Card.Description>
                      <Button.Group>
                      <Button onClick={this.getLinks} icon>
                        <Icon name='play' />
                      </Button>
                      <Button icon>
                        <Icon name='stop' />
                      </Button>
                      <Button icon>
                        <Icon name='redo' />
                      </Button>
                     
                    </Button.Group>
                    
                  </Card.Description>
                    </Card.Content>
                    <Card.Content extra>
                      <input type="text" onChange={this.onInputChange} name="filter-field" placeholder="filter results" />
                    </Card.Content>
                  </Form>
                </Card>
              </Grid.Column>
              <Grid.Column>
              {links && <TodoList items={this.filterList(this.state.links)} removeItem={this.removeItem} />}
              </Grid.Column>
              <Grid.Column>
              </Grid.Column>
            </Grid.Row>
          </Grid>
        </Container>



      </div>
    );
  }
}

export default App;
