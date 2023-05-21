import React from 'react';
import cl from '../css/unregStocks.module.css';
import { NavLink } from "react-router-dom";


class UnregStocks extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/"><h1 className={cl.head}>GOsling</h1></NavLink>
                    <NavLink to="/registration"><button className={cl.reg}>Регистрация</button></NavLink>
                    <NavLink to="/authorization"><button className={cl.ent}>Войти</button></NavLink>
                </div>
                <div className={cl.together}>
                    <div className={cl.stock}>
                        <p >Акции</p>
                        <hr />

                    </div>
                </div>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default UnregStocks