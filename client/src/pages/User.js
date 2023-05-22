import React from 'react';
import cl from '../css/user.module.css';
import { NavLink } from "react-router-dom";
import photoArrow from "../img/arrow.png"
import AllAccounts from './AllAccounts';

class User extends React.Component {

    render() {
        return (
            <div>
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
                            <AllAccounts/>
                        </div>

                        <div className={cl.actLoan}>
                            <p>Активные кредиты</p>
                            <hr />
                        </div>
                    </div>

                    <div className={cl.history}>
                        <p>Последние переводы</p>
                        <hr />

                        <div className={cl.exampleHistory}>
                            <p className={cl.aboutTransfer}>123121 <img src={photoArrow} className={cl.arrow} /> 123113 <b>:</b> 311 BYN</p>
                        </div>

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