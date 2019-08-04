package service

import (
	"context"
	data "week3/data"
	gw "week3/transport"
)

type feedBackService struct {
	db *data.FeedbackDB
}

// AddFeedback Add new feedback
func (s *feedBackService) AddFeedback(ctx context.Context, in *gw.AddPassengerFeedbackRequest) (*gw.AddPassengerFeedbackResponse, error) {
	feedbacks := s.db.Query(0, in.GetFeedback().GetBookingCode())
	if len(feedbacks) > 0 {
		return &gw.AddPassengerFeedbackResponse{Errors: []*gw.Error{&gw.Error{
			Code:    1,
			Message: "Feeback for this booking has been made",
		}}}, nil
	}

	// Map to DB model
	s.db.Add(&data.Feedback{
		FeedbackID:  in.GetFeedback().GetFeedbackID(),
		BookingCode: in.GetFeedback().GetBookingCode(),
		PassengerID: in.GetFeedback().GetPassengerID(),
		Feedback:    in.GetFeedback().GetFeedback(),
	})
	return &gw.AddPassengerFeedbackResponse{Errors: []*gw.Error{}}, nil
}

// ListFeedBack List all feedbacks belong to passengers
func (s *feedBackService) ListFeedBack(ctx context.Context, in *gw.ListPassengerFeedbackRequest) (*gw.ListPassengerFeedbackResponse, error) {
	total := s.db.Count()
	feedbacks := s.db.Query(in.GetPassengerID(), in.GetBookingCode())
	ret := []*gw.PassengerFeedback{}

	// Mapping db model to gRPC model
	for _, feedback := range feedbacks {
		ret = append(ret, &gw.PassengerFeedback{
			FeedbackID:  feedback.FeedbackID,
			BookingCode: feedback.BookingCode,
			PassengerID: feedback.PassengerID,
			Feedback:    feedback.Feedback,
		})
	}
	return &gw.ListPassengerFeedbackResponse{Errors: nil, Paging: &gw.Paging{Total: total, Offset: 0, Limit: total}, Data: ret}, nil
}

// DeleteFeedback Delete feedback by passenger code
func (s *feedBackService) DeleteFeedback(ctx context.Context, in *gw.DeleteFeedBackRequest) (*gw.DeleteFeedBackResponse, error) {
	feedbacks := s.db.Query(in.GetPassengerID(), "")

	ret := []string{}
	for _, feedback := range feedbacks {
		s.db.DeleteByID(feedback.FeedbackID)
		ret = append(ret, feedback.FeedbackID)
	}
	return &gw.DeleteFeedBackResponse{FeedbackIds: ret}, nil
}

// NewFeedbackService Create new FeedbackService instance
func NewFeedbackService(db *data.FeedbackDB) gw.FeedbackServiceServer {
	return &feedBackService{db: db}
}
