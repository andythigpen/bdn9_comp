package serial

import (
	"errors"
	"fmt"
)

const (
	LAYER_DEFAULT = iota
	LAYER_PROGRAMMING
	LAYER_DEBUGGING
	LAYER_SLACK
	LAYER_TEAMS
	LAYER_MAX // not a valid layer
)

var layerDesc = map[Layer]string{
	LAYER_DEFAULT:     "DEFAULT      - Main (default) layer",
	LAYER_PROGRAMMING: "PROGRAMMING  - Nvim test, runner, and debug start macros",
	LAYER_DEBUGGING:   "DEBUGGING    - Nvim debugging macros",
	LAYER_SLACK:       "SLACK        - Slack call macros",
	LAYER_TEAMS:       "TEAMS        - Teams call macros",
}

type Layer int

func (l Layer) IsValid() error {
	if l >= LAYER_DEFAULT && l < LAYER_MAX {
		return nil
	}
	return errors.New("Invalid layer")
}

func (l Layer) String() string {
	return fmt.Sprintf("%2d: %s", l, layerDesc[l])
}
