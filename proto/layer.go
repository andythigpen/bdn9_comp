package proto

import (
	"errors"
	"fmt"
)

const KEY_MAX = 12
const KEY_UPPER_MAX = 9

var layerDesc = map[Layer]string{
	Layer_LAYER_DEFAULT: "DEFAULT      - Main (default) layer",
	Layer_LAYER_TEAMS:   "TEAMS        - Teams call macros",
	Layer_LAYER_SLACK:   "SLACK        - Slack call macros",
}

func (l Layer) IsValid() error {
	if l >= Layer_LAYER_DEFAULT && l < Layer_LAYER_MAX {
		return nil
	}
	return errors.New("Invalid layer")
}

func (l Layer) Description() string {
	return fmt.Sprintf("%2d: %s", l, layerDesc[l])
}
