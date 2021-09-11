package logtic

import "github.com/fatih/color"

func colorHiBlackString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return color.HiBlackString(m)
}

func colorBlueString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return color.BlueString(m)
}

func colorYellowString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return color.YellowString(m)
}

func colorRedString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return color.RedString(m)
}
