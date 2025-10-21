// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHAssetResourceCreationOptions] class.
var (
	PHAssetResourceCreationOptionsClass     _PHAssetResourceCreationOptionsClass
	PHAssetResourceCreationOptionsClassOnce sync.Once
)

func getPHAssetResourceCreationOptionsClass() _PHAssetResourceCreationOptionsClass {
	PHAssetResourceCreationOptionsClassOnce.Do(func() {
		PHAssetResourceCreationOptionsClass = _PHAssetResourceCreationOptionsClass{objc.GetClass("PHAssetResourceCreationOptions")}
	})
	return PHAssetResourceCreationOptionsClass
}

type _PHAssetResourceCreationOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResourceCreationOptions] class.
type IPHAssetResourceCreationOptions interface {
	objectivec.IObject
}

// A set of options affecting the creation of a new Photos asset from underlying resources.
//
// You use this class when creating an asset for addition to the Photos library with a object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions
type PHAssetResourceCreationOptions struct {
	objectivec.Object
}

// PHAssetResourceCreationOptionsFrom constructs a [PHAssetResourceCreationOptions] from an unsafe.Pointer.
//
// A set of options affecting the creation of a new Photos asset from underlying resources.
func PHAssetResourceCreationOptionsFrom(ptr unsafe.Pointer) PHAssetResourceCreationOptions {
	return PHAssetResourceCreationOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceCreationOptionsClass) Alloc() PHAssetResourceCreationOptions {
	rv := objc.Send[PHAssetResourceCreationOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceCreationOptionsClass) New() PHAssetResourceCreationOptions {
	rv := objc.Send[PHAssetResourceCreationOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResourceCreationOptions) Init() PHAssetResourceCreationOptions {
	rv := objc.Send[PHAssetResourceCreationOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResourceCreationOptions) Autorelease() PHAssetResourceCreationOptions {
	rv := objc.Send[PHAssetResourceCreationOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResourceCreationOptions creates a new PHAssetResourceCreationOptions instance.
func NewPHAssetResourceCreationOptions() PHAssetResourceCreationOptions {
	return getPHAssetResourceCreationOptionsClass().New()
}


// The type of data being provided for this asset resource. If not specified, one will be inferred from the PHAssetResourceType or file URL extension (if provided).
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/contentType
func (p_ PHAssetResourceCreationOptions) ContentType() UTType {
	rv := objc.Send[UTType](p_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
// The type of data being provided for this asset resource. If not specified, one will be inferred from the PHAssetResourceType or file URL extension (if provided).

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/contentType
func (p_ PHAssetResourceCreationOptions) SetContentType(value UTType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentType:"), value)
}

// The filename for the asset resource being created.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/originalFilename
func (p_ PHAssetResourceCreationOptions) OriginalFilename() string {
	rv := objc.Send[string](p_.ID, objc.Sel("originalFilename"))
	return rv
}


// SetOriginalFilename sets the value of the originalFilename property.
// The filename for the asset resource being created.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/originalFilename
func (p_ PHAssetResourceCreationOptions) SetOriginalFilename(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOriginalFilename:"), objc.String(value))
}

// A Boolean value that determines whether Photos moves or duplicates files when creating an asset resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/shouldMoveFile
func (p_ PHAssetResourceCreationOptions) ShouldMoveFile() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldMoveFile"))
	return rv
}


// SetShouldMoveFile sets the value of the shouldMoveFile property.
// A Boolean value that determines whether Photos moves or duplicates files when creating an asset resource.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/shouldMoveFile
func (p_ PHAssetResourceCreationOptions) SetShouldMoveFile(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldMoveFile:"), value)
}

// The uniform type identifier for the resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/uniformTypeIdentifier
func (p_ PHAssetResourceCreationOptions) UniformTypeIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("uniformTypeIdentifier"))
	return rv
}


// SetUniformTypeIdentifier sets the value of the uniformTypeIdentifier property.
// The uniform type identifier for the resource.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResourceCreationOptions/uniformTypeIdentifier
func (p_ PHAssetResourceCreationOptions) SetUniformTypeIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUniformTypeIdentifier:"), objc.String(value))
}



