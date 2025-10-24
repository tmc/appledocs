// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLResponse */


/* debug [class_header]: Header for NSURLResponse */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLResponse */
// An interface definition for the [URLResponse] class.
type IURLResponse interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLResponse */
	// properties:
	ExpectedContentLength() objectivec.IObject
	MIMEType() IString
	SuggestedFilename() IString
	SetSuggestedFilename(value IString)
	TextEncodingName() IString
	SetTextEncodingName(value IString)
	Url() IURL
	SetUrl(value IURL)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLResponse */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLResponse */
// Alloc allocates a new instance without initialization.
func (uc _URLResponseClass) Alloc() URLResponse {
	rv := objc.Send[URLResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLResponse */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLResponse *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLResponse */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLResponse */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLResponse */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLResponse */

// The expected length of the response’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResponse/expectedContentLength
func (u_ URLResponse) ExpectedContentLength() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("expectedContentLength"))
	return rv
}/* debug [instance_properties/getter]: expectedContentLength */


// The MIME type of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResponse/mimeType
func (u_ URLResponse) MIMEType() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("MIMEType"))
	return rv
}/* debug [instance_properties/getter]: MIMEType */


// A suggested filename for the response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/suggestedfilename
func (u_ URLResponse) SuggestedFilename() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("suggestedFilename"))
	return rv
}/* debug [instance_properties/getter]: suggestedFilename */


// A suggested filename for the response data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/suggestedfilename
func (u_ URLResponse) SetSuggestedFilename(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setSuggestedFilename:"), value)
}/* debug [instance_properties/setter]: suggestedFilename */


// The name of the text encoding provided by the response’s originating source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/textencodingname
func (u_ URLResponse) TextEncodingName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("textEncodingName"))
	return rv
}/* debug [instance_properties/getter]: textEncodingName */


// The name of the text encoding provided by the response’s originating source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/textencodingname
func (u_ URLResponse) SetTextEncodingName(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setTextEncodingName:"), value)
}/* debug [instance_properties/setter]: textEncodingName */


// The URL for the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/url
func (u_ URLResponse) Url() IURL {
	rv := objc.Send[URL](u_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// The URL for the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlresponse/url
func (u_ URLResponse) SetUrl(value IURL) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUrl:"), value)
}/* debug [instance_properties/setter]: url */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLResponse */



