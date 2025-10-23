// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLResponse] class.
var (
	URLResponseClass     _URLResponseClass
	URLResponseClassOnce sync.Once
)

func getURLResponseClass() _URLResponseClass {
	URLResponseClassOnce.Do(func() {
		URLResponseClass = _URLResponseClass{objc.GetClass("NSURLResponse")}
	})
	return URLResponseClass
}

type _URLResponseClass struct {
	class objc.Class
}

// An interface definition for the [URLResponse] class.
type IURLResponse interface {
	objectivec.IObject
	ExpectedContentLength() unsafe.Pointer
	SetExpectedContentLength(value unsafe.Pointer)
	MimeType() string
	SetMimeType(value string)
	SuggestedFilename() string
	SetSuggestedFilename(value string)
	TextEncodingName() string
	SetTextEncodingName(value string)
	Url() IURL
	SetUrl(value IURL)
}

// The metadata associated with the response to a URL load request, independent of protocol and URL scheme.
//
// The related class is a commonly used subclass of whose objects represent a response to an HTTP URL load request and store additional protocol-specific information such as the response headers. Whenever you make an HTTP request, the object you get back is actually an instance of the class.


// The metadata associated with the response to a URL load request, independent of protocol and URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResponse
type URLResponse struct {
	objectivec.Object
}

// URLResponseFrom constructs a [URLResponse] from an unsafe.Pointer.
//
// The metadata associated with the response to a URL load request, independent of protocol and URL scheme.
func URLResponseFrom(ptr unsafe.Pointer) URLResponse {
	return URLResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLResponseClass) Alloc() URLResponse {
	rv := objc.Send[URLResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLResponseClass) New() URLResponse {
	rv := objc.Send[URLResponse](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLResponse) Init() URLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLResponse) Autorelease() URLResponse {
	rv := objc.Send[URLResponse](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLResponse creates a new URLResponse instance.
func NewURLResponse() URLResponse {
	return getURLResponseClass().New()
}



// The expected length of the response’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/expectedcontentlength
func (u_ URLResponse) ExpectedContentLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("expectedContentLength"))
	return rv
}


// The expected length of the response’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/expectedcontentlength
func (u_ URLResponse) SetExpectedContentLength(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setExpectedContentLength:"), value)
}


// The MIME type of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/mimetype
func (u_ URLResponse) MimeType() string {
	rv := objc.Send[string](u_.ID, objc.Sel("mimeType"))
	return rv
}


// The MIME type of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/mimetype
func (u_ URLResponse) SetMimeType(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setMimeType:"), objc.String(value))
}


// A suggested filename for the response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/suggestedfilename
func (u_ URLResponse) SuggestedFilename() string {
	rv := objc.Send[string](u_.ID, objc.Sel("suggestedFilename"))
	return rv
}


// A suggested filename for the response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/suggestedfilename
func (u_ URLResponse) SetSuggestedFilename(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSuggestedFilename:"), objc.String(value))
}


// The name of the text encoding provided by the response’s originating source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/textencodingname
func (u_ URLResponse) TextEncodingName() string {
	rv := objc.Send[string](u_.ID, objc.Sel("textEncodingName"))
	return rv
}


// The name of the text encoding provided by the response’s originating source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/textencodingname
func (u_ URLResponse) SetTextEncodingName(value string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTextEncodingName:"), objc.String(value))
}


// The URL for the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/url
func (u_ URLResponse) Url() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("url"))
	return rv
}


// The URL for the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/url
func (u_ URLResponse) SetUrl(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}



