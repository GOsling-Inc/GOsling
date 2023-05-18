import React from 'react';
import cl from '../css/roles.module.css';
import { NavLink } from "react-router-dom";


class Roles extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>

                <div>
                    <div className={cl.form}>
                        <input placeholder="Номер счёта" className={cl.input}></input>
                        <button className={cl.button}>Информация</button>
                        <button className={cl.button}>Сделать менеджером</button>
                    </div>
                    <div className={cl.information}>



                    </div>
                </div>
            </div>
        );
    }
}

export default Roles