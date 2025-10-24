// Code generated from Apple documentation for HIDDriverKit. DO NOT EDIT.

package hiddriverkit
import (
	"unsafe"
)


// C struct types
// IOHIDDevice - An object containing the low-level behavior for all HID device providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice
type IOHIDDevice struct {
}/* debug [types.gen.go/struct]: IOHIDDevice */

// CompleteReport - Completes all async requests made when getting or setting a report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/CompleteReport
type CompleteReport struct {
}/* debug [types.gen.go/struct]: CompleteReport */

// getReport - Gets a report from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/getReport
type getReport struct {
}/* debug [types.gen.go/struct]: getReport */

// handleReport - Handles an asynchronous report received from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/handleReport
type handleReport struct {
}/* debug [types.gen.go/struct]: handleReport */

// setProperty - Updates the specified property on the corresponding kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/setProperty
type setProperty struct {
}/* debug [types.gen.go/struct]: setProperty */

// setReport - Sends a report to the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/setReport
type setReport struct {
}/* debug [types.gen.go/struct]: setReport */

// IOHIDDigitizerCollection - A collection of elements that contain digitizer-related data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection
type IOHIDDigitizerCollection struct {
}/* debug [types.gen.go/struct]: IOHIDDigitizerCollection */

// addElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/addElement
type addElement struct {
}/* debug [types.gen.go/struct]: addElement */

// free
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/free
type free struct {
}/* debug [types.gen.go/struct]: free */

// getElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getElements
type getElements struct {
}/* debug [types.gen.go/struct]: getElements */

// getInRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getInRange
type getInRange struct {
}/* debug [types.gen.go/struct]: getInRange */

// getParentCollection
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getParentCollection
type getParentCollection struct {
}/* debug [types.gen.go/struct]: getParentCollection */

// getTouch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getTouch
type getTouch struct {
}/* debug [types.gen.go/struct]: getTouch */

// getType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getType
type getType struct {
}/* debug [types.gen.go/struct]: getType */

// getX
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getX
type getX struct {
}/* debug [types.gen.go/struct]: getX */

// getY
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getY
type getY struct {
}/* debug [types.gen.go/struct]: getY */

// getZ
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getZ
type getZ struct {
}/* debug [types.gen.go/struct]: getZ */

// initWithType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/initWithType
type initWithType struct {
}/* debug [types.gen.go/struct]: initWithType */

// setInRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setInRange
type setInRange struct {
}/* debug [types.gen.go/struct]: setInRange */

// setTouch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setTouch
type setTouch struct {
}/* debug [types.gen.go/struct]: setTouch */

// setX
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setX
type setX struct {
}/* debug [types.gen.go/struct]: setX */

// setY
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setY
type setY struct {
}/* debug [types.gen.go/struct]: setY */

// setZ
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setZ
type setZ struct {
}/* debug [types.gen.go/struct]: setZ */

// withType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/withType
type withType struct {
}/* debug [types.gen.go/struct]: withType */

// IOHIDDigitizerStylusData - A structure containing digitizer stylus data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerStylusData
type IOHIDDigitizerStylusData struct {
	BarrelPressure Fixed // The barrel pressure value in the range   to  .
	BarrelSwitch uint32 // A single-bit Boolean that indicates whether the barrel switch button is pressed.
	Effect uint32 // An optional stylus effect defined by vendor.
	Eraser uint32 // A Boolean value that indicates whether the inverted stylus is in contact with the surface of the digitizer.
	Identifier uint32 // A unique stylus identifier.
	InRange uint32 // A single-bit Boolean that indicates whether the stylus is in range.
	Invert uint32 // A single-bit Boolean that indicates whether the stylus is inverted.
	PointerType uint32 // An optional pointer type defined by vendor.
	PositionChanged uint32 // A single-bit Boolean that indicates whether the x or y position changed since the last event was dispatched.
	RangeChanged uint32 // A single-bit Boolean that indicates whether the in-range status changed since the last event was dispatched.
	TiltX Fixed // The tilt of the stylus across the x-axis.
	TiltY Fixed // The tilt of the stylus across the y-axis.
	Tip uint32 // A single-bit Boolean that indicates whether the tip of the stylus is in contact with the surface of the digitizer.
	TipChanged uint32 // A single-bit Boolean that indicates whether the tip contact status changed since the last event was dispatched.
	TipPressure Fixed // A tip pressure value in the range   to  .
	Twist Fixed // The clockwise rotation of the stylus.
	UniqueID uint64 // An optional unique identifier for the stylus.
	X Fixed // An x-axis value in the range   to  .
	Y Fixed // A y-axis value in the range   to  .
}/* debug [types.gen.go/struct]: IOHIDDigitizerStylusData */

// IOHIDDigitizerTouchData - A structure containing the current digitizer touch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerTouchData
type IOHIDDigitizerTouchData struct {
	Identifier uint32 // A unique contact identifier.
	InRange uint32 // A single-bit Boolean that indicates whether the finger is in range.
	PositionChanged uint32 // A single-bit Boolean that indicates whether the x or y position changed since the last event was dispatched.
	RangeChanged uint32 // A single-bit Boolean that indicates whether the range variable changed since the last event was dispatched.
	Touch uint32 // A single-bit Boolean that indicates whether the finger is in contact with the surface of the digitizer.
	TouchChanged uint32 // A single-bit Boolean that indicates whether the touch variable changed since the last event was dispatched.
	TouchValid uint32 // A single-bit Boolean that indicates whether the touch contact was intended.
	X Fixed // An x-coordinate value in the range   to  .
	Y Fixed // A y-coordinate value in the range   to  .
}/* debug [types.gen.go/struct]: IOHIDDigitizerTouchData */

// IOHIDElement - An object that contains parsed information from a HID input report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement
type IOHIDElement struct {
}/* debug [types.gen.go/struct]: IOHIDElement */

// commit - Commits the element value to and from the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/commit
type commit struct {
}/* debug [types.gen.go/struct]: commit */

// conformsTo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/conformsTo
type conformsTo struct {
}/* debug [types.gen.go/struct]: conformsTo */

// getChildElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getChildElements
type getChildElements struct {
}/* debug [types.gen.go/struct]: getChildElements */

// getCollectionType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getCollectionType
type getCollectionType struct {
}/* debug [types.gen.go/struct]: getCollectionType */

// getCookie
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getCookie
type getCookie struct {
}/* debug [types.gen.go/struct]: getCookie */

// getDataValue - Gets the data value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getDataValue
type getDataValue struct {
}/* debug [types.gen.go/struct]: getDataValue */

// getFlags
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getFlags
type getFlags struct {
}/* debug [types.gen.go/struct]: getFlags */

// getLogicalMax
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getLogicalMax
type getLogicalMax struct {
}/* debug [types.gen.go/struct]: getLogicalMax */

// getLogicalMin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getLogicalMin
type getLogicalMin struct {
}/* debug [types.gen.go/struct]: getLogicalMin */

// getParentElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getParentElement
type getParentElement struct {
}/* debug [types.gen.go/struct]: getParentElement */

// getPhysicalMax
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getPhysicalMax
type getPhysicalMax struct {
}/* debug [types.gen.go/struct]: getPhysicalMax */

// getPhysicalMin
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getPhysicalMin
type getPhysicalMin struct {
}/* debug [types.gen.go/struct]: getPhysicalMin */

// getReportCount
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getReportCount
type getReportCount struct {
}/* debug [types.gen.go/struct]: getReportCount */

// getReportID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getReportID
type getReportID struct {
}/* debug [types.gen.go/struct]: getReportID */

// getReportSize
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getReportSize
type getReportSize struct {
}/* debug [types.gen.go/struct]: getReportSize */

// getScaledFixedValue - Returns a fixed number that represents the scaled version of the element’s logical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getScaledFixedValue
type getScaledFixedValue struct {
}/* debug [types.gen.go/struct]: getScaledFixedValue */

// getScaledValue - Returns a scaled version of the logical value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getScaledValue
type getScaledValue struct {
}/* debug [types.gen.go/struct]: getScaledValue */

// getTimeStamp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getTimeStamp
type getTimeStamp struct {
}/* debug [types.gen.go/struct]: getTimeStamp */

// getUnit - Returns the units that you use to interpret the element’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getUnit
type getUnit struct {
}/* debug [types.gen.go/struct]: getUnit */

// getUnitExponent - Returns the exponent that you use to interpret the element’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getUnitExponent
type getUnitExponent struct {
}/* debug [types.gen.go/struct]: getUnitExponent */

// getUsage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getUsage
type getUsage struct {
}/* debug [types.gen.go/struct]: getUsage */

// getUsagePage
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getUsagePage
type getUsagePage struct {
}/* debug [types.gen.go/struct]: getUsagePage */

// getValue - Gets the logical value that the device reported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getValue
type getValue struct {
}/* debug [types.gen.go/struct]: getValue */

// setDataValue - Sets the data value of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/setDataValue
type setDataValue struct {
}/* debug [types.gen.go/struct]: setDataValue */

// setValue - Sets the value of the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/setValue
type setValue struct {
}/* debug [types.gen.go/struct]: setValue */

// IOHIDEventService - The base class for implementing a device or operating system service that dispatches events to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService
type IOHIDEventService struct {
}/* debug [types.gen.go/struct]: IOHIDEventService */

// dispatchAbsolutePointerEvent - Dispatches an absolute pointer event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchAbsolutePointerEvent
type dispatchAbsolutePointerEvent struct {
}/* debug [types.gen.go/struct]: dispatchAbsolutePointerEvent */

// dispatchDigitizerStylusEvent - Dispatches a digitizer stylus event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchDigitizerStylusEvent
type dispatchDigitizerStylusEvent struct {
}/* debug [types.gen.go/struct]: dispatchDigitizerStylusEvent */

// dispatchDigitizerTouchEvent - Dispatches a digitizer touch event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchDigitizerTouchEvent
type dispatchDigitizerTouchEvent struct {
}/* debug [types.gen.go/struct]: dispatchDigitizerTouchEvent */

// dispatchEvent - Dispatches a HID event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchEvent
type dispatchEvent struct {
}/* debug [types.gen.go/struct]: dispatchEvent */

// dispatchKeyboardEvent - Dispatches a keyboard-related event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchKeyboardEvent
type dispatchKeyboardEvent struct {
}/* debug [types.gen.go/struct]: dispatchKeyboardEvent */

// dispatchRelativePointerEvent - Dispatches a relative pointer event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchRelativePointerEvent
type dispatchRelativePointerEvent struct {
}/* debug [types.gen.go/struct]: dispatchRelativePointerEvent */

// dispatchRelativeScrollWheelEvent - Dispatches a relative scroll wheel event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchRelativeScrollWheelEvent
type dispatchRelativeScrollWheelEvent struct {
}/* debug [types.gen.go/struct]: dispatchRelativeScrollWheelEvent */

// handleCopyMatchingEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/handleCopyMatchingEvent
type handleCopyMatchingEvent struct {
}/* debug [types.gen.go/struct]: handleCopyMatchingEvent */

// init
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/init
type init struct {
}/* debug [types.gen.go/struct]: init */

// SetLED - Configures the on/off state for an LED on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/SetLED
type SetLED struct {
}/* debug [types.gen.go/struct]: SetLED */

// SetLEDState
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/SetLEDState
type SetLEDState struct {
}/* debug [types.gen.go/struct]: SetLEDState */

// SetProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/SetProperties
type SetProperties struct {
}/* debug [types.gen.go/struct]: SetProperties */

// Start
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/Start
type Start struct {
}/* debug [types.gen.go/struct]: Start */

// Stop
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/Stop
type Stop struct {
}/* debug [types.gen.go/struct]: Stop */

// IOHIDInterface - A provider object for a HID device’s interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface
type IOHIDInterface struct {
}/* debug [types.gen.go/struct]: IOHIDInterface */

// AddReportToPool - Adds a memory descriptor to the report pool.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/AddReportToPool
type AddReportToPool struct {
}/* debug [types.gen.go/struct]: AddReportToPool */

// Close - Closes the interface and stops the delivery of input reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/Close
type Close struct {
}/* debug [types.gen.go/struct]: Close */

// commitElements - Gets or sets the contents of the interface’s stored elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/commitElements
type commitElements struct {
}/* debug [types.gen.go/struct]: commitElements */

// GetReport - Retrieves a new input report from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/GetReport
type GetReport struct {
}/* debug [types.gen.go/struct]: GetReport */

// Open - Opens a session to the device and begins the delivery of input reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/Open
type Open struct {
}/* debug [types.gen.go/struct]: Open */

// processReport - Parses the contents of the specified report and updates the interface’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/processReport
type processReport struct {
}/* debug [types.gen.go/struct]: processReport */

// ReportAvailable - Notifies the interface that an updated report is available from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/ReportAvailable
type ReportAvailable struct {
}/* debug [types.gen.go/struct]: ReportAvailable */

// SetReport - Sends a report to the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/SetReport
type SetReport struct {
}/* debug [types.gen.go/struct]: SetReport */

// IOUserHIDDevice - A provider object for devices that support interactions with users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice
type IOUserHIDDevice struct {
}/* debug [types.gen.go/struct]: IOUserHIDDevice */

// handleStart - Performs any custom initialization associated with starting the device service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice/handleStart
type handleStart struct {
}/* debug [types.gen.go/struct]: handleStart */

// newDeviceDescription - Creates and returns a new dictionary that describes the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice/newDeviceDescription
type newDeviceDescription struct {
}/* debug [types.gen.go/struct]: newDeviceDescription */

// newReportDescriptor - Returns the data in the HID device’s report descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice/newReportDescriptor
type newReportDescriptor struct {
}/* debug [types.gen.go/struct]: newReportDescriptor */

// IOUserHIDEventDriver - A complete driver object that dispatches keyboard, digitizer, scrolling, and pointer events originating from a HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver
type IOUserHIDEventDriver struct {
}/* debug [types.gen.go/struct]: IOUserHIDEventDriver */

// calibrateCenteredPreferredStateElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/calibrateCenteredPreferredStateElement
type calibrateCenteredPreferredStateElement struct {
}/* debug [types.gen.go/struct]: calibrateCenteredPreferredStateElement */

// calibrateJustifiedPreferredStateElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/calibrateJustifiedPreferredStateElement
type calibrateJustifiedPreferredStateElement struct {
}/* debug [types.gen.go/struct]: calibrateJustifiedPreferredStateElement */

// checkGameControllerElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/checkGameControllerElement
type checkGameControllerElement struct {
}/* debug [types.gen.go/struct]: checkGameControllerElement */

// copyKeyboardEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/copyKeyboardEvent
type copyKeyboardEvent struct {
}/* debug [types.gen.go/struct]: copyKeyboardEvent */

// createEventForDigitizerCollection - Creates a HID event object that represents a digitizer collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/createEventForDigitizerCollection
type createEventForDigitizerCollection struct {
}/* debug [types.gen.go/struct]: createEventForDigitizerCollection */

// getButtonStateFromElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/getButtonStateFromElements
type getButtonStateFromElements struct {
}/* debug [types.gen.go/struct]: getButtonStateFromElements */

// handleAbsolutePointerReport - Iterates through absolute pointer elements and dispatches them if the element value has been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleAbsolutePointerReport
type handleAbsolutePointerReport struct {
}/* debug [types.gen.go/struct]: handleAbsolutePointerReport */

// handleDigitizerReport - Processes the digitizer elements and dispatches events for any updated values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleDigitizerReport
type handleDigitizerReport struct {
}/* debug [types.gen.go/struct]: handleDigitizerReport */

// handleGameControllerReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleGameControllerReport
type handleGameControllerReport struct {
}/* debug [types.gen.go/struct]: handleGameControllerReport */

// handleKeyboardReport - Iterates through keyboard elements and dispatches them if the element value has been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleKeyboardReport
type handleKeyboardReport struct {
}/* debug [types.gen.go/struct]: handleKeyboardReport */

// handleProximityReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleProximityReport
type handleProximityReport struct {
}/* debug [types.gen.go/struct]: handleProximityReport */

// handleRelativePointerReport - Iterates through relative pointer elements and dispatches them if the element value has been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleRelativePointerReport
type handleRelativePointerReport struct {
}/* debug [types.gen.go/struct]: handleRelativePointerReport */

