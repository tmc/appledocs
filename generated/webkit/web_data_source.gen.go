// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebDataSource */


/* debug [class_header]: Header for WebDataSource */
// The class instance for the [WebDataSource] class.
var (
	WebDataSourceClass     _WebDataSourceClass
	WebDataSourceClassOnce sync.Once
)

func getWebDataSourceClass() _WebDataSourceClass {
	WebDataSourceClassOnce.Do(func() {
		WebDataSourceClass = _WebDataSourceClass{objc.GetClass("WebDataSource")}
	})
	return WebDataSourceClass
}

type _WebDataSourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebDataSource */
// An interface definition for the [WebDataSource] class.
type IWebDataSource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebDataSource */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	InitialRequest() foundation.URLRequest
	Loading() bool
	MainResource() IWebResource
	PageTitle() objc.IObject /* cross-framework: NSString */
	Representation() unsafe.Pointer
	Request() foundation.MutableURLRequest
	Response() foundation.URLResponse
	Subresources() objc.IObject /* cross-framework: NSArray */
	TextEncodingName() objc.IObject /* cross-framework: NSString */
	UnreachableURL() objc.IObject /* cross-framework: NSURL */
	WebArchive() IWebArchive
	WebFrame() IWebFrame
	IsLoading() bool
	SetIsLoading(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebDataSource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebDataSource */
// Alloc allocates a new instance without initialization.
func (wc _WebDataSourceClass) Alloc() WebDataSource {
	rv := objc.Send[WebDataSource](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebDataSourceClass) New() WebDataSource {
	rv := objc.Send[WebDataSource](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebDataSource) Init() WebDataSource {
	rv := objc.Send[WebDataSource](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebDataSource) Autorelease() WebDataSource {
	rv := objc.Send[WebDataSource](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebDataSource creates a new WebDataSource instance.
func NewWebDataSource() WebDataSource {
	return getWebDataSourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebDataSource */
// encapsulates the web content to be displayed in a web frame view. A object has a representation object, conforming to the protocol, that holds the data in an appropriate format depending on the MIME type. You can extend WebKit to support new MIME types by implementing your own view and representation classes, and specifying the mapping between them using the class method.
//
// objects have an associated initial request, possibly a modified request, and a response object. Since the data source may be in the process of being loaded, you should check the state of a data source using before accessing its data. Use to get the raw data. Use the method to get the actual representation object and query it for more details.


// encapsulates the web content to be displayed in a web frame view. A object has a representation object, conforming to the protocol, that holds the data in an appropriate format depending on the MIME type. You can extend WebKit to support new MIME types by implementing your own view and representation classes, and specifying the mapping between them using the class method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource
type WebDataSource struct {
	objectivec.Object
}

// WebDataSourceFrom constructs a [WebDataSource] from an unsafe.Pointer.
//
// encapsulates the web content to be displayed in a web frame view. A object has a representation object, conforming to the protocol, that holds the data in an appropriate format depending on the MIME type. You can extend WebKit to support new MIME types by implementing your own view and representation classes, and specifying the mapping between them using the class method.
func WebDataSourceFrom(ptr unsafe.Pointer) WebDataSource {
	return WebDataSource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebDataSource */

// initializes a data source with a URL request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/init(request:)
func NewWebDataSourceWithRequest(request foundation.URLRequest) WebDataSource {
	instance := getWebDataSourceClass().Alloc()
	rv := objc.Send[WebDataSource](instance.ID, objc.Sel("initWithRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebDataSourceWithRequest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebDataSource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebDataSource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebDataSource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebDataSource */

// The raw data that represents the data source’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/data
func (w_ WebDataSource) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](w_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// A reference to the original request that was used to load the web content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/initialRequest
func (w_ WebDataSource) InitialRequest() foundation.URLRequest {
	rv := objc.Send[foundation.URLRequest](w_.ID, objc.Sel("initialRequest"))
	return rv
}/* debug [instance_properties/getter]: initialRequest */


// A Boolean that indicates whether the data source is loading its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/isLoading
func (w_ WebDataSource) Loading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("loading"))
	return rv
}/* debug [instance_properties/getter]: loading */


// A object representing the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/mainResource
func (w_ WebDataSource) MainResource() IWebResource {
	rv := objc.Send[WebResource](w_.ID, objc.Sel("mainResource"))
	return rv
}/* debug [instance_properties/getter]: mainResource */


// The title of the data source’s page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/pageTitle
func (w_ WebDataSource) PageTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("pageTitle"))
	return rv
}/* debug [instance_properties/getter]: pageTitle */


// The data source’s representation depending on its MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/representation
func (w_ WebDataSource) Representation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("representation"))
	return rv
}/* debug [instance_properties/getter]: representation */


// The request that was used to create the data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/request
func (w_ WebDataSource) Request() foundation.MutableURLRequest {
	rv := objc.Send[foundation.MutableURLRequest](w_.ID, objc.Sel("request"))
	return rv
}/* debug [instance_properties/getter]: request */


// The response for this data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/response
func (w_ WebDataSource) Response() foundation.URLResponse {
	rv := objc.Send[foundation.URLResponse](w_.ID, objc.Sel("response"))
	return rv
}/* debug [instance_properties/getter]: response */


// The data source’s subresources that have finished downloading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/subresources
func (w_ WebDataSource) Subresources() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("subresources"))
	return rv
}/* debug [instance_properties/getter]: subresources */


// The text encoding for the data source’s web view, if set, or the text encoding of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/textEncodingName
func (w_ WebDataSource) TextEncodingName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("textEncodingName"))
	return rv
}/* debug [instance_properties/getter]: textEncodingName */


// The data source’s unreachable URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/unreachableURL
func (w_ WebDataSource) UnreachableURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("unreachableURL"))
	return rv
}/* debug [instance_properties/getter]: unreachableURL */


// A web archive representing the data source, its subresources, and subframes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/webArchive
func (w_ WebDataSource) WebArchive() IWebArchive {
	rv := objc.Send[WebArchive](w_.ID, objc.Sel("webArchive"))
	return rv
}/* debug [instance_properties/getter]: webArchive */


// The web frame that represents this data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDataSource/webFrame
func (w_ WebDataSource) WebFrame() IWebFrame {
	rv := objc.Send[WebFrame](w_.ID, objc.Sel("webFrame"))
	return rv
}/* debug [instance_properties/getter]: webFrame */


// A Boolean that indicates whether the data source is loading its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdatasource/isloading
func (w_ WebDataSource) IsLoading() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLoading"))
	return rv
}/* debug [instance_properties/getter]: isLoading */


// A Boolean that indicates whether the data source is loading its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdatasource/isloading
func (w_ WebDataSource) SetIsLoading(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLoading:"), value)
}/* debug [instance_properties/setter]: isLoading */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebDataSource */


