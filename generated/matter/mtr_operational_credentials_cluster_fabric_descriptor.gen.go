// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalCredentialsClusterFabricDescriptor] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorClass     _MTROperationalCredentialsClusterFabricDescriptorClass
	MTROperationalCredentialsClusterFabricDescriptorClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorClass() _MTROperationalCredentialsClusterFabricDescriptorClass {
	MTROperationalCredentialsClusterFabricDescriptorClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorClass = _MTROperationalCredentialsClusterFabricDescriptorClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptor")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorClass
}

type _MTROperationalCredentialsClusterFabricDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptor] class.
type IMTROperationalCredentialsClusterFabricDescriptor interface {
	IMTROperationalCredentialsClusterFabricDescriptorStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor
type MTROperationalCredentialsClusterFabricDescriptor struct {
	MTROperationalCredentialsClusterFabricDescriptorStruct
}

// MTROperationalCredentialsClusterFabricDescriptorFrom constructs a [MTROperationalCredentialsClusterFabricDescriptor] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptor {
	return MTROperationalCredentialsClusterFabricDescriptor{
		MTROperationalCredentialsClusterFabricDescriptorStruct: MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) Alloc() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) New() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Init() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Autorelease() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptor creates a new MTROperationalCredentialsClusterFabricDescriptor instance.
func NewMTROperationalCredentialsClusterFabricDescriptor() MTROperationalCredentialsClusterFabricDescriptor {
	return getMTROperationalCredentialsClusterFabricDescriptorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptor) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/fabricindex
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/label
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Label() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/label
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetLabel(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptor) RootPublicKey() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}


// SetRootPublicKey sets the value of the rootPublicKey property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusterfabricdescriptor/rootpublickey
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetRootPublicKey(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}



