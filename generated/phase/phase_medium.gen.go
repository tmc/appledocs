// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PHASEMedium] class.
type IPHASEMedium interface {
	objectivec.IObject
}

// A property or quality of the environment that affects how sound travels.
//
// This class defines choices for the engine’s . Currently, this property provides only sound traveling through air.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEMediumClass) Alloc() PHASEMedium {
	rv := objc.Send[PHASEMedium](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a medium.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEMedium/init(engine:preset:)
func NewPHASEMediumWithEnginePreset(engine IPHASEEngine, preset IPHASEMediumPreset) PHASEMedium {
	instance := getPHASEMediumClass().Alloc()
	rv := objc.Send[PHASEMedium](instance.ID, objc.Sel("initWithEngine:preset:"), engine, preset)
	rv.Autorelease()
	return rv
}


// The physical matter through which sound travels.
//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseengine/defaultmedium
func (p_ PHASEMedium) DefaultMedium() PHASEMedium {
	rv := objc.Send[PHASEMedium](p_.ID, objc.Sel("defaultMedium"))
	return rv
}


// SetDefaultMedium sets the value of the defaultMedium property.
// The physical matter through which sound travels.

//
// [Full Topic]: https://developer.apple.com/documentation/phase/phaseengine/defaultmedium
func (p_ PHASEMedium) SetDefaultMedium(value IPHASEMedium) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDefaultMedium:"), value)
}


