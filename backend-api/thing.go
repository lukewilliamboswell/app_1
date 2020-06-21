package thingapi

import (
	thing "backend-api/gen/thing"
	"context"
	"log"
)

// thing service example implementation.
// The example methods log the requests and return zero values.
type thingsrvc struct {
	logger *log.Logger
}

// NewThing returns the thing service implementation.
func NewThing(logger *log.Logger) thing.Service {
	return &thingsrvc{logger}
}

// List implements list.
func (s *thingsrvc) List(ctx context.Context) (res thing.StoredThingCollection, err error) {
	s.logger.Print("thing.list")

	storedThing := thing.StoredThing{
		ID:       23,
		Name:     "Hello World",
		Features: []string{},
	}

	var storedThingCollection = []*thing.StoredThing{&storedThing}

	return thing.StoredThingCollection(storedThingCollection), nil
}
