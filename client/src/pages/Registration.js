import React from 'react';
import cl from '../css/registration.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Registration extends React.Component {
    constructor(props) {
        super(props);
        this.state = { firstName: "", lastName: "", date: "", email: "", password: "", passwordConfirm: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/sign-up", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
            },
            body: JSON.stringify({ "Name": this.state.firstName, "Surname": this.state.lastName, "Birthdate": this.state.date, "Email": this.state.email, "Password": this.state.password })
        })
        const data = await response.json()
        if (data["error"] == "") {
            const cookies = new Cookies();
            cookies.set('Token', data["data"]["Token"], { path: '/' });
            window.location.href = "/user";
        }
        else {
            this.setState({ error: data["error"] })
        }
    }



    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/"><h1 className={cl.head}>GOsling</h1></NavLink>
                    <NavLink to="/authorization"><button className={cl.ent}>Войти</button></NavLink>
                </div>
                <form className={cl.form} onSubmit={this.onSubmit}>
                    <div>{this.state.error}</div>
                    <p className={cl.author}>Регистрация</p>
                    <hr />
                    <input required type="text" value={this.state.firstName} onChange={(e) => this.setState({firstName: e.target.value})} placeholder="Имя" ></input>
                    <input required type="text" value={this.state.lastName} onChange={(e) => this.setState({lastName: e.target.value})} placeholder="Фамилия"></input>
                    <input required type="date" value={this.state.date} onChange={(e) => this.setState({date: e.target.value})} className={cl.date}></input>
                    <input required type="email" value={this.state.email} onChange={(e) => this.setState({email: e.target.value})} placeholder="Почта"></input>
                    <input required type="password" value={this.state.password} onChange={(e) => this.setState({password: e.target.value})} placeholder="Пароль"></input>
                    <input required type="password" value={this.state.passwordConfirm} onChange={(e) => this.setState({passwordConfirm: e.target.value})} placeholder="Подтверждение пароля"></input>
                    <button type="submit">Зарегистрироваться</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Registration