import React from 'react';
import cl from '../css/registration.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';

class Registration extends React.Component {
    constructor(props) {
        super(props);
        this.state = { Name: "", Surname: "", Email: "", Password: "", passwordConfirm: "", Birthdate: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        if (this.state.Password == this.state.passwordConfirm) {
            this.setState({ error: "" })
            const response = await fetch("http://localhost:1337/sign-up", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                },
                body: JSON.stringify({ "Name": this.state.Name, "Surname": this.state.Surname, "Email": this.state.Email, "Password": this.state.Password, "Birthdate": this.state.Birthdate })
            })
            const data = await response.json()
            if (data["error"] == "") {
                const cookies = new Cookies();
                cookies.set('Token', data["data"]["Token"], { path: '/' });
                window.location.href = "/user";
            }
            else {
                this.setState({ error: data["error"] })
                console.log(data["error"])
            }
        }
        else {
            this.setState({ error: "Пароли не совпадают" })
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
                    <p className={cl.author}>Регистрация</p>
                    <hr />
                    <div style={{ marginTop: 0, height: 0 }}><div style={{ color: "red", position: "relative", top: 15 }}>{this.state.error}</div></div>
                    <input required type="text" value={this.state.Name} onChange={(e) => this.setState({ Name: e.target.value })} placeholder="Имя" ></input>
                    <input required type="text" value={this.state.Surname} onChange={(e) => this.setState({ Surname: e.target.value })} placeholder="Фамилия"></input>
                    <input required type="date" value={this.state.Birthdate} onChange={(e) => this.setState({ Birthdate: e.target.value })} className={cl.date}></input>
                    <input required type="email" value={this.state.Email} onChange={(e) => this.setState({ Email: e.target.value })} placeholder="Почта"></input>
                    <input required type="password" value={this.state.Password} onChange={(e) => this.setState({ Password: e.target.value })} placeholder="Пароль"></input>
                    <input required type="password" value={this.state.passwordConfirm} onChange={(e) => this.setState({ passwordConfirm: e.target.value })} placeholder="Подтверждение пароля"></input>
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