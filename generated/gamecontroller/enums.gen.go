// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

/* debug [enums.gen.go]: Generating 11 enums for GameController */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum GCDeviceBatteryState (4 cases) */
// GCDeviceBatteryState - A state that indicates whether a device’s battery has power and is charging.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State
type GCDeviceBatteryState uint

const (
	// GCDeviceBatteryStateCharging - The device’s battery has power and is charging, but isn’t fully charged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/charging
	GCDeviceBatteryStateCharging GCDeviceBatteryState = 0
	// GCDeviceBatteryStateDischarging - The device’s battery is discharging.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/discharging
	GCDeviceBatteryStateDischarging GCDeviceBatteryState = 0
	// GCDeviceBatteryStateFull - The device’s battery has power and is fully charged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/full
	GCDeviceBatteryStateFull GCDeviceBatteryState = 0
	// GCDeviceBatteryStateUnknown - The state of the device’s battery is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State/unknown
	GCDeviceBatteryStateUnknown GCDeviceBatteryState = 0
)

/* debug [enums.gen.go]: Processing enum GCSystemGestureState (3 cases) */
// GCSystemGestureState - A state for handling input when an element is part of a system gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/SystemGestureState
type GCSystemGestureState uint

const (
	// GCSystemGestureStateAlwaysReceive - A state that sends input to your app and a gesture recognizer simultaneously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/SystemGestureState/alwaysReceive
	GCSystemGestureStateAlwaysReceive GCSystemGestureState = 0
	// GCSystemGestureStateDisabled - A state that sends input to your app directly and not to a gesture recognizer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/SystemGestureState/disabled
	GCSystemGestureStateDisabled GCSystemGestureState = 0
	// GCSystemGestureStateEnabled - A state that sends input to your app only after a gesture recognizer doesn’t identify a gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerElement/SystemGestureState/enabled
	GCSystemGestureStateEnabled GCSystemGestureState = 0
)

/* debug [enums.gen.go]: Processing enum GCControllerPlayerIndex (5 cases) */
// GCControllerPlayerIndex - The possible values for controller player indices.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex
type GCControllerPlayerIndex uint

const (
	// GCControllerPlayerIndex1 - Player one is using the controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex/index1
	GCControllerPlayerIndex1 GCControllerPlayerIndex = 0
	// GCControllerPlayerIndex2 - Player two is using the controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex/index2
	GCControllerPlayerIndex2 GCControllerPlayerIndex = 0
	// GCControllerPlayerIndex3 - Player three is using the controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex/index3
	GCControllerPlayerIndex3 GCControllerPlayerIndex = 0
	// GCControllerPlayerIndex4 - Player four is using the controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex/index4
	GCControllerPlayerIndex4 GCControllerPlayerIndex = 0
	// GCControllerPlayerIndexUnset - The default index for a player on a controller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex/indexUnset
	GCControllerPlayerIndexUnset GCControllerPlayerIndex = 0
)

/* debug [enums.gen.go]: Processing enum GCTouchState (3 cases) */
// GCTouchState - The possible states of the user’s touch.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/TouchState-swift.enum
type GCTouchState uint

const (
	// GCTouchStateDown - The user starts touching the surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/TouchState-swift.enum/down
	GCTouchStateDown GCTouchState = 0
	// GCTouchStateMoving - The user continues touching the surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/TouchState-swift.enum/moving
	GCTouchStateMoving GCTouchState = 0
	// GCTouchStateUp - The user stops or isn’t touching the surface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/TouchState-swift.enum/up
	GCTouchStateUp GCTouchState = 0
)

/* debug [enums.gen.go]: Processing enum GCDevicePhysicalInputElementChange (3 cases) */
// GCDevicePhysicalInputElementChange - Possible values that describe whether the input value of an element changes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange
type GCDevicePhysicalInputElementChange uint

const (
	// GCDevicePhysicalInputElementChanged - There’s a change to the input value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange/changed
	GCDevicePhysicalInputElementChanged GCDevicePhysicalInputElementChange = 0
	// GCDevicePhysicalInputElementNoChange - There’s no change to the input value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange/noChange
	GCDevicePhysicalInputElementNoChange GCDevicePhysicalInputElementChange = 0
	// GCDevicePhysicalInputElementUnknownChange - It’s unknown whether there’s a change to the input value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange/unknownChange
	GCDevicePhysicalInputElementUnknownChange GCDevicePhysicalInputElementChange = 0
)

/* debug [enums.gen.go]: Processing enum GCDualSenseAdaptiveTriggerMode (5 cases) */
// GCDualSenseAdaptiveTriggerMode - The possible modes of an adaptive trigger.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum
type GCDualSenseAdaptiveTriggerMode uint

const (
	// GCDualSenseAdaptiveTriggerModeFeedback - Provides feedback when the user depresses the trigger equal to, or greater than, the start position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum/feedback
	GCDualSenseAdaptiveTriggerModeFeedback GCDualSenseAdaptiveTriggerMode = 0
	// GCDualSenseAdaptiveTriggerModeOff - Provides no adaptive trigger effects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum/off
	GCDualSenseAdaptiveTriggerModeOff GCDualSenseAdaptiveTriggerMode = 0
	// GCDualSenseAdaptiveTriggerModeSlopeFeedback - Provides feedback when the user tilts the trigger between the start and the end positions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum/slopeFeedback
	GCDualSenseAdaptiveTriggerModeSlopeFeedback GCDualSenseAdaptiveTriggerMode = 0
	// GCDualSenseAdaptiveTriggerModeVibration - Vibrates when the user depresses the trigger equal to, or greater than, the start position.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum/vibration
	GCDualSenseAdaptiveTriggerModeVibration GCDualSenseAdaptiveTriggerMode = 0
	// GCDualSenseAdaptiveTriggerModeWeapon - Provides feedback when the user depresses the trigger between the start and the end positions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum/weapon
	GCDualSenseAdaptiveTriggerModeWeapon GCDualSenseAdaptiveTriggerMode = 0
)

/* debug [enums.gen.go]: Processing enum GCDualSenseAdaptiveTriggerStatus (11 cases) */
// GCDualSenseAdaptiveTriggerStatus - The possible states of an adaptive trigger.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum
type GCDualSenseAdaptiveTriggerStatus uint

const (
	// GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied - The trigger is in feedback mode and is applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/feedbackLoadApplied
	GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad - The trigger is in feedback mode, but isn’t applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/feedbackNoLoad
	GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad - The trigger is in slope mode, and is applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/slopeFeedbackApplyingLoad
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished - The trigger is in slope mode, and stopped applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/slopeFeedbackFinished
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady - The trigger is in slope mode, but isn’t applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/slopeFeedbackReady
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusUnknown - The trigger status is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/unknown
	GCDualSenseAdaptiveTriggerStatusUnknown GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating - The trigger is in vibration mode and is vibrating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/vibrationIsVibrating
	GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating - The trigger is in vibration mode, but isn’t vibrating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/vibrationNotVibrating
	GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusWeaponFired - The trigger is in weapon mode, has fired, and has stopped applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/weaponFired
	GCDualSenseAdaptiveTriggerStatusWeaponFired GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusWeaponFiring - The trigger is in weapon mode, firing, and is applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/weaponFiring
	GCDualSenseAdaptiveTriggerStatusWeaponFiring GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusWeaponReady - The trigger is in weapon mode and ready to fire, but isn’t applying the resistive load.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum/weaponReady
	GCDualSenseAdaptiveTriggerStatusWeaponReady GCDualSenseAdaptiveTriggerStatus = 0
)

/* debug [enums.gen.go]: Processing enum GCExtendedGamepadSnapshotDataVersion (2 cases) */
// GCExtendedGamepadSnapshotDataVersion enum type
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotDataVersion
type GCExtendedGamepadSnapshotDataVersion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotDataVersion/version1
	GCExtendedGamepadSnapshotDataVersion1 GCExtendedGamepadSnapshotDataVersion = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotDataVersion/version2
	GCExtendedGamepadSnapshotDataVersion2 GCExtendedGamepadSnapshotDataVersion = 0
)

/* debug [enums.gen.go]: Processing enum GCMicroGamepadSnapshotDataVersion (1 cases) */
// GCMicroGamepadSnapshotDataVersion enum type
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotDataVersion
type GCMicroGamepadSnapshotDataVersion uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotDataVersion/version1
	GCMicroGamepadSnapshotDataVersion1 GCMicroGamepadSnapshotDataVersion = 0
)

/* debug [enums.gen.go]: Processing enum GCPhysicalInputSourceDirection (5 cases) */
// GCPhysicalInputSourceDirection - The directions that a physical input source involves.
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection
type GCPhysicalInputSourceDirection uint

const (
	// GCPhysicalInputSourceDirectionDown - The physical input source supports the down direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection/down
	GCPhysicalInputSourceDirectionDown GCPhysicalInputSourceDirection = 0
	// GCPhysicalInputSourceDirectionNotApplicable - The physical input source doesn’t support directions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection/GCPhysicalInputSourceDirectionNotApplicable
	GCPhysicalInputSourceDirectionNotApplicable GCPhysicalInputSourceDirection = 0
	// GCPhysicalInputSourceDirectionLeft - The physical input source supports the left direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection/left
	GCPhysicalInputSourceDirectionLeft GCPhysicalInputSourceDirection = 0
	// GCPhysicalInputSourceDirectionRight - The physical input source supports the right direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection/right
	GCPhysicalInputSourceDirectionRight GCPhysicalInputSourceDirection = 0
	// GCPhysicalInputSourceDirectionUp - The physical input source contains a value for the up direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection/up
	GCPhysicalInputSourceDirectionUp GCPhysicalInputSourceDirection = 0
)

/* debug [enums.gen.go]: Processing enum GCUIEventTypes (2 cases) */
// GCUIEventTypes enum type
//
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCUIEventTypes
type GCUIEventTypes uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCUIEventTypes/gamepad
	GCUIEventTypeGamepad GCUIEventTypes = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameController/GCUIEventTypes/GCUIEventTypeNone
	GCUIEventTypeNone GCUIEventTypes = 0
)


