import React from 'react';
import cl from '../css/user.module.css';
import Cookies from 'universal-cookie';


class AllAccounts extends React.Component {
    constructor(props) {
        super(props);

        const cookies = new Cookies();
        fetch("http://localhost:1337/user/accounts", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({accounts: data.data}))

        this.state = {
            accounts: [ ]
        };

    }



    render() {
        if (this.state.accounts.length > 0)
            return (
                <div>
                    {this.state.accounts.map((el) => (
                        <div className={cl.exampleAccount} key={el.id}>
                            <div className={cl.divName}>
                                <p>{el.name}</p>
                            </div>

                            <div className={cl.balance}>
                                <div className={cl.aboutAccount}>
                                    <p className={cl.numberAccount} >Номер счёта: {el.id}</p>
                                    <p className={cl.remainderName} >Остаток на счёте:</p>
                                </div>
                                <p className={cl.remainder} >{el.amount} {el.unit}</p>
                                <hr />
                            </div>

                            <div className={cl.close} >
                                <button className={cl.close1}>Закрыть</button>
                                <button className={cl.close2}>Заморозить</button>
                                <p >Вид: {el.type}</p>
                            </div>

                        </div>))}
                </div>
            )
        else
            return (
                <div >
                    <p style={{ textAlign: "center", paddingTop: 50, fontSize: 22, letterSpacing: 1 }}>У вас нет активных счётов</p>
                </div>
            )
    }
}

export default AllAccounts