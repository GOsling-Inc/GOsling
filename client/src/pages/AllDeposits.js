import React from 'react';
import cl from '../css/exampleDeposits.module.css';
import Cookies from 'universal-cookie';


class AllDeposits extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            deposits: []
        };
        const cookies = new Cookies();
        fetch("http://localhost:1337/user/deposits", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({ deposits: data.data }))

    }



    render() {
        if (this.state.deposits != null)
            return (
                <div>
                    {this.state.deposits.map((el) => (
                        <div className={cl.exampleLoan} key={el.id}>

                            <div className={cl.divName}>
                                <p>Вклад на номер счёта:</p>
                            </div>

                            <div className={cl.blocks}>
                                <div className={cl.percent}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Процент</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>1%</p>
                                    </div>
                                </div>

                                <div className={cl.amount}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Сумма</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>321 BYN</p>
                                    </div>
                                </div>

                                <div className={cl.time}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Срок</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>1 год</p>
                                    </div>
                                </div>
                            </div>
                            

                        </div>))}
                </div>
            )
        else
            return (
                <div >
                    <p style={{ textAlign: "center", paddingTop: 50, fontSize: 22, letterSpacing: 1 }}>У вас нет активных вкладов</p>
                </div>
            )
    }
}

export default AllDeposits