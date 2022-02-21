package proto

import (
	"errors"
	"fmt"
)

var rgbDesc = map[RGBMode]string{
	RGBMode_RGB_MATRIX_SOLID_COLOR:               "SOLID_COLOR               - Static single hue, no speed support",
	RGBMode_RGB_MATRIX_ALPHAS_MODS:               "ALPHAS_MODS               - Static dual hue, speed is hue for secondary hue",
	RGBMode_RGB_MATRIX_GRADIENT_UP_DOWN:          "GRADIENT_UP_DOWN          - Static gradient top to bottom, speed controls how much gradient changes",
	RGBMode_RGB_MATRIX_GRADIENT_LEFT_RIGHT:       "GRADIENT_LEFT_RIGHT       - Static gradient left to right, speed controls how much gradient changes",
	RGBMode_RGB_MATRIX_BREATHING:                 "BREATHING                 - Single hue brightness cycling animation",
	RGBMode_RGB_MATRIX_BAND_SAT:                  "BAND_SAT                  - Single hue band fading saturation scrolling left to right",
	RGBMode_RGB_MATRIX_BAND_VAL:                  "BAND_VAL                  - Single hue band fading brightness scrolling left to right",
	RGBMode_RGB_MATRIX_BAND_PINWHEEL_SAT:         "BAND_PINWHEEL_SAT         - Single hue 3 blade spinning pinwheel fades saturation",
	RGBMode_RGB_MATRIX_BAND_PINWHEEL_VAL:         "BAND_PINWHEEL_VAL         - Single hue 3 blade spinning pinwheel fades brightness",
	RGBMode_RGB_MATRIX_BAND_SPIRAL_SAT:           "BAND_SPIRAL_SAT           - Single hue spinning spiral fades saturation",
	RGBMode_RGB_MATRIX_BAND_SPIRAL_VAL:           "BAND_SPIRAL_VAL           - Single hue spinning spiral fades brightness",
	RGBMode_RGB_MATRIX_CYCLE_ALL:                 "CYCLE_ALL                 - Full keyboard solid hue cycling through full gradient",
	RGBMode_RGB_MATRIX_CYCLE_LEFT_RIGHT:          "CYCLE_LEFT_RIGHT          - Full gradient scrolling left to right",
	RGBMode_RGB_MATRIX_CYCLE_UP_DOWN:             "CYCLE_UP_DOWN             - Full gradient scrolling top to bottom",
	RGBMode_RGB_MATRIX_CYCLE_OUT_IN:              "CYCLE_OUT_IN              - Full gradient scrolling out to in",
	RGBMode_RGB_MATRIX_CYCLE_OUT_IN_DUAL:         "CYCLE_OUT_IN_DUAL         - Full dual gradients scrolling out to in",
	RGBMode_RGB_MATRIX_RAINBOW_MOVING_CHEVRON:    "RAINBOW_MOVING_CHEVRON    - Full gradent Chevron shapped scrolling left to right",
	RGBMode_RGB_MATRIX_CYCLE_PINWHEEL:            "CYCLE_PINWHEEL            - Full gradient spinning pinwheel around center of keyboard",
	RGBMode_RGB_MATRIX_CYCLE_SPIRAL:              "CYCLE_SPIRAL              - Full gradient spinning spiral around center of keyboard",
	RGBMode_RGB_MATRIX_DUAL_BEACON:               "DUAL_BEACON               - Full gradient spinning around center of keyboard",
	RGBMode_RGB_MATRIX_RAINBOW_BEACON:            "RAINBOW_BEACON            - Full tighter gradient spinning around center of keyboard",
	RGBMode_RGB_MATRIX_RAINBOW_PINWHEELS:         "RAINBOW_PINWHEELS         - Full dual gradients spinning two halfs of keyboard",
	RGBMode_RGB_MATRIX_RAINDROPS:                 "RAINDROPS                 - Randomly changes a single key's hue",
	RGBMode_RGB_MATRIX_JELLYBEAN_RAINDROPS:       "JELLYBEAN_RAINDROPS       - Randomly changes a single key's hue and saturation",
	RGBMode_RGB_MATRIX_TYPING_HEATMAP:            "TYPING_HEATMAP            - How hot is your WPM!",
	RGBMode_RGB_MATRIX_DIGITAL_RAIN:              "DIGITAL_RAIN              - That famous computer simulation",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_SIMPLE:     "SOLID_REACTIVE_SIMPLE     - Pulses keys hit to hue & value then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE:            "SOLID_REACTIVE            - Static single hue, pulses keys hit to shifted hue then fades to current hue",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_WIDE:       "SOLID_REACTIVE_WIDE       - Hue & value pulse near a single key hit then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_MULTIWIDE:  "SOLID_REACTIVE_MULTIWIDE  - Hue & value pulse near multiple key hits then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_CROSS:      "SOLID_REACTIVE_CROSS      - Hue & value pulse the same column and row of a single key hit then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_MULTICROSS: "SOLID_REACTIVE_MULTICROSS - Hue & value pulse the same column and row of multiple key hits then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_NEXUS:      "SOLID_REACTIVE_NEXUS      - Hue & value pulse away on the same column and row of a single key hit then fades value out",
	RGBMode_RGB_MATRIX_SOLID_REACTIVE_MULTINEXUS: "SOLID_REACTIVE_MULTINEXUS - Hue & value pulse away on the same column and row of multiple key hits then fades value out",
	RGBMode_RGB_MATRIX_SPLASH:                    "SPLASH                    - Full gradient & value pulse away from a single key hit then fades value out",
	RGBMode_RGB_MATRIX_MULTISPLASH:               "MULTISPLASH               - Full gradient & value pulse away from multiple key hits then fades value out",
	RGBMode_RGB_MATRIX_SOLID_SPLASH:              "SOLID_SPLASH              - Hue & value pulse away from a single key hit then fades value out",
	RGBMode_RGB_MATRIX_SOLID_MULTISPLASH:         "SOLID_MULTISPLASH         - Hue & value pulse away from multiple key hits then fades value out",
}

func (m RGBMode) IsValid() error {
	if m > RGBMode_RGB_MATRIX_INVALID && m < RGBMode_RGB_MATRIX_MAX {
		return nil
	}
	return errors.New("Invalid RGB Mode")
}

func (m RGBMode) Description() string {
	return fmt.Sprintf("%2d: %s", m, rgbDesc[m])
}
