package models

const (
	// user status
	USER_NORMAL = 1
	USER_ADMIN  = 2

	// user action info
	U_DO_SUCCESS = 0
	U_DO_REMAIN  = 1

	U_DEL_ERROR   = 100
	U_DEL_SELF    = 101
	U_DEL_MANAGER = 102

	U_PASS_WRONG = 100
	U_PASS_UPERR = 110

	// message request code
	MSG_OK   = 0
	MSG_FAIL = 404
)
