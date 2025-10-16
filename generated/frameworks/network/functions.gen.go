// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

// Network Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (409 total):

// nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor _, :  nw_advertise_descriptor_t) ->  nw_txt_record_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_advertise_descriptor_create_application_service(application_service_name _, :  UnsafePointer< CChar>) ->  nw_advertise_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_advertise_descriptor_create_bonjour_service(name _, type :  UnsafePointer< CChar>?,  _, domain :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>?) ->  nw_advertise_descriptor_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_get_application_service_name(advertise_descriptor _, :  nw_advertise_descriptor_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor _, :  nw_advertise_descriptor_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor _, no_auto_rename :  nw_advertise_descriptor_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_set_txt_record(advertise_descriptor _, txt_record :  nw_advertise_descriptor_t,  _, txt_length :  UnsafeRawPointer?,  _, :  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_advertise_descriptor_set_txt_record_object(advertise_descriptor _, txt_record :  nw_advertise_descriptor_t,  _, :  nw_txt_record_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_create_application_service(application_service_name _, :  UnsafePointer< CChar>) ->  nw_browse_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_browse_descriptor_create_bonjour_service(type _, domain :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>?) ->  nw_browse_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_get_application_service_name(descriptor _, :  nw_browse_descriptor_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_browse_descriptor_get_bonjour_service_domain(descriptor _, :  nw_browse_descriptor_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_get_bonjour_service_type(descriptor _, :  nw_browse_descriptor_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_get_include_txt_record(descriptor _, :  nw_browse_descriptor_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_descriptor_set_include_txt_record(descriptor _, include_txt_record :  nw_browse_descriptor_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_copy_endpoint(result _, :  nw_browse_result_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_copy_txt_record_object(result _, :  nw_browse_result_t) ->  nw_txt_record_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_enumerate_interfaces(result _, enumerator :  nw_browse_result_t,  _, : ( nw_interface_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_get_changes(old_result _, new_result :  nw_browse_result_t?,  _, :  nw_browse_result_t?) ->  nw_browse_result_change_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browse_result_get_interfaces_count(result _, :  nw_browse_result_t) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_cancel(browser _, :  nw_browser_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_copy_browse_descriptor(browser _, :  nw_browser_t) ->  nw_browse_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_copy_parameters(browser _, :  nw_browser_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_create(descriptor _, parameters :  nw_browse_descriptor_t,  _, :  nw_parameters_t?) ->  nw_browser_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_set_browse_results_changed_handler(browser _, handler :  nw_browser_t,  _, :  nw_browser_browse_results_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_set_queue(browser _, queue :  nw_browser_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_set_state_changed_handler(browser _, state_changed_handler :  nw_browser_t,  _, :  nw_browser_state_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_browser_start(browser _, :  nw_browser_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_connection_access_establishment_report(connection _, queue :  nw_connection_t,  _, access_block :  dispatch_queue_t,  _, :  @escaping  nw_establishment_report_access_block_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_connection_batch(connection _, batch_block :  nw_connection_t,  _, : () ->  Void) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_cancel(connection _, :  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_cancel_current_endpoint(connection _, :  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_current_path(connection _, :  nw_connection_t) ->  nw_path_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_description(connection _, :  nw_connection_t) ->  UnsafeMutablePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_endpoint(connection _, :  nw_connection_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_parameters(connection _, :  nw_connection_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_copy_protocol_metadata(connection _, definition :  nw_connection_t,  _, :  nw_protocol_definition_t) ->  nw_protocol_metadata_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_create(endpoint _, parameters :  nw_endpoint_t,  _, :  nw_parameters_t) ->  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_create_new_data_transfer_report(connection _, :  nw_connection_t) ->  nw_data_transfer_report_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_connection_force_cancel(connection _, :  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_get_maximum_datagram_size(connection _, :  nw_connection_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_group_cancel(group _, :  nw_connection_group_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_descriptor(group _, :  nw_connection_group_t) ->  nw_group_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_local_endpoint_for_message(group _, context :  nw_connection_group_t,  _, :  nw_content_context_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_parameters(group _, :  nw_connection_group_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_path_for_message(group _, context :  nw_connection_group_t,  _, :  nw_content_context_t) ->  nw_path_t?) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_copy_protocol_metadata(group _, definition :  nw_connection_group_t,  _, :  nw_protocol_definition_t) ->  nw_protocol_metadata_t?) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_copy_protocol_metadata_for_message(group _, context :  nw_connection_group_t,  _, definition :  nw_content_context_t,  _, :  nw_protocol_definition_t) ->  nw_protocol_metadata_t?) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_copy_remote_endpoint_for_message(group _, context :  nw_connection_group_t,  _, :  nw_content_context_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_create(group_descriptor _, parameters :  nw_group_descriptor_t,  _, :  nw_parameters_t) ->  nw_connection_group_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_extract_connection(group _, endpoint :  nw_connection_group_t,  _, protocol_options :  nw_endpoint_t?,  _, :  nw_protocol_options_t?) ->  nw_connection_t?) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_extract_connection_for_message(group _, context :  nw_connection_group_t,  _, :  nw_content_context_t) ->  nw_connection_t?) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_reinsert_extracted_connection(group _, connection :  nw_connection_group_t,  _, :  nw_connection_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_reply(group _, inbound_message :  nw_connection_group_t,  _, outbound_message :  nw_content_context_t,  _, content :  nw_content_context_t,  _, :  dispatch_data_t?)) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_send_message(group _, content :  nw_connection_group_t,  _, endpoint :  dispatch_data_t?,  _, context :  nw_endpoint_t?,  _, completion :  nw_content_context_t,  _, :  @escaping  nw_connection_group_send_completion_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_set_new_connection_handler(group _, new_connection_handler :  nw_connection_group_t,  _, :  nw_connection_group_new_connection_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_connection_group_set_queue(group _, queue :  nw_connection_group_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_set_receive_handler(group _, maximum_message_size :  nw_connection_group_t,  _, reject_oversized_messages :  UInt32,  _, receive_handler :  Bool,  _, :  nw_connection_group_receive_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_set_state_changed_handler(group _, state_changed_handler :  nw_connection_group_t,  _, :  nw_connection_group_state_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_group_start(group _, :  nw_connection_group_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_connection_receive(connection _, minimum_incomplete_length :  nw_connection_t,  _, maximum_length :  UInt32,  _, completion :  UInt32,  _, :  @escaping  nw_connection_receive_completion_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_receive_message(connection _, completion :  nw_connection_t,  _, :  @escaping  nw_connection_receive_completion_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_restart(connection _, :  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_send(connection _, content :  nw_connection_t,  _, context :  dispatch_data_t?,  _, is_complete :  nw_content_context_t,  _, completion :  Bool,  _, :  @escaping  nw_connection_send_completion_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_better_path_available_handler(connection _, handler :  nw_connection_t,  _, :  nw_connection_boolean_event_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_path_changed_handler(connection _, handler :  nw_connection_t,  _, :  nw_connection_path_event_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_queue(connection _, queue :  nw_connection_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_state_changed_handler(connection _, handler :  nw_connection_t,  _, :  nw_connection_state_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_set_viability_changed_handler(connection _, handler :  nw_connection_t,  _, :  nw_connection_boolean_event_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_connection_start(connection _, :  nw_connection_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_copy_antecedent(context _, :  nw_content_context_t) ->  nw_content_context_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_copy_protocol_metadata(context _, protocol :  nw_content_context_t,  _, :  nw_protocol_definition_t) ->  nw_protocol_metadata_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_create(context_identifier _, :  UnsafePointer< CChar>) ->  nw_content_context_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_foreach_protocol_metadata(context _, foreach_block :  nw_content_context_t,  _, :  @escaping  nw_protocol_definition_t,  nw_protocol_metadata_t) ->  Void) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_expiration_milliseconds(context _, :  nw_content_context_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_identifier(context _, :  nw_content_context_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_is_final(context _, :  nw_content_context_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_get_relative_priority(context _, :  nw_content_context_t) ->  Double) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_antecedent(context _, antecedent_context :  nw_content_context_t,  _, :  nw_content_context_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_expiration_milliseconds(context _, expiration_milliseconds :  nw_content_context_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_is_final(context _, is_final :  nw_content_context_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_metadata_for_protocol(context _, protocol_metadata :  nw_content_context_t,  _, :  nw_protocol_metadata_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_content_context_set_relative_priority(context _, relative_priority :  nw_content_context_t,  _, :  Double) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_data_transfer_report_collect(report _, queue :  nw_data_transfer_report_t,  _, collect_block :  dispatch_queue_t,  _, :  @escaping  nw_data_transfer_report_collect_block_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_copy_path_interface(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  nw_interface_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_duration_milliseconds(report _, :  nw_data_transfer_report_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_path_count(report _, :  nw_data_transfer_report_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_path_radio_type(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  nw_interface_radio_type_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_data_transfer_report_get_received_application_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_ip_packet_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_transport_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_transport_duplicate_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_application_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_ip_packet_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_transport_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_state(report _, :  nw_data_transfer_report_t) ->  nw_data_transfer_report_state_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_transport_rtt_variance(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report _, path_index :  nw_data_transfer_report_t,  _, :  UInt32) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_endpoint_copy_address_string(endpoint _, :  nw_endpoint_t) ->  UnsafeMutablePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_copy_port_string(endpoint _, :  nw_endpoint_t) ->  UnsafeMutablePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_copy_txt_record(endpoint _, :  nw_endpoint_t) ->  nw_txt_record_t?) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_endpoint_create_address(address _, :  UnsafePointer< sockaddr>) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_create_bonjour_service(name _, type :  UnsafePointer< CChar>,  _, domain :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_create_host(hostname _, port :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_create_url(url _, :  UnsafePointer< CChar>) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_endpoint_get_address(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< sockaddr>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_bonjour_service_domain(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_bonjour_service_name(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_bonjour_service_type(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_hostname(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_port(endpoint _, :  nw_endpoint_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_signature(endpoint _, out_signature_length :  nw_endpoint_t,  _, :  UnsafeMutablePointer< Int>) ->  UnsafePointer< UInt8>?) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_endpoint_get_type(endpoint _, :  nw_endpoint_t) ->  nw_endpoint_type_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_endpoint_get_url(endpoint _, :  nw_endpoint_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_error_copy_cf_error(error _, :  nw_error_t) ->  Unmanaged< CFError>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_error_get_error_code(error _, :  nw_error_t) ->  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_error_get_error_domain(error _, :  nw_error_t) ->  nw_error_domain_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_establishment_report_copy_proxy_endpoint(report _, :  nw_establishment_report_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_enumerate_protocols(report _, enumerate_block :  nw_establishment_report_t,  _, : ( nw_protocol_definition_t,  UInt64,  UInt64) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_enumerate_resolution_reports(report _, enumerate_block :  nw_establishment_report_t,  _, : ( nw_resolution_report_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_establishment_report_enumerate_resolutions(report _, enumerate_block :  nw_establishment_report_t,  _, : ( nw_report_resolution_source_t,  UInt64,  UInt32,  nw_endpoint_t,  nw_endpoint_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_attempt_started_after_milliseconds(report _, :  nw_establishment_report_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_duration_milliseconds(report _, :  nw_establishment_report_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_previous_attempt_count(report _, :  nw_establishment_report_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_proxy_configured(report _, :  nw_establishment_report_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_establishment_report_get_used_proxy(report _, :  nw_establishment_report_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ethernet_channel_cancel(ethernet_channel _, :  nw_ethernet_channel_t) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_create(ether_type _, interface :  UInt16,  _, :  nw_interface_t) ->  nw_ethernet_channel_t) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_create_with_parameters(ether_type _, interface :  UInt16,  _, parameters :  nw_interface_t,  _, :  nw_parameters_t) ->  nw_ethernet_channel_t) func
//
// Availability:
//   - macOS 13.0+

// nw_ethernet_channel_get_maximum_payload_size(ethernet_channel _, :  nw_ethernet_channel_t) ->  UInt32) func
//
// Availability:
//   - macOS 13.0+

// nw_ethernet_channel_send(ethernet_channel _, content :  nw_ethernet_channel_t,  _, vlan_tag :  dispatch_data_t,  _, remote_address :  UInt16,  _, completion :  UnsafeMutablePointer< UInt8>,  _, :  @escaping  nw_ethernet_channel_send_completion_t) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_set_queue(ethernet_channel _, queue :  nw_ethernet_channel_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_set_receive_handler(ethernet_channel _, handler :  nw_ethernet_channel_t,  _, :  nw_ethernet_channel_receive_handler_t?)) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_set_state_changed_handler(ethernet_channel _, handler :  nw_ethernet_channel_t,  _, :  nw_ethernet_channel_state_changed_handler_t?)) func
//
// Availability:
//   - macOS 10.15+

// nw_ethernet_channel_start(ethernet_channel _, :  nw_ethernet_channel_t) func
//
// Availability:
//   - macOS 10.15+

// nw_framer_async(framer _, async_block :  nw_framer_t,  _, :  @escaping  nw_framer_block_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_local_endpoint(framer _, :  nw_framer_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_options(framer _, :  nw_framer_t) ->  nw_protocol_options_t) func
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+

