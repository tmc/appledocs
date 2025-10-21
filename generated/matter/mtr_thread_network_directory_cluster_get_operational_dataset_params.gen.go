// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] class.
var (
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass     _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass() _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass {
	MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass = _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass{objc.GetClass("MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams")}
	})
	return MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass
}

type _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] class.
type IMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams
type MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsFrom constructs a [MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	return MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass) Alloc() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass) New() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) Init() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) Autorelease() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams creates a new MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams instance.
func NewMTRThreadNetworkDirectoryClusterGetOperationalDatasetParams() MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams {
	return getMTRThreadNetworkDirectoryClusterGetOperationalDatasetParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) ExtendedPanID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("extendedPanID"))
	return rv
}


// SetExtendedPanID sets the value of the extendedPanID property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/extendedPanID
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetExtendedPanID(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/serverSideProcessingTimeout
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams/timedInvokeTimeoutMs
func (m_ MTRThreadNetworkDirectoryClusterGetOperationalDatasetParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



