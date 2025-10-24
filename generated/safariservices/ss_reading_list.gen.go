// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SSReadingList */


/* debug [class_header]: Header for SSReadingList */
// The class instance for the [SSReadingList] class.
var (
	SSReadingListClass     _SSReadingListClass
	SSReadingListClassOnce sync.Once
)

func getSSReadingListClass() _SSReadingListClass {
	SSReadingListClassOnce.Do(func() {
		SSReadingListClass = _SSReadingListClass{objc.GetClass("SSReadingList")}
	})
	return SSReadingListClass
}

type _SSReadingListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SSReadingList */
// An interface definition for the [SSReadingList] class.
type ISSReadingList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SSReadingList */
	// properties:
	SSReadingListErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SSReadingList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SSReadingList */
// Alloc allocates a new instance without initialization.
func (sc _SSReadingListClass) Alloc() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SSReadingListClass) New() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SSReadingList) Init() SSReadingList {
	rv := objc.Send[SSReadingList](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SSReadingList) Autorelease() SSReadingList {
	rv := objc.Send[SSReadingList](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSSReadingList creates a new SSReadingList instance.
func NewSSReadingList() SSReadingList {
	return getSSReadingListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SSReadingList */
// An object for adding items to a user’s Safari Reading List.


// An object for adding items to a user’s Safari Reading List.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList
type SSReadingList struct {
	objectivec.Object
}

// SSReadingListFrom constructs a [SSReadingList] from an unsafe.Pointer.
//
// An object for adding items to a user’s Safari Reading List.
func SSReadingListFrom(ptr unsafe.Pointer) SSReadingList {
	return SSReadingList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SSReadingList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SSReadingList */

// Returns the Safari Reading List singleton object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList/default()
func (sc _SSReadingListClass) DefaultReadingList() SSReadingList {
	rv := objc.Send[SSReadingList](objc.ID(sc.class), objc.Sel("defaultReadingList"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultReadingList) */


// Determines whether a URL can be added to the Reading List.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SSReadingList/supportsURL(_:)
func (sc _SSReadingListClass) SupportsURL(URL objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("supportsURL:"), URL)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SupportsURL) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SSReadingList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SSReadingList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SSReadingList */

// The domain for Safari Reading List errors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/ssreadinglisterrordomain
func (s_ SSReadingList) SSReadingListErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SSReadingListErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: SSReadingListErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SSReadingList */


