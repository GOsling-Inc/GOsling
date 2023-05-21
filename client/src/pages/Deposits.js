import React from 'react';
import cs from '../css/deposits.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"
import { useState } from "react";
import Cookies from 'universal-cookie';


class Deposits extends React.Component {
    constructor(props) {
        super(props);
        this.state = { valutaType: "BYN", sum: "", depositType: "Начальный", accountNumber: "",error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/user/deposits", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": Cookies.get('Token')
            },
            body: JSON.stringify({  "valutaType": this.state.valutaType, "sum": this.state.sum, "depositType": this.state.depositType, "accountNumber": this.state.accountNumber })
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
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <form className={cs.form} onSubmit={this.onSubmit}>
                <div>{this.state.error}</div>
                    <p className={cs.author}>Вклад</p>
                    <hr />
                    <p style={{marginTop: 20, fontSize: 17, marginLeft: 25, letterSpacing: 1}}>Выберите тип валюты</p>
                    <select className={cs.valuta} value={this.state.valutaType} onChange={(e) => this.setState({ valutaType: e.target.value })}>
                        <option value="BYN">BYN</option>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                    <input required type="text" value={this.state.sum} onChange={(e) => this.setState({ sum: e.target.value })} placeholder="Сумма" className={cs.name}></input>

                    <p style={{marginTop: 30, fontSize: 17, marginLeft: 25, letterSpacing: 1}}>Выберите вид вклада
                        <div className={cs.cl}><img src={photoQuestion} className={cs.what} />
                            <div className={cs.clue}>
                                <p style={{letterSpacing: 1, fontSize: 16}}>1) Начальный:<br />
                                    - 1 год<br />
                                    - 3% годовых<br /><br />
                                    2) Базовый:<br />
                                    - 2 года<br />
                                    - 4% годовых<br /><br />
                                    3) Продвинутый:<br />
                                    - 3 года<br />
                                    - 5% годовых<br /><br />
                                    4) Детский:<br />
                                    - до момента совершеннолетия<br />
                                    - 2% годовых<br /><br />
                                    5) Базовый USD/EUR:<br />
                                    - 1 год.<br />
                                    - 4% годовых.    <br /><br />
                                    6) Продвинутый USD/EUR:<br />
                                    - 2 года. <br />
                                    - 5% годовых.<br /><br />
                                </p>
                            </div>
                        </div></p>

                    <select className={cs.acc} value={this.state.depositType} onChange={(e) => this.setState({ depositType: e.target.value })}>
                        <option value="Начальный">Начальный</option>
                        <option value="Базовый">Базовый</option>
                        <option value="Продвинутый">Продвинутый</option>
                        <option value="Детский">Детский</option>
                        <option value="Базовый USD/EUR">Базовый USD/EUR</option>
                        <option value="Продвинутый USD/EUR">Продвинутый USD/EUR</option>
                    </select>

                    <input required type="text" value={this.state.accountNumber} onChange={(e) => this.setState({ accountNumber: e.target.value })} placeholder="Номер счёта" className={cs.name}></input>

                    <button className={cs.open} type="submit">Оформить</button>
                </form>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                    <NavLink className={cs.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Deposits