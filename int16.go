package null

import (
	"database/sql"
	"strconv"

	"github.com/guregu/null/v5/internal"
)

// Int16 is an nullable int16.
// It does not consider zero values to be null.
// It will decode to null, not zero, if null.
type Int16 struct {
	sql.NullInt16
}

// NewInt16 creates a new Int16.
func NewInt16(i int16, valid bool) Int16 {
	return Int16{
		NullInt16: sql.NullInt16{
			Int16: i,
			Valid: valid,
		},
	}
}

// Int16From creates a new Int16 that will always be valid.
func Int16From(i int16) Int16 {
	return NewInt16(i, true)
}

// Int16FromPtr creates a new Int16 that be null if i is nil.
func Int16FromPtr(i *int16) Int16 {
	if i == nil {
		return NewInt16(0, false)
	}
	return NewInt16(*i, true)
}

// ValueOrZero returns the inner value if valid, otherwise zero.
func (i Int16) ValueOrZero() int16 {
	if !i.Valid {
		return 0
	}
	return i.Int16
}

// ValueOr returns the inner value if valid, otherwise v.
func (i Int16) ValueOr(v int16) int16 {
	if !i.Valid {
		return v
	}
	return i.Int16
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports number, string, and null input.
// 0 will not be considered a null Int16.
func (i *Int16) UnmarshalJSON(data []byte) error {
	return internal.UnmarshalIntJSON(data, &i.Int16, &i.Valid, 16, strconv.ParseInt)
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null Int16 if the input is blank.
// It will return an error if the input is not an integer, blank, or "null".
func (i *Int16) UnmarshalText(text []byte) error {
	return internal.UnmarshalIntText(text, &i.Int16, &i.Valid, 16, strconv.ParseInt)
}

// MarshalJSON implements json.Marshaler.
// It will encode null if this Int16 is null.
func (i Int16) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(i.Int16), 10)), nil
}

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string if this Int16 is null.
func (i Int16) MarshalText() ([]byte, error) {
	if !i.Valid {
		return []byte{}, nil
	}
	return []byte(strconv.FormatInt(int64(i.Int16), 10)), nil
}

// SetValid changes this Int16's value and also sets it to be non-null.
func (i *Int16) SetValid(n int16) {
	i.Int16 = n
	i.Valid = true
}

// Ptr returns a pointer to this Int16's value, or a nil pointer if this Int16 is null.
func (i Int16) Ptr() *int16 {
	if !i.Valid {
		return nil
	}
	return &i.Int16
}

// IsZero returns true for invalid Int16s, for future omitempty support (Go 1.4?)
// A non-null Int16 with a 0 value will not be considered zero.
func (i Int16) IsZero() bool {
	return !i.Valid
}

// Equal returns true if both ints have the same value or are both null.
func (i Int16) Equal(other Int16) bool {
	return i.Valid == other.Valid && (!i.Valid || i.Int16 == other.Int16)
}

func (i Int16) value() (int64, bool) {
	return int64(i.Int16), i.Valid
}
