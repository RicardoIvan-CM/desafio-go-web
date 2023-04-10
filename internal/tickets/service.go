package tickets

import "context"

type Service interface {
	GetTotalTickets(context.Context, string) (int, error)
	AverageDestination(context.Context, string) (float64, error)
}

type DefaultService struct {
	repository Repository
}

func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		repository: repo,
	}
}

func (s *DefaultService) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *DefaultService) AverageDestination(ctx context.Context, destination string) (float64, error) {
	var ticketsDestino int
	var ticketsTotales int
	var result float64
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	ticketsDestino = len(tickets)
	tickets, err = s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	ticketsTotales = len(tickets)
	result = float64(ticketsDestino) / float64(ticketsTotales)
	return result, nil
}
