// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SharedTextureHandle] class.
var (
	SharedTextureHandleClass     _SharedTextureHandleClass
	SharedTextureHandleClassOnce sync.Once
)

func getSharedTextureHandleClass() _SharedTextureHandleClass {
	SharedTextureHandleClassOnce.Do(func() {
		SharedTextureHandleClass = _SharedTextureHandleClass{objc.GetClass("MTLSharedTextureHandle")}
	})
	return SharedTextureHandleClass
}

type _SharedTextureHandleClass struct {
	class objc.Class
}

// An interface definition for the [SharedTextureHandle] class.
type ISharedTextureHandle interface {
	objectivec.IObject
}

// A texture handle that can be shared across process address space boundaries.
//
// objects may be passed between processes using XPC connections and then used to create a reference to the texture in another process. The texture in the other process must be created using the same on which the shared texture was originally created. To identify which device it was created on, you can use the property of the object.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle
type SharedTextureHandle struct {
	objectivec.Object
}

// SharedTextureHandleFrom constructs a [SharedTextureHandle] from an unsafe.Pointer.
//
// A texture handle that can be shared across process address space boundaries.
func SharedTextureHandleFrom(ptr unsafe.Pointer) SharedTextureHandle {
	return SharedTextureHandle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SharedTextureHandleClass) Alloc() SharedTextureHandle {
	rv := objc.Send[SharedTextureHandle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SharedTextureHandleClass) New() SharedTextureHandle {
	rv := objc.Send[SharedTextureHandle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharedTextureHandle) Init() SharedTextureHandle {
	rv := objc.Send[SharedTextureHandle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharedTextureHandle) Autorelease() SharedTextureHandle {
	rv := objc.Send[SharedTextureHandle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharedTextureHandle creates a new SharedTextureHandle instance.
func NewSharedTextureHandle() SharedTextureHandle {
	return getSharedTextureHandleClass().New()
}


// A string that identifies the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedtexturehandle/label
func (s_ SharedTextureHandle) Label() string {
	rv := objc.Send[string](s_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
// A string that identifies the texture.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlsharedtexturehandle/label
func (s_ SharedTextureHandle) SetLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The device object that created the texture.
//
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedTextureHandle/device
func (s_ SharedTextureHandle) Device() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("device"))
	return rv
}



