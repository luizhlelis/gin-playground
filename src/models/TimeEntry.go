package models

// TimeEntry entry that register employee worked hours
type TimeEntry struct {
	TemperatureCelsius    int    `form:"temperatureCelsius" json:"temperatureCelsius" xml:"temperatureCelsius" binding:"required"`
	TemperatureFahrenheit int    `form:"temperatureFahrenheit" json:"temperatureFahrenheit" xml:"temperatureFahrenheit" binding:"-"`
	Summary               string `form:"summary" json:"summary" xml:"summary" binding:"required"`
}

// NewTimeEntry returns a pointer to a new TimeEntry struct
func NewTimeEntry(temperatureCelsius int, summary string) (*TimeEntry, error) {

	timeEntry := &TimeEntry{
		TemperatureCelsius:    temperatureCelsius,
		TemperatureFahrenheit: 32 + (int)(float64(temperatureCelsius)/0.5556),
		Summary:               summary,
	}

	return timeEntry, nil
}
