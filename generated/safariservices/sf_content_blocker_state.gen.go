// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFContentBlockerState */


/* debug [class_header]: Header for SFContentBlockerState */
// The class instance for the [SFContentBlockerState] class.
var (
	SFContentBlockerStateClass     _SFContentBlockerStateClass
	SFContentBlockerStateClassOnce sync.Once
)

func getSFContentBlockerStateClass() _SFContentBlockerStateClass {
	SFContentBlockerStateClassOnce.Do(func() {
		SFContentBlockerStateClass = _SFContentBlockerStateClass{objc.GetClass("SFContentBlockerState")}
	})
	return SFContentBlockerStateClass
}

type _SFContentBlockerStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFContentBlockerState */
// An interface definition for the [SFContentBlockerState] class.
type ISFContentBlockerState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFContentBlockerState */
	// properties:
	Enabled() bool
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFContentBlockerState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFContentBlockerState */
// Alloc allocates a new instance without initialization.
func (sc _SFContentBlockerStateClass) Alloc() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFContentBlockerStateClass) New() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFContentBlockerState) Init() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFContentBlockerState) Autorelease() SFContentBlockerState {
	rv := objc.Send[SFContentBlockerState](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFContentBlockerState creates a new SFContentBlockerState instance.
func NewSFContentBlockerState() SFContentBlockerState {
	return getSFContentBlockerStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFContentBlockerState */
// The state of a content blocker extension.


// The state of a content blocker extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerState
type SFContentBlockerState struct {
	objectivec.Object
}

// SFContentBlockerStateFrom constructs a [SFContentBlockerState] from an unsafe.Pointer.
//
// The state of a content blocker extension.
func SFContentBlockerStateFrom(ptr unsafe.Pointer) SFContentBlockerState {
	return SFContentBlockerState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFContentBlockerState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFContentBlockerState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFContentBlockerState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFContentBlockerState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFContentBlockerState */

// A Boolean value that indicates whether the content blocker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerState/isEnabled
func (s_ SFContentBlockerState) Enabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// A Boolean value that indicates whether the content blocker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfcontentblockerstate/isenabled
func (s_ SFContentBlockerState) IsEnabled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value that indicates whether the content blocker is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfcontentblockerstate/isenabled
func (s_ SFContentBlockerState) SetIsEnabled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFContentBlockerState */



