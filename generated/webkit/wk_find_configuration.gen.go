// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [FindConfiguration] class.
type IFindConfiguration interface {
	objectivec.IObject
	Wraps() bool
	SetWraps(value bool)
	Backwards() bool
	SetBackwards(value bool)
	CaseSensitive() bool
	SetCaseSensitive(value bool)
}

// The configuration parameters to use when searching the contents of the web view.
//
// Create a object and configure its attributes to specify how to perform searches within the web view’s contents. To initiate a search, call the appropriate method of and pass this object along with the search string.
//
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

// Alloc allocates a new instance without initialization.
func (fc _FindConfigurationClass) Alloc() FindConfiguration {
	rv := objc.Send[FindConfiguration](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A Boolean value that indicates whether the search wraps around to the other side of the page.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/wraps
func (f_ FindConfiguration) Wraps() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("wraps"))
	return rv
}


// SetWraps sets the value of the wraps property.
// A Boolean value that indicates whether the search wraps around to the other side of the page.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/wraps
func (f_ FindConfiguration) SetWraps(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setWraps:"), value)
}

// A Boolean value that indicates the search direction, relative to the current selection.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/backwards
func (f_ FindConfiguration) Backwards() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("backwards"))
	return rv
}


// SetBackwards sets the value of the backwards property.
// A Boolean value that indicates the search direction, relative to the current selection.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/backwards
func (f_ FindConfiguration) SetBackwards(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setBackwards:"), value)
}

// A Boolean value that indicates whether to consider case when matching the search string.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/casesensitive
func (f_ FindConfiguration) CaseSensitive() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("caseSensitive"))
	return rv
}


// SetCaseSensitive sets the value of the caseSensitive property.
// A Boolean value that indicates whether to consider case when matching the search string.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkfindconfiguration/casesensitive
func (f_ FindConfiguration) SetCaseSensitive(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setCaseSensitive:"), value)
}