// nw_framer_copy_parameters(framer _, :  nw_framer_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_copy_remote_endpoint(framer _, :  nw_framer_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_create_definition(identifier _, flags :  UnsafePointer< CChar>,  _, start_handler :  UInt32,  _, :  @escaping  nw_framer_start_handler_t) ->  nw_protocol_definition_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_create_options(framer_definition _, :  nw_protocol_definition_t) ->  nw_protocol_options_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_deliver_input(framer _, input_buffer :  nw_framer_t,  _, input_length :  UnsafePointer< UInt8>,  _, message :  Int,  _, is_complete :  nw_framer_message_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_deliver_input_no_copy(framer _, input_length :  nw_framer_t,  _, message :  Int,  _, is_complete :  nw_framer_message_t,  _, :  Bool) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_mark_failed_with_error(framer _, error_code :  nw_framer_t,  _, :  Int32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_mark_ready(framer _, :  nw_framer_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_access_value(message _, key :  nw_framer_message_t,  _, access_value :  UnsafePointer< CChar>,  _, : ( UnsafeRawPointer?) ->  Bool) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_copy_object_value(message _, key :  nw_framer_message_t,  _, :  UnsafePointer< CChar>) ->  Any?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_create(framer _, :  nw_framer_t) ->  nw_framer_message_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_set_object_value(message _, key :  nw_framer_message_t,  _, value :  UnsafePointer< CChar>,  _, :  Any?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_message_set_value(message _, key :  nw_framer_message_t,  _, value :  UnsafePointer< CChar>,  _, dispose_value :  UnsafeMutableRawPointer?,  _, :  nw_framer_message_dispose_value_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_options_copy_object_value(options _, key :  nw_protocol_options_t,  _, :  UnsafePointer< CChar>) ->  Any?) func
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+

// nw_framer_options_set_object_value(options _, key :  nw_protocol_options_t,  _, value :  UnsafePointer< CChar>,  _, :  Any?)) func
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS 12.3+
//   - tvOS 15.4+
//   - visionOS 1.0+
//   - watchOS 8.4+

// nw_framer_parse_input(framer _, minimum_incomplete_length :  nw_framer_t,  _, maximum_length :  Int,  _, temp_buffer :  Int,  _, parse :  UnsafeMutablePointer< UInt8>?,  _, : ( UnsafeMutablePointer< UInt8>?,  Int,  Bool) ->  Int) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_parse_output(framer _, minimum_incomplete_length :  nw_framer_t,  _, maximum_length :  Int,  _, temp_buffer :  Int,  _, parse :  UnsafeMutablePointer< UInt8>?,  _, : ( UnsafeMutablePointer< UInt8>?,  Int,  Bool) ->  Int) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_pass_through_input(framer _, :  nw_framer_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_pass_through_output(framer _, :  nw_framer_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_prepend_application_protocol(framer _, protocol_options :  nw_framer_t,  _, :  nw_protocol_options_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_protocol_create_message(definition _, :  nw_protocol_definition_t) ->  nw_framer_message_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_schedule_wakeup(framer _, milliseconds :  nw_framer_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_cleanup_handler(framer _, cleanup_handler :  nw_framer_t,  _, :  @escaping  nw_framer_cleanup_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_input_handler(framer _, input_handler :  nw_framer_t,  _, :  @escaping  nw_framer_input_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_output_handler(framer _, output_handler :  nw_framer_t,  _, :  @escaping  nw_framer_output_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_stop_handler(framer _, stop_handler :  nw_framer_t,  _, :  @escaping  nw_framer_stop_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_set_wakeup_handler(framer _, wakeup_handler :  nw_framer_t,  _, :  @escaping  nw_framer_wakeup_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_write_output(framer _, output_buffer :  nw_framer_t,  _, output_length :  UnsafePointer< UInt8>,  _, :  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_write_output_data(framer _, output_data :  nw_framer_t,  _, :  dispatch_data_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_framer_write_output_no_copy(framer _, output_length :  nw_framer_t,  _, :  Int) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_group_descriptor_add_endpoint(descriptor _, endpoint :  nw_group_descriptor_t,  _, :  nw_endpoint_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_group_descriptor_create_multicast(multicast_group _, :  nw_endpoint_t) ->  nw_group_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_group_descriptor_create_multiplex(remote_endpoint _, :  nw_endpoint_t) ->  nw_group_descriptor_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_group_descriptor_enumerate_endpoints(descriptor _, enumerate_block :  nw_group_descriptor_t,  _, : ( nw_endpoint_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_interface_get_index(interface _, :  nw_interface_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_interface_get_name(interface _, :  nw_interface_t) ->  UnsafePointer< CChar>) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_interface_get_type(interface _, :  nw_interface_t) ->  nw_interface_type_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_create_metadata() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_get_ecn_flag(metadata _, :  nw_protocol_metadata_t) ->  nw_ip_ecn_flag_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_get_receive_time(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_get_service_class(metadata _, :  nw_protocol_metadata_t) ->  nw_service_class_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_set_ecn_flag(metadata _, ecn_flag :  nw_protocol_metadata_t,  _, :  nw_ip_ecn_flag_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_metadata_set_service_class(metadata _, service_class :  nw_protocol_metadata_t,  _, :  nw_service_class_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_calculate_receive_time(options _, calculate_receive_time :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_disable_fragmentation(options _, disable_fragmentation :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_disable_multicast_loopback(options _, disable_multicast_loopback :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_ip_options_set_hop_limit(options _, hop_limit :  nw_protocol_options_t,  _, :  UInt8) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_local_address_preference(options _, preference :  nw_protocol_options_t,  _, :  nw_ip_local_address_preference_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ip_options_set_use_minimum_mtu(options _, use_minimum_mtu :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ip_options_set_version(options _, version :  nw_protocol_options_t,  _, :  nw_ip_version_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_cancel(listener _, :  nw_listener_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_create(parameters _, :  nw_parameters_t) ->  nw_listener_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_create_with_connection(connection _, parameters :  nw_connection_t,  _, :  nw_parameters_t) ->  nw_listener_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_create_with_launchd_key(parameters _, launchd_key :  nw_parameters_t,  _, :  UnsafePointer< CChar>) ->  nw_listener_t) func
//
// Availability:
//   - macOS 10.14+

