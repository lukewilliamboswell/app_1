// Code generated by goa v3.1.2, DO NOT EDIT.
//
// thing views
//
// Command:
// $ goa gen backend-api/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// StoredThingCollection is the viewed result type that is projected based on a
// view.
type StoredThingCollection struct {
	// Type to project
	Projected StoredThingCollectionView
	// View to render
	View string
}

// StoredThingCollectionView is a type that runs validations on a projected
// type.
type StoredThingCollectionView []*StoredThingView

// StoredThingView is a type that runs validations on a projected type.
type StoredThingView struct {
	ID *uint64
	// Name of Thing
	Name *string
	// Features of Thing
	Features []string
}

var (
	// StoredThingCollectionMap is a map of attribute names in result type
	// StoredThingCollection indexed by view name.
	StoredThingCollectionMap = map[string][]string{
		"tiny": []string{
			"id",
			"name",
		},
		"default": []string{
			"id",
			"name",
			"features",
		},
	}
	// StoredThingMap is a map of attribute names in result type StoredThing
	// indexed by view name.
	StoredThingMap = map[string][]string{
		"tiny": []string{
			"id",
			"name",
		},
		"default": []string{
			"id",
			"name",
			"features",
		},
	}
)

// ValidateStoredThingCollection runs the validations defined on the viewed
// result type StoredThingCollection.
func ValidateStoredThingCollection(result StoredThingCollection) (err error) {
	switch result.View {
	case "tiny":
		err = ValidateStoredThingCollectionViewTiny(result.Projected)
	case "default", "":
		err = ValidateStoredThingCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"tiny", "default"})
	}
	return
}

// ValidateStoredThingCollectionViewTiny runs the validations defined on
// StoredThingCollectionView using the "tiny" view.
func ValidateStoredThingCollectionViewTiny(result StoredThingCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredThingViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredThingCollectionView runs the validations defined on
// StoredThingCollectionView using the "default" view.
func ValidateStoredThingCollectionView(result StoredThingCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateStoredThingView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateStoredThingViewTiny runs the validations defined on StoredThingView
// using the "tiny" view.
func ValidateStoredThingViewTiny(result *StoredThingView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	return
}

// ValidateStoredThingView runs the validations defined on StoredThingView
// using the "default" view.
func ValidateStoredThingView(result *StoredThingView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Name != nil {
		if utf8.RuneCountInString(*result.Name) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.name", *result.Name, utf8.RuneCountInString(*result.Name), 100, false))
		}
	}
	return
}
