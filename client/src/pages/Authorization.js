import React from 'react';
import cl from '../css/authoriz.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Authorization extends React.Component {
    constructor(props) {
        super(props);
        this.state = {email: "", password: "", error: ""};
        this.onSubmit = this.onSubmit.bind(this) 
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await  fetch("http://localhost:1337/sign-in", {
            method: "POST",
            headers: {
              'Accept': 'application/json',
              'Content-type': 'application/json',
            },
            body: JSON.stringify({"email": this.state.email, "password": this.state.password})
        })
        
        const data = await response.json()
        if (data["error"] == "") {
            const cookies = new Cookies();
            cookies.set('Token', data["data"]["Token"], { path: '/' });
            window.location.href = "/user";
        }
        else {
            this.setState({error: data["error"]})
        }
    }

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/"><h1 className={cl.head}>GOsling</h1></NavLink>
                    <NavLink to="/registration"><button className={cl.reg}>Регистрация</button></NavLink>
                </div>
                <form className={cl.form} onSubmit={this.onSubmit}>
                    <div>{this.state.error}</div>
                    <p className={cl.author}>Авторизация</p>
                    <hr />
                    <input type="email" value={this.state.email} onChange={(e) => this.setState({email: e.target.value})} placeholder="Почта" ></input>
                    <input type="text" value={this.state.password} onChange={(e) => this.setState({password: e.target.value})} placeholder="Пароль"></input>
                    <button type="submit">Войти</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Authorization