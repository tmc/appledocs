// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRClusterPath] class.
var (
	MTRClusterPathClass     _MTRClusterPathClass
	MTRClusterPathClassOnce sync.Once
)

func getMTRClusterPathClass() _MTRClusterPathClass {
	MTRClusterPathClassOnce.Do(func() {
		MTRClusterPathClass = _MTRClusterPathClass{objc.GetClass("MTRClusterPath")}
	})
	return MTRClusterPathClass
}

type _MTRClusterPathClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPath] class.
type IMTRClusterPath interface {
	objectivec.IObject
	Cluster() foundation.Number
	SetCluster(value foundation.INumber)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPath
type MTRClusterPath struct {
	objectivec.Object
}

// MTRClusterPathFrom constructs a [MTRClusterPath] from an unsafe.Pointer.
func MTRClusterPathFrom(ptr unsafe.Pointer) MTRClusterPath {
	return MTRClusterPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPathClass) Alloc() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPathClass) New() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPath) Init() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPath) Autorelease() MTRClusterPath {
	rv := objc.Send[MTRClusterPath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPath creates a new MTRClusterPath instance.
func NewMTRClusterPath() MTRClusterPath {
	return getMTRClusterPathClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrclusterpath/cluster
func (m_ MTRClusterPath) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrclusterpath/cluster
func (m_ MTRClusterPath) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrclusterpath/endpoint
func (m_ MTRClusterPath) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrclusterpath/endpoint
func (m_ MTRClusterPath) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}



