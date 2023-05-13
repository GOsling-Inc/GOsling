import React from 'react';
import cl from '../css/registration.module.css';
import { NavLink } from "react-router-dom";


class Registration extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/"><h1 className={cl.head}>GOsling</h1></NavLink>
                    <NavLink to="/authorization"><button className={cl.ent}>Войти</button></NavLink>
                </div>
                <form className={cl.form}>
                    <p className={cl.author}>Регистрация</p>
                    <hr />
                    <input type="text" placeholder="Имя" ></input>
                    <input type="text" placeholder="Фамилия"></input>
                    <input type="date" className={cl.date}></input>
                    <input type="email" placeholder="Почта"></input>
                    <input type="text" placeholder="Пароль"></input>
                    <input type="text" placeholder="Подтверждение пароля"></input>
                    <button>Зарегистрироваться</button>
                </form>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Registration