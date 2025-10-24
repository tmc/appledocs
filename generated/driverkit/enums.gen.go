// Code generated from Apple documentation for DriverKit. DO NOT EDIT.

package driverkit

/* debug [enums.gen.go]: Generating 5 enums for DriverKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum IOStateReporter_valueSelector (3 cases) */
// IOStateReporter_valueSelector enum type
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_valueSelector
type IOStateReporter_valueSelector uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_valueSelector/kInTransitions
	kInTransitions IOStateReporter_valueSelector = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_valueSelector/kLastTransitionTime
	kLastTransitionTime IOStateReporter_valueSelector = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/IOStateReporter_valueSelector/kResidencyTime
	kResidencyTime IOStateReporter_valueSelector = 0
)

/* debug [enums.gen.go]: Processing enum SCSIServiceResponse (6 cases) */
// SCSIServiceResponse - Attributes for task service response.
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse
type SCSIServiceResponse uint

const (
	// kSCSIServiceResponse_FUNCTION_COMPLETE - The task management function completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_FUNCTION_COMPLETE
	kSCSIServiceResponse_FUNCTION_COMPLETE SCSIServiceResponse = 0
	// kSCSIServiceResponse_FUNCTION_REJECTED - The task management function was rejected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_FUNCTION_REJECTED
	kSCSIServiceResponse_FUNCTION_REJECTED SCSIServiceResponse = 0
	// kSCSIServiceResponse_LINK_COMMAND_COMPLETE - The linked command completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_LINK_COMMAND_COMPLETE
	kSCSIServiceResponse_LINK_COMMAND_COMPLETE SCSIServiceResponse = 0
	// kSCSIServiceResponse_Request_In_Process - Not defined in SAM specification, but is a service response used for asynchronous commands that are not yet completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_Request_In_Process
	kSCSIServiceResponse_Request_In_Process SCSIServiceResponse = 0
	// kSCSIServiceResponse_SERVICE_DELIVERY_OR_TARGET_FAILURE - The service request failed because of a delivery or target failure.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_SERVICE_DELIVERY_OR_TARGET_FAILURE
	kSCSIServiceResponse_SERVICE_DELIVERY_OR_TARGET_FAILURE SCSIServiceResponse = 0
	// kSCSIServiceResponse_TASK_COMPLETE - The task completed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSIServiceResponse/kSCSIServiceResponse_TASK_COMPLETE
	kSCSIServiceResponse_TASK_COMPLETE SCSIServiceResponse = 0
)

/* debug [enums.gen.go]: Processing enum SCSITaskAttribute (4 cases) */
// SCSITaskAttribute - Attributes for task delivery.
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskAttribute
type SCSITaskAttribute uint

const (
	// kSCSITask_ACA - The task has an auto-contingent-allegiance attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskAttribute/kSCSITask_ACA
	kSCSITask_ACA SCSITaskAttribute = 0
	// kSCSITask_HEAD_OF_QUEUE - The task has a head-of-queue attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskAttribute/kSCSITask_HEAD_OF_QUEUE
	kSCSITask_HEAD_OF_QUEUE SCSITaskAttribute = 0
	// kSCSITask_ORDERED - The task has an ordered attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskAttribute/kSCSITask_ORDERED
	kSCSITask_ORDERED SCSITaskAttribute = 0
	// kSCSITask_SIMPLE - The task has a simple attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskAttribute/kSCSITask_SIMPLE
	kSCSITask_SIMPLE SCSITaskAttribute = 0
)

/* debug [enums.gen.go]: Processing enum SCSITaskState (5 cases) */
// SCSITaskState - Attributes for task state.
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState
type SCSITaskState uint

const (
	// kSCSITaskState_BLOCKED - The task is blocked.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState/kSCSITaskState_BLOCKED
	kSCSITaskState_BLOCKED SCSITaskState = 0
	// kSCSITaskState_DORMANT - The task is dormant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState/kSCSITaskState_DORMANT
	kSCSITaskState_DORMANT SCSITaskState = 0
	// kSCSITaskState_ENABLED - The task is enabled and queued.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState/kSCSITaskState_ENABLED
	kSCSITaskState_ENABLED SCSITaskState = 0
	// kSCSITaskState_ENDED - The task is complete.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState/kSCSITaskState_ENDED
	kSCSITaskState_ENDED SCSITaskState = 0
	// kSCSITaskState_NEW_TASK - The task state is new task.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskState/kSCSITaskState_NEW_TASK
	kSCSITaskState_NEW_TASK SCSITaskState = 0
)

/* debug [enums.gen.go]: Processing enum SCSITaskStatus (15 cases) */
// SCSITaskStatus - Attributes for task status.
//
// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus
type SCSITaskStatus uint

const (
	// kSCSITaskStatus_ACA_ACTIVE - The task completed with a status of ACA_ACTIVE. The device server may need the initiator to clear the Auto-Contingent Allegiance condition before it will respond to new commands.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_ACA_ACTIVE
	kSCSITaskStatus_ACA_ACTIVE SCSITaskStatus = 0
	// kSCSITaskStatus_BUSY - The task completed with a status of BUSY. The device server might need time to process a request and a delay may be required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_BUSY
	kSCSITaskStatus_BUSY SCSITaskStatus = 0
	// kSCSITaskStatus_CHECK_CONDITION - The task completed with a status of CHECK_CONDITION. Additional information about the condition should be available in the sense data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_CHECK_CONDITION
	kSCSITaskStatus_CHECK_CONDITION SCSITaskStatus = 0
	// kSCSITaskStatus_CONDITION_MET - The task completed with a status of CONDITION_MET.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_CONDITION_MET
	kSCSITaskStatus_CONDITION_MET SCSITaskStatus = 0
	// kSCSITaskStatus_DeliveryFailure - If the task is unable to be delivered to the device due to a failure in the SCSI Protocol Layer, such as a bus reset or communications error, but the device is is known to be functioning properly, the task status shall be set to kSCSITaskStatus_DeliveryFailure. This can also be reported if the task could not be delivered due to a protocol error that has since been corrected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_DeliveryFailure
	kSCSITaskStatus_DeliveryFailure SCSITaskStatus = 0
	// kSCSITaskStatus_DeviceNotPresent - If the task is unable to be delivered because the device has been detached, the task status shall be set to kSCSITaskStatus_DeviceNotPresent. This will allow the SCSI Application Layer to halt the sending of tasks to the device and, if supported, perform any device failover or system cleanup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_DeviceNotPresent
	kSCSITaskStatus_DeviceNotPresent SCSITaskStatus = 0
	// kSCSITaskStatus_DeviceNotResponding - If a task is unable to be delivered due to a failure of the device not accepting the task or the device acknowledging the attempt to send it the device the task status shall be set to kSCSITaskStatus_DeviceNotResponding. This will allow the SCSI Application driver to perform the necessary steps to try to recover the device. This shall only be reported after the SCSI Protocol Layer driver has attempted all protocol specific attempts to recover the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_DeviceNotResponding
	kSCSITaskStatus_DeviceNotResponding SCSITaskStatus = 0
	// kSCSITaskStatus_GOOD - The task completed with a status of GOOD.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_GOOD
	kSCSITaskStatus_GOOD SCSITaskStatus = 0
	// kSCSITaskStatus_INTERMEDIATE - The task completed with a status of INTERMEDIATE.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_INTERMEDIATE
	kSCSITaskStatus_INTERMEDIATE SCSITaskStatus = 0
	// kSCSITaskStatus_INTERMEDIATE_CONDITION_MET - The task completed with a status of INTERMEDIATE_CONDITION_MET.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_INTERMEDIATE_CONDITION_MET
	kSCSITaskStatus_INTERMEDIATE_CONDITION_MET SCSITaskStatus = 0
	// kSCSITaskStatus_No_Status - This status is not defined by the SCSI specifications, but is here to provide a status that can be returned in cases where there is not status available from the device or protocol, for example, when the service response is neither TASK_COMPLETED nor LINK_COMMAND_COMPLETE or when the service response is SERVICE_DELIVERY_OR_TARGET_FAILURE and the reason for failure could not be determined.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_No_Status
	kSCSITaskStatus_No_Status SCSITaskStatus = 0
	// kSCSITaskStatus_ProtocolTimeoutOccurred - If a task is aborted by the SCSI Protocol Layer due to it exceeding a timeout value specified by the support for the protocol or a related specification, the task status shall be set to kSCSITaskStatus_ProtocolTimeoutOccurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_ProtocolTimeoutOccurred
	kSCSITaskStatus_ProtocolTimeoutOccurred SCSITaskStatus = 0
	// kSCSITaskStatus_RESERVATION_CONFLICT - The task completed with a status of RESERVATION_CONFLICT.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_RESERVATION_CONFLICT
	kSCSITaskStatus_RESERVATION_CONFLICT SCSITaskStatus = 0
	// kSCSITaskStatus_TASK_SET_FULL - The task completed with a status of TASK_SET_FULL. The device server may need to complete a task before the initiator sends another.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_TASK_SET_FULL
	kSCSITaskStatus_TASK_SET_FULL SCSITaskStatus = 0
	// kSCSITaskStatus_TaskTimeoutOccurred - If a task is aborted by the SCSI Protocol Layer due to it exceeding the timeout value specified by the task, the task status shall be set to kSCSITaskStatus_TaskTimeoutOccurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/DriverKit/SCSITaskStatus/kSCSITaskStatus_TaskTimeoutOccurred
	kSCSITaskStatus_TaskTimeoutOccurred SCSITaskStatus = 0
)


