// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariViewControllerActivityButton */


/* debug [class_header]: Header for SFSafariViewControllerActivityButton */
// The class instance for the [SFSafariViewControllerActivityButton] class.
var (
	SFSafariViewControllerActivityButtonClass     _SFSafariViewControllerActivityButtonClass
	SFSafariViewControllerActivityButtonClassOnce sync.Once
)

func getSFSafariViewControllerActivityButtonClass() _SFSafariViewControllerActivityButtonClass {
	SFSafariViewControllerActivityButtonClassOnce.Do(func() {
		SFSafariViewControllerActivityButtonClass = _SFSafariViewControllerActivityButtonClass{objc.GetClass("SFSafariViewControllerActivityButton")}
	})
	return SFSafariViewControllerActivityButtonClass
}

type _SFSafariViewControllerActivityButtonClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariViewControllerActivityButton */
// An interface definition for the [SFSafariViewControllerActivityButton] class.
type ISFSafariViewControllerActivityButton interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariViewControllerActivityButton */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariViewControllerActivityButton */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariViewControllerActivityButton */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerActivityButtonClass) Alloc() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariViewControllerActivityButtonClass) New() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewControllerActivityButton) Init() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewControllerActivityButton) Autorelease() SFSafariViewControllerActivityButton {
	rv := objc.Send[SFSafariViewControllerActivityButton](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewControllerActivityButton creates a new SFSafariViewControllerActivityButton instance.
func NewSFSafariViewControllerActivityButton() SFSafariViewControllerActivityButton {
	return getSFSafariViewControllerActivityButtonClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariViewControllerActivityButton */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton
type SFSafariViewControllerActivityButton struct {
	objectivec.Object
}

// SFSafariViewControllerActivityButtonFrom constructs a [SFSafariViewControllerActivityButton] from an unsafe.Pointer.
func SFSafariViewControllerActivityButtonFrom(ptr unsafe.Pointer) SFSafariViewControllerActivityButton {
	return SFSafariViewControllerActivityButton{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariViewControllerActivityButton */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/ActivityButton/init(templateImage:extensionIdentifier:)
func NewSFSafariViewControllerActivityButtonWithTemplateImageExtensionIdentifier(templateImage appkit.Image, extensionIdentifier objc.IObject /* cross-framework: NSString */) SFSafariViewControllerActivityButton {
	instance := getSFSafariViewControllerActivityButtonClass().Alloc()
	rv := objc.Send[SFSafariViewControllerActivityButton](instance.ID, objc.Sel("initWithTemplateImage:extensionIdentifier:"), templateImage, extensionIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSafariViewControllerActivityButtonWithTemplateImageExtensionIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariViewControllerActivityButton */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariViewControllerActivityButton */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariViewControllerActivityButton */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariViewControllerActivityButton */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariViewControllerActivityButton */


