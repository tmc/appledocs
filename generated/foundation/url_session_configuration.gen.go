// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [URLSessionConfiguration] class.
var URLSessionConfigurationClass objc.Class

func init() {
	URLSessionConfigurationClass = objc.GetClass("NSURLSessionConfiguration")
}

type URLSessionConfiguration struct {
	objc.ID
}

func URLSessionConfigurationFrom(ptr unsafe.Pointer) URLSessionConfiguration {
	return URLSessionConfiguration{
		ID: objc.ID(ptr),
	}
}


// Returns a session configuration object that allows HTTP and HTTPS uploads or downloads to be performed in the background. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc URLSessionConfiguration) BackgroundSessionConfiguration(identifier string) unsafe.Pointer {
	sel := objc.RegisterName("backgroundSessionConfiguration:")
	ret := objc.ID(URLSessionConfigurationClass).Send(sel, identifier)
	return unsafe.Pointer(ret)
}

