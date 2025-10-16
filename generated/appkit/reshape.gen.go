
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [reshape] class.
var reshapeClass _reshapeClass

func init() {
	reshapeClass = _reshapeClass{objc.GetClass("reshape")}
}

type _reshapeClass struct {
	objc.Class
}

// An interface definition for the [reshape] class.
type Ireshape interface {
	ID() objc.ID
}

type reshape struct {
	id objc.ID
}

func reshapeFrom(ptr unsafe.Pointer) reshape {
	return reshape{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ reshape) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _reshapeClass) Alloc() reshape {
	rv := objc.Send[reshape](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _reshapeClass) New() reshape {
	rv := objc.Send[reshape](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newreshape creates and returns a new initialized instance.
func Newreshape() reshape {
	return reshapeClass.New()
}

// Init initializes the instance.
func (r_ reshape) Init() reshape {
	rv := objc.Send[reshape](r_.ID(), selInit)
	return rv
}
