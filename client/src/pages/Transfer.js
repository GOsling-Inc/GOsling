import React from 'react';
import cl from '../css/transfer.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Transfer extends React.Component {
    constructor(props) {
        super(props);
        this.state = { sendAccount: "", beneficAccount: "", sum: "", error: "" };
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
            body: JSON.stringify({ "sendAccount": this.state.sendAccount, "beneficAccount": this.state.beneficAccount, "sum": this.state.sum })
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
                    <input required type="text" value={this.state.sendAccount} onChange={(e) => this.setState({sendAccount: e.target.value})} placeholder="Счет отправителя" className={cl.name} style={{ marginTop: 40 }}></input>
                    <input required type="text" value={this.state.beneficAccount} onChange={(e) => this.setState({beneficAccount: e.target.value})} placeholder="Счет получателя" className={cl.name}></input>
                    <input required type="text" value={this.state.sum} onChange={(e) => this.setState({sum: e.target.value})} placeholder="Сумма перевода" className={cl.name}></input>
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