import React from 'react';
import cl from '../css/roles.module.css';
import { NavLink } from "react-router-dom";
import Сookies from 'universal-cookie';
import AllUsers from './AllUsers';

class Roles extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            accounts: [],
            role: "",
            id: "",
            error: ""
        };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        if (this.state.role == "manager") {
            const cookies = new Сookies();
            const response = await fetch("http://localhost:1337/manage/update-user", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id, "role": this.state.role })
            })
            const data = await response.json()
            if (data["error"] == "") {
            }
            else {
                this.setState({ error: data["error"] })
            }
        }
        else if (this.state.role == "user") {
            const cookies = new Сookies();
            const response = await fetch("http://localhost:1337/manage/update-user", {
                method: "POST",
                headers: {
                    'Accept': 'application/json',
                    'Content-type': 'application/json',
                    "Token": cookies.get('Token')
                },
                body: JSON.stringify({ "id": this.state.id, "role": this.state.role })
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
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 180, left: 170 }}>{this.state.error}</p></div>
                    <form className={cl.form} onSubmit={this.onSubmit}>
                        <input required type="text" onChange={(e) => this.setState({ id: e.target.value })} placeholder="Номер пользователя" className={cl.input}></input>
                        <button className={cl.button1} type="submit" onClick={(e) => this.setState({ role: "manager" })}>Сделать manager</button>
                        <button className={cl.button2} type="submit" onClick={(e) => this.setState({ role: "user" })}>Сделать user</button>
                    </form>
                    <div className={cl.information}>

                        <AllUsers />

                    </div>
                </div>
            </div>
        );
    }
}

export default Roles