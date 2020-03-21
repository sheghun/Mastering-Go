package appliances

type Stove struct {
	typeName string
}

func (s *Stove) Start() {
	s.typeName = "Stove"
}

func (s *Stove) GetPurpose() string {
	return "I am a " + s.typeName + " I cool stuff down!"
}
