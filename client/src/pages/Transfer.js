import React from 'react';
import cl from '../css/transfer.module.css';
import { NavLink } from "react-router-dom";


class Transfer extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                <NavLink to="/user"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                <div className={cl.form}>
                    <p className={cl.author}>Перевод средств</p>
                    <hr />
                    <input type="text" placeholder="Счет отправителя" className={cl.name} style={{marginTop: 40}}></input>
                    <input type="text" placeholder="Счет получателя" className={cl.name}></input>
                    <input type="text" placeholder="Сумма перевода" className={cl.name}></input>
                    <button className={cl.open}>Перевести</button>
                </div>
                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                    <NavLink className={cl.support} to="/support">Служба поддержки</NavLink>
                </div>
            </div>
        );
    }
}

export default Transfer