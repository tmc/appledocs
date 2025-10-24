// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariPageProperties */


/* debug [class_header]: Header for SFSafariPageProperties */
// The class instance for the [SFSafariPageProperties] class.
var (
	SFSafariPagePropertiesClass     _SFSafariPagePropertiesClass
	SFSafariPagePropertiesClassOnce sync.Once
)

func getSFSafariPagePropertiesClass() _SFSafariPagePropertiesClass {
	SFSafariPagePropertiesClassOnce.Do(func() {
		SFSafariPagePropertiesClass = _SFSafariPagePropertiesClass{objc.GetClass("SFSafariPageProperties")}
	})
	return SFSafariPagePropertiesClass
}

type _SFSafariPagePropertiesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariPageProperties */
// An interface definition for the [SFSafariPageProperties] class.
type ISFSafariPageProperties interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariPageProperties */
	// properties:
	Active() bool
	Title() objc.IObject /* cross-framework: NSString */
	Url() objc.IObject /* cross-framework: NSURL */
	UsesPrivateBrowsing() bool
	SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */
	IsActive() bool
	SetIsActive(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariPageProperties */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariPageProperties */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariPagePropertiesClass) Alloc() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariPagePropertiesClass) New() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariPageProperties) Init() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariPageProperties) Autorelease() SFSafariPageProperties {
	rv := objc.Send[SFSafariPageProperties](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariPageProperties creates a new SFSafariPageProperties instance.
func NewSFSafariPageProperties() SFSafariPageProperties {
	return getSFSafariPagePropertiesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariPageProperties */
// An object that captures information about a webpage.
//
// Use the properties object to retrieve page information, such as the current URL, page title, active status, and private browsing status.


// An object that captures information about a webpage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties
type SFSafariPageProperties struct {
	objectivec.Object
}

// SFSafariPagePropertiesFrom constructs a [SFSafariPageProperties] from an unsafe.Pointer.
//
// An object that captures information about a webpage.
func SFSafariPagePropertiesFrom(ptr unsafe.Pointer) SFSafariPageProperties {
	return SFSafariPageProperties{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariPageProperties *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariPageProperties */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariPageProperties */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariPageProperties */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariPageProperties */

// A Boolean value that indicates whether the page is currently active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/isActive
func (s_ SFSafariPageProperties) Active() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("active"))
	return rv
}/* debug [instance_properties/getter]: active */


// The title of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/title
func (s_ SFSafariPageProperties) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// Indicates the URL of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/url
func (s_ SFSafariPageProperties) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// A Boolean value that indicates whether the page is using Safari Private Browsing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariPageProperties/usesPrivateBrowsing
func (s_ SFSafariPageProperties) UsesPrivateBrowsing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesPrivateBrowsing"))
	return rv
}/* debug [instance_properties/getter]: usesPrivateBrowsing */


// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariPageProperties) SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}/* debug [instance_properties/getter]: SFExtensionProfileKey */


// A Boolean value that indicates whether the page is currently active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafaripageproperties/isactive
func (s_ SFSafariPageProperties) IsActive() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isActive"))
	return rv
}/* debug [instance_properties/getter]: isActive */


// A Boolean value that indicates whether the page is currently active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafaripageproperties/isactive
func (s_ SFSafariPageProperties) SetIsActive(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsActive:"), value)
}/* debug [instance_properties/setter]: isActive */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariPageProperties */



