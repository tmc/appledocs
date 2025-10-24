// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebResource */


/* debug [class_header]: Header for WebResource */
// The class instance for the [WebResource] class.
var (
	WebResourceClass     _WebResourceClass
	WebResourceClassOnce sync.Once
)

func getWebResourceClass() _WebResourceClass {
	WebResourceClassOnce.Do(func() {
		WebResourceClass = _WebResourceClass{objc.GetClass("WebResource")}
	})
	return WebResourceClass
}

type _WebResourceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebResource */
// An interface definition for the [WebResource] class.
type IWebResource interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebResource */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	FrameName() objc.IObject /* cross-framework: NSString */
	MIMEType() objc.IObject /* cross-framework: NSString */
	TextEncodingName() objc.IObject /* cross-framework: NSString */
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebResource */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebResource */
// Alloc allocates a new instance without initialization.
func (wc _WebResourceClass) Alloc() WebResource {
	rv := objc.Send[WebResource](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebResourceClass) New() WebResource {
	rv := objc.Send[WebResource](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebResource) Init() WebResource {
	rv := objc.Send[WebResource](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebResource) Autorelease() WebResource {
	rv := objc.Send[WebResource](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebResource creates a new WebResource instance.
func NewWebResource() WebResource {
	return getWebResourceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebResource */
// A object represents a downloaded URL. It encapsulates the data of the download as well as other resource properties such as the URL, MIME type, and frame name.
//
// Use the method to initialize a newly created object. Use the other methods in this class to get the properties of a object.


// A object represents a downloaded URL. It encapsulates the data of the download as well as other resource properties such as the URL, MIME type, and frame name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource
type WebResource struct {
	objectivec.Object
}

// WebResourceFrom constructs a [WebResource] from an unsafe.Pointer.
//
// A object represents a downloaded URL. It encapsulates the data of the download as well as other resource properties such as the URL, MIME type, and frame name.
func WebResourceFrom(ptr unsafe.Pointer) WebResource {
	return WebResource{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebResource */

// Initializes and returns a web resource instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/init(data:url:mimeType:textEncodingName:frameName:)
func NewWebResourceWithDataURLMIMETypeTextEncodingNameFrameName(data objc.IObject /* cross-framework: NSData */, URL objc.IObject /* cross-framework: NSURL */, MIMEType objc.IObject /* cross-framework: NSString */, textEncodingName objc.IObject /* cross-framework: NSString */, frameName objc.IObject /* cross-framework: NSString */) WebResource {
	instance := getWebResourceClass().Alloc()
	rv := objc.Send[WebResource](instance.ID, objc.Sel("initWithData:URL:MIMEType:textEncodingName:frameName:"), data, URL, MIMEType, textEncodingName, frameName)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWebResourceWithDataURLMIMETypeTextEncodingNameFrameName */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebResource */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebResource */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebResource */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebResource */

// The receiver’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/data
func (w_ WebResource) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](w_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The name of the frame. If the receiver does not represent the contents of an entire HTML frame, this is .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/frameName
func (w_ WebResource) FrameName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("frameName"))
	return rv
}/* debug [instance_properties/getter]: frameName */


// The receiver’s MIME type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/mimeType
func (w_ WebResource) MIMEType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("MIMEType"))
	return rv
}/* debug [instance_properties/getter]: MIMEType */


// The receiver’s text encoding name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/textEncodingName
func (w_ WebResource) TextEncodingName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("textEncodingName"))
	return rv
}/* debug [instance_properties/getter]: textEncodingName */


// The receiver’s URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebResource/url
func (w_ WebResource) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebResource */


