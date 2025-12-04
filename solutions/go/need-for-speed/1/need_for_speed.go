package speed

// Car struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

// Track struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive moves the car if battery allows
func Drive(car Car) Car {
	if car.battery < car.batteryDrain {
		return car
	}
	car.battery -= car.batteryDrain
	car.distance += car.speed
	return car
}

// CanFinish checks if the car can finish the given track
func CanFinish(car Car, track Track) bool {
	// How many drives needed to finish?
	drivesNeeded := track.distance / car.speed

	// Total battery required
	totalDrain := drivesNeeded * car.batteryDrain

	return car.battery >= totalDrain
}
