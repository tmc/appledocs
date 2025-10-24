// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFAddToHomeScreenInfo */


/* debug [class_header]: Header for SFAddToHomeScreenInfo */
// The class instance for the [SFAddToHomeScreenInfo] class.
var (
	SFAddToHomeScreenInfoClass     _SFAddToHomeScreenInfoClass
	SFAddToHomeScreenInfoClassOnce sync.Once
)

func getSFAddToHomeScreenInfoClass() _SFAddToHomeScreenInfoClass {
	SFAddToHomeScreenInfoClassOnce.Do(func() {
		SFAddToHomeScreenInfoClass = _SFAddToHomeScreenInfoClass{objc.GetClass("SFAddToHomeScreenInfo")}
	})
	return SFAddToHomeScreenInfoClass
}

type _SFAddToHomeScreenInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFAddToHomeScreenInfo */
// An interface definition for the [SFAddToHomeScreenInfo] class.
type ISFAddToHomeScreenInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFAddToHomeScreenInfo */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFAddToHomeScreenInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFAddToHomeScreenInfo */
// Alloc allocates a new instance without initialization.
func (sc _SFAddToHomeScreenInfoClass) Alloc() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFAddToHomeScreenInfoClass) New() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFAddToHomeScreenInfo) Init() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFAddToHomeScreenInfo) Autorelease() SFAddToHomeScreenInfo {
	rv := objc.Send[SFAddToHomeScreenInfo](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFAddToHomeScreenInfo creates a new SFAddToHomeScreenInfo instance.
func NewSFAddToHomeScreenInfo() SFAddToHomeScreenInfo {
	return getSFAddToHomeScreenInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFAddToHomeScreenInfo */
// A class that provides information about a web app that someone adds to their Home Screen.


// A class that provides information about a web app that someone adds to their Home Screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo
type SFAddToHomeScreenInfo struct {
	objectivec.Object
}

// SFAddToHomeScreenInfoFrom constructs a [SFAddToHomeScreenInfo] from an unsafe.Pointer.
//
// A class that provides information about a web app that someone adds to their Home Screen.
func SFAddToHomeScreenInfoFrom(ptr unsafe.Pointer) SFAddToHomeScreenInfo {
	return SFAddToHomeScreenInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFAddToHomeScreenInfo */

// Initializes a Home Screen information object with the supplied web app manifest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFAddToHomeScreenInfo/init(manifest:)
func NewSFAddToHomeScreenInfoWithManifest(manifest unsafe.Pointer) SFAddToHomeScreenInfo {
	instance := getSFAddToHomeScreenInfoClass().Alloc()
	rv := objc.Send[SFAddToHomeScreenInfo](instance.ID, objc.Sel("initWithManifest:"), manifest)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFAddToHomeScreenInfoWithManifest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFAddToHomeScreenInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFAddToHomeScreenInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFAddToHomeScreenInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFAddToHomeScreenInfo */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFAddToHomeScreenInfo */


