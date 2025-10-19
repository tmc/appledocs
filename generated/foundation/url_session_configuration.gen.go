// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionConfiguration] class.
var (
	uRLSessionConfigurationClass     _URLSessionConfigurationClass
	uRLSessionConfigurationClassOnce sync.Once
)

func getURLSessionConfigurationClass() _URLSessionConfigurationClass {
	uRLSessionConfigurationClassOnce.Do(func() {
		uRLSessionConfigurationClass = _URLSessionConfigurationClass{objc.GetClass("NSURLSessionConfiguration")}
	})
	return uRLSessionConfigurationClass
}

type _URLSessionConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [URLSessionConfiguration] class.
type IURLSessionConfiguration interface {
	objectivec.IObject
}

// A configuration object that defines behavior and policies for a URL session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration
type URLSessionConfiguration struct {
	objectivec.Object
}

// URLSessionConfigurationFrom constructs a [URLSessionConfiguration] from an unsafe.Pointer.
//
// A configuration object that defines behavior and policies for a URL session.
func URLSessionConfigurationFrom(ptr unsafe.Pointer) URLSessionConfiguration {
	return URLSessionConfiguration{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLSessionConfigurationClass) Alloc() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLSessionConfigurationClass) New() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLSessionConfiguration) Init() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLSessionConfiguration) Autorelease() URLSessionConfiguration {
	rv := objc.Send[URLSessionConfiguration](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLSessionConfiguration creates a new URLSessionConfiguration instance.
func NewURLSessionConfiguration() URLSessionConfiguration {
	return getURLSessionConfigurationClass().New()
}


// Returns a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfiguration(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("backgroundSessionConfiguration:"), objc.String(identifier))
	return rv
}


