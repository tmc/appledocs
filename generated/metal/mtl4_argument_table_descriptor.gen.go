// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTL4ArgumentTableDescriptor] class.
var (
	MTL4ArgumentTableDescriptorClass     _MTL4ArgumentTableDescriptorClass
	MTL4ArgumentTableDescriptorClassOnce sync.Once
)

func getMTL4ArgumentTableDescriptorClass() _MTL4ArgumentTableDescriptorClass {
	MTL4ArgumentTableDescriptorClassOnce.Do(func() {
		MTL4ArgumentTableDescriptorClass = _MTL4ArgumentTableDescriptorClass{objc.GetClass("MTL4ArgumentTableDescriptor")}
	})
	return MTL4ArgumentTableDescriptorClass
}

type _MTL4ArgumentTableDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [MTL4ArgumentTableDescriptor] class.
type IMTL4ArgumentTableDescriptor interface {
	objectivec.IObject
	InitializeBindings() bool
	SetInitializeBindings(value bool)
	Label() string
	SetLabel(value string)
	MaxBufferBindCount() int
	SetMaxBufferBindCount(value int)
	MaxSamplerStateBindCount() int
	SetMaxSamplerStateBindCount(value int)
	MaxTextureBindCount() int
	SetMaxTextureBindCount(value int)
	SupportAttributeStrides() bool
	SetSupportAttributeStrides(value bool)
	MTL4CommandQueueErrorDomain() string
}

// Groups parameters for the creation of a Metal argument table.
//
// Argument tables provide resource bindings to your Metal pipeline states.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4ArgumentTableDescriptor
type MTL4ArgumentTableDescriptor struct {
	objectivec.Object
}

// MTL4ArgumentTableDescriptorFrom constructs a [MTL4ArgumentTableDescriptor] from an unsafe.Pointer.
//
// Groups parameters for the creation of a Metal argument table.
func MTL4ArgumentTableDescriptorFrom(ptr unsafe.Pointer) MTL4ArgumentTableDescriptor {
	return MTL4ArgumentTableDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTL4ArgumentTableDescriptorClass) Alloc() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTL4ArgumentTableDescriptorClass) New() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTL4ArgumentTableDescriptor) Init() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTL4ArgumentTableDescriptor) Autorelease() MTL4ArgumentTableDescriptor {
	rv := objc.Send[MTL4ArgumentTableDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4ArgumentTableDescriptor creates a new MTL4ArgumentTableDescriptor instance.
func NewMTL4ArgumentTableDescriptor() MTL4ArgumentTableDescriptor {
	return getMTL4ArgumentTableDescriptorClass().New()
}


// Configures whether Metal initializes the bindings to nil values upon creation of argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/initializebindings
func (m_ MTL4ArgumentTableDescriptor) InitializeBindings() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("initializeBindings"))
	return rv
}


// SetInitializeBindings sets the value of the initializeBindings property.
// Configures whether Metal initializes the bindings to nil values upon creation of argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/initializebindings
func (m_ MTL4ArgumentTableDescriptor) SetInitializeBindings(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setInitializeBindings:"), value)
}

// Assigns an optional label with the argument table for debug purposes.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/label
func (m_ MTL4ArgumentTableDescriptor) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// Assigns an optional label with the argument table for debug purposes.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/label
func (m_ MTL4ArgumentTableDescriptor) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// Determines the number of buffer-binding slots for the argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxbufferbindcount
func (m_ MTL4ArgumentTableDescriptor) MaxBufferBindCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxBufferBindCount"))
	return rv
}


// SetMaxBufferBindCount sets the value of the maxBufferBindCount property.
// Determines the number of buffer-binding slots for the argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxbufferbindcount
func (m_ MTL4ArgumentTableDescriptor) SetMaxBufferBindCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxBufferBindCount:"), value)
}

// Determines the number of sampler state-binding slots for the argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxsamplerstatebindcount
func (m_ MTL4ArgumentTableDescriptor) MaxSamplerStateBindCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxSamplerStateBindCount"))
	return rv
}


// SetMaxSamplerStateBindCount sets the value of the maxSamplerStateBindCount property.
// Determines the number of sampler state-binding slots for the argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxsamplerstatebindcount
func (m_ MTL4ArgumentTableDescriptor) SetMaxSamplerStateBindCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxSamplerStateBindCount:"), value)
}

// Determines the number of texture-binding slots for the argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxtexturebindcount
func (m_ MTL4ArgumentTableDescriptor) MaxTextureBindCount() int {
	rv := objc.Send[int](m_.ID, objc.Sel("maxTextureBindCount"))
	return rv
}


// SetMaxTextureBindCount sets the value of the maxTextureBindCount property.
// Determines the number of texture-binding slots for the argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/maxtexturebindcount
func (m_ MTL4ArgumentTableDescriptor) SetMaxTextureBindCount(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMaxTextureBindCount:"), value)
}

// Controls whether Metal should reserve memory for attribute strides in the argument table.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/supportattributestrides
func (m_ MTL4ArgumentTableDescriptor) SupportAttributeStrides() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("supportAttributeStrides"))
	return rv
}


// SetSupportAttributeStrides sets the value of the supportAttributeStrides property.
// Controls whether Metal should reserve memory for attribute strides in the argument table.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4argumenttabledescriptor/supportattributestrides
func (m_ MTL4ArgumentTableDescriptor) SetSupportAttributeStrides(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSupportAttributeStrides:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtl4commandqueueerrordomain
func (m_ MTL4ArgumentTableDescriptor) MTL4CommandQueueErrorDomain() string {
	rv := objc.Send[string](m_.ID, objc.Sel("MTL4CommandQueueErrorDomain"))
	return rv
}



