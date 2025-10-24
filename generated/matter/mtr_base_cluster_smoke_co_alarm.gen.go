// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterSmokeCOAlarm] class.
var (
	MTRBaseClusterSmokeCOAlarmClass     _MTRBaseClusterSmokeCOAlarmClass
	MTRBaseClusterSmokeCOAlarmClassOnce sync.Once
)

func getMTRBaseClusterSmokeCOAlarmClass() _MTRBaseClusterSmokeCOAlarmClass {
	MTRBaseClusterSmokeCOAlarmClassOnce.Do(func() {
		MTRBaseClusterSmokeCOAlarmClass = _MTRBaseClusterSmokeCOAlarmClass{objc.GetClass("MTRBaseClusterSmokeCOAlarm")}
	})
	return MTRBaseClusterSmokeCOAlarmClass
}

type _MTRBaseClusterSmokeCOAlarmClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterSmokeCOAlarm] class.
type IMTRBaseClusterSmokeCOAlarm interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterSmokeCOAlarm
type MTRBaseClusterSmokeCOAlarm struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterSmokeCOAlarmFrom constructs a [MTRBaseClusterSmokeCOAlarm] from an unsafe.Pointer.
func MTRBaseClusterSmokeCOAlarmFrom(ptr unsafe.Pointer) MTRBaseClusterSmokeCOAlarm {
	return MTRBaseClusterSmokeCOAlarm{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterSmokeCOAlarmClass) Alloc() MTRBaseClusterSmokeCOAlarm {
	rv := objc.Send[MTRBaseClusterSmokeCOAlarm](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterSmokeCOAlarmClass) New() MTRBaseClusterSmokeCOAlarm {
	rv := objc.Send[MTRBaseClusterSmokeCOAlarm](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterSmokeCOAlarm) Init() MTRBaseClusterSmokeCOAlarm {
	rv := objc.Send[MTRBaseClusterSmokeCOAlarm](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterSmokeCOAlarm) Autorelease() MTRBaseClusterSmokeCOAlarm {
	rv := objc.Send[MTRBaseClusterSmokeCOAlarm](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterSmokeCOAlarm creates a new MTRBaseClusterSmokeCOAlarm instance.
func NewMTRBaseClusterSmokeCOAlarm() MTRBaseClusterSmokeCOAlarm {
	return getMTRBaseClusterSmokeCOAlarmClass().New()
}