// nw_listener_create_with_port(port _, parameters :  UnsafePointer< CChar>,  _, :  nw_parameters_t) ->  nw_listener_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_get_new_connection_limit(listener _, :  nw_listener_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.15+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_get_port(listener _, :  nw_listener_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_advertise_descriptor(listener _, advertise_descriptor :  nw_listener_t,  _, :  nw_advertise_descriptor_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_advertised_endpoint_changed_handler(listener _, handler :  nw_listener_t,  _, :  nw_listener_advertised_endpoint_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_new_connection_group_handler(listener _, handler :  nw_listener_t,  _, :  nw_listener_new_connection_group_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_listener_set_new_connection_handler(listener _, handler :  nw_listener_t,  _, :  nw_listener_new_connection_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_new_connection_limit(listener _, new_connection_limit :  nw_listener_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.15+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_queue(listener _, queue :  nw_listener_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_set_state_changed_handler(listener _, handler :  nw_listener_t,  _, :  nw_listener_state_changed_handler_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_listener_start(listener _, :  nw_listener_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor _, :  nw_group_descriptor_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor _, disable_unicast_traffic :  nw_group_descriptor_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_multicast_group_descriptor_set_specific_source(multicast_descriptor _, source :  nw_group_descriptor_t,  _, :  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_parameters_clear_prohibited_interface_types(parameters _, :  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_clear_prohibited_interfaces(parameters _, :  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy(parameters _, :  nw_parameters_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy_default_protocol_stack(parameters _, :  nw_parameters_t) ->  nw_protocol_stack_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy_local_endpoint(parameters _, :  nw_parameters_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_copy_required_interface(parameters _, :  nw_parameters_t) ->  nw_interface_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_create() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_create_application_service() func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_parameters_create_custom_ip(custom_ip_protocol_number _, configure_ip :  UInt8,  _, :  @escaping  nw_parameters_configure_protocol_block_t) ->  nw_parameters_t) func
//
// Availability:
//   - macOS 10.15+

