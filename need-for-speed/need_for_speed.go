package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain,battery: 100}
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {

	if car.batteryDrain > car.battery {
		return car
	} else {
		return Car{speed: car.speed, distance: car.speed, battery: car.battery - car.batteryDrain, batteryDrain: car.batteryDrain}
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return track.distance <= (((car.battery /car.batteryDrain) * car.speed) + car.distance)
}
