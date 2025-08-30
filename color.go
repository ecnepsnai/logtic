package logtic

// IColor describes an interface for applying color to a message. By default, logtic only supports ANSI color codes.
// On platforms where these codes don't work - such as the legacy Console Host on Windows, you may implement your own
// color interface that applies colors as need for your platform.
type IColor interface {
	// Apply a bright black (gray) color to m
	HiBlack(m string) string
	// Apply a blue color to m
	Blue(m string) string
	// Apply a yellow color to m
	Yellow(m string) string
	// Apply a red color to m
	Red(m string) string
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorGray   = "\033[90m"
)

type tDefaultColor struct{}

func (*tDefaultColor) HiBlack(m string) string {
	return colorGray + m + colorReset
}
func (*tDefaultColor) Blue(m string) string {
	return colorBlue + m + colorReset
}
func (*tDefaultColor) Yellow(m string) string {
	return colorYellow + m + colorReset
}
func (*tDefaultColor) Red(m string) string {
	return colorRed + m + colorReset
}

func (s *Source) colorHiBlack(m string) string {
	if s.instance.Color == nil {
		return m
	}

	return s.instance.Color.HiBlack(m)
}

func (s *Source) colorBlue(m string) string {
	if s.instance.Color == nil {
		return m
	}

	return s.instance.Color.Blue(m)
}

func (s *Source) colorYellow(m string) string {
	if s.instance.Color == nil {
		return m
	}

	return s.instance.Color.Yellow(m)
}

func (s *Source) colorRed(m string) string {
	if s.instance.Color == nil {
		return m
	}

	return s.instance.Color.Red(m)
}
