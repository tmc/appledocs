// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHLocalIdentifierMapping] class.
var (
	PHLocalIdentifierMappingClass     _PHLocalIdentifierMappingClass
	PHLocalIdentifierMappingClassOnce sync.Once
)

func getPHLocalIdentifierMappingClass() _PHLocalIdentifierMappingClass {
	PHLocalIdentifierMappingClassOnce.Do(func() {
		PHLocalIdentifierMappingClass = _PHLocalIdentifierMappingClass{objc.GetClass("PHLocalIdentifierMapping")}
	})
	return PHLocalIdentifierMappingClass
}

type _PHLocalIdentifierMappingClass struct {
	class objc.Class
}

// An interface definition for the [PHLocalIdentifierMapping] class.
type IPHLocalIdentifierMapping interface {
	objectivec.IObject
}

// An object that contains the local identifier result from looking up a cloud identifier, or an error indicating why the lookup failed.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHLocalIdentifierMapping
type PHLocalIdentifierMapping struct {
	objectivec.Object
}

// PHLocalIdentifierMappingFrom constructs a [PHLocalIdentifierMapping] from an unsafe.Pointer.
//
// An object that contains the local identifier result from looking up a cloud identifier, or an error indicating why the lookup failed.
func PHLocalIdentifierMappingFrom(ptr unsafe.Pointer) PHLocalIdentifierMapping {
	return PHLocalIdentifierMapping{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHLocalIdentifierMappingClass) Alloc() PHLocalIdentifierMapping {
	rv := objc.Send[PHLocalIdentifierMapping](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHLocalIdentifierMappingClass) New() PHLocalIdentifierMapping {
	rv := objc.Send[PHLocalIdentifierMapping](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHLocalIdentifierMapping) Init() PHLocalIdentifierMapping {
	rv := objc.Send[PHLocalIdentifierMapping](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHLocalIdentifierMapping) Autorelease() PHLocalIdentifierMapping {
	rv := objc.Send[PHLocalIdentifierMapping](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHLocalIdentifierMapping creates a new PHLocalIdentifierMapping instance.
func NewPHLocalIdentifierMapping() PHLocalIdentifierMapping {
	return getPHLocalIdentifierMappingClass().New()
}




