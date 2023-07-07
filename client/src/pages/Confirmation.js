import React from 'react';
import cl from '../css/confirmation.module.css';
import { NavLink } from "react-router-dom";
import Cookies from 'universal-cookie';
import AllConfirms from './AllConfirms';

class Confirmation extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            accounts: [],
            status: "",
            id: "",
            table: "",
            error: ""
        };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        if (this.state.status == "ACTIVE") {
            const cookies = new Cookies();
            const response = await fetch("http://localhost:1337/manage/confirmation", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id, "value": this.state.table, "status": this.state.status })
            })
            const data = await response.json()
            if (data["error"] == "") {
            }
            else {
                this.setState({ error: data["error"] })
            }
        }
        else if (this.state.status == "BLOCKED") {
            const cookies = new Cookies();
            const response = await fetch("http://localhost:1337/manage/confirmation", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id, "value": this.state.table, "status": this.state.status })
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
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 140, left: 260 }}>{this.state.error}</p></div>
                    <form className={cl.form} onSubmit={this.onSubmit}>
                        <input required type="text" onChange={(e) => this.setState({ id: e.target.value })} placeholder="Id" className={cl.input}></input>
                        <input required type="text" onChange={(e) => this.setState({ table: e.target.value })} placeholder="table" className={cl.input}></input>
                        <button className={cl.button1} type="submit" onClick={(e) => this.setState({ status: "ACTIVE" })}>Подтвердить</button>
                        <button className={cl.button2} type="submit" onClick={(e) => this.setState({ status: "BLOCKED" })}>Отказать</button>
                    </form>
                    <div className={cl.loan}>
                        <p className={cl.loanTitle}>Подтверждение</p>
                        <hr />
                        <AllConfirms />

                    </div>
                </div>



            </div>
        );
    }
}

export default Confirmation