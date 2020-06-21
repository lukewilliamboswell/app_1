package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("thing", func() {
	Title("thingulator Service")
	Description("Service for adding numbers, a Goa teaser")
	Server("thing", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var Thing = Type("Thing", func() {
	Description("Thing describes my thing that is stored in the DB.")
	Attribute("name", String, "Name of Thing", func() {
		MaxLength(100)
		Example("Apple")
	})
	Attribute("features", ArrayOf(String), "Features of Thing", func() {
	})
	Required("name")
})

var ThingResult = ResultType("vnd.goa.thing", func() {
	Reference(Thing)
	TypeName("StoredThing")

	Attributes(func() {
		Attribute("id", UInt64)
		Attribute("name", String)
		Attribute("features", ArrayOf(String))
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})

	Required("id", "name")
})

var _ = Service("thing", func() {
	Description("Gets all the things out of the DB")

	HTTP(func() {
		Path("/things")
	})

	Method("list", func() {
		Result(CollectionOf(ThingResult), func() {
			View("tiny")
		})
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

})
