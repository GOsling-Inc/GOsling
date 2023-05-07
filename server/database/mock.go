package database

import (
	"errors"

	"github.com/GOsling-Inc/GOsling/models"
)

var inv_id = 1

type MockDatabase struct {
	users            []models.User
	accounts         []models.Account
	transfers        []models.Trasfer
	exchanges        []models.Exchange
	loans            []models.Loan
	deposits         []models.Deposit
	insurances       []models.Insurance
	investments      []models.Investment
	orders           []models.Order
	order_operations []models.Order_operation
}

func NewMock() *MockDatabase {

	return &MockDatabase{
		users:            make([]models.User, 10),
		accounts:         make([]models.Account, 10),
		transfers:        make([]models.Trasfer, 10),
		exchanges:        make([]models.Exchange, 10),
		loans:            make([]models.Loan, 10),
		deposits:         make([]models.Deposit, 10),
		insurances:       make([]models.Insurance, 10),
		investments:      make([]models.Investment, 10),
		orders:           make([]models.Order, 10),
		order_operations: make([]models.Order_operation, 10),
	}
}

func (d *MockDatabase) GetUserByMail(mail string) (models.User, error) {
	for _, usr := range d.users {
		if usr.Email == mail {
			return usr, nil
		}
	}
	return models.User{}, errors.New("not found")
}

func (d *MockDatabase) GetUserById(id string) (models.User, error) {
	for _, usr := range d.users {
		if usr.Id == id {
			return usr, nil
		}
	}
	return models.User{}, errors.New("not found")
}

func (d *MockDatabase) AddUser(user *models.User) error {
	d.users = append(d.users, *user)
	return nil
}

func (d *MockDatabase) UpdatePasswordUser(id, password string) error {
	for i, usr := range d.users {
		if usr.Id == id {
			d.users[i].Password = password
			return nil
		}
	}
	return errors.New("not found")
}

func (d *MockDatabase) UpdateUserData(id, name, surname, birthdate string) error {
	for i, usr := range d.users {
		if usr.Id == id {
			d.users[i].Name = name
			d.users[i].Surname = surname
			d.users[i].Birthdate = birthdate
			return nil
		}
	}
	return errors.New("not found")
}

func (d *MockDatabase) AddAccount(account models.Account) error {
	d.accounts = append(d.accounts, account)
	return nil
}

func (d *MockDatabase) GetUserAccounts(userId string) ([]models.Account, error) {
	var accounts []models.Account
	for _, acc := range d.accounts {
		if acc.UserId == userId {
			accounts = append(accounts, acc)
		}
	}
	return accounts, nil
}

func (d *MockDatabase) GetAccountById(id string) (models.Account, error) {
	for _, acc := range d.accounts {
		if acc.Id == id {
			return acc, nil
		}
	}
	return models.Account{}, errors.New("not found")
}

func (d *MockDatabase) DeleteAccount(id string) error {
	check := false
	for i, acc := range d.accounts {
		if acc.Id == id {
			check = true
			d.accounts[i] = d.accounts[len(d.accounts)-1]
		}
	}
	if !check {
		return errors.New("error")
	}
	d.accounts = d.accounts[:len(d.accounts)-1]
	return nil
}

func (d *MockDatabase) Transfer(senderId, receiverId string, amount float64) error {
	for i, acc := range d.accounts {
		if acc.Id == senderId {
			d.accounts[i].Amount -= amount
		}
		if acc.Id == receiverId {
			d.accounts[i].Amount += amount
		}
	}
	return nil
}

func (d *MockDatabase) AddTransfer(transfer models.Trasfer) error {
	d.transfers = append(d.transfers, transfer)
	return nil
}

func (d *MockDatabase) Exchange(senderId, receiverId string, sender_amount, receiver_amount float64) error {
	for i, acc := range d.accounts {
		if acc.Id == senderId {
			d.accounts[i].Amount -= sender_amount
		}
		if acc.Id == receiverId {
			d.accounts[i].Amount += receiver_amount
		}
	}
	return nil
}

func (d *MockDatabase) AddExchange(exchange models.Exchange) error {
	d.exchanges = append(d.exchanges, exchange)
	return nil
}

func (d *MockDatabase) AddDeposit(deposit models.Deposit) error {
	d.deposits = append(d.deposits, deposit)
	return nil
}

func (d *MockDatabase) GetUserDeposits(userId string) ([]models.Deposit, error) {
	var deposits []models.Deposit
	for _, dep := range d.deposits {
		if dep.UserId == userId {
			deposits = append(deposits, dep)
		}
	}
	return deposits, nil
}

func (d *MockDatabase) GetDepositById(id string) (models.Deposit, error) {
	for _, dep := range d.deposits {
		if dep.Id == id {
			return dep, nil
		}
	}
	return models.Deposit{}, errors.New("not found")
}

func (d *MockDatabase) AddInsurance(insurance models.Insurance) error {
	d.insurances = append(d.insurances, insurance)
	return nil
}

func (d *MockDatabase) GetUserInsurances(userId string) ([]models.Insurance, error) {
	var insurances []models.Insurance
	for _, ins := range d.insurances {
		if ins.UserId == userId {
			insurances = append(insurances, ins)
		}
	}
	return insurances, nil
}

func (d *MockDatabase) GetInsuranceById(id string) (models.Insurance, error) {
	for _, ins := range d.insurances {
		if ins.Id == id {
			return ins, nil
		}
	}
	return models.Insurance{}, errors.New("not found")
}

func (d *MockDatabase) AddInvestment(investment models.Investment) error {
	d.investments = append(d.investments, investment)
	return nil
}

func (d *MockDatabase) GetInvestments() ([]models.Investment, error) {
	return d.investments, nil
}

func (d *MockDatabase) GetInvestment(name string) (models.Investment, error) {
	for _, inv := range d.investments {
		if inv.Name == name {
			return inv, nil
		}
	}
	return models.Investment{}, errors.New("not found")
}

func (d *MockDatabase) CreateOrder(order models.Order) error {
	d.orders = append(d.orders, order)
	return nil
}

func (d *MockDatabase) GetOrders() ([]models.Order, error) {
	return d.orders, nil
}

func (d *MockDatabase) GetOrder(id int) (models.Order, error) {
	for _, ord := range d.orders {
		if ord.Id == id {
			return ord, nil
		}
	}
	return models.Order{}, errors.New("not found")
}

func (d *MockDatabase) Buy(accountId string, order models.Order, count int) error {
	for i, acc := range d.accounts {
		if acc.Id == accountId {
			d.accounts[i].Amount -= order.Price * float64(count)
		}
		if acc.Id == order.AccountId {
			d.accounts[i].Amount += order.Price * float64(count)
		}
	}
	for i, ord := range d.orders {
		if ord.Id == order.Id {
			d.orders[i].Count -= count
			if d.orders[i].Count == 0 {
				d.orders[i] = d.orders[len(d.orders)-1]
				d.orders = d.orders[:len(d.orders)-1]
			}
		}
	}
	for i, inv := range d.investments {
		if inv.Name == order.Name {
			d.investments[i].Investors[order.UserId] -= count
			d.investments[i].Investors[accountId] += count
		}
	}
	d.order_operations = append(d.order_operations, models.Order_operation{
		Id:     inv_id,
		Name:   order.Name,
		Buyer:  accountId,
		Seller: order.AccountId,
		Count:  count,
		Price:  order.Price,
	})
	inv_id++

	return nil
}

func (d *MockDatabase) Sell(accountId string, order models.Order, count int) error {
	for i, acc := range d.accounts {
		if acc.Id == accountId {
			d.accounts[i].Amount += order.Price * float64(count)
		}
		if acc.Id == order.AccountId {
			d.accounts[i].Amount -= order.Price * float64(count)
		}
	}
	for i, ord := range d.orders {
		if ord.Id == order.Id {
			d.orders[i].Count -= count
			if d.orders[i].Count == 0 {
				d.orders[i] = d.orders[len(d.orders)-1]
				d.orders = d.orders[:len(d.orders)-1]
			}
		}
	}
	for i, inv := range d.investments {
		if inv.Name == order.Name {
			d.investments[i].Investors[order.UserId] += count
			d.investments[i].Investors[accountId] -= count
		}
	}
	d.order_operations = append(d.order_operations, models.Order_operation{
		Id:     inv_id,
		Name:   order.Name,
		Buyer:  order.AccountId,
		Seller: accountId,
		Count:  count,
		Price:  order.Price,
	})
	inv_id++

	return nil
}

