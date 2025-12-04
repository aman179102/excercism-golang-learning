package birdwatcher

// TotalBirdCount return the total bird count by summing all days.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, birds := range birdsPerDay {
		total += birds
	}
	return total
}

// BirdsInWeek returns the total bird count for the given week (1-indexed).
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	end := start + 7
	total := 0
	for _, birds := range birdsPerDay[start:end] {
		total += birds
	}
	return total
}

// FixBirdCountLog increments the bird count for every second day (0th, 2nd, 4th…)
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
