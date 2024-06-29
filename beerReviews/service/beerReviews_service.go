// breweryReviews/service/review_service.go
package service

import (
	"Beer.app/beerReviews/repository"
	"Beer.app/models"
)

type BeerReviewService interface {
	GetBeerReview(id int) (*models.Review, error)
	CreateBeerReview(review *models.Review) error
}

type beerReviewService struct {
	repo repository.BeerReviewRepository
}

func NewBeerReviewService(repo repository.BeerReviewRepository) BeerReviewService {
	return &beerReviewService{repo: repo}
}

func (s *beerReviewService) GetBeerReview(id int) (*models.Review, error) {
	return s.repo.GetBeerReviewByID(id)
}

func (s *beerReviewService) CreateBeerReview(review *models.Review) error {
	return s.repo.CreateBeerReview(review)
}
