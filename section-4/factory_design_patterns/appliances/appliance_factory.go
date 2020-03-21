package appliances

import "errors"

type Appliance interface {
	Start()
	GetPurpose() string
}

const (
	STOVE = iota
	FRIDGE
	MICROWAVE
)

func CreateAppliance(applianceType int) (Appliance, error) {
	switch applianceType {
	case STOVE:
		return new(Stove), nil
	case FRIDGE:
		return new(Fridge), nil
	case MICROWAVE:
		return new(Fridge), nil
	default:
		return nil, errors.New("invalid Appliance type")
	}
}
