package logtic

func colorHiBlackString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return colorGray + m + colorReset
}

func colorBlueString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return colorBlue + m + colorReset
}

func colorYellowString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return colorYellow + m + colorReset
}

func colorRedString(m string) string {
	if !Log.Options.Color {
		return m
	}
	return colorRed + m + colorReset
}
