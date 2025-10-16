
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [adjustSubviews] class.
var adjustSubviewsClass _adjustSubviewsClass

func init() {
	adjustSubviewsClass = _adjustSubviewsClass{objc.GetClass("adjustSubviews")}
}

type _adjustSubviewsClass struct {
	objc.Class
}

// An interface definition for the [adjustSubviews] class.
type IadjustSubviews interface {
	ID() objc.ID
}

type adjustSubviews struct {
	id objc.ID
}

func adjustSubviewsFrom(ptr unsafe.Pointer) adjustSubviews {
	return adjustSubviews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ adjustSubviews) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _adjustSubviewsClass) Alloc() adjustSubviews {
	rv := objc.Send[adjustSubviews](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _adjustSubviewsClass) New() adjustSubviews {
	rv := objc.Send[adjustSubviews](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewadjustSubviews creates and returns a new initialized instance.
func NewadjustSubviews() adjustSubviews {
	return adjustSubviewsClass.New()
}

// Init initializes the instance.
func (a_ adjustSubviews) Init() adjustSubviews {
	rv := objc.Send[adjustSubviews](a_.ID(), selInit)
	return rv
}
