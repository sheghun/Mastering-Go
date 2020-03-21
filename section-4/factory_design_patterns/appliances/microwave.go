package appliances

type Microwave struct {
	typeName string
}

func (m *Microwave) Start() {
	m.typeName = "Microwave"
}

func (m *Microwave) GetPurpose() string {
	return "I am a " + m.typeName + "I heat food!!"
}
