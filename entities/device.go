package entities

type Device struct {
	ID               int
	UUID             string
	Name             string
	FramewareVersion string
	Status           string
	CreatedAt        int64
}
