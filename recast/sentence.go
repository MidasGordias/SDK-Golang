package recast

// Sentence is generated by the Recast.AI Response and provide utilities for managing the sentence data
type Sentence struct {
	source   string
	_type    string
	action   string
	agent    string
	polarity string
	entities map[string][]*Entity
}

func newSentence(data map[string]interface{}) *Sentence {
	s := &Sentence{}
	s.source, _ = data["source"].(string)
	s._type, _ = data["type"].(string)
	s.action, _ = data["action"].(string)
	s.agent, _ = data["agent"].(string)
	s.polarity, _ = data["polarity"].(string)
	entities := data["entities"].(map[string]interface{})
	s.entities = make(map[string][]*Entity)
	for name, ents := range entities {
		s.entities[name] = make([]*Entity, len(ents.([]interface{})))
		for i, ent := range ents.([]interface{}) {
			currentEntityData := ent.(map[string]interface{})
			s.entities[name][i] = &Entity{currentEntityData, name}
		}
	}
	return s
}

// Type returns the type of the sentence
// Refert ot the Recast.Ai manual
func (s *Sentence) Type() string {
	return s._type
}

// Source returns the source of the sentence
func (s *Sentence) Source() string {
	return s.source
}

// Action returns the action of the sentence
func (s *Sentence) Action() string {
	return s.action
}

// Agent returns the agent of the sentence
// Note that the value returned may be empty
func (s *Sentence) Agent() string {
	return s.agent
}

// Polarity returns the polarity of the sentence
// Refer to the Recast.Ai manual
func (s *Sentence) Polarity() string {
	return s.polarity
}

//Entity returns the first entity matching with name
func (s *Sentence) Entity(name string) *Entity {
	if entities, exists := s.entities[name]; exists {
		return entities[0]
	}
	return nil
}

// Entities returns a slice of Entity containing all entities matching with name
func (s *Sentence) Entities(name string) []*Entity {
	if ents, exists := s.entities[name]; exists {
		return ents
	}
	return nil

}

// AllEntities returns a map containing slices of entities matching names, with names as keys
// AllEntities called with no arguments returns a map of all entities detected in the sentence
func (s *Sentence) AllEntities(names ...string) map[string][]*Entity {
	if len(names) == 0 {
		return s.entities
	}
	entities := make(map[string][]*Entity)
	for _, name := range names {
		if ents, exists := s.entities[name]; exists {
			entities[name] = make([]*Entity, len(ents))
			copy(entities[name], ents)
		}
	}
	return entities
}
