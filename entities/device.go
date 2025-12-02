package entities

type Device struct {
	ID               int
	UUID             string `validate:"required,uuid4"`
	Name             string
	FramewareVersion string
	Status           string `validate:"required,oneof=Active Inactive"`
	CreatedAt        int64
}

type RegisterDeviceInput struct {
	UUID             string
	Name             string
	FramewareVersion string
	Status           string
}

func NewDevice(dto RegisterDeviceInput) *Device {
	return &Device{
		UUID:             dto.UUID,
		Name:             dto.Name,
		FramewareVersion: dto.FramewareVersion,
		Status:           dto.Status,
	}
}
