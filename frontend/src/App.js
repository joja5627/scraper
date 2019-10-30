import React, { Component } from "react";
import "./App.css";
import { Button } from 'semantic-ui-react'
import { Progress } from "semantic-ui-react";


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

          <button type="button" className="close" onClick={this.onClickClose}>&times;</button>
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
      this.setState({ clListingScrapePercentage: event.payload })
    } else if (event.messageType === "contactInfoPercentComplete") {
      this.setState({ contactInfoScrapePercentage: event.payload })
    } else if(event.messageType === "listingReqComplete"){
      console.log(event)
      // var flatLinks = JSON.parse(event.payload).flat();
      // var uniqueLinks = Array.from(new Set(flatLinks));
      this.setState({ links: JSON.parse(event.payload) })
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
    var todoItems = [];
    todoItems.push({ index: 1, value: "learn react", done: false });
    todoItems.push({ index: 2, value: "Go shopping", done: true });
    todoItems.push({ index: 3, value: "buy flowers", done: true });
    console.log()
    return (
      <div className="App">

        <div className="container h-100">
          <div style={progressBarStyle}>
            <div className="row">
              <h2> Parsing Steps </h2>
              <div className="row">

                <h2>Listing Scrape Progress </h2>
                <div className="margin-left-30 padding-10">
                  {this.state.clListingScrapePercentage && <ProgressBar uploadState={"uploading"} percentUploaded={this.state.clListingScrapePercentage} />}
                </div>

              </div>
              <div className="row">
                <h2>Contact Info Scrape Progress</h2>
                <div className="margin-left-30 padding-10">
                  {this.state.contactInfoScrapePercentage && <ProgressBar uploadState={"uploading"} percentUploaded={this.state.contactInfoScrapePercentage} />}
                </div>
              </div>
            </div>
          </div>

          <div className="row">
            <div className="col">
              <input type="text" onChange={this.onInputChange} name="filter-field" placeholder="filter results" />

              <div className="row">
              
                <button className="ui fade animated button" onClick={this.getLinks}>
                  <div className="visible content">scrape cl</div>
                  <div className="hidden content">scraping</div>
                </button>
                {links && <TodoList items={this.filterList(this.state.links)} removeItem={this.removeItem} />}
              </div>
            </div>
          </div>
        </div>


      </div>
    );
  }
}

export default App;
