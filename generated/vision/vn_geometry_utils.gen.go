// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [GeometryUtils] class.
var (
	GeometryUtilsClass     _GeometryUtilsClass
	GeometryUtilsClassOnce sync.Once
)

func getGeometryUtilsClass() _GeometryUtilsClass {
	GeometryUtilsClassOnce.Do(func() {
		GeometryUtilsClass = _GeometryUtilsClass{objc.GetClass("VNGeometryUtils")}
	})
	return GeometryUtilsClass
}

type _GeometryUtilsClass struct {
	class objc.Class
}

// An interface definition for the [GeometryUtils] class.
type IGeometryUtils interface {
	objectivec.IObject
	// properties:
	// methods:
}

// Utility methods to determine the geometries of various Vision types.


// Utility methods to determine the geometries of various Vision types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNGeometryUtils
type GeometryUtils struct {
	objectivec.Object
}

// GeometryUtilsFrom constructs a [GeometryUtils] from an unsafe.Pointer.
//
// Utility methods to determine the geometries of various Vision types.
func GeometryUtilsFrom(ptr unsafe.Pointer) GeometryUtils {
	return GeometryUtils{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (gc _GeometryUtilsClass) Alloc() GeometryUtils {
	rv := objc.Send[GeometryUtils](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GeometryUtilsClass) New() GeometryUtils {
	rv := objc.Send[GeometryUtils](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GeometryUtils) Init() GeometryUtils {
	rv := objc.Send[GeometryUtils](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GeometryUtils) Autorelease() GeometryUtils {
	rv := objc.Send[GeometryUtils](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGeometryUtils creates a new GeometryUtils instance.
func NewGeometryUtils() GeometryUtils {
	return getGeometryUtilsClass().New()
}




