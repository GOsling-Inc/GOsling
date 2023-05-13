import React from 'react';
import cs from '../css/new-account.module.css';
import { NavLink } from "react-router-dom";
import photoQuestion from "../img/question.png"

class NewAccount extends React.Component {

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <div className={cs.form}>
                    <p className={cs.author}>Открытие счёта</p>
                    <hr />
                    <input type="text" placeholder="Название счёта" className={cs.name}></input>
                    <p className={cs.pType}>Выберите тип валюты</p>
                    <select className={cs.valuta}>
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

                    <select className={cs.acc}>
                        <option value="Базовый счёт">Базовый счёт</option>
                        <option value="Бизнес счёт">Бизнес счёт</option>
                        <option value="Инвестиционный счёт">Инвестиционный счёт</option>
                    </select>
                    <button className={cs.open}>Открыть</button>
                </div>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                    <NavLink className={cs.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default NewAccount