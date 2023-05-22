import React from 'react';
import cs from '../css/deposits.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"
import { useState } from "react";
import Cookies from 'universal-cookie';


class Deposits extends React.Component {
    constructor(props) {
        super(props);
        this.state = { Amount: 0, Percent: 0, Period: "", AccountId: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Cookies();
        const response = await fetch("http://localhost:1337/user/deposits", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
            body: JSON.stringify({ "AccountId": this.state.AccountId, "Period": this.state.Period, "Amount ": this.state.Amount, "Percent ": this.state.Percent })
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
                    <p className={cs.author}>Вклад</p>
                    <hr />
                    <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center"}}>{this.state.error}</p></div>
                    <input required type="number" onChange={(e) => this.setState({ Amount: e.target.value })} placeholder="Сумма" className={cs.name}></input>

                    <p style={{ marginTop: 0 }}>
                        <div className={cs.cl}><img src={photoQuestion} className={cs.what} />
                            <div className={cs.clue}>
                                <p style={{ letterSpacing: 1, fontSize: 16 }}>1) Начальный:<br />
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

                    <div>
                        <input required type="number" onChange={(e) => this.setState({ Percent: e.target.value })} placeholder="Процент" className={cs.percent}></input>
                        <input required type="text" value={this.state.Period} onChange={(e) => this.setState({ Period: e.target.value })} placeholder="Срок" className={cs.period}></input>
                    </div>

                    <input required type="text" value={this.state.AccountId} onChange={(e) => this.setState({ AccountId: e.target.value })} placeholder="Номер счёта" className={cs.name}></input>

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