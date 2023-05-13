import React from 'react';
import cs from '../css/deposits.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"

class Deposits extends React.Component {

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <div className={cs.form}>
                    <p className={cs.author}>Вклад</p>
                    <hr />
                    <p style={{marginTop: 20, fontSize: 17, marginLeft: 25, letterSpacing: 1}}>Выберите тип валюты</p>
                    <select className={cs.valuta}>
                        <option value="BYN">BYN</option>
                        <option value="USD">USD</option>
                        <option value="EUR">EUR</option>
                    </select>
                    <input placeholder="Сумма" className={cs.name}></input>

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

                    <select className={cs.acc}>
                        <option value="Начальный">Начальный</option>
                        <option value="Базовый">Базовый</option>
                        <option value="Продвинутый">Продвинутый</option>
                        <option value="Детский">Детский</option>
                        <option value="Базовый USD/EUR">Базовый USD/EUR</option>
                        <option value="Продвинутый USD/EUR">Продвинутый USD/EUR</option>
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

export default Deposits