// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var URLSessionConfigurationClass _URLSessionConfigurationClass

func init() {
	URLSessionConfigurationClass = _URLSessionConfigurationClass{objc.GetClass("NSURLSessionConfiguration")}
}

type _URLSessionConfigurationClass struct {
	class objc.Class
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
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/backgroundSessionConfiguration(_:)
func (uc _URLSessionConfigurationClass) BackgroundSessionConfiguration(identifier string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("backgroundSessionConfiguration:"), identifier)
	return rv
}


