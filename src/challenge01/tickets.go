package main

import (
    "fmt"
    "math/rand"
)

const DISTANCE = 62100000.0  // km

func main() {
    spacelines := []string{"Space Adventures", "SpaceX", "Virgin Galactic"}

    fmt.Println("Spaceline         Days  Trip type    Price")
    fmt.Println("==========================================")
    for i := 0; i < 10; i++ {
        spaceline := spacelines[rand.Intn(len(spacelines))]
        velocity := 16 + rand.Intn(30 - 16 + 1)  // km/sec
        days := int(DISTANCE / float64(velocity) / 60 / 60 / 24)
        isRoundTrip := rand.Intn(2) == 1
        tripType := "One-way"
        price := 36 + int((50.0 - 36.0) / (30.0 - 16.0) * (float64(velocity) - 16.0))
        if isRoundTrip {
            tripType = "Round-trip"
            price *= 2
        }

        fmt.Printf("%-16v %5v  %-12v $%4v\n", spaceline, days, tripType, price)
    }
}
