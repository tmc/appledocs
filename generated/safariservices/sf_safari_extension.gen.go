// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariExtension */


/* debug [class_header]: Header for SFSafariExtension */
// The class instance for the [SFSafariExtension] class.
var (
	SFSafariExtensionClass     _SFSafariExtensionClass
	SFSafariExtensionClassOnce sync.Once
)

func getSFSafariExtensionClass() _SFSafariExtensionClass {
	SFSafariExtensionClassOnce.Do(func() {
		SFSafariExtensionClass = _SFSafariExtensionClass{objc.GetClass("SFSafariExtension")}
	})
	return SFSafariExtensionClass
}

type _SFSafariExtensionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariExtension */
// An interface definition for the [SFSafariExtension] class.
type ISFSafariExtension interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariExtension */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariExtension */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariExtension */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionClass) Alloc() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariExtensionClass) New() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtension) Init() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtension) Autorelease() SFSafariExtension {
	rv := objc.Send[SFSafariExtension](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtension creates a new SFSafariExtension instance.
func NewSFSafariExtension() SFSafariExtension {
	return getSFSafariExtensionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariExtension */
// A proxy for the Safari extension.


// A proxy for the Safari extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtension
type SFSafariExtension struct {
	objectivec.Object
}

// SFSafariExtensionFrom constructs a [SFSafariExtension] from an unsafe.Pointer.
//
// A proxy for the Safari extension.
func SFSafariExtensionFrom(ptr unsafe.Pointer) SFSafariExtension {
	return SFSafariExtension{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariExtension *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariExtension */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtension/getBaseURI(completionHandler:)
func (sc _SFSafariExtensionClass) GetBaseURIWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getBaseURIWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetBaseURIWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariExtension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariExtension */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariExtension */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariExtension */



