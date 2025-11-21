package moment

type Originator interface {
	SetMomento(m Momento)
	CreateMomento()
}

type Momento interface {
	GetState()
	SetState()
}

type Directory struct {
}

func (d Directory) SetMomento(m Momento) {

}

func (d Directory) CreateMomento() *Momento {
	return nil
}

type momentoDirectory struct {
}

func (md momentoDirectory) GetState() {

}

func (md momentoDirectory) SetState() {

}

func main() {

}