func (d *MockDatabase) AddLoan(loan models.Loan) error {
	d.loans = append(d.loans, loan)
	return nil
}

func (d *MockDatabase) GetLoanById(id string) (models.Loan, error) {
	for _, ln := range d.loans {
		if ln.Id == id {
			return ln, nil
		}
	}
	return models.Loan{}, errors.New("not found")
}

func (d *MockDatabase) GetUserLoans(userId string) ([]models.Loan, error) {
	var loans []models.Loan
	for _, ln := range d.loans {
		if ln.UserId == userId {
			loans = append(loans, ln)
		}
	}
	return loans, nil
}

func (d *MockDatabase) GetUnconfirmed() []models.Unconfirmed {
	var unc []models.Unconfirmed
	for _, ln := range d.loans {
		if ln.State == "PENDING" {
			unc = append(unc, models.Unconfirmed{
				Table: "loans",
				Id:    ln.Id,
			})
		}
	}
	for _, d := range d.deposits {
		if d.State == "PENDING" {
			unc = append(unc, models.Unconfirmed{
				Table: "deposits",
				Id:    d.Id,
			})
		}
	}
	for _, ins := range d.insurances {
		if ins.State == "PENDING" {
			unc = append(unc, models.Unconfirmed{
				Table: "insurances",
				Id:    ins.Id,
			})
		}
	}
	return unc

}

func (d *MockDatabase) ConfirmLoan(loan models.Loan) error {
	for i, ln := range d.loans {
		if ln.Id == loan.Id {
			d.loans[i] = loan
		}
	}
	for i, acc := range d.accounts {
		if acc.Id == loan.AccountId {
			d.accounts[i].Amount += loan.Amount
		}
	}

	return nil
}

func (d *MockDatabase) ConfirmInsurance(insurance models.Insurance) error {
	for i, ins := range d.insurances {
		if ins.Id == insurance.Id {
			d.insurances[i].State = insurance.State
		}
	}
	return nil
}

func (d *MockDatabase) ConfirmDeposit(deposit models.Deposit) error {
	for i, dp := range d.deposits {
		if dp.Id == deposit.Id {
			d.deposits[i] = deposit
		}
	}
	for i, acc := range d.accounts {
		if acc.Id == deposit.AccountId {
			d.accounts[i].Amount -= deposit.Amount
		}
	}

	return nil
}

func (d *MockDatabase) GetTransactions() []models.Trasfer {
	return d.transfers
}

func (d *MockDatabase) CancelTransaction(trasaction models.Trasfer) error {
	return d.Transfer(trasaction.Receiver, trasaction.Sender, trasaction.Amount)
}

func (d *MockDatabase) GetAccounts() []models.Account {
	return d.accounts
}

func (d *MockDatabase) GetTransferById(id string) (models.Trasfer, error) {
	for _, trs := range d.transfers {
		if trs.Id == id {
			return trs, nil
		}
	}
	return models.Trasfer{}, errors.New("not found")
}

func (d *MockDatabase) UpdateAccount(id, state string) error {
	for i, acc := range d.accounts {
		if acc.Id == id {
			d.accounts[i].State = state
			return nil
		}
	}
	return errors.New("not found")
}

func (d *MockDatabase) GetUsers() []models.User {
	return d.users
}

func (d *MockDatabase) UpdateRole(id, role string) error {
	for i, user := range d.users {
		if user.Id == id {
			d.users[i].Role = role
			return nil
		}
	}
	return errors.New("not found")
}

func (d *MockDatabase) UpdateLoans() error {
	return nil
}

func (d *MockDatabase) UpdateInsurances() error {
	return nil
}

func (d *MockDatabase) UpdateDeposits() error {
	return nil
}
