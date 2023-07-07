import React from 'react';
import cl from './css/unreg.module.css';
import photoForStocks from "./img/stocks.jpg"
import photoForExchange from "./img/exchange.jpg"
import photoForBankIcon from "./img/bank.png"
import {NavLink} from "react-router-dom";


class Unreg extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <h1 className={cl.head}>GOsling</h1>
                    <NavLink to="/registration"><button className={cl.reg}>Регистрация</button></NavLink>
                    <NavLink to="/authorization"><button className={cl.ent}>Войти</button></NavLink>
                </div>

                <div className={cl.but}>
                <NavLink to="/stocks"><div className={cl.stocks} >
                        <div><img src={photoForStocks} alt="Акции" className={cl.photoForStocks} /></div>
                        <p style={{ textAlign: "center", fontSize: 20, letterSpacing: 4 }}>Акции</p>
                    </div></NavLink>

                    <NavLink to="/exchanges"><div className={cl.exchange}>
                        <div><img src={photoForExchange} alt="Курс" className={cl.photoForExchange} /></div>
                        <p style={{ textAlign: "center", fontSize: 20, letterSpacing: 4, wordSpacing: 8 }}>Курсы валют</p>
                    </div></NavLink>
                </div>

                <div className={cl.about}>
                    <div style={{ position: "absolute", margin: 15 }}><img src={photoForBankIcon} className={cl.bankIcon} /></div>
                    <p className={cl.fulling}>“GOsling” - веб-сайт, выступающий в роли “вымышленного” банка,
                        предназначенный для совершения различных финансовых операций.
                        Управления своими сбережениями, вкладами, акциями.
                        Веб-сайт не рассчитан на эксплуатацию в реальных условиях, а только выступает в качестве проекта для учебной практики, т.к. банк является вымышленным и не подключён к банковской сети, все совпадения случайны.
                    </p>
                </div>

                <div className={cl.help}>
                    <p className={cl.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default Unreg