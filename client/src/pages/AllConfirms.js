import React from 'react';
import cl from '../css/confirmation.module.css';
import Cookies from 'universal-cookie';


class AllConfirms extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            confirms: []
        };
        const cookies = new Cookies();
        fetch("http://localhost:1337/manage/confirms", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({ confirms: data.data }))

    }

    render() {
        if (this.state.confirms != null)
            return (
                <div>
                    {this.state.confirms.map((el) => (
                        <div className={cl.exampleConfirm}  key={el.id}>
                        <div className={cl.table}>
                            <p className={cl.tableText}>Table: {el.table}</p>
                        </div>
                        <div className={cl.id}>
                            <p className={cl.idText}>Id: {el.id}</p>
                        </div>
                    </div>))}
                </div>
            )

    }
}

export default AllConfirms