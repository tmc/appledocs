
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [hidesOnDeactivate] class.
var hidesOnDeactivateClass _hidesOnDeactivateClass

func init() {
	hidesOnDeactivateClass = _hidesOnDeactivateClass{objc.GetClass("hidesOnDeactivate")}
}

type _hidesOnDeactivateClass struct {
	objc.Class
}

// An interface definition for the [hidesOnDeactivate] class.
type IhidesOnDeactivate interface {
	ID() objc.ID
}

type hidesOnDeactivate struct {
	id objc.ID
}

func hidesOnDeactivateFrom(ptr unsafe.Pointer) hidesOnDeactivate {
	return hidesOnDeactivate{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ hidesOnDeactivate) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _hidesOnDeactivateClass) Alloc() hidesOnDeactivate {
	rv := objc.Send[hidesOnDeactivate](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _hidesOnDeactivateClass) New() hidesOnDeactivate {
	rv := objc.Send[hidesOnDeactivate](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewhidesOnDeactivate creates and returns a new initialized instance.
func NewhidesOnDeactivate() hidesOnDeactivate {
	return hidesOnDeactivateClass.New()
}

// Init initializes the instance.
func (h_ hidesOnDeactivate) Init() hidesOnDeactivate {
	rv := objc.Send[hidesOnDeactivate](h_.ID(), selInit)
	return rv
}
