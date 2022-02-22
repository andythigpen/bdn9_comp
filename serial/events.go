package serial

import "errors"

const (
	EVENT_INVALID = iota
	EVENT_UNUSED
	EVENT_MUTE_SLACK
	EVENT_MUTE_TEAMS
	EVENT_FOCUS_SLACK
	EVENT_FOCUS_TEAMS
	EVENT_START_SLACK
	EVENT_START_TEAMS
	EVENT_END_CALL
	EVENT_MAX
)

type Event int

func (ev Event) IsValid() error {
	if ev > EVENT_INVALID && ev < EVENT_MAX {
		return nil
	}
	return errors.New("Invalid event")
}

type EventHandler interface {
	HandleEvent(d BDN9SerialDevice, ev Event)
}
