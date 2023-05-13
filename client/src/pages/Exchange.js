import React from 'react';
import cs from '../css/exchange.module.css';
import { NavLink } from "react-router-dom";

class Exchange extends React.Component {

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/user"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <div className={cs.together}>
                    <div className={cs.form}>
                        <p className={cs.author}>Валюта</p>
                        <hr />
                        <p className={cs.selectType} >Выберите тип 1 валюты</p>
                        <select className={cs.valuta}>
                            <option value="BYN">BYN</option>
                            <option value="USD">USD</option>
                            <option value="EUR">EUR</option>
                        </select>
                        <p className={cs.selectType} >Выберите тип 2 валюты</p>
                        <select className={cs.valuta}>
                            <option value="BYN">BYN</option>
                            <option value="USD">USD</option>
                            <option value="EUR">EUR</option>
                        </select>
                        <input placeholder="Сумма" className={cs.name}></input>
                        <input placeholder="Номер счёта 1 валюты" className={cs.name}></input>
                        <input placeholder="Номер счёта 2 валюты " className={cs.name}></input>

                        <button className={cs.open}>Выполнить</button>
                    </div>

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
                            </div>
                            <div className={cs.sell}>
                                <div className={cs.exch}>
                                    <p >Продажа</p>
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