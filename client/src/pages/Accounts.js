import React from 'react';
import cl from '../css/accounts.module.css';
import { NavLink } from "react-router-dom";
import Сookies from 'universal-cookie';
import AllAccForManager from './AllAccForManager';


class Accounts extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            accounts: [],
            what: "",
            id: "",
            error: ""
        };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        if (this.state.what == "freeze") {
            const cookies = new Сookies();
            const response = await fetch("http://localhost:1337/manage/freeze-account", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id })
            })
            const data = await response.json()
            if (data["error"] == "") {
            }
            else {
                this.setState({ error: data["error"] })
            }
        }
        else if (this.state.what == "block") {
            const cookies = new Сookies();
            const response = await fetch("http://localhost:1337/manage/block-account", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id })
            })
            const data = await response.json()
            if (data["error"] == "") {
            }
            else {
                this.setState({ error: data["error"] })
            }

        }
    }


    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>

                <div>
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 100, left: 250 }}>{this.state.error}</p></div>
                    <div className={cl.information}>

                        <AllAccForManager />

                    </div>
                    <form className={cl.form2} onSubmit={this.onSubmit}>
                        <input required type="text" onChange={(e) => this.setState({ id: e.target.value })} placeholder="Номер счёта" className={cl.input}></input>
                        <button className={cl.button} type="submit" onClick={(e) => this.setState({ what: "freeze" })}>Заморозить</button>
                    </form>
                    <form className={cl.form3} onSubmit={this.onSubmit}>
                        <input required type="text" onChange={(e) => this.setState({ id: e.target.value })} placeholder="Номер счёта" className={cl.input}></input>
                        <button className={cl.button} type="submit" onClick={(e) => this.setState({ what: "block" })}>Заблокировать</button>
                    </form>

                </div>
            </div>
        );
    }
}

export default Accounts