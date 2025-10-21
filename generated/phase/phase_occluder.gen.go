// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEOccluder] class.
var (
	PHASEOccluderClass     _PHASEOccluderClass
	PHASEOccluderClassOnce sync.Once
)

func getPHASEOccluderClass() _PHASEOccluderClass {
	PHASEOccluderClassOnce.Do(func() {
		PHASEOccluderClass = _PHASEOccluderClass{objc.GetClass("PHASEOccluder")}
	})
	return PHASEOccluderClass
}

type _PHASEOccluderClass struct {
	class objc.Class
}

// An interface definition for the [PHASEOccluder] class.
type IPHASEOccluder interface {
	IPHASEObject
}

// An object with a shape and position that blocks audio from reaching the listener.
//
// The framework lowers the volume of an audio signal when an instance of this class positions somewhere along the path between the sound source and the listener. For an example that demonstrates sound occlusion, see .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder
type PHASEOccluder struct {
	PHASEObject
}

// PHASEOccluderFrom constructs a [PHASEOccluder] from an unsafe.Pointer.
//
// An object with a shape and position that blocks audio from reaching the listener.
func PHASEOccluderFrom(ptr unsafe.Pointer) PHASEOccluder {
	return PHASEOccluder{
		PHASEObject: PHASEObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEOccluderClass) Alloc() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEOccluderClass) New() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEOccluder) Init() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEOccluder) Autorelease() PHASEOccluder {
	rv := objc.Send[PHASEOccluder](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEOccluder creates a new PHASEOccluder instance.
func NewPHASEOccluder() PHASEOccluder {
	return getPHASEOccluderClass().New()
}


// Creates an occluder with the given engine and shapes.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder/init(engine:shapes:)
func NewPHASEOccluderWithEngineShapes(engine unsafe.Pointer, shapes unsafe.Pointer) PHASEOccluder {
	instance := getPHASEOccluderClass().Alloc()
	rv := objc.Send[PHASEOccluder](instance.ID, objc.Sel("initWithEngine:shapes:"), engine, shapes)
	rv.Autorelease()
	return rv
}


// An array of shapes that collectively define the occluder’s audio-deflecting surface and texture.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEOccluder/shapes
func (p_ PHASEOccluder) Shapes() []PHASEShape {
	rv := objc.Send[[]PHASEShape](p_.ID, objc.Sel("shapes"))
	return rv
}


