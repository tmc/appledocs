// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEStreamNode */


/* debug [class_header]: Header for PHASEStreamNode */
// The class instance for the [PHASEStreamNode] class.
var (
	PHASEStreamNodeClass     _PHASEStreamNodeClass
	PHASEStreamNodeClassOnce sync.Once
)

func getPHASEStreamNodeClass() _PHASEStreamNodeClass {
	PHASEStreamNodeClassOnce.Do(func() {
		PHASEStreamNodeClass = _PHASEStreamNodeClass{objc.GetClass("PHASEStreamNode")}
	})
	return PHASEStreamNodeClass
}

type _PHASEStreamNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEStreamNode */
// An interface definition for the [PHASEStreamNode] class.
type IPHASEStreamNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEStreamNode */
	// properties:
	Format() avfaudio.AudioFormat
	GainMetaParameter() IPHASENumberMetaParameter
	Mixer() IPHASEMixer
	RateMetaParameter() IPHASENumberMetaParameter
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEStreamNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEStreamNode */
// Alloc allocates a new instance without initialization.
func (pc _PHASEStreamNodeClass) Alloc() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEStreamNodeClass) New() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEStreamNode) Init() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEStreamNode) Autorelease() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEStreamNode creates a new PHASEStreamNode instance.
func NewPHASEStreamNode() PHASEStreamNode {
	return getPHASEStreamNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEStreamNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode
type PHASEStreamNode struct {
	objectivec.Object
}

// PHASEStreamNodeFrom constructs a [PHASEStreamNode] from an unsafe.Pointer.
func PHASEStreamNodeFrom(ptr unsafe.Pointer) PHASEStreamNode {
	return PHASEStreamNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEStreamNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEStreamNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEStreamNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEStreamNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEStreamNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/format
func (p_ PHASEStreamNode) Format() avfaudio.AudioFormat {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}/* debug [instance_properties/getter]: format */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/gainMetaParameter
func (p_ PHASEStreamNode) GainMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}/* debug [instance_properties/getter]: gainMetaParameter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/mixer
func (p_ PHASEStreamNode) Mixer() IPHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("mixer"))
	return rv
}/* debug [instance_properties/getter]: mixer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/rateMetaParameter
func (p_ PHASEStreamNode) RateMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("rateMetaParameter"))
	return rv
}/* debug [instance_properties/getter]: rateMetaParameter */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEStreamNode */



