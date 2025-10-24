// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LADomainState */


/* debug [class_header]: Header for LADomainState */
// The class instance for the [DomainState] class.
var (
	DomainStateClass     _DomainStateClass
	DomainStateClassOnce sync.Once
)

func getDomainStateClass() _DomainStateClass {
	DomainStateClassOnce.Do(func() {
		DomainStateClass = _DomainStateClass{objc.GetClass("LADomainState")}
	})
	return DomainStateClass
}

type _DomainStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DomainState */
// An interface definition for the [DomainState] class.
type IDomainState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DomainState */
	// properties:
	Biometry() ILADomainStateBiometry
	Companion() ILADomainStateCompanion
	StateHash() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DomainState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DomainState */
// Alloc allocates a new instance without initialization.
func (dc _DomainStateClass) Alloc() DomainState {
	rv := objc.Send[DomainState](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DomainStateClass) New() DomainState {
	rv := objc.Send[DomainState](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainState) Init() DomainState {
	rv := objc.Send[DomainState](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainState) Autorelease() DomainState {
	rv := objc.Send[DomainState](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainState creates a new DomainState instance.
func NewDomainState() DomainState {
	return getDomainStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DomainState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState
type DomainState struct {
	objectivec.Object
}

// DomainStateFrom constructs a [DomainState] from an unsafe.Pointer.
func DomainStateFrom(ptr unsafe.Pointer) DomainState {
	return DomainState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DomainState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DomainState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DomainState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DomainState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DomainState */

// Contains biometric domain state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/biometry
func (d_ DomainState) Biometry() ILADomainStateBiometry {
	rv := objc.Send[DomainStateBiometry](d_.ID, objc.Sel("biometry"))
	return rv
}/* debug [instance_properties/getter]: biometry */


// Contains companion domain state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/companion
func (d_ DomainState) Companion() ILADomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](d_.ID, objc.Sel("companion"))
	return rv
}/* debug [instance_properties/getter]: companion */


// Contains combined state hash data for biometry and companion state hashes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainState/stateHash
func (d_ DomainState) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("stateHash"))
	return rv
}/* debug [instance_properties/getter]: stateHash */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LADomainState */



