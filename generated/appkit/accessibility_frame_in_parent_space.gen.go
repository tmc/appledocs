
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [accessibilityFrameInParentSpace] class.
var accessibilityFrameInParentSpaceClass _accessibilityFrameInParentSpaceClass

func init() {
	accessibilityFrameInParentSpaceClass = _accessibilityFrameInParentSpaceClass{objc.GetClass("accessibilityFrameInParentSpace")}
}

type _accessibilityFrameInParentSpaceClass struct {
	objc.Class
}

// An interface definition for the [accessibilityFrameInParentSpace] class.
type IaccessibilityFrameInParentSpace interface {
	ID() objc.ID
}

type accessibilityFrameInParentSpace struct {
	id objc.ID
}

func accessibilityFrameInParentSpaceFrom(ptr unsafe.Pointer) accessibilityFrameInParentSpace {
	return accessibilityFrameInParentSpace{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ accessibilityFrameInParentSpace) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _accessibilityFrameInParentSpaceClass) Alloc() accessibilityFrameInParentSpace {
	rv := objc.Send[accessibilityFrameInParentSpace](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _accessibilityFrameInParentSpaceClass) New() accessibilityFrameInParentSpace {
	rv := objc.Send[accessibilityFrameInParentSpace](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewaccessibilityFrameInParentSpace creates and returns a new initialized instance.
func NewaccessibilityFrameInParentSpace() accessibilityFrameInParentSpace {
	return accessibilityFrameInParentSpaceClass.New()
}

// Init initializes the instance.
func (a_ accessibilityFrameInParentSpace) Init() accessibilityFrameInParentSpace {
	rv := objc.Send[accessibilityFrameInParentSpace](a_.ID(), selInit)
	return rv
}
