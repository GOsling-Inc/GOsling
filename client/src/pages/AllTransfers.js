import React from 'react';
import cl from '../css/user.module.css';
import Cookies from 'universal-cookie';
import photoArrow from "../img/arrow.png"


class AllTransfers extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            transfers: []
        };
        const cookies = new Cookies();
        fetch("http://localhost:1337/user/transfers", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({ transfers: data.data }))

    }



    render() {
        if (this.state.transfers != null)
            return (
                <div>
                    {this.state.transfers.map((el) => (
                        <div className={cl.exampleHistory} key={el.id}>
                            <p className={cl.aboutTransfer}>{el.sender} <img src={photoArrow} className={cl.arrow} /> {el.receiver} <b>:</b> {el.amount}</p>
                        </div>))}
                </div>
            )
        else
            return (
                <div >
                    <p style={{ textAlign: "center", paddingTop: 30, fontSize: 20, letterSpacing: 1 }}>У вас не было переводов</p>
                </div>
            )
    }
}

export default AllTransfers