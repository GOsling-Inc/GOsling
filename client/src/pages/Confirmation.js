import React from 'react';
import cl from '../css/confirmation.module.css';
import { NavLink } from "react-router-dom";


class Confirmation extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>

                <div>
                    <div className={cl.loan}>
                        <p >Кредиты</p>
                        <hr />


                    </div>

                    <div className={cl.insurance}>
                        <p >Страхование</p>
                        <hr />


                    </div>

                    <div className={cl.account}>
                        <p >Счета</p>
                        <hr />


                    </div>
                </div>

            </div>
        );
    }
}

export default Confirmation