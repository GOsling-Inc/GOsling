import React from 'react';
import cl from '../css/stocks.module.css';
import { NavLink } from "react-router-dom";


class Stocks extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/user"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>
                <div className={cl.together}>
                    <div className={cl.form}>
                        <p style={{wordSpacing: 5}}>Ваши акции</p>
                        <hr />

                    </div>

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

export default Stocks