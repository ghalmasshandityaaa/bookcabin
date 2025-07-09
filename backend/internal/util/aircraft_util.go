package util

import (
	"bookcabin-backend/internal/model"
	"fmt"
	"math/rand"
	"time"
)

const RANDOM_SEAT_COUNT = 3

type AircraftConfig struct {
	MaxRow int
	Seats  []string
}

var aircraftConfigs = map[model.AircraftType]AircraftConfig{
	model.ATR: {
		MaxRow: 18,
		Seats:  []string{"A", "C", "D", "F"},
	},
	model.Airbus320: {
		MaxRow: 32,
		Seats:  []string{"A", "B", "C", "D", "E", "F"},
	},
	model.Boeing737Max: {
		MaxRow: 32,
		Seats:  []string{"A", "B", "C", "D", "E", "F"},
	},
}

func GenerateUniqueSeats(assignedSeats []string, aircraftType model.AircraftType) ([]string, error) {
	config, ok := aircraftConfigs[aircraftType]
	if !ok {
		return nil, fmt.Errorf("aircraft/unsupported-type")
	}

	// Convert assignedSeats to map for O(1) lookup
	assignedMap := make(map[string]struct{}, len(assignedSeats))
	for _, s := range assignedSeats {
		assignedMap[s] = struct{}{}
	}

	// Generate all valid seats
	var allSeats []string
	for row := 1; row <= config.MaxRow; row++ {
		for _, seatLetter := range config.Seats {
			seat := fmt.Sprintf("%d%s", row, seatLetter)
			if _, taken := assignedMap[seat]; !taken {
				allSeats = append(allSeats, seat)
			}
		}
	}

	if len(allSeats) < RANDOM_SEAT_COUNT {
		return nil, fmt.Errorf("seats/not-enough")
	}

	// Use local RNG (thread-safe, no rand.Seed required)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(allSeats), func(i, j int) {
		allSeats[i], allSeats[j] = allSeats[j], allSeats[i]
	})

	return allSeats[:RANDOM_SEAT_COUNT], nil
}
