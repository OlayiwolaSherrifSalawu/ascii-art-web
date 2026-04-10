package ascii

func Color(color string) (string, string) {
	colors := map[string]string{
		"black":  "\033[30m",
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",

		"orange":  "\033[38;5;208m",
		"brown":   "\033[38;5;94m",
		"pink":    "\033[38;5;205m",
		"gold":    "\033[38;5;220m",
		"lime":    "\033[38;5;118m",
		"teal":    "\033[38;5;30m",
		"navy":    "\033[38;5;17m",
		"magenta": "\033[38;5;201m",
		"sky":     "\033[38;5;117m",
		"gray":    "\033[38;5;244m",
	}

	reset := "\033[0m"

	val, ok := colors[color]
	if !ok {
		return "", ""
	}
	return val, reset
}
