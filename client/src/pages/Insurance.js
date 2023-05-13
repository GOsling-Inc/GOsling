import React from 'react';
import cl from '../css/insurance.module.css';
import { NavLink } from "react-router-dom";


class Insurance extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/user"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                <div className={cl.form}>
                    <p className={cl.author}>Страхование</p>
                    <hr />
                    <input placeholder="Наименование имущества" className={cl.name}></input>
                    <input placeholder="Сумма" className={cl.name}></input>

                    <button className={cl.open}>Застраховать</button>
                </div>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Insurance