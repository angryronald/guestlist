package guest

import "context"

func (s *Service) GetAvailableSpace(ctx context.Context) (int, error) {
	allGuests, err := s.ListGuests(ctx, false)
	if err != nil {
		return 0, err
	}

	allArrivedGuests, err := s.ListGuests(ctx, true)
	if err != nil {
		return 0, err
	}

	var totalSpace int
	for _, guest := range allGuests {
		totalSpace = totalSpace + guest.AccompanyingGuests
	}

	var totalUsedSpace int
	for _, guest := range allArrivedGuests {
		totalUsedSpace = totalUsedSpace + guest.ActualAccompanyingGuests
	}

	return totalSpace - totalUsedSpace, nil
}
