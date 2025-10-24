// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterSmokeCOAlarm] class.
var (
	MTRClusterSmokeCOAlarmClass     _MTRClusterSmokeCOAlarmClass
	MTRClusterSmokeCOAlarmClassOnce sync.Once
)

func getMTRClusterSmokeCOAlarmClass() _MTRClusterSmokeCOAlarmClass {
	MTRClusterSmokeCOAlarmClassOnce.Do(func() {
		MTRClusterSmokeCOAlarmClass = _MTRClusterSmokeCOAlarmClass{objc.GetClass("MTRClusterSmokeCOAlarm")}
	})
	return MTRClusterSmokeCOAlarmClass
}

type _MTRClusterSmokeCOAlarmClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterSmokeCOAlarm] class.
type IMTRClusterSmokeCOAlarm interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterSmokeCOAlarm
type MTRClusterSmokeCOAlarm struct {
	MTRGenericCluster
}

// MTRClusterSmokeCOAlarmFrom constructs a [MTRClusterSmokeCOAlarm] from an unsafe.Pointer.
func MTRClusterSmokeCOAlarmFrom(ptr unsafe.Pointer) MTRClusterSmokeCOAlarm {
	return MTRClusterSmokeCOAlarm{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterSmokeCOAlarmClass) Alloc() MTRClusterSmokeCOAlarm {
	rv := objc.Send[MTRClusterSmokeCOAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterSmokeCOAlarmClass) New() MTRClusterSmokeCOAlarm {
	rv := objc.Send[MTRClusterSmokeCOAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterSmokeCOAlarm) Init() MTRClusterSmokeCOAlarm {
	rv := objc.Send[MTRClusterSmokeCOAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterSmokeCOAlarm) Autorelease() MTRClusterSmokeCOAlarm {
	rv := objc.Send[MTRClusterSmokeCOAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterSmokeCOAlarm creates a new MTRClusterSmokeCOAlarm instance.
func NewMTRClusterSmokeCOAlarm() MTRClusterSmokeCOAlarm {
	return getMTRClusterSmokeCOAlarmClass().New()
}
