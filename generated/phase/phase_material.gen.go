// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEMaterial */


/* debug [class_header]: Header for PHASEMaterial */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMaterial */
// An interface definition for the [PHASEMaterial] class.
type IPHASEMaterial interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEMaterial */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMaterial */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMaterial */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMaterialClass) Alloc() PHASEMaterial {
	rv := objc.Send[PHASEMaterial](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMaterial */
// Surface characteristics that determine the acoustic properties of an object.
//
// To specify the physical texture of a sound source or occluder, define the argument of the initializer, . The contains the surface types with which you define the argument of this class’s initializer.


// Surface characteristics that determine the acoustic properties of an object.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMaterial */

// Creates a material with the given preset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMaterial/init(engine:preset:)
func NewPHASEMaterialWithEnginePreset(engine IPHASEEngine, preset PHASEMaterialPreset) PHASEMaterial {
	instance := getPHASEMaterialClass().Alloc()
	rv := objc.Send[PHASEMaterial](instance.ID, objc.Sel("initWithEngine:preset:"), engine, preset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEMaterialWithEnginePreset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMaterial */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMaterial */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMaterial */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMaterial */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMaterial */


