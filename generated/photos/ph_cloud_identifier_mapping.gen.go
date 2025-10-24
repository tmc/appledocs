// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHCloudIdentifierMapping] class.
var (
	PHCloudIdentifierMappingClass     _PHCloudIdentifierMappingClass
	PHCloudIdentifierMappingClassOnce sync.Once
)

func getPHCloudIdentifierMappingClass() _PHCloudIdentifierMappingClass {
	PHCloudIdentifierMappingClassOnce.Do(func() {
		PHCloudIdentifierMappingClass = _PHCloudIdentifierMappingClass{objc.GetClass("PHCloudIdentifierMapping")}
	})
	return PHCloudIdentifierMappingClass
}

type _PHCloudIdentifierMappingClass struct {
	class objc.Class
}

// An interface definition for the [PHCloudIdentifierMapping] class.
type IPHCloudIdentifierMapping interface {
	objectivec.IObject
	// properties:
	UserInfo() objc.IObject /* cross-framework: NSString */
	SetUserInfo(value objc.IObject /* cross-framework: NSString */)
	PHLocalIdentifierNotFound() objc.IObject  /* cross-framework: NSString */
	PHLocalIdentifiersErrorKey() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An object that contains the cloud identifier result from looking up a local identifier, or an error indicating why the lookup failed.
//
// The error property exhibits two common errors— and . When encountering multiple identifiers, use the error’s property to retrieve a list of matched local identifiers. You can access them using .

// An object that contains the cloud identifier result from looking up a local identifier, or an error indicating why the lookup failed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCloudIdentifierMapping
type PHCloudIdentifierMapping struct {
	objectivec.Object
}

// PHCloudIdentifierMappingFrom constructs a [PHCloudIdentifierMapping] from an unsafe.Pointer.
//
// An object that contains the cloud identifier result from looking up a local identifier, or an error indicating why the lookup failed.
func PHCloudIdentifierMappingFrom(ptr unsafe.Pointer) PHCloudIdentifierMapping {
	return PHCloudIdentifierMapping{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCloudIdentifierMappingClass) Alloc() PHCloudIdentifierMapping {
	rv := objc.Send[PHCloudIdentifierMapping](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCloudIdentifierMappingClass) New() PHCloudIdentifierMapping {
	rv := objc.Send[PHCloudIdentifierMapping](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCloudIdentifierMapping) Init() PHCloudIdentifierMapping {
	rv := objc.Send[PHCloudIdentifierMapping](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCloudIdentifierMapping) Autorelease() PHCloudIdentifierMapping {
	rv := objc.Send[PHCloudIdentifierMapping](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCloudIdentifierMapping creates a new PHCloudIdentifierMapping instance.
func NewPHCloudIdentifierMapping() PHCloudIdentifierMapping {
	return getPHCloudIdentifierMappingClass().New()
}

// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (p_ PHCloudIdentifierMapping) UserInfo() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userInfo"))
	return rv
}

// The user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
func (p_ PHCloudIdentifierMapping) SetUserInfo(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserInfo:"), value)
}

// A constant value that indicates that the system can’t resolve a local object from a global identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlocalidentifiernotfound
func (p_ PHCloudIdentifierMapping) PHLocalIdentifierNotFound() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("PHLocalIdentifierNotFound"))
	return rv
}

// An error key that retrieves an array of string values representing local identifiers matched to a cloud identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phlocalidentifierserrorkey
func (p_ PHCloudIdentifierMapping) PHLocalIdentifiersErrorKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("PHLocalIdentifiersErrorKey"))
	return rv
}
