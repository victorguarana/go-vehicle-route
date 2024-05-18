package csp

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/itinerary"
	"github.com/victorguarana/go-vehicle-route/src/slc"
)

func CoveringWithDrones(itineraryList []itinerary.Itinerary, m gps.Map, neighborhoodDistance float64) {
	clientsNeighborhood := gps.MapNeighborhood(m.Clients, neighborhoodDistance)
	for itineraryIndex := 0; len(clientsNeighborhood) > 0; itineraryIndex++ {
		itn := slc.CircularSelection(itineraryList, itineraryIndex)
		actualClient := gps.ClosestPointWithMostNeighbors(itn.ActualCarPoint(), clientsNeighborhood)
		itn.MoveCar(actualClient)
		deliverNeighborsWithDrones(itn, clientsNeighborhood[actualClient])
		removeClientAndItsNeighborsFromMap(actualClient, clientsNeighborhood)
		gps.RemovePointFromNearbyMap(actualClient, clientsNeighborhood)
	}

	finishItineraryOnClosestDeposits(itineraryList, m)
}

// TODO: Implement an function similar to this one
// but finish when drone can not support next client
func deliverNeighborsWithDrones(itn itinerary.Itinerary, neighbors []gps.Point) {
	if len(neighbors) <= 0 {
		return
	}
	droneNumbers := itn.DroneNumbers()
	actualCarPoint := itn.ActualCarPoint()
	actualCarStop := itn.ActualCarStop()
	// Need to set PackageSize to 0 to avoid drone support
	actualCarPoint.PackageSize = 0
	for neighborIndex, droneIndex := 0, 0; neighborIndex < len(neighbors); droneIndex++ {
		actualNeighbor := neighbors[neighborIndex]
		actualDroneNumber := slc.CircularSelection(droneNumbers, droneIndex)
		if itn.DroneSupport(actualDroneNumber, actualNeighbor, actualCarPoint) {
			itn.MoveDrone(actualDroneNumber, actualNeighbor)
			neighborIndex++
		} else {
			itn.LandDrone(actualDroneNumber, actualCarStop)
		}
	}
	itn.LandAllDrones(actualCarStop)
}

func removeClientAndItsNeighborsFromMap(client gps.Point, clientsNeighborhood gps.Neighborhood) {
	neighbors := clientsNeighborhood[client]
	for _, neighbor := range neighbors {
		gps.RemovePointFromNearbyMap(neighbor, clientsNeighborhood)
	}
	gps.RemovePointFromNearbyMap(client, clientsNeighborhood)
}

func finishItineraryOnClosestDeposits(itineraryList []itinerary.Itinerary, m gps.Map) {
	for _, itinerary := range itineraryList {
		position := itinerary.ActualCarPoint()
		closestDeposit := gps.ClosestPoint(position, m.Deposits)
		itinerary.MoveCar(closestDeposit)
	}
}