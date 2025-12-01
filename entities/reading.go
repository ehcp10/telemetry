package entities

type Reading struct {
	ID        int
	DeviceID  int
	Type      string
	Value     float64
	Timestamp int64
}
