package entities

type Reading struct {
	ID        int
	DeviceID  int
	Type      string `validate:"required,oneof=Temperature Humidity Pressure"`
	Value     float64
	Timestamp int64
}

type RadingInput struct {
	DeviceID  int
	Type      string `validate:"required,oneof=Temperature Humidity Pressure"`
	Value     float64
	Timestamp int64
}

func NewReading(dto RadingInput) *Reading {
	return &Reading{
		DeviceID:  dto.DeviceID,
		Type:      dto.Type,
		Value:     dto.Value,
		Timestamp: dto.Timestamp,
	}
}
