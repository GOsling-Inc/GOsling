import React from 'react';
import cs from '../css/exchange.module.css';
import { NavLink } from "react-router-dom";
import { useState } from "react";
import Cookies from 'universal-cookie';


class Exchange extends React.Component {
    constructor(props) {
        super(props);

        fetch("http://localhost:1337/exchanges", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
            },
        }).then(res => res.json()).then(data => console.log(data))

        this.state = {
            accounts: [],
            SenderAmount: 0,
            Sender: "",
            Receiver: "",
            error: ""
        };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const cookies = new Cookies();
        const response = await fetch("http://localhost:1337/user/exchange", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
            body: JSON.stringify({ "Sender": this.state.Sender, "Receiver": this.state.Receiver, "SenderAmount": this.state.SenderAmount })
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
                        <p className={cs.author}>Валюта</p>
                        <hr />
                        <div style={{ marginTop: 0, height: 0 }}><p style={{ color: "red", position: "relative", top: 15, textAlign: "center" }}>{this.state.error}</p></div>
                        <input required type="number" onChange={(e) => this.setState({ SenderAmount: e.target.value })} placeholder="Сумма" className={cs.name}></input>
                        <input required type="text" value={this.state.Sender} onChange={(e) => this.setState({ Sender: e.target.value })} placeholder="Номер счёта 1 валюты" className={cs.name}></input>
                        <input required type="text" value={this.state.Receiver} onChange={(e) => this.setState({ Receiver: e.target.value })} placeholder="Номер счёта 2 валюты" className={cs.name}></input>

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