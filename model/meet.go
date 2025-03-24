package model

import (
	"fmt"
	"time"
)

type Meet struct {
	ID          string
	Title       string
	Status      MeetStatus
	From        time.Time
	To          time.Time
	Description string
	Link        string

	// many 2 many
	Participants []Participant

	OrganizerID Participant
	CreatedBy   Participant
	CreatedAt   time.Time
}

// NewMeet init meet with given fields.
// Set status SCHEDULED.
func NewMeet(
	title string,
	from, to time.Time,
	description string,
	creatorID string,
	organizerID string,
	attenderIDs []string,
) Meet {
	return newMeet(title, CREATED, from, to, description, creatorID, organizerID, attenderIDs)
}

// NewScheduledMeet init meet with given fields.
// Set status CREATED.
func NewScheduledMeet(
	title string,
	from, to time.Time,
	description string,
	creatorID string,
	organizerID string,
	attenderIDs []string,
) Meet {
	return newMeet(title, SCHEDULED, from, to, description, creatorID, organizerID, attenderIDs)
}

// newMeet init meet with given fields.
func newMeet(
	title string,
	status MeetStatus,
	from, to time.Time,
	description string,
	creatorID string,
	organizerID string,
	attenderIDs []string,
) Meet {
	meet := Meet{
		Title:       title,
		Status:      status,
		From:        from,
		To:          to,
		Description: description,
		OrganizerID: Participant{
			ID: organizerID,
		},
		CreatedBy: Participant{
			ID: creatorID,
		},
	}

	// add participants to meet
	for _, id := range attenderIDs {
		meet.Participants = append(meet.Participants, Participant{ID: id})
	}
	meet.Participants = append(
		meet.Participants,
		[]Participant{
			{ID: organizerID},
			{ID: creatorID},
		}...)

	return meet
}

func (m *Meet) AddParticipants(participantIDs []string) error {
	if m == nil {
		return nil
	}

	if m.Participants == nil {
		m.Participants = make([]Participant, 0, len(participantIDs))
	}

	for _, participant := range participantIDs {
		err := m.AddParticipant(participant)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *Meet) AddParticipant(participantID string) error {
	if m == nil {
		return nil
	}
	if participantID == "" {
		return fmt.Errorf("validation error: empty participant ID")
	}

	// find or upsert in cache
	//_, found :=
	//if found {
	//	return fmt.Errorf("duplicate error: participant id:%s already exists", participant.ID)
	//}
	//m.Participants[participant.ID] = *participant

	m.Participants = append(m.Participants, Participant{ID: participantID})
	return nil
}

func (m *Meet) DeleteParticipant(participantID string) error {
	if participantID == "" {
		return fmt.Errorf("validation error: empty participant ID")
	}

	for i, participant := range m.Participants {
		if participant.ID == participantID {
			m.Participants[i] = m.Participants[len(m.Participants)-1]
			m.Participants = m.Participants[:len(m.Participants)-1]
		}
	}

	return nil
}
