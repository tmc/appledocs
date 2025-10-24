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
	// properties:
	Direction() unsafe.Pointer
	SetDirection(value unsafe.Pointer)
	Identifier() objc.IObject /* cross-framework: UUID */
	SetIdentifier(value objc.IObject /* cross-framework: UUID */)
	SourceAppAuditToken() objc.IObject /* cross-framework: Data */
	SetSourceAppAuditToken(value objc.IObject /* cross-framework: Data */)
	SourceAppIdentifier() objc.IObject /* cross-framework: NSString */
	SetSourceAppIdentifier(value objc.IObject /* cross-framework: NSString */)
	SourceAppUniqueIdentifier() objc.IObject /* cross-framework: Data */
	SetSourceAppUniqueIdentifier(value objc.IObject /* cross-framework: Data */)
	SourceAppVersion() objc.IObject /* cross-framework: NSString */
	SetSourceAppVersion(value objc.IObject /* cross-framework: NSString */)
	SourceProcessAuditToken() objc.IObject /* cross-framework: Data */
	SetSourceProcessAuditToken(value objc.IObject /* cross-framework: Data */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
	NEFilterFlowBytesMax() uint64
	SetNEFilterFlowBytesMax(value uint64)
	// methods:
}

// The abstract base class for types that represent flows of network data.


// The abstract base class for types that represent flows of network data.
//
// [Full Topic]
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



// The initial direction of the flow: incoming or outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/direction
func (n_ NEFilterFlow) Direction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("direction"))
	return rv
}


// The initial direction of the flow: incoming or outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/direction
func (n_ NEFilterFlow) SetDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDirection:"), value)
}


// The unique identifier of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/identifier
func (n_ NEFilterFlow) Identifier() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](n_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/identifier
func (n_ NEFilterFlow) SetIdentifier(value objc.IObject /* cross-framework: UUID */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setIdentifier:"), value)
}


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappaudittoken
func (n_ NEFilterFlow) SourceAppAuditToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("sourceAppAuditToken"))
	return rv
}


// The audit token of the source application of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappaudittoken
func (n_ NEFilterFlow) SetSourceAppAuditToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppAuditToken:"), value)
}


// A string containing the identifier of the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappidentifier
func (n_ NEFilterFlow) SourceAppIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppIdentifier"))
	return rv
}


// A string containing the identifier of the source app of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappidentifier
func (n_ NEFilterFlow) SetSourceAppIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppIdentifier:"), value)
}


// A byte string that uniquely identifies the binary for each build of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappuniqueidentifier
func (n_ NEFilterFlow) SourceAppUniqueIdentifier() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("sourceAppUniqueIdentifier"))
	return rv
}


// A byte string that uniquely identifies the binary for each build of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappuniqueidentifier
func (n_ NEFilterFlow) SetSourceAppUniqueIdentifier(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppUniqueIdentifier:"), value)
}


// The short version string of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappversion
func (n_ NEFilterFlow) SourceAppVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("sourceAppVersion"))
	return rv
}


// The short version string of the app that is the source of the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceappversion
func (n_ NEFilterFlow) SetSourceAppVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceAppVersion:"), value)
}


// The audit token of the process that created the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceprocessaudittoken
func (n_ NEFilterFlow) SourceProcessAuditToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](n_.ID, objc.Sel("sourceProcessAuditToken"))
	return rv
}


// The audit token of the process that created the flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/sourceprocessaudittoken
func (n_ NEFilterFlow) SetSourceProcessAuditToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setSourceProcessAuditToken:"), value)
}


// The flow’s HTTP URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/url
func (n_ NEFilterFlow) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](n_.ID, objc.Sel("url"))
	return rv
}


// The flow’s HTTP URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflow/url
func (n_ NEFilterFlow) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setUrl:"), value)
}


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) NEFilterFlowBytesMax() uint64 {
	rv := objc.Send[uint64](n_.ID, objc.Sel("NEFilterFlowBytesMax"))
	return rv
}


// The maximum number of bytes to pass or peek for a flow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/nefilterflowbytesmax
func (n_ NEFilterFlow) SetNEFilterFlowBytesMax(value uint64) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setNEFilterFlowBytesMax:"), value)
}



