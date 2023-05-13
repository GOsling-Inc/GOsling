import React from 'react';
import cs from '../css/loan.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"

class Loan extends React.Component {

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <div className={cs.form}>
                    <p className={cs.author}>Кредитование</p>
                    <hr />
                    <p className={cs.pType} style={{marginTop: 20}}>Выберите тип валюты</p>
                    <select className={cs.valuta}>
                        <option value="BYN">BYN</option>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                    <input placeholder="Сумма" className={cs.name}></input>

                    <p className={cs.pType} style={{marginTop: 30}}>Выберите вид кредита
                        <div className={cs.cl}><img src={photoQuestion} className={cs.what} />
                            <div className={cs.clue}>
                                <p style={{letterSpacing: 1, fontSize: 18}}>1) Потребительский:<br /><br />
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

                    <select className={cs.acc}>
                        <option value="Потребительский">Потребительский</option>
                        <option value="Ссуда под развитие бизнеса">Ссуда под развитие бизнеса</option>
                        <option value="Ипотека">Ипотека</option>
                    </select>

                    <input placeholder="Номер счёта" className={cs.name}></input>

                    <button className={cs.open}>Оформить</button>
                </div>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                    <NavLink className={cs.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Loan