// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariViewControllerConfiguration */


/* debug [class_header]: Header for SFSafariViewControllerConfiguration */
// The class instance for the [SFSafariViewControllerConfiguration] class.
var (
	SFSafariViewControllerConfigurationClass     _SFSafariViewControllerConfigurationClass
	SFSafariViewControllerConfigurationClassOnce sync.Once
)

func getSFSafariViewControllerConfigurationClass() _SFSafariViewControllerConfigurationClass {
	SFSafariViewControllerConfigurationClassOnce.Do(func() {
		SFSafariViewControllerConfigurationClass = _SFSafariViewControllerConfigurationClass{objc.GetClass("SFSafariViewControllerConfiguration")}
	})
	return SFSafariViewControllerConfigurationClass
}

type _SFSafariViewControllerConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariViewControllerConfiguration */
// An interface definition for the [SFSafariViewControllerConfiguration] class.
type ISFSafariViewControllerConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariViewControllerConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariViewControllerConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariViewControllerConfiguration */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerConfigurationClass) Alloc() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariViewControllerConfigurationClass) New() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerConfiguration) Init() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerConfiguration) Autorelease() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerConfiguration creates a new SFSafariViewControllerConfiguration instance.
func NewSFSafariViewControllerConfiguration() SFSafariViewControllerConfiguration {
	return getSFSafariViewControllerConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariViewControllerConfiguration */
// A configuration object that defines how a Safari view controller should be initialized.
//
// Use a configuration object with the method to initialize your view controller.


// A configuration object that defines how a Safari view controller should be initialized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/Configuration-swift.class
type SFSafariViewControllerConfiguration struct {
	objectivec.Object
}

// SFSafariViewControllerConfigurationFrom constructs a [SFSafariViewControllerConfiguration] from an unsafe.Pointer.
//
// A configuration object that defines how a Safari view controller should be initialized.
func SFSafariViewControllerConfigurationFrom(ptr unsafe.Pointer) SFSafariViewControllerConfiguration {
	return SFSafariViewControllerConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariViewControllerConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariViewControllerConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariViewControllerConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariViewControllerConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariViewControllerConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariViewControllerConfiguration */


