package etc

import (
	"database/sql"

	"github.com/golang/protobuf/ptypes/wrappers"
)

// Null String returns value if it is valid
func NullString(s *wrappers.StringValue) (ns sql.NullString) {
	if s != nil {
		ns.String = s.Value
		ns.Valid = true
	}
	return ns
}

//NullFloat64 returns value if it is valid
func NullDouble(s *wrappers.DoubleValue) (ns sql.NullFloat64) {
	if s != nil {
		ns.Float64 = s.Value
		ns.Valid = true
	}
	return ns
}

// String Value... [returns value of the string or nil]
func StringValue(ns sql.NullString) *wrappers.StringValue {
	if ns.Valid {
		s := wrappers.StringValue{Value: ns.String}
		return &s
	}
	return nil
}

//DoubleValue ...
func DoubleValue(ns sql.NullFloat64) *wrappers.DoubleValue {
	if ns.Valid {
		s := wrappers.DoubleValue{Value: ns.Float64}
		return &s
	}
	return nil
}

// The samething as DoubleValue
func Int64Value(ns sql.NullInt64) *wrappers.Int64Value {
	if ns.Valid {
		s := wrappers.Int64Value{Value: ns.Int64}
		return &s
	}
	return nil
}
