// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "model/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "go run model/internal/cmd/pdatagen/main.go".

package plog

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	"go.opentelemetry.io/collector/pdata/internal/data"
	otlplogs "go.opentelemetry.io/collector/pdata/internal/data/protogen/logs/v1"
	"go.opentelemetry.io/collector/pdata/pcommon"
)

func TestResourceLogsSlice(t *testing.T) {
	es := NewResourceLogsSlice()
	assert.Equal(t, 0, es.Len())
	es = newResourceLogsSlice(&[]*otlplogs.ResourceLogs{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newResourceLogs(&otlplogs.ResourceLogs{})
	testVal := ResourceLogs(internal.GenerateTestResourceLogs())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestResourceLogs(internal.ResourceLogs(el))
		assert.Equal(t, testVal, el)
	}
}

func TestResourceLogsSlice_CopyTo(t *testing.T) {
	dest := NewResourceLogsSlice()
	// Test CopyTo to empty
	NewResourceLogsSlice().CopyTo(dest)
	assert.Equal(t, NewResourceLogsSlice(), dest)

	// Test CopyTo larger slice
	ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()).CopyTo(dest)
	assert.Equal(t, ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()), dest)

	// Test CopyTo same size slice
	ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()).CopyTo(dest)
	assert.Equal(t, ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()), dest)
}

