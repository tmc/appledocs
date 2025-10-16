
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [minFullScreenContentSize] class.
var minFullScreenContentSizeClass _minFullScreenContentSizeClass

func init() {
	minFullScreenContentSizeClass = _minFullScreenContentSizeClass{objc.GetClass("minFullScreenContentSize")}
}

type _minFullScreenContentSizeClass struct {
	objc.Class
}

// An interface definition for the [minFullScreenContentSize] class.
type IminFullScreenContentSize interface {
	ID() objc.ID
}

type minFullScreenContentSize struct {
	id objc.ID
}

func minFullScreenContentSizeFrom(ptr unsafe.Pointer) minFullScreenContentSize {
	return minFullScreenContentSize{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (m_ minFullScreenContentSize) ID() objc.ID {
	return m_.id
}

// Alloc allocates a new instance without initialization.
func (mc _minFullScreenContentSizeClass) Alloc() minFullScreenContentSize {
	rv := objc.Send[minFullScreenContentSize](objc.ID(mc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (mc _minFullScreenContentSizeClass) New() minFullScreenContentSize {
	rv := objc.Send[minFullScreenContentSize](objc.ID(mc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewminFullScreenContentSize creates and returns a new initialized instance.
func NewminFullScreenContentSize() minFullScreenContentSize {
	return minFullScreenContentSizeClass.New()
}

// Init initializes the instance.
func (m_ minFullScreenContentSize) Init() minFullScreenContentSize {
	rv := objc.Send[minFullScreenContentSize](m_.ID(), selInit)
	return rv
}
