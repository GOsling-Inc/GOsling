import React from 'react';
import cs from '../css/loan.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"
import { useState } from "react";
import Cookies from 'universal-cookie';


class Loan extends React.Component {
    constructor(props) {
        super(props);
        this.state = { valutaType: "BYN", sum: "", loanType: "Потребительский", accountNumber: "", error: "" };
        this.onSubmit = this.onSubmit.bind(this)
    }

    async onSubmit(e) {
        e.preventDefault();
        const response = await fetch("http://localhost:1337/user/loan", {
            method: "POST",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": Cookies.get('Token')
            },
            body: JSON.stringify({ "valutaType": this.state.valutaType, "sum": this.state.sum, "loanType": this.state.loanType, "accountNumber": this.state.accountNumber })
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
                    <p className={cs.author}>Кредитование</p>
                    <hr />
                    <p className={cs.pType} style={{ marginTop: 20 }}>Выберите тип валюты</p>
                    <select className={cs.valuta} value={this.state.valutaType} onChange={(e) => this.setState({ valutaType: e.target.value })}>
                        <option value="BYN">BYN</option>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                    <input required type="text" value={this.state.sum} onChange={(e) => this.setState({ sum: e.target.value })} placeholder="Сумма" className={cs.name}></input>

                    <p className={cs.pType} style={{ marginTop: 30 }}>Выберите вид кредита
                        <div className={cs.cl}><img src={photoQuestion} className={cs.what} />
                            <div className={cs.clue}>
                                <p style={{ letterSpacing: 1, fontSize: 18 }}>1) Потребительский:<br /><br />
                                    - срок кредита до 5 лет.<br />
                                    - сумма кредита до 5000 BYN.<br />
                                    - процентная ставка 9%.<br /><br />
                                    2) Ссуда под развитие бизнеса:<br /><br />
                                    - срок кредита до 7 лет.<br />
                                    - сумма кредита от 20000 BYN до 100000 BYN.<br />
                                    - процентная ставка 7%.<br /><br />
                                    3) Ипотека:<br /><br />
                                    - срок кредита до 25 лет.<br />
                                    - сумма кредита до 1000000 BYN.<br />
                                    - процентная ставка 8%.<br />

                                </p>
                            </div>
                        </div></p>

                    <select className={cs.acc} value={this.state.loanType} onChange={(e) => this.setState({ loanType: e.target.value })}>
                        <option value="Потребительский">Потребительский</option>
                        <option value="Ссуда под развитие бизнеса">Ссуда под развитие бизнеса</option>
                        <option value="Ипотека">Ипотека</option>
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

export default Loan