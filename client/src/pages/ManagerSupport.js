import React from 'react';
import cl from '../css/managerSupport.module.css';
import { NavLink } from "react-router-dom";


class ManagerSupport extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <NavLink to="/manage"><h1 className={cl.head}>GOsling</h1></NavLink>
                </div>

                <div className={cl.stock}>
                    <p >Сообщения</p>
                    <hr />

                    <div className={cl.example}>

                    </div>
                </div>

            </div>
        );
    }
}

export default ManagerSupport