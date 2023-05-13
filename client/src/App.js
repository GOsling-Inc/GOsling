import React from "react";
import Unreg from "./Unreg";
import {BrowserRouter, Routes, Route, Link} from "react-router-dom";
import Registration from "./pages/Registration";
import Authorization from "./pages/Authorization";
import Support from "./pages/Support";
import User from "./pages/User";
import Transfer from "./pages/Transfer";
import Deposits from "./pages/Deposits";
import Exchange from "./pages/Exchange";
import Insurance from "./pages/Insurance";
import Loan from "./pages/Loan";
import NewAccount from "./pages/NewAccount";
import Stocks from "./pages/Stocks";

function App() {
    return(
        <div>
            
            <BrowserRouter>
                <Routes>
                    <Route exact path="/" Component={Unreg} />
                    <Route exact path="/authorization" Component={Authorization} />
                    <Route exact path="/registration" Component={Registration} />
                    <Route exact path="/support" Component={Support} />
                    <Route exact path="/user" Component={User} />
                    <Route exact path="/user/transfer" Component={Transfer} />
                    <Route exact path="/user/deposits" Component={Deposits} />
                    <Route exact path="/user/exchange" Component={Exchange} />
                    <Route exact path="/user/insurance" Component={Insurance} />
                    <Route exact path="/user/loan" Component={Loan} />
                    <Route exact path="/user/new-account" Component={NewAccount} />
                    <Route exact path="/user/stocks" Component={Stocks} />
                </Routes>
            </BrowserRouter>

        </div>
    )
}

export default App;