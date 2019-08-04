package data

// Repository instance
var Repository *FeedbackDB

func init() {
	Repository = &FeedbackDB{userDB: make(map[string]*Feedback)}
}

// Feedback Feedback
type Feedback struct {
	FeedbackID  string
	BookingCode string
	PassengerID int32
	Feedback    string
}

// FeedbackDB FeedbackDB struct
type FeedbackDB struct {
	userDB map[string]*Feedback
}

// Query Search feedback by passenger's ID or booking's code
func (i *FeedbackDB) Query(passengerID int32, bookingCode string) []*Feedback {
	ret := make([]*Feedback, 0)

	for _, v := range i.userDB {
		if passengerID > 0 && bookingCode != "" && v.PassengerID == passengerID && v.BookingCode == bookingCode {
			ret = append(ret, v)
		} else if passengerID > 0 && v.PassengerID == passengerID {
			ret = append(ret, v)
		} else if bookingCode != "" && v.BookingCode == bookingCode {
			ret = append(ret, v)
		} else if passengerID == 0 && bookingCode == "" {
			ret = append(ret, v)
		}
	}
	return ret
}

// Add Add new feedback
func (i *FeedbackDB) Add(data *Feedback) {
	i.userDB[data.FeedbackID] = data
}

// FindByID Find feedback by ID
func (i *FeedbackDB) FindByID(id string) *Feedback {
	if feedback, ok := i.userDB[id]; ok {
		return feedback
	}
	return nil
}

// DeleteByID Delete feedback by ID
func (i *FeedbackDB) DeleteByID(id string) {
	if _, ok := i.userDB[id]; ok {
		delete(i.userDB, id)
	}
}

// Count Total records in datastore
func (i *FeedbackDB) Count() int32 {
	return int32(len(i.userDB))
}
