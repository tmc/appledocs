// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServerCluster] class.
var (
	MTRServerClusterClass     _MTRServerClusterClass
	MTRServerClusterClassOnce sync.Once
)

func getMTRServerClusterClass() _MTRServerClusterClass {
	MTRServerClusterClassOnce.Do(func() {
		MTRServerClusterClass = _MTRServerClusterClass{objc.GetClass("MTRServerCluster")}
	})
	return MTRServerClusterClass
}

type _MTRServerClusterClass struct {
	class objc.Class
}

// An interface definition for the [MTRServerCluster] class.
type IMTRServerCluster interface {
	objectivec.IObject
	AccessGrants() MTRAccessGrant
	SetAccessGrants(value IMTRAccessGrant)
	Attributes() MTRServerAttribute
	SetAttributes(value MTRServerAttribute)
	ClusterID() foundation.Number
	SetClusterID(value foundation.INumber)
	ClusterRevision() foundation.Number
	SetClusterRevision(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServerCluster
type MTRServerCluster struct {
	objectivec.Object
}

// MTRServerClusterFrom constructs a [MTRServerCluster] from an unsafe.Pointer.
func MTRServerClusterFrom(ptr unsafe.Pointer) MTRServerCluster {
	return MTRServerCluster{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServerClusterClass) Alloc() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServerClusterClass) New() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServerCluster) Init() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServerCluster) Autorelease() MTRServerCluster {
	rv := objc.Send[MTRServerCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServerCluster creates a new MTRServerCluster instance.
func NewMTRServerCluster() MTRServerCluster {
	return getMTRServerClusterClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/accessgrants
func (m_ MTRServerCluster) AccessGrants() MTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](m_.ID, objc.Sel("accessGrants"))
	return rv
}


// SetAccessGrants sets the value of the accessGrants property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/accessgrants
func (m_ MTRServerCluster) SetAccessGrants(value IMTRAccessGrant) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessGrants:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/attributes
func (m_ MTRServerCluster) Attributes() MTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/attributes
func (m_ MTRServerCluster) SetAttributes(value MTRServerAttribute) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterid
func (m_ MTRServerCluster) ClusterID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("clusterID"))
	return rv
}


// SetClusterID sets the value of the clusterID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterid
func (m_ MTRServerCluster) SetClusterID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusterID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterrevision
func (m_ MTRServerCluster) ClusterRevision() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("clusterRevision"))
	return rv
}


// SetClusterRevision sets the value of the clusterRevision property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterrevision
func (m_ MTRServerCluster) SetClusterRevision(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusterRevision:"), value)
}



