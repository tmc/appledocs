
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [effectsViewInserter] class.
var effectsViewInserterClass _effectsViewInserterClass

func init() {
	effectsViewInserterClass = _effectsViewInserterClass{objc.GetClass("effectsViewInserter")}
}

type _effectsViewInserterClass struct {
	objc.Class
}

// An interface definition for the [effectsViewInserter] class.
type IeffectsViewInserter interface {
	ID() objc.ID
}

type effectsViewInserter struct {
	id objc.ID
}

func effectsViewInserterFrom(ptr unsafe.Pointer) effectsViewInserter {
	return effectsViewInserter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ effectsViewInserter) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _effectsViewInserterClass) Alloc() effectsViewInserter {
	rv := objc.Send[effectsViewInserter](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _effectsViewInserterClass) New() effectsViewInserter {
	rv := objc.Send[effectsViewInserter](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NeweffectsViewInserter creates and returns a new initialized instance.
func NeweffectsViewInserter() effectsViewInserter {
	return effectsViewInserterClass.New()
}

// Init initializes the instance.
func (e_ effectsViewInserter) Init() effectsViewInserter {
	rv := objc.Send[effectsViewInserter](e_.ID(), selInit)
	return rv
}
