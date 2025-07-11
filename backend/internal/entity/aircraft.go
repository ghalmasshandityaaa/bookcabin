package entity

type AircraftType string

const (
	ATR          AircraftType = "ATR"
	Airbus320    AircraftType = "Airbus 320"
	Boeing737Max AircraftType = "Boeing 737 Max"
)

type Aircraft map[AircraftType]AircraftConfig
type AircraftConfig struct {
	MaxRow int
	Seats  []string
}

func NewAircraft() *Aircraft {
	return &Aircraft{
		ATR: {
			MaxRow: 18,
			Seats:  []string{"A", "C", "D", "F"},
		},
		Airbus320: {
			MaxRow: 32,
			Seats:  []string{"A", "B", "C", "D", "E", "F"},
		},
		Boeing737Max: {
			MaxRow: 32,
			Seats:  []string{"A", "B", "C", "D", "E", "F"},
		},
	}
}

func (a *Aircraft) GetAircraftConfig(aircraftType AircraftType) *AircraftConfig {
	config, ok := (*a)[aircraftType]
	if !ok {
		return nil
	}
	return &config
}

// func (a *Aircraft) ToJSON() []map[string]any {
// 	var result []map[string]any
// 	for aircraftType, config := range *a {
// 		result = append(result, config.ToJSON(aircraftType))
// 	}
// 	return result
// }

// func (ac *AircraftConfig) ToJSON(aircraftType AircraftType) map[string]any {
// 	return map[string]any{
// 		"aircraft": string(aircraftType),
// 		"max_row":  ac.MaxRow,
// 		"seats":    ac.Seats,
// 	}
// }
