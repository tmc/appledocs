// Code generated from Apple documentation for EventKit. DO NOT EDIT.

package eventkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [EKSource] class.
type IEKSource interface {
	IEKObject
	// properties:
	Calendars() unsafe.Pointer
	IsDelegate() bool /* primitive/slice/pointer. */
	SourceIdentifier() string /* primitive/slice/pointer. */
	SourceType() EKSourceType
	Title() string /* primitive/slice/pointer. */
	// methods:
	CalendarsForEntityType(entityType EKEntityType) unsafe.Pointer
}

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

// Alloc allocates a new instance without initialization.
func (ec _EKSourceClass) Alloc() EKSource {
	rv := objc.Send[EKSource](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns the calendars that belong to this source object that support a particular entity type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/calendars(for:)
func (e_ EKSource) CalendarsForEntityType(entityType EKEntityType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("calendarsForEntityType:"), entityType)
	return rv
}


// The calendars that belong to this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/calendars
func (e_ EKSource) Calendars() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](e_.ID, objc.Sel("calendars"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/isDelegate
func (e_ EKSource) IsDelegate() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](e_.ID, objc.Sel("isDelegate"))
	return rv
}


// A unique identifier for the source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/sourceIdentifier
func (e_ EKSource) SourceIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("sourceIdentifier"))
	return rv
}


// The type of this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/sourceType
func (e_ EKSource) SourceType() EKSourceType {
	rv := objc.Send[EKSourceType](e_.ID, objc.Sel("sourceType"))
	return rv
}


// The name of this source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/EventKit/EKSource/title
func (e_ EKSource) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](e_.ID, objc.Sel("title"))
	return rv
}



