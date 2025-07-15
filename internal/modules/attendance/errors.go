package attendance

import "errors"

var ErrAlreadyCheckedIn = errors.New("already checked in")
var ErrAlreadyCheckedOut = errors.New("already checked out")
