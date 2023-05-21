import React from 'react';
import cl from '../css/transactions.module.css';
import { NavLink } from "react-router-dom";


class Transactions extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                
                <div className={cl.form}>
                    <input placeholder="Номер транзакции" className={cl.input}></input>
                    <button className={cl.button}>Отменить</button>
                </div>
            </div>
        );
    }
}

export default Transactions