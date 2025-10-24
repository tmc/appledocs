// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKFindResult */


/* debug [class_header]: Header for WKFindResult */
// The class instance for the [FindResult] class.
var (
	FindResultClass     _FindResultClass
	FindResultClassOnce sync.Once
)

func getFindResultClass() _FindResultClass {
	FindResultClassOnce.Do(func() {
		FindResultClass = _FindResultClass{objc.GetClass("WKFindResult")}
	})
	return FindResultClass
}

type _FindResultClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FindResult */
// An interface definition for the [FindResult] class.
type IFindResult interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FindResult */
	// properties:
	MatchFound() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FindResult */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FindResult */
// Alloc allocates a new instance without initialization.
func (fc _FindResultClass) Alloc() FindResult {
	rv := objc.Send[FindResult](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FindResultClass) New() FindResult {
	rv := objc.Send[FindResult](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FindResult) Init() FindResult {
	rv := objc.Send[FindResult](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FindResult) Autorelease() FindResult {
	rv := objc.Send[FindResult](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFindResult creates a new FindResult instance.
func NewFindResult() FindResult {
	return getFindResultClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FindResult */
// An object that contains the results of searching the web view’s contents.
//
// When you perform a search using the methods of , the web view creates a object and delivers it to your completion handler. You don’t create instances of this class directly. Use the objects that the web view provides to determine whether it found a match for the content.


// An object that contains the results of searching the web view’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindResult
type FindResult struct {
	objectivec.Object
}

// FindResultFrom constructs a [FindResult] from an unsafe.Pointer.
//
// An object that contains the results of searching the web view’s contents.
func FindResultFrom(ptr unsafe.Pointer) FindResult {
	return FindResult{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FindResult *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FindResult */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FindResult */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FindResult */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FindResult */

// A Boolean value that indicates whether the web view found a match during the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindResult/matchFound
func (f_ FindResult) MatchFound() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("matchFound"))
	return rv
}/* debug [instance_properties/getter]: matchFound */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKFindResult */



