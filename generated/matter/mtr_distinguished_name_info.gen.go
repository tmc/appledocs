// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDistinguishedNameInfo] class.
var (
	MTRDistinguishedNameInfoClass     _MTRDistinguishedNameInfoClass
	MTRDistinguishedNameInfoClassOnce sync.Once
)

func getMTRDistinguishedNameInfoClass() _MTRDistinguishedNameInfoClass {
	MTRDistinguishedNameInfoClassOnce.Do(func() {
		MTRDistinguishedNameInfoClass = _MTRDistinguishedNameInfoClass{objc.GetClass("MTRDistinguishedNameInfo")}
	})
	return MTRDistinguishedNameInfoClass
}

type _MTRDistinguishedNameInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRDistinguishedNameInfo] class.
type IMTRDistinguishedNameInfo interface {
	objectivec.IObject
	// properties:
	CaseAuthenticatedTags() objc.IObject /* cross-framework: NSNumber */
	SetCaseAuthenticatedTags(value objc.IObject /* cross-framework: NSNumber */)
	FabricID() objc.IObject /* cross-framework: NSNumber */
	SetFabricID(value objc.IObject /* cross-framework: NSNumber */)
	IntermediateCACertificateID() objc.IObject /* cross-framework: NSNumber */
	SetIntermediateCACertificateID(value objc.IObject /* cross-framework: NSNumber */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	RootCACertificateID() objc.IObject /* cross-framework: NSNumber */
	SetRootCACertificateID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDistinguishedNameInfo
type MTRDistinguishedNameInfo struct {
	objectivec.Object
}

// MTRDistinguishedNameInfoFrom constructs a [MTRDistinguishedNameInfo] from an unsafe.Pointer.
func MTRDistinguishedNameInfoFrom(ptr unsafe.Pointer) MTRDistinguishedNameInfo {
	return MTRDistinguishedNameInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDistinguishedNameInfoClass) Alloc() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDistinguishedNameInfoClass) New() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDistinguishedNameInfo) Init() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDistinguishedNameInfo) Autorelease() MTRDistinguishedNameInfo {
	rv := objc.Send[MTRDistinguishedNameInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDistinguishedNameInfo creates a new MTRDistinguishedNameInfo instance.
func NewMTRDistinguishedNameInfo() MTRDistinguishedNameInfo {
	return getMTRDistinguishedNameInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/caseauthenticatedtags
func (m_ MTRDistinguishedNameInfo) CaseAuthenticatedTags() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/caseauthenticatedtags
func (m_ MTRDistinguishedNameInfo) SetCaseAuthenticatedTags(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAuthenticatedTags:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/fabricid
func (m_ MTRDistinguishedNameInfo) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/fabricid
func (m_ MTRDistinguishedNameInfo) SetFabricID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/intermediatecacertificateid
func (m_ MTRDistinguishedNameInfo) IntermediateCACertificateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("intermediateCACertificateID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/intermediatecacertificateid
func (m_ MTRDistinguishedNameInfo) SetIntermediateCACertificateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCACertificateID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/nodeid
func (m_ MTRDistinguishedNameInfo) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/nodeid
func (m_ MTRDistinguishedNameInfo) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/rootcacertificateid
func (m_ MTRDistinguishedNameInfo) RootCACertificateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rootCACertificateID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/rootcacertificateid
func (m_ MTRDistinguishedNameInfo) SetRootCACertificateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCACertificateID:"), value)
}



