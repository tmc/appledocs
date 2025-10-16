
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [arrangedSubviews] class.
var arrangedSubviewsClass _arrangedSubviewsClass

func init() {
	arrangedSubviewsClass = _arrangedSubviewsClass{objc.GetClass("arrangedSubviews")}
}

type _arrangedSubviewsClass struct {
	objc.Class
}

// An interface definition for the [arrangedSubviews] class.
type IarrangedSubviews interface {
	ID() objc.ID
}

type arrangedSubviews struct {
	id objc.ID
}

func arrangedSubviewsFrom(ptr unsafe.Pointer) arrangedSubviews {
	return arrangedSubviews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ arrangedSubviews) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _arrangedSubviewsClass) Alloc() arrangedSubviews {
	rv := objc.Send[arrangedSubviews](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _arrangedSubviewsClass) New() arrangedSubviews {
	rv := objc.Send[arrangedSubviews](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewarrangedSubviews creates and returns a new initialized instance.
func NewarrangedSubviews() arrangedSubviews {
	return arrangedSubviewsClass.New()
}

// Init initializes the instance.
func (a_ arrangedSubviews) Init() arrangedSubviews {
	rv := objc.Send[arrangedSubviews](a_.ID(), selInit)
	return rv
}
