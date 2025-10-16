
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tagNames] class.
var tagNamesClass _tagNamesClass

func init() {
	tagNamesClass = _tagNamesClass{objc.GetClass("tagNames")}
}

type _tagNamesClass struct {
	objc.Class
}

// An interface definition for the [tagNames] class.
type ItagNames interface {
	ID() objc.ID
}

type tagNames struct {
	id objc.ID
}

func tagNamesFrom(ptr unsafe.Pointer) tagNames {
	return tagNames{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tagNames) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tagNamesClass) Alloc() tagNames {
	rv := objc.Send[tagNames](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tagNamesClass) New() tagNames {
	rv := objc.Send[tagNames](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtagNames creates and returns a new initialized instance.
func NewtagNames() tagNames {
	return tagNamesClass.New()
}

// Init initializes the instance.
func (t_ tagNames) Init() tagNames {
	rv := objc.Send[tagNames](t_.ID(), selInit)
	return rv
}
