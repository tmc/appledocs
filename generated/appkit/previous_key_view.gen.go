
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [previousKeyView] class.
var previousKeyViewClass _previousKeyViewClass

func init() {
	previousKeyViewClass = _previousKeyViewClass{objc.GetClass("previousKeyView")}
}

type _previousKeyViewClass struct {
	objc.Class
}

// An interface definition for the [previousKeyView] class.
type IpreviousKeyView interface {
	ID() objc.ID
}

type previousKeyView struct {
	id objc.ID
}

func previousKeyViewFrom(ptr unsafe.Pointer) previousKeyView {
	return previousKeyView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ previousKeyView) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _previousKeyViewClass) Alloc() previousKeyView {
	rv := objc.Send[previousKeyView](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _previousKeyViewClass) New() previousKeyView {
	rv := objc.Send[previousKeyView](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreviousKeyView creates and returns a new initialized instance.
func NewpreviousKeyView() previousKeyView {
	return previousKeyViewClass.New()
}

// Init initializes the instance.
func (p_ previousKeyView) Init() previousKeyView {
	rv := objc.Send[previousKeyView](p_.ID(), selInit)
	return rv
}
