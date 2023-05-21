import React from 'react';
import cl from '../css/insurance.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Insurance extends React.Component {
    constructor(props) {
        super(props);
        this.state = { name: "", sum: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/user/insurance", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": Cookies.get('Token')
            },
            body: JSON.stringify({ "name": this.state.name, "sum": this.state.sum })
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
                    <p className={cl.author}>Страхование</p>
                    <hr />
                    <input required type="text" value={this.state.name} onChange={(e) => this.setState({ name: e.target.value })} placeholder="Наименование имущества" className={cl.name}></input>
                    <input required type="text" value={this.state.sum} onChange={(e) => this.setState({ sum: e.target.value })} placeholder="Сумма" className={cl.name}></input>

                    <button type="submit" className={cl.open}>Застраховать</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Insurance