
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [arrangesAllSubviews] class.
var arrangesAllSubviewsClass _arrangesAllSubviewsClass

func init() {
	arrangesAllSubviewsClass = _arrangesAllSubviewsClass{objc.GetClass("arrangesAllSubviews")}
}

type _arrangesAllSubviewsClass struct {
	objc.Class
}

// An interface definition for the [arrangesAllSubviews] class.
type IarrangesAllSubviews interface {
	ID() objc.ID
}

type arrangesAllSubviews struct {
	id objc.ID
}

func arrangesAllSubviewsFrom(ptr unsafe.Pointer) arrangesAllSubviews {
	return arrangesAllSubviews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ arrangesAllSubviews) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _arrangesAllSubviewsClass) Alloc() arrangesAllSubviews {
	rv := objc.Send[arrangesAllSubviews](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _arrangesAllSubviewsClass) New() arrangesAllSubviews {
	rv := objc.Send[arrangesAllSubviews](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewarrangesAllSubviews creates and returns a new initialized instance.
func NewarrangesAllSubviews() arrangesAllSubviews {
	return arrangesAllSubviewsClass.New()
}

// Init initializes the instance.
func (a_ arrangesAllSubviews) Init() arrangesAllSubviews {
	rv := objc.Send[arrangesAllSubviews](a_.ID(), selInit)
	return rv
}
