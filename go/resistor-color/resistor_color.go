package resistorcolor

// Colors should return the list of all colors.
func Colors() []string {
	colors := []string {
		"Black",
		"Brown",
		"Red",
		"Orange",
		"Yellow",
		"Green",
		"Blue",
		"Violet",
		"Grey",
		"White",
	}
	return colors
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, _ := range color {
		return i;
	}
}