// nw_parameters_create_quic(configure_quic _, :  @escaping  nw_parameters_configure_protocol_block_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_parameters_create_secure_tcp(configure_tls _, configure_tcp :  @escaping  nw_parameters_configure_protocol_block_t,  _, :  @escaping  nw_parameters_configure_protocol_block_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_create_secure_udp(configure_dtls _, configure_udp :  @escaping  nw_parameters_configure_protocol_block_t,  _, :  @escaping  nw_parameters_configure_protocol_block_t) ->  nw_parameters_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_allow_ultra_constrained(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// nw_parameters_get_attribution(parameters _, :  nw_parameters_t) ->  nw_parameters_attribution_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_parameters_get_expired_dns_behavior(parameters _, :  nw_parameters_t) ->  nw_parameters_expired_dns_behavior_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_fast_open_enabled(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_include_peer_to_peer(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_local_only(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_multipath_service(parameters _, :  nw_parameters_t) ->  nw_multipath_service_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_prefer_no_proxy(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_prohibit_constrained(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_parameters_get_prohibit_expensive(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_required_interface_type(parameters _, :  nw_parameters_t) ->  nw_interface_type_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_reuse_local_address(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_get_service_class(parameters _, :  nw_parameters_t) ->  nw_service_class_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_iterate_prohibited_interface_types(parameters _, iterate_block :  nw_parameters_t,  _, : ( nw_interface_type_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_iterate_prohibited_interfaces(parameters _, iterate_block :  nw_parameters_t,  _, : ( nw_interface_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_prohibit_interface(parameters _, interface :  nw_parameters_t,  _, :  nw_interface_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_prohibit_interface_type(parameters _, interface_type :  nw_parameters_t,  _, :  nw_interface_type_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_require_interface(parameters _, interface :  nw_parameters_t,  _, :  nw_interface_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_requires_dnssec_validation(parameters _, :  nw_parameters_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_parameters_set_allow_ultra_constrained(parameters _, allow_ultra_constrained :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// nw_parameters_set_attribution(parameters _, attribution :  nw_parameters_t,  _, :  nw_parameters_attribution_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_parameters_set_expired_dns_behavior(parameters _, expired_dns_behavior :  nw_parameters_t,  _, :  nw_parameters_expired_dns_behavior_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_fast_open_enabled(parameters _, fast_open_enabled :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_include_peer_to_peer(parameters _, include_peer_to_peer :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_local_endpoint(parameters _, local_endpoint :  nw_parameters_t,  _, :  nw_endpoint_t?)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_local_only(parameters _, local_only :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_multipath_service(parameters _, multipath_service :  nw_parameters_t,  _, :  nw_multipath_service_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_prefer_no_proxy(parameters _, prefer_no_proxy :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_privacy_context(parameters _, privacy_context :  nw_parameters_t,  _, :  nw_privacy_context_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_parameters_set_prohibit_constrained(parameters _, prohibit_constrained :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_parameters_set_prohibit_expensive(parameters _, prohibit_expensive :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_required_interface_type(parameters _, interface_type :  nw_parameters_t,  _, :  nw_interface_type_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_requires_dnssec_validation(parameters _, requires_dnssec_validation :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_parameters_set_reuse_local_address(parameters _, reuse_local_address :  nw_parameters_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_parameters_set_service_class(parameters _, service_class :  nw_parameters_t,  _, :  nw_service_class_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_copy_effective_local_endpoint(path _, :  nw_path_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_copy_effective_remote_endpoint(path _, :  nw_path_t) ->  nw_endpoint_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_enumerate_gateways(path _, enumerate_block :  nw_path_t,  _, : ( nw_endpoint_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_path_enumerate_interfaces(path _, enumerate_block :  nw_path_t,  _, : ( nw_interface_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_get_link_quality(path _, :  nw_path_t) ->  nw_link_quality_t) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// nw_path_get_status(path _, :  nw_path_t) ->  nw_path_status_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_get_unsatisfied_reason(path _, :  nw_path_t) ->  nw_path_unsatisfied_reason_t) func
//
// Availability:
//   - Mac Catalyst 14.2+
//   - iOS 14.2+
//   - iPadOS 14.2+
//   - macOS 11.0+
//   - tvOS 14.2+
//   - visionOS 1.0+
//   - watchOS 7.1+

// nw_path_has_dns(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_has_ipv4(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_has_ipv6(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_constrained(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_path_is_equal(path _, other_path :  nw_path_t,  _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_expensive(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_is_ultra_constrained(path _, :  nw_path_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// nw_path_monitor_cancel(monitor _, :  nw_path_monitor_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_create() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_create_for_ethernet_channel() func
//
// Availability:
//   - macOS 13.0+

// nw_path_monitor_create_with_type(required_interface_type _, :  nw_interface_type_t) ->  nw_path_monitor_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_prohibit_interface_type(monitor _, interface_type :  nw_path_monitor_t,  _, :  nw_interface_type_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_path_monitor_set_cancel_handler(monitor _, cancel_handler :  nw_path_monitor_t,  _, :  @escaping  nw_path_monitor_cancel_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_set_queue(monitor _, queue :  nw_path_monitor_t,  _, :  dispatch_queue_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_set_update_handler(monitor _, update_handler :  nw_path_monitor_t,  _, :  @escaping  nw_path_monitor_update_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_monitor_start(monitor _, :  nw_path_monitor_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_path_uses_interface_type(path _, interface_type :  nw_path_t,  _, :  nw_interface_type_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_copy_ip_definition() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_copy_quic_definition() func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_protocol_copy_tcp_definition() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_copy_tls_definition() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_copy_udp_definition() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_copy_ws_definition() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_protocol_metadata_copy_definition(metadata _, :  nw_protocol_metadata_t) ->  nw_protocol_definition_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_framer_message(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_protocol_metadata_is_ip(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_quic(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_protocol_metadata_is_tcp(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_tls(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_udp(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_protocol_metadata_is_ws(metadata _, :  nw_protocol_metadata_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_protocol_options_is_quic(options _, :  nw_protocol_options_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_proxy_config_add_excluded_domain(config _, excluded_domain :  nw_proxy_config_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_add_match_domain(config _, match_domain :  nw_proxy_config_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_clear_excluded_domains(config _, :  nw_proxy_config_t) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_clear_match_domains(config _, :  nw_proxy_config_t) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_enumerate_excluded_domains(config _, enumerator :  nw_proxy_config_t,  _, : ( UnsafePointer< CChar>) ->  Void) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_proxy_config_enumerate_match_domains(config _, enumerator :  nw_proxy_config_t,  _, : ( UnsafePointer< CChar>) ->  Void) func
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//   - watchOS 10.0+

// nw_quic_add_tls_application_protocol(options _, application_protocol :  nw_protocol_options_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_copy_sec_protocol_metadata(metadata _, :  nw_protocol_metadata_t) ->  sec_protocol_metadata_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_copy_sec_protocol_options(options _, :  nw_protocol_options_t) ->  sec_protocol_options_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_create_options() func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_application_error(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_application_error_reason(metadata _, :  nw_protocol_metadata_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_idle_timeout(options _, :  nw_protocol_options_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_data(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_stream_data_bidirectional_local(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_stream_data_bidirectional_remote(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_stream_data_unidirectional(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_streams_bidirectional(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_initial_max_streams_unidirectional(options _, :  nw_protocol_options_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_keepalive_interval(metadata _, :  nw_protocol_metadata_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_local_max_streams_bidirectional(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_local_max_streams_unidirectional(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_max_datagram_frame_size(options _, :  nw_protocol_options_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_get_max_udp_payload_size(options _, :  nw_protocol_options_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_remote_idle_timeout(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_remote_max_streams_bidirectional(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_remote_max_streams_unidirectional(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_application_error(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_id(metadata _, :  nw_protocol_metadata_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_is_datagram(options _, :  nw_protocol_options_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_get_stream_is_unidirectional(options _, :  nw_protocol_options_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_type(stream_metadata _, :  nw_protocol_metadata_t) ->  UInt8) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_get_stream_usable_datagram_frame_size(metadata _, :  nw_protocol_metadata_t) ->  UInt16) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_application_error(metadata _, application_error :  nw_protocol_metadata_t,  _, reason :  UInt64,  _, :  UnsafePointer< CChar>?)) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_idle_timeout(options _, idle_timeout :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_data(options _, initial_max_data :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_stream_data_bidirectional_local(options _, initial_max_stream_data_bidirectional_local :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_stream_data_bidirectional_remote(options _, initial_max_stream_data_bidirectional_remote :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_stream_data_unidirectional(options _, initial_max_stream_data_unidirectional :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_streams_bidirectional(options _, initial_max_streams_bidirectional :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_initial_max_streams_unidirectional(options _, initial_max_streams_unidirectional :  nw_protocol_options_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_keepalive_interval(metadata _, keepalive_interval :  nw_protocol_metadata_t,  _, :  UInt16) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_local_max_streams_bidirectional(metadata _, max_streams_bidirectional :  nw_protocol_metadata_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_local_max_streams_unidirectional(metadata _, max_streams_unidirectional :  nw_protocol_metadata_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_max_datagram_frame_size(options _, max_datagram_frame_size :  nw_protocol_options_t,  _, :  UInt16) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_max_udp_payload_size(options _, max_udp_payload_size :  nw_protocol_options_t,  _, :  UInt16) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_stream_application_error(metadata _, application_error :  nw_protocol_metadata_t,  _, :  UInt64) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_quic_set_stream_is_datagram(options _, is_datagram :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+

// nw_quic_set_stream_is_unidirectional(options _, is_unidirectional :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_release(obj void *, );)
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_resolution_report_copy_preferred_endpoint(resolution_report _, :  nw_resolution_report_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_copy_successful_endpoint(resolution_report _, :  nw_resolution_report_t) ->  nw_endpoint_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_endpoint_count(resolution_report _, :  nw_resolution_report_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_milliseconds(resolution_report _, :  nw_resolution_report_t) ->  UInt64) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_protocol(resolution_report _, :  nw_resolution_report_t) ->  nw_report_resolution_protocol_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_resolution_report_get_source(resolution_report _, :  nw_resolution_report_t) ->  nw_report_resolution_source_t) func
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+

