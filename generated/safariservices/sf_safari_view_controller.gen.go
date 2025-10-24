// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class SFSafariViewController */


/* debug [class_header]: Header for SFSafariViewController */
// The class instance for the [SFSafariViewController] class.
var (
	SFSafariViewControllerClass     _SFSafariViewControllerClass
	SFSafariViewControllerClassOnce sync.Once
)

func getSFSafariViewControllerClass() _SFSafariViewControllerClass {
	SFSafariViewControllerClassOnce.Do(func() {
		SFSafariViewControllerClass = _SFSafariViewControllerClass{objc.GetClass("SFSafariViewController")}
	})
	return SFSafariViewControllerClass
}

type _SFSafariViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariViewController */
// An interface definition for the [SFSafariViewController] class.
type ISFSafariViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for SFSafariViewController */
	// properties:
	EventAttribution() EventAttribution /* not a class type */
	SetEventAttribution(value EventAttribution /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariViewController */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerClass) Alloc() SFSafariViewController {
	rv := objc.Send[SFSafariViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariViewControllerClass) New() SFSafariViewController {
	rv := objc.Send[SFSafariViewController](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariViewController) Init() SFSafariViewController {
	rv := objc.Send[SFSafariViewController](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariViewController) Autorelease() SFSafariViewController {
	rv := objc.Send[SFSafariViewController](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariViewController creates a new SFSafariViewController instance.
func NewSFSafariViewController() SFSafariViewController {
	return getSFSafariViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariViewController */
// An object that provides a visible standard interface for browsing the web.
//
// An object presents a self-contained web interface inside your app. Present this view controller to let people view websites from anywhere on the internet without leaving your app. The web interface supports Safari features such as Reader, AutoFill, Fraudulent Website Warning, and content blocking. Interactions with the web interface aren’t visible to your app, and you can’t access AutoFill data, browsing history, or website data. You don’t need to secure data between your app and Safari. To share data between your app and Safari, use instead. Present an when you don’t need to customize or interact with the web content. After you present the content, interactions with the web content occur solely within the view controller. When the person dismisses the view controller, control returns to your app’s interface. If you need to customize the controls of the web interface, or you want to interact with content in that interface, display the content using a object instead. UI features include the following: A read-only address field with a security indicator and a Reader button An Action button that invokes an activity view controller offering custom services from your app and activities, such as messaging, from the system and other extensions A Done button, back and forward navigation buttons, and a button to open the page directly in Safari Peek and Pop for links and detected data using 3D Touch When a person peeks and pops a link in , the view controller loads and displays the link destination. When a person peeks and pops a link in a class, the web view opens the link in Safari by default.


// An object that provides a visible standard interface for browsing the web.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController
type SFSafariViewController struct {
	ViewController
}

// SFSafariViewControllerFrom constructs a [SFSafariViewController] from an unsafe.Pointer.
//
// An object that provides a visible standard interface for browsing the web.
func SFSafariViewControllerFrom(ptr unsafe.Pointer) SFSafariViewController {
	return SFSafariViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariViewController */

// Initializes a Safari view controller that loads the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:)
func NewSFSafariViewControllerWithURL(URL objc.IObject /* cross-framework: NSURL */) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSafariViewControllerWithURL */


// Initializes and configures a Safari view controller that loads the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:configuration:)
func NewSFSafariViewControllerWithURLConfiguration(URL objc.IObject /* cross-framework: NSURL */, configuration ISFSafariViewControllerConfiguration) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:configuration:"), URL, configuration)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSafariViewControllerWithURLConfiguration */


// Initializes a Safari view controller that will load the specified URL, entering Reader mode if Reader mode is requested and available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:entersReaderIfAvailable:)
func NewSFSafariViewControllerWithURLEntersReaderIfAvailable(URL objc.IObject /* cross-framework: NSURL */, entersReaderIfAvailable bool) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:entersReaderIfAvailable:"), URL, entersReaderIfAvailable)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSFSafariViewControllerWithURLEntersReaderIfAvailable */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/prewarmConnections(to:)
func (sc _SFSafariViewControllerClass) PrewarmConnectionsToURLs(URLs []foundation.URL) ISFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("prewarmConnectionsToURLs:"), URLs)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrewarmConnectionsToURLs) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariViewController */

// An object you use to send tap event attribution data to the browser for Private Click Measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/configuration-swift.class/eventattribution
func (s_ SFSafariViewController) EventAttribution() EventAttribution /* not a class type */ {
	rv := objc.Send[EventAttribution](s_.ID, objc.Sel("eventAttribution"))
	return rv
}/* debug [instance_properties/getter]: eventAttribution */


// An object you use to send tap event attribution data to the browser for Private Click Measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/configuration-swift.class/eventattribution
func (s_ SFSafariViewController) SetEventAttribution(value EventAttribution /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEventAttribution:"), value)
}/* debug [instance_properties/setter]: eventAttribution */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariViewController */


