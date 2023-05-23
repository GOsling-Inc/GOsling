import React from 'react';
import cl from '../css/roles.module.css';
import Cookies from 'universal-cookie';


class AllUsers extends React.Component {
    constructor(props) {
        super(props);
        
        this.state = {
            users: []
        };
        const cookies = new Cookies();
        fetch("http://localhost:1337/manage/users", {
            method: "GET",
            headers: {
                'Accept': 'application/json',
                'Content-type': 'application/json',
                "Token": cookies.get('Token')
            },
        }).then(res => res.json()).then(data => this.setState({ users: data.data }))

    }

    render() {
        if (this.state.users != null)
            return (
                <div>
                    {this.state.users.map((el) => (
                         <div className={cl.exampleProfile}  key={el.id}>
                         <div className={cl.idRole}>
                             <div className={cl.left}>
                                 <p >id:</p>
                                 <p>{el.id}</p>
                             </div>
                             <div className={cl.right}>
                                 <p>role:</p>
                                 <p>{el.role}</p>
                             </div>
                         </div>
                         <div className={cl.name}>
                             <div className={cl.left}>
                                 <p>Name:</p>
                                 <p>{el.name}</p>
                             </div>
                             <div className={cl.right}>
                                 <p>Surname:</p>
                                 <p>{el.surname}</p>
                             </div>
                         </div>
                         <div className={cl.emailBirth}>
                             <div className={cl.left}>
                                 <p>email:</p>
                                 <p>{el.email}</p>
                             </div>
                             <div className={cl.right}>
                                 <p>Birthdate:</p>
                                 <p>{el.birthdate}</p>
                             </div>
                         </div>
                     </div>))}
                </div>
            )

    }
}

export default AllUsers