package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birds int
	for i := 0; i < len(birdsPerDay); i++ {
		birds += birdsPerDay[i]
	}
	return birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var weeklyBirds int
	if week == 1 {
		for i := 0; i <= len(birdsPerDay); i++ {
			weeklyBirds += birdsPerDay[i]
		}
	} else if week == 2 {
		for i := birdsPerDay[:8]; i <= len(birdsPerDay); i++ {
			weeklyBirds += birdsPerDay[i]
		}
	}
	return weeklyBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	panic("Please implement the FixBirdCountLog() function")
}
