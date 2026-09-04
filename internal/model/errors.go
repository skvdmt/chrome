package model

import (
	"errors"
	"fmt"
)

var ERR_TARGET_NOT_FOUND = errors.New("target not found")
var ERR_USER_DATA_DIR_NOT_SET = fmt.Errorf(
	"arg %s is not set", ARG_NAME_USER_DATA_DIR,
)
var ERR_RESPONSE_TIMEOUT = errors.New("response timeout")
var ERR_UNKNOWN_DATA = errors.New("unknown data incomint from chrome server")
var ERR_NO_TARGETS = errors.New("no targets")
var ERR_NO_SESSION = errors.New("no session")
var ERR_NODE_NOT_FOUND = errors.New("node not found")
