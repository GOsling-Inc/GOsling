import React from 'react';
import cs from '../css/new-account.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"
import { useState } from "react";
import Сookies from 'universal-cookie';


class NewAccount extends React.Component {
constructor(props) {
        super(props);
        this.state = { Name: "", Unit: "BYN", Type: "Базовый счёт", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Сookies();
        const response = await fetch("http://localhost:1337/user/new-account", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
            body: JSON.stringify({ "Name": this.state.Name, "Unit": this.state.Unit, "Type": this.state.Type })
        })
        const data = await response.json()
        if (data["error"] == "") {
            window.location.href = "/user";
        }
        else {
            this.setState({ error: data["error"] })
        }
    }


    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <form className={cs.form} onSubmit={this.onSubmit}>
                    <p className={cs.author}>Открытие счёта</p>
                    <hr />
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center"}}>{this.state.error}</p></div>
                    <input required type="text" value={this.state.Name} onChange={(e) => this.setState({ Name: e.target.value })} placeholder="Название счёта" className={cs.name}></input>
                    <p className={cs.pType}>Выберите тип валюты</p>
                    <select className={cs.valuta} value={this.state.Unit} onChange={(e) => this.setState({ Unit: e.target.value })}>
                        <option value="BYN">BYN</option>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                    <p className={cs.pType}>Выберите тип счёта
                        <div className={cs.cl}><img src={photoQuestion} className={cs.what} />
                            <div className={cs.clue}>
                                <p >1) базовый счет: <br />
                                    - переводы до 10000 BYN (или эквивалент в переводе на другую валюту по курсу банка “GOsling”) для переводов<br />
                                    - комиссия 2% для переводов меньше 1000 BYN (5% в ином случае)<br />
                                    - бесплатное обслуживание<br />
                                    <br />
                                    <hr />
                                    <br />
                                </p>
                                <p >2) бизнес счет:<br />
                                    - неограниченная сумма переводов.<br />
                                    - стоимость обслуживания: 2% от общей суммы переводов за месяц (не ниже 100 BYN)<br />
                                    - комиссия 1.5% для любых переводов<br />
                                    <br />
                                    <hr />
                                    <br />
                                </p>
                                <p >3) инвестиционный счет:<br />
                                    - стоимость обслуживания: 3% от общей суммы сделок за месяц (при условии общей суммы сделок не менее 1000 BYN, <br />в противном случае 50 BYN)<br />
                                    - переводы до 100000 BYN (или эквивалент в переводе на другую валюту по курсу банка “GOsling”) для переводов <br />
                                    - комиссия 5% для переводов меньше 500  BYN (7% в ином случае)<br />
                                </p>
                            </div>
                        </div></p>

                    <select className={cs.acc} value={this.state.Type} onChange={(e) => this.setState({ Type: e.target.value })}>
                        <option value="Базовый счёт">Базовый счёт</option>
                        <option value="Бизнес счёт">Бизнес счёт</option>
                        <option value="Инвестиционный счёт">Инвестиционный счёт</option>
                    </select>
                    <button className={cs.open} type="submit">Открыть</button>
                </form>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                    <NavLink className={cs.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default NewAccount