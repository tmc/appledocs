// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MovementDisorderManager] class.
var (
	MovementDisorderManagerClass     _MovementDisorderManagerClass
	MovementDisorderManagerClassOnce sync.Once
)

func getMovementDisorderManagerClass() _MovementDisorderManagerClass {
	MovementDisorderManagerClassOnce.Do(func() {
		MovementDisorderManagerClass = _MovementDisorderManagerClass{objc.GetClass("CMMovementDisorderManager")}
	})
	return MovementDisorderManagerClass
}

type _MovementDisorderManagerClass struct {
	class objc.Class
}

// An interface definition for the [MovementDisorderManager] class.
type IMovementDisorderManager interface {
	objectivec.IObject
	LastProcessedDate() unsafe.Pointer
	MonitorKinesiasForDuration(duration TimeInterval)
	MonitorKinesiasExpirationDate() unsafe.Pointer
	QueryDyskineticSymptomFromDateToDateWithHandler(fromDate unsafe.Pointer, toDate unsafe.Pointer, handler unsafe.Pointer)
	QueryTremorFromDateToDateWithHandler(fromDate unsafe.Pointer, toDate unsafe.Pointer, handler unsafe.Pointer)
}

// A manager for recording and querying movement disorder data.
//
// Use to measure a resting Parkinsonian tremor in the 3-7 Hz range and choreiform dyskinetic symptoms. When collecting data, the user should wear Apple Watch on their most affected arm. requires an entitlement from Apple. To apply for the entitlement, see .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager
type MovementDisorderManager struct {
	objectivec.Object
}

// MovementDisorderManagerFrom constructs a [MovementDisorderManager] from an unsafe.Pointer.
//
// A manager for recording and querying movement disorder data.
func MovementDisorderManagerFrom(ptr unsafe.Pointer) MovementDisorderManager {
	return MovementDisorderManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MovementDisorderManagerClass) Alloc() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MovementDisorderManagerClass) New() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MovementDisorderManager) Init() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MovementDisorderManager) Autorelease() MovementDisorderManager {
	rv := objc.Send[MovementDisorderManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovementDisorderManager creates a new MovementDisorderManager instance.
func NewMovementDisorderManager() MovementDisorderManager {
	return getMovementDisorderManagerClass().New()
}


// A value indicating whether the user has authorized the app to monitor and query for movement disorder data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/authorizationStatus()
func (mc _MovementDisorderManagerClass) AuthorizationStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("authorizationStatus"))
	return rv
}

// A Boolean value indicating whether the current device supports the movement disorder manager.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/isAvailable()
func (mc _MovementDisorderManagerClass) IsAvailable() bool {
	rv := objc.Send[bool](objc.ID(mc.class), objc.Sel("isAvailable"))
	return rv
}

// Returns a string that describes the movement disorder algorithm’s current version.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/version()
func (mc _MovementDisorderManagerClass) Version() string {
	rv := objc.Send[string](objc.ID(mc.class), objc.Sel("version"))
	return rv
}

// Returns the date of the most recently calculated results.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/lastProcessedDate()
func (m_ MovementDisorderManager) LastProcessedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("lastProcessedDate"))
	return rv
}

// Calculate and store tremor and dyskinetic symptom results for the duration of the specified time interval.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/monitorKinesias(forDuration:)
func (m_ MovementDisorderManager) MonitorKinesiasForDuration(duration TimeInterval) {
	objc.Send[objc.ID](m_.ID, objc.Sel("monitorKinesiasForDuration:"), duration)
}

// Returns the expiration date for the most recent monitoring period.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/monitorKinesiasExpirationDate()
func (m_ MovementDisorderManager) MonitorKinesiasExpirationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("monitorKinesiasExpirationDate"))
	return rv
}

// Query for dyskinetic symptoms from the provided time interval.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/queryDyskineticSymptom(from:to:withHandler:)
func (m_ MovementDisorderManager) QueryDyskineticSymptomFromDateToDateWithHandler(fromDate unsafe.Pointer, toDate unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryDyskineticSymptomFromDate:toDate:withHandler:"), fromDate, toDate, handler)
}

// Query for tremor results from the provided time interval.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMovementDisorderManager/queryTremor(from:to:withHandler:)
func (m_ MovementDisorderManager) QueryTremorFromDateToDateWithHandler(fromDate unsafe.Pointer, toDate unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("queryTremorFromDate:toDate:withHandler:"), fromDate, toDate, handler)
}



