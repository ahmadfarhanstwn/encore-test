package site

import "context"

type AddParams struct {
	Url string `json:"url"`
}

//encore:api public method=POST path=/site
func (s *Service) Add(ctx context.Context, params *AddParams) (*Site, error) {
	site := &Site{Url: params.Url}
	if err := s.db.Create(site).Error; err != nil {
		return nil, err
	}
	return site, nil
}
