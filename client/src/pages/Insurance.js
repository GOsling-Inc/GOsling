import React from 'react';
import cl from '../css/insurance.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Insurance extends React.Component {
    constructor(props) {
        super(props);
        this.state = { AccountId: "", Amount: 0, Period: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Cookies();
        const response = await fetch("http://localhost:1337/user/new-insurance", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
            body: JSON.stringify({ "AccountId": this.state.AccountId, "Period": this.state.Period, "Amount": this.state.Amount })
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
                    <p className={cl.author}>Страхование</p>
                    <hr />
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center"}}>{this.state.error}</p></div>
                    <input required type="text" value={this.state.AccountId} onChange={(e) => this.setState({ AccountId: e.target.value })} placeholder="Номер счёта" className={cl.name}></input>
                    <div>
                        <input required type="number" onChange={(e) => this.setState({ Amount: e.target.value })} placeholder="Сумма" className={cl.amount}></input>
                        <input required type="text" value={this.state.Period} onChange={(e) => this.setState({ Period: e.target.value })} placeholder="Срок" className={cl.period}></input>
                    </div>
                    <button type="submit" className={cl.open}>Застраховать</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Insurance