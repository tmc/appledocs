// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MCPeerID] class.
var (
	MCPeerIDClass     _MCPeerIDClass
	MCPeerIDClassOnce sync.Once
)

func getMCPeerIDClass() _MCPeerIDClass {
	MCPeerIDClassOnce.Do(func() {
		MCPeerIDClass = _MCPeerIDClass{objc.GetClass("MCPeerID")}
	})
	return MCPeerIDClass
}

type _MCPeerIDClass struct {
	class objc.Class
}

// An interface definition for the [MCPeerID] class.
type IMCPeerID interface {
	objectivec.IObject
}

// An object represents a peer in a multipeer session.
//
// You create a single peer ID object that represents the instance of your app running on the local device. The Multipeer Connectivity framework is responsible for creating peer ID objects that represent other devices. To create a new peer ID for the local app and associate a display name with that ID, call . The peer’s name must be no longer than 63 bytes in UTF-8 encoding. Each peer ID your app creates with is unique, even when supplying the same display name. If you want a device’s peer ID to be stable over time, don’t create a new peer ID every time your app begins advertising or browsing. Instead, archive the ID when you create it, and then unarchive it the next time you need it. If you need the peer ID to be tied to the display name, you can archive the name as well, and only create a new peer ID when the name changes, as illustrated in the following code fragment:
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID
type MCPeerID struct {
	objectivec.Object
}

// MCPeerIDFrom constructs a [MCPeerID] from an unsafe.Pointer.
//
// An object represents a peer in a multipeer session.
func MCPeerIDFrom(ptr unsafe.Pointer) MCPeerID {
	return MCPeerID{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MCPeerIDClass) Alloc() MCPeerID {
	rv := objc.Send[MCPeerID](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MCPeerIDClass) New() MCPeerID {
	rv := objc.Send[MCPeerID](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MCPeerID) Init() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MCPeerID) Autorelease() MCPeerID {
	rv := objc.Send[MCPeerID](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMCPeerID creates a new MCPeerID instance.
func NewMCPeerID() MCPeerID {
	return getMCPeerIDClass().New()
}


// Initializes a peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID/init(displayName:)
func NewMCPeerIDWithDisplayName(myDisplayName string) MCPeerID {
	instance := getMCPeerIDClass().Alloc()
	rv := objc.Send[MCPeerID](instance.ID, objc.Sel("initWithDisplayName:"), objc.String(myDisplayName))
	rv.Autorelease()
	return rv
}


// The display name for this peer.
//
// [Full Topic]: https://developer.apple.com/documentation/MultipeerConnectivity/MCPeerID/displayName
func (m_ MCPeerID) DisplayName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("displayName"))
	return rv
}


