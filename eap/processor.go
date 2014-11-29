package eap

type Processor interface {
	Tick(*EntityManager, int)
}
