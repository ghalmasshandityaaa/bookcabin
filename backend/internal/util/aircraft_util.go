package util

import (
	"bookcabin-backend/internal/entity"
	"fmt"
	"math/rand"
	"time"
)

const RANDOM_SEAT_COUNT = 3

func GenerateUniqueSeats(assignedSeats []string, config *entity.AircraftConfig) ([]string, error) {
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
