// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class WebDownload */


/* debug [class_header]: Header for WebDownload */
// The class instance for the [WebDownload] class.
var (
	WebDownloadClass     _WebDownloadClass
	WebDownloadClassOnce sync.Once
)

func getWebDownloadClass() _WebDownloadClass {
	WebDownloadClassOnce.Do(func() {
		WebDownloadClass = _WebDownloadClass{objc.GetClass("WebDownload")}
	})
	return WebDownloadClass
}

type _WebDownloadClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebDownload */
// An interface definition for the [WebDownload] class.
type IWebDownload interface {
	foundation.IURLDownload
	
/* debug [class_interface_properties]: Properties for WebDownload */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebDownload */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebDownload */
// Alloc allocates a new instance without initialization.
func (wc _WebDownloadClass) Alloc() WebDownload {
	rv := objc.Send[WebDownload](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebDownloadClass) New() WebDownload {
	rv := objc.Send[WebDownload](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebDownload) Init() WebDownload {
	rv := objc.Send[WebDownload](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebDownload) Autorelease() WebDownload {
	rv := objc.Send[WebDownload](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebDownload creates a new WebDownload instance.
func NewWebDownload() WebDownload {
	return getWebDownloadClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebDownload */
// objects initiate download client requests on behalf of a delegate. A download request involves loading the data, decoding it (if necessary), and saving it to a file. Instances of this class behave similar to except delegates of may implement an additional delegate method. The method allows the delegate to specify the window to be used for authentication sheets. If the delegate does not implement this method, the object will prompt the user for authentication using the standard WebKit authentication panel, as either a sheet or window. There are no additional methods defined in this class. See for the delegate method.


// objects initiate download client requests on behalf of a delegate. A download request involves loading the data, decoding it (if necessary), and saving it to a file. Instances of this class behave similar to except delegates of may implement an additional delegate method. The method allows the delegate to specify the window to be used for authentication sheets. If the delegate does not implement this method, the object will prompt the user for authentication using the standard WebKit authentication panel, as either a sheet or window. There are no additional methods defined in this class. See for the delegate method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WebDownload
type WebDownload struct {
	foundation.URLDownload
}

// WebDownloadFrom constructs a [WebDownload] from an unsafe.Pointer.
//
// objects initiate download client requests on behalf of a delegate. A download request involves loading the data, decoding it (if necessary), and saving it to a file. Instances of this class behave similar to except delegates of may implement an additional delegate method. The method allows the delegate to specify the window to be used for authentication sheets. If the delegate does not implement this method, the object will prompt the user for authentication using the standard WebKit authentication panel, as either a sheet or window. There are no additional methods defined in this class. See for the delegate method.
func WebDownloadFrom(ptr unsafe.Pointer) WebDownload {
	return WebDownload{
		URLDownload: foundation.URLDownloadFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebDownload *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebDownload */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebDownload */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebDownload */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebDownload */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WebDownload */



