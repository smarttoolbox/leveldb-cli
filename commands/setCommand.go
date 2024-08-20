// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package commands

import (
	"fmt"

	"github.com/smarttoolbox/leveldb-cli/cliutil"
)

// The command set a value.
// It sets the value for the selected key.
//
// Returns a string containing information about the result of the operation.
func Set(key, value, format string) string {
	if !isConnected {
		return AppError(ErrDbDoesNotOpen)
	}

	if key == "" {
		return AppError(ErrKeyIsEmpty)
	}
	v := []byte(value)
	if format != "" {
		v = cliutil.FromString(format, value)
	}
	err := dbh.Put([]byte(key), v, nil)
	if err != nil {
		return fmt.Sprintf(
			AppError(ErrUnableToWrite),
			err.Error(),
		)
	}

	return "Success"
}
