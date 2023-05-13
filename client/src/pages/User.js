import React from 'react';
import cl from '../css/user.module.css';
import { NavLink } from "react-router-dom";


class User extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <h1 className={cl.head}>GOsling</h1>
                </div>
                <div className={cl.action}>
                    <NavLink to="/user/stocks"><button>Инвестирование</button></NavLink>
                    <NavLink to="/user/new-account"><button>Открытие счёта</button></NavLink>
                    <NavLink to="/user/transfer"><button>Перевод со счёта на счёт</button></NavLink>
                    <NavLink to="/user/deposits"><button>Вклады</button></NavLink>
                    <NavLink to="/user/exchange"><button>Покупка/продажа валюты</button></NavLink>
                    <NavLink to="/user/loan"><button>Кредитование</button></NavLink>
                    <NavLink to="/user/insurance"><button>Страхование</button></NavLink>
                </div>

                <div className={cl.event}>
                    <div className={cl.blocks}>
                        <div className={cl.actAccount}>
                            <p className={cl.actAccountTitle}>Активные счета</p>
                            <hr />

                                <div className={cl.example}>
                                    <div className={cl.divName}>
                                        <p>Example</p>
                                    </div>

                                    <div className={cl.balance}>
                                        <p className={cl.remainderName} >Остаток на счёте:</p>
                                        <p className={cl.remainder} >1233.08 BYN</p>
                                        <hr />
                                    </div>

                                    <div className={cl.close} >
                                        <button className={cl.close1}>Закрыть</button>
                                        <button className={cl.close2}>Заморозить</button>
                                        <p >Вид счёта: базовый</p>
                                    </div>

                                </div>

                        </div>
                        <div className={cl.actLoan}>
                            <p>Активные кредиты</p>
                            <hr />
                        </div>
                    </div>
                    <div className={cl.history}>
                        <p>Последние переводы</p>
                        <hr/>
                    </div>
                </div>

                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default User