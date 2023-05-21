import React from 'react';
import cs from '../css/exchange.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';


class Exchange extends React.Component {
    constructor(props) {
        super(props);
        this.state = { valutaType1: "BYN", valutaType2: "BYN", sum: "", accountNumber1: "", accountNumber2: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/user/exchange", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": Cookies.get('Token')
            },
            body: JSON.stringify({ "valutaType1": this.state.valutaType1, "valutaType2": this.state.valutaType2, "sum": this.state.sum, "accountNumber1": this.state.accountNumber1, "accountNumber2": this.state.accountNumber2 })
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

                <div className={cs.together}>
                    <form className={cs.form} onSubmit={this.onSubmit}>
                        <div>{this.state.error}</div>
                        <p className={cs.author}>Валюта</p>
                        <hr />
                        <p className={cs.selectType} >Выберите тип 1 валюты</p>
                        <select className={cs.valuta} value={this.state.valutaType1} onChange={(e) => this.setState({ valutaType1: e.target.value })}>
                            <option value="BYN">BYN</option>
                            <option value="USD">USD</option>
                            <option value="EUR">EUR</option>
                        </select>
                        <p className={cs.selectType} >Выберите тип 2 валюты</p>
                        <select className={cs.valuta} value={this.state.valutaType2} onChange={(e) => this.setState({ valutaType2: e.target.value })}>
                            <option value="BYN">BYN</option>
                            <option value="USD">USD</option>
                            <option value="EUR">EUR</option>
                        </select>
                        <input required type="text" value={this.state.sum} onChange={(e) => this.setState({ sum: e.target.value })} placeholder="Сумма" className={cs.name}></input>
                        <input required type="text" value={this.state.accountNumber1} onChange={(e) => this.setState({ accountNumber1: e.target.value })} placeholder="Номер счёта 1 валюты" className={cs.name}></input>
                        <input required type="text" value={this.state.accountNumber2} onChange={(e) => this.setState({ accountNumber2: e.target.value })} placeholder="Номер счёта 2 валюты " className={cs.name}></input>

                        <button className={cs.open} type="submit">Выполнить</button>
                    </form>

                    <div className={cs.rate}>
                        <p className={cs.pRate}>Курсы</p>
                        <div>
                            <div className={cs.val}>
                                <div className={cs.exch}>
                                    <p className={cs.pValuta}>Валюта</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p >USD</p>
                                </div>
                                <div className={cs.euro}>
                                    <p >EUR</p>
                                </div>
                            </div>
                            <div className={cs.buy}>
                                <div className={cs.exch}>
                                    <p >Покупка</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p style={{ fontSize: 22 }}>?</p>
                                </div>
                                <div className={cs.euro}>
                                    <p style={{ fontSize: 22 }}>?</p>
                                </div>
                            </div>
                            <div className={cs.sell}>
                                <div className={cs.exch}>
                                    <p >Продажа</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p style={{ fontSize: 22 }}>?</p>
                                </div>
                                <div className={cs.euro}>
                                    <p style={{ fontSize: 22 }}>?</p>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                    <NavLink className={cs.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Exchange