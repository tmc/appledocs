// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEShapeElement] class.
var (
	PHASEShapeElementClass     _PHASEShapeElementClass
	PHASEShapeElementClassOnce sync.Once
)

func getPHASEShapeElementClass() _PHASEShapeElementClass {
	PHASEShapeElementClassOnce.Do(func() {
		PHASEShapeElementClass = _PHASEShapeElementClass{objc.GetClass("PHASEShapeElement")}
	})
	return PHASEShapeElementClass
}

type _PHASEShapeElementClass struct {
	class objc.Class
}

// An interface definition for the [PHASEShapeElement] class.
type IPHASEShapeElement interface {
	objectivec.IObject
}

// An object that describes the characteristics of a physical surface.
//
// This class defines the material that makes up a object. You don’t instantiate instances of this class yourself; the framework creates an instance of this class for every material you pass into the initializer. The shape’s array provides read-only access to the instances.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/Element
type PHASEShapeElement struct {
	objectivec.Object
}

// PHASEShapeElementFrom constructs a [PHASEShapeElement] from an unsafe.Pointer.
//
// An object that describes the characteristics of a physical surface.
func PHASEShapeElementFrom(ptr unsafe.Pointer) PHASEShapeElement {
	return PHASEShapeElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEShapeElementClass) Alloc() PHASEShapeElement {
	rv := objc.Send[PHASEShapeElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEShapeElementClass) New() PHASEShapeElement {
	rv := objc.Send[PHASEShapeElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEShapeElement) Init() PHASEShapeElement {
	rv := objc.Send[PHASEShapeElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEShapeElement) Autorelease() PHASEShapeElement {
	rv := objc.Send[PHASEShapeElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEShapeElement creates a new PHASEShapeElement instance.
func NewPHASEShapeElement() PHASEShapeElement {
	return getPHASEShapeElementClass().New()
}


// A surface characteristic that determines the acoustic properties of an object.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/Element/material
func (p_ PHASEShapeElement) Material() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("material"))
	return rv
}


// SetMaterial sets the value of the material property.
// A surface characteristic that determines the acoustic properties of an object.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/Element/material
func (p_ PHASEShapeElement) SetMaterial(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaterial:"), value)
}

// An array of objects that collectively describe the physical characteristics of a surface.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseshape/elements
func (p_ PHASEShapeElement) Elements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("elements"))
	return rv
}


// SetElements sets the value of the elements property.
// An array of objects that collectively describe the physical characteristics of a surface.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseshape/elements
func (p_ PHASEShapeElement) SetElements(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setElements:"), value)
}



