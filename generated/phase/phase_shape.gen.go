// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEShape] class.
var (
	PHASEShapeClass     _PHASEShapeClass
	PHASEShapeClassOnce sync.Once
)

func getPHASEShapeClass() _PHASEShapeClass {
	PHASEShapeClassOnce.Do(func() {
		PHASEShapeClass = _PHASEShapeClass{objc.GetClass("PHASEShape")}
	})
	return PHASEShapeClass
}

type _PHASEShapeClass struct {
	class objc.Class
}

// An interface definition for the [PHASEShape] class.
type IPHASEShape interface {
	objectivec.IObject
	// properties:
	Elements() []IPHASEShapeElement
	Shapes() IPHASEShape
	SetShapes(value IPHASEShape)
	// methods:
}

// A collection of points that connect to form a 3D volume.
//
// To define your scene’s important 3D volumes, create one or more of the following surfaces and add them to your scene’s array: The audio-emitting surface of a volumetric The audio-deflecting surface and texture of a


// A collection of points that connect to form a 3D volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape
type PHASEShape struct {
	objectivec.Object
}

// PHASEShapeFrom constructs a [PHASEShape] from an unsafe.Pointer.
//
// A collection of points that connect to form a 3D volume.
func PHASEShapeFrom(ptr unsafe.Pointer) PHASEShape {
	return PHASEShape{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEShapeClass) Alloc() PHASEShape {
	rv := objc.Send[PHASEShape](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEShapeClass) New() PHASEShape {
	rv := objc.Send[PHASEShape](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEShape) Init() PHASEShape {
	rv := objc.Send[PHASEShape](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEShape) Autorelease() PHASEShape {
	rv := objc.Send[PHASEShape](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEShape creates a new PHASEShape instance.
func NewPHASEShape() PHASEShape {
	return getPHASEShapeClass().New()
}



// Creates an object that the given geometric data shapes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/init(engine:mesh:)
func NewPHASEShapeWithEngineMesh(engine IPHASEEngine, mesh unsafe.Pointer) PHASEShape {
	instance := getPHASEShapeClass().Alloc()
	rv := objc.Send[PHASEShape](instance.ID, objc.Sel("initWithEngine:mesh:"), engine, mesh)
	rv.Autorelease()
	return rv
}


// Creates an object of a specific material that the given geometric data shapes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/init(engine:mesh:materials:)
func NewPHASEShapeWithEngineMeshMaterials(engine IPHASEEngine, mesh unsafe.Pointer, materials []IPHASEMaterial) PHASEShape {
	instance := getPHASEShapeClass().Alloc()
	rv := objc.Send[PHASEShape](instance.ID, objc.Sel("initWithEngine:mesh:materials:"), engine, mesh, materials)
	rv.Autorelease()
	return rv
}



// An array of objects that collectively describe the physical characteristics of a surface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEShape/elements
func (p_ PHASEShape) Elements() []IPHASEShapeElement {
	rv := objc.Send[[]PHASEShapeElement](p_.ID, objc.Sel("elements"))
	return rv
}


// An array of shapes that collectively define the audio-emitting surface area of a volumetric source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesource/shapes
func (p_ PHASEShape) Shapes() IPHASEShape {
	rv := objc.Send[PHASEShape](p_.ID, objc.Sel("shapes"))
	return rv
}


// An array of shapes that collectively define the audio-emitting surface area of a volumetric source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasesource/shapes
func (p_ PHASEShape) SetShapes(value IPHASEShape) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShapes:"), value)
}


