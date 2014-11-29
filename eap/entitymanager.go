package eap

import (
	"errors"
	"fmt"
)

type EntityManager struct {
	entities    []Entity
	assemblages map[Entity]map[string]Aspect

	aspects map[string][]Aspect
}

func NewEntityManager() *EntityManager {
	em := &EntityManager{}

	em.assemblages = make(map[Entity]map[string]Aspect)
	em.aspects = make(map[string][]Aspect)

	return em
}

func (em *EntityManager) CreateEntity() Entity {
	ent := Entity(0)

	em.assemblages[ent] = make(map[string]Aspect)

	return ent
}

func (em *EntityManager) KillEntity(entity Entity) {
}

func (em *EntityManager) AddAspect(entity Entity, aspect Aspect) error {
	var assemblage map[string]Aspect
	var ok bool

	if assemblage, ok = em.assemblages[entity]; !ok {
		return errors.New(fmt.Sprintf("Entity #%d assemblage error", entity))
	}

	aspectType := aspect.GetType()
	if _, ok := em.aspects[aspectType]; !ok {
		em.aspects[aspectType] = make([]Aspect, 0, 100)
	}
	em.aspects[aspectType] = append(em.aspects[aspectType], aspect)

	if _, ok := assemblage[aspectType]; ok {
		return errors.New(fmt.Sprintf("Entity #%d already has a %s aspect", entity, aspectType))
	}
	assemblage[aspectType] = aspect

	return nil
}

func (em EntityManager) GetAspect(entity Entity, aspectType string) (Aspect, error) {
	return nil, nil
}

func (em EntityManager) GetAspectsFromType(aspectType string) ([]Aspect, error) {
	if aspects, ok := em.aspects[aspectType]; ok {
		return aspects, nil
	}

	return nil, errors.New("Unknown aspect type " + aspectType)
}
