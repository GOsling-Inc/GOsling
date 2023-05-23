import React from 'react';
import cl from '../css/exampleInsurances.module.css';
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
        }).then(res => res.json()).then(data => this.setState({ insurances: data.data}))

    }



    render() {
        if (this.state.insurances != null)
            return (
                <div>
                    {this.state.insurances.map((el) => (
                        <div className={cl.exampleLoan} key={el.id}>

                            <div className={cl.divName}>
                                <p>Страховка на номер счёта:</p>
                            </div>

                            <div className={cl.blocks}>


                                <div className={cl.amount}>
                                    <div className={cl.textAboutBack1}>
                                        <p className={cl.textAbout}>Сумма</p>
                                    </div>
                                    <div className={cl.aboutAm}>
                                        <p className={cl.textIn}>321 BYN</p>
                                    </div>
                                </div>

                                <div className={cl.time}>
                                    <div className={cl.textAboutBack2}>
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
                    <p style={{ textAlign: "center", paddingTop: 50, fontSize: 22, letterSpacing: 1 }}>У вас нет активных страховок</p>
                </div>
            )
    }
}

export default AllInsurances