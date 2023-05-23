import React from 'react';
import cl from '../css/exampleLoan.module.css';
import Cookies from 'universal-cookie';


class AllInsurances extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            insurances: []
        };
        const cookies = new Cookies();
        fetch("http://localhost:1337/user/insurances", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({ insurances: data.data }))

    }



    render() {
        if (this.state.insurances != null)
            return (
                <div>
                    {this.state.insurances.map((el) => (
                        <div className={cl.exampleLoan} key={el.id}>

                            <div className={cl.divName}>
                                <p>Страховка на номер счёта: {el.accountId}</p>
                            </div>

                            <div className={cl.blocks}>

                                <div className={cl.amount}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Сумма</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>{el.amount}</p>
                                    </div>
                                </div>

                                <div className={cl.amount}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Статус</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>{el.state}</p>
                                    </div>
                                </div>

                                <div className={cl.percent}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Осталось</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>{el.remaining}</p>
                                    </div>
                                </div>

                                <div className={cl.percent}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Часть</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>{el.part}</p>
                                    </div>
                                </div>

                                <div className={cl.percent}>
                                    <div className={cl.textAboutBack}>
                                        <p className={cl.textAbout}>Срок</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>{el.deadline}</p>
                                    </div>
                                </div>
                            </div>


                        </div>))}
                </div>
            )
        else
            return (
                <div >
                    <p style={{ textAlign: "center", paddingTop: 50, fontSize: 22, letterSpacing: 1 }}>У вас нет активных страховок</p>
                </div>
            )
    }
}

export default AllInsurances