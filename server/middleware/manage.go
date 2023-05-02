package middleware

import "github.com/GOsling-Inc/GOsling/services"

type ManagerMiddleware struct {
	service *services.Service
}

func NewManagerMiddleware(s *services.Service) *ManagerMiddleware {
	return &ManagerMiddleware{
		service: s,
	}
}

func (m *ManagerMiddleware) Confirm(id, table, state string) (int, error) {
	switch table {
	case "loans":
		err := m.service.ConfirmLoan(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	case "deposits":
		err := m.service.ConfirmDeposit(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	case "insurances":
		err := m.service.ConfirmInsurance(id, state)
		if err != nil {
			return INTERNAL, err
		}
		return OK, nil
	}
	return 0, nil
}
