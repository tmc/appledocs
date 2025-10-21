// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHProjectInfo] class.
var (
	PHProjectInfoClass     _PHProjectInfoClass
	PHProjectInfoClassOnce sync.Once
)

func getPHProjectInfoClass() _PHProjectInfoClass {
	PHProjectInfoClassOnce.Do(func() {
		PHProjectInfoClass = _PHProjectInfoClass{objc.GetClass("PHProjectInfo")}
	})
	return PHProjectInfoClass
}

type _PHProjectInfoClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectInfo] class.
type IPHProjectInfo interface {
	objectivec.IObject
}

// Information about the project extension.
//
// macOS Photos automatically generates a object when creating a new project. Photos passes along the project information with a object. This object contains metadata about the project’s creation source, sections, product type, branding, and page numbers. Your extension leverages project information to influence project layout, autoflow, and theme selection. The properties of this class are immutable, and your extension can’t instantiate the object directly.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectInfo
type PHProjectInfo struct {
	objectivec.Object
}

// PHProjectInfoFrom constructs a [PHProjectInfo] from an unsafe.Pointer.
//
// Information about the project extension.
func PHProjectInfoFrom(ptr unsafe.Pointer) PHProjectInfo {
	return PHProjectInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectInfoClass) Alloc() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectInfoClass) New() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectInfo) Init() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectInfo) Autorelease() PHProjectInfo {
	rv := objc.Send[PHProjectInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectInfo creates a new PHProjectInfo instance.
func NewPHProjectInfo() PHProjectInfo {
	return getPHProjectInfoClass().New()
}




