
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [pulldownTarget] class.
var pulldownTargetClass _pulldownTargetClass

func init() {
	pulldownTargetClass = _pulldownTargetClass{objc.GetClass("pulldownTarget")}
}

type _pulldownTargetClass struct {
	objc.Class
}

// An interface definition for the [pulldownTarget] class.
type IpulldownTarget interface {
	ID() objc.ID
}

type pulldownTarget struct {
	id objc.ID
}

func pulldownTargetFrom(ptr unsafe.Pointer) pulldownTarget {
	return pulldownTarget{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ pulldownTarget) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _pulldownTargetClass) Alloc() pulldownTarget {
	rv := objc.Send[pulldownTarget](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _pulldownTargetClass) New() pulldownTarget {
	rv := objc.Send[pulldownTarget](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpulldownTarget creates and returns a new initialized instance.
func NewpulldownTarget() pulldownTarget {
	return pulldownTargetClass.New()
}

// Init initializes the instance.
func (p_ pulldownTarget) Init() pulldownTarget {
	rv := objc.Send[pulldownTarget](p_.ID(), selInit)
	return rv
}
