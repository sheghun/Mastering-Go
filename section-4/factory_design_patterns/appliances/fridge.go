package appliances

type Fridge struct {
	typeName string
}

func (f *Fridge) Start() {
	f.typeName = "Fridge"
}

func (f *Fridge) GetPurpose() string {
	return "I am a " + f.typeName + " I cool stuff down!"
}
