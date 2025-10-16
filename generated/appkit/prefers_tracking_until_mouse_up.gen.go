
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prefersTrackingUntilMouseUp] class.
var prefersTrackingUntilMouseUpClass _prefersTrackingUntilMouseUpClass

func init() {
	prefersTrackingUntilMouseUpClass = _prefersTrackingUntilMouseUpClass{objc.GetClass("prefersTrackingUntilMouseUp")}
}

type _prefersTrackingUntilMouseUpClass struct {
	objc.Class
}

// An interface definition for the [prefersTrackingUntilMouseUp] class.
type IprefersTrackingUntilMouseUp interface {
	ID() objc.ID
}

type prefersTrackingUntilMouseUp struct {
	id objc.ID
}

func prefersTrackingUntilMouseUpFrom(ptr unsafe.Pointer) prefersTrackingUntilMouseUp {
	return prefersTrackingUntilMouseUp{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prefersTrackingUntilMouseUp) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prefersTrackingUntilMouseUpClass) Alloc() prefersTrackingUntilMouseUp {
	rv := objc.Send[prefersTrackingUntilMouseUp](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prefersTrackingUntilMouseUpClass) New() prefersTrackingUntilMouseUp {
	rv := objc.Send[prefersTrackingUntilMouseUp](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprefersTrackingUntilMouseUp creates and returns a new initialized instance.
func NewprefersTrackingUntilMouseUp() prefersTrackingUntilMouseUp {
	return prefersTrackingUntilMouseUpClass.New()
}

// Init initializes the instance.
func (p_ prefersTrackingUntilMouseUp) Init() prefersTrackingUntilMouseUp {
	rv := objc.Send[prefersTrackingUntilMouseUp](p_.ID(), selInit)
	return rv
}
