package serial

import (
	"errors"
	"fmt"
)

const (
	RGB_MATRIX_INVALID = iota
	RGB_MATRIX_SOLID_COLOR
	RGB_MATRIX_ALPHAS_MODS
	RGB_MATRIX_GRADIENT_UP_DOWN
	RGB_MATRIX_GRADIENT_LEFT_RIGHT
	RGB_MATRIX_BREATHING
	RGB_MATRIX_BAND_SAT
	RGB_MATRIX_BAND_VAL
	RGB_MATRIX_BAND_PINWHEEL_SAT
	RGB_MATRIX_BAND_PINWHEEL_VAL
	RGB_MATRIX_BAND_SPIRAL_SAT
	RGB_MATRIX_BAND_SPIRAL_VAL
	RGB_MATRIX_CYCLE_ALL
	RGB_MATRIX_CYCLE_LEFT_RIGHT
	RGB_MATRIX_CYCLE_UP_DOWN
	RGB_MATRIX_CYCLE_OUT_IN
	RGB_MATRIX_CYCLE_OUT_IN_DUAL
	RGB_MATRIX_RAINBOW_MOVING_CHEVRON
	RGB_MATRIX_CYCLE_PINWHEEL
	RGB_MATRIX_CYCLE_SPIRAL
	RGB_MATRIX_DUAL_BEACON
	RGB_MATRIX_RAINBOW_BEACON
	RGB_MATRIX_RAINBOW_PINWHEELS
	RGB_MATRIX_RAINDROPS
	RGB_MATRIX_JELLYBEAN_RAINDROPS
	RGB_MATRIX_TYPING_HEATMAP
	RGB_MATRIX_DIGITAL_RAIN
	RGB_MATRIX_SOLID_REACTIVE_SIMPLE
	RGB_MATRIX_SOLID_REACTIVE
	RGB_MATRIX_SOLID_REACTIVE_WIDE
	RGB_MATRIX_SOLID_REACTIVE_MULTIWIDE
	RGB_MATRIX_SOLID_REACTIVE_CROSS
	RGB_MATRIX_SOLID_REACTIVE_MULTICROSS
	RGB_MATRIX_SOLID_REACTIVE_NEXUS
	RGB_MATRIX_SOLID_REACTIVE_MULTINEXUS
	RGB_MATRIX_SPLASH
	RGB_MATRIX_MULTISPLASH
	RGB_MATRIX_SOLID_SPLASH
	RGB_MATRIX_SOLID_MULTISPLASH
	RGB_MATRIX_MAX
)

