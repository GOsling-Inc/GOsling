import React from "react"


class Authorization extends React.Component {

      render() {
        return (<div>
            <h1>GOsling</h1>
            <form className="form">
        <input placeholder="Логин"></input>
        <input placeholder="Пароль"></input>
        <button>Войти</button>
        </form>
    </div>)
  }
}

export default Authorization