import React from 'react';
import cs from '../css/loan.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"
import { useState } from "react";
import Cookies from 'universal-cookie';


class Loan extends React.Component {
    constructor(props) {
        super(props);
        this.state = { Amount: 0, Period: "", Percent: 0, AccountId: "", error: "" };
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
            body: JSON.stringify({ "AccountId": this.state.AccountId, "Period": this.state.Period, "Amount ": this.state.Amount, "Percent  ": this.state.Percent })
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
                    <input required type="number" onChange={(e) => this.setState({ Amount: e.target.value })} placeholder="Сумма" className={cs.name}></input>

                    <p className={cs.pType} style={{ marginTop: 0 }}>
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

export default Loan