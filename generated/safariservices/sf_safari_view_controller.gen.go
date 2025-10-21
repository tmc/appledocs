// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [SFSafariViewController] class.
type ISFSafariViewController interface {
	appkit.IViewController
}

// An object that provides a visible standard interface for browsing the web.
//
// An object presents a self-contained web interface inside your app. Present this view controller to let people view websites from anywhere on the internet without leaving your app. The web interface supports Safari features such as Reader, AutoFill, Fraudulent Website Warning, and content blocking. Interactions with the web interface aren’t visible to your app, and you can’t access AutoFill data, browsing history, or website data. You don’t need to secure data between your app and Safari. To share data between your app and Safari, use instead. Present an when you don’t need to customize or interact with the web content. After you present the content, interactions with the web content occur solely within the view controller. When the person dismisses the view controller, control returns to your app’s interface. If you need to customize the controls of the web interface, or you want to interact with content in that interface, display the content using a object instead. UI features include the following: A read-only address field with a security indicator and a Reader button An Action button that invokes an activity view controller offering custom services from your app and activities, such as messaging, from the system and other extensions A Done button, back and forward navigation buttons, and a button to open the page directly in Safari Peek and Pop for links and detected data using 3D Touch When a person peeks and pops a link in , the view controller loads and displays the link destination. When a person peeks and pops a link in a class, the web view opens the link in Safari by default.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController
type SFSafariViewController struct {
	appkit.ViewController
}

// SFSafariViewControllerFrom constructs a [SFSafariViewController] from an unsafe.Pointer.
//
// An object that provides a visible standard interface for browsing the web.
func SFSafariViewControllerFrom(ptr unsafe.Pointer) SFSafariViewController {
	return SFSafariViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SFSafariViewControllerClass) Alloc() SFSafariViewController {
	rv := objc.Send[SFSafariViewController](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a Safari view controller that loads the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:)
func NewSFSafariViewControllerWithURL(URL foundation.IURL) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:"), URL)
	rv.Autorelease()
	return rv
}



// Initializes and configures a Safari view controller that loads the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:configuration:)
func NewSFSafariViewControllerWithURLConfiguration(URL foundation.IURL, configuration ISFSafariViewControllerConfiguration) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:configuration:"), URL, configuration)
	rv.Autorelease()
	return rv
}



// Initializes a Safari view controller that will load the specified URL, entering Reader mode if Reader mode is requested and available.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/init(url:entersReaderIfAvailable:)
func NewSFSafariViewControllerWithURLEntersReaderIfAvailable(URL foundation.IURL, entersReaderIfAvailable bool) SFSafariViewController {
	instance := getSFSafariViewControllerClass().Alloc()
	rv := objc.Send[SFSafariViewController](instance.ID, objc.Sel("initWithURL:entersReaderIfAvailable:"), URL, entersReaderIfAvailable)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/prewarmConnections(to:)
func (sc _SFSafariViewControllerClass) PrewarmConnectionsToURLs(URLs []foundation.IURL) SFSafariViewControllerPrewarmingToken {
	rv := objc.Send[SFSafariViewControllerPrewarmingToken](objc.ID(sc.class), objc.Sel("prewarmConnectionsToURLs:"), URLs)
	return rv
}

// A copy of the Safari view controller’s initialized configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/configuration-swift.property
func (s_ SFSafariViewController) Configuration() SFSafariViewControllerConfiguration {
	rv := objc.Send[SFSafariViewControllerConfiguration](s_.ID, objc.Sel("configuration"))
	return rv
}

// An object that provides behavior for the Safari view controller’s Done and Action buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/delegate
func (s_ SFSafariViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// An object that provides behavior for the Safari view controller’s Done and Action buttons.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/delegate
func (s_ SFSafariViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// The style of dismiss button to use in the navigation bar to close the Safari view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/dismissButtonStyle-swift.property
func (s_ SFSafariViewController) DismissButtonStyle() SFSafariViewControllerDismissButtonStyle {
	rv := objc.Send[SFSafariViewControllerDismissButtonStyle](s_.ID, objc.Sel("dismissButtonStyle"))
	return rv
}


// SetDismissButtonStyle sets the value of the dismissButtonStyle property.
// The style of dismiss button to use in the navigation bar to close the Safari view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/dismissButtonStyle-swift.property
func (s_ SFSafariViewController) SetDismissButtonStyle(value SFSafariViewControllerDismissButtonStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDismissButtonStyle:"), value)
}

// The color to tint the background of the navigation bar and the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredBarTintColor
func (s_ SFSafariViewController) PreferredBarTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("preferredBarTintColor"))
	return rv
}


// SetPreferredBarTintColor sets the value of the preferredBarTintColor property.
// The color to tint the background of the navigation bar and the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredBarTintColor
func (s_ SFSafariViewController) SetPreferredBarTintColor(value appkit.IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredBarTintColor:"), value)
}

// The color to tint the control buttons on the navigation bar and the toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredControlTintColor
func (s_ SFSafariViewController) PreferredControlTintColor() appkit.Color {
	rv := objc.Send[appkit.Color](s_.ID, objc.Sel("preferredControlTintColor"))
	return rv
}


// SetPreferredControlTintColor sets the value of the preferredControlTintColor property.
// The color to tint the control buttons on the navigation bar and the toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariViewController/preferredControlTintColor
func (s_ SFSafariViewController) SetPreferredControlTintColor(value appkit.IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPreferredControlTintColor:"), value)
}

// An object you use to send tap event attribution data to the browser for Private Click Measurement.
//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/configuration-swift.class/eventattribution
func (s_ SFSafariViewController) EventAttribution() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("eventAttribution"))
	return rv
}


// SetEventAttribution sets the value of the eventAttribution property.
// An object you use to send tap event attribution data to the browser for Private Click Measurement.

//
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfsafariviewcontroller/configuration-swift.class/eventattribution
func (s_ SFSafariViewController) SetEventAttribution(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEventAttribution:"), value)
}


