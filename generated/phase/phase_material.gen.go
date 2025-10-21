// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHASEMaterial] class.
var (
	PHASEMaterialClass     _PHASEMaterialClass
	PHASEMaterialClassOnce sync.Once
)

func getPHASEMaterialClass() _PHASEMaterialClass {
	PHASEMaterialClassOnce.Do(func() {
		PHASEMaterialClass = _PHASEMaterialClass{objc.GetClass("PHASEMaterial")}
	})
	return PHASEMaterialClass
}

type _PHASEMaterialClass struct {
	class objc.Class
}

// An interface definition for the [PHASEMaterial] class.
type IPHASEMaterial interface {
	objectivec.IObject
}

// Surface characteristics that determine the acoustic properties of an object.
//
// To specify the physical texture of a sound source or occluder, define the argument of the initializer, . The contains the surface types with which you define the argument of this class’s initializer.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterial
type PHASEMaterial struct {
	objectivec.Object
}

// PHASEMaterialFrom constructs a [PHASEMaterial] from an unsafe.Pointer.
//
// Surface characteristics that determine the acoustic properties of an object.
func PHASEMaterialFrom(ptr unsafe.Pointer) PHASEMaterial {
	return PHASEMaterial{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEMaterialClass) Alloc() PHASEMaterial {
	rv := objc.Send[PHASEMaterial](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEMaterialClass) New() PHASEMaterial {
	rv := objc.Send[PHASEMaterial](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMaterial) Init() PHASEMaterial {
	rv := objc.Send[PHASEMaterial](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMaterial) Autorelease() PHASEMaterial {
	rv := objc.Send[PHASEMaterial](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMaterial creates a new PHASEMaterial instance.
func NewPHASEMaterial() PHASEMaterial {
	return getPHASEMaterialClass().New()
}


// Creates a material with the given preset.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterial/init(engine:preset:)
func NewPHASEMaterialWithEnginePreset(engine unsafe.Pointer, preset unsafe.Pointer) PHASEMaterial {
	instance := getPHASEMaterialClass().Alloc()
	rv := objc.Send[PHASEMaterial](instance.ID, objc.Sel("initWithEngine:preset:"), engine, preset)
	rv.Autorelease()
	return rv
}



