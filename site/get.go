package site

import "context"

type Site struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

//encore:api public method=GET path=/site/:siteID
func (s *Service) Get(ctx context.Context, siteID int) (*Site, error) {
	var site Site
	if err := s.db.Where("Id = $1", siteID).First(&site).Error; err != nil {
		return nil, err
	}

	return &site, nil
}
