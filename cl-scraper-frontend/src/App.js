import React, { Component } from "react";
import logo from "./logo.svg";
import solute_logo from "./S&GLogoDope.png";


import "./App.css";




const buttonStyle = {
    margin: "10px"
};
const headerMarginStyle = {
    marginBottom: "100px"
};


class App extends Component {
    constructor(props) {
        super(props);
        this.state = {
            links: [],
            eventTarget: null
        };
    }

    getLinks = () => {
        fetch("http://localhost:8080/scrape")
            .then(res => {
                return res.json()
            })
            .then(res => {
                var flatLinks = JSON.parse(res.links).flat();
                var uniqueLinks =  Array.from(new Set(flatLinks));
                this.setState({ links: uniqueLinks })
            })
            .catch(err => err);
    }
    buildLinkEl = links => {
        return links.map(function (link, index) {
            return <span key={index}><a key={index} href={`${link}`}> {link}</a><br></br></span>
        })
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
                <header style={headerMarginStyle} className="App-header">
                    <img src={solute_logo} className="App-logo" alt="logo" />
                </header>
                <div className="container h-100">
                    <div className="row align-items-center h-100">
                        <div className="col-10 mx-auto">
                        <input type="text" onChange={this.onInputChange} name="filter-field" placeholder="filter results" />

                            <div className="row justify-content-center align-self-center">
                               
                                <button className="btn btn-2 btn-2a" onClick={this.getLinks}> pull links </button>
                                {links && this.buildLinkEl(this.filterList(this.state.links))}
                            </div>
                        </div>
                    </div>
                </div>


            </div>
        );
    }
}

export default App;
