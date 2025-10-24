// Code generated from Apple documentation for HIDDriverKit. DO NOT EDIT.

package hiddriverkit


// C struct types
// IOHIDCompletion - A structure specifying the action to perform when a set/get report completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDCompletion
type IOHIDCompletion struct {
}// IOHIDDevice - An object containing the low-level behavior for all HID device providers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice
type IOHIDDevice struct {
}// CompleteReport - Completes all async requests made when getting or setting a report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/CompleteReport
type CompleteReport struct {
}// handleReport - Handles an asynchronous report received from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/handleReport
type handleReport struct {
}// setProperty - Updates the specified property on the corresponding kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDevice/setProperty
type setProperty struct {
}// IOHIDDigitizerCollection - A collection of elements that contain digitizer-related data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection
type IOHIDDigitizerCollection struct {
}// addElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/addElement
type addElement struct {
}// getTouch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getTouch
type getTouch struct {
}// getX
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getX
type getX struct {
}// getZ
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/getZ
type getZ struct {
}// setInRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setInRange
type setInRange struct {
}// setTouch
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/setTouch
type setTouch struct {
}// withType
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerCollection/withType
type withType struct {
}// IOHIDDigitizerStylusData - A structure containing digitizer stylus data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerStylusData
type IOHIDDigitizerStylusData struct {
}// IOHIDDigitizerTouchData - A structure containing the current digitizer touch data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDDigitizerTouchData
type IOHIDDigitizerTouchData struct {
}// IOHIDElement - An object that contains parsed information from a HID input report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement
type IOHIDElement struct {
}// getTimeStamp
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getTimeStamp
type getTimeStamp struct {
}// getUnitExponent - Returns the exponent that you use to interpret the element’s value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getUnitExponent
type getUnitExponent struct {
}// getValue - Gets the logical value that the device reported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDElement/getValue
type getValue struct {
}// IOHIDEventService - The base class for implementing a device or operating system service that dispatches events to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService
type IOHIDEventService struct {
}// Start
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/Start
type Start struct {
}// dispatchDigitizerStylusEvent - Dispatches a digitizer stylus event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchDigitizerStylusEvent
type dispatchDigitizerStylusEvent struct {
}// dispatchDigitizerTouchEvent - Dispatches a digitizer touch event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchDigitizerTouchEvent
type dispatchDigitizerTouchEvent struct {
}// dispatchRelativeScrollWheelEvent - Dispatches a relative scroll wheel event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDEventService/dispatchRelativeScrollWheelEvent
type dispatchRelativeScrollWheelEvent struct {
}// IOHIDInterface - A provider object for a HID device’s interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface
type IOHIDInterface struct {
}// GetReport - Retrieves a new input report from the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/GetReport
type GetReport struct {
}// Open - Opens a session to the device and begins the delivery of input reports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/Open
type Open struct {
}// getElements - Returns the array of elements that the interface uses to store  parsed report data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOHIDInterface/getElements
type getElements struct {
}// IOUserHIDDevice - A provider object for devices that support interactions with users.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice
type IOUserHIDDevice struct {
}// Start - Starts the device service and associates it with the specified provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice/Start
type Start struct {
}// newDeviceDescription - Creates and returns a new dictionary that describes the HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDDevice/newDeviceDescription
type newDeviceDescription struct {
}// IOUserHIDEventDriver - A complete driver object that dispatches keyboard, digitizer, scrolling, and pointer events originating from a HID device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver
type IOUserHIDEventDriver struct {
}// calibrateJustifiedPreferredStateElement
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/calibrateJustifiedPreferredStateElement
type calibrateJustifiedPreferredStateElement struct {
}// handleGameControllerReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/handleGameControllerReport
type handleGameControllerReport struct {
}// processGameControllerElements
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventDriver/processGameControllerElements
type processGameControllerElements struct {
}// IOUserHIDEventService - A service that parses HID report data into elements that you can use to dispatch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService
type IOUserHIDEventService struct {
}// Start - Starts the current event service and associates it with the specified provider object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/Start
type Start struct {
}// dispatchDigitizerStylusEvent - Dispatches a digitizer stylus event to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/dispatchDigitizerStylusEvent
type dispatchDigitizerStylusEvent struct {
}// dispatchExtendedGameControllerEventWithOptionalButtons
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/dispatchExtendedGameControllerEventWithOptionalButtons
type dispatchExtendedGameControllerEventWithOptionalButtons struct {
}// getElements - Returns an array of elements that contain the parsed data from the HID device’s report.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/getElements
type getElements struct {
}// processReport
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserHIDEventService/processReport
type processReport struct {
}// IOUserUSBHostHIDDevice - A provider object for USB devices that support HID interactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice
type IOUserUSBHostHIDDevice struct {
}// getAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/getAction
type getAction struct {
}// initInputReport - Starts reading the input report from the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/initInputReport
type initInputReport struct {
}// initPipes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/initPipes
type initPipes struct {
}// newReportDescriptor - Returns the data in the HID device’s report descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/newReportDescriptor
type newReportDescriptor struct {
}// returnAction
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/returnAction
type returnAction struct {
}// setProtocol - Sets the active protocol to use for communicating with the USB device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HIDDriverKit/IOUserUSBHostHIDDevice/setProtocol
type setProtocol struct {
}



