
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [level] class.
var levelClass _levelClass

func init() {
	levelClass = _levelClass{objc.GetClass("level")}
}

type _levelClass struct {
	objc.Class
}

// An interface definition for the [level] class.
type Ilevel interface {
	ID() objc.ID
}

type level struct {
	id objc.ID
}

func levelFrom(ptr unsafe.Pointer) level {
	return level{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ level) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _levelClass) Alloc() level {
	rv := objc.Send[level](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _levelClass) New() level {
	rv := objc.Send[level](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newlevel creates and returns a new initialized instance.
func Newlevel() level {
	return levelClass.New()
}

// Init initializes the instance.
func (l_ level) Init() level {
	rv := objc.Send[level](l_.ID(), selInit)
	return rv
}