func TestResourceLogsSlice_EnsureCapacity(t *testing.T) {
	es := ResourceLogsSlice(internal.GenerateTestResourceLogsSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.ResourceLogs]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.ResourceLogs]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlplogs.ResourceLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestResourceLogsSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := ResourceLogsSlice(internal.GenerateTestResourceLogsSlice())
	dest := NewResourceLogsSlice()
	src := ResourceLogsSlice(internal.GenerateTestResourceLogsSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	ResourceLogsSlice(internal.GenerateTestResourceLogsSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestResourceLogsSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewResourceLogsSlice()
	emptySlice.RemoveIf(func(el ResourceLogs) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := ResourceLogsSlice(internal.GenerateTestResourceLogsSlice())
	pos := 0
	filtered.RemoveIf(func(el ResourceLogs) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestResourceLogs_MoveTo(t *testing.T) {
	ms := ResourceLogs(internal.GenerateTestResourceLogs())
	dest := NewResourceLogs()
	ms.MoveTo(dest)
	assert.Equal(t, NewResourceLogs(), ms)
	assert.Equal(t, ResourceLogs(internal.GenerateTestResourceLogs()), dest)
}

func TestResourceLogs_CopyTo(t *testing.T) {
	ms := NewResourceLogs()
	orig := NewResourceLogs()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = ResourceLogs(internal.GenerateTestResourceLogs())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestResourceLogs_Resource(t *testing.T) {
	ms := NewResourceLogs()
	internal.FillTestResource(internal.Resource(ms.Resource()))
	assert.Equal(t, pcommon.Resource(internal.GenerateTestResource()), ms.Resource())
}

func TestResourceLogs_SchemaUrl(t *testing.T) {
	ms := NewResourceLogs()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestResourceLogs_ScopeLogs(t *testing.T) {
	ms := NewResourceLogs()
	assert.Equal(t, NewScopeLogsSlice(), ms.ScopeLogs())
	internal.FillTestScopeLogsSlice(internal.ScopeLogsSlice(ms.ScopeLogs()))
	assert.Equal(t, ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()), ms.ScopeLogs())
}

func TestScopeLogsSlice(t *testing.T) {
	es := NewScopeLogsSlice()
	assert.Equal(t, 0, es.Len())
	es = newScopeLogsSlice(&[]*otlplogs.ScopeLogs{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newScopeLogs(&otlplogs.ScopeLogs{})
	testVal := ScopeLogs(internal.GenerateTestScopeLogs())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestScopeLogs(internal.ScopeLogs(el))
		assert.Equal(t, testVal, el)
	}
}

func TestScopeLogsSlice_CopyTo(t *testing.T) {
	dest := NewScopeLogsSlice()
	// Test CopyTo to empty
	NewScopeLogsSlice().CopyTo(dest)
	assert.Equal(t, NewScopeLogsSlice(), dest)

	// Test CopyTo larger slice
	ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()).CopyTo(dest)
	assert.Equal(t, ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()), dest)

	// Test CopyTo same size slice
	ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()).CopyTo(dest)
	assert.Equal(t, ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()), dest)
}

func TestScopeLogsSlice_EnsureCapacity(t *testing.T) {
	es := ScopeLogsSlice(internal.GenerateTestScopeLogsSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.ScopeLogs]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.ScopeLogs]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.ScopeLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlplogs.ScopeLogs]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestScopeLogsSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := ScopeLogsSlice(internal.GenerateTestScopeLogsSlice())
	dest := NewScopeLogsSlice()
	src := ScopeLogsSlice(internal.GenerateTestScopeLogsSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	ScopeLogsSlice(internal.GenerateTestScopeLogsSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestScopeLogsSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewScopeLogsSlice()
	emptySlice.RemoveIf(func(el ScopeLogs) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := ScopeLogsSlice(internal.GenerateTestScopeLogsSlice())
	pos := 0
	filtered.RemoveIf(func(el ScopeLogs) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestScopeLogs_MoveTo(t *testing.T) {
	ms := ScopeLogs(internal.GenerateTestScopeLogs())
	dest := NewScopeLogs()
	ms.MoveTo(dest)
	assert.Equal(t, NewScopeLogs(), ms)
	assert.Equal(t, ScopeLogs(internal.GenerateTestScopeLogs()), dest)
}

func TestScopeLogs_CopyTo(t *testing.T) {
	ms := NewScopeLogs()
	orig := NewScopeLogs()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = ScopeLogs(internal.GenerateTestScopeLogs())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestScopeLogs_Scope(t *testing.T) {
	ms := NewScopeLogs()
	internal.FillTestInstrumentationScope(internal.InstrumentationScope(ms.Scope()))
	assert.Equal(t, pcommon.InstrumentationScope(internal.GenerateTestInstrumentationScope()), ms.Scope())
}

func TestScopeLogs_SchemaUrl(t *testing.T) {
	ms := NewScopeLogs()
	assert.Equal(t, "", ms.SchemaUrl())
	ms.SetSchemaUrl("https://opentelemetry.io/schemas/1.5.0")
	assert.Equal(t, "https://opentelemetry.io/schemas/1.5.0", ms.SchemaUrl())
}

func TestScopeLogs_LogRecords(t *testing.T) {
	ms := NewScopeLogs()
	assert.Equal(t, NewLogRecordSlice(), ms.LogRecords())
	internal.FillTestLogRecordSlice(internal.LogRecordSlice(ms.LogRecords()))
	assert.Equal(t, LogRecordSlice(internal.GenerateTestLogRecordSlice()), ms.LogRecords())
}

func TestLogRecordSlice(t *testing.T) {
	es := NewLogRecordSlice()
	assert.Equal(t, 0, es.Len())
	es = newLogRecordSlice(&[]*otlplogs.LogRecord{})
	assert.Equal(t, 0, es.Len())

	es.EnsureCapacity(7)
	emptyVal := newLogRecord(&otlplogs.LogRecord{})
	testVal := LogRecord(internal.GenerateTestLogRecord())
	assert.Equal(t, 7, cap(*es.getOrig()))
	for i := 0; i < es.Len(); i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, el)
		internal.FillTestLogRecord(internal.LogRecord(el))
		assert.Equal(t, testVal, el)
	}
}

func TestLogRecordSlice_CopyTo(t *testing.T) {
	dest := NewLogRecordSlice()
	// Test CopyTo to empty
	NewLogRecordSlice().CopyTo(dest)
	assert.Equal(t, NewLogRecordSlice(), dest)

	// Test CopyTo larger slice
	LogRecordSlice(internal.GenerateTestLogRecordSlice()).CopyTo(dest)
	assert.Equal(t, LogRecordSlice(internal.GenerateTestLogRecordSlice()), dest)

	// Test CopyTo same size slice
	LogRecordSlice(internal.GenerateTestLogRecordSlice()).CopyTo(dest)
	assert.Equal(t, LogRecordSlice(internal.GenerateTestLogRecordSlice()), dest)
}

func TestLogRecordSlice_EnsureCapacity(t *testing.T) {
	es := LogRecordSlice(internal.GenerateTestLogRecordSlice())
	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	expectedEs := make(map[*otlplogs.LogRecord]bool)
	for i := 0; i < es.Len(); i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, es.Len(), len(expectedEs))
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	foundEs := make(map[*otlplogs.LogRecord]bool, es.Len())
	for i := 0; i < es.Len(); i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	oldLen := es.Len()
	expectedEs = make(map[*otlplogs.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		expectedEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, oldLen, len(expectedEs))
	es.EnsureCapacity(ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.getOrig()))
	foundEs = make(map[*otlplogs.LogRecord]bool, oldLen)
	for i := 0; i < oldLen; i++ {
		foundEs[es.At(i).getOrig()] = true
	}
	assert.Equal(t, expectedEs, foundEs)
}

func TestLogRecordSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := LogRecordSlice(internal.GenerateTestLogRecordSlice())
	dest := NewLogRecordSlice()
	src := LogRecordSlice(internal.GenerateTestLogRecordSlice())
	src.MoveAndAppendTo(dest)
	assert.Equal(t, LogRecordSlice(internal.GenerateTestLogRecordSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, LogRecordSlice(internal.GenerateTestLogRecordSlice()), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	LogRecordSlice(internal.GenerateTestLogRecordSlice()).MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestLogRecordSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewLogRecordSlice()
	emptySlice.RemoveIf(func(el LogRecord) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := LogRecordSlice(internal.GenerateTestLogRecordSlice())
	pos := 0
	filtered.RemoveIf(func(el LogRecord) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestLogRecord_MoveTo(t *testing.T) {
	ms := LogRecord(internal.GenerateTestLogRecord())
	dest := NewLogRecord()
	ms.MoveTo(dest)
	assert.Equal(t, NewLogRecord(), ms)
	assert.Equal(t, LogRecord(internal.GenerateTestLogRecord()), dest)
}

func TestLogRecord_CopyTo(t *testing.T) {
	ms := NewLogRecord()
	orig := NewLogRecord()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = LogRecord(internal.GenerateTestLogRecord())
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
}

func TestLogRecord_ObservedTimestamp(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, pcommon.Timestamp(0), ms.ObservedTimestamp())
	testValObservedTimestamp := pcommon.Timestamp(1234567890)
	ms.SetObservedTimestamp(testValObservedTimestamp)
	assert.Equal(t, testValObservedTimestamp, ms.ObservedTimestamp())
}

func TestLogRecord_Timestamp(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, pcommon.Timestamp(0), ms.Timestamp())
	testValTimestamp := pcommon.Timestamp(1234567890)
	ms.SetTimestamp(testValTimestamp)
	assert.Equal(t, testValTimestamp, ms.Timestamp())
}

func TestLogRecord_TraceID(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, pcommon.TraceID(data.NewTraceID([16]byte{})), ms.TraceID())
	testValTraceID := pcommon.TraceID(data.NewTraceID([16]byte{1, 2, 3, 4, 5, 6, 7, 8, 8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetTraceID(testValTraceID)
	assert.Equal(t, testValTraceID, ms.TraceID())
}

func TestLogRecord_SpanID(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, pcommon.SpanID(data.NewSpanID([8]byte{})), ms.SpanID())
	testValSpanID := pcommon.SpanID(data.NewSpanID([8]byte{8, 7, 6, 5, 4, 3, 2, 1}))
	ms.SetSpanID(testValSpanID)
	assert.Equal(t, testValSpanID, ms.SpanID())
}

func TestLogRecord_FlagsStruct(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, LogRecordFlags(0), ms.FlagsStruct())
	testValFlagsStruct := LogRecordFlags(1)
	ms.SetFlagsStruct(testValFlagsStruct)
	assert.Equal(t, testValFlagsStruct, ms.FlagsStruct())
}

func TestLogRecord_SeverityText(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, "", ms.SeverityText())
	ms.SetSeverityText("INFO")
	assert.Equal(t, "INFO", ms.SeverityText())
}

func TestLogRecord_SeverityNumber(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, SeverityNumber(otlplogs.SeverityNumber(0)), ms.SeverityNumber())
	testValSeverityNumber := SeverityNumber(otlplogs.SeverityNumber(5))
	ms.SetSeverityNumber(testValSeverityNumber)
	assert.Equal(t, testValSeverityNumber, ms.SeverityNumber())
}

func TestLogRecord_Body(t *testing.T) {
	ms := NewLogRecord()
	internal.FillTestValue(internal.Value(ms.Body()))
	assert.Equal(t, pcommon.Value(internal.GenerateTestValue()), ms.Body())
}

func TestLogRecord_Attributes(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, pcommon.NewMap(), ms.Attributes())
	internal.FillTestMap(internal.Map(ms.Attributes()))
	assert.Equal(t, pcommon.Map(internal.GenerateTestMap()), ms.Attributes())
}

func TestLogRecord_DroppedAttributesCount(t *testing.T) {
	ms := NewLogRecord()
	assert.Equal(t, uint32(0), ms.DroppedAttributesCount())
	ms.SetDroppedAttributesCount(uint32(17))
	assert.Equal(t, uint32(17), ms.DroppedAttributesCount())
}