// Code generated from Apple documentation for Network. DO NOT EDIT.

package network

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// Network Functions (415 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_nw_advertise_descriptor_copy_txt_record_object func(unsafe.Pointer) unsafe.Pointer
	_nw_advertise_descriptor_create_application_service func(unsafe.Pointer) unsafe.Pointer
	_nw_advertise_descriptor_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_advertise_descriptor_get_application_service_name func(unsafe.Pointer) unsafe.Pointer
	_nw_advertise_descriptor_get_no_auto_rename func(unsafe.Pointer) bool
	_nw_advertise_descriptor_set_no_auto_rename func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_advertise_descriptor_set_txt_record func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_advertise_descriptor_set_txt_record_object func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_create_application_service func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_get_application_service_name func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_get_bonjour_service_domain func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_get_bonjour_service_type func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_descriptor_get_include_txt_record func(unsafe.Pointer) bool
	_nw_browse_descriptor_set_include_txt_record func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_browse_result_copy_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_result_copy_txt_record_object func(unsafe.Pointer) unsafe.Pointer
	_nw_browse_result_enumerate_interfaces func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browse_result_get_changes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browse_result_get_interfaces_count func(unsafe.Pointer) unsafe.Pointer
	_nw_browser_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_browser_copy_browse_descriptor func(unsafe.Pointer) unsafe.Pointer
	_nw_browser_copy_parameters func(unsafe.Pointer) unsafe.Pointer
	_nw_browser_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browser_set_browse_results_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browser_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browser_set_state_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_browser_start func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_access_establishment_report func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_batch func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_cancel_current_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_copy_current_path func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_copy_description func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_copy_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_copy_parameters func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_copy_protocol_metadata func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_create_new_data_transfer_report func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_force_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_get_maximum_datagram_size func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_descriptor func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_local_endpoint_for_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_parameters func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_path_for_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_protocol_metadata func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_protocol_metadata_for_message func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_copy_remote_endpoint_for_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_extract_connection func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_extract_connection_for_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_reinsert_extracted_connection func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_connection_group_reply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_send_message func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_set_new_connection_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_set_receive_handler func(unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_set_state_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_group_start func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_receive func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_receive_message func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_restart func(unsafe.Pointer) unsafe.Pointer
	_nw_connection_send func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_nw_connection_set_better_path_available_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_set_path_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_set_state_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_set_viability_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_connection_start func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_copy_antecedent func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_copy_protocol_metadata func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_content_context_create func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_foreach_protocol_metadata func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_get_expiration_milliseconds func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_get_identifier func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_get_is_final func(unsafe.Pointer) bool
	_nw_content_context_get_relative_priority func(unsafe.Pointer) unsafe.Pointer
	_nw_content_context_set_antecedent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_content_context_set_expiration_milliseconds func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_content_context_set_is_final func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_content_context_set_metadata_for_protocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_content_context_set_relative_priority func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_collect func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_copy_path_interface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_duration_milliseconds func(unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_path_count func(unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_path_radio_type func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_received_application_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_received_ip_packet_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_received_transport_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_received_transport_duplicate_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_received_transport_out_of_order_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_sent_application_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_sent_ip_packet_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_sent_transport_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_state func(unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_transport_rtt_variance func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_copy_address_string func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_copy_port_string func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_copy_txt_record func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_create_address func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_create_bonjour_service func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_create_host func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_create_url func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_address func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_domain func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_name func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_bonjour_service_type func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_hostname func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_port func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_signature func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_type func(unsafe.Pointer) unsafe.Pointer
	_nw_endpoint_get_url func(unsafe.Pointer) unsafe.Pointer
	_nw_error_copy_cf_error func(unsafe.Pointer) unsafe.Pointer
	_nw_error_get_error_code func(unsafe.Pointer) int
	_nw_error_get_error_domain func(unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_copy_proxy_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_enumerate_protocols func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_enumerate_resolution_reports func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_enumerate_resolutions func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_get_attempt_started_after_milliseconds func(unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_get_duration_milliseconds func(unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_get_previous_attempt_count func(unsafe.Pointer) unsafe.Pointer
	_nw_establishment_report_get_proxy_configured func(unsafe.Pointer) bool
	_nw_establishment_report_get_used_proxy func(unsafe.Pointer) bool
	_nw_ethernet_channel_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_create_with_parameters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_get_maximum_payload_size func(unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_send func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_set_receive_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_set_state_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ethernet_channel_start func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_async func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_copy_local_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_copy_options func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_copy_parameters func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_copy_remote_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_create_definition func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_create_options func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_deliver_input func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_nw_framer_deliver_input_no_copy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool) bool
	_nw_framer_mark_failed_with_error func(unsafe.Pointer, int) unsafe.Pointer
	_nw_framer_mark_ready func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_message_access_value func(unsafe.Pointer, unsafe.Pointer, bool) bool
	_nw_framer_message_copy_object_value func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_message_create func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_message_set_object_value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_message_set_value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_options_copy_object_value func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_options_set_object_value func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_parse_input func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_nw_framer_parse_output func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_nw_framer_pass_through_input func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_pass_through_output func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_prepend_application_protocol func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_framer_protocol_create_message func(unsafe.Pointer) unsafe.Pointer
	_nw_framer_schedule_wakeup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_set_cleanup_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_set_input_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_set_output_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_set_stop_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_set_wakeup_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_write_output func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_write_output_data func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_framer_write_output_no_copy func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_group_descriptor_add_endpoint func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_group_descriptor_create_multicast func(unsafe.Pointer) unsafe.Pointer
	_nw_group_descriptor_create_multiplex func(unsafe.Pointer) unsafe.Pointer
	_nw_group_descriptor_enumerate_endpoints func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_interface_get_index func(unsafe.Pointer) unsafe.Pointer
	_nw_interface_get_name func(unsafe.Pointer) unsafe.Pointer
	_nw_interface_get_type func(unsafe.Pointer) unsafe.Pointer
	_nw_ip_create_metadata func() unsafe.Pointer
	_nw_ip_metadata_get_ecn_flag func(unsafe.Pointer) unsafe.Pointer
	_nw_ip_metadata_get_receive_time func(unsafe.Pointer) unsafe.Pointer
	_nw_ip_metadata_get_service_class func(unsafe.Pointer) unsafe.Pointer
	_nw_ip_metadata_set_ecn_flag func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ip_metadata_set_service_class func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ip_options_set_calculate_receive_time func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ip_options_set_disable_fragmentation func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ip_options_set_disable_multicast_loopback func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ip_options_set_hop_limit func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ip_options_set_local_address_preference func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ip_options_set_use_minimum_mtu func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ip_options_set_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_listener_create func(unsafe.Pointer) unsafe.Pointer
	_nw_listener_create_with_connection func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_create_with_launchd_key func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_create_with_port func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_get_new_connection_limit func(unsafe.Pointer) unsafe.Pointer
	_nw_listener_get_port func(unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_advertise_descriptor func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_advertised_endpoint_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_new_connection_group_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_new_connection_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_new_connection_limit func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_set_state_changed_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_listener_start func(unsafe.Pointer) unsafe.Pointer
	_nw_multicast_group_descriptor_get_disable_unicast_traffic func(unsafe.Pointer) bool
	_nw_multicast_group_descriptor_set_disable_unicast_traffic func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_multicast_group_descriptor_set_specific_source func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_clear_prohibited_interface_types func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_clear_prohibited_interfaces func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_copy func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_copy_default_protocol_stack func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_copy_local_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_copy_required_interface func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_create func() unsafe.Pointer
	_nw_parameters_create_application_service func() unsafe.Pointer
	_nw_parameters_create_custom_ip func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_create_quic func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_create_secure_tcp func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_create_secure_udp func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_get_allow_ultra_constrained func(unsafe.Pointer) bool
	_nw_parameters_get_attribution func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_get_expired_dns_behavior func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_get_fast_open_enabled func(unsafe.Pointer) bool
	_nw_parameters_get_include_peer_to_peer func(unsafe.Pointer) bool
	_nw_parameters_get_local_only func(unsafe.Pointer) bool
	_nw_parameters_get_multipath_service func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_get_prefer_no_proxy func(unsafe.Pointer) bool
	_nw_parameters_get_prohibit_constrained func(unsafe.Pointer) bool
	_nw_parameters_get_prohibit_expensive func(unsafe.Pointer) bool
	_nw_parameters_get_required_interface_type func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_get_reuse_local_address func(unsafe.Pointer) bool
	_nw_parameters_get_service_class func(unsafe.Pointer) unsafe.Pointer
	_nw_parameters_iterate_prohibited_interface_types func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_iterate_prohibited_interfaces func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_prohibit_interface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_prohibit_interface_type func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_require_interface func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_requires_dnssec_validation func(unsafe.Pointer) bool
	_nw_parameters_set_allow_ultra_constrained func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_attribution func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_expired_dns_behavior func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_fast_open_enabled func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_include_peer_to_peer func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_local_endpoint func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_local_only func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_multipath_service func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_prefer_no_proxy func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_privacy_context func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_prohibit_constrained func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_prohibit_expensive func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_required_interface_type func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_parameters_set_requires_dnssec_validation func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_reuse_local_address func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_parameters_set_service_class func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_copy_effective_local_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_path_copy_effective_remote_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_path_enumerate_gateways func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_enumerate_interfaces func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_get_link_quality func(unsafe.Pointer) unsafe.Pointer
	_nw_path_get_status func(unsafe.Pointer) unsafe.Pointer
	_nw_path_get_unsatisfied_reason func(unsafe.Pointer) unsafe.Pointer
	_nw_path_has_dns func(unsafe.Pointer) bool
	_nw_path_has_ipv4 func(unsafe.Pointer) bool
	_nw_path_has_ipv6 func(unsafe.Pointer) bool
	_nw_path_is_constrained func(unsafe.Pointer) bool
	_nw_path_is_equal func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_path_is_expensive func(unsafe.Pointer) bool
	_nw_path_is_ultra_constrained func(unsafe.Pointer) bool
	_nw_path_monitor_cancel func(unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_create func() unsafe.Pointer
	_nw_path_monitor_create_for_ethernet_channel func() unsafe.Pointer
	_nw_path_monitor_create_with_type func(unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_prohibit_interface_type func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_set_cancel_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_set_queue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_set_update_handler func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_path_monitor_start func(unsafe.Pointer) unsafe.Pointer
	_nw_path_uses_interface_type func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_protocol_copy_ip_definition func() unsafe.Pointer
	_nw_protocol_copy_quic_definition func() unsafe.Pointer
	_nw_protocol_copy_tcp_definition func() unsafe.Pointer
	_nw_protocol_copy_tls_definition func() unsafe.Pointer
	_nw_protocol_copy_udp_definition func() unsafe.Pointer
	_nw_protocol_copy_ws_definition func() unsafe.Pointer
	_nw_protocol_metadata_copy_definition func(unsafe.Pointer) unsafe.Pointer
	_nw_protocol_metadata_is_framer_message func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_ip func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_quic func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_tcp func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_tls func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_udp func(unsafe.Pointer) bool
	_nw_protocol_metadata_is_ws func(unsafe.Pointer) bool
	_nw_protocol_options_is_quic func(unsafe.Pointer) bool
	_nw_protocol_stack_clear_application_protocols func(unsafe.Pointer) unsafe.Pointer
	_nw_protocol_stack_copy_internet_protocol func(unsafe.Pointer) unsafe.Pointer
	_nw_protocol_stack_copy_transport_protocol func(unsafe.Pointer) unsafe.Pointer
	_nw_protocol_stack_iterate_application_protocols func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_protocol_stack_prepend_application_protocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_protocol_stack_set_transport_protocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_add_excluded_domain func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_add_match_domain func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_clear_excluded_domains func(unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_clear_match_domains func(unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_enumerate_excluded_domains func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_proxy_config_enumerate_match_domains func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_add_tls_application_protocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_copy_sec_protocol_metadata func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_copy_sec_protocol_options func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_create_options func() unsafe.Pointer
	_nw_quic_get_application_error func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_application_error_reason func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_idle_timeout func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_data func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_stream_data_bidirectional_local func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_stream_data_bidirectional_remote func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_stream_data_unidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_streams_bidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_initial_max_streams_unidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_keepalive_interval func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_local_max_streams_bidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_local_max_streams_unidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_max_datagram_frame_size func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_max_udp_payload_size func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_remote_idle_timeout func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_remote_max_streams_bidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_remote_max_streams_unidirectional func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_stream_application_error func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_stream_id func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_stream_is_datagram func(unsafe.Pointer) bool
	_nw_quic_get_stream_is_unidirectional func(unsafe.Pointer) bool
	_nw_quic_get_stream_type func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_get_stream_usable_datagram_frame_size func(unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_application_error func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_idle_timeout func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_data func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_stream_data_bidirectional_local func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_stream_data_bidirectional_remote func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_stream_data_unidirectional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_streams_bidirectional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_initial_max_streams_unidirectional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_keepalive_interval func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_local_max_streams_bidirectional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_local_max_streams_unidirectional func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_max_datagram_frame_size func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_max_udp_payload_size func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_stream_application_error func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_quic_set_stream_is_datagram func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_quic_set_stream_is_unidirectional func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_release func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_copy_preferred_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_copy_successful_endpoint func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_get_endpoint_count func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_get_milliseconds func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_get_protocol func(unsafe.Pointer) unsafe.Pointer
	_nw_resolution_report_get_source func(unsafe.Pointer) unsafe.Pointer
	_nw_retain func(unsafe.Pointer) unsafe.Pointer
	_nw_tcp_create_options func() unsafe.Pointer
	_nw_tcp_get_available_receive_buffer func(unsafe.Pointer) unsafe.Pointer
	_nw_tcp_get_available_send_buffer func(unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_connection_timeout func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_disable_ack_stretching func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_disable_ecn func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_enable_fast_open func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_enable_keepalive func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_keepalive_count func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_keepalive_idle_time func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_keepalive_interval func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_maximum_segment_size func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_multipath_force_version func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_no_delay func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_no_options func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_no_push func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tcp_options_set_persist_timeout func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_retransmit_connection_drop_time func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_tcp_options_set_retransmit_fin_drop func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_tls_copy_sec_protocol_metadata func(unsafe.Pointer) unsafe.Pointer
	_nw_tls_copy_sec_protocol_options func(unsafe.Pointer) unsafe.Pointer
	_nw_tls_create_options func() unsafe.Pointer
	_nw_txt_record_access_bytes func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_access_key func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_apply func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_copy func(unsafe.Pointer) unsafe.Pointer
	_nw_txt_record_create_dictionary func() unsafe.Pointer
	_nw_txt_record_create_with_bytes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_txt_record_find_key func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_txt_record_get_key_count func(unsafe.Pointer) unsafe.Pointer
	_nw_txt_record_is_dictionary func(unsafe.Pointer) bool
	_nw_txt_record_is_equal func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_remove_key func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_txt_record_set_key func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_nw_udp_create_metadata func() unsafe.Pointer
	_nw_udp_create_options func() unsafe.Pointer
	_nw_udp_options_set_prefer_no_checksum func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ws_create_metadata func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_create_options func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_metadata_copy_server_response func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_metadata_get_close_code func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_metadata_get_opcode func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_metadata_set_close_code func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_metadata_set_pong_handler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_options_add_additional_header func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_options_add_subprotocol func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_options_set_auto_reply_ping func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ws_options_set_client_request_handler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_options_set_maximum_message_size func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_options_set_skip_handshake func(unsafe.Pointer, bool) unsafe.Pointer
	_nw_ws_request_enumerate_additional_headers func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_ws_request_enumerate_subprotocols func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_ws_response_add_additional_header func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_response_create func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_nw_ws_response_enumerate_additional_headers func(unsafe.Pointer, unsafe.Pointer) bool
	_nw_ws_response_get_selected_subprotocol func(unsafe.Pointer) unsafe.Pointer
	_nw_ws_response_get_status func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_nw_advertise_descriptor_copy_txt_record_object, lib, "nw_advertise_descriptor_copy_txt_record_object")
	tryRegister(&_nw_advertise_descriptor_create_application_service, lib, "nw_advertise_descriptor_create_application_service")
	tryRegister(&_nw_advertise_descriptor_create_bonjour_service, lib, "nw_advertise_descriptor_create_bonjour_service")
	tryRegister(&_nw_advertise_descriptor_get_application_service_name, lib, "nw_advertise_descriptor_get_application_service_name")
	tryRegister(&_nw_advertise_descriptor_get_no_auto_rename, lib, "nw_advertise_descriptor_get_no_auto_rename")
	tryRegister(&_nw_advertise_descriptor_set_no_auto_rename, lib, "nw_advertise_descriptor_set_no_auto_rename")
	tryRegister(&_nw_advertise_descriptor_set_txt_record, lib, "nw_advertise_descriptor_set_txt_record")
	tryRegister(&_nw_advertise_descriptor_set_txt_record_object, lib, "nw_advertise_descriptor_set_txt_record_object")
	tryRegister(&_nw_browse_descriptor_create_application_service, lib, "nw_browse_descriptor_create_application_service")
	tryRegister(&_nw_browse_descriptor_create_bonjour_service, lib, "nw_browse_descriptor_create_bonjour_service")
	tryRegister(&_nw_browse_descriptor_get_application_service_name, lib, "nw_browse_descriptor_get_application_service_name")
	tryRegister(&_nw_browse_descriptor_get_bonjour_service_domain, lib, "nw_browse_descriptor_get_bonjour_service_domain")
	tryRegister(&_nw_browse_descriptor_get_bonjour_service_type, lib, "nw_browse_descriptor_get_bonjour_service_type")
	tryRegister(&_nw_browse_descriptor_get_include_txt_record, lib, "nw_browse_descriptor_get_include_txt_record")
	tryRegister(&_nw_browse_descriptor_set_include_txt_record, lib, "nw_browse_descriptor_set_include_txt_record")
	tryRegister(&_nw_browse_result_copy_endpoint, lib, "nw_browse_result_copy_endpoint")
	tryRegister(&_nw_browse_result_copy_txt_record_object, lib, "nw_browse_result_copy_txt_record_object")
	tryRegister(&_nw_browse_result_enumerate_interfaces, lib, "nw_browse_result_enumerate_interfaces")
	tryRegister(&_nw_browse_result_get_changes, lib, "nw_browse_result_get_changes")
	tryRegister(&_nw_browse_result_get_interfaces_count, lib, "nw_browse_result_get_interfaces_count")
	tryRegister(&_nw_browser_cancel, lib, "nw_browser_cancel")
	tryRegister(&_nw_browser_copy_browse_descriptor, lib, "nw_browser_copy_browse_descriptor")
	tryRegister(&_nw_browser_copy_parameters, lib, "nw_browser_copy_parameters")
	tryRegister(&_nw_browser_create, lib, "nw_browser_create")
	tryRegister(&_nw_browser_set_browse_results_changed_handler, lib, "nw_browser_set_browse_results_changed_handler")
	tryRegister(&_nw_browser_set_queue, lib, "nw_browser_set_queue")
	tryRegister(&_nw_browser_set_state_changed_handler, lib, "nw_browser_set_state_changed_handler")
	tryRegister(&_nw_browser_start, lib, "nw_browser_start")
	tryRegister(&_nw_connection_access_establishment_report, lib, "nw_connection_access_establishment_report")
	tryRegister(&_nw_connection_batch, lib, "nw_connection_batch")
	tryRegister(&_nw_connection_cancel, lib, "nw_connection_cancel")
	tryRegister(&_nw_connection_cancel_current_endpoint, lib, "nw_connection_cancel_current_endpoint")
	tryRegister(&_nw_connection_copy_current_path, lib, "nw_connection_copy_current_path")
	tryRegister(&_nw_connection_copy_description, lib, "nw_connection_copy_description")
	tryRegister(&_nw_connection_copy_endpoint, lib, "nw_connection_copy_endpoint")
	tryRegister(&_nw_connection_copy_parameters, lib, "nw_connection_copy_parameters")
	tryRegister(&_nw_connection_copy_protocol_metadata, lib, "nw_connection_copy_protocol_metadata")
	tryRegister(&_nw_connection_create, lib, "nw_connection_create")
	tryRegister(&_nw_connection_create_new_data_transfer_report, lib, "nw_connection_create_new_data_transfer_report")
	tryRegister(&_nw_connection_force_cancel, lib, "nw_connection_force_cancel")
	tryRegister(&_nw_connection_get_maximum_datagram_size, lib, "nw_connection_get_maximum_datagram_size")
	tryRegister(&_nw_connection_group_cancel, lib, "nw_connection_group_cancel")
	tryRegister(&_nw_connection_group_copy_descriptor, lib, "nw_connection_group_copy_descriptor")
	tryRegister(&_nw_connection_group_copy_local_endpoint_for_message, lib, "nw_connection_group_copy_local_endpoint_for_message")
	tryRegister(&_nw_connection_group_copy_parameters, lib, "nw_connection_group_copy_parameters")
	tryRegister(&_nw_connection_group_copy_path_for_message, lib, "nw_connection_group_copy_path_for_message")
	tryRegister(&_nw_connection_group_copy_protocol_metadata, lib, "nw_connection_group_copy_protocol_metadata")
	tryRegister(&_nw_connection_group_copy_protocol_metadata_for_message, lib, "nw_connection_group_copy_protocol_metadata_for_message")
	tryRegister(&_nw_connection_group_copy_remote_endpoint_for_message, lib, "nw_connection_group_copy_remote_endpoint_for_message")
	tryRegister(&_nw_connection_group_create, lib, "nw_connection_group_create")
	tryRegister(&_nw_connection_group_extract_connection, lib, "nw_connection_group_extract_connection")
	tryRegister(&_nw_connection_group_extract_connection_for_message, lib, "nw_connection_group_extract_connection_for_message")
	tryRegister(&_nw_connection_group_reinsert_extracted_connection, lib, "nw_connection_group_reinsert_extracted_connection")
	tryRegister(&_nw_connection_group_reply, lib, "nw_connection_group_reply")
	tryRegister(&_nw_connection_group_send_message, lib, "nw_connection_group_send_message")
	tryRegister(&_nw_connection_group_set_new_connection_handler, lib, "nw_connection_group_set_new_connection_handler")
	tryRegister(&_nw_connection_group_set_queue, lib, "nw_connection_group_set_queue")
	tryRegister(&_nw_connection_group_set_receive_handler, lib, "nw_connection_group_set_receive_handler")
	tryRegister(&_nw_connection_group_set_state_changed_handler, lib, "nw_connection_group_set_state_changed_handler")
	tryRegister(&_nw_connection_group_start, lib, "nw_connection_group_start")
	tryRegister(&_nw_connection_receive, lib, "nw_connection_receive")
	tryRegister(&_nw_connection_receive_message, lib, "nw_connection_receive_message")
	tryRegister(&_nw_connection_restart, lib, "nw_connection_restart")
	tryRegister(&_nw_connection_send, lib, "nw_connection_send")
	tryRegister(&_nw_connection_set_better_path_available_handler, lib, "nw_connection_set_better_path_available_handler")
	tryRegister(&_nw_connection_set_path_changed_handler, lib, "nw_connection_set_path_changed_handler")
	tryRegister(&_nw_connection_set_queue, lib, "nw_connection_set_queue")
	tryRegister(&_nw_connection_set_state_changed_handler, lib, "nw_connection_set_state_changed_handler")
	tryRegister(&_nw_connection_set_viability_changed_handler, lib, "nw_connection_set_viability_changed_handler")
	tryRegister(&_nw_connection_start, lib, "nw_connection_start")
	tryRegister(&_nw_content_context_copy_antecedent, lib, "nw_content_context_copy_antecedent")
	tryRegister(&_nw_content_context_copy_protocol_metadata, lib, "nw_content_context_copy_protocol_metadata")
	tryRegister(&_nw_content_context_create, lib, "nw_content_context_create")
	tryRegister(&_nw_content_context_foreach_protocol_metadata, lib, "nw_content_context_foreach_protocol_metadata")
	tryRegister(&_nw_content_context_get_expiration_milliseconds, lib, "nw_content_context_get_expiration_milliseconds")
	tryRegister(&_nw_content_context_get_identifier, lib, "nw_content_context_get_identifier")
	tryRegister(&_nw_content_context_get_is_final, lib, "nw_content_context_get_is_final")
	tryRegister(&_nw_content_context_get_relative_priority, lib, "nw_content_context_get_relative_priority")
	tryRegister(&_nw_content_context_set_antecedent, lib, "nw_content_context_set_antecedent")
	tryRegister(&_nw_content_context_set_expiration_milliseconds, lib, "nw_content_context_set_expiration_milliseconds")
	tryRegister(&_nw_content_context_set_is_final, lib, "nw_content_context_set_is_final")
	tryRegister(&_nw_content_context_set_metadata_for_protocol, lib, "nw_content_context_set_metadata_for_protocol")
	tryRegister(&_nw_content_context_set_relative_priority, lib, "nw_content_context_set_relative_priority")
	tryRegister(&_nw_data_transfer_report_collect, lib, "nw_data_transfer_report_collect")
	tryRegister(&_nw_data_transfer_report_copy_path_interface, lib, "nw_data_transfer_report_copy_path_interface")
	tryRegister(&_nw_data_transfer_report_get_duration_milliseconds, lib, "nw_data_transfer_report_get_duration_milliseconds")
	tryRegister(&_nw_data_transfer_report_get_path_count, lib, "nw_data_transfer_report_get_path_count")
	tryRegister(&_nw_data_transfer_report_get_path_radio_type, lib, "nw_data_transfer_report_get_path_radio_type")
	tryRegister(&_nw_data_transfer_report_get_received_application_byte_count, lib, "nw_data_transfer_report_get_received_application_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_ip_packet_count, lib, "nw_data_transfer_report_get_received_ip_packet_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_byte_count, lib, "nw_data_transfer_report_get_received_transport_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_duplicate_byte_count, lib, "nw_data_transfer_report_get_received_transport_duplicate_byte_count")
	tryRegister(&_nw_data_transfer_report_get_received_transport_out_of_order_byte_count, lib, "nw_data_transfer_report_get_received_transport_out_of_order_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_application_byte_count, lib, "nw_data_transfer_report_get_sent_application_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_ip_packet_count, lib, "nw_data_transfer_report_get_sent_ip_packet_count")
	tryRegister(&_nw_data_transfer_report_get_sent_transport_byte_count, lib, "nw_data_transfer_report_get_sent_transport_byte_count")
	tryRegister(&_nw_data_transfer_report_get_sent_transport_retransmitted_byte_count, lib, "nw_data_transfer_report_get_sent_transport_retransmitted_byte_count")
	tryRegister(&_nw_data_transfer_report_get_state, lib, "nw_data_transfer_report_get_state")
	tryRegister(&_nw_data_transfer_report_get_transport_minimum_rtt_milliseconds, lib, "nw_data_transfer_report_get_transport_minimum_rtt_milliseconds")
	tryRegister(&_nw_data_transfer_report_get_transport_rtt_variance, lib, "nw_data_transfer_report_get_transport_rtt_variance")
	tryRegister(&_nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds, lib, "nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds")
	tryRegister(&_nw_endpoint_copy_address_string, lib, "nw_endpoint_copy_address_string")
	tryRegister(&_nw_endpoint_copy_port_string, lib, "nw_endpoint_copy_port_string")
	tryRegister(&_nw_endpoint_copy_txt_record, lib, "nw_endpoint_copy_txt_record")
	tryRegister(&_nw_endpoint_create_address, lib, "nw_endpoint_create_address")
	tryRegister(&_nw_endpoint_create_bonjour_service, lib, "nw_endpoint_create_bonjour_service")
	tryRegister(&_nw_endpoint_create_host, lib, "nw_endpoint_create_host")
	tryRegister(&_nw_endpoint_create_url, lib, "nw_endpoint_create_url")
	tryRegister(&_nw_endpoint_get_address, lib, "nw_endpoint_get_address")
	tryRegister(&_nw_endpoint_get_bonjour_service_domain, lib, "nw_endpoint_get_bonjour_service_domain")
	tryRegister(&_nw_endpoint_get_bonjour_service_name, lib, "nw_endpoint_get_bonjour_service_name")
	tryRegister(&_nw_endpoint_get_bonjour_service_type, lib, "nw_endpoint_get_bonjour_service_type")
	tryRegister(&_nw_endpoint_get_hostname, lib, "nw_endpoint_get_hostname")
	tryRegister(&_nw_endpoint_get_port, lib, "nw_endpoint_get_port")
	tryRegister(&_nw_endpoint_get_signature, lib, "nw_endpoint_get_signature")
	tryRegister(&_nw_endpoint_get_type, lib, "nw_endpoint_get_type")
	tryRegister(&_nw_endpoint_get_url, lib, "nw_endpoint_get_url")
	tryRegister(&_nw_error_copy_cf_error, lib, "nw_error_copy_cf_error")
	tryRegister(&_nw_error_get_error_code, lib, "nw_error_get_error_code")
	tryRegister(&_nw_error_get_error_domain, lib, "nw_error_get_error_domain")
	tryRegister(&_nw_establishment_report_copy_proxy_endpoint, lib, "nw_establishment_report_copy_proxy_endpoint")
	tryRegister(&_nw_establishment_report_enumerate_protocols, lib, "nw_establishment_report_enumerate_protocols")
	tryRegister(&_nw_establishment_report_enumerate_resolution_reports, lib, "nw_establishment_report_enumerate_resolution_reports")
	tryRegister(&_nw_establishment_report_enumerate_resolutions, lib, "nw_establishment_report_enumerate_resolutions")
	tryRegister(&_nw_establishment_report_get_attempt_started_after_milliseconds, lib, "nw_establishment_report_get_attempt_started_after_milliseconds")
	tryRegister(&_nw_establishment_report_get_duration_milliseconds, lib, "nw_establishment_report_get_duration_milliseconds")
	tryRegister(&_nw_establishment_report_get_previous_attempt_count, lib, "nw_establishment_report_get_previous_attempt_count")
	tryRegister(&_nw_establishment_report_get_proxy_configured, lib, "nw_establishment_report_get_proxy_configured")
	tryRegister(&_nw_establishment_report_get_used_proxy, lib, "nw_establishment_report_get_used_proxy")
	tryRegister(&_nw_ethernet_channel_cancel, lib, "nw_ethernet_channel_cancel")
	tryRegister(&_nw_ethernet_channel_create, lib, "nw_ethernet_channel_create")
	tryRegister(&_nw_ethernet_channel_create_with_parameters, lib, "nw_ethernet_channel_create_with_parameters")
	tryRegister(&_nw_ethernet_channel_get_maximum_payload_size, lib, "nw_ethernet_channel_get_maximum_payload_size")
	tryRegister(&_nw_ethernet_channel_send, lib, "nw_ethernet_channel_send")
	tryRegister(&_nw_ethernet_channel_set_queue, lib, "nw_ethernet_channel_set_queue")
	tryRegister(&_nw_ethernet_channel_set_receive_handler, lib, "nw_ethernet_channel_set_receive_handler")
	tryRegister(&_nw_ethernet_channel_set_state_changed_handler, lib, "nw_ethernet_channel_set_state_changed_handler")
	tryRegister(&_nw_ethernet_channel_start, lib, "nw_ethernet_channel_start")
	tryRegister(&_nw_framer_async, lib, "nw_framer_async")
	tryRegister(&_nw_framer_copy_local_endpoint, lib, "nw_framer_copy_local_endpoint")
	tryRegister(&_nw_framer_copy_options, lib, "nw_framer_copy_options")
	tryRegister(&_nw_framer_copy_parameters, lib, "nw_framer_copy_parameters")
	tryRegister(&_nw_framer_copy_remote_endpoint, lib, "nw_framer_copy_remote_endpoint")
	tryRegister(&_nw_framer_create_definition, lib, "nw_framer_create_definition")
	tryRegister(&_nw_framer_create_options, lib, "nw_framer_create_options")
	tryRegister(&_nw_framer_deliver_input, lib, "nw_framer_deliver_input")
	tryRegister(&_nw_framer_deliver_input_no_copy, lib, "nw_framer_deliver_input_no_copy")
	tryRegister(&_nw_framer_mark_failed_with_error, lib, "nw_framer_mark_failed_with_error")
	tryRegister(&_nw_framer_mark_ready, lib, "nw_framer_mark_ready")
	tryRegister(&_nw_framer_message_access_value, lib, "nw_framer_message_access_value")
	tryRegister(&_nw_framer_message_copy_object_value, lib, "nw_framer_message_copy_object_value")
	tryRegister(&_nw_framer_message_create, lib, "nw_framer_message_create")
	tryRegister(&_nw_framer_message_set_object_value, lib, "nw_framer_message_set_object_value")
	tryRegister(&_nw_framer_message_set_value, lib, "nw_framer_message_set_value")
	tryRegister(&_nw_framer_options_copy_object_value, lib, "nw_framer_options_copy_object_value")
	tryRegister(&_nw_framer_options_set_object_value, lib, "nw_framer_options_set_object_value")
	tryRegister(&_nw_framer_parse_input, lib, "nw_framer_parse_input")
	tryRegister(&_nw_framer_parse_output, lib, "nw_framer_parse_output")
	tryRegister(&_nw_framer_pass_through_input, lib, "nw_framer_pass_through_input")
	tryRegister(&_nw_framer_pass_through_output, lib, "nw_framer_pass_through_output")
	tryRegister(&_nw_framer_prepend_application_protocol, lib, "nw_framer_prepend_application_protocol")
	tryRegister(&_nw_framer_protocol_create_message, lib, "nw_framer_protocol_create_message")
	tryRegister(&_nw_framer_schedule_wakeup, lib, "nw_framer_schedule_wakeup")
	tryRegister(&_nw_framer_set_cleanup_handler, lib, "nw_framer_set_cleanup_handler")
	tryRegister(&_nw_framer_set_input_handler, lib, "nw_framer_set_input_handler")
	tryRegister(&_nw_framer_set_output_handler, lib, "nw_framer_set_output_handler")
	tryRegister(&_nw_framer_set_stop_handler, lib, "nw_framer_set_stop_handler")
	tryRegister(&_nw_framer_set_wakeup_handler, lib, "nw_framer_set_wakeup_handler")
	tryRegister(&_nw_framer_write_output, lib, "nw_framer_write_output")
	tryRegister(&_nw_framer_write_output_data, lib, "nw_framer_write_output_data")
	tryRegister(&_nw_framer_write_output_no_copy, lib, "nw_framer_write_output_no_copy")
	tryRegister(&_nw_group_descriptor_add_endpoint, lib, "nw_group_descriptor_add_endpoint")
	tryRegister(&_nw_group_descriptor_create_multicast, lib, "nw_group_descriptor_create_multicast")
	tryRegister(&_nw_group_descriptor_create_multiplex, lib, "nw_group_descriptor_create_multiplex")
	tryRegister(&_nw_group_descriptor_enumerate_endpoints, lib, "nw_group_descriptor_enumerate_endpoints")
	tryRegister(&_nw_interface_get_index, lib, "nw_interface_get_index")
	tryRegister(&_nw_interface_get_name, lib, "nw_interface_get_name")
	tryRegister(&_nw_interface_get_type, lib, "nw_interface_get_type")
	tryRegister(&_nw_ip_create_metadata, lib, "nw_ip_create_metadata")
	tryRegister(&_nw_ip_metadata_get_ecn_flag, lib, "nw_ip_metadata_get_ecn_flag")
	tryRegister(&_nw_ip_metadata_get_receive_time, lib, "nw_ip_metadata_get_receive_time")
	tryRegister(&_nw_ip_metadata_get_service_class, lib, "nw_ip_metadata_get_service_class")
	tryRegister(&_nw_ip_metadata_set_ecn_flag, lib, "nw_ip_metadata_set_ecn_flag")
	tryRegister(&_nw_ip_metadata_set_service_class, lib, "nw_ip_metadata_set_service_class")
	tryRegister(&_nw_ip_options_set_calculate_receive_time, lib, "nw_ip_options_set_calculate_receive_time")
	tryRegister(&_nw_ip_options_set_disable_fragmentation, lib, "nw_ip_options_set_disable_fragmentation")
	tryRegister(&_nw_ip_options_set_disable_multicast_loopback, lib, "nw_ip_options_set_disable_multicast_loopback")
	tryRegister(&_nw_ip_options_set_hop_limit, lib, "nw_ip_options_set_hop_limit")
	tryRegister(&_nw_ip_options_set_local_address_preference, lib, "nw_ip_options_set_local_address_preference")
	tryRegister(&_nw_ip_options_set_use_minimum_mtu, lib, "nw_ip_options_set_use_minimum_mtu")
	tryRegister(&_nw_ip_options_set_version, lib, "nw_ip_options_set_version")
	tryRegister(&_nw_listener_cancel, lib, "nw_listener_cancel")
	tryRegister(&_nw_listener_create, lib, "nw_listener_create")
	tryRegister(&_nw_listener_create_with_connection, lib, "nw_listener_create_with_connection")
	tryRegister(&_nw_listener_create_with_launchd_key, lib, "nw_listener_create_with_launchd_key")
	tryRegister(&_nw_listener_create_with_port, lib, "nw_listener_create_with_port")
	tryRegister(&_nw_listener_get_new_connection_limit, lib, "nw_listener_get_new_connection_limit")
	tryRegister(&_nw_listener_get_port, lib, "nw_listener_get_port")
	tryRegister(&_nw_listener_set_advertise_descriptor, lib, "nw_listener_set_advertise_descriptor")
	tryRegister(&_nw_listener_set_advertised_endpoint_changed_handler, lib, "nw_listener_set_advertised_endpoint_changed_handler")
	tryRegister(&_nw_listener_set_new_connection_group_handler, lib, "nw_listener_set_new_connection_group_handler")
	tryRegister(&_nw_listener_set_new_connection_handler, lib, "nw_listener_set_new_connection_handler")
	tryRegister(&_nw_listener_set_new_connection_limit, lib, "nw_listener_set_new_connection_limit")
	tryRegister(&_nw_listener_set_queue, lib, "nw_listener_set_queue")
	tryRegister(&_nw_listener_set_state_changed_handler, lib, "nw_listener_set_state_changed_handler")
	tryRegister(&_nw_listener_start, lib, "nw_listener_start")
	tryRegister(&_nw_multicast_group_descriptor_get_disable_unicast_traffic, lib, "nw_multicast_group_descriptor_get_disable_unicast_traffic")
	tryRegister(&_nw_multicast_group_descriptor_set_disable_unicast_traffic, lib, "nw_multicast_group_descriptor_set_disable_unicast_traffic")
	tryRegister(&_nw_multicast_group_descriptor_set_specific_source, lib, "nw_multicast_group_descriptor_set_specific_source")
	tryRegister(&_nw_parameters_clear_prohibited_interface_types, lib, "nw_parameters_clear_prohibited_interface_types")
	tryRegister(&_nw_parameters_clear_prohibited_interfaces, lib, "nw_parameters_clear_prohibited_interfaces")
	tryRegister(&_nw_parameters_copy, lib, "nw_parameters_copy")
	tryRegister(&_nw_parameters_copy_default_protocol_stack, lib, "nw_parameters_copy_default_protocol_stack")
	tryRegister(&_nw_parameters_copy_local_endpoint, lib, "nw_parameters_copy_local_endpoint")
	tryRegister(&_nw_parameters_copy_required_interface, lib, "nw_parameters_copy_required_interface")
	tryRegister(&_nw_parameters_create, lib, "nw_parameters_create")
	tryRegister(&_nw_parameters_create_application_service, lib, "nw_parameters_create_application_service")
	tryRegister(&_nw_parameters_create_custom_ip, lib, "nw_parameters_create_custom_ip")
	tryRegister(&_nw_parameters_create_quic, lib, "nw_parameters_create_quic")
	tryRegister(&_nw_parameters_create_secure_tcp, lib, "nw_parameters_create_secure_tcp")
	tryRegister(&_nw_parameters_create_secure_udp, lib, "nw_parameters_create_secure_udp")
	tryRegister(&_nw_parameters_get_allow_ultra_constrained, lib, "nw_parameters_get_allow_ultra_constrained")
	tryRegister(&_nw_parameters_get_attribution, lib, "nw_parameters_get_attribution")
	tryRegister(&_nw_parameters_get_expired_dns_behavior, lib, "nw_parameters_get_expired_dns_behavior")
	tryRegister(&_nw_parameters_get_fast_open_enabled, lib, "nw_parameters_get_fast_open_enabled")
	tryRegister(&_nw_parameters_get_include_peer_to_peer, lib, "nw_parameters_get_include_peer_to_peer")
	tryRegister(&_nw_parameters_get_local_only, lib, "nw_parameters_get_local_only")
	tryRegister(&_nw_parameters_get_multipath_service, lib, "nw_parameters_get_multipath_service")
	tryRegister(&_nw_parameters_get_prefer_no_proxy, lib, "nw_parameters_get_prefer_no_proxy")
	tryRegister(&_nw_parameters_get_prohibit_constrained, lib, "nw_parameters_get_prohibit_constrained")
	tryRegister(&_nw_parameters_get_prohibit_expensive, lib, "nw_parameters_get_prohibit_expensive")
	tryRegister(&_nw_parameters_get_required_interface_type, lib, "nw_parameters_get_required_interface_type")
	tryRegister(&_nw_parameters_get_reuse_local_address, lib, "nw_parameters_get_reuse_local_address")
	tryRegister(&_nw_parameters_get_service_class, lib, "nw_parameters_get_service_class")
	tryRegister(&_nw_parameters_iterate_prohibited_interface_types, lib, "nw_parameters_iterate_prohibited_interface_types")
	tryRegister(&_nw_parameters_iterate_prohibited_interfaces, lib, "nw_parameters_iterate_prohibited_interfaces")
	tryRegister(&_nw_parameters_prohibit_interface, lib, "nw_parameters_prohibit_interface")
	tryRegister(&_nw_parameters_prohibit_interface_type, lib, "nw_parameters_prohibit_interface_type")
	tryRegister(&_nw_parameters_require_interface, lib, "nw_parameters_require_interface")
	tryRegister(&_nw_parameters_requires_dnssec_validation, lib, "nw_parameters_requires_dnssec_validation")
	tryRegister(&_nw_parameters_set_allow_ultra_constrained, lib, "nw_parameters_set_allow_ultra_constrained")
	tryRegister(&_nw_parameters_set_attribution, lib, "nw_parameters_set_attribution")
	tryRegister(&_nw_parameters_set_expired_dns_behavior, lib, "nw_parameters_set_expired_dns_behavior")
	tryRegister(&_nw_parameters_set_fast_open_enabled, lib, "nw_parameters_set_fast_open_enabled")
	tryRegister(&_nw_parameters_set_include_peer_to_peer, lib, "nw_parameters_set_include_peer_to_peer")
	tryRegister(&_nw_parameters_set_local_endpoint, lib, "nw_parameters_set_local_endpoint")
	tryRegister(&_nw_parameters_set_local_only, lib, "nw_parameters_set_local_only")
	tryRegister(&_nw_parameters_set_multipath_service, lib, "nw_parameters_set_multipath_service")
	tryRegister(&_nw_parameters_set_prefer_no_proxy, lib, "nw_parameters_set_prefer_no_proxy")
	tryRegister(&_nw_parameters_set_privacy_context, lib, "nw_parameters_set_privacy_context")
	tryRegister(&_nw_parameters_set_prohibit_constrained, lib, "nw_parameters_set_prohibit_constrained")
	tryRegister(&_nw_parameters_set_prohibit_expensive, lib, "nw_parameters_set_prohibit_expensive")
	tryRegister(&_nw_parameters_set_required_interface_type, lib, "nw_parameters_set_required_interface_type")
	tryRegister(&_nw_parameters_set_requires_dnssec_validation, lib, "nw_parameters_set_requires_dnssec_validation")
	tryRegister(&_nw_parameters_set_reuse_local_address, lib, "nw_parameters_set_reuse_local_address")
	tryRegister(&_nw_parameters_set_service_class, lib, "nw_parameters_set_service_class")
	tryRegister(&_nw_path_copy_effective_local_endpoint, lib, "nw_path_copy_effective_local_endpoint")
	tryRegister(&_nw_path_copy_effective_remote_endpoint, lib, "nw_path_copy_effective_remote_endpoint")
	tryRegister(&_nw_path_enumerate_gateways, lib, "nw_path_enumerate_gateways")
	tryRegister(&_nw_path_enumerate_interfaces, lib, "nw_path_enumerate_interfaces")
	tryRegister(&_nw_path_get_link_quality, lib, "nw_path_get_link_quality")
	tryRegister(&_nw_path_get_status, lib, "nw_path_get_status")
	tryRegister(&_nw_path_get_unsatisfied_reason, lib, "nw_path_get_unsatisfied_reason")
	tryRegister(&_nw_path_has_dns, lib, "nw_path_has_dns")
	tryRegister(&_nw_path_has_ipv4, lib, "nw_path_has_ipv4")
	tryRegister(&_nw_path_has_ipv6, lib, "nw_path_has_ipv6")
	tryRegister(&_nw_path_is_constrained, lib, "nw_path_is_constrained")
	tryRegister(&_nw_path_is_equal, lib, "nw_path_is_equal")
	tryRegister(&_nw_path_is_expensive, lib, "nw_path_is_expensive")
	tryRegister(&_nw_path_is_ultra_constrained, lib, "nw_path_is_ultra_constrained")
	tryRegister(&_nw_path_monitor_cancel, lib, "nw_path_monitor_cancel")
	tryRegister(&_nw_path_monitor_create, lib, "nw_path_monitor_create")
	tryRegister(&_nw_path_monitor_create_for_ethernet_channel, lib, "nw_path_monitor_create_for_ethernet_channel")
	tryRegister(&_nw_path_monitor_create_with_type, lib, "nw_path_monitor_create_with_type")
	tryRegister(&_nw_path_monitor_prohibit_interface_type, lib, "nw_path_monitor_prohibit_interface_type")
	tryRegister(&_nw_path_monitor_set_cancel_handler, lib, "nw_path_monitor_set_cancel_handler")
	tryRegister(&_nw_path_monitor_set_queue, lib, "nw_path_monitor_set_queue")
	tryRegister(&_nw_path_monitor_set_update_handler, lib, "nw_path_monitor_set_update_handler")
	tryRegister(&_nw_path_monitor_start, lib, "nw_path_monitor_start")
	tryRegister(&_nw_path_uses_interface_type, lib, "nw_path_uses_interface_type")
	tryRegister(&_nw_protocol_copy_ip_definition, lib, "nw_protocol_copy_ip_definition")
	tryRegister(&_nw_protocol_copy_quic_definition, lib, "nw_protocol_copy_quic_definition")
	tryRegister(&_nw_protocol_copy_tcp_definition, lib, "nw_protocol_copy_tcp_definition")
	tryRegister(&_nw_protocol_copy_tls_definition, lib, "nw_protocol_copy_tls_definition")
	tryRegister(&_nw_protocol_copy_udp_definition, lib, "nw_protocol_copy_udp_definition")
	tryRegister(&_nw_protocol_copy_ws_definition, lib, "nw_protocol_copy_ws_definition")
	tryRegister(&_nw_protocol_metadata_copy_definition, lib, "nw_protocol_metadata_copy_definition")
	tryRegister(&_nw_protocol_metadata_is_framer_message, lib, "nw_protocol_metadata_is_framer_message")
	tryRegister(&_nw_protocol_metadata_is_ip, lib, "nw_protocol_metadata_is_ip")
	tryRegister(&_nw_protocol_metadata_is_quic, lib, "nw_protocol_metadata_is_quic")
	tryRegister(&_nw_protocol_metadata_is_tcp, lib, "nw_protocol_metadata_is_tcp")
	tryRegister(&_nw_protocol_metadata_is_tls, lib, "nw_protocol_metadata_is_tls")
	tryRegister(&_nw_protocol_metadata_is_udp, lib, "nw_protocol_metadata_is_udp")
	tryRegister(&_nw_protocol_metadata_is_ws, lib, "nw_protocol_metadata_is_ws")
	tryRegister(&_nw_protocol_options_is_quic, lib, "nw_protocol_options_is_quic")
	tryRegister(&_nw_protocol_stack_clear_application_protocols, lib, "nw_protocol_stack_clear_application_protocols")
	tryRegister(&_nw_protocol_stack_copy_internet_protocol, lib, "nw_protocol_stack_copy_internet_protocol")
	tryRegister(&_nw_protocol_stack_copy_transport_protocol, lib, "nw_protocol_stack_copy_transport_protocol")
	tryRegister(&_nw_protocol_stack_iterate_application_protocols, lib, "nw_protocol_stack_iterate_application_protocols")
	tryRegister(&_nw_protocol_stack_prepend_application_protocol, lib, "nw_protocol_stack_prepend_application_protocol")
	tryRegister(&_nw_protocol_stack_set_transport_protocol, lib, "nw_protocol_stack_set_transport_protocol")
	tryRegister(&_nw_proxy_config_add_excluded_domain, lib, "nw_proxy_config_add_excluded_domain")
	tryRegister(&_nw_proxy_config_add_match_domain, lib, "nw_proxy_config_add_match_domain")
	tryRegister(&_nw_proxy_config_clear_excluded_domains, lib, "nw_proxy_config_clear_excluded_domains")
	tryRegister(&_nw_proxy_config_clear_match_domains, lib, "nw_proxy_config_clear_match_domains")
	tryRegister(&_nw_proxy_config_enumerate_excluded_domains, lib, "nw_proxy_config_enumerate_excluded_domains")
	tryRegister(&_nw_proxy_config_enumerate_match_domains, lib, "nw_proxy_config_enumerate_match_domains")
	tryRegister(&_nw_quic_add_tls_application_protocol, lib, "nw_quic_add_tls_application_protocol")
	tryRegister(&_nw_quic_copy_sec_protocol_metadata, lib, "nw_quic_copy_sec_protocol_metadata")
	tryRegister(&_nw_quic_copy_sec_protocol_options, lib, "nw_quic_copy_sec_protocol_options")
	tryRegister(&_nw_quic_create_options, lib, "nw_quic_create_options")
	tryRegister(&_nw_quic_get_application_error, lib, "nw_quic_get_application_error")
	tryRegister(&_nw_quic_get_application_error_reason, lib, "nw_quic_get_application_error_reason")
	tryRegister(&_nw_quic_get_idle_timeout, lib, "nw_quic_get_idle_timeout")
	tryRegister(&_nw_quic_get_initial_max_data, lib, "nw_quic_get_initial_max_data")
	tryRegister(&_nw_quic_get_initial_max_stream_data_bidirectional_local, lib, "nw_quic_get_initial_max_stream_data_bidirectional_local")
	tryRegister(&_nw_quic_get_initial_max_stream_data_bidirectional_remote, lib, "nw_quic_get_initial_max_stream_data_bidirectional_remote")
	tryRegister(&_nw_quic_get_initial_max_stream_data_unidirectional, lib, "nw_quic_get_initial_max_stream_data_unidirectional")
	tryRegister(&_nw_quic_get_initial_max_streams_bidirectional, lib, "nw_quic_get_initial_max_streams_bidirectional")
	tryRegister(&_nw_quic_get_initial_max_streams_unidirectional, lib, "nw_quic_get_initial_max_streams_unidirectional")
	tryRegister(&_nw_quic_get_keepalive_interval, lib, "nw_quic_get_keepalive_interval")
	tryRegister(&_nw_quic_get_local_max_streams_bidirectional, lib, "nw_quic_get_local_max_streams_bidirectional")
	tryRegister(&_nw_quic_get_local_max_streams_unidirectional, lib, "nw_quic_get_local_max_streams_unidirectional")
	tryRegister(&_nw_quic_get_max_datagram_frame_size, lib, "nw_quic_get_max_datagram_frame_size")
	tryRegister(&_nw_quic_get_max_udp_payload_size, lib, "nw_quic_get_max_udp_payload_size")
	tryRegister(&_nw_quic_get_remote_idle_timeout, lib, "nw_quic_get_remote_idle_timeout")
	tryRegister(&_nw_quic_get_remote_max_streams_bidirectional, lib, "nw_quic_get_remote_max_streams_bidirectional")
	tryRegister(&_nw_quic_get_remote_max_streams_unidirectional, lib, "nw_quic_get_remote_max_streams_unidirectional")
	tryRegister(&_nw_quic_get_stream_application_error, lib, "nw_quic_get_stream_application_error")
	tryRegister(&_nw_quic_get_stream_id, lib, "nw_quic_get_stream_id")
	tryRegister(&_nw_quic_get_stream_is_datagram, lib, "nw_quic_get_stream_is_datagram")
	tryRegister(&_nw_quic_get_stream_is_unidirectional, lib, "nw_quic_get_stream_is_unidirectional")
	tryRegister(&_nw_quic_get_stream_type, lib, "nw_quic_get_stream_type")
	tryRegister(&_nw_quic_get_stream_usable_datagram_frame_size, lib, "nw_quic_get_stream_usable_datagram_frame_size")
	tryRegister(&_nw_quic_set_application_error, lib, "nw_quic_set_application_error")
	tryRegister(&_nw_quic_set_idle_timeout, lib, "nw_quic_set_idle_timeout")
	tryRegister(&_nw_quic_set_initial_max_data, lib, "nw_quic_set_initial_max_data")
	tryRegister(&_nw_quic_set_initial_max_stream_data_bidirectional_local, lib, "nw_quic_set_initial_max_stream_data_bidirectional_local")
	tryRegister(&_nw_quic_set_initial_max_stream_data_bidirectional_remote, lib, "nw_quic_set_initial_max_stream_data_bidirectional_remote")
	tryRegister(&_nw_quic_set_initial_max_stream_data_unidirectional, lib, "nw_quic_set_initial_max_stream_data_unidirectional")
	tryRegister(&_nw_quic_set_initial_max_streams_bidirectional, lib, "nw_quic_set_initial_max_streams_bidirectional")
	tryRegister(&_nw_quic_set_initial_max_streams_unidirectional, lib, "nw_quic_set_initial_max_streams_unidirectional")
	tryRegister(&_nw_quic_set_keepalive_interval, lib, "nw_quic_set_keepalive_interval")
	tryRegister(&_nw_quic_set_local_max_streams_bidirectional, lib, "nw_quic_set_local_max_streams_bidirectional")
	tryRegister(&_nw_quic_set_local_max_streams_unidirectional, lib, "nw_quic_set_local_max_streams_unidirectional")
	tryRegister(&_nw_quic_set_max_datagram_frame_size, lib, "nw_quic_set_max_datagram_frame_size")
	tryRegister(&_nw_quic_set_max_udp_payload_size, lib, "nw_quic_set_max_udp_payload_size")
	tryRegister(&_nw_quic_set_stream_application_error, lib, "nw_quic_set_stream_application_error")
	tryRegister(&_nw_quic_set_stream_is_datagram, lib, "nw_quic_set_stream_is_datagram")
	tryRegister(&_nw_quic_set_stream_is_unidirectional, lib, "nw_quic_set_stream_is_unidirectional")
	tryRegister(&_nw_release, lib, "nw_release")
	tryRegister(&_nw_resolution_report_copy_preferred_endpoint, lib, "nw_resolution_report_copy_preferred_endpoint")
	tryRegister(&_nw_resolution_report_copy_successful_endpoint, lib, "nw_resolution_report_copy_successful_endpoint")
	tryRegister(&_nw_resolution_report_get_endpoint_count, lib, "nw_resolution_report_get_endpoint_count")
	tryRegister(&_nw_resolution_report_get_milliseconds, lib, "nw_resolution_report_get_milliseconds")
	tryRegister(&_nw_resolution_report_get_protocol, lib, "nw_resolution_report_get_protocol")
	tryRegister(&_nw_resolution_report_get_source, lib, "nw_resolution_report_get_source")
	tryRegister(&_nw_retain, lib, "nw_retain")
	tryRegister(&_nw_tcp_create_options, lib, "nw_tcp_create_options")
	tryRegister(&_nw_tcp_get_available_receive_buffer, lib, "nw_tcp_get_available_receive_buffer")
	tryRegister(&_nw_tcp_get_available_send_buffer, lib, "nw_tcp_get_available_send_buffer")
	tryRegister(&_nw_tcp_options_set_connection_timeout, lib, "nw_tcp_options_set_connection_timeout")
	tryRegister(&_nw_tcp_options_set_disable_ack_stretching, lib, "nw_tcp_options_set_disable_ack_stretching")
	tryRegister(&_nw_tcp_options_set_disable_ecn, lib, "nw_tcp_options_set_disable_ecn")
	tryRegister(&_nw_tcp_options_set_enable_fast_open, lib, "nw_tcp_options_set_enable_fast_open")
	tryRegister(&_nw_tcp_options_set_enable_keepalive, lib, "nw_tcp_options_set_enable_keepalive")
	tryRegister(&_nw_tcp_options_set_keepalive_count, lib, "nw_tcp_options_set_keepalive_count")
	tryRegister(&_nw_tcp_options_set_keepalive_idle_time, lib, "nw_tcp_options_set_keepalive_idle_time")
	tryRegister(&_nw_tcp_options_set_keepalive_interval, lib, "nw_tcp_options_set_keepalive_interval")
	tryRegister(&_nw_tcp_options_set_maximum_segment_size, lib, "nw_tcp_options_set_maximum_segment_size")
	tryRegister(&_nw_tcp_options_set_multipath_force_version, lib, "nw_tcp_options_set_multipath_force_version")
	tryRegister(&_nw_tcp_options_set_no_delay, lib, "nw_tcp_options_set_no_delay")
	tryRegister(&_nw_tcp_options_set_no_options, lib, "nw_tcp_options_set_no_options")
	tryRegister(&_nw_tcp_options_set_no_push, lib, "nw_tcp_options_set_no_push")
	tryRegister(&_nw_tcp_options_set_persist_timeout, lib, "nw_tcp_options_set_persist_timeout")
	tryRegister(&_nw_tcp_options_set_retransmit_connection_drop_time, lib, "nw_tcp_options_set_retransmit_connection_drop_time")
	tryRegister(&_nw_tcp_options_set_retransmit_fin_drop, lib, "nw_tcp_options_set_retransmit_fin_drop")
	tryRegister(&_nw_tls_copy_sec_protocol_metadata, lib, "nw_tls_copy_sec_protocol_metadata")
	tryRegister(&_nw_tls_copy_sec_protocol_options, lib, "nw_tls_copy_sec_protocol_options")
	tryRegister(&_nw_tls_create_options, lib, "nw_tls_create_options")
	tryRegister(&_nw_txt_record_access_bytes, lib, "nw_txt_record_access_bytes")
	tryRegister(&_nw_txt_record_access_key, lib, "nw_txt_record_access_key")
	tryRegister(&_nw_txt_record_apply, lib, "nw_txt_record_apply")
	tryRegister(&_nw_txt_record_copy, lib, "nw_txt_record_copy")
	tryRegister(&_nw_txt_record_create_dictionary, lib, "nw_txt_record_create_dictionary")
	tryRegister(&_nw_txt_record_create_with_bytes, lib, "nw_txt_record_create_with_bytes")
	tryRegister(&_nw_txt_record_find_key, lib, "nw_txt_record_find_key")
	tryRegister(&_nw_txt_record_get_key_count, lib, "nw_txt_record_get_key_count")
	tryRegister(&_nw_txt_record_is_dictionary, lib, "nw_txt_record_is_dictionary")
	tryRegister(&_nw_txt_record_is_equal, lib, "nw_txt_record_is_equal")
	tryRegister(&_nw_txt_record_remove_key, lib, "nw_txt_record_remove_key")
	tryRegister(&_nw_txt_record_set_key, lib, "nw_txt_record_set_key")
	tryRegister(&_nw_udp_create_metadata, lib, "nw_udp_create_metadata")
	tryRegister(&_nw_udp_create_options, lib, "nw_udp_create_options")
	tryRegister(&_nw_udp_options_set_prefer_no_checksum, lib, "nw_udp_options_set_prefer_no_checksum")
	tryRegister(&_nw_ws_create_metadata, lib, "nw_ws_create_metadata")
	tryRegister(&_nw_ws_create_options, lib, "nw_ws_create_options")
	tryRegister(&_nw_ws_metadata_copy_server_response, lib, "nw_ws_metadata_copy_server_response")
	tryRegister(&_nw_ws_metadata_get_close_code, lib, "nw_ws_metadata_get_close_code")
	tryRegister(&_nw_ws_metadata_get_opcode, lib, "nw_ws_metadata_get_opcode")
	tryRegister(&_nw_ws_metadata_set_close_code, lib, "nw_ws_metadata_set_close_code")
	tryRegister(&_nw_ws_metadata_set_pong_handler, lib, "nw_ws_metadata_set_pong_handler")
	tryRegister(&_nw_ws_options_add_additional_header, lib, "nw_ws_options_add_additional_header")
	tryRegister(&_nw_ws_options_add_subprotocol, lib, "nw_ws_options_add_subprotocol")
	tryRegister(&_nw_ws_options_set_auto_reply_ping, lib, "nw_ws_options_set_auto_reply_ping")
	tryRegister(&_nw_ws_options_set_client_request_handler, lib, "nw_ws_options_set_client_request_handler")
	tryRegister(&_nw_ws_options_set_maximum_message_size, lib, "nw_ws_options_set_maximum_message_size")
	tryRegister(&_nw_ws_options_set_skip_handshake, lib, "nw_ws_options_set_skip_handshake")
	tryRegister(&_nw_ws_request_enumerate_additional_headers, lib, "nw_ws_request_enumerate_additional_headers")
	tryRegister(&_nw_ws_request_enumerate_subprotocols, lib, "nw_ws_request_enumerate_subprotocols")
	tryRegister(&_nw_ws_response_add_additional_header, lib, "nw_ws_response_add_additional_header")
	tryRegister(&_nw_ws_response_create, lib, "nw_ws_response_create")
	tryRegister(&_nw_ws_response_enumerate_additional_headers, lib, "nw_ws_response_enumerate_additional_headers")
	tryRegister(&_nw_ws_response_get_selected_subprotocol, lib, "nw_ws_response_get_selected_subprotocol")
	tryRegister(&_nw_ws_response_get_status, lib, "nw_ws_response_get_status")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Accesses the TXT record to advertise with the service. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_copy_txt_record_object(_:)
func nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor unsafe.Pointer) unsafe.Pointer {
	return _nw_advertise_descriptor_copy_txt_record_object(advertise_descriptor)
	}


// nw_advertise_descriptor_create_application_service is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_application_service(_:)
func nw_advertise_descriptor_create_application_service(application_service_name unsafe.Pointer) unsafe.Pointer {
	return _nw_advertise_descriptor_create_application_service(application_service_name)
	}


// Initializes a Bonjour service to advertise. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_create_bonjour_service(_:_:_:)
func nw_advertise_descriptor_create_bonjour_service(name unsafe.Pointer, type_ unsafe.Pointer, domain unsafe.Pointer) unsafe.Pointer {
	return _nw_advertise_descriptor_create_bonjour_service(name, type_, domain)
	}


// nw_advertise_descriptor_get_application_service_name is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_application_service_name(_:)
func nw_advertise_descriptor_get_application_service_name(advertise_descriptor unsafe.Pointer) unsafe.Pointer {
	return _nw_advertise_descriptor_get_application_service_name(advertise_descriptor)
	}


// Checks whether the service prohibits automatic renaming in the event of a name conflict. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_get_no_auto_rename(_:)
func nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor unsafe.Pointer) bool {
	return _nw_advertise_descriptor_get_no_auto_rename(advertise_descriptor)
	}


// Sets a Boolean to indicate whether the service prohibits automatic renaming in the event of a name conflict. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_no_auto_rename(_:_:)
func nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor unsafe.Pointer, no_auto_rename bool) {
	_nw_advertise_descriptor_set_no_auto_rename(advertise_descriptor, no_auto_rename)
	}


// Sets the TXT record as a raw buffer to advertise with the service. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record(_:_:_:)
func nw_advertise_descriptor_set_txt_record(advertise_descriptor unsafe.Pointer, txt_record unsafe.Pointer, txt_length unsafe.Pointer) {
	_nw_advertise_descriptor_set_txt_record(advertise_descriptor, txt_record, txt_length)
	}


// Sets the TXT record to advertise with the service. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_advertise_descriptor_set_txt_record_object(_:_:)
func nw_advertise_descriptor_set_txt_record_object(advertise_descriptor unsafe.Pointer, txt_record unsafe.Pointer) {
	_nw_advertise_descriptor_set_txt_record_object(advertise_descriptor, txt_record)
	}


// nw_browse_descriptor_create_application_service is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_application_service(_:)
func nw_browse_descriptor_create_application_service(application_service_name unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_descriptor_create_application_service(application_service_name)
	}


// Initializes a service descriptor used to discover a Bonjour service. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_create_bonjour_service(_:_:)
func nw_browse_descriptor_create_bonjour_service(type_ unsafe.Pointer, domain unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_descriptor_create_bonjour_service(type_, domain)
	}


// nw_browse_descriptor_get_application_service_name is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_application_service_name(_:)
func nw_browse_descriptor_get_application_service_name(descriptor unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_descriptor_get_application_service_name(descriptor)
	}


// Accesses the Bonjour service domain set on a browse descriptor. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_domain(_:)
func nw_browse_descriptor_get_bonjour_service_domain(descriptor unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_descriptor_get_bonjour_service_domain(descriptor)
	}


// Accesses the Bonjour service type set on a browse descriptor. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_bonjour_service_type(_:)
func nw_browse_descriptor_get_bonjour_service_type(descriptor unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_descriptor_get_bonjour_service_type(descriptor)
	}


// Checks if the browse descriptor requires including associated TXT records with all results. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_get_include_txt_record(_:)
func nw_browse_descriptor_get_include_txt_record(descriptor unsafe.Pointer) bool {
	return _nw_browse_descriptor_get_include_txt_record(descriptor)
	}


// Requires including associated TXT records with all results generated for this service descriptor. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_descriptor_set_include_txt_record(_:_:)
func nw_browse_descriptor_set_include_txt_record(descriptor unsafe.Pointer, include_txt_record bool) {
	_nw_browse_descriptor_set_include_txt_record(descriptor, include_txt_record)
	}


// The discovered service endpoint. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_copy_endpoint(_:)
func nw_browse_result_copy_endpoint(result unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_result_copy_endpoint(result)
	}


// Accesses the TXT record associated with a discovered service. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_copy_txt_record_object(_:)
func nw_browse_result_copy_txt_record_object(result unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_result_copy_txt_record_object(result)
	}


// Enumerates the list of interfaces on which the service was discovered. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_enumerate_interfaces(_:_:)
func nw_browse_result_enumerate_interfaces(result unsafe.Pointer, enumerator unsafe.Pointer) {
	_nw_browse_result_enumerate_interfaces(result, enumerator)
	}


// Compares two discovered services and calculates changes between them. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_get_changes(_:_:)
func nw_browse_result_get_changes(old_result unsafe.Pointer, new_result unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_result_get_changes(old_result, new_result)
	}


// Accesses the number of interfaces associated with a discovered service. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browse_result_get_interfaces_count(_:)
func nw_browse_result_get_interfaces_count(result unsafe.Pointer) unsafe.Pointer {
	return _nw_browse_result_get_interfaces_count(result)
	}


// Stops browsing for services. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_cancel(_:)
func nw_browser_cancel(browser unsafe.Pointer) {
	_nw_browser_cancel(browser)
	}


// Accesses the service descriptor with which the browser was created. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_copy_browse_descriptor(_:)
func nw_browser_copy_browse_descriptor(browser unsafe.Pointer) unsafe.Pointer {
	return _nw_browser_copy_browse_descriptor(browser)
	}


// Accesses the parameters with which the browser was created. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_copy_parameters(_:)
func nw_browser_copy_parameters(browser unsafe.Pointer) unsafe.Pointer {
	return _nw_browser_copy_parameters(browser)
	}


// Initializes a browser with a type of service to discover. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_create(_:_:)
func nw_browser_create(descriptor unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_browser_create(descriptor, parameters)
	}


// Sets the handler to receive updates about discovered services. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_browse_results_changed_handler(_:_:)
func nw_browser_set_browse_results_changed_handler(browser unsafe.Pointer, handler unsafe.Pointer) {
	_nw_browser_set_browse_results_changed_handler(browser, handler)
	}


// Sets the queue on which all browser events will be delivered. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_queue(_:_:)
func nw_browser_set_queue(browser unsafe.Pointer, queue unsafe.Pointer) {
	_nw_browser_set_queue(browser, queue)
	}


// Sets a handler to receive browser state updates. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_set_state_changed_handler(_:_:)
func nw_browser_set_state_changed_handler(browser unsafe.Pointer, state_changed_handler unsafe.Pointer) {
	_nw_browser_set_state_changed_handler(browser, state_changed_handler)
	}


// Starts browsing for services. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_browser_start(_:)
func nw_browser_start(browser unsafe.Pointer) {
	_nw_browser_start(browser)
	}


// Requests a copy of the connection’s establishment report once the connection is in the ready state. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_access_establishment_report(_:_:_:)
func nw_connection_access_establishment_report(connection unsafe.Pointer, queue unsafe.Pointer, access_block unsafe.Pointer) {
	_nw_connection_access_establishment_report(connection, queue, access_block)
	}


// Defines a block in which calls to send and receive are processed as a batch to improve performance. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_batch(_:_:)
func nw_connection_batch(connection unsafe.Pointer, batch_block unsafe.Pointer) {
	_nw_connection_batch(connection, batch_block)
	}


// Cancels the connection and gracefully disconnects any established network protocols. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_cancel(_:)
func nw_connection_cancel(connection unsafe.Pointer) {
	_nw_connection_cancel(connection)
	}


// Causes the current endpoint to be rejected, allowing the connection to try another resolved address. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_cancel_current_endpoint(_:)
func nw_connection_cancel_current_endpoint(connection unsafe.Pointer) {
	_nw_connection_cancel_current_endpoint(connection)
	}


// Accesses the network path the connection is using. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_current_path(_:)
func nw_connection_copy_current_path(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_copy_current_path(connection)
	}


// Copies the description of the connection as a string. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_description(_:)
func nw_connection_copy_description(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_copy_description(connection)
	}


// Accesses the endpoint with which the connection was created. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_endpoint(_:)
func nw_connection_copy_endpoint(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_copy_endpoint(connection)
	}


// Accesses the parameters with which the connection was created. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_parameters(_:)
func nw_connection_copy_parameters(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_copy_parameters(connection)
	}


// Retrieves the connection-wide metadata for a specific protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_copy_protocol_metadata(_:_:)
func nw_connection_copy_protocol_metadata(connection unsafe.Pointer, definition unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_copy_protocol_metadata(connection, definition)
	}


// Initializes a new connection to a remote endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_create(_:_:)
func nw_connection_create(endpoint unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_create(endpoint, parameters)
	}


// Begins a new data transfer report, which can later be collected. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_create_new_data_transfer_report(_:)
func nw_connection_create_new_data_transfer_report(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_create_new_data_transfer_report(connection)
	}


// Cancels the connection and immediately disconnects any established network protocols. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_force_cancel(_:)
func nw_connection_force_cancel(connection unsafe.Pointer) {
	_nw_connection_force_cancel(connection)
	}


// Accesses the maximum size of a datagram message that can be sent on a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_get_maximum_datagram_size(_:)
func nw_connection_get_maximum_datagram_size(connection unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_get_maximum_datagram_size(connection)
	}


// Cancels the connection group object and leaves the network group. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_cancel(_:)
func nw_connection_group_cancel(group unsafe.Pointer) {
	_nw_connection_group_cancel(group)
	}


// Accesses the descriptor of the group you use to initialize the connection group. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_descriptor(_:)
func nw_connection_group_copy_descriptor(group unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_descriptor(group)
	}


// Accesses the local address and port you use to receive the message. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_local_endpoint_for_message(_:_:)
func nw_connection_group_copy_local_endpoint_for_message(group unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_local_endpoint_for_message(group, context)
	}


// Accesses the parameters with which you initialize the connection group. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_parameters(_:)
func nw_connection_group_copy_parameters(group unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_parameters(group)
	}


// Accesses the network path on which you receive the message. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_path_for_message(_:_:)
func nw_connection_group_copy_path_for_message(group unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_path_for_message(group, context)
	}


// nw_connection_group_copy_protocol_metadata is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata(_:_:)
func nw_connection_group_copy_protocol_metadata(group unsafe.Pointer, definition unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_protocol_metadata(group, definition)
	}


// nw_connection_group_copy_protocol_metadata_for_message is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_protocol_metadata_for_message(_:_:_:)
func nw_connection_group_copy_protocol_metadata_for_message(group unsafe.Pointer, context unsafe.Pointer, definition unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_protocol_metadata_for_message(group, context, definition)
	}


// Accesses the endpoint that originates the message you receive. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_copy_remote_endpoint_for_message(_:_:)
func nw_connection_group_copy_remote_endpoint_for_message(group unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_copy_remote_endpoint_for_message(group, context)
	}


// Initializes a new connection group with a group identifier. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_create(_:_:)
func nw_connection_group_create(group_descriptor unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_create(group_descriptor, parameters)
	}


// nw_connection_group_extract_connection is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection(_:_:_:)
func nw_connection_group_extract_connection(group unsafe.Pointer, endpoint unsafe.Pointer, protocol_options unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_extract_connection(group, endpoint, protocol_options)
	}


// Converts a message you receive from an endpoint into a connection object that you use for long-term communication with that endpoint. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_extract_connection_for_message(_:_:)
func nw_connection_group_extract_connection_for_message(group unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _nw_connection_group_extract_connection_for_message(group, context)
	}


// nw_connection_group_reinsert_extracted_connection is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_reinsert_extracted_connection(_:_:)
func nw_connection_group_reinsert_extracted_connection(group unsafe.Pointer, connection unsafe.Pointer) bool {
	return _nw_connection_group_reinsert_extracted_connection(group, connection)
	}


// Sends a reply to the specific endpoint that originates a group message you receive. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_reply(_:_:_:_:)
func nw_connection_group_reply(group unsafe.Pointer, inbound_message unsafe.Pointer, outbound_message unsafe.Pointer, content unsafe.Pointer) {
	_nw_connection_group_reply(group, inbound_message, outbound_message, content)
	}


// Sends data to the entire group, or to a specific member of the group. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_send_message(_:_:_:_:_:)
func nw_connection_group_send_message(group unsafe.Pointer, content unsafe.Pointer, endpoint unsafe.Pointer, context unsafe.Pointer, completion unsafe.Pointer) {
	_nw_connection_group_send_message(group, content, endpoint, context, completion)
	}


// nw_connection_group_set_new_connection_handler is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_new_connection_handler(_:_:)
func nw_connection_group_set_new_connection_handler(group unsafe.Pointer, new_connection_handler unsafe.Pointer) {
	_nw_connection_group_set_new_connection_handler(group, new_connection_handler)
	}


// Sets the queue on which you handle connection group events. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_queue(_:_:)
func nw_connection_group_set_queue(group unsafe.Pointer, queue unsafe.Pointer) {
	_nw_connection_group_set_queue(group, queue)
	}


// Sets a handler that receives inbound messages from members of the group. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_receive_handler(_:_:_:_:)
func nw_connection_group_set_receive_handler(group unsafe.Pointer, maximum_message_size unsafe.Pointer, reject_oversized_messages bool, receive_handler unsafe.Pointer) {
	_nw_connection_group_set_receive_handler(group, maximum_message_size, reject_oversized_messages, receive_handler)
	}


// Sets a handler that receives connection group state updates. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_set_state_changed_handler(_:_:)
func nw_connection_group_set_state_changed_handler(group unsafe.Pointer, state_changed_handler unsafe.Pointer) {
	_nw_connection_group_set_state_changed_handler(group, state_changed_handler)
	}


// Joins the group and registers to receive messages. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_group_start(_:)
func nw_connection_group_start(group unsafe.Pointer) {
	_nw_connection_group_start(group)
	}


// Schedules a single receive completion handler, with a range indicating how many bytes the handler can receive at one time. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_receive(_:_:_:_:)
func nw_connection_receive(connection unsafe.Pointer, minimum_incomplete_length unsafe.Pointer, maximum_length unsafe.Pointer, completion unsafe.Pointer) {
	_nw_connection_receive(connection, minimum_incomplete_length, maximum_length, completion)
	}


// Schedules a single receive completion handler for a complete message, as opposed to a range of bytes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_receive_message(_:_:)
func nw_connection_receive_message(connection unsafe.Pointer, completion unsafe.Pointer) {
	_nw_connection_receive_message(connection, completion)
	}


// Restarts a connection that is in the waiting state. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_restart(_:)
func nw_connection_restart(connection unsafe.Pointer) {
	_nw_connection_restart(connection)
	}


// Sends data on a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_send(_:_:_:_:_:)
func nw_connection_send(connection unsafe.Pointer, content unsafe.Pointer, context unsafe.Pointer, is_complete bool, completion unsafe.Pointer) {
	_nw_connection_send(connection, content, context, is_complete, completion)
	}


// Sets a handler that receives updates when an alternative network path is preferred over the current path. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_better_path_available_handler(_:_:)
func nw_connection_set_better_path_available_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	_nw_connection_set_better_path_available_handler(connection, handler)
	}


// Sets a handler that receives network path updates. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_path_changed_handler(_:_:)
func nw_connection_set_path_changed_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	_nw_connection_set_path_changed_handler(connection, handler)
	}


// Sets the queue on which all connection events are delivered. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_queue(_:_:)
func nw_connection_set_queue(connection unsafe.Pointer, queue unsafe.Pointer) {
	_nw_connection_set_queue(connection, queue)
	}


// Sets a handler to receive connection state updates. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_state_changed_handler(_:_:)
func nw_connection_set_state_changed_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	_nw_connection_set_state_changed_handler(connection, handler)
	}


// Sets a handler that receives updates when data can be sent and received. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_set_viability_changed_handler(_:_:)
func nw_connection_set_viability_changed_handler(connection unsafe.Pointer, handler unsafe.Pointer) {
	_nw_connection_set_viability_changed_handler(connection, handler)
	}


// Starts establishing a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_connection_start(_:)
func nw_connection_start(connection unsafe.Pointer) {
	_nw_connection_start(connection)
	}


// Accesses the optional message context that must be sent before the context you are sending. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_copy_antecedent(_:)
func nw_content_context_copy_antecedent(context unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_copy_antecedent(context)
	}


// Retreives the metadata associated with a specific protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_copy_protocol_metadata(_:_:)
func nw_content_context_copy_protocol_metadata(context unsafe.Pointer, protocol_ unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_copy_protocol_metadata(context, protocol_)
	}


// Initializes a custom message context. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_create(_:)
func nw_content_context_create(context_identifier unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_create(context_identifier)
	}


// Iterates through all protocol metadata associated with the message context. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_foreach_protocol_metadata(_:_:)
func nw_content_context_foreach_protocol_metadata(context unsafe.Pointer) {
	_nw_content_context_foreach_protocol_metadata(context)
	}


// Accesses the expiration set for this message context. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_expiration_milliseconds(_:)
func nw_content_context_get_expiration_milliseconds(context unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_get_expiration_milliseconds(context)
	}


// Accesses the identifier used to create this message context. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_identifier(_:)
func nw_content_context_get_identifier(context unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_get_identifier(context)
	}


// Checks whether this context represents the final message being received. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_is_final(_:)
func nw_content_context_get_is_final(context unsafe.Pointer) bool {
	return _nw_content_context_get_is_final(context)
	}


// Accesses the relative value of priority used to reorder contexts when sending. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_get_relative_priority(_:)
func nw_content_context_get_relative_priority(context unsafe.Pointer) unsafe.Pointer {
	return _nw_content_context_get_relative_priority(context)
	}


// Set an optional message context that must be sent before the context you are sending. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_antecedent(_:_:)
func nw_content_context_set_antecedent(context unsafe.Pointer, antecedent_context unsafe.Pointer) {
	_nw_content_context_set_antecedent(context, antecedent_context)
	}


// Sets the number of milliseconds after which sending the data associated with this context must begin, otherwise the data is discarded. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_expiration_milliseconds(_:_:)
func nw_content_context_set_expiration_milliseconds(context unsafe.Pointer, expiration_milliseconds unsafe.Pointer) {
	_nw_content_context_set_expiration_milliseconds(context, expiration_milliseconds)
	}


// Sets a Boolean indicating if this context represents the final message being sent. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_is_final(_:_:)
func nw_content_context_set_is_final(context unsafe.Pointer, is_final bool) {
	_nw_content_context_set_is_final(context, is_final)
	}


// Sets protocol metadata to configure per-message or per-packet properties. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_metadata_for_protocol(_:_:)
func nw_content_context_set_metadata_for_protocol(context unsafe.Pointer, protocol_metadata unsafe.Pointer) {
	_nw_content_context_set_metadata_for_protocol(context, protocol_metadata)
	}


// Sets the relative value of priority used to reorder contexts when sending. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_content_context_set_relative_priority(_:_:)
func nw_content_context_set_relative_priority(context unsafe.Pointer, relative_priority unsafe.Pointer) {
	_nw_content_context_set_relative_priority(context, relative_priority)
	}


// Stops an outstanding data transfer report and calculates the results. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_collect(_:_:_:)
func nw_data_transfer_report_collect(report unsafe.Pointer, queue unsafe.Pointer, collect_block unsafe.Pointer) {
	_nw_data_transfer_report_collect(report, queue, collect_block)
	}


// Accesses the network interface the path used. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_copy_path_interface(_:_:)
func nw_data_transfer_report_copy_path_interface(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_copy_path_interface(report, path_index)
	}


// Checks the duration of the data transfer report, from when it was started to when it was collected. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_duration_milliseconds(_:)
func nw_data_transfer_report_get_duration_milliseconds(report unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_duration_milliseconds(report)
	}


// Checks the number of valid paths in the report. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_count(_:)
func nw_data_transfer_report_get_path_count(report unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_path_count(report)
	}


// nw_data_transfer_report_get_path_radio_type is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_path_radio_type(_:_:)
func nw_data_transfer_report_get_path_radio_type(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_path_radio_type(report, path_index)
	}


// Accesses the number of bytes the connection delivered. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_application_byte_count(_:_:)
func nw_data_transfer_report_get_received_application_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_received_application_byte_count(report, path_index)
	}


// Accesses the number of IP packets the connection received. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_ip_packet_count(_:_:)
func nw_data_transfer_report_get_received_ip_packet_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_received_ip_packet_count(report, path_index)
	}


// Accesses the number of bytes the transport protocol delivered. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_received_transport_byte_count(report, path_index)
	}


// Accesses the number of duplicated bytes the transport protocol detected. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_duplicate_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_duplicate_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_received_transport_duplicate_byte_count(report, path_index)
	}


// Accesses the number of bytes the transport protocol received out of order. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_received_transport_out_of_order_byte_count(_:_:)
func nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_received_transport_out_of_order_byte_count(report, path_index)
	}


// Accesses the number of bytes sent on the connection. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_application_byte_count(_:_:)
func nw_data_transfer_report_get_sent_application_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_sent_application_byte_count(report, path_index)
	}


// Accesses the number of IP packets the connection sent. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_ip_packet_count(_:_:)
func nw_data_transfer_report_get_sent_ip_packet_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_sent_ip_packet_count(report, path_index)
	}


// Accesses the number of bytes sent into the transport protocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_byte_count(_:_:)
func nw_data_transfer_report_get_sent_transport_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_sent_transport_byte_count(report, path_index)
	}


// Accesses the number of bytes the transport protocol retransmitted. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(_:_:)
func nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_sent_transport_retransmitted_byte_count(report, path_index)
	}


// Checks whether a data transfer report is collected. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_state(_:)
func nw_data_transfer_report_get_state(report unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_state(report)
	}


// Accesses the minimum round-trip time the transport protocol measured, in milliseconds. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(_:_:)
func nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_transport_minimum_rtt_milliseconds(report, path_index)
	}


// Accesses the variance of the round-trip time the transport protocol measured. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_rtt_variance(_:_:)
func nw_data_transfer_report_get_transport_rtt_variance(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_transport_rtt_variance(report, path_index)
	}


// Accesses the smoothed round-trip time the transport protocol measured, in milliseconds. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(_:_:)
func nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report unsafe.Pointer, path_index unsafe.Pointer) unsafe.Pointer {
	return _nw_data_transfer_report_get_transport_smoothed_rtt_milliseconds(report, path_index)
	}


// Copies the address of an endpoint as a string. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_address_string(_:)
func nw_endpoint_copy_address_string(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_copy_address_string(endpoint)
	}


// Copies the port of an endpoint as a string. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_port_string(_:)
func nw_endpoint_copy_port_string(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_copy_port_string(endpoint)
	}


// nw_endpoint_copy_txt_record is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_copy_txt_record(_:)
func nw_endpoint_copy_txt_record(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_copy_txt_record(endpoint)
	}


// Creates a network endpoint with an address structure. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_address(_:)
func nw_endpoint_create_address(address unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_create_address(address)
	}


// Creates a network endpoint with a Bonjour service name, type, and domain. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_bonjour_service(_:_:_:)
func nw_endpoint_create_bonjour_service(name unsafe.Pointer, type_ unsafe.Pointer, domain unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_create_bonjour_service(name, type_, domain)
	}


// Creates a network endpoint with a hostname and port, where the hostname may be interpreted as an IP address. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_host(_:_:)
func nw_endpoint_create_host(hostname unsafe.Pointer, port unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_create_host(hostname, port)
	}


// Creates a network endpoint with a URL string. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_create_url(_:)
func nw_endpoint_create_url(url unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_create_url(url)
	}


// Accesses the address structure stored in an address endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_address(_:)
func nw_endpoint_get_address(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_address(endpoint)
	}


// Accesses the Bonjour service domain stored in an endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_domain(_:)
func nw_endpoint_get_bonjour_service_domain(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_domain(endpoint)
	}


// Accesses the Bonjour service name stored in an endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_name(_:)
func nw_endpoint_get_bonjour_service_name(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_name(endpoint)
	}


// Accesses the Bonjour service type stored in an endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_bonjour_service_type(_:)
func nw_endpoint_get_bonjour_service_type(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_bonjour_service_type(endpoint)
	}


// Accesses the hostname stored in an endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_hostname(_:)
func nw_endpoint_get_hostname(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_hostname(endpoint)
	}


// Accesses the port stored in an endpoint, in host-byte order. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_port(_:)
func nw_endpoint_get_port(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_port(endpoint)
	}


// nw_endpoint_get_signature is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_signature(_:_:)
func nw_endpoint_get_signature(endpoint unsafe.Pointer, out_signature_length unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_signature(endpoint, out_signature_length)
	}


// Accesses the type of a endpoint. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_type(_:)
func nw_endpoint_get_type(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_type(endpoint)
	}


// Accesses the URL string stored in an endpoint. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_endpoint_get_url(_:)
func nw_endpoint_get_url(endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_endpoint_get_url(endpoint)
	}


// Returns a copy of a network error. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_copy_cf_error(_:)
func nw_error_copy_cf_error(error_ unsafe.Pointer) unsafe.Pointer {
	return _nw_error_copy_cf_error(error_)
	}


// Accesses the specific code of the network error. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_get_error_code(_:)
func nw_error_get_error_code(error_ unsafe.Pointer) int {
	return _nw_error_get_error_code(error_)
	}


// Accesses the domain of the network error. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_error_get_error_domain(_:)
func nw_error_get_error_domain(error_ unsafe.Pointer) unsafe.Pointer {
	return _nw_error_get_error_domain(error_)
	}


// Accesses the endpoint of the proxy the connection used. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_copy_proxy_endpoint(_:)
func nw_establishment_report_copy_proxy_endpoint(report unsafe.Pointer) unsafe.Pointer {
	return _nw_establishment_report_copy_proxy_endpoint(report)
	}


// Iterates a list of protocol handshakes in order from first completed to last completed. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_protocols(_:_:)
func nw_establishment_report_enumerate_protocols(report unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_protocols(report, enumerate_block)
	}


// nw_establishment_report_enumerate_resolution_reports is a Network function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolution_reports(_:_:)
func nw_establishment_report_enumerate_resolution_reports(report unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_resolution_reports(report, enumerate_block)
	}


// Iterates a list of resolution steps performed during connection establishment, in order from first resolved to last resolved. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_enumerate_resolutions(_:_:)
func nw_establishment_report_enumerate_resolutions(report unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_establishment_report_enumerate_resolutions(report, enumerate_block)
	}


// Accesses the time between the call to start and the beginning of the successful connection attempt, in milliseconds. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_attempt_started_after_milliseconds(_:)
func nw_establishment_report_get_attempt_started_after_milliseconds(report unsafe.Pointer) unsafe.Pointer {
	return _nw_establishment_report_get_attempt_started_after_milliseconds(report)
	}


// Checks the total duration of the successful connection establishment attempt, from the preparing state to the ready state. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_duration_milliseconds(_:)
func nw_establishment_report_get_duration_milliseconds(report unsafe.Pointer) unsafe.Pointer {
	return _nw_establishment_report_get_duration_milliseconds(report)
	}


// Checks the number of attempts made before the successful attempt, when the connection moved from the preparing state back to the waiting state. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_previous_attempt_count(_:)
func nw_establishment_report_get_previous_attempt_count(report unsafe.Pointer) unsafe.Pointer {
	return _nw_establishment_report_get_previous_attempt_count(report)
	}


// Checks whether a proxy was configured on the connection. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_proxy_configured(_:)
func nw_establishment_report_get_proxy_configured(report unsafe.Pointer) bool {
	return _nw_establishment_report_get_proxy_configured(report)
	}


// Checks whether the connection used a proxy. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_establishment_report_get_used_proxy(_:)
func nw_establishment_report_get_used_proxy(report unsafe.Pointer) bool {
	return _nw_establishment_report_get_used_proxy(report)
	}


// Unregisters the channel from the interface. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_cancel(_:)
func nw_ethernet_channel_cancel(ethernet_channel unsafe.Pointer) {
	_nw_ethernet_channel_cancel(ethernet_channel)
	}


// Initializes an Ethernet channel on a specific interface with a custom Ethernet type. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create(_:_:)
func nw_ethernet_channel_create(ether_type unsafe.Pointer, interface_ unsafe.Pointer) unsafe.Pointer {
	return _nw_ethernet_channel_create(ether_type, interface_)
	}


// nw_ethernet_channel_create_with_parameters is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_create_with_parameters(_:_:_:)
func nw_ethernet_channel_create_with_parameters(ether_type unsafe.Pointer, interface_ unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_ethernet_channel_create_with_parameters(ether_type, interface_, parameters)
	}


// nw_ethernet_channel_get_maximum_payload_size is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_get_maximum_payload_size(_:)
func nw_ethernet_channel_get_maximum_payload_size(ethernet_channel unsafe.Pointer) unsafe.Pointer {
	return _nw_ethernet_channel_get_maximum_payload_size(ethernet_channel)
	}


// Sends a single Ethernet frame over a channel to a specific Ethernet address. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_send(_:_:_:_:_:)
func nw_ethernet_channel_send(ethernet_channel unsafe.Pointer, content unsafe.Pointer, vlan_tag unsafe.Pointer, remote_address unsafe.Pointer, completion unsafe.Pointer) {
	_nw_ethernet_channel_send(ethernet_channel, content, vlan_tag, remote_address, completion)
	}


// Sets the queue on which all channel events are delivered. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_queue(_:_:)
func nw_ethernet_channel_set_queue(ethernet_channel unsafe.Pointer, queue unsafe.Pointer) {
	_nw_ethernet_channel_set_queue(ethernet_channel, queue)
	}


// Sets a handler to receive inbound Ethernet frames. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_receive_handler(_:_:)
func nw_ethernet_channel_set_receive_handler(ethernet_channel unsafe.Pointer, handler unsafe.Pointer) {
	_nw_ethernet_channel_set_receive_handler(ethernet_channel, handler)
	}


// Sets a handler to receive channel state updates. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_set_state_changed_handler(_:_:)
func nw_ethernet_channel_set_state_changed_handler(ethernet_channel unsafe.Pointer, handler unsafe.Pointer) {
	_nw_ethernet_channel_set_state_changed_handler(ethernet_channel, handler)
	}


// Starts the process of registering the channel. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ethernet_channel_start(_:)
func nw_ethernet_channel_start(ethernet_channel unsafe.Pointer) {
	_nw_ethernet_channel_start(ethernet_channel)
	}


// Requests that a block be executed on the connection’s internal scheduling context. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_async(_:_:)
func nw_framer_async(framer unsafe.Pointer, async_block unsafe.Pointer) {
	_nw_framer_async(framer, async_block)
	}


// Accesses the local endpoint of the connection in which your protocol is running. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_local_endpoint(_:)
func nw_framer_copy_local_endpoint(framer unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_copy_local_endpoint(framer)
	}


// nw_framer_copy_options is a Network function. [Full Topic]
//
// Added in macOS 12.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_options(_:)
func nw_framer_copy_options(framer unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_copy_options(framer)
	}


// Accesses the parameters of the connection in which your protocol is running. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_parameters(_:)
func nw_framer_copy_parameters(framer unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_copy_parameters(framer)
	}


// Accesses the remote endpoint of the connection in which your protocol is running. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_copy_remote_endpoint(_:)
func nw_framer_copy_remote_endpoint(framer unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_copy_remote_endpoint(framer)
	}


// Initializes a new protocol definition based on your protocol implementation. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_create_definition(_:_:_:)
func nw_framer_create_definition(identifier unsafe.Pointer, flags unsafe.Pointer, start_handler unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_create_definition(identifier, flags, start_handler)
	}


// Initializes a set of protocol options with a custom framer definition. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_create_options(_:)
func nw_framer_create_options(framer_definition unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_create_options(framer_definition)
	}


// Delivers an inbound message containing arbitrary data from your protocol to the application. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_deliver_input(_:_:_:_:_:)
func nw_framer_deliver_input(framer unsafe.Pointer, input_buffer unsafe.Pointer, input_length unsafe.Pointer, message unsafe.Pointer, is_complete bool) {
	_nw_framer_deliver_input(framer, input_buffer, input_length, message, is_complete)
	}


// Delivers an inbound message containing a specific number of next received bytes. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_deliver_input_no_copy(_:_:_:_:)
func nw_framer_deliver_input_no_copy(framer unsafe.Pointer, input_length unsafe.Pointer, message unsafe.Pointer, is_complete bool) bool {
	return _nw_framer_deliver_input_no_copy(framer, input_length, message, is_complete)
	}


// Indicates to a connection that your protocol has encountered an error, or has gracefully closed. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_mark_failed_with_error(_:_:)
func nw_framer_mark_failed_with_error(framer unsafe.Pointer, error_code int) {
	_nw_framer_mark_failed_with_error(framer, error_code)
	}


// Indicates to a connection that your protocol’s handshake is complete. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_mark_ready(_:)
func nw_framer_mark_ready(framer unsafe.Pointer) {
	_nw_framer_mark_ready(framer)
	}


// Accesses a custom value stored in a framer message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_access_value(_:_:_:)
func nw_framer_message_access_value(message unsafe.Pointer, key unsafe.Pointer, access_value bool) bool {
	return _nw_framer_message_access_value(message, key, access_value)
	}


// Accesses an NSObject value stored in a framer message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_copy_object_value(_:_:)
func nw_framer_message_copy_object_value(message unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_message_copy_object_value(message, key)
	}


// Initializes an empty message from within a framer implementation. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_create(_:)
func nw_framer_message_create(framer unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_message_create(framer)
	}


// Sets an NSObject value to be stored in a framer message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_set_object_value(_:_:_:)
func nw_framer_message_set_object_value(message unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_nw_framer_message_set_object_value(message, key, value)
	}


// Sets a value to be stored in a framer message, with a completion to call to disposed the stored value when the message is released. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_message_set_value(_:_:_:_:)
func nw_framer_message_set_value(message unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, dispose_value unsafe.Pointer) {
	_nw_framer_message_set_value(message, key, value, dispose_value)
	}


// nw_framer_options_copy_object_value is a Network function. [Full Topic]
//
// Added in macOS 12.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_options_copy_object_value(_:_:)
func nw_framer_options_copy_object_value(options unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_options_copy_object_value(options, key)
	}


// nw_framer_options_set_object_value is a Network function. [Full Topic]
//
// Added in macOS 12.3.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_options_set_object_value(_:_:_:)
func nw_framer_options_set_object_value(options unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_nw_framer_options_set_object_value(options, key, value)
	}


// Examines the content of input data while inside your input handler block. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_parse_input(_:_:_:_:_:)
func nw_framer_parse_input(framer unsafe.Pointer, minimum_incomplete_length unsafe.Pointer, maximum_length unsafe.Pointer, temp_buffer unsafe.Pointer, parse unsafe.Pointer) bool {
	return _nw_framer_parse_input(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	}


// Examines the content of output data while inside your output handler. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_parse_output(_:_:_:_:_:)
func nw_framer_parse_output(framer unsafe.Pointer, minimum_incomplete_length unsafe.Pointer, maximum_length unsafe.Pointer, temp_buffer unsafe.Pointer, parse unsafe.Pointer) bool {
	return _nw_framer_parse_output(framer, minimum_incomplete_length, maximum_length, temp_buffer, parse)
	}


// Indicates that your protocol no longer needs to handle input data. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_pass_through_input(_:)
func nw_framer_pass_through_input(framer unsafe.Pointer) {
	_nw_framer_pass_through_input(framer)
	}


// Indicates that your protocol no longer needs to handle output data. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_pass_through_output(_:)
func nw_framer_pass_through_output(framer unsafe.Pointer) {
	_nw_framer_pass_through_output(framer)
	}


// Dynamically adds another protocol that will run above your protocol after your protocol calls . [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_prepend_application_protocol(_:_:)
func nw_framer_prepend_application_protocol(framer unsafe.Pointer, protocol_options unsafe.Pointer) bool {
	return _nw_framer_prepend_application_protocol(framer, protocol_options)
	}


// Initializes an empty message for a custom framer definition. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_protocol_create_message(_:)
func nw_framer_protocol_create_message(definition unsafe.Pointer) unsafe.Pointer {
	return _nw_framer_protocol_create_message(definition)
	}


// Requests that the be called on your protocol at a specific time in the future. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_schedule_wakeup(_:_:)
func nw_framer_schedule_wakeup(framer unsafe.Pointer, milliseconds unsafe.Pointer) {
	_nw_framer_schedule_wakeup(framer, milliseconds)
	}


// Sets a block to handle the final cleanup of allocations made by your protocol instance. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_cleanup_handler(_:_:)
func nw_framer_set_cleanup_handler(framer unsafe.Pointer, cleanup_handler unsafe.Pointer) {
	_nw_framer_set_cleanup_handler(framer, cleanup_handler)
	}


// Sets a block to handle new inbound data. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_input_handler(_:_:)
func nw_framer_set_input_handler(framer unsafe.Pointer, input_handler unsafe.Pointer) {
	_nw_framer_set_input_handler(framer, input_handler)
	}


// Sets a block to handle new outbound messages. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_output_handler(_:_:)
func nw_framer_set_output_handler(framer unsafe.Pointer, output_handler unsafe.Pointer) {
	_nw_framer_set_output_handler(framer, output_handler)
	}


// Sets a block to handle when the connection is being closed. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_stop_handler(_:_:)
func nw_framer_set_stop_handler(framer unsafe.Pointer, stop_handler unsafe.Pointer) {
	_nw_framer_set_stop_handler(framer, stop_handler)
	}


// Sets a handler to receive scheduled wakeup events. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_set_wakeup_handler(_:_:)
func nw_framer_set_wakeup_handler(framer unsafe.Pointer, wakeup_handler unsafe.Pointer) {
	_nw_framer_set_wakeup_handler(framer, wakeup_handler)
	}


// Sends arbitrary output data in a buffer from your protocol to the next protocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output(_:_:_:)
func nw_framer_write_output(framer unsafe.Pointer, output_buffer unsafe.Pointer, output_length unsafe.Pointer) {
	_nw_framer_write_output(framer, output_buffer, output_length)
	}


// Sends arbitrary output data from your protocol to the next protocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output_data(_:_:)
func nw_framer_write_output_data(framer unsafe.Pointer, output_data unsafe.Pointer) {
	_nw_framer_write_output_data(framer, output_data)
	}


// Sends a specific number of bytes from a message while inside your output handler. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_framer_write_output_no_copy(_:_:)
func nw_framer_write_output_no_copy(framer unsafe.Pointer, output_length unsafe.Pointer) bool {
	return _nw_framer_write_output_no_copy(framer, output_length)
	}


// Adds a multicast address endpoint you specify to define an extra IP multicast group to join. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_add_endpoint(_:_:)
func nw_group_descriptor_add_endpoint(descriptor unsafe.Pointer, endpoint unsafe.Pointer) bool {
	return _nw_group_descriptor_add_endpoint(descriptor, endpoint)
	}


// Creates group descriptor you use to join an IP multicast group on a local network. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multicast(_:)
func nw_group_descriptor_create_multicast(multicast_group unsafe.Pointer) unsafe.Pointer {
	return _nw_group_descriptor_create_multicast(multicast_group)
	}


// nw_group_descriptor_create_multiplex is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_create_multiplex(_:)
func nw_group_descriptor_create_multiplex(remote_endpoint unsafe.Pointer) unsafe.Pointer {
	return _nw_group_descriptor_create_multiplex(remote_endpoint)
	}


// Sets a handler to list all endpoints added to the group descriptor. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_group_descriptor_enumerate_endpoints(_:_:)
func nw_group_descriptor_enumerate_endpoints(descriptor unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_group_descriptor_enumerate_endpoints(descriptor, enumerate_block)
	}


// Accesses the system interface index associated with the interface. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_index(_:)
func nw_interface_get_index(interface_ unsafe.Pointer) unsafe.Pointer {
	return _nw_interface_get_index(interface_)
	}


// Accesses the name of the interface. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_name(_:)
func nw_interface_get_name(interface_ unsafe.Pointer) unsafe.Pointer {
	return _nw_interface_get_name(interface_)
	}


// Accesses the type of the interface, such as Wi-Fi or Loopback. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_interface_get_type(_:)
func nw_interface_get_type(interface_ unsafe.Pointer) unsafe.Pointer {
	return _nw_interface_get_type(interface_)
	}


// Initializes an IP packet configuration with default settings. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_create_metadata()
func nw_ip_create_metadata() unsafe.Pointer {
	return _nw_ip_create_metadata()
	}


// Checks the Explicit Congestion Notification flag value received on an IP packet. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_ecn_flag(_:)
func nw_ip_metadata_get_ecn_flag(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ip_metadata_get_ecn_flag(metadata)
	}


// Access the time at which a packet was received, in nanoseconds, based on . [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_receive_time(_:)
func nw_ip_metadata_get_receive_time(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ip_metadata_get_receive_time(metadata)
	}


// Accesses a specific service class to mark on an IP packet. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_get_service_class(_:)
func nw_ip_metadata_get_service_class(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ip_metadata_get_service_class(metadata)
	}


// Sets a specific Explicit Congestion Notification flag value to set on an IP packet. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_ecn_flag(_:_:)
func nw_ip_metadata_set_ecn_flag(metadata unsafe.Pointer, ecn_flag unsafe.Pointer) {
	_nw_ip_metadata_set_ecn_flag(metadata, ecn_flag)
	}


// Sets a specific service class to mark on an IP packet. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_metadata_set_service_class(_:_:)
func nw_ip_metadata_set_service_class(metadata unsafe.Pointer, service_class unsafe.Pointer) {
	_nw_ip_metadata_set_service_class(metadata, service_class)
	}


// Configures a connection to deliver receive timestamps for IP packets. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_calculate_receive_time(_:_:)
func nw_ip_options_set_calculate_receive_time(options unsafe.Pointer, calculate_receive_time bool) {
	_nw_ip_options_set_calculate_receive_time(options, calculate_receive_time)
	}


// Configures a connection to disable fragmentation on outbound packets. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_fragmentation(_:_:)
func nw_ip_options_set_disable_fragmentation(options unsafe.Pointer, disable_fragmentation bool) {
	_nw_ip_options_set_disable_fragmentation(options, disable_fragmentation)
	}


// nw_ip_options_set_disable_multicast_loopback is a Network function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_disable_multicast_loopback(_:_:)
func nw_ip_options_set_disable_multicast_loopback(options unsafe.Pointer, disable_multicast_loopback bool) {
	_nw_ip_options_set_disable_multicast_loopback(options, disable_multicast_loopback)
	}


// Configures the default hop limit for packets generated by a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_hop_limit(_:_:)
func nw_ip_options_set_hop_limit(options unsafe.Pointer, hop_limit unsafe.Pointer) {
	_nw_ip_options_set_hop_limit(options, hop_limit)
	}


// Configures a connection to prefer certain types of local addresses, such as temporary or stable. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_local_address_preference(_:_:)
func nw_ip_options_set_local_address_preference(options unsafe.Pointer, preference unsafe.Pointer) {
	_nw_ip_options_set_local_address_preference(options, preference)
	}


// Configures a connection to use the minimum MTU value, which is 1280 bytes for IPv6. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_use_minimum_mtu(_:_:)
func nw_ip_options_set_use_minimum_mtu(options unsafe.Pointer, use_minimum_mtu bool) {
	_nw_ip_options_set_use_minimum_mtu(options, use_minimum_mtu)
	}


// Sets a required IP version to disable all other versions for a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ip_options_set_version(_:_:)
func nw_ip_options_set_version(options unsafe.Pointer, version unsafe.Pointer) {
	_nw_ip_options_set_version(options, version)
	}


// Stops listening for inbound connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_cancel(_:)
func nw_listener_cancel(listener unsafe.Pointer) {
	_nw_listener_cancel(listener)
	}


// Initializes a network listener, which will select a random port. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create(_:)
func nw_listener_create(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_create(parameters)
	}


// Initializes a network listener to receive new streams on a multiplexed connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_connection(_:_:)
func nw_listener_create_with_connection(connection unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_create_with_connection(connection, parameters)
	}


// nw_listener_create_with_launchd_key is a Network function. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_launchd_key(_:_:)
func nw_listener_create_with_launchd_key(parameters unsafe.Pointer, launchd_key unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_create_with_launchd_key(parameters, launchd_key)
	}


// Initializes a network listener with a specified local port. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_create_with_port(_:_:)
func nw_listener_create_with_port(port unsafe.Pointer, parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_create_with_port(port, parameters)
	}


// Checks the remaining number of inbound connections to deliver before rejecting connections. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_get_new_connection_limit(_:)
func nw_listener_get_new_connection_limit(listener unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_get_new_connection_limit(listener)
	}


// The port on which the listener can accept connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_get_port(_:)
func nw_listener_get_port(listener unsafe.Pointer) unsafe.Pointer {
	return _nw_listener_get_port(listener)
	}


// Sets a Bonjour service that advertises the listener on the local network. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_advertise_descriptor(_:_:)
func nw_listener_set_advertise_descriptor(listener unsafe.Pointer, advertise_descriptor unsafe.Pointer) {
	_nw_listener_set_advertise_descriptor(listener, advertise_descriptor)
	}


// Sets a handler that receives updates for the service endpoint being advertised. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_advertised_endpoint_changed_handler(_:_:)
func nw_listener_set_advertised_endpoint_changed_handler(listener unsafe.Pointer, handler unsafe.Pointer) {
	_nw_listener_set_advertised_endpoint_changed_handler(listener, handler)
	}


// nw_listener_set_new_connection_group_handler is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_group_handler(_:_:)
func nw_listener_set_new_connection_group_handler(listener unsafe.Pointer, handler unsafe.Pointer) {
	_nw_listener_set_new_connection_group_handler(listener, handler)
	}


// Sets a handler that receives inbound connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_handler(_:_:)
func nw_listener_set_new_connection_handler(listener unsafe.Pointer, handler unsafe.Pointer) {
	_nw_listener_set_new_connection_handler(listener, handler)
	}


// Resets the number of inbound connections to deliver before rejecting connections. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_new_connection_limit(_:_:)
func nw_listener_set_new_connection_limit(listener unsafe.Pointer, new_connection_limit unsafe.Pointer) {
	_nw_listener_set_new_connection_limit(listener, new_connection_limit)
	}


// Sets the queue on which all listener events are delivered. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_queue(_:_:)
func nw_listener_set_queue(listener unsafe.Pointer, queue unsafe.Pointer) {
	_nw_listener_set_queue(listener, queue)
	}


// Sets a handler to receive listener state updates. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_set_state_changed_handler(_:_:)
func nw_listener_set_state_changed_handler(listener unsafe.Pointer, handler unsafe.Pointer) {
	_nw_listener_set_state_changed_handler(listener, handler)
	}


// Registers for listening for inbound connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_listener_start(_:)
func nw_listener_start(listener unsafe.Pointer) {
	_nw_listener_start(listener)
	}


// Checks a Boolean that indicates whether a connection group should reject unicast traffic. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_get_disable_unicast_traffic(_:)
func nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor unsafe.Pointer) bool {
	return _nw_multicast_group_descriptor_get_disable_unicast_traffic(multicast_descriptor)
	}


// Sets a Boolean that indicates whether a connection group should reject unicast traffic. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_disable_unicast_traffic(_:_:)
func nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor unsafe.Pointer, disable_unicast_traffic bool) {
	_nw_multicast_group_descriptor_set_disable_unicast_traffic(multicast_descriptor, disable_unicast_traffic)
	}


// Sets an optional address endpoint used to filter received multicast packets. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_multicast_group_descriptor_set_specific_source(_:_:)
func nw_multicast_group_descriptor_set_specific_source(multicast_descriptor unsafe.Pointer, source unsafe.Pointer) {
	_nw_multicast_group_descriptor_set_specific_source(multicast_descriptor, source)
	}


// Removes all prohibited interface types. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interface_types(_:)
func nw_parameters_clear_prohibited_interface_types(parameters unsafe.Pointer) {
	_nw_parameters_clear_prohibited_interface_types(parameters)
	}


// Removes all prohibited interface types. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_clear_prohibited_interfaces(_:)
func nw_parameters_clear_prohibited_interfaces(parameters unsafe.Pointer) {
	_nw_parameters_clear_prohibited_interfaces(parameters)
	}


// Peforms a deep copy of a parameters object. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy(_:)
func nw_parameters_copy(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_copy(parameters)
	}


// Accesses the protocol stack used by connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy_default_protocol_stack(_:)
func nw_parameters_copy_default_protocol_stack(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_copy_default_protocol_stack(parameters)
	}


// Accesses the local IP address and port used for connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy_local_endpoint(_:)
func nw_parameters_copy_local_endpoint(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_copy_local_endpoint(parameters)
	}


// Accesses the interface required on connections, listeners, and browsers. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_copy_required_interface(_:)
func nw_parameters_copy_required_interface(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_copy_required_interface(parameters)
	}


// Initializes parameters for connections, listeners, and browsers with no protocols specified. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create()
func nw_parameters_create() unsafe.Pointer {
	return _nw_parameters_create()
	}


// nw_parameters_create_application_service is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_application_service()
func nw_parameters_create_application_service() unsafe.Pointer {
	return _nw_parameters_create_application_service()
	}


// Initializes parameters for connections and listeners using a custom IP protocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_custom_ip(_:_:)
func nw_parameters_create_custom_ip(custom_ip_protocol_number unsafe.Pointer, configure_ip unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_create_custom_ip(custom_ip_protocol_number, configure_ip)
	}


// Initializes parameters for QUIC connections and listeners. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_quic(_:)
func nw_parameters_create_quic(configure_quic unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_create_quic(configure_quic)
	}


// Initializes parameters for TLS or TCP connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_tcp(_:_:)
func nw_parameters_create_secure_tcp(configure_tls unsafe.Pointer, configure_tcp unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_create_secure_tcp(configure_tls, configure_tcp)
	}


// Initializes parameters for DTLS or UDP connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_create_secure_udp(_:_:)
func nw_parameters_create_secure_udp(configure_dtls unsafe.Pointer, configure_udp unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_create_secure_udp(configure_dtls, configure_udp)
	}


// nw_parameters_get_allow_ultra_constrained is a Network function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_allow_ultra_constrained(_:)
func nw_parameters_get_allow_ultra_constrained(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_allow_ultra_constrained(parameters)
	}


// Gets a flag that indicates whether the network request originates from the developer or the user. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_attribution(_:)
func nw_parameters_get_attribution(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_get_attribution(parameters)
	}


// Checks the behavior for how expired DNS answers should be used. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_expired_dns_behavior(_:)
func nw_parameters_get_expired_dns_behavior(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_get_expired_dns_behavior(parameters)
	}


// Checks if sending application data with protocol handshakes is enabled. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_fast_open_enabled(_:)
func nw_parameters_get_fast_open_enabled(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_fast_open_enabled(parameters)
	}


// Checks whether a connection is allowed to use peer-to-peer link technologies. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_include_peer_to_peer(_:)
func nw_parameters_get_include_peer_to_peer(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_include_peer_to_peer(parameters)
	}


// Checks if a listener is restricted to accepting connections from the local link. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_local_only(_:)
func nw_parameters_get_local_only(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_local_only(parameters)
	}


// Checks if multipath is enabled on a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_multipath_service(_:)
func nw_parameters_get_multipath_service(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_get_multipath_service(parameters)
	}


// Checks if proxies are ignored by default. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prefer_no_proxy(_:)
func nw_parameters_get_prefer_no_proxy(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_prefer_no_proxy(parameters)
	}


// Checks if connections, listeners, and browsers are prevented from using network paths marked as constrained by Low Data Mode. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_constrained(_:)
func nw_parameters_get_prohibit_constrained(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_prohibit_constrained(parameters)
	}


// Checks if connections, listeners, and browsers are prevented from using network paths marked as expensive. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_prohibit_expensive(_:)
func nw_parameters_get_prohibit_expensive(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_prohibit_expensive(parameters)
	}


// Accesses the interface type required on connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_required_interface_type(_:)
func nw_parameters_get_required_interface_type(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_get_required_interface_type(parameters)
	}


// Checks whether a connection allows reusing local addresses and ports. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_reuse_local_address(_:)
func nw_parameters_get_reuse_local_address(parameters unsafe.Pointer) bool {
	return _nw_parameters_get_reuse_local_address(parameters)
	}


// Checks the level of service quality used for connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_get_service_class(_:)
func nw_parameters_get_service_class(parameters unsafe.Pointer) unsafe.Pointer {
	return _nw_parameters_get_service_class(parameters)
	}


// Examines the list of prohibited interface types. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interface_types(_:_:)
func nw_parameters_iterate_prohibited_interface_types(parameters unsafe.Pointer, iterate_block unsafe.Pointer) {
	_nw_parameters_iterate_prohibited_interface_types(parameters, iterate_block)
	}


// Examines the list of prohibited interfaces. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_iterate_prohibited_interfaces(_:_:)
func nw_parameters_iterate_prohibited_interfaces(parameters unsafe.Pointer, iterate_block unsafe.Pointer) {
	_nw_parameters_iterate_prohibited_interfaces(parameters, iterate_block)
	}


// Prevents connections and listeners from using a specific interface. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface(_:_:)
func nw_parameters_prohibit_interface(parameters unsafe.Pointer, interface_ unsafe.Pointer) {
	_nw_parameters_prohibit_interface(parameters, interface_)
	}


// Prevents connections, listeners, and browsers from using a specific interface type. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_prohibit_interface_type(_:_:)
func nw_parameters_prohibit_interface_type(parameters unsafe.Pointer, interface_type unsafe.Pointer) {
	_nw_parameters_prohibit_interface_type(parameters, interface_type)
	}


// Sets a specific interface to require on connections, listeners, and browsers. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_require_interface(_:_:)
func nw_parameters_require_interface(parameters unsafe.Pointer, interface_ unsafe.Pointer) {
	_nw_parameters_require_interface(parameters, interface_)
	}


// Checks whether a connection requires DNSSEC validation when resolving endpoints. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_requires_dnssec_validation(_:)
func nw_parameters_requires_dnssec_validation(parameters unsafe.Pointer) bool {
	return _nw_parameters_requires_dnssec_validation(parameters)
	}


// nw_parameters_set_allow_ultra_constrained is a Network function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_allow_ultra_constrained(_:_:)
func nw_parameters_set_allow_ultra_constrained(parameters unsafe.Pointer, allow_ultra_constrained bool) {
	_nw_parameters_set_allow_ultra_constrained(parameters, allow_ultra_constrained)
	}


// Sets a flag that indicates whether the network request originates from the developer or the user. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_attribution(_:_:)
func nw_parameters_set_attribution(parameters unsafe.Pointer, attribution unsafe.Pointer) {
	_nw_parameters_set_attribution(parameters, attribution)
	}


// Sets the behavior for how expired DNS answers should be used. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_expired_dns_behavior(_:_:)
func nw_parameters_set_expired_dns_behavior(parameters unsafe.Pointer, expired_dns_behavior unsafe.Pointer) {
	_nw_parameters_set_expired_dns_behavior(parameters, expired_dns_behavior)
	}


// Enables sending application data with protocol handshakes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_fast_open_enabled(_:_:)
func nw_parameters_set_fast_open_enabled(parameters unsafe.Pointer, fast_open_enabled bool) {
	_nw_parameters_set_fast_open_enabled(parameters, fast_open_enabled)
	}


// Enables peer-to-peer link technologies for connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_include_peer_to_peer(_:_:)
func nw_parameters_set_include_peer_to_peer(parameters unsafe.Pointer, include_peer_to_peer bool) {
	_nw_parameters_set_include_peer_to_peer(parameters, include_peer_to_peer)
	}


// Sets a specific local IP address and port to use for connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_local_endpoint(_:_:)
func nw_parameters_set_local_endpoint(parameters unsafe.Pointer, local_endpoint unsafe.Pointer) {
	_nw_parameters_set_local_endpoint(parameters, local_endpoint)
	}


// Restricts listeners to only accepting connections from the local link. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_local_only(_:_:)
func nw_parameters_set_local_only(parameters unsafe.Pointer, local_only bool) {
	_nw_parameters_set_local_only(parameters, local_only)
	}


// Enables multipath protocols to allow connections to use multiple interfaces. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_multipath_service(_:_:)
func nw_parameters_set_multipath_service(parameters unsafe.Pointer, multipath_service unsafe.Pointer) {
	_nw_parameters_set_multipath_service(parameters, multipath_service)
	}


// Sets a Boolean that indicates that connections should ignore proxies when they are enabled on the system. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prefer_no_proxy(_:_:)
func nw_parameters_set_prefer_no_proxy(parameters unsafe.Pointer, prefer_no_proxy bool) {
	_nw_parameters_set_prefer_no_proxy(parameters, prefer_no_proxy)
	}


// Associates a privacy context with any connections or listeners that use the parameters. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_privacy_context(_:_:)
func nw_parameters_set_privacy_context(parameters unsafe.Pointer, privacy_context unsafe.Pointer) {
	_nw_parameters_set_privacy_context(parameters, privacy_context)
	}


// Prevents connections, listeners, and browsers from using network paths marked as constrained by Low Data Mode. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_constrained(_:_:)
func nw_parameters_set_prohibit_constrained(parameters unsafe.Pointer, prohibit_constrained bool) {
	_nw_parameters_set_prohibit_constrained(parameters, prohibit_constrained)
	}


// Prevents connections, listeners, and browsers from using network paths marked as expensive. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_prohibit_expensive(_:_:)
func nw_parameters_set_prohibit_expensive(parameters unsafe.Pointer, prohibit_expensive bool) {
	_nw_parameters_set_prohibit_expensive(parameters, prohibit_expensive)
	}


// Sets an interface type to require on connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_required_interface_type(_:_:)
func nw_parameters_set_required_interface_type(parameters unsafe.Pointer, interface_type unsafe.Pointer) {
	_nw_parameters_set_required_interface_type(parameters, interface_type)
	}


// Determines whether a connection requires DNSSEC validation when resolving endpoints. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_requires_dnssec_validation(_:_:)
func nw_parameters_set_requires_dnssec_validation(parameters unsafe.Pointer, requires_dnssec_validation bool) {
	_nw_parameters_set_requires_dnssec_validation(parameters, requires_dnssec_validation)
	}


// Allows reusing local addresses and ports across connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_reuse_local_address(_:_:)
func nw_parameters_set_reuse_local_address(parameters unsafe.Pointer, reuse_local_address bool) {
	_nw_parameters_set_reuse_local_address(parameters, reuse_local_address)
	}


// Sets a level of service quality to use for connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_parameters_set_service_class(_:_:)
func nw_parameters_set_service_class(parameters unsafe.Pointer, service_class unsafe.Pointer) {
	_nw_parameters_set_service_class(parameters, service_class)
	}


// Accesses the local endpoint in use by a connection’s network path. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_copy_effective_local_endpoint(_:)
func nw_path_copy_effective_local_endpoint(path unsafe.Pointer) unsafe.Pointer {
	return _nw_path_copy_effective_local_endpoint(path)
	}


// Accesses the remote endpoint in use by a connection’s network path. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_copy_effective_remote_endpoint(_:)
func nw_path_copy_effective_remote_endpoint(path unsafe.Pointer) unsafe.Pointer {
	return _nw_path_copy_effective_remote_endpoint(path)
	}


// Enumerates the list of gateways configured on the interfaces available to a path. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_enumerate_gateways(_:_:)
func nw_path_enumerate_gateways(path unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_path_enumerate_gateways(path, enumerate_block)
	}


// Enumerates the list of all interfaces available to the path, in order of preference. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_enumerate_interfaces(_:_:)
func nw_path_enumerate_interfaces(path unsafe.Pointer, enumerate_block unsafe.Pointer) {
	_nw_path_enumerate_interfaces(path, enumerate_block)
	}


// nw_path_get_link_quality is a Network function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_link_quality(_:)
func nw_path_get_link_quality(path unsafe.Pointer) unsafe.Pointer {
	return _nw_path_get_link_quality(path)
	}


// Checks whether a path can be used by connections. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_status(_:)
func nw_path_get_status(path unsafe.Pointer) unsafe.Pointer {
	return _nw_path_get_status(path)
	}


// nw_path_get_unsatisfied_reason is a Network function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_get_unsatisfied_reason(_:)
func nw_path_get_unsatisfied_reason(path unsafe.Pointer) unsafe.Pointer {
	return _nw_path_get_unsatisfied_reason(path)
	}


// Checks whether the path has a DNS server configured. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_dns(_:)
func nw_path_has_dns(path unsafe.Pointer) bool {
	return _nw_path_has_dns(path)
	}


// Checks whether the path can route IPv4 traffic. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_ipv4(_:)
func nw_path_has_ipv4(path unsafe.Pointer) bool {
	return _nw_path_has_ipv4(path)
	}


// Checks whether the path can route IPv6 traffic. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_has_ipv6(_:)
func nw_path_has_ipv6(path unsafe.Pointer) bool {
	return _nw_path_has_ipv6(path)
	}


// Checks whether the path uses an interface in Low Data Mode. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_constrained(_:)
func nw_path_is_constrained(path unsafe.Pointer) bool {
	return _nw_path_is_constrained(path)
	}


// Compares if two paths are identical. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_equal(_:_:)
func nw_path_is_equal(path unsafe.Pointer, other_path unsafe.Pointer) bool {
	return _nw_path_is_equal(path, other_path)
	}


// Checks whether the path uses an interface that is considered expensive, such as Cellular or a Personal Hotspot. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_expensive(_:)
func nw_path_is_expensive(path unsafe.Pointer) bool {
	return _nw_path_is_expensive(path)
	}


// nw_path_is_ultra_constrained is a Network function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_is_ultra_constrained(_:)
func nw_path_is_ultra_constrained(path unsafe.Pointer) bool {
	return _nw_path_is_ultra_constrained(path)
	}


// Stops receiving network path updates. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_cancel(_:)
func nw_path_monitor_cancel(monitor unsafe.Pointer) {
	_nw_path_monitor_cancel(monitor)
	}


// Initializes a path monitor to observe all available interface types. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create()
func nw_path_monitor_create() unsafe.Pointer {
	return _nw_path_monitor_create()
	}


// nw_path_monitor_create_for_ethernet_channel is a Network function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create_for_ethernet_channel()
func nw_path_monitor_create_for_ethernet_channel() unsafe.Pointer {
	return _nw_path_monitor_create_for_ethernet_channel()
	}


// Initializes a path monitor to observe a specific interface type. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_create_with_type(_:)
func nw_path_monitor_create_with_type(required_interface_type unsafe.Pointer) unsafe.Pointer {
	return _nw_path_monitor_create_with_type(required_interface_type)
	}


// Prohibit a path monitor from using a specific interface type. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_prohibit_interface_type(_:_:)
func nw_path_monitor_prohibit_interface_type(monitor unsafe.Pointer, interface_type unsafe.Pointer) {
	_nw_path_monitor_prohibit_interface_type(monitor, interface_type)
	}


// Sets a handler to determine when a monitor is fully cancelled and will no longer deliver events. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_cancel_handler(_:_:)
func nw_path_monitor_set_cancel_handler(monitor unsafe.Pointer, cancel_handler unsafe.Pointer) {
	_nw_path_monitor_set_cancel_handler(monitor, cancel_handler)
	}


// Sets a queue on which to deliver path events. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_queue(_:_:)
func nw_path_monitor_set_queue(monitor unsafe.Pointer, queue unsafe.Pointer) {
	_nw_path_monitor_set_queue(monitor, queue)
	}


// Sets a handler to receive network path updates. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_set_update_handler(_:_:)
func nw_path_monitor_set_update_handler(monitor unsafe.Pointer, update_handler unsafe.Pointer) {
	_nw_path_monitor_set_update_handler(monitor, update_handler)
	}


// Starts monitoring path changes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_monitor_start(_:)
func nw_path_monitor_start(monitor unsafe.Pointer) {
	_nw_path_monitor_start(monitor)
	}


// Checks if connections using the path may send traffic over a specific interface type. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_path_uses_interface_type(_:_:)
func nw_path_uses_interface_type(path unsafe.Pointer, interface_type unsafe.Pointer) bool {
	return _nw_path_uses_interface_type(path, interface_type)
	}


// Accesses the system definition of the Internet Protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_ip_definition()
func nw_protocol_copy_ip_definition() unsafe.Pointer {
	return _nw_protocol_copy_ip_definition()
	}


// Accesses the system definition of the QUIC transport protocol. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_quic_definition()
func nw_protocol_copy_quic_definition() unsafe.Pointer {
	return _nw_protocol_copy_quic_definition()
	}


// Accesses the system definition of the Transport Control Protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_tcp_definition()
func nw_protocol_copy_tcp_definition() unsafe.Pointer {
	return _nw_protocol_copy_tcp_definition()
	}


// Accesses the system definition of the Transport Layer Security protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_tls_definition()
func nw_protocol_copy_tls_definition() unsafe.Pointer {
	return _nw_protocol_copy_tls_definition()
	}


// Accesses the system definition of the User Datagram Protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_udp_definition()
func nw_protocol_copy_udp_definition() unsafe.Pointer {
	return _nw_protocol_copy_udp_definition()
	}


// Accesses the system definition of the WebSocket protocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_copy_ws_definition()
func nw_protocol_copy_ws_definition() unsafe.Pointer {
	return _nw_protocol_copy_ws_definition()
	}


// Accesses the protocol definition associated with the metadata object. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_copy_definition(_:)
func nw_protocol_metadata_copy_definition(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_protocol_metadata_copy_definition(metadata)
	}


// Checks if a metadata object represents a custom framer protocol message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_framer_message(_:)
func nw_protocol_metadata_is_framer_message(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_framer_message(metadata)
	}


// Checks whether a metadata object represents an IP packet. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ip(_:)
func nw_protocol_metadata_is_ip(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_ip(metadata)
	}


// Checks whether a metadata object contains QUIC connection state. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_quic(_:)
func nw_protocol_metadata_is_quic(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_quic(metadata)
	}


// Checks whether a metadata object contains TCP connection state. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tcp(_:)
func nw_protocol_metadata_is_tcp(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_tcp(metadata)
	}


// Checks whether a metadata object contains TLS connection state. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_tls(_:)
func nw_protocol_metadata_is_tls(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_tls(metadata)
	}


// Checks whether a metadata object represents a UDP datagram. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_udp(_:)
func nw_protocol_metadata_is_udp(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_udp(metadata)
	}


// Checks whether a metadata object represents a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_metadata_is_ws(_:)
func nw_protocol_metadata_is_ws(metadata unsafe.Pointer) bool {
	return _nw_protocol_metadata_is_ws(metadata)
	}


// Checks whether an options object uses the QUIC protocol. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_options_is_quic(_:)
func nw_protocol_options_is_quic(options unsafe.Pointer) bool {
	return _nw_protocol_options_is_quic(options)
	}


// Removes all application protocols from the protocol stack. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_clear_application_protocols(_:)
func nw_protocol_stack_clear_application_protocols(stack unsafe.Pointer) {
	_nw_protocol_stack_clear_application_protocols(stack)
	}


// Accesses the protocol stack’s Internet Protocol options. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_internet_protocol(_:)
func nw_protocol_stack_copy_internet_protocol(stack unsafe.Pointer) unsafe.Pointer {
	return _nw_protocol_stack_copy_internet_protocol(stack)
	}


// Accesses the options for the protocol stack’s transport protocol. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_copy_transport_protocol(_:)
func nw_protocol_stack_copy_transport_protocol(stack unsafe.Pointer) unsafe.Pointer {
	return _nw_protocol_stack_copy_transport_protocol(stack)
	}


// Iterates through the array of application protocol options that will be used by connections and listeners. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_iterate_application_protocols(_:_:)
func nw_protocol_stack_iterate_application_protocols(stack unsafe.Pointer, iterate_block unsafe.Pointer) {
	_nw_protocol_stack_iterate_application_protocols(stack, iterate_block)
	}


// Adds a protocol onto the top of the protocol stack. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_prepend_application_protocol(_:_:)
func nw_protocol_stack_prepend_application_protocol(stack unsafe.Pointer, protocol_ unsafe.Pointer) {
	_nw_protocol_stack_prepend_application_protocol(stack, protocol_)
	}


// Replaces the protocol stack’s transport protocol with a new set of options. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_protocol_stack_set_transport_protocol(_:_:)
func nw_protocol_stack_set_transport_protocol(stack unsafe.Pointer, protocol_ unsafe.Pointer) {
	_nw_protocol_stack_set_transport_protocol(stack, protocol_)
	}


// nw_proxy_config_add_excluded_domain is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_add_excluded_domain(_:_:)
func nw_proxy_config_add_excluded_domain(config unsafe.Pointer, excluded_domain unsafe.Pointer) {
	_nw_proxy_config_add_excluded_domain(config, excluded_domain)
	}


// nw_proxy_config_add_match_domain is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_add_match_domain(_:_:)
func nw_proxy_config_add_match_domain(config unsafe.Pointer, match_domain unsafe.Pointer) {
	_nw_proxy_config_add_match_domain(config, match_domain)
	}


// nw_proxy_config_clear_excluded_domains is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_excluded_domains(_:)
func nw_proxy_config_clear_excluded_domains(config unsafe.Pointer) {
	_nw_proxy_config_clear_excluded_domains(config)
	}


// nw_proxy_config_clear_match_domains is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_clear_match_domains(_:)
func nw_proxy_config_clear_match_domains(config unsafe.Pointer) {
	_nw_proxy_config_clear_match_domains(config)
	}


// nw_proxy_config_enumerate_excluded_domains is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_excluded_domains(_:_:)
func nw_proxy_config_enumerate_excluded_domains(config unsafe.Pointer, enumerator unsafe.Pointer) {
	_nw_proxy_config_enumerate_excluded_domains(config, enumerator)
	}


// nw_proxy_config_enumerate_match_domains is a Network function. [Full Topic]
//
// Added in macOS 14.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_proxy_config_enumerate_match_domains(_:_:)
func nw_proxy_config_enumerate_match_domains(config unsafe.Pointer, enumerator unsafe.Pointer) {
	_nw_proxy_config_enumerate_match_domains(config, enumerator)
	}


// Adds a supported Application-Layer Protocol Negotiation value. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_add_tls_application_protocol(_:_:)
func nw_quic_add_tls_application_protocol(options unsafe.Pointer, application_protocol unsafe.Pointer) {
	_nw_quic_add_tls_application_protocol(options, application_protocol)
	}


// Accesses the result of the QUIC handshake. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_metadata(_:)
func nw_quic_copy_sec_protocol_metadata(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_copy_sec_protocol_metadata(metadata)
	}


// Accesses the handshake security options QUIC will use. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_copy_sec_protocol_options(_:)
func nw_quic_copy_sec_protocol_options(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_copy_sec_protocol_options(options)
	}


// Initializes a default set of QUIC connection options. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_create_options()
func nw_quic_create_options() unsafe.Pointer {
	return _nw_quic_create_options()
	}


// Accesses the QUIC application error code received from the peer. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_application_error(_:)
func nw_quic_get_application_error(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_application_error(metadata)
	}


// Accesses the QUIC application error reason received from the peer. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_application_error_reason(_:)
func nw_quic_get_application_error_reason(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_application_error_reason(metadata)
	}


// Accesses the idle timeout for the QUIC connection, in milliseconds. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_idle_timeout(_:)
func nw_quic_get_idle_timeout(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_idle_timeout(options)
	}


// Accesses a QUIC connection’s initial maximum data transport parameter. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_data(_:)
func nw_quic_get_initial_max_data(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_data(options)
	}


// Accesses a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_local(_:)
func nw_quic_get_initial_max_stream_data_bidirectional_local(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_stream_data_bidirectional_local(options)
	}


// Accesses a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_bidirectional_remote(_:)
func nw_quic_get_initial_max_stream_data_bidirectional_remote(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_stream_data_bidirectional_remote(options)
	}


// Accesses a QUIC connection’s initial maximum stream data limit for unidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_stream_data_unidirectional(_:)
func nw_quic_get_initial_max_stream_data_unidirectional(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_stream_data_unidirectional(options)
	}


// Accesses a QUIC connection’s initial maximum number of bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_bidirectional(_:)
func nw_quic_get_initial_max_streams_bidirectional(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_streams_bidirectional(options)
	}


// Accesses a QUIC connection’s initial maximum number of unidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_initial_max_streams_unidirectional(_:)
func nw_quic_get_initial_max_streams_unidirectional(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_initial_max_streams_unidirectional(options)
	}


// Accesses the keepalive interval for the QUIC connection, in seconds. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_keepalive_interval(_:)
func nw_quic_get_keepalive_interval(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_keepalive_interval(metadata)
	}


// Accesses the maximum number of bidirectional streams that the peer can create on a QUIC connection. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_bidirectional(_:)
func nw_quic_get_local_max_streams_bidirectional(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_local_max_streams_bidirectional(metadata)
	}


// Accesses the maximum number of unidirectional streams that the peer can create on a QUIC connection. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_local_max_streams_unidirectional(_:)
func nw_quic_get_local_max_streams_unidirectional(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_local_max_streams_unidirectional(metadata)
	}


// Accesses a QUIC connection’s maximum DATAGRAM frame size. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_max_datagram_frame_size(_:)
func nw_quic_get_max_datagram_frame_size(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_max_datagram_frame_size(options)
	}


// Accesses the maximum length of a QUIC packet that can be received on a connection, in bytes. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_max_udp_payload_size(_:)
func nw_quic_get_max_udp_payload_size(options unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_max_udp_payload_size(options)
	}


// Accesses the idle timeout value from the peer’s transport parameters, in milliseconds. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_idle_timeout(_:)
func nw_quic_get_remote_idle_timeout(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_remote_idle_timeout(metadata)
	}


// Accesses the maximum number of bidirectional streams advertised by peer that the connection is allowed to create. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_bidirectional(_:)
func nw_quic_get_remote_max_streams_bidirectional(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_remote_max_streams_bidirectional(metadata)
	}


// Accesses the maximum number of unidirectional streams advertised by peer that the connection is allowed to create. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_remote_max_streams_unidirectional(_:)
func nw_quic_get_remote_max_streams_unidirectional(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_remote_max_streams_unidirectional(metadata)
	}


// Accesses the QUIC application error code received from the peer for the stream. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_application_error(_:)
func nw_quic_get_stream_application_error(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_stream_application_error(metadata)
	}


// Accesses the QUIC stream identifier. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_id(_:)
func nw_quic_get_stream_id(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_stream_id(metadata)
	}


// Checks if a QUIC stream is a datagram flow, instead of a byte stream. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_datagram(_:)
func nw_quic_get_stream_is_datagram(options unsafe.Pointer) bool {
	return _nw_quic_get_stream_is_datagram(options)
	}


// Checks if a QUIC stream is unidirectional, instead of bidirectional. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_is_unidirectional(_:)
func nw_quic_get_stream_is_unidirectional(options unsafe.Pointer) bool {
	return _nw_quic_get_stream_is_unidirectional(options)
	}


// Accesses the stream type of the QUIC stream. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_type(_:)
func nw_quic_get_stream_type(stream_metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_stream_type(stream_metadata)
	}


// Accesses the maximum usable size of a datagram frame on a QUIC datagram flow. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_get_stream_usable_datagram_frame_size(_:)
func nw_quic_get_stream_usable_datagram_frame_size(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_quic_get_stream_usable_datagram_frame_size(metadata)
	}


// Sets the QUIC application error code to send for the connection. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_application_error(_:_:_:)
func nw_quic_set_application_error(metadata unsafe.Pointer, application_error unsafe.Pointer, reason unsafe.Pointer) {
	_nw_quic_set_application_error(metadata, application_error, reason)
	}


// Sets the idle timeout for the QUIC connection, in milliseconds. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_idle_timeout(_:_:)
func nw_quic_set_idle_timeout(options unsafe.Pointer, idle_timeout unsafe.Pointer) {
	_nw_quic_set_idle_timeout(options, idle_timeout)
	}


// Sets a QUIC connection’s initial maximum data transport parameter. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_data(_:_:)
func nw_quic_set_initial_max_data(options unsafe.Pointer, initial_max_data unsafe.Pointer) {
	_nw_quic_set_initial_max_data(options, initial_max_data)
	}


// Sets a QUIC connection’s initial maximum stream data limit for locally-initiated bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_local(_:_:)
func nw_quic_set_initial_max_stream_data_bidirectional_local(options unsafe.Pointer, initial_max_stream_data_bidirectional_local unsafe.Pointer) {
	_nw_quic_set_initial_max_stream_data_bidirectional_local(options, initial_max_stream_data_bidirectional_local)
	}


// Sets a QUIC connection’s initial maximum stream data limit for remote-initiated bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_bidirectional_remote(_:_:)
func nw_quic_set_initial_max_stream_data_bidirectional_remote(options unsafe.Pointer, initial_max_stream_data_bidirectional_remote unsafe.Pointer) {
	_nw_quic_set_initial_max_stream_data_bidirectional_remote(options, initial_max_stream_data_bidirectional_remote)
	}


// Sets a QUIC connection’s initial maximum stream data limit for unidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_stream_data_unidirectional(_:_:)
func nw_quic_set_initial_max_stream_data_unidirectional(options unsafe.Pointer, initial_max_stream_data_unidirectional unsafe.Pointer) {
	_nw_quic_set_initial_max_stream_data_unidirectional(options, initial_max_stream_data_unidirectional)
	}


// Sets a QUIC connection’s initial maximum number of bidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_bidirectional(_:_:)
func nw_quic_set_initial_max_streams_bidirectional(options unsafe.Pointer, initial_max_streams_bidirectional unsafe.Pointer) {
	_nw_quic_set_initial_max_streams_bidirectional(options, initial_max_streams_bidirectional)
	}


// Sets a QUIC connection’s initial maximum number of unidirectional streams. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_initial_max_streams_unidirectional(_:_:)
func nw_quic_set_initial_max_streams_unidirectional(options unsafe.Pointer, initial_max_streams_unidirectional unsafe.Pointer) {
	_nw_quic_set_initial_max_streams_unidirectional(options, initial_max_streams_unidirectional)
	}


// Sets the keepalive interval for the QUIC connection, in seconds. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_keepalive_interval(_:_:)
func nw_quic_set_keepalive_interval(metadata unsafe.Pointer, keepalive_interval unsafe.Pointer) {
	_nw_quic_set_keepalive_interval(metadata, keepalive_interval)
	}


// Sets the maximum number of bidirectional streams that the peer can create on a QUIC connection. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_bidirectional(_:_:)
func nw_quic_set_local_max_streams_bidirectional(metadata unsafe.Pointer, max_streams_bidirectional unsafe.Pointer) {
	_nw_quic_set_local_max_streams_bidirectional(metadata, max_streams_bidirectional)
	}


// Sets the maximum number of unidirectional streams that the peer can create on a QUIC connection. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_local_max_streams_unidirectional(_:_:)
func nw_quic_set_local_max_streams_unidirectional(metadata unsafe.Pointer, max_streams_unidirectional unsafe.Pointer) {
	_nw_quic_set_local_max_streams_unidirectional(metadata, max_streams_unidirectional)
	}


// Sets a QUIC connection’s maximum DATAGRAM frame size. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_max_datagram_frame_size(_:_:)
func nw_quic_set_max_datagram_frame_size(options unsafe.Pointer, max_datagram_frame_size unsafe.Pointer) {
	_nw_quic_set_max_datagram_frame_size(options, max_datagram_frame_size)
	}


// Sets the maximum length of a QUIC packet that can be received on a connection, in bytes. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_max_udp_payload_size(_:_:)
func nw_quic_set_max_udp_payload_size(options unsafe.Pointer, max_udp_payload_size unsafe.Pointer) {
	_nw_quic_set_max_udp_payload_size(options, max_udp_payload_size)
	}


// Sets the QUIC application error code to send for the stream. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_application_error(_:_:)
func nw_quic_set_stream_application_error(metadata unsafe.Pointer, application_error unsafe.Pointer) {
	_nw_quic_set_stream_application_error(metadata, application_error)
	}


// Configures a QUIC stream as a datagram flow, instead of a byte stream. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_datagram(_:_:)
func nw_quic_set_stream_is_datagram(options unsafe.Pointer, is_datagram bool) {
	_nw_quic_set_stream_is_datagram(options, is_datagram)
	}


// Configures a QUIC stream as unidirectional, instead of bidirectional. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_quic_set_stream_is_unidirectional(_:_:)
func nw_quic_set_stream_is_unidirectional(options unsafe.Pointer, is_unidirectional bool) {
	_nw_quic_set_stream_is_unidirectional(options, is_unidirectional)
	}


// Releases a reference count on a Network.framework object. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_release
func nw_release(obj unsafe.Pointer) {
	_nw_release(obj)
	}


// Accesses the resolved endpoint that the connection used for its first connection attempt. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_preferred_endpoint(_:)
func nw_resolution_report_copy_preferred_endpoint(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_copy_preferred_endpoint(resolution_report)
	}


// Accesses the resolved endpoint that led to the established connection. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_copy_successful_endpoint(_:)
func nw_resolution_report_copy_successful_endpoint(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_copy_successful_endpoint(resolution_report)
	}


// Accesses the number of endpoints resolved in this step. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_endpoint_count(_:)
func nw_resolution_report_get_endpoint_count(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_get_endpoint_count(resolution_report)
	}


// Accesses the duration of this resolution step, from when the query was issued to when the response was complete. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_milliseconds(_:)
func nw_resolution_report_get_milliseconds(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_get_milliseconds(resolution_report)
	}


// Accesses the transport protocol your connection used for DNS resolution. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_protocol(_:)
func nw_resolution_report_get_protocol(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_get_protocol(resolution_report)
	}


// Accesses the source of the DNS response. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_resolution_report_get_source(_:)
func nw_resolution_report_get_source(resolution_report unsafe.Pointer) unsafe.Pointer {
	return _nw_resolution_report_get_source(resolution_report)
	}


// Adds a reference count to a Network.framework object. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_retain
func nw_retain(obj unsafe.Pointer) unsafe.Pointer {
	return _nw_retain(obj)
	}


// Initializes a default set of TCP connection options. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_create_options()
func nw_tcp_create_options() unsafe.Pointer {
	return _nw_tcp_create_options()
	}


// Accesses the number of available bytes in the TCP receive buffer. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_get_available_receive_buffer(_:)
func nw_tcp_get_available_receive_buffer(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_tcp_get_available_receive_buffer(metadata)
	}


// Accesses the number of available bytes in the TCP send buffer. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_get_available_send_buffer(_:)
func nw_tcp_get_available_send_buffer(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_tcp_get_available_send_buffer(metadata)
	}


// Sets the number of seconds that TCP waits before timing out its handshake. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_connection_timeout(_:_:)
func nw_tcp_options_set_connection_timeout(options unsafe.Pointer, connection_timeout unsafe.Pointer) {
	_nw_tcp_options_set_connection_timeout(options, connection_timeout)
	}


// Disables TCP acknowledgment stretching. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ack_stretching(_:_:)
func nw_tcp_options_set_disable_ack_stretching(options unsafe.Pointer, disable_ack_stretching bool) {
	_nw_tcp_options_set_disable_ack_stretching(options, disable_ack_stretching)
	}


// Disables negotiation of Explicit Congestion Notification markings. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_disable_ecn(_:_:)
func nw_tcp_options_set_disable_ecn(options unsafe.Pointer, disable_ecn bool) {
	_nw_tcp_options_set_disable_ecn(options, disable_ecn)
	}


// Enables TCP Fast Open on a connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_fast_open(_:_:)
func nw_tcp_options_set_enable_fast_open(options unsafe.Pointer, enable_fast_open bool) {
	_nw_tcp_options_set_enable_fast_open(options, enable_fast_open)
	}


// Enables TCP keepalives. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_enable_keepalive(_:_:)
func nw_tcp_options_set_enable_keepalive(options unsafe.Pointer, enable_keepalive bool) {
	_nw_tcp_options_set_enable_keepalive(options, enable_keepalive)
	}


// Sets the number of keepalive probes that TCP sends before terminating the connection. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_count(_:_:)
func nw_tcp_options_set_keepalive_count(options unsafe.Pointer, keepalive_count unsafe.Pointer) {
	_nw_tcp_options_set_keepalive_count(options, keepalive_count)
	}


// Sets the number of seconds of idleness that TCP waits before sending keepalive probes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_idle_time(_:_:)
func nw_tcp_options_set_keepalive_idle_time(options unsafe.Pointer, keepalive_idle_time unsafe.Pointer) {
	_nw_tcp_options_set_keepalive_idle_time(options, keepalive_idle_time)
	}


// Sets the number of seconds that TCP waits between sending keepalive probes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_keepalive_interval(_:_:)
func nw_tcp_options_set_keepalive_interval(options unsafe.Pointer, keepalive_interval unsafe.Pointer) {
	_nw_tcp_options_set_keepalive_interval(options, keepalive_interval)
	}


// Sets TCP’s maximum segment size in bytes. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_maximum_segment_size(_:_:)
func nw_tcp_options_set_maximum_segment_size(options unsafe.Pointer, maximum_segment_size unsafe.Pointer) {
	_nw_tcp_options_set_maximum_segment_size(options, maximum_segment_size)
	}


// nw_tcp_options_set_multipath_force_version is a Network function. [Full Topic]
//
// Added in macOS 12.0.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_multipath_force_version(_:_:)
func nw_tcp_options_set_multipath_force_version(options unsafe.Pointer, multipath_force_version unsafe.Pointer) {
	_nw_tcp_options_set_multipath_force_version(options, multipath_force_version)
	}


// Disables Nagle’s algorithm for TCP. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_delay(_:_:)
func nw_tcp_options_set_no_delay(options unsafe.Pointer, no_delay bool) {
	_nw_tcp_options_set_no_delay(options, no_delay)
	}


// Sets TCP into no-options mode. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_options(_:_:)
func nw_tcp_options_set_no_options(options unsafe.Pointer, no_options bool) {
	_nw_tcp_options_set_no_options(options, no_options)
	}


// Sets TCP into no-push mode. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_no_push(_:_:)
func nw_tcp_options_set_no_push(options unsafe.Pointer, no_push bool) {
	_nw_tcp_options_set_no_push(options, no_push)
	}


// Sets the TCP persist timeout in seconds, as defined by RFC 6429. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_persist_timeout(_:_:)
func nw_tcp_options_set_persist_timeout(options unsafe.Pointer, persist_timeout unsafe.Pointer) {
	_nw_tcp_options_set_persist_timeout(options, persist_timeout)
	}


// Sets the number of seconds that TCP waits between retransmission attempts. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_connection_drop_time(_:_:)
func nw_tcp_options_set_retransmit_connection_drop_time(options unsafe.Pointer, retransmit_connection_drop_time unsafe.Pointer) {
	_nw_tcp_options_set_retransmit_connection_drop_time(options, retransmit_connection_drop_time)
	}


// Causes TCP to drop its connection after not receiving an ACK after a FIN. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tcp_options_set_retransmit_fin_drop(_:_:)
func nw_tcp_options_set_retransmit_fin_drop(options unsafe.Pointer, retransmit_fin_drop bool) {
	_nw_tcp_options_set_retransmit_fin_drop(options, retransmit_fin_drop)
	}


// Accesses the result of the TLS handshake. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_metadata(_:)
func nw_tls_copy_sec_protocol_metadata(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_tls_copy_sec_protocol_metadata(metadata)
	}


// Accesses the handshake security options TLS will use. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_copy_sec_protocol_options(_:)
func nw_tls_copy_sec_protocol_options(options unsafe.Pointer) unsafe.Pointer {
	return _nw_tls_copy_sec_protocol_options(options)
	}


// Initializes a default set of TLS connection options. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_tls_create_options()
func nw_tls_create_options() unsafe.Pointer {
	return _nw_tls_create_options()
	}


// Accesses the raw bytes contained within a TXT record. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_access_bytes(_:_:)
func nw_txt_record_access_bytes(txt_record unsafe.Pointer, access_bytes unsafe.Pointer) bool {
	return _nw_txt_record_access_bytes(txt_record, access_bytes)
	}


// Accesses the value for a specific key in a TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_access_key(_:_:_:)
func nw_txt_record_access_key(txt_record unsafe.Pointer, key unsafe.Pointer, access_value unsafe.Pointer) bool {
	return _nw_txt_record_access_key(txt_record, key, access_value)
	}


// Iterates through all keys in a TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_apply(_:_:)
func nw_txt_record_apply(txt_record unsafe.Pointer, applier unsafe.Pointer) bool {
	return _nw_txt_record_apply(txt_record, applier)
	}


// Performs a deep copy of a TXT record. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_copy(_:)
func nw_txt_record_copy(txt_record unsafe.Pointer) unsafe.Pointer {
	return _nw_txt_record_copy(txt_record)
	}


// Initializes a TXT record as a dictionary of strings. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_create_dictionary()
func nw_txt_record_create_dictionary() unsafe.Pointer {
	return _nw_txt_record_create_dictionary()
	}


// Initializes a TXT record with raw bytes. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_create_with_bytes(_:_:)
func nw_txt_record_create_with_bytes(txt_bytes unsafe.Pointer, txt_len unsafe.Pointer) unsafe.Pointer {
	return _nw_txt_record_create_with_bytes(txt_bytes, txt_len)
	}


// Checks the status of value associated with a key in a TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_find_key(_:_:)
func nw_txt_record_find_key(txt_record unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _nw_txt_record_find_key(txt_record, key)
	}


// Accesses the number of keys stored in the TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_get_key_count(_:)
func nw_txt_record_get_key_count(txt_record unsafe.Pointer) unsafe.Pointer {
	return _nw_txt_record_get_key_count(txt_record)
	}


// Checks whether a TXT record conforms to a dictionary format. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_is_dictionary(_:)
func nw_txt_record_is_dictionary(txt_record unsafe.Pointer) bool {
	return _nw_txt_record_is_dictionary(txt_record)
	}


// Checks whether two TXT records are equivalent. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_is_equal(_:_:)
func nw_txt_record_is_equal(left unsafe.Pointer, right unsafe.Pointer) bool {
	return _nw_txt_record_is_equal(left, right)
	}


// Removes a data value in a TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_remove_key(_:_:)
func nw_txt_record_remove_key(txt_record unsafe.Pointer, key unsafe.Pointer) bool {
	return _nw_txt_record_remove_key(txt_record, key)
	}


// Sets a data value in a TXT record dictionary. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_txt_record_set_key(_:_:_:_:)
func nw_txt_record_set_key(txt_record unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, value_len unsafe.Pointer) bool {
	return _nw_txt_record_set_key(txt_record, key, value, value_len)
	}


// Initializes a default UDP message. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_create_metadata()
func nw_udp_create_metadata() unsafe.Pointer {
	return _nw_udp_create_metadata()
	}


// Initializes a default set of UDP connection options. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_create_options()
func nw_udp_create_options() unsafe.Pointer {
	return _nw_udp_create_options()
	}


// Configures the connection to not send UDP checksums. [Full Topic]
//
// Added in macOS 10.14.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_udp_options_set_prefer_no_checksum(_:_:)
func nw_udp_options_set_prefer_no_checksum(options unsafe.Pointer, prefer_no_checksum bool) {
	_nw_udp_options_set_prefer_no_checksum(options, prefer_no_checksum)
	}


// Initializes a WebSocket message with a specific type code. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_create_metadata(_:)
func nw_ws_create_metadata(opcode unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_create_metadata(opcode)
	}


// Initializes a default set of WebSocket connection options. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_create_options(_:)
func nw_ws_create_options(version unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_create_options(version)
	}


// Accesses the WebSocket server’s response sent during the handshake. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_copy_server_response(_:)
func nw_ws_metadata_copy_server_response(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_metadata_copy_server_response(metadata)
	}


// Accesses the close code on a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_close_code(_:)
func nw_ws_metadata_get_close_code(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_metadata_get_close_code(metadata)
	}


// Checks the type code on a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_get_opcode(_:)
func nw_ws_metadata_get_opcode(metadata unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_metadata_get_opcode(metadata)
	}


// Sets a close code on a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_close_code(_:_:)
func nw_ws_metadata_set_close_code(metadata unsafe.Pointer, close_code unsafe.Pointer) {
	_nw_ws_metadata_set_close_code(metadata, close_code)
	}


// Sets a handler on a Ping message to be invoked when the corresponding Pong message is received. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_metadata_set_pong_handler(_:_:_:)
func nw_ws_metadata_set_pong_handler(metadata unsafe.Pointer, client_queue unsafe.Pointer, pong_handler unsafe.Pointer) {
	_nw_ws_metadata_set_pong_handler(metadata, client_queue, pong_handler)
	}


// Adds additional HTTP header fields to be sent by the client during the WebSocket handshake. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_add_additional_header(_:_:_:)
func nw_ws_options_add_additional_header(options unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	_nw_ws_options_add_additional_header(options, name, value)
	}


// Adds to the list of supported application protocols that will be presented to a WebSocket server during connection establishment. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_add_subprotocol(_:_:)
func nw_ws_options_add_subprotocol(options unsafe.Pointer, subprotocol unsafe.Pointer) {
	_nw_ws_options_add_subprotocol(options, subprotocol)
	}


// Configures the connection to automatically reply to Ping messages instead of delivering them to you. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_auto_reply_ping(_:_:)
func nw_ws_options_set_auto_reply_ping(options unsafe.Pointer, auto_reply_ping bool) {
	_nw_ws_options_set_auto_reply_ping(options, auto_reply_ping)
	}


// Sets a handler to react to as a server to inbound WebSocket client handshakes. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_client_request_handler(_:_:_:)
func nw_ws_options_set_client_request_handler(options unsafe.Pointer, client_queue unsafe.Pointer, handler unsafe.Pointer) {
	_nw_ws_options_set_client_request_handler(options, client_queue, handler)
	}


// Sets the maximum allowed message size, in bytes, to be received by the WebSocket connection. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_maximum_message_size(_:_:)
func nw_ws_options_set_maximum_message_size(options unsafe.Pointer, maximum_message_size unsafe.Pointer) {
	_nw_ws_options_set_maximum_message_size(options, maximum_message_size)
	}


// Specifies whether the WebSocket protocol skips its handshake and begins framing data once the underlying connection is established. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_options_set_skip_handshake(_:_:)
func nw_ws_options_set_skip_handshake(options unsafe.Pointer, skip_handshake bool) {
	_nw_ws_options_set_skip_handshake(options, skip_handshake)
	}


// Enumerates additional HTTP headers in a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_additional_headers(_:_:)
func nw_ws_request_enumerate_additional_headers(request unsafe.Pointer, enumerator unsafe.Pointer) bool {
	return _nw_ws_request_enumerate_additional_headers(request, enumerator)
	}


// Enumerates the supported subprotocols in a WebSocket message. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_request_enumerate_subprotocols(_:_:)
func nw_ws_request_enumerate_subprotocols(request unsafe.Pointer, enumerator unsafe.Pointer) bool {
	return _nw_ws_request_enumerate_subprotocols(request, enumerator)
	}


// Adds an additional HTTP header to a WebSocket server response. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_add_additional_header(_:_:_:)
func nw_ws_response_add_additional_header(response unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) {
	_nw_ws_response_add_additional_header(response, name, value)
	}


// Initializes a WebSocket server response with a status and selected subprotocol. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_create(_:_:)
func nw_ws_response_create(status unsafe.Pointer, selected_subprotocol unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_response_create(status, selected_subprotocol)
	}


// Enumerates the additional HTTP headers in a WebSocket server response. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_enumerate_additional_headers(_:_:)
func nw_ws_response_enumerate_additional_headers(response unsafe.Pointer, enumerator unsafe.Pointer) bool {
	return _nw_ws_response_enumerate_additional_headers(response, enumerator)
	}


// Accesses the selected subprotocol in a WebSocket server response. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_get_selected_subprotocol(_:)
func nw_ws_response_get_selected_subprotocol(response unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_response_get_selected_subprotocol(response)
	}


// Accesses the status of a WebSocket server response. [Full Topic]
//
// Added in macOS 10.15.
//
// [Full Topic]: https://developer.apple.com/documentation/Network/nw_ws_response_get_status(_:)
func nw_ws_response_get_status(response unsafe.Pointer) unsafe.Pointer {
	return _nw_ws_response_get_status(response)
	}




