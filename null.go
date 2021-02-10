// Copyright 2021 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import "database/sql/driver"

type (
	// NullUUID represents a UUID that may be null.
	// NullUUID implements the Scanner interface so
	// it can be used as a scan destination:
	//
	//  var u uuid.NullUUID
	//  err := db.QueryRow("SELECT name FROM foo WHERE id=?", id).Scan(&u)
	//  ...
	//  if u.Valid {
	//     // use u.UUID
	//  } else {
	//     // NULL value
	//  }
	//
	NullUUID struct {
		UUID  UUID
		Valid bool // Valid is true if UUID is not NULL
	}
)

// Scan implements the Scanner interface.
func (nu *NullUUID) Scan(value interface{}) error {
	if value == nil {
		nu.UUID, nu.Valid = Nil, false
		return nil
	}

	// Delegate to UUID Scan function
	nu.Valid = true
	return nu.UUID.Scan(value)
}

// Value implements the driver Valuer interface.
func (nu NullUUID) Value() (driver.Value, error) {
	if !nu.Valid {
		return nil, nil
	}
	// Delegate to UUID Value function
	return nu.UUID.Value()
}
