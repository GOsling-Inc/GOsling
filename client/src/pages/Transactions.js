import React from 'react';
import cl from '../css/transactions.module.css';
import { NavLink } from "react-router-dom";
import Сookies from 'universal-cookie';


class Transactions extends React.Component {
    constructor(props) {
        super(props);
        this.state = { id: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Сookies();
        const response = await fetch("http://localhost:1337/manage/cancel-transaction", {
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

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                
                <form className={cl.form} onSubmit={this.onSubmit}>
                    <input required type="text" value={this.state.id} onChange={(e) => this.setState({ id: e.target.value })} placeholder="Номер транзакции" className={cl.input}></input>
                    <button className={cl.button} type="submit">Отменить</button>
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center"}}>{this.state.error}</p></div>
                </form>
            </div>
        );
    }
}

export default Transactions