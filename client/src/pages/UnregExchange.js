import React from 'react';
import cs from '../css/unregExchange.module.css';
import { NavLink } from "react-router-dom";

class UnregExchange extends React.Component {

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/"><h1 className={cs.head}>GOsling</h1></NavLink>
                </div>

                <div className={cs.together}>
                    <div className={cs.rate}>
                        <p className={cs.pRate}>Курсы</p>
                        <div>
                            <div className={cs.val}>
                                <div className={cs.exch}>
                                    <p className={cs.pValuta}>Валюта</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p >USD</p>
                                </div>
                                <div className={cs.euro}>
                                    <p >EUR</p>
                                </div>
                            </div>
                            <div className={cs.buy}>
                                <div className={cs.exch}>
                                    <p >Покупка</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p style={{fontSize: 22}}>?</p>
                                </div>
                                <div className={cs.euro}>
                                    <p style={{fontSize: 22}}>?</p>
                                </div>

                            </div>
                            <div className={cs.sell}>
                                <div className={cs.exch}>
                                    <p >Продажа</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p style={{fontSize: 22}}>?</p>
                                </div>
                                <div className={cs.euro}>
                                    <p style={{fontSize: 22}}>?</p>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>

                <div className={cs.help}>
                    <p className={cs.info}>© 2023. GOsling</p>
                </div>
            </div>
        );
    }
}

export default UnregExchange