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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/intermediatecacertificateid
func (m_ MTRDistinguishedNameInfo) IntermediateCACertificateID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("intermediateCACertificateID"))
	return rv
}


// SetIntermediateCACertificateID sets the value of the intermediateCACertificateID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/intermediatecacertificateid
func (m_ MTRDistinguishedNameInfo) SetIntermediateCACertificateID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCACertificateID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/nodeid
func (m_ MTRDistinguishedNameInfo) NodeID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nodeID"))
	return rv
}


// SetNodeID sets the value of the nodeID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/nodeid
func (m_ MTRDistinguishedNameInfo) SetNodeID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/fabricid
func (m_ MTRDistinguishedNameInfo) FabricID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricID"))
	return rv
}


// SetFabricID sets the value of the fabricID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/fabricid
func (m_ MTRDistinguishedNameInfo) SetFabricID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/rootcacertificateid
func (m_ MTRDistinguishedNameInfo) RootCACertificateID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rootCACertificateID"))
	return rv
}


// SetRootCACertificateID sets the value of the rootCACertificateID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/rootcacertificateid
func (m_ MTRDistinguishedNameInfo) SetRootCACertificateID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCACertificateID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/caseauthenticatedtags
func (m_ MTRDistinguishedNameInfo) CaseAuthenticatedTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}


// SetCaseAuthenticatedTags sets the value of the caseAuthenticatedTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdistinguishednameinfo/caseauthenticatedtags
func (m_ MTRDistinguishedNameInfo) SetCaseAuthenticatedTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAuthenticatedTags:"), value)
}



