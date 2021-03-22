package serial

import "errors"

const (
	LAYER_DEFAULT = iota
	LAYER_PROGRAMMING
	LAYER_DEBUGGING
	LAYER_SLACK
	LAYER_TEAMS
	LAYER_MAX // not a valid layer
)

type Layer int

func (l Layer) IsValid() error {
	if l >= LAYER_DEFAULT && l < LAYER_MAX {
		return nil
	}
	return errors.New("Invalid layer")
}
