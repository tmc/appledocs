// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKFindConfiguration */


/* debug [class_header]: Header for WKFindConfiguration */
// The class instance for the [FindConfiguration] class.
var (
	FindConfigurationClass     _FindConfigurationClass
	FindConfigurationClassOnce sync.Once
)

func getFindConfigurationClass() _FindConfigurationClass {
	FindConfigurationClassOnce.Do(func() {
		FindConfigurationClass = _FindConfigurationClass{objc.GetClass("WKFindConfiguration")}
	})
	return FindConfigurationClass
}

type _FindConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FindConfiguration */
// An interface definition for the [FindConfiguration] class.
type IFindConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FindConfiguration */
	// properties:
	Backwards() bool
	SetBackwards(value bool)
	CaseSensitive() bool
	SetCaseSensitive(value bool)
	Wraps() bool
	SetWraps(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FindConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FindConfiguration */
// Alloc allocates a new instance without initialization.
func (fc _FindConfigurationClass) Alloc() FindConfiguration {
	rv := objc.Send[FindConfiguration](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FindConfigurationClass) New() FindConfiguration {
	rv := objc.Send[FindConfiguration](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FindConfiguration) Init() FindConfiguration {
	rv := objc.Send[FindConfiguration](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FindConfiguration) Autorelease() FindConfiguration {
	rv := objc.Send[FindConfiguration](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFindConfiguration creates a new FindConfiguration instance.
func NewFindConfiguration() FindConfiguration {
	return getFindConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FindConfiguration */
// The configuration parameters to use when searching the contents of the web view.
//
// Create a object and configure its attributes to specify how to perform searches within the web view’s contents. To initiate a search, call the appropriate method of and pass this object along with the search string.


// The configuration parameters to use when searching the contents of the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration
type FindConfiguration struct {
	objectivec.Object
}

// FindConfigurationFrom constructs a [FindConfiguration] from an unsafe.Pointer.
//
// The configuration parameters to use when searching the contents of the web view.
func FindConfigurationFrom(ptr unsafe.Pointer) FindConfiguration {
	return FindConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FindConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FindConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FindConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FindConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FindConfiguration */

// A Boolean value that indicates the search direction, relative to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/backwards
func (f_ FindConfiguration) Backwards() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("backwards"))
	return rv
}/* debug [instance_properties/getter]: backwards */


// A Boolean value that indicates the search direction, relative to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/backwards
func (f_ FindConfiguration) SetBackwards(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBackwards:"), value)
}/* debug [instance_properties/setter]: backwards */


// A Boolean value that indicates whether to consider case when matching the search string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/caseSensitive
func (f_ FindConfiguration) CaseSensitive() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("caseSensitive"))
	return rv
}/* debug [instance_properties/getter]: caseSensitive */


// A Boolean value that indicates whether to consider case when matching the search string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/caseSensitive
func (f_ FindConfiguration) SetCaseSensitive(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCaseSensitive:"), value)
}/* debug [instance_properties/setter]: caseSensitive */


// A Boolean value that indicates whether the search wraps around to the other side of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/wraps
func (f_ FindConfiguration) Wraps() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("wraps"))
	return rv
}/* debug [instance_properties/getter]: wraps */


// A Boolean value that indicates whether the search wraps around to the other side of the page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/wraps
func (f_ FindConfiguration) SetWraps(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWraps:"), value)
}/* debug [instance_properties/setter]: wraps */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKFindConfiguration */



