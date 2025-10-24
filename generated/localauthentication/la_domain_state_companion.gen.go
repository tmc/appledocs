// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class LADomainStateCompanion */


/* debug [class_header]: Header for LADomainStateCompanion */
// The class instance for the [DomainStateCompanion] class.
var (
	DomainStateCompanionClass     _DomainStateCompanionClass
	DomainStateCompanionClassOnce sync.Once
)

func getDomainStateCompanionClass() _DomainStateCompanionClass {
	DomainStateCompanionClassOnce.Do(func() {
		DomainStateCompanionClass = _DomainStateCompanionClass{objc.GetClass("LADomainStateCompanion")}
	})
	return DomainStateCompanionClass
}

type _DomainStateCompanionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DomainStateCompanion */
// An interface definition for the [DomainStateCompanion] class.
type IDomainStateCompanion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DomainStateCompanion */
	// properties:
	AvailableCompanionTypes() unsafe.Pointer
	StateHash() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DomainStateCompanion */
	// methods:
	StateHashForCompanionType(companionType CompanionType) foundation.Data
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DomainStateCompanion */
// Alloc allocates a new instance without initialization.
func (dc _DomainStateCompanionClass) Alloc() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DomainStateCompanionClass) New() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DomainStateCompanion) Init() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DomainStateCompanion) Autorelease() DomainStateCompanion {
	rv := objc.Send[DomainStateCompanion](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDomainStateCompanion creates a new DomainStateCompanion instance.
func NewDomainStateCompanion() DomainStateCompanion {
	return getDomainStateCompanionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DomainStateCompanion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion
type DomainStateCompanion struct {
	objectivec.Object
}

// DomainStateCompanionFrom constructs a [DomainStateCompanion] from an unsafe.Pointer.
func DomainStateCompanionFrom(ptr unsafe.Pointer) DomainStateCompanion {
	return DomainStateCompanion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DomainStateCompanion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DomainStateCompanion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DomainStateCompanion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DomainStateCompanion */

// Returns state hash data for the given companion type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash(for:)
func (d_ DomainStateCompanion) StateHashForCompanionType(companionType CompanionType) foundation.Data {
	rv := objc.Send[foundation.Data](d_.ID, objc.Sel("stateHashForCompanionType:"), companionType)
	return rv
}/* debug [instance_methods/method]: StateHashForCompanionType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DomainStateCompanion */

// Indicates types of companions paired with the device. The elements are NSNumber-wrapped instances of @c .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/availableCompanionTypes-1ggnh
func (d_ DomainStateCompanion) AvailableCompanionTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("availableCompanionTypes"))
	return rv
}/* debug [instance_properties/getter]: availableCompanionTypes */


// Contains combined state hash data for all available companion types. . Returns if no companion devices are paired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LADomainStateCompanion/stateHash
func (d_ DomainStateCompanion) StateHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](d_.ID, objc.Sel("stateHash"))
	return rv
}/* debug [instance_properties/getter]: stateHash */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class LADomainStateCompanion */



