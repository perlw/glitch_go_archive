package eap

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type EntityManager struct {
	entities   []Entity
	nextEntity Entity

	assemblages map[Entity]map[string]Aspect

	aspects map[string][]Aspect

	entityTemplates map[string][]Entity
}

func findFreeEntity(entities []Entity, nextEntity Entity) Entity {
	for index, entity := range entities[nextEntity-1:] {
		if entity == 0 {
			return Entity(index + 1)
		}
	}

	return Entity(len(entities) + 1)
}

func NewEntityManager() *EntityManager {
	em := &EntityManager{}
	em.entities = []Entity{}
	em.nextEntity = 1

	em.assemblages = make(map[Entity]map[string]Aspect)
	em.aspects = make(map[string][]Aspect)
	em.entityTemplates = make(map[string][]Entity)

	return em
}

func (em *EntityManager) CreateEntity() Entity {
	ent := em.nextEntity

	if int(ent-1) == len(em.entities) {
		em.entities = append(em.entities, ent)
	} else {
		em.entities[ent-1] = ent
	}
	em.nextEntity = findFreeEntity(em.entities, em.nextEntity)

	em.assemblages[ent-1] = make(map[string]Aspect)

	return ent
}

func (em *EntityManager) KillEntity(entity Entity) error {
	var assemblage map[string]Aspect
	var ok bool

	if assemblage, ok = em.assemblages[entity-1]; !ok {
		return errors.New(fmt.Sprintf("Entity #%d assemblage error", entity))
	}

	for index := range assemblage {
		aspects := em.aspects[index]
		for aindex, aspect := range aspects {
			if aspect == assemblage[index] {
				em.aspects[index] = append(em.aspects[index][:aindex], em.aspects[index][aindex+1:]...)
				break
			}
		}
		assemblage[index] = nil
	}
	em.assemblages[entity-1] = nil
	em.entities[entity-1] = 0

	em.nextEntity = entity

	return nil
}

func (em *EntityManager) AddAspect(entity Entity, aspect Aspect) error {
	var assemblage map[string]Aspect
	var ok bool

	if assemblage, ok = em.assemblages[entity-1]; !ok {
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

	aspect.SetEntity(entity)

	return nil
}

func (em EntityManager) GetAspect(entity Entity, aspectType string) (Aspect, error) {
	var aspect Aspect
	var aspects map[string]Aspect
	var ok bool

	if aspects, ok = em.assemblages[entity-1]; !ok {
		return nil, errors.New(fmt.Sprintf("Could not find assemblage for entity #%d", entity))
	}

	if aspect, ok = aspects[aspectType]; !ok {
		return nil, errors.New(fmt.Sprintf("No aspect of type %s bound to entity #%d", aspectType, entity))
	}

	return aspect, nil
}

func (em EntityManager) GetAspectsFromType(aspectType string) ([]Aspect, error) {
	if aspects, ok := em.aspects[aspectType]; ok {
		return aspects, nil
	}

	return nil, errors.New(fmt.Sprintf("Unknown aspect type %s", aspectType))
}

// TODO: Add in prebuilding templates, ids, etc
// TODO: Dirtychecking to rebuild list, need template struct
func (em *EntityManager) GetEntitiesFromAspects(aspectTypes []string) []Entity {
	var entities []Entity
	var ok bool

	sort.Strings(aspectTypes)
	name := strings.Join(aspectTypes, "#")

	if entities, ok = em.entityTemplates[name]; !ok {
		entities = []Entity{}

		allAspects := []Aspect{}

		for _, aspectType := range aspectTypes {
			if aspects, err := em.GetAspectsFromType(aspectType); err == nil {
				allAspects = append(allAspects, aspects...)
			}
		}

		numAspects := len(aspectTypes)
		entityCount := make(map[Entity]int)
		for _, aspect := range allAspects {
			var count int
			var ok bool

			entity := aspect.GetEntity()
			if count, ok = entityCount[entity]; !ok {
				count = 0
			}
			count++
			entityCount[entity] = count

			if count == numAspects {
				entities = append(entities, entity)
			}
		}

		em.entityTemplates[name] = entities
	}

	return entities
}
