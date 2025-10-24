//go:build darwin && ios

// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/gamekit"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NISession


// Stops a running session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/invalidate()
func (n_ NISession) Invalidate() {
	objc.Send[objc.ID](n_.ID, objc.Sel("invalidate"))
}

// Stops sending distance and direction updates to the peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/pause()
func (n_ NISession) Pause() {
	objc.Send[objc.ID](n_.ID, objc.Sel("pause"))
}

// Starts a session with a nearby peer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/run(_:)
func (n_ NISession) RunWithConfiguration(configuration INIConfiguration) {
	objc.Send[objc.ID](n_.ID, objc.Sel("runWithConfiguration:"), configuration)
}

// Provides the framework with an existing AR session to use for Camera Assistance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/setARSession(_:)
func (n_ NISession) SetARSession(session gamekit.Session) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setARSession:"), session)
}

// Returns a world transform to integrate a nearby object in an AR experience.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/worldTransformForObject:
func (n_ NISession) WorldTransformForObject(object ININearbyObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("worldTransformForObject:"), object)
	return rv
}

// iOS-only properties

// The configuration run by the session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/configuration
func (n_ NISession) Configuration() INIConfiguration {
	rv := objc.Send[NIConfiguration](n_.ID, objc.Sel("configuration"))
	return rv
}

// An object that the framework notifies of session events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/delegate
func (n_ NISession) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegate"))
	return rv
}
func (n_ NISession) SetDelegate(value unsafe.Pointer) {
	n_.ID.Send(objc.RegisterName("setDelegate:"), value)
}

// The dispatch queue on which the session invokes delegate callbacks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/delegateQueue
func (n_ NISession) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("delegateQueue"))
	return rv
}
func (n_ NISession) SetDelegateQueue(value unsafe.Pointer) {
	n_.ID.Send(objc.RegisterName("setDelegateQueue:"), value)
}

// A temporary, random identifier for a device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/discoveryToken
func (n_ NISession) DiscoveryToken() INIDiscoveryToken {
	rv := objc.Send[NIDiscoveryToken](n_.ID, objc.Sel("discoveryToken"))
	return rv
}





