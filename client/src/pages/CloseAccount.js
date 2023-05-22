import React from 'react';
import cl from '../css/insurance.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class CloseAccount extends React.Component {
    constructor(props) {
        super(props);
        this.state = { accountId : "", password : "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Cookies();
        const response = await fetch("http://localhost:1337/user/delete-account", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
            body: JSON.stringify({ "accountId ": this.state.accountId , "password  ": this.state.password })
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
                    <p className={cl.author}>Закрыть счёт</p>
                    <hr />
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center"}}>{this.state.error}</p></div>
                    <input required type="text" value={this.state.accountId } onChange={(e) => this.setState({ accountId : e.target.value })} placeholder="Номер счёта" className={cl.name}></input>
                    <input required type="password" onChange={(e) => this.setState({ password : e.target.value })} placeholder="Пароль" className={cl.name}></input>
                    <button type="submit" className={cl.open}>Закрыть</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default CloseAccount