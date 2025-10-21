// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAttributeRequestPath] class.
var (
	MTRAttributeRequestPathClass     _MTRAttributeRequestPathClass
	MTRAttributeRequestPathClassOnce sync.Once
)

func getMTRAttributeRequestPathClass() _MTRAttributeRequestPathClass {
	MTRAttributeRequestPathClassOnce.Do(func() {
		MTRAttributeRequestPathClass = _MTRAttributeRequestPathClass{objc.GetClass("MTRAttributeRequestPath")}
	})
	return MTRAttributeRequestPathClass
}

type _MTRAttributeRequestPathClass struct {
	class objc.Class
}

// An interface definition for the [MTRAttributeRequestPath] class.
type IMTRAttributeRequestPath interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributeRequestPath
type MTRAttributeRequestPath struct {
	objectivec.Object
}

// MTRAttributeRequestPathFrom constructs a [MTRAttributeRequestPath] from an unsafe.Pointer.
func MTRAttributeRequestPathFrom(ptr unsafe.Pointer) MTRAttributeRequestPath {
	return MTRAttributeRequestPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAttributeRequestPathClass) Alloc() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAttributeRequestPathClass) New() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributeRequestPath) Init() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributeRequestPath) Autorelease() MTRAttributeRequestPath {
	rv := objc.Send[MTRAttributeRequestPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributeRequestPath creates a new MTRAttributeRequestPath instance.
func NewMTRAttributeRequestPath() MTRAttributeRequestPath {
	return getMTRAttributeRequestPathClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/attribute
func (m_ MTRAttributeRequestPath) Attribute() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("attribute"))
	return rv
}


// SetAttribute sets the value of the attribute property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/attribute
func (m_ MTRAttributeRequestPath) SetAttribute(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttribute:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/cluster
func (m_ MTRAttributeRequestPath) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/cluster
func (m_ MTRAttributeRequestPath) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/endpoint
func (m_ MTRAttributeRequestPath) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrattributerequestpath/endpoint
func (m_ MTRAttributeRequestPath) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



