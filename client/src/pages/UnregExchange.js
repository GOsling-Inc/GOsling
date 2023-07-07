import React from 'react';
import cs from '../css/unregExchange.module.css';
import { NavLink } from "react-router-dom";

class UnregExchange extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            BU: 0,
            BE: 0,
        };

        fetch("http://localhost:1337/exchanges", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
            },
        }).then(res => res.json()).then(data => {
            this.setState({BU: data["BYN/USD"]})
            this.setState({BE: data["BYN/EUR"]})
        })

    }

    render() {
        return (
            <div>
                <div className={cs.headBack}>
                    <NavLink to="/"><h1 className={cs.head}>GOsling</h1></NavLink>
                    <NavLink to="/registration"><button className={cs.reg}>Регистрация</button></NavLink>
                    <NavLink to="/authorization"><button className={cs.ent}>Войти</button></NavLink>
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
                                    <p >BYN</p>
                                </div>

                                <div className={cs.dollar}>
                                    <p style={{fontSize: 22}}>{this.state.BU}</p>
                                </div>
                                <div className={cs.euro}>
                                    <p style={{fontSize: 22}}>{this.state.BE}</p>
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