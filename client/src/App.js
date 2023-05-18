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
import UnregExchange from "./pages/UnregExchange";
import UnregStocks from "./pages/UnregStocks";
import Manage from "./pages/Manage";
import Confirmation from "./pages/Confirmation"
import Accounts from "./pages/Accounts"
import Transactions from "./pages/Transactions"
import Roles from "./pages/Roles"
import ManagerSupport from "./pages/ManagerSupport"

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
                    <Route exact path="/exchanges" Component={UnregExchange} />
                    <Route exact path="/stocks" Component={UnregStocks} />
                    <Route exact path="/manage" Component={Manage} />
                    <Route exact path="/manage/confirmation" Component={Confirmation} />
                    <Route exact path="/manage/accounts" Component={Accounts} />
                    <Route exact path="/manage/transactions" Component={Transactions} />
                    <Route exact path="/manage/roles" Component={Roles} />
                    <Route exact path="/manage/support" Component={ManagerSupport} />
                </Routes>
            </BrowserRouter>

        </div>
    )
}

export default App;