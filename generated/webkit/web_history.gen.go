// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebHistory */

/* debug [class_header]: Header for WebHistory */
// The class instance for the [WebHistory] class.
var (
	WebHistoryClass     _WebHistoryClass
	WebHistoryClassOnce sync.Once
)

func getWebHistoryClass() _WebHistoryClass {
	WebHistoryClassOnce.Do(func() {
		WebHistoryClass = _WebHistoryClass{objc.GetClass("WebHistory")}
	})
	return WebHistoryClass
}

type _WebHistoryClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebHistory */
// An interface definition for the [WebHistory] class.
type IWebHistory interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebHistory */
	// properties:
	HistoryAgeInDaysLimit() int
	SetHistoryAgeInDaysLimit(value int)
	HistoryItemLimit() int
	SetHistoryItemLimit(value int)
	OrderedLastVisitedDays() objc.IObject /* cross-framework: NSArray */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebHistory */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebHistory */
// Alloc allocates a new instance without initialization.
func (wc _WebHistoryClass) Alloc() WebHistory {
	rv := objc.Send[WebHistory](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebHistoryClass) New() WebHistory {
	rv := objc.Send[WebHistory](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebHistory) Init() WebHistory {
	rv := objc.Send[WebHistory](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebHistory) Autorelease() WebHistory {
	rv := objc.Send[WebHistory](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebHistory creates a new WebHistory instance.
func NewWebHistory() WebHistory {
	return getWebHistoryClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebHistory */
// objects are used to maintain the pages visited by users. Visited pages are represented by objects. You add and remove history items using the and methods. These methods post appropriate notifications when items are added or removed so you can update the display. organizes the objects by the day they were visited, ordered from most recent to oldest. You can request all the days that contain history items using the method or request the items visited on a particular day using the method. objects can be loaded and saved by specifying a file URL (see ).

// objects are used to maintain the pages visited by users. Visited pages are represented by objects. You add and remove history items using the and methods. These methods post appropriate notifications when items are added or removed so you can update the display. organizes the objects by the day they were visited, ordered from most recent to oldest. You can request all the days that contain history items using the method or request the items visited on a particular day using the method. objects can be loaded and saved by specifying a file URL (see ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory
type WebHistory struct {
	objectivec.Object
}

// WebHistoryFrom constructs a [WebHistory] from an unsafe.Pointer.
//
// objects are used to maintain the pages visited by users. Visited pages are represented by objects. You add and remove history items using the and methods. These methods post appropriate notifications when items are added or removed so you can update the display. organizes the objects by the day they were visited, ordered from most recent to oldest. You can request all the days that contain history items using the method or request the items visited on a particular day using the method. objects can be loaded and saved by specifying a file URL (see ).
func WebHistoryFrom(ptr unsafe.Pointer) WebHistory {
	return WebHistory{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebHistory */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebHistory */

// Returns a shared web history object, if one exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/optionalShared()
func (wc _WebHistoryClass) OptionalSharedHistory() WebHistory {
	rv := objc.Send[WebHistory](objc.ID(wc.class), objc.Sel("optionalSharedHistory"))
	return rv
} /* debug [class_methods/method]: Class method for%!(EXTRA string=OptionalSharedHistory) */

// Sets the web history object to share.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/setOptionalShared(_:)
func (wc _WebHistoryClass) SetOptionalSharedHistory(history IWebHistory) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("setOptionalSharedHistory:"), history)
} /* debug [class_methods/method]: Class method for%!(EXTRA string=SetOptionalSharedHistory) */

/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebHistory */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebHistory */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebHistory */

// The maximum age of web history items that can be retrieved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/historyAgeInDaysLimit
func (w_ WebHistory) HistoryAgeInDaysLimit() int {
	rv := objc.Send[int](w_.ID, objc.Sel("historyAgeInDaysLimit"))
	return rv
} /* debug [instance_properties/getter]: historyAgeInDaysLimit */

// The maximum age of web history items that can be retrieved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/historyAgeInDaysLimit
func (w_ WebHistory) SetHistoryAgeInDaysLimit(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHistoryAgeInDaysLimit:"), value)
} /* debug [instance_properties/setter]: historyAgeInDaysLimit */

// The maximum number of web history items that can be stored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/historyItemLimit
func (w_ WebHistory) HistoryItemLimit() int {
	rv := objc.Send[int](w_.ID, objc.Sel("historyItemLimit"))
	return rv
} /* debug [instance_properties/getter]: historyItemLimit */

// The maximum number of web history items that can be stored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/historyItemLimit
func (w_ WebHistory) SetHistoryItemLimit(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHistoryItemLimit:"), value)
} /* debug [instance_properties/setter]: historyItemLimit */

// An array of all calendar days represented in the web history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebHistory/orderedLastVisitedDays
func (w_ WebHistory) OrderedLastVisitedDays() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("orderedLastVisitedDays"))
	return rv
} /* debug [instance_properties/getter]: orderedLastVisitedDays */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebHistory */
