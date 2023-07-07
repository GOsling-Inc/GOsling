import React from 'react';
import cl from '../css/support.module.css';
import { NavLink } from "react-router-dom";


class Support extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/user"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>

                <div className={cl.form}>
                    <p >Опишите вашу проблему</p>
                    <hr />
                    <input type="text" placeholder="Заголовок" className={cl.name} maxlength="70"></input>
                    <textarea placeholder="Содержание" className={cl.description}></textarea>
                    <button>Отправить</button>
                </div>

                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Support