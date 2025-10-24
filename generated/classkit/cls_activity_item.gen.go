// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CLSActivityItem */


/* debug [class_header]: Header for CLSActivityItem */
// The class instance for the [SActivityItem] class.
var (
	SActivityItemClass     _SActivityItemClass
	SActivityItemClassOnce sync.Once
)

func getSActivityItemClass() _SActivityItemClass {
	SActivityItemClassOnce.Do(func() {
		SActivityItemClass = _SActivityItemClass{objc.GetClass("CLSActivityItem")}
	})
	return SActivityItemClass
}

type _SActivityItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SActivityItem */
// An interface definition for the [SActivityItem] class.
type ISActivityItem interface {
	ISObject
	
/* debug [class_interface_properties]: Properties for SActivityItem */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SActivityItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SActivityItem */
// Alloc allocates a new instance without initialization.
func (sc _SActivityItemClass) Alloc() SActivityItem {
	rv := objc.Send[SActivityItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SActivityItemClass) New() SActivityItem {
	rv := objc.Send[SActivityItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SActivityItem) Init() SActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SActivityItem) Autorelease() SActivityItem {
	rv := objc.Send[SActivityItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSActivityItem creates a new SActivityItem instance.
func NewSActivityItem() SActivityItem {
	return getSActivityItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SActivityItem */
// An abstract base class for gathering information about an activity.
//
// You don’t typically use an instance of this class directly. Instead, use one of its subclasses to represent a particular activity metric. For example, use a to add a score to a activity.


// An abstract base class for gathering information about an activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSActivityItem
type SActivityItem struct {
	SObject
}

// SActivityItemFrom constructs a [SActivityItem] from an unsafe.Pointer.
//
// An abstract base class for gathering information about an activity.
func SActivityItemFrom(ptr unsafe.Pointer) SActivityItem {
	return SActivityItem{
		SObject: SObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SActivityItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SActivityItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SActivityItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SActivityItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SActivityItem */

// An identifier for the activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivityitem/identifier
func (s_ SActivityItem) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// An identifier for the activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivityitem/identifier
func (s_ SActivityItem) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A human readable name for the activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivityitem/title
func (s_ SActivityItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// A human readable name for the activity item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/classkit/clsactivityitem/title
func (s_ SActivityItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSActivityItem */



