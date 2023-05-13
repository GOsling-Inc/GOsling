import React from 'react';
import cl from '../css/authoriz.module.css';
import { NavLink } from "react-router-dom";


class Authorization extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/"><h1 className={cl.head}>GOsling</h1></NavLink>
                    <NavLink to="/registration"><button className={cl.reg}>Регистрация</button></NavLink>
                </div>
                <form className={cl.form}>
                    <p className={cl.author}>Авторизация</p>
                    <hr />
                    <input type="email" placeholder="Почта" ></input>
                    <input type="text" placeholder="Пароль"></input>
                    <NavLink to="/user"><button>Войти</button></NavLink>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Authorization