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
	// properties:
	AccessGrants() IMTRAccessGrant
	SetAccessGrants(value IMTRAccessGrant)
	Attributes() IMTRServerAttribute
	SetAttributes(value IMTRServerAttribute)
	ClusterID() objc.IObject /* cross-framework: NSNumber */
	SetClusterID(value objc.IObject /* cross-framework: NSNumber */)
	ClusterRevision() objc.IObject /* cross-framework: NSNumber */
	SetClusterRevision(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/accessgrants
func (m_ MTRServerCluster) AccessGrants() IMTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](m_.ID, objc.Sel("accessGrants"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/accessgrants
func (m_ MTRServerCluster) SetAccessGrants(value IMTRAccessGrant) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAccessGrants:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/attributes
func (m_ MTRServerCluster) Attributes() IMTRServerAttribute {
	rv := objc.Send[MTRServerAttribute](m_.ID, objc.Sel("attributes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/attributes
func (m_ MTRServerCluster) SetAttributes(value IMTRServerAttribute) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterid
func (m_ MTRServerCluster) ClusterID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clusterID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterid
func (m_ MTRServerCluster) SetClusterID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusterID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterrevision
func (m_ MTRServerCluster) ClusterRevision() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("clusterRevision"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrservercluster/clusterrevision
func (m_ MTRServerCluster) SetClusterRevision(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setClusterRevision:"), value)
}



