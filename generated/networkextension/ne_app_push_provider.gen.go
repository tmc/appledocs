// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NEAppPushProvider] class.
var (
	nEAppPushProviderClass     _NEAppPushProviderClass
	nEAppPushProviderClassOnce sync.Once
)

func getNEAppPushProviderClass() _NEAppPushProviderClass {
	nEAppPushProviderClassOnce.Do(func() {
		nEAppPushProviderClass = _NEAppPushProviderClass{objc.GetClass("NEAppPushProvider")}
	})
	return nEAppPushProviderClass
}

type _NEAppPushProviderClass struct {
	class objc.Class
}

// An interface definition for the [NEAppPushProvider] class.
type INEAppPushProvider interface {
	INEProvider
	ReportIncomingCallWithUserInfo(userInfo unsafe.Pointer)
	StartWithCompletionHandler(completionHandler unsafe.Pointer)
}

// An object that creates and maintains a persistent network connection to a local push server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider
type NEAppPushProvider struct {
	NEProvider
}

// NEAppPushProviderFrom constructs a [NEAppPushProvider] from an unsafe.Pointer.
//
// An object that creates and maintains a persistent network connection to a local push server.
func NEAppPushProviderFrom(ptr unsafe.Pointer) NEAppPushProvider {
	return NEAppPushProvider{
		NEProvider: NEProviderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (nc _NEAppPushProviderClass) Alloc() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEAppPushProviderClass) New() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppPushProvider) Init() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppPushProvider) Autorelease() NEAppPushProvider {
	rv := objc.Send[NEAppPushProvider](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppPushProvider creates a new NEAppPushProvider instance.
func NewNEAppPushProvider() NEAppPushProvider {
	return getNEAppPushProviderClass().New()
}


// Informs the manager about an incoming call. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/reportIncomingCall(userInfo:)
func (n_ NEAppPushProvider) ReportIncomingCallWithUserInfo(userInfo unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reportIncomingCallWithUserInfo:"), userInfo)
}
// Indicates that the framework has started the provider, and provides a completion handler for subclasses to signal their readiness. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/start(completionHandler:)
func (n_ NEAppPushProvider) StartWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("startWithCompletionHandler:"), completionHandler)
}