// handleScrollReport - Iterates through scroll elements and dispatches them if the element value has been updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleScrollReport
type handleScrollReport struct {
}/* debug [types.gen.go/struct]: handleScrollReport */

// parseDigitizerElement - Parses an element to see if it supports digitizer usages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseDigitizerElement
type parseDigitizerElement struct {
}/* debug [types.gen.go/struct]: parseDigitizerElement */

// parseElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseElement
type parseElement struct {
}/* debug [types.gen.go/struct]: parseElement */

// parseElements - Parses the specified array of elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseElements
type parseElements struct {
}/* debug [types.gen.go/struct]: parseElements */

// parseGameControllerElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseGameControllerElement
type parseGameControllerElement struct {
}/* debug [types.gen.go/struct]: parseGameControllerElement */

// parseKeyboardElement - Parses an element to see if it contains keyboard-related information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseKeyboardElement
type parseKeyboardElement struct {
}/* debug [types.gen.go/struct]: parseKeyboardElement */

// parseLEDElement - Parses an element to see if it supports LED usages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseLEDElement
type parseLEDElement struct {
}/* debug [types.gen.go/struct]: parseLEDElement */

// parsePointerElement - Parses an element to see if it supports pointer usages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parsePointerElement
type parsePointerElement struct {
}/* debug [types.gen.go/struct]: parsePointerElement */

// parseProximityElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseProximityElement
type parseProximityElement struct {
}/* debug [types.gen.go/struct]: parseProximityElement */

// parseRemainingElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseRemainingElement
type parseRemainingElement struct {
}/* debug [types.gen.go/struct]: parseRemainingElement */

// parseRemainingElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseRemainingElements
type parseRemainingElements struct {
}/* debug [types.gen.go/struct]: parseRemainingElements */

// parseScrollElement - Parses an element to see if it supports scroll usages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/parseScrollElement
type parseScrollElement struct {
}/* debug [types.gen.go/struct]: parseScrollElement */

// postProcessElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/postProcessElements
type postProcessElements struct {
}/* debug [types.gen.go/struct]: postProcessElements */

// postProcessElements_internal
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/postProcessElements_internal
type postProcessElements_internal struct {
}/* debug [types.gen.go/struct]: postProcessElements_internal */

// processDigitizerElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/processDigitizerElements
type processDigitizerElements struct {
}/* debug [types.gen.go/struct]: processDigitizerElements */

// processGameControllerElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/processGameControllerElements
type processGameControllerElements struct {
}/* debug [types.gen.go/struct]: processGameControllerElements */

// setAccelerationProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setAccelerationProperties
type setAccelerationProperties struct {
}/* debug [types.gen.go/struct]: setAccelerationProperties */

// setDigitizerProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setDigitizerProperties
type setDigitizerProperties struct {
}/* debug [types.gen.go/struct]: setDigitizerProperties */

// setGameControllerProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setGameControllerProperties
type setGameControllerProperties struct {
}/* debug [types.gen.go/struct]: setGameControllerProperties */

// setKeyboardProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setKeyboardProperties
type setKeyboardProperties struct {
}/* debug [types.gen.go/struct]: setKeyboardProperties */

// setLEDProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setLEDProperties
type setLEDProperties struct {
}/* debug [types.gen.go/struct]: setLEDProperties */

// setRelativeProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setRelativeProperties
type setRelativeProperties struct {
}/* debug [types.gen.go/struct]: setRelativeProperties */

// setScrollProperties
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setScrollProperties
type setScrollProperties struct {
}/* debug [types.gen.go/struct]: setScrollProperties */

// setSurfaceDimensions
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setSurfaceDimensions
type setSurfaceDimensions struct {
}/* debug [types.gen.go/struct]: setSurfaceDimensions */

// setTipThreshold
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/setTipThreshold
type setTipThreshold struct {
}/* debug [types.gen.go/struct]: setTipThreshold */

// IOUserHIDEventService - A service that parses HID report data into elements that you can use to dispatch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService
type IOUserHIDEventService struct {
}/* debug [types.gen.go/struct]: IOUserHIDEventService */

// createReportPool
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/createReportPool
type createReportPool struct {
}/* debug [types.gen.go/struct]: createReportPool */

// dispatchExtendedGameControllerEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/dispatchExtendedGameControllerEvent
type dispatchExtendedGameControllerEvent struct {
}/* debug [types.gen.go/struct]: dispatchExtendedGameControllerEvent */

// dispatchExtendedGameControllerEventWithOptionalButtons
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/dispatchExtendedGameControllerEventWithOptionalButtons
type dispatchExtendedGameControllerEventWithOptionalButtons struct {
}/* debug [types.gen.go/struct]: dispatchExtendedGameControllerEventWithOptionalButtons */

// dispatchStandardGameControllerEvent
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/dispatchStandardGameControllerEvent
type dispatchStandardGameControllerEvent struct {
}/* debug [types.gen.go/struct]: dispatchStandardGameControllerEvent */

// IOUserUSBHostHIDDevice - A provider object for USB devices that support HID interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice
type IOUserUSBHostHIDDevice struct {
}/* debug [types.gen.go/struct]: IOUserUSBHostHIDDevice */

// cancelInputReportRetry - Cancels a retry attempt for an input report request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/cancelInputReportRetry
type cancelInputReportRetry struct {
}/* debug [types.gen.go/struct]: cancelInputReportRetry */

// CompleteInputReport - Processes the results of an asynchronous request for an input report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/CompleteInputReport
type CompleteInputReport struct {
}/* debug [types.gen.go/struct]: CompleteInputReport */

// CompleteOutputReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/CompleteOutputReport
type CompleteOutputReport struct {
}/* debug [types.gen.go/struct]: CompleteOutputReport */

// CompleteOutputRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/CompleteOutputRequest
type CompleteOutputRequest struct {
}/* debug [types.gen.go/struct]: CompleteOutputRequest */

// CompleteZLP
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/CompleteZLP
type CompleteZLP struct {
}/* debug [types.gen.go/struct]: CompleteZLP */

// copyStringAtIndex
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/copyStringAtIndex
type copyStringAtIndex struct {
}/* debug [types.gen.go/struct]: copyStringAtIndex */

// getAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/getAction
type getAction struct {
}/* debug [types.gen.go/struct]: getAction */

// getHIDDescriptorInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/getHIDDescriptorInfo
type getHIDDescriptorInfo struct {
}/* debug [types.gen.go/struct]: getHIDDescriptorInfo */

// initInputReport - Starts reading the input report from the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/initInputReport
type initInputReport struct {
}/* debug [types.gen.go/struct]: initInputReport */

// initPipes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/initPipes
type initPipes struct {
}/* debug [types.gen.go/struct]: initPipes */

// isBulkPipeSupported
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/isBulkPipeSupported
type isBulkPipeSupported struct {
}/* debug [types.gen.go/struct]: isBulkPipeSupported */

// reset - Resets the USB device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/reset
type reset struct {
}/* debug [types.gen.go/struct]: reset */

// returnAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/returnAction
type returnAction struct {
}/* debug [types.gen.go/struct]: returnAction */

// scheduleInputReportRetry - Retries a previous request for an input report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/scheduleInputReportRetry
type scheduleInputReportRetry struct {
}/* debug [types.gen.go/struct]: scheduleInputReportRetry */

// setIdle - Sets the device’s idle time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/setIdle
type setIdle struct {
}/* debug [types.gen.go/struct]: setIdle */

// setIdlePolicy - Sets the amount of idle time that must pass before suspending the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/setIdlePolicy
type setIdlePolicy struct {
}/* debug [types.gen.go/struct]: setIdlePolicy */

// setProtocol - Sets the active protocol to use for communicating with the USB device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/setProtocol
type setProtocol struct {
}/* debug [types.gen.go/struct]: setProtocol */

// TimerOccurred - Handles timeout-related actions when retrying input report requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/TimerOccurred
type TimerOccurred struct {
}/* debug [types.gen.go/struct]: TimerOccurred */

// IOHIDCompletion - A structure specifying the action to perform when a set/get report completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDCompletion
type IOHIDCompletion struct {
	Action HIDCompletionAction // The function to call.
	Parameter unsafe.Pointer // The parameter to pass to the action function.
	Target unsafe.Pointer // The target to pass to the action function.
}/* debug [types.gen.go/struct]: IOHIDCompletion */





