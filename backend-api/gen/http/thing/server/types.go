// Code generated by goa v3.1.2, DO NOT EDIT.
//
// thing HTTP server types
//
// Command:
// $ goa gen backend-api/design

package server

import (
	thingviews "backend-api/gen/thing/views"
)

// StoredThingResponseTinyCollection is the type of the "thing" service "list"
// endpoint HTTP response body.
type StoredThingResponseTinyCollection []*StoredThingResponseTiny

// StoredThingResponseTiny is used to define fields on response body types.
type StoredThingResponseTiny struct {
	ID uint64 `form:"id" json:"id" xml:"id"`
	// Name of Thing
	Name string `form:"name" json:"name" xml:"name"`
}

// NewStoredThingResponseTinyCollection builds the HTTP response body from the
// result of the "list" endpoint of the "thing" service.
func NewStoredThingResponseTinyCollection(res thingviews.StoredThingCollectionView) StoredThingResponseTinyCollection {
	body := make([]*StoredThingResponseTiny, len(res))
	for i, val := range res {
		body[i] = marshalThingviewsStoredThingViewToStoredThingResponseTiny(val)
	}
	return body
}
