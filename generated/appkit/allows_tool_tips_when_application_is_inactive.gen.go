
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsToolTipsWhenApplicationIsInactive] class.
var allowsToolTipsWhenApplicationIsInactiveClass _allowsToolTipsWhenApplicationIsInactiveClass

func init() {
	allowsToolTipsWhenApplicationIsInactiveClass = _allowsToolTipsWhenApplicationIsInactiveClass{objc.GetClass("allowsToolTipsWhenApplicationIsInactive")}
}

type _allowsToolTipsWhenApplicationIsInactiveClass struct {
	objc.Class
}

// An interface definition for the [allowsToolTipsWhenApplicationIsInactive] class.
type IallowsToolTipsWhenApplicationIsInactive interface {
	ID() objc.ID
}

type allowsToolTipsWhenApplicationIsInactive struct {
	id objc.ID
}

func allowsToolTipsWhenApplicationIsInactiveFrom(ptr unsafe.Pointer) allowsToolTipsWhenApplicationIsInactive {
	return allowsToolTipsWhenApplicationIsInactive{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsToolTipsWhenApplicationIsInactive) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsToolTipsWhenApplicationIsInactiveClass) Alloc() allowsToolTipsWhenApplicationIsInactive {
	rv := objc.Send[allowsToolTipsWhenApplicationIsInactive](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsToolTipsWhenApplicationIsInactiveClass) New() allowsToolTipsWhenApplicationIsInactive {
	rv := objc.Send[allowsToolTipsWhenApplicationIsInactive](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsToolTipsWhenApplicationIsInactive creates and returns a new initialized instance.
func NewallowsToolTipsWhenApplicationIsInactive() allowsToolTipsWhenApplicationIsInactive {
	return allowsToolTipsWhenApplicationIsInactiveClass.New()
}

// Init initializes the instance.
func (a_ allowsToolTipsWhenApplicationIsInactive) Init() allowsToolTipsWhenApplicationIsInactive {
	rv := objc.Send[allowsToolTipsWhenApplicationIsInactive](a_.ID(), selInit)
	return rv
}
