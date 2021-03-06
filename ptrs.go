package ptrs

import "time"

// NewIntPtr returns pointer of which value is v.
func NewIntPtr(v int) *int {
	return &v
}

// NewIntPtrZeroNil returns pointer of which value is v when v is not zero value,
// otherwise returns nil.
func NewIntPtrZeroNil(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}

// NewStringPtrEmptyNil returns pointer of which value is v when v is not zero value,
// otherwise returns nil.
func NewStringPtrEmptyNil(v string) *string {
	if v == "" {
		return nil
	}
	return &v
}

// StringFromPtr returns string value of p when p is not nil, otherwise returns "".
func StringFromPtr(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

// TimePtrToUnixSecPtr returns pointer of which value is v.
func TimePtrToUnixSecPtr(t *time.Time) *int64 {
	if t == nil || t.IsZero() {
		return nil
	}
	unix := t.Unix()
	return &unix
}

// TimeToUnixSecPtr returns pointer of which value is v.
func TimeToUnixSecPtr(t time.Time) *int64 {
	if t.IsZero() {
		return nil
	}
	unix := t.Unix()
	return &unix
}
