// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessGrant] class.
var (
	MTRAccessGrantClass     _MTRAccessGrantClass
	MTRAccessGrantClassOnce sync.Once
)

func getMTRAccessGrantClass() _MTRAccessGrantClass {
	MTRAccessGrantClassOnce.Do(func() {
		MTRAccessGrantClass = _MTRAccessGrantClass{objc.GetClass("MTRAccessGrant")}
	})
	return MTRAccessGrantClass
}

type _MTRAccessGrantClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessGrant] class.
type IMTRAccessGrant interface {
	objectivec.IObject
	// properties:
	AuthenticationMode() MTRAccessControlEntryAuthMode
	SetAuthenticationMode(value MTRAccessControlEntryAuthMode)
	GrantedPrivilege() MTRAccessControlEntryPrivilege
	SetGrantedPrivilege(value MTRAccessControlEntryPrivilege)
	SubjectID() objc.IObject /* cross-framework: NSNumber */
	SetSubjectID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessGrant
type MTRAccessGrant struct {
	objectivec.Object
}

// MTRAccessGrantFrom constructs a [MTRAccessGrant] from an unsafe.Pointer.
func MTRAccessGrantFrom(ptr unsafe.Pointer) MTRAccessGrant {
	return MTRAccessGrant{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessGrantClass) Alloc() MTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessGrantClass) New() MTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessGrant) Init() MTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessGrant) Autorelease() MTRAccessGrant {
	rv := objc.Send[MTRAccessGrant](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessGrant creates a new MTRAccessGrant instance.
func NewMTRAccessGrant() MTRAccessGrant {
	return getMTRAccessGrantClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/authenticationmode
func (m_ MTRAccessGrant) AuthenticationMode() MTRAccessControlEntryAuthMode {
	rv := objc.Send[MTRAccessControlEntryAuthMode](m_.ID, objc.Sel("authenticationMode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/authenticationmode
func (m_ MTRAccessGrant) SetAuthenticationMode(value MTRAccessControlEntryAuthMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAuthenticationMode:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/grantedprivilege
func (m_ MTRAccessGrant) GrantedPrivilege() MTRAccessControlEntryPrivilege {
	rv := objc.Send[MTRAccessControlEntryPrivilege](m_.ID, objc.Sel("grantedPrivilege"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/grantedprivilege
func (m_ MTRAccessGrant) SetGrantedPrivilege(value MTRAccessControlEntryPrivilege) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGrantedPrivilege:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/subjectid
func (m_ MTRAccessGrant) SubjectID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("subjectID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccessgrant/subjectid
func (m_ MTRAccessGrant) SetSubjectID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubjectID:"), value)
}
