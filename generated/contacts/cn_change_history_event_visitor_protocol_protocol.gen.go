// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

// PCNChangeHistoryEventVisitor is the CNChangeHistoryEventVisitor protocol interface.
//
// An interface for receiving notice of changes to contacts and groups.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// See: doc://com.apple.contacts/documentation/Contacts/CNChangeHistoryEventVisitor
type PCNChangeHistoryEventVisitor interface {
	// Required methods
	VisitUpdateContactEvent(event ICNChangeHistoryUpdateContactEvent)/* debug [protocol_interface/required_method]: VisitUpdateContactEvent */
	VisitDropEverythingEvent(event ICNChangeHistoryDropEverythingEvent)/* debug [protocol_interface/required_method]: VisitDropEverythingEvent */
	VisitAddContactEvent(event ICNChangeHistoryAddContactEvent)/* debug [protocol_interface/required_method]: VisitAddContactEvent */
	VisitDeleteContactEvent(event ICNChangeHistoryDeleteContactEvent)/* debug [protocol_interface/required_method]: VisitDeleteContactEvent */
	// Optional methods
	VisitUpdateGroupEvent(event ICNChangeHistoryUpdateGroupEvent)
	HasVisitUpdateGroupEvent() bool
	VisitDeleteGroupEvent(event ICNChangeHistoryDeleteGroupEvent)
	HasVisitDeleteGroupEvent() bool
	VisitAddGroupEvent(event ICNChangeHistoryAddGroupEvent)
	HasVisitAddGroupEvent() bool
	VisitAddMemberToGroupEvent(event ICNChangeHistoryAddMemberToGroupEvent)
	HasVisitAddMemberToGroupEvent() bool
	VisitAddSubgroupToGroupEvent(event ICNChangeHistoryAddSubgroupToGroupEvent)
	HasVisitAddSubgroupToGroupEvent() bool
	VisitRemoveMemberFromGroupEvent(event ICNChangeHistoryRemoveMemberFromGroupEvent)
	HasVisitRemoveMemberFromGroupEvent() bool
	VisitRemoveSubgroupFromGroupEvent(event ICNChangeHistoryRemoveSubgroupFromGroupEvent)
	HasVisitRemoveSubgroupFromGroupEvent() bool
}
