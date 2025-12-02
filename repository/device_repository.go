package repository

type DeviceRepository interface {
	Save(device *Device) (*Device, error)
	FindByID(id int) (*Device, error)
	FindByUUID(uuid string) (*Device, error)
	Get() ([]*Device, error)
	Exists() (uuid string) (bool, error)	
	List() ([]*Device, error)
}