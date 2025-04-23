package meet

import (
	"fmt"
)

// meetStatus is a domain-specific enumeration.
// It is unexported to prevent invalid or inconsistent values being created outside this package.
// It enforces safe construction via NewMeetStatus() and safe status transition using Transition().
// safe transition logic with explicit
type meetStatus string

const (
	CREATED     meetStatus = "Created"
	CANCELED    meetStatus = "Canceled"
	SCHEDULED   meetStatus = "Scheduled"
	IN_PROGRESS meetStatus = "InProgress"
	COMPLETED   meetStatus = "Completed"
)

var validMeetStatuses = map[meetStatus]struct{}{
	CREATED:     {},
	CANCELED:    {},
	SCHEDULED:   {},
	IN_PROGRESS: {},
	COMPLETED:   {},
}

func (m meetStatus) String() string {
	_, ok := validMeetStatuses[m]
	if !ok {
		return fmt.Sprintf("MeetStatus(%s)", string(m))
	}
	return string(m)
}

// NewMeetStatus validates and returns a meetStatus.
// Returns an error if input is invalid.
func NewMeetStatus(s string) (meetStatus, error) {
	status := meetStatus(s)
	if _, ok := validMeetStatuses[status]; !ok {
		return "", fmt.Errorf("invalid meet status: %q", s)
	}
	return status, nil
}

// Transition validates the current and target status and applies allowed transitions.
// Returns an error if transition is not allowed.
func (m meetStatus) Transition(target meetStatus) (meetStatus, error) {
	if _, ok := validMeetStatuses[m]; !ok {
		return "", fmt.Errorf("invalid current meet status: %q", m)
	}
	if _, ok := validMeetStatuses[target]; !ok {
		return "", fmt.Errorf("invalid target meet status: %q", target)
	}

	switch m {
	case CREATED:
		if target == CANCELED || target == SCHEDULED {
			return target, nil
		}
	case SCHEDULED:
		if target == CANCELED || target == IN_PROGRESS {
			return target, nil
		}
	case IN_PROGRESS:
		if target == COMPLETED {
			return target, nil
		}
	}
	return "", fmt.Errorf("transition %s → %s is not allowed", m, target)
}
