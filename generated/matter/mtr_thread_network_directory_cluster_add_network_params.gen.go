// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDirectoryClusterAddNetworkParams] class.
var (
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClass     _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterAddNetworkParamsClass() _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass {
	MTRThreadNetworkDirectoryClusterAddNetworkParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterAddNetworkParamsClass = _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterAddNetworkParams")}
	})
	return MTRThreadNetworkDirectoryClusterAddNetworkParamsClass
}

type _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDirectoryClusterAddNetworkParams] class.
type IMTRThreadNetworkDirectoryClusterAddNetworkParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams
type MTRThreadNetworkDirectoryClusterAddNetworkParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterAddNetworkParamsFrom constructs a [MTRThreadNetworkDirectoryClusterAddNetworkParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterAddNetworkParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterAddNetworkParams {
	return MTRThreadNetworkDirectoryClusterAddNetworkParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass) Alloc() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDirectoryClusterAddNetworkParamsClass) New() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) Init() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) Autorelease() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterAddNetworkParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterAddNetworkParams creates a new MTRThreadNetworkDirectoryClusterAddNetworkParams instance.
func NewMTRThreadNetworkDirectoryClusterAddNetworkParams() MTRThreadNetworkDirectoryClusterAddNetworkParams {
	return getMTRThreadNetworkDirectoryClusterAddNetworkParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) OperationalDataset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalDataset"))
	return rv
}


// SetOperationalDataset sets the value of the operationalDataset property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/operationalDataset
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetOperationalDataset(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalDataset:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterAddNetworkParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterAddNetworkParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



