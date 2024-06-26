package measure

import (
	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/itinerary"
)

func SpentFuel(itineraryInfo itinerary.Info) float64 {
	iterator := itineraryInfo.RouteIterator()
	var totalFuelSpent float64
	carEfficiency := itineraryInfo.CarEfficiency()
	droneEfficiency := itineraryInfo.DroneEfficiency()
	for iterator.HasNext() {
		actual := iterator.Actual()
		next := iterator.Next()
		totalFuelSpent += gps.ManhattanDistanceBetweenPoints(actual.Point(), next.Point()) / carEfficiency
		if subRoutes := actual.StartingSubRoutes(); len(subRoutes) > 0 {
			for _, subRoute := range subRoutes {
				totalFuelSpent += calcSubRouteDistance(subRoute) / droneEfficiency
			}
		}
		iterator.GoToNext()
	}
	return totalFuelSpent
}
