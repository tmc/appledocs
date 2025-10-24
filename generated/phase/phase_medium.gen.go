// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEMedium */


/* debug [class_header]: Header for PHASEMedium */
// The class instance for the [PHASEMedium] class.
var (
	PHASEMediumClass     _PHASEMediumClass
	PHASEMediumClassOnce sync.Once
)

func getPHASEMediumClass() _PHASEMediumClass {
	PHASEMediumClassOnce.Do(func() {
		PHASEMediumClass = _PHASEMediumClass{objc.GetClass("PHASEMedium")}
	})
	return PHASEMediumClass
}

type _PHASEMediumClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEMedium */
// An interface definition for the [PHASEMedium] class.
type IPHASEMedium interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEMedium */
	// properties:
	DefaultMedium() IPHASEMedium
	SetDefaultMedium(value IPHASEMedium)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEMedium */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEMedium */
// Alloc allocates a new instance without initialization.
func (pc _PHASEMediumClass) Alloc() PHASEMedium {
	rv := objc.Send[PHASEMedium](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEMediumClass) New() PHASEMedium {
	rv := objc.Send[PHASEMedium](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEMedium) Init() PHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEMedium) Autorelease() PHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEMedium creates a new PHASEMedium instance.
func NewPHASEMedium() PHASEMedium {
	return getPHASEMediumClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEMedium */
// A property or quality of the environment that affects how sound travels.
//
// This class defines choices for the engine’s . Currently, this property provides only sound traveling through air.


// A property or quality of the environment that affects how sound travels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMedium
type PHASEMedium struct {
	objectivec.Object
}

// PHASEMediumFrom constructs a [PHASEMedium] from an unsafe.Pointer.
//
// A property or quality of the environment that affects how sound travels.
func PHASEMediumFrom(ptr unsafe.Pointer) PHASEMedium {
	return PHASEMedium{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEMedium */

// Creates a medium.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMedium/init(engine:preset:)
func NewPHASEMediumWithEnginePreset(engine IPHASEEngine, preset PHASEMediumPreset) PHASEMedium {
	instance := getPHASEMediumClass().Alloc()
	rv := objc.Send[PHASEMedium](instance.ID, objc.Sel("initWithEngine:preset:"), engine, preset)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEMediumWithEnginePreset */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEMedium */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEMedium */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEMedium */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEMedium */

// The physical matter through which sound travels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseengine/defaultmedium
func (p_ PHASEMedium) DefaultMedium() IPHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("defaultMedium"))
	return rv
}/* debug [instance_properties/getter]: defaultMedium */


// The physical matter through which sound travels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseengine/defaultmedium
func (p_ PHASEMedium) SetDefaultMedium(value IPHASEMedium) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultMedium:"), value)
}/* debug [instance_properties/setter]: defaultMedium */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEMedium */


