// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package common

import "gnbsim/logger"

type EventType uint32

/* Using the MSB to represent the interface to which the event belongs.*/
const (
	/* Network interface events */
	UU_EVENT EventType = 0x1000000
	N1_EVENT EventType = 0x2000000
	N2_EVENT EventType = 0x3000000
	N3_EVENT EventType = 0x4000000

	/* Application's interfaces events */
	PROFILE_SIMUE_EVENT EventType = 0x5000000
	SIMUE_REALUE_EVENT  EventType = 0x6000000
	COMMON_EVENT        EventType = 0x7000000
)

const (
	INIT_EVENT EventType = COMMON_EVENT + 1 + iota
	QUIT_EVENT
	ERROR_EVENT
)

/* Events between Profile and SimUe */
const (
	PROFILE_START_EVENT EventType = PROFILE_SIMUE_EVENT + 1 + iota
	PROFILE_PASS_EVENT
	PROFILE_FAIL_EVENT
)

/* Events between SimUe and RealUE */
const (
	DATA_PKT_GEN_REQUEST_EVENT EventType = SIMUE_REALUE_EVENT + 1 + iota
	DATA_PKT_GEN_SUCCESS_EVENT
	DATA_PKT_GEN_FAILURE_EVENT
)

/* Events between UE and GNodeB (UU) */
const (
	CONNECTION_REQUEST_EVENT EventType = UU_EVENT + 1 + iota
	CONNECTION_RELEASE_REQUEST_EVENT

	UL_INFO_TRANSFER_EVENT
	DL_INFO_TRANSFER_EVENT

	// For exchanging user data
	UL_UE_DATA_TRANSFER_EVENT
	DL_UE_DATA_TRANSFER_EVENT
	END_MARKER_EVENT

	// For setting up channels between UE and GNB to exchange user data
	DATA_BEARER_SETUP_REQUEST_EVENT
	DATA_BEARER_SETUP_RESPONSE_EVENT

	// gNB acknowledges simue for releasing UE context using this event
	CTX_RELEASE_ACKNOWLEDGEMENT_EVENT

	// SimUe commands gNB to trigger RAN Connection release which further
	// triggers gNB initiated UE Context Release Request
	TRIGGER_AN_RELEASE_EVENT
)

/* Events betweem UE and AMF (N1)
 * Following events offsets are numbered same as the NAS Message types in
 * section 9.7 of 3GPP TS 24.501. This makes it easy to generate events
 * corresponding to received NAS messages.
 */
const UE_5GS_MOBILITY_MANAGEMENT_EVENTS EventType = N1_EVENT + 64
const UE_5GS_SESSION_MANAGEMENT_EVENTS EventType = N1_EVENT + 192

// 5GS Mobility Management events.
const (
	_ EventType = UE_5GS_MOBILITY_MANAGEMENT_EVENTS + iota

	REG_REQUEST_EVENT //65
	REG_ACCEPT_EVENT
	REG_COMPLETE_EVENT
	REG_REJECT_EVENT
	DEREG_REQUEST_UE_ORIG_EVENT
	DEREG_ACCEPT_UE_ORIG_EVENT
	DEREG_REQUEST_UE_TERM_EVENT
	DEREG_ACCEPT_UE_TERM_EVENT //72

	SERVICE_REQUEST_EVENT = UE_5GS_MOBILITY_MANAGEMENT_EVENTS + 3 + iota //76
	SERVICE_REJECT_EVENT
	SERVICE_ACCEPT_EVENT //78

	AUTH_REQUEST_EVENT = UE_5GS_MOBILITY_MANAGEMENT_EVENTS + 10 + iota //86
	AUTH_RESPONSE_EVENT
	AUTH_REJECT_EVENT
	AUTH_FAILURE_EVENT
	AUTH_RESULT_EVENT
	ID_REQUEST_EVENT
	ID_RESPONSE_EVENT
	SEC_MOD_COMMAND_EVENT
	SEC_MOD_COMPLETE_EVENT
	SEC_MOD_REJECT_EVENT //95

	FIVE_GMM_STATUS_EVENT = UE_5GS_MOBILITY_MANAGEMENT_EVENTS + 14 + iota //86
	NOTIFICATION_EVENT
	NOTIFICATION_RESPONSE_EVENT
	UL_NAS_TRANSPORT_EVENT
	DL_NAS_TRANSPORT_EVENT
)

// 5GS Session Management events.
const (
	_ EventType = UE_5GS_SESSION_MANAGEMENT_EVENTS + iota

	PDU_SESS_EST_REQUEST_EVENT //193
	PDU_SESS_EST_ACCEPT_EVENT
	PDU_SESS_EST_REJECT_EVENT

	PDU_SESS_AUTH_COMMAND_EVENT = UE_5GS_SESSION_MANAGEMENT_EVENTS + 1 + iota
	PDU_SESS_AUTH_COMPLETE_EVENT
	PDU_SESS_AUTH_RESULT_EVENT

	PDU_SESS_MOD_REQUEST_EVENT = UE_5GS_SESSION_MANAGEMENT_EVENTS + 2 + iota
	PDU_SESS_MOD_REJECT_EVENT
	PDU_SESS_MOD_COMMAND_EVENT
	PDU_SESS_MOD_COMPLETE_EVENT
	PDU_SESS_MOD_CMD_REJECT_EVENT

	PDU_SESS_REL_REQUEST_EVENT = UE_5GS_SESSION_MANAGEMENT_EVENTS + 5 + iota
	PDU_SESS_REL_REJECT_EVENT
	PDU_SESS_REL_COMMAND_EVENT
	PDU_SESS_REL_COMPLETE_EVENT

	FIVEGSM_STATUS_EVENT = UE_5GS_SESSION_MANAGEMENT_EVENTS + 6 + iota //214
)

// Events between GNodeB and AMF (N2)
const (
	DOWNLINK_NAS_TRANSPORT_EVENT EventType = N2_EVENT + 1 + iota
	INITIAL_CTX_SETUP_REQUEST_EVENT
	PDU_SESS_RESOURCE_SETUP_REQUEST_EVENT
	UE_CTX_RELEASE_COMMAND_EVENT
)

// Events between GNodeB and UPF (N3)
const (
	DL_UE_DATA_TRANSPORT_EVENT EventType = N3_EVENT + 1 + iota
)

var evtStrMap map[EventType]string = map[EventType]string{
	INIT_EVENT:                            "INIT-EVENT",
	QUIT_EVENT:                            "QUIT-EVENT",
	ERROR_EVENT:                           "ERROR-EVENT",
	PROFILE_START_EVENT:                   "PROFILE-START-EVENT",
	PROFILE_PASS_EVENT:                    "PROFILE-PASS-EVENT",
	PROFILE_FAIL_EVENT:                    "PROFILE-FAIL-EVENT",
	DATA_PKT_GEN_REQUEST_EVENT:            "DATA-PACKET-GENERATION-REQUEST-EVENT",
	DATA_PKT_GEN_SUCCESS_EVENT:            "DATA-PACKET-SUCCESS-EVENT",
	DATA_PKT_GEN_FAILURE_EVENT:            "DATA-PACKET-FAILURE-EVENT",
	CONNECTION_REQUEST_EVENT:              "CONNECTION-REQUEST-EVENT",
	CONNECTION_RELEASE_REQUEST_EVENT:      "CONNECTION-RELEASE-REQUEST-EVENT",
	UL_INFO_TRANSFER_EVENT:                "UL-INFO-TRANSFER-EVENT",
	DL_INFO_TRANSFER_EVENT:                "DL-INFO-TRANSFER-EVENT",
	UL_UE_DATA_TRANSFER_EVENT:             "UL-UE-DATA-TRANSFER-EVENT",
	DL_UE_DATA_TRANSFER_EVENT:             "DL-UE-DATA-TRANSFER-EVENT",
	END_MARKER_EVENT:                      "END-MARKER-EVENT",
	DATA_BEARER_SETUP_REQUEST_EVENT:       "DATA-BEARER-SETUP-REQUEST-EVENT",
	DATA_BEARER_SETUP_RESPONSE_EVENT:      "DATA-BEARER-SETUP-RESPONSE-EVENT",
	CTX_RELEASE_ACKNOWLEDGEMENT_EVENT:     "CONTEXT-RELEASE-ACKNOWLEDGEMENT-EVENT",
	TRIGGER_AN_RELEASE_EVENT:              "TRIGGER-AN-RELEASE-EVENT",
	REG_REQUEST_EVENT:                     "REGESTRATION-REQUEST-EVENT",
	REG_ACCEPT_EVENT:                      "REGESTRATION-ACCEPT-EVENT",
	REG_COMPLETE_EVENT:                    "REGESTRATION-COMPLETE-EVENT",
	REG_REJECT_EVENT:                      "REGESTRATION-REJECT-EVENT",
	DEREG_REQUEST_UE_ORIG_EVENT:           "DEREGISTRATION-REQUEST-UE-ORIG-EVENT",
	DEREG_ACCEPT_UE_ORIG_EVENT:            "DEREGISTRATION-ACCEPT-UE-ORIG-EVENT",
	DEREG_REQUEST_UE_TERM_EVENT:           "DEREGISTRATION-REQUEST-UE-TERM-EVENT",
	DEREG_ACCEPT_UE_TERM_EVENT:            "DEREGISTRATION-ACCEPT-UE-TERM-EVENT",
	SERVICE_REQUEST_EVENT:                 "SERVICE-REQUEST-EVENT",
	SERVICE_REJECT_EVENT:                  "SERVICE-REJECT-EVENT",
	SERVICE_ACCEPT_EVENT:                  "SERVICE-ACCEPT-EVENT",
	AUTH_REQUEST_EVENT:                    "AUTHENTICATION-REQUEST-EVENT",
	AUTH_RESPONSE_EVENT:                   "AUTHENTICATION-RESPONSE-EVENT",
	AUTH_REJECT_EVENT:                     "AUTHENTICATION-REJECT-EVENT",
	AUTH_FAILURE_EVENT:                    "AUTHENTICATION-FAILURE-EVENT",
	AUTH_RESULT_EVENT:                     "AUTHENTICATION-RESULT-EVENT",
	ID_REQUEST_EVENT:                      "ID-REQUEST-EVENT",
	ID_RESPONSE_EVENT:                     "ID-RESPONSE-EVENT",
	SEC_MOD_COMMAND_EVENT:                 "SECURITY-MODE-COMMAND-EVENT",
	SEC_MOD_COMPLETE_EVENT:                "SECURITY-MODE-COMPLETE-EVENT",
	SEC_MOD_REJECT_EVENT:                  "SECURITY-MODE-REJECT-EVENT",
	FIVE_GMM_STATUS_EVENT:                 "FIVE-GMM-STATUS-EVENT",
	NOTIFICATION_EVENT:                    "NOTIFICATION-EVENT",
	NOTIFICATION_RESPONSE_EVENT:           "NOTIFICATION-RESPONSE-EVENT",
	UL_NAS_TRANSPORT_EVENT:                "UL-NAS-TRANSPORT-EVENT",
	DL_NAS_TRANSPORT_EVENT:                "DL-NAS-TRANSPORT-EVENT",
	PDU_SESS_EST_REQUEST_EVENT:            "PDU-SESSION-ESTABLISHMENT-REQUEST-EVENT",
	PDU_SESS_EST_ACCEPT_EVENT:             "PDU-SESSION-ESTABLISHMENT-ACCEPT-EVENT",
	PDU_SESS_EST_REJECT_EVENT:             "PDU-SESSION-ESTABLISHMENT-REJECT-EVENT",
	PDU_SESS_AUTH_COMMAND_EVENT:           "PDU-SESSION-AUTHENTICATION-COMMAND-EVENT",
	PDU_SESS_AUTH_COMPLETE_EVENT:          "PDU-SESSION-AUTHENTICATION-COMPLETE-EVENT",
	PDU_SESS_AUTH_RESULT_EVENT:            "PDU-SESSION-AUTHENTICATION-RESULT-EVENT",
	PDU_SESS_MOD_REQUEST_EVENT:            "PDU-SESSION-MODIFICATION-REQUEST-EVENT",
	PDU_SESS_MOD_REJECT_EVENT:             "PDU-SESSION-MODIFICATION-REJECT-EVENT",
	PDU_SESS_MOD_COMMAND_EVENT:            "PDU-SESSION-MODIFICATION-COMMAND-EVENT",
	PDU_SESS_MOD_COMPLETE_EVENT:           "PDU-SESSION-MODIFICATION-COMPLETE-EVENT",
	PDU_SESS_MOD_CMD_REJECT_EVENT:         "PDU-SESSION-MODIFICATION-COMMAND-REJECT-EVENT",
	PDU_SESS_REL_REQUEST_EVENT:            "PDU-SESSION-RELEASE-REQUEST-EVENT",
	PDU_SESS_REL_REJECT_EVENT:             "PDU-SESSION-RELEASE-REJECT-EVENT",
	PDU_SESS_REL_COMMAND_EVENT:            "PDU-SESSION-RELEASE-COMMAND-EVENT",
	PDU_SESS_REL_COMPLETE_EVENT:           "PDU-SESSION-RELEASE-COMPLETE-EVENT",
	FIVEGSM_STATUS_EVENT:                  "FIVEGSM-STATUS-EVENT",
	DOWNLINK_NAS_TRANSPORT_EVENT:          "DOWNLINK-NAS-TRANSPORT-EVENT",
	INITIAL_CTX_SETUP_REQUEST_EVENT:       "INITIAL-CONTEXT-SETUP-REQUEST-EVENT",
	PDU_SESS_RESOURCE_SETUP_REQUEST_EVENT: "PDU-SESSION-RESOURCE-SETUP-REQUEST-EVENT",
	UE_CTX_RELEASE_COMMAND_EVENT:          "UE-CONTEXT-RELEASE-COMMAND-EVENT",
	DL_UE_DATA_TRANSPORT_EVENT:            "DL-UE-DATA-TRANSPORT-EVENT",
}

func GetEvtString(id EventType) string {
	evtStr, ok := evtStrMap[id]
	if !ok {
		logger.AppLog.Fatalln("Invaid Event ID:", id)
	}
	return evtStr
}
