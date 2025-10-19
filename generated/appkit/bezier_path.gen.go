// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BezierPath] class.
var (
	bezierPathClass     _BezierPathClass
	bezierPathClassOnce sync.Once
)

func getBezierPathClass() _BezierPathClass {
	bezierPathClassOnce.Do(func() {
		bezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}
	})
	return bezierPathClass
}

type _BezierPathClass struct {
	class objc.Class
}

// An interface definition for the [BezierPath] class.
type IBezierPath interface {
	objectivec.IObject
}

// An object that can create paths using PostScript-style commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath

type BezierPath struct {
	objectivec.Object
}

// BezierPathFrom constructs a [BezierPath] from an unsafe.Pointer.
//
// An object that can create paths using PostScript-style commands.
func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (bc _BezierPathClass) Alloc() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (bc _BezierPathClass) New() BezierPath {
	rv := objc.Send[BezierPath](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BezierPath) Init() BezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BezierPath) Autorelease() BezierPath {
	rv := objc.Send[BezierPath](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBezierPath creates a new BezierPath instance.
func NewBezierPath() BezierPath {
	return getBezierPathClass().New()
}




