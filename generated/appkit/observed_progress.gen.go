
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [observedProgress] class.
var observedProgressClass _observedProgressClass

func init() {
	observedProgressClass = _observedProgressClass{objc.GetClass("observedProgress")}
}

type _observedProgressClass struct {
	objc.Class
}

// An interface definition for the [observedProgress] class.
type IobservedProgress interface {
	ID() objc.ID
}

type observedProgress struct {
	id objc.ID
}

func observedProgressFrom(ptr unsafe.Pointer) observedProgress {
	return observedProgress{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (o_ observedProgress) ID() objc.ID {
	return o_.id
}

// Alloc allocates a new instance without initialization.
func (oc _observedProgressClass) Alloc() observedProgress {
	rv := objc.Send[observedProgress](objc.ID(oc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (oc _observedProgressClass) New() observedProgress {
	rv := objc.Send[observedProgress](objc.ID(oc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewobservedProgress creates and returns a new initialized instance.
func NewobservedProgress() observedProgress {
	return observedProgressClass.New()
}

// Init initializes the instance.
func (o_ observedProgress) Init() observedProgress {
	rv := objc.Send[observedProgress](o_.ID(), selInit)
	return rv
}