// nw_retain(obj void *, );) void  *
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_create_options() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_get_available_receive_buffer(metadata _, :  nw_protocol_metadata_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_get_available_send_buffer(metadata _, :  nw_protocol_metadata_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_connection_timeout(options _, connection_timeout :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_disable_ack_stretching(options _, disable_ack_stretching :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_disable_ecn(options _, disable_ecn :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_enable_fast_open(options _, enable_fast_open :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_enable_keepalive(options _, enable_keepalive :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_keepalive_count(options _, keepalive_count :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_keepalive_idle_time(options _, keepalive_idle_time :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_keepalive_interval(options _, keepalive_interval :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_maximum_segment_size(options _, maximum_segment_size :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_multipath_force_version(options _, multipath_force_version :  nw_protocol_options_t,  _, :  nw_multipath_version_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// nw_tcp_options_set_no_delay(options _, no_delay :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_no_options(options _, no_options :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_no_push(options _, no_push :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_persist_timeout(options _, persist_timeout :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_retransmit_connection_drop_time(options _, retransmit_connection_drop_time :  nw_protocol_options_t,  _, :  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tcp_options_set_retransmit_fin_drop(options _, retransmit_fin_drop :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tls_copy_sec_protocol_metadata(metadata _, :  nw_protocol_metadata_t) ->  sec_protocol_metadata_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tls_copy_sec_protocol_options(options _, :  nw_protocol_options_t) ->  sec_protocol_options_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_tls_create_options() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_txt_record_access_bytes(txt_record _, access_bytes :  nw_txt_record_t,  _, :  @escaping  nw_txt_record_access_bytes_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_access_key(txt_record _, key :  nw_txt_record_t,  _, access_value :  UnsafePointer< CChar>,  _, :  @escaping  nw_txt_record_access_key_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_apply(txt_record _, applier :  nw_txt_record_t,  _, :  @escaping  nw_txt_record_applier_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_copy(txt_record _, :  nw_txt_record_t?) ->  nw_txt_record_t?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_create_dictionary() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_create_with_bytes(txt_bytes _, txt_len :  UnsafePointer< UInt8>,  _, :  Int) ->  nw_txt_record_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_find_key(txt_record _, key :  nw_txt_record_t,  _, :  UnsafePointer< CChar>) ->  nw_txt_record_find_key_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_get_key_count(txt_record _, :  nw_txt_record_t?) ->  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_is_dictionary(txt_record _, :  nw_txt_record_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_is_equal(left _, right :  nw_txt_record_t?,  _, :  nw_txt_record_t?) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_remove_key(txt_record _, key :  nw_txt_record_t,  _, :  UnsafePointer< CChar>) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_txt_record_set_key(txt_record _, key :  nw_txt_record_t,  _, value :  UnsafePointer< CChar>,  _, value_len :  UnsafePointer< UInt8>?,  _, :  Int) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_udp_create_metadata() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_udp_create_options() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_udp_options_set_prefer_no_checksum(options _, prefer_no_checksum :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//   - watchOS 5.0+

// nw_ws_create_metadata(opcode _, :  nw_ws_opcode_t) ->  nw_protocol_metadata_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_create_options(version _, :  nw_ws_version_t) ->  nw_protocol_options_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_copy_server_response(metadata _, :  nw_protocol_metadata_t) ->  nw_ws_response_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_get_close_code(metadata _, :  nw_protocol_metadata_t) ->  nw_ws_close_code_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_get_opcode(metadata _, :  nw_protocol_metadata_t) ->  nw_ws_opcode_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_set_close_code(metadata _, close_code :  nw_protocol_metadata_t,  _, :  nw_ws_close_code_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_metadata_set_pong_handler(metadata _, client_queue :  nw_protocol_metadata_t,  _, pong_handler :  dispatch_queue_t,  _, :  @escaping  nw_ws_pong_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_add_additional_header(options _, name :  nw_protocol_options_t,  _, value :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_add_subprotocol(options _, subprotocol :  nw_protocol_options_t,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_auto_reply_ping(options _, auto_reply_ping :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_client_request_handler(options _, client_queue :  nw_protocol_options_t,  _, handler :  dispatch_queue_t,  _, :  @escaping  nw_ws_client_request_handler_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_maximum_message_size(options _, maximum_message_size :  nw_protocol_options_t,  _, :  Int) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_options_set_skip_handshake(options _, skip_handshake :  nw_protocol_options_t,  _, :  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_request_enumerate_additional_headers(request _, enumerator :  nw_ws_request_t,  _, : ( UnsafePointer< CChar>,  UnsafePointer< CChar>) ->  Bool) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_request_enumerate_subprotocols(request _, enumerator :  nw_ws_request_t,  _, : ( UnsafePointer< CChar>) ->  Bool) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_add_additional_header(response _, name :  nw_ws_response_t,  _, value :  UnsafePointer< CChar>,  _, :  UnsafePointer< CChar>)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_create(status _, selected_subprotocol :  nw_ws_response_status_t,  _, :  UnsafePointer< CChar>?) ->  nw_ws_response_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_enumerate_additional_headers(response _, enumerator :  nw_ws_response_t,  _, : ( UnsafePointer< CChar>,  UnsafePointer< CChar>) ->  Bool) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_get_selected_subprotocol(response _, :  nw_ws_response_t) ->  UnsafePointer< CChar>?) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+

// nw_ws_response_get_status(response _, :  nw_ws_response_t?) ->  nw_ws_response_status_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
