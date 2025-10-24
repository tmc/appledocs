// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class EKSource */


/* debug [class_header]: Header for EKSource */
// The class instance for the [EKSource] class.
var (
	EKSourceClass     _EKSourceClass
	EKSourceClassOnce sync.Once
)

func getEKSourceClass() _EKSourceClass {
	EKSourceClassOnce.Do(func() {
		EKSourceClass = _EKSourceClass{objc.GetClass("EKSource")}
	})
	return EKSourceClass
}

type _EKSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EKSource */
// An interface definition for the [EKSource] class.
type IEKSource interface {
	IEKObject
	
/* debug [class_interface_properties]: Properties for EKSource */
	// properties:
	IsDelegate() bool
	SourceIdentifier() objc.IObject /* cross-framework: NSString */
	SourceType() EKSourceType
	Title() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EKSource */
	// methods:
	CalendarsForEntityType(entityType EKEntityType) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EKSource */
// Alloc allocates a new instance without initialization.
func (ec _EKSourceClass) Alloc() EKSource {
	rv := objc.Send[EKSource](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EKSourceClass) New() EKSource {
	rv := objc.Send[EKSource](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EKSource) Init() EKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EKSource) Autorelease() EKSource {
	rv := objc.Send[EKSource](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEKSource creates a new EKSource instance.
func NewEKSource() EKSource {
	return getEKSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EKSource */
// An abstract superclass that represents the account a calendar belongs to.
//
// You do not create instances of this class; instead, you retrieve objects from an object. Use the sources property to get all the objects for an event store, and use the methods in this class to access properties of the source object.


// An abstract superclass that represents the account a calendar belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource
type EKSource struct {
	EKObject
}

// EKSourceFrom constructs a [EKSource] from an unsafe.Pointer.
//
// An abstract superclass that represents the account a calendar belongs to.
func EKSourceFrom(ptr unsafe.Pointer) EKSource {
	return EKSource{
		EKObject: EKObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EKSource *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EKSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EKSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EKSource */

// Returns the calendars that belong to this source object that support a particular entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/calendars(for:)
func (e_ EKSource) CalendarsForEntityType(entityType EKEntityType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("calendarsForEntityType:"), entityType)
	return rv
}/* debug [instance_methods/method]: CalendarsForEntityType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EKSource */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/isDelegate
func (e_ EKSource) IsDelegate() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDelegate"))
	return rv
}/* debug [instance_properties/getter]: isDelegate */


// A unique identifier for the source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/sourceIdentifier
func (e_ EKSource) SourceIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("sourceIdentifier"))
	return rv
}/* debug [instance_properties/getter]: sourceIdentifier */


// The type of this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/sourceType
func (e_ EKSource) SourceType() EKSourceType {
	rv := objc.Send[EKSourceType](e_.ID, objc.Sel("sourceType"))
	return rv
}/* debug [instance_properties/getter]: sourceType */


// The name of this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/title
func (e_ EKSource) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EKSource */


