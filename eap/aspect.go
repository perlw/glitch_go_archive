package eap

type Aspect interface {
	GetType() string
	GetEntity() Entity
	SetEntity(Entity)
}
