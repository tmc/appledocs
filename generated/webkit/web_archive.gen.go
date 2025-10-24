// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WebArchive */

/* debug [class_header]: Header for WebArchive */
// The class instance for the [WebArchive] class.
var (
	WebArchiveClass     _WebArchiveClass
	WebArchiveClassOnce sync.Once
)

func getWebArchiveClass() _WebArchiveClass {
	WebArchiveClassOnce.Do(func() {
		WebArchiveClass = _WebArchiveClass{objc.GetClass("WebArchive")}
	})
	return WebArchiveClass
}

type _WebArchiveClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebArchive */
// An interface definition for the [WebArchive] class.
type IWebArchive interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebArchive */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	MainResource() IWebResource
	SubframeArchives() objc.IObject     /* cross-framework: NSArray */
	Subresources() objc.IObject         /* cross-framework: NSArray */
	WebArchivePboardType() objc.IObject /* cross-framework: NSString */
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebArchive */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebArchive */
// Alloc allocates a new instance without initialization.
func (wc _WebArchiveClass) Alloc() WebArchive {
	rv := objc.Send[WebArchive](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebArchiveClass) New() WebArchive {
	rv := objc.Send[WebArchive](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebArchive) Init() WebArchive {
	rv := objc.Send[WebArchive](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebArchive) Autorelease() WebArchive {
	rv := objc.Send[WebArchive](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebArchive creates a new WebArchive instance.
func NewWebArchive() WebArchive {
	return getWebArchiveClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebArchive */
// A WebArchive object represents a webpage that can be archived—for example, archived on disk or on the pasteboard. A WebArchive object contains the main resource, as well as the subresources and subframes of the main resource. The main resource can be an entire webpage, a portion of a webpage, or some other kind of data such as an image. Use this class to archive webpages, or place a portion of a webpage on the pasteboard, or to represent rich web content in any application.

// A WebArchive object represents a webpage that can be archived—for example, archived on disk or on the pasteboard. A WebArchive object contains the main resource, as well as the subresources and subframes of the main resource. The main resource can be an entire webpage, a portion of a webpage, or some other kind of data such as an image. Use this class to archive webpages, or place a portion of a webpage on the pasteboard, or to represent rich web content in any application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive
type WebArchive struct {
	objectivec.Object
}

// WebArchiveFrom constructs a [WebArchive] from an unsafe.Pointer.
//
// A WebArchive object represents a webpage that can be archived—for example, archived on disk or on the pasteboard. A WebArchive object contains the main resource, as well as the subresources and subframes of the main resource. The main resource can be an entire webpage, a portion of a webpage, or some other kind of data such as an image. Use this class to archive webpages, or place a portion of a webpage on the pasteboard, or to represent rich web content in any application.
func WebArchiveFrom(ptr unsafe.Pointer) WebArchive {
	return WebArchive{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebArchive */

// Initializes and returns the receiver, specifying the initial content data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/init(data:)
func NewWebArchiveWithData(data objc.IObject /* cross-framework: NSData */) WebArchive {
	instance := getWebArchiveClass().Alloc()
	rv := objc.Send[WebArchive](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewWebArchiveWithData */

// Initializes the receiver with a resource and optional subresources and subframe archives..
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/init(mainResource:subresources:subframeArchives:)
func NewWebArchiveWithMainResourceSubresourcesSubframeArchives(mainResource IWebResource, subresources objc.IObject /* cross-framework: NSArray */, subframeArchives objc.IObject /* cross-framework: NSArray */) WebArchive {
	instance := getWebArchiveClass().Alloc()
	rv := objc.Send[WebArchive](instance.ID, objc.Sel("initWithMainResource:subresources:subframeArchives:"), mainResource, subresources, subframeArchives)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewWebArchiveWithMainResourceSubresourcesSubframeArchives */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebArchive */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebArchive */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebArchive */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebArchive */

// The data representation of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/data
func (w_ WebArchive) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](w_.ID, objc.Sel("data"))
	return rv
} /* debug [instance_properties/getter]: data */

// The receiver’s main resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/mainResource
func (w_ WebArchive) MainResource() IWebResource {
	rv := objc.Send[WebResource](w_.ID, objc.Sel("mainResource"))
	return rv
} /* debug [instance_properties/getter]: mainResource */

// Archives representing the receiver’s subresources or if there are none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/subframeArchives
func (w_ WebArchive) SubframeArchives() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("subframeArchives"))
	return rv
} /* debug [instance_properties/getter]: subframeArchives */

// The receiver’s subresources, or if there are none.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebArchive/subresources
func (w_ WebArchive) Subresources() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](w_.ID, objc.Sel("subresources"))
	return rv
} /* debug [instance_properties/getter]: subresources */

// The pasteboard type constant used when adding or accessing a WebArchive on the pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/webarchivepboardtype
func (w_ WebArchive) WebArchivePboardType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("WebArchivePboardType"))
	return rv
} /* debug [instance_properties/getter]: WebArchivePboardType */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WebArchive */
