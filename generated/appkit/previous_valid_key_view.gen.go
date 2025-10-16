
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [previousValidKeyView] class.
var previousValidKeyViewClass _previousValidKeyViewClass

func init() {
	previousValidKeyViewClass = _previousValidKeyViewClass{objc.GetClass("previousValidKeyView")}
}

type _previousValidKeyViewClass struct {
	objc.Class
}

// An interface definition for the [previousValidKeyView] class.
type IpreviousValidKeyView interface {
	ID() objc.ID
}

type previousValidKeyView struct {
	id objc.ID
}

func previousValidKeyViewFrom(ptr unsafe.Pointer) previousValidKeyView {
	return previousValidKeyView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ previousValidKeyView) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _previousValidKeyViewClass) Alloc() previousValidKeyView {
	rv := objc.Send[previousValidKeyView](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _previousValidKeyViewClass) New() previousValidKeyView {
	rv := objc.Send[previousValidKeyView](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewpreviousValidKeyView creates and returns a new initialized instance.
func NewpreviousValidKeyView() previousValidKeyView {
	return previousValidKeyViewClass.New()
}

// Init initializes the instance.
func (p_ previousValidKeyView) Init() previousValidKeyView {
	rv := objc.Send[previousValidKeyView](p_.ID(), selInit)
	return rv
}
