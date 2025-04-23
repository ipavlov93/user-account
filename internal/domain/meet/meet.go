package meet

import (
	"event-calendar/internal/domain"
	"fmt"
	"time"
)

type Meet struct {
	ID          int64
	Title       string
	Status      meetStatus
	StartedAt   time.Time
	FinishedAt  time.Time
	Description string
	Link        string

	// many 2 many
	Participants []domain.User

	OrganizerID domain.User
	CreatedBy   domain.User
}

// NewMeet init meet with given fields.
// Set status CREATED.
func NewMeet(
	title string,
	startedAt time.Time,
	finishedAt time.Time,
	description string,
	creatorID int64,
	organizerID int64,
	attenderIDs []int64,
) Meet {
	return newMeet(title, CREATED, startedAt, finishedAt, description, creatorID, organizerID, attenderIDs)
}

// NewScheduledMeet init meet with given fields.
// Set status SCHEDULED.
func NewScheduledMeet(
	title string,
	startedAt time.Time,
	finishedAt time.Time,
	description string,
	creatorID int64,
	organizerID int64,
	attenderIDs []int64,
) Meet {
	return newMeet(title, SCHEDULED, startedAt, finishedAt, description, creatorID, organizerID, attenderIDs)
}

// newMeet init meet with given fields.
func newMeet(
	title string,
	status meetStatus,
	startedAt time.Time,
	finishedAt time.Time,
	description string,
	creatorID int64,
	organizerID int64,
	attenderIDs []int64,
) Meet {
	meet := Meet{
		Title:       title,
		Status:      status,
		StartedAt:   startedAt,
		FinishedAt:  finishedAt,
		Description: description,
		OrganizerID: domain.User{
			ID: organizerID,
		},
		CreatedBy: domain.User{
			ID: creatorID,
		},
	}

	// add participants to meet
	for _, id := range attenderIDs {
		meet.Participants = append(meet.Participants, domain.User{ID: id})
	}
	meet.Participants = append(
		meet.Participants,
		[]domain.User{
			{ID: organizerID},
			{ID: creatorID},
		}...)

	return meet
}

func (m *Meet) AddParticipants(participantIDs []int64) error {
	if m == nil {
		return nil
	}

	if m.Participants == nil {
		m.Participants = make([]domain.User, 0, len(participantIDs))
	}

	for _, participant := range participantIDs {
		err := m.AddParticipant(participant)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Meet) AddParticipant(participantID int64) error {
	if m == nil {
		return nil
	}
	if participantID < 1 {
		return fmt.Errorf("validation error: invalid participant ID")
	}

	// find or upsert in cache
	//_, found :=
	//if found {
	//	return fmt.Errorf("duplicate error: participant id:%s already exists", participant.ID)
	//}
	//m.Participants[participant.ID] = *participant

	m.Participants = append(m.Participants, domain.User{ID: participantID})
	return nil
}

func (m *Meet) DeleteParticipant(participantID int64) error {
	if participantID < 1 {
		return nil
	}

	for i, participant := range m.Participants {
		if participant.ID == participantID {
			m.Participants[i] = m.Participants[len(m.Participants)-1]
			m.Participants = m.Participants[:len(m.Participants)-1]
		}
	}

	return nil
}
