
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BezierPath] class.
var BezierPathClass _BezierPathClass

func init() {
	BezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}
}

type _BezierPathClass struct {
	objc.Class
}

// An interface definition for the [BezierPath] class.
type IBezierPath interface {
	ID() objc.ID
}

type BezierPath struct {
	id objc.ID
}

func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ BezierPath) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BezierPathClass) New() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBezierPath creates and returns a new initialized instance.
func NewBezierPath() BezierPath {
	return BezierPathClass.New()
}

// Init initializes the instance.
func (b_ BezierPath) Init() BezierPath {
	rv := objc.Send[BezierPath](b_.ID(), selInit)
	return rv
}
