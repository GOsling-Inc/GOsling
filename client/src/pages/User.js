import React from 'react';
import cl from '../css/user.module.css';
import { NavLink } from "react-router-dom";
import AllAccounts from './AllAccounts';
import AllLoans from './AllLoans';
import AllTransfers from './AllTransfers';
import AllDeposits from './AllDeposits';
import AllInsurances from './AllInsurances';


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
                            <AllAccounts />
                        </div>

                        <div className={cl.actLoan}>
                            <p className={cl.actLoanText}>Активные кредиты</p>
                            <hr />
                            <AllLoans />
                        </div>
                    </div>

                    <div className={cl.history}>
                        <p>Последние переводы</p>
                        <hr />
                        <AllTransfers />
                    </div>
                </div>

                <div className={cl.together}>
                    <div className={cl.actDeposit}>
                        <p className={cl.actDepositTitle}>Активные вклады</p>
                        <hr />
                        <AllDeposits />
                    </div>

                    <div className={cl.actInsurance}>
                        <p className={cl.actInsuranceText}>Активные страховки</p>
                        <hr />
                        <AllInsurances />
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