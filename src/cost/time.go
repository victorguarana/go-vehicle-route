package cost

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/routes"
)

type subRouteTimes map[routes.ISubRoute]float64

// TODO: How to get the exact drone that made that flight?
// Actual implementation is considering that the vehicles always have default speed
func CalcTime(itn itinerary.Itinerary) float64 {
	var subRoutesFlyingTimes = make(subRouteTimes)
	var mainRouteTravelTime = make(subRouteTimes)
	var totalTime float64
	carSpeed := itn.CarSpeed()
	droneSpeed := itn.DroneSpeed()
	iterator := itn.RouteIterator()
	for {
		actual := iterator.Actual()
		if subRoutes := actual.StartingSubRoutes(); len(subRoutes) > 0 {
			calcSubRouteTimes(mainRouteTravelTime, subRoutesFlyingTimes, subRoutes, droneSpeed)
		}
		if subRoutes := actual.ReturningSubRoutes(); len(subRoutes) > 0 {
			totalTime += maxAdditionalTimeWaitingSubRoutes(mainRouteTravelTime, subRoutesFlyingTimes, subRoutes)
			removeReturningSubRoutes(mainRouteTravelTime, subRoutesFlyingTimes, subRoutes)
		}
		if !iterator.HasNext() {
			break
		}
		next := iterator.Next()
		travelTime := gps.DistanceBetweenPoints(actual.Point(), next.Point()) / carSpeed
		updateMainRouteTravelTimes(mainRouteTravelTime, travelTime)
		totalTime += travelTime
		iterator.GoToNext()
	}

	return totalTime
}

func calcSubRouteTimes(mainRouteTravelTime subRouteTimes, subRouteFlyingTimes subRouteTimes, subRoutes []routes.ISubRoute, droneSpeed float64) {
	for _, subRoute := range subRoutes {
		mainRouteTravelTime[subRoute] = 0
		subRouteFlyingTimes[subRoute] = calcSubRouteDistance(subRoute) / droneSpeed
	}
}

func maxAdditionalTimeWaitingSubRoutes(mainRouteTravelTime subRouteTimes, subRoutesFlyingTimes subRouteTimes, subRoutes []routes.ISubRoute) float64 {
	var maxWaitingTime float64
	for _, subRoute := range subRoutes {
		timeDifference := subRoutesFlyingTimes[subRoute] - mainRouteTravelTime[subRoute]
		if timeDifference > maxWaitingTime {
			maxWaitingTime = timeDifference
		}
	}
	return maxWaitingTime
}

func updateMainRouteTravelTimes(mainRouteTravelTime subRouteTimes, travelTime float64) {
	for subRoute := range mainRouteTravelTime {
		mainRouteTravelTime[subRoute] += travelTime
	}
}

func removeReturningSubRoutes(mainRouteTravelTime subRouteTimes, subRoutesFlyingTimes subRouteTimes, subRoutes []routes.ISubRoute) {
	for _, subRoute := range subRoutes {
		delete(mainRouteTravelTime, subRoute)
		delete(subRoutesFlyingTimes, subRoute)
	}
}
