// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariExtensionState */


/* debug [class_header]: Header for SFSafariExtensionState */
// The class instance for the [SFSafariExtensionState] class.
var (
	SFSafariExtensionStateClass     _SFSafariExtensionStateClass
	SFSafariExtensionStateClassOnce sync.Once
)

func getSFSafariExtensionStateClass() _SFSafariExtensionStateClass {
	SFSafariExtensionStateClassOnce.Do(func() {
		SFSafariExtensionStateClass = _SFSafariExtensionStateClass{objc.GetClass("SFSafariExtensionState")}
	})
	return SFSafariExtensionStateClass
}

type _SFSafariExtensionStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariExtensionState */
// An interface definition for the [SFSafariExtensionState] class.
type ISFSafariExtensionState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariExtensionState */
	// properties:
	Enabled() bool
	SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariExtensionState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariExtensionState */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionStateClass) Alloc() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariExtensionStateClass) New() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionState) Init() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionState) Autorelease() SFSafariExtensionState {
	rv := objc.Send[SFSafariExtensionState](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionState creates a new SFSafariExtensionState instance.
func NewSFSafariExtensionState() SFSafariExtensionState {
	return getSFSafariExtensionStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariExtensionState */
// The state of a Safari app extension.


// The state of a Safari app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionState
type SFSafariExtensionState struct {
	objectivec.Object
}

// SFSafariExtensionStateFrom constructs a [SFSafariExtensionState] from an unsafe.Pointer.
//
// The state of a Safari app extension.
func SFSafariExtensionStateFrom(ptr unsafe.Pointer) SFSafariExtensionState {
	return SFSafariExtensionState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariExtensionState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariExtensionState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariExtensionState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariExtensionState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariExtensionState */

// A Boolean value that indicates whether the user has enabled the app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionState/isEnabled
func (s_ SFSafariExtensionState) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariExtensionState) SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}/* debug [instance_properties/getter]: SFExtensionProfileKey */


// A Boolean value that indicates whether the user has enabled the app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariextensionstate/isenabled
func (s_ SFSafariExtensionState) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the user has enabled the app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariextensionstate/isenabled
func (s_ SFSafariExtensionState) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariExtensionState */



