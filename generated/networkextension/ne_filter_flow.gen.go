// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NEFilterFlow] class.
var (
	NEFilterFlowClass     _NEFilterFlowClass
	NEFilterFlowClassOnce sync.Once
)

func getNEFilterFlowClass() _NEFilterFlowClass {
	NEFilterFlowClassOnce.Do(func() {
		NEFilterFlowClass = _NEFilterFlowClass{objc.GetClass("NEFilterFlow")}
	})
	return NEFilterFlowClass
}

type _NEFilterFlowClass struct {
	class objc.Class
}

// An interface definition for the [NEFilterFlow] class.
type INEFilterFlow interface {
	objectivec.IObject
}

// The abstract base class for types that represent flows of network data.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterFlow
type NEFilterFlow struct {
	objectivec.Object
}

// NEFilterFlowFrom constructs a [NEFilterFlow] from an unsafe.Pointer.
//
// The abstract base class for types that represent flows of network data.
func NEFilterFlowFrom(ptr unsafe.Pointer) NEFilterFlow {
	return NEFilterFlow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NEFilterFlowClass) Alloc() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NEFilterFlowClass) New() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterFlow) Init() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterFlow) Autorelease() NEFilterFlow {
	rv := objc.Send[NEFilterFlow](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterFlow creates a new NEFilterFlow instance.
func NewNEFilterFlow() NEFilterFlow {
	return getNEFilterFlowClass().New()
}


// The audit token of the source application of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappaudittoken
func (n_ NEFilterFlow) SourceAppAuditToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}


// SetSourceAppAuditToken sets the value of the sourceAppAuditToken property.
// The audit token of the source application of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappaudittoken
func (n_ NEFilterFlow) SetSourceAppAuditToken(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppAuditToken:"), value)
}

// A string containing the identifier of the source app of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappidentifier
func (n_ NEFilterFlow) SourceAppIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("sourceAppIdentifier"))
	return rv
}


// SetSourceAppIdentifier sets the value of the sourceAppIdentifier property.
// A string containing the identifier of the source app of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappidentifier
func (n_ NEFilterFlow) SetSourceAppIdentifier(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppIdentifier:"), objc.String(value))
}

// The short version string of the app that is the source of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappversion
func (n_ NEFilterFlow) SourceAppVersion() string {
	rv := objc.Send[string](n_.ID, objc.Sel("sourceAppVersion"))
	return rv
}


// SetSourceAppVersion sets the value of the sourceAppVersion property.
// The short version string of the app that is the source of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappversion
func (n_ NEFilterFlow) SetSourceAppVersion(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppVersion:"), objc.String(value))
}

// The audit token of the process that created the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceprocessaudittoken
func (n_ NEFilterFlow) SourceProcessAuditToken() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("sourceProcessAuditToken"))
	return rv
}


// SetSourceProcessAuditToken sets the value of the sourceProcessAuditToken property.
// The audit token of the process that created the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceprocessaudittoken
func (n_ NEFilterFlow) SetSourceProcessAuditToken(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceProcessAuditToken:"), value)
}

// The initial direction of the flow: incoming or outgoing.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/direction
func (n_ NEFilterFlow) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("direction"))
	return rv
}


// SetDirection sets the value of the direction property.
// The initial direction of the flow: incoming or outgoing.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/direction
func (n_ NEFilterFlow) SetDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDirection:"), value)
}

// A byte string that uniquely identifies the binary for each build of the app that is the source of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappuniqueidentifier
func (n_ NEFilterFlow) SourceAppUniqueIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return rv
}


// SetSourceAppUniqueIdentifier sets the value of the sourceAppUniqueIdentifier property.
// A byte string that uniquely identifies the binary for each build of the app that is the source of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappuniqueidentifier
func (n_ NEFilterFlow) SetSourceAppUniqueIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppUniqueIdentifier:"), value)
}

// The flow’s HTTP URL.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/url
func (n_ NEFilterFlow) Url() foundation.URL {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("url"))
	return rv
}


// SetUrl sets the value of the url property.
// The flow’s HTTP URL.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/url
func (n_ NEFilterFlow) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUrl:"), value)
}

// The unique identifier of the flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/identifier
func (n_ NEFilterFlow) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// The unique identifier of the flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/identifier
func (n_ NEFilterFlow) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentifier:"), value)
}

// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) NEFilterFlowBytesMax() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("NEFilterFlowBytesMax"))
	return rv
}


// SetNEFilterFlowBytesMax sets the value of the NEFilterFlowBytesMax property.
// The maximum number of bytes to pass or peek for a flow.

//
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) SetNEFilterFlowBytesMax(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterFlowBytesMax:"), value)
}



