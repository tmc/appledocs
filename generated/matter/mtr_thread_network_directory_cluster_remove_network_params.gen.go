// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] class.
var (
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass     _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass() _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass {
	MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass = _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterRemoveNetworkParams")}
	})
	return MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass
}

type _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] class.
type IMTRThreadNetworkDirectoryClusterRemoveNetworkParams interface {
	objectivec.IObject
	// properties:
	ExtendedPanID() objc.IObject /* cross-framework: NSData */
	SetExtendedPanID(value objc.IObject /* cross-framework: NSData */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams
type MTRThreadNetworkDirectoryClusterRemoveNetworkParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterRemoveNetworkParamsFrom constructs a [MTRThreadNetworkDirectoryClusterRemoveNetworkParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterRemoveNetworkParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	return MTRThreadNetworkDirectoryClusterRemoveNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass) Alloc() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass) New() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) Init() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) Autorelease() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterRemoveNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterRemoveNetworkParams creates a new MTRThreadNetworkDirectoryClusterRemoveNetworkParams instance.
func NewMTRThreadNetworkDirectoryClusterRemoveNetworkParams() MTRThreadNetworkDirectoryClusterRemoveNetworkParams {
	return getMTRThreadNetworkDirectoryClusterRemoveNetworkParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) ExtendedPanID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("extendedPanID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetExtendedPanID(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterRemoveNetworkParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterRemoveNetworkParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



