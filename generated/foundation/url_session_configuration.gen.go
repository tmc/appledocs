// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLSessionConfiguration] class.
var uRLSessionConfigurationClass = _URLSessionConfigurationClass{objc.GetClass("NSURLSessionConfiguration")}

type _URLSessionConfigurationClass struct {
	class objc.Class
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

// Returns a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfiguration(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("backgroundSessionConfiguration:"), identifier)
	return rv
}


