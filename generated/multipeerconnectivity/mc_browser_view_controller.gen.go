// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [MCBrowserViewController] class.
type IMCBrowserViewController interface {
	appkit.IViewController
}

// The class presents nearby devices to the user and enables the user to invite nearby devices to a session. To use this class in iOS or tvOS, call methods from the underlying class ( and for storyboards or and for nib-based views) to present and dismiss the view controller. In macOS, use the comparable methods and instead.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MCBrowserViewControllerClass) Alloc() MCBrowserViewController {
	rv := objc.Send[MCBrowserViewController](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes a browser view controller with the provided browser and session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/init(browser:session:)
func NewMCBrowserViewControllerWithBrowserSession(browser unsafe.Pointer, session unsafe.Pointer) MCBrowserViewController {
	instance := getMCBrowserViewControllerClass().Alloc()
	rv := objc.Send[MCBrowserViewController](instance.ID, objc.Sel("initWithBrowser:session:"), browser, session)
	rv.Autorelease()
	return rv
}

// Initializes a browser view controller using the provided service type and session.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/init(serviceType:session:)
func NewMCBrowserViewControllerWithServiceTypeSession(serviceType string, session unsafe.Pointer) MCBrowserViewController {
	instance := getMCBrowserViewControllerClass().Alloc()
	rv := objc.Send[MCBrowserViewController](instance.ID, objc.Sel("initWithServiceType:session:"), objc.String(serviceType), session)
	rv.Autorelease()
	return rv
}


// The browser object that is used for discovering peers.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/browser
func (m_ MCBrowserViewController) Browser() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("browser"))
	return rv
}

// The delegate object that handles browser-view-controller-related events.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/delegate
func (m_ MCBrowserViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate object that handles browser-view-controller-related events.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/delegate
func (m_ MCBrowserViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}
// The maximum number of peers allowed in a session, including the local peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/maximumNumberOfPeers
func (m_ MCBrowserViewController) MaximumNumberOfPeers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("maximumNumberOfPeers"))
	return rv
}


// SetMaximumNumberOfPeers sets the value of the maximumNumberOfPeers property.
// The maximum number of peers allowed in a session, including the local peer.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/maximumNumberOfPeers
func (m_ MCBrowserViewController) SetMaximumNumberOfPeers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaximumNumberOfPeers:"), value)
}
// The minimum number of peers that need to be in a session, including the local peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/minimumNumberOfPeers
func (m_ MCBrowserViewController) MinimumNumberOfPeers() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("minimumNumberOfPeers"))
	return rv
}


// SetMinimumNumberOfPeers sets the value of the minimumNumberOfPeers property.
// The minimum number of peers that need to be in a session, including the local peer.

//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/minimumNumberOfPeers
func (m_ MCBrowserViewController) SetMinimumNumberOfPeers(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinimumNumberOfPeers:"), value)
}
// The multipeer session to which the invited peers are connected.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCBrowserViewController/session
func (m_ MCBrowserViewController) Session() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("session"))
	return rv
}


