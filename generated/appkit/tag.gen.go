
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [tag] class.
var tagClass _tagClass

func init() {
	tagClass = _tagClass{objc.GetClass("tag")}
}

type _tagClass struct {
	objc.Class
}

// An interface definition for the [tag] class.
type Itag interface {
	ID() objc.ID
}

type tag struct {
	id objc.ID
}

func tagFrom(ptr unsafe.Pointer) tag {
	return tag{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ tag) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _tagClass) Alloc() tag {
	rv := objc.Send[tag](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _tagClass) New() tag {
	rv := objc.Send[tag](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtag creates and returns a new initialized instance.
func Newtag() tag {
	return tagClass.New()
}

// Init initializes the instance.
func (t_ tag) Init() tag {
	rv := objc.Send[tag](t_.ID(), selInit)
	return rv
}