var rgbDesc = map[RGBMode]string{
	RGB_MATRIX_SOLID_COLOR:               "SOLID_COLOR               - Static single hue, no speed support",
	RGB_MATRIX_ALPHAS_MODS:               "ALPHAS_MODS               - Static dual hue, speed is hue for secondary hue",
	RGB_MATRIX_GRADIENT_UP_DOWN:          "GRADIENT_UP_DOWN          - Static gradient top to bottom, speed controls how much gradient changes",
	RGB_MATRIX_GRADIENT_LEFT_RIGHT:       "GRADIENT_LEFT_RIGHT       - Static gradient left to right, speed controls how much gradient changes",
	RGB_MATRIX_BREATHING:                 "BREATHING                 - Single hue brightness cycling animation",
	RGB_MATRIX_BAND_SAT:                  "BAND_SAT                  - Single hue band fading saturation scrolling left to right",
	RGB_MATRIX_BAND_VAL:                  "BAND_VAL                  - Single hue band fading brightness scrolling left to right",
	RGB_MATRIX_BAND_PINWHEEL_SAT:         "BAND_PINWHEEL_SAT         - Single hue 3 blade spinning pinwheel fades saturation",
	RGB_MATRIX_BAND_PINWHEEL_VAL:         "BAND_PINWHEEL_VAL         - Single hue 3 blade spinning pinwheel fades brightness",
	RGB_MATRIX_BAND_SPIRAL_SAT:           "BAND_SPIRAL_SAT           - Single hue spinning spiral fades saturation",
	RGB_MATRIX_BAND_SPIRAL_VAL:           "BAND_SPIRAL_VAL           - Single hue spinning spiral fades brightness",
	RGB_MATRIX_CYCLE_ALL:                 "CYCLE_ALL                 - Full keyboard solid hue cycling through full gradient",
	RGB_MATRIX_CYCLE_LEFT_RIGHT:          "CYCLE_LEFT_RIGHT          - Full gradient scrolling left to right",
	RGB_MATRIX_CYCLE_UP_DOWN:             "CYCLE_UP_DOWN             - Full gradient scrolling top to bottom",
	RGB_MATRIX_CYCLE_OUT_IN:              "CYCLE_OUT_IN              - Full gradient scrolling out to in",
	RGB_MATRIX_CYCLE_OUT_IN_DUAL:         "CYCLE_OUT_IN_DUAL         - Full dual gradients scrolling out to in",
	RGB_MATRIX_RAINBOW_MOVING_CHEVRON:    "RAINBOW_MOVING_CHEVRON    - Full gradent Chevron shapped scrolling left to right",
	RGB_MATRIX_CYCLE_PINWHEEL:            "CYCLE_PINWHEEL            - Full gradient spinning pinwheel around center of keyboard",
	RGB_MATRIX_CYCLE_SPIRAL:              "CYCLE_SPIRAL              - Full gradient spinning spiral around center of keyboard",
	RGB_MATRIX_DUAL_BEACON:               "DUAL_BEACON               - Full gradient spinning around center of keyboard",
	RGB_MATRIX_RAINBOW_BEACON:            "RAINBOW_BEACON            - Full tighter gradient spinning around center of keyboard",
	RGB_MATRIX_RAINBOW_PINWHEELS:         "RAINBOW_PINWHEELS         - Full dual gradients spinning two halfs of keyboard",
	RGB_MATRIX_RAINDROPS:                 "RAINDROPS                 - Randomly changes a single key's hue",
	RGB_MATRIX_JELLYBEAN_RAINDROPS:       "JELLYBEAN_RAINDROPS       - Randomly changes a single key's hue and saturation",
	RGB_MATRIX_TYPING_HEATMAP:            "TYPING_HEATMAP            - How hot is your WPM!",
	RGB_MATRIX_DIGITAL_RAIN:              "DIGITAL_RAIN              - That famous computer simulation",
	RGB_MATRIX_SOLID_REACTIVE_SIMPLE:     "SOLID_REACTIVE_SIMPLE     - Pulses keys hit to hue & value then fades value out",
	RGB_MATRIX_SOLID_REACTIVE:            "SOLID_REACTIVE            - Static single hue, pulses keys hit to shifted hue then fades to current hue",
	RGB_MATRIX_SOLID_REACTIVE_WIDE:       "SOLID_REACTIVE_WIDE       - Hue & value pulse near a single key hit then fades value out",
	RGB_MATRIX_SOLID_REACTIVE_MULTIWIDE:  "SOLID_REACTIVE_MULTIWIDE  - Hue & value pulse near multiple key hits then fades value out",
	RGB_MATRIX_SOLID_REACTIVE_CROSS:      "SOLID_REACTIVE_CROSS      - Hue & value pulse the same column and row of a single key hit then fades value out",
	RGB_MATRIX_SOLID_REACTIVE_MULTICROSS: "SOLID_REACTIVE_MULTICROSS - Hue & value pulse the same column and row of multiple key hits then fades value out",
	RGB_MATRIX_SOLID_REACTIVE_NEXUS:      "SOLID_REACTIVE_NEXUS      - Hue & value pulse away on the same column and row of a single key hit then fades value out",
	RGB_MATRIX_SOLID_REACTIVE_MULTINEXUS: "SOLID_REACTIVE_MULTINEXUS - Hue & value pulse away on the same column and row of multiple key hits then fades value out",
	RGB_MATRIX_SPLASH:                    "SPLASH                    - Full gradient & value pulse away from a single key hit then fades value out",
	RGB_MATRIX_MULTISPLASH:               "MULTISPLASH               - Full gradient & value pulse away from multiple key hits then fades value out",
	RGB_MATRIX_SOLID_SPLASH:              "SOLID_SPLASH              - Hue & value pulse away from a single key hit then fades value out",
	RGB_MATRIX_SOLID_MULTISPLASH:         "SOLID_MULTISPLASH         - Hue & value pulse away from multiple key hits then fades value out",
}

type RGBMode int

func (m RGBMode) IsValid() error {
	if m > RGB_MATRIX_INVALID && m < RGB_MATRIX_MAX {
		return nil
	}
	return errors.New("Invalid RGB Mode")
}

func (m RGBMode) String() string {
	return fmt.Sprintf("%2d: %s", m, rgbDesc[m])
}
