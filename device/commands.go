package device

import "errors"

const (
	COMMAND_INVALID           = iota // sets the RGB matrix animation mode
	COMMAND_RESET                    // resets the keyboard for flashing
	COMMAND_SET_MODE                 // sets the RGB matrix animation mode
	COMMAND_TOGGLE_MATRIX            // turns on/off the LED matrix
	COMMAND_SET_MATRIX_HSV           // sets HSV for the entire matrix
	COMMAND_TOGGLE_INDICATOR         // turns on/off a specific LED indicator
	COMMAND_SET_INDICATOR_RGB        // sets RGB for a specific LED
	COMMAND_ENABLE_INDICATOR         // turns on a specific LED indicator
	COMMAND_DISABLE_INDICATOR        // turns off a specific LED indicator
	COMMAND_ACTIVATE_LAYER           // enables a specific layer
	COMMAND_SET_SPEED                // sets the animation speed
	COMMAND_SET_MUTE_STATUS          // sets the mute status
	COMMAND_END_CALL                 // clears call status and goes back to default layer
	COMMAND_ECHO                     // test command, just echos back what you send
	COMMAND_CONNECT                  // send after connecting
	COMMAND_DISCONNECT               // send before disconnecting
	COMMAND_MAX                      // not a valid command
)

type Command int

func (c Command) IsValid() error {
	if c > COMMAND_INVALID && c < COMMAND_MAX {
		return nil
	}
	return errors.New("Invalid command")
}
