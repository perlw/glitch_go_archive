package eap

import "errors"

type EntityManager struct {
	Entities    []Entity
	Assemblages map[Entity]map[string][]*Aspect

	Aspects map[string][]*Aspect
}

func NewEntityManager() *EntityManager {
	em := &EntityManager{}

	em.Assemblages = make(map[Entity]map[string][]*Aspect)
	em.Aspects = make(map[string][]*Aspect)

	return em
}

func (em *EntityManager) CreateEntity() Entity {
	return 0
}

func (em *EntityManager) KillEntity(entity Entity) {
}

func (em *EntityManager) AddAspect(entity Entity, aspect Aspect) error {
	var assemblage map[string][]*Aspect
	var ok bool

	if assemblage, ok = em.Assemblages[entity]; !ok {
		return errors.New("foo")
	}

	aspectType := aspect.GetType()
	if _, ok := em.Aspects[aspectType]; !ok {
		em.Aspects[aspectType] = make([]*Aspect, 10, 100)
	}
	em.Aspects[aspectType] = append(em.Aspects[aspectType], &aspect)

	if _, ok := assemblage[aspectType]; !ok {
		assemblage[aspectType] = make([]*Aspect, 10, 100)
	}
	assemblage[aspectType] = append(assemblage[aspectType], &aspect)

	return nil
}

func (em EntityManager) GetAspect(entity Entity, aspectType string) (*Aspect, error) {
	return nil, nil
}
