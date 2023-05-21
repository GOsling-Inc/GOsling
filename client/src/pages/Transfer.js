import React from 'react';
import cl from '../css/transfer.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Transfer extends React.Component {
    constructor(props) {
        super(props);
        this.state = { Sender: "", Receiver: "", Amount: 0, error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/user/transfer", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": Cookies.get('Token')
            },
            body: JSON.stringify({ "Sender": this.state.Sender, "Receiver": this.state.Receiver, "sum": this.state.Amount })
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
                    <NavLink to="/user"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                <form className={cl.form} onSubmit={this.onSubmit}>
                    <div>{this.state.error}</div>
                    <p className={cl.author}>Перевод средств</p>
                    <hr />
                    <input required type="text" value={this.state.Sender} onChange={(e) => this.setState({ Sender: e.target.value })} placeholder="Счет отправителя" className={cl.name} style={{ marginTop: 40 }}></input>
                    <input required type="text" value={this.state.Receiver} onChange={(e) => this.setState({ Receiver: e.target.value })} placeholder="Счет получателя" className={cl.name}></input>
                    <input required type="number" onChange={(e) => this.setState({ Amount: e.target.value })} placeholder="Сумма перевода" className={cl.name}></input>
                    <button className={cl.open} type="submit">Перевести</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Transfer