// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package enums

import (
	"fmt"
)

var (
	DLQOperationType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Merge":       1,
		"Purge":       2,
	}
)

// DLQOperationTypeFromString parses a DLQOperationType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to DLQOperationType
func DLQOperationTypeFromString(s string) (DLQOperationType, error) {
	if v, ok := DLQOperationType_value[s]; ok {
		return DLQOperationType(v), nil
	} else if v, ok := DLQOperationType_shorthandValue[s]; ok {
		return DLQOperationType(v), nil
	}
	return DLQOperationType(0), fmt.Errorf("%s is not a valid DLQOperationType", s)
}

var (
	DLQOperationState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Running":     1,
		"Completed":   2,
		"Failed":      3,
	}
)

// DLQOperationStateFromString parses a DLQOperationState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to DLQOperationState
func DLQOperationStateFromString(s string) (DLQOperationState, error) {
	if v, ok := DLQOperationState_value[s]; ok {
		return DLQOperationState(v), nil
	} else if v, ok := DLQOperationState_shorthandValue[s]; ok {
		return DLQOperationState(v), nil
	}
	return DLQOperationState(0), fmt.Errorf("%s is not a valid DLQOperationState", s)
}