// MIT License

// Copyright (c) 2022 Project-Sparrow
// Copyright (c) 2023 Kelwing <kelwing@kelnet.org>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package snowflake

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
)

var nullBytes = []byte("null")

// NullSnowflake is a nullable Snowflake
type NullSnowflake struct {
	Snowflake Snowflake
	Valid     bool
}

// NewNullSnowflake creates a new NullSnowflake
func NewNullSnowflake(s Snowflake, valid bool) NullSnowflake {
	return NullSnowflake{
		Snowflake: s,
		Valid:     valid,
	}
}

// NullSnowflakeFromPtr creates a new NullSnowflake from a Snowflake pointer
func NullSnowflakeFromPtr(s *Snowflake) NullSnowflake {
	if s == nil {
		return NewNullSnowflake(Snowflake(0), false)
	}

	return NewNullSnowflake(*s, true)
}

// NullSnowflakeFromStringPtr creates a new NullSnowflake from a string pointer
// Always returns a NullSnowflake, will be invalid if the string cannot be converted to an int
func NullSnowflakeFromStringPtr(s *string) NullSnowflake {
	if s == nil {
		return NewNullSnowflake(Snowflake(0), false)
	}

	snowflake, err := SnowflakeFromString(*s)
	if err != nil {
		return NewNullSnowflake(Snowflake(0), false)
	}

	return NewNullSnowflake(snowflake, true)
}

// Scan implements sql.Scanner interface
func (s *NullSnowflake) Scan(value any) error {
	if value == nil {
		s.Snowflake, s.Valid = Snowflake(0), false
		return nil
	}
	s.Valid = true
	return (&s.Snowflake).Scan(value)
}

// Value implements driver.Valuer interface
func (s NullSnowflake) Value() (driver.Value, error) {
	if !s.Valid {
		return nil, nil
	}
	return s.Snowflake.Value()
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (s NullSnowflake) ValueOrZero() Snowflake {
	if !s.Valid {
		return Snowflake(0)
	}
	return s.Snowflake
}

// UnmarshalJSON implements json.Unmarshaler.
func (s *NullSnowflake) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, nullBytes) {
		s.Valid = false
		return nil
	}

	if err := json.Unmarshal(data, &s.Snowflake); err != nil {
		return err
	}

	s.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler.
func (s NullSnowflake) MarshalJSON() ([]byte, error) {
	if !s.Valid {
		return nullBytes, nil
	}

	return s.Snowflake.MarshalJSON()
}
