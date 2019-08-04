package service

import (
	"context"
	gw "week3/transport"
)

type feedBackService struct {
	// db *sql.DB
}

// AddFeedback
func (s *feedBackService) AddFeedback(ctx context.Context, in *gw.AddPassengerFeedbackRequest) (*gw.AddPassengerFeedbackResponse, error) {
	return &gw.AddPassengerFeedbackResponse{Errors: []*gw.Error{}}, nil
}

// ListFeedBack
func (s *feedBackService) ListFeedBack(ctx context.Context, in *gw.ListPassengerFeedbackRequest) (*gw.ListPassengerFeedbackResponse, error) {
	return &gw.ListPassengerFeedbackResponse{Errors: nil, Paging: &gw.Paging{Total: 10, Offset: 0, Limit: 0}, Data: nil}, nil
}

// NewFeedbackService NewFeedbackService
func NewFeedbackService() gw.FeedbackServiceServer {
	// return &feedBackService{db: db}
	return &feedBackService{}
}
