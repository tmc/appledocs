// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREventRequestPath] class.
var (
	MTREventRequestPathClass     _MTREventRequestPathClass
	MTREventRequestPathClassOnce sync.Once
)

func getMTREventRequestPathClass() _MTREventRequestPathClass {
	MTREventRequestPathClassOnce.Do(func() {
		MTREventRequestPathClass = _MTREventRequestPathClass{objc.GetClass("MTREventRequestPath")}
	})
	return MTREventRequestPathClass
}

type _MTREventRequestPathClass struct {
	class objc.Class
}

// An interface definition for the [MTREventRequestPath] class.
type IMTREventRequestPath interface {
	objectivec.IObject
	// properties:
	Cluster() objc.IObject /* cross-framework: NSNumber */
	SetCluster(value objc.IObject /* cross-framework: NSNumber */)
	Endpoint() objc.IObject /* cross-framework: NSNumber */
	SetEndpoint(value objc.IObject /* cross-framework: NSNumber */)
	Event() objc.IObject /* cross-framework: NSNumber */
	SetEvent(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREventRequestPath
type MTREventRequestPath struct {
	objectivec.Object
}

// MTREventRequestPathFrom constructs a [MTREventRequestPath] from an unsafe.Pointer.
func MTREventRequestPathFrom(ptr unsafe.Pointer) MTREventRequestPath {
	return MTREventRequestPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREventRequestPathClass) Alloc() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREventRequestPathClass) New() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREventRequestPath) Init() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREventRequestPath) Autorelease() MTREventRequestPath {
	rv := objc.Send[MTREventRequestPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREventRequestPath creates a new MTREventRequestPath instance.
func NewMTREventRequestPath() MTREventRequestPath {
	return getMTREventRequestPathClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/cluster
func (m_ MTREventRequestPath) Cluster() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cluster"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/cluster
func (m_ MTREventRequestPath) SetCluster(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/endpoint
func (m_ MTREventRequestPath) Endpoint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("endpoint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/endpoint
func (m_ MTREventRequestPath) SetEndpoint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/event
func (m_ MTREventRequestPath) Event() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("event"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtreventrequestpath/event
func (m_ MTREventRequestPath) SetEvent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEvent:"), value)
}



