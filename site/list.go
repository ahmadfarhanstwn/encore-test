package site

import "context"

type GetAllResponse struct {
	Sites []*Site `json:"sites"`
}

//encore:api public method=GET path=/site
func (s *Service) List(ctx context.Context) (*GetAllResponse, error) {
	var sites []*Site
	if err := s.db.Find(&sites).Error; err != nil {
		return nil, err
	}
	return &GetAllResponse{Sites: sites}, nil
}
