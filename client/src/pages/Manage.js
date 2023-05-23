import React from 'react';
import cl from '../css/manage.module.css';
import { NavLink } from "react-router-dom";


class Manage extends React.Component {

    render() {
        return (
            <div>
                <div className={cl.headBack}>
                    <h1 className={cl.head}>GOsling</h1>
                </div>

                <div className={cl.together}>

                    <div className={cl.blocks1}>
                    <NavLink to="/manage/accounts"><div className={cl.accounts}>
                            <p className={cl.text}>Счета пользователей</p>
                        <hr/>
                        </div></NavLink>

                        <NavLink to="/manage/confirmation"><div className={cl.confirmations}>
                            <p className={cl.text}>Подтверждение</p>
                            <hr/>
                        </div></NavLink>
                    </div>

                    <div className={cl.blocks2}>
                    <NavLink to="/manage/transactions"><div className={cl.transactions}>
                            <p className={cl.text}>Отмена транзакций</p>
                            <hr/>
                        </div></NavLink>

                        <NavLink to="/manage/roles"><div className={cl.roles}>
                        <p className={cl.text}>Роли</p>
                        <hr/>
                    </div></NavLink>
                    </div>

                    

                </div>

            </div>
        );
    }
}

export default Manage