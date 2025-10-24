// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MCBrowserViewController */


/* debug [class_header]: Header for MCBrowserViewController */
// The class instance for the [MCBrowserViewController] class.
var (
	MCBrowserViewControllerClass     _MCBrowserViewControllerClass
	MCBrowserViewControllerClassOnce sync.Once
)

func getMCBrowserViewControllerClass() _MCBrowserViewControllerClass {
	MCBrowserViewControllerClassOnce.Do(func() {
		MCBrowserViewControllerClass = _MCBrowserViewControllerClass{objc.GetClass("MCBrowserViewController")}
	})
	return MCBrowserViewControllerClass
}

type _MCBrowserViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MCBrowserViewController */
// An interface definition for the [MCBrowserViewController] class.
type IMCBrowserViewController interface {
	appkit.IViewController
	
/* debug [class_interface_properties]: Properties for MCBrowserViewController */
	// properties:
	Browser() IMCNearbyServiceBrowser
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	MaximumNumberOfPeers() uint
	SetMaximumNumberOfPeers(value uint)
	MinimumNumberOfPeers() uint
	SetMinimumNumberOfPeers(value uint)
	Session() IMCSession
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MCBrowserViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MCBrowserViewController */
// Alloc allocates a new instance without initialization.
func (mc _MCBrowserViewControllerClass) Alloc() MCBrowserViewController {
	rv := objc.Send[MCBrowserViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MCBrowserViewControllerClass) New() MCBrowserViewController {
	rv := objc.Send[MCBrowserViewController](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCBrowserViewController) Init() MCBrowserViewController {
	rv := objc.Send[MCBrowserViewController](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCBrowserViewController) Autorelease() MCBrowserViewController {
	rv := objc.Send[MCBrowserViewController](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCBrowserViewController creates a new MCBrowserViewController instance.
func NewMCBrowserViewController() MCBrowserViewController {
	return getMCBrowserViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MCBrowserViewController */
// The class presents nearby devices to the user and enables the user to invite nearby devices to a session. To use this class in iOS or tvOS, call methods from the underlying class ( and for storyboards or and for nib-based views) to present and dismiss the view controller. In macOS, use the comparable methods and instead.


// The class presents nearby devices to the user and enables the user to invite nearby devices to a session. To use this class in iOS or tvOS, call methods from the underlying class ( and for storyboards or and for nib-based views) to present and dismiss the view controller. In macOS, use the comparable methods and instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController
type MCBrowserViewController struct {
	appkit.ViewController
}

// MCBrowserViewControllerFrom constructs a [MCBrowserViewController] from an unsafe.Pointer.
//
// The class presents nearby devices to the user and enables the user to invite nearby devices to a session. To use this class in iOS or tvOS, call methods from the underlying class ( and for storyboards or and for nib-based views) to present and dismiss the view controller. In macOS, use the comparable methods and instead.
func MCBrowserViewControllerFrom(ptr unsafe.Pointer) MCBrowserViewController {
	return MCBrowserViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MCBrowserViewController */

// Initializes a browser view controller with the provided browser and session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/init(browser:session:)
func NewMCBrowserViewControllerWithBrowserSession(browser IMCNearbyServiceBrowser, session IMCSession) MCBrowserViewController {
	instance := getMCBrowserViewControllerClass().Alloc()
	rv := objc.Send[MCBrowserViewController](instance.ID, objc.Sel("initWithBrowser:session:"), browser, session)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCBrowserViewControllerWithBrowserSession */


// Initializes a browser view controller using the provided service type and session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/init(serviceType:session:)
func NewMCBrowserViewControllerWithServiceTypeSession(serviceType objc.IObject /* cross-framework: NSString */, session IMCSession) MCBrowserViewController {
	instance := getMCBrowserViewControllerClass().Alloc()
	rv := objc.Send[MCBrowserViewController](instance.ID, objc.Sel("initWithServiceType:session:"), serviceType, session)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMCBrowserViewControllerWithServiceTypeSession */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MCBrowserViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MCBrowserViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MCBrowserViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MCBrowserViewController */

// The browser object that is used for discovering peers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/browser
func (m_ MCBrowserViewController) Browser() IMCNearbyServiceBrowser {
	rv := objc.Send[MCNearbyServiceBrowser](m_.ID, objc.Sel("browser"))
	return rv
}/* debug [instance_properties/getter]: browser */


// The delegate object that handles browser-view-controller-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/delegate
func (m_ MCBrowserViewController) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate object that handles browser-view-controller-related events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/delegate
func (m_ MCBrowserViewController) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The maximum number of peers allowed in a session, including the local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/maximumNumberOfPeers
func (m_ MCBrowserViewController) MaximumNumberOfPeers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maximumNumberOfPeers"))
	return rv
}/* debug [instance_properties/getter]: maximumNumberOfPeers */


// The maximum number of peers allowed in a session, including the local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/maximumNumberOfPeers
func (m_ MCBrowserViewController) SetMaximumNumberOfPeers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumNumberOfPeers:"), value)
}/* debug [instance_properties/setter]: maximumNumberOfPeers */


// The minimum number of peers that need to be in a session, including the local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/minimumNumberOfPeers
func (m_ MCBrowserViewController) MinimumNumberOfPeers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("minimumNumberOfPeers"))
	return rv
}/* debug [instance_properties/getter]: minimumNumberOfPeers */


// The minimum number of peers that need to be in a session, including the local peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/minimumNumberOfPeers
func (m_ MCBrowserViewController) SetMinimumNumberOfPeers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumNumberOfPeers:"), value)
}/* debug [instance_properties/setter]: minimumNumberOfPeers */


// The multipeer session to which the invited peers are connected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/session
func (m_ MCBrowserViewController) Session() IMCSession {
	rv := objc.Send[MCSession](m_.ID, objc.Sel("session"))
	return rv
}/* debug [instance_properties/getter]: session */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MCBrowserViewController */


