// Copyright 2026 mlrd.tech, Inc.
// http://www.apache.org/licenses/LICENSE-2.0

package dt_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/go-test/deep"

	"github.com/mlrd-tech/dt"
)

func TestM(t *testing.T) {
	testcases := []struct {
		in  []any
		out map[string]types.AttributeValue
	}{
		{
			in: []any{":s", dt.S("hello")},
			out: map[string]types.AttributeValue{
				":s": &types.AttributeValueMemberS{Value: "hello"},
			},
		},
		{
			in: []any{
				":s", dt.S("hello"),
				":n", dt.N("42"),
				":b", dt.BOOL(true),
			},
			out: map[string]types.AttributeValue{
				":s": &types.AttributeValueMemberS{Value: "hello"},
				":n": &types.AttributeValueMemberN{Value: "42"},
				":b": &types.AttributeValueMemberBOOL{Value: true},
			},
		},
	}

	for _, tc := range testcases {
		got := dt.M(tc.in...)
		if diff := deep.Equal(got, tc.out); diff != nil {
			t.Error(diff)
		}
	}
}

func TestS(t *testing.T) {
	got := dt.S("test")
	want := &types.AttributeValueMemberS{Value: "test"}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestN(t *testing.T) {
	testcases := []struct {
		name string
		in   any
		out  types.AttributeValue
	}{
		{
			name: "string",
			in:   "123",
			out:  &types.AttributeValueMemberN{Value: "123"},
		},
		{
			name: "int",
			in:   42,
			out:  &types.AttributeValueMemberN{Value: "42"},
		},
		{
			name: "int64",
			in:   int64(999),
			out:  &types.AttributeValueMemberN{Value: "999"},
		},
		{
			name: "int32",
			in:   int32(100),
			out:  &types.AttributeValueMemberN{Value: "100"},
		},
		{
			name: "uint",
			in:   uint(50),
			out:  &types.AttributeValueMemberN{Value: "50"},
		},
		{
			name: "uint64",
			in:   uint64(1000),
			out:  &types.AttributeValueMemberN{Value: "1000"},
		},
		{
			name: "float64",
			in:   45.67,
			out:  &types.AttributeValueMemberN{Value: "45.67"},
		},
		{
			name: "float32",
			in:   float32(12.34),
			out:  &types.AttributeValueMemberN{Value: "12.34"},
		},
		{
			name: "negative int",
			in:   -100,
			out:  &types.AttributeValueMemberN{Value: "-100"},
		},
		{
			name: "zero",
			in:   0,
			out:  &types.AttributeValueMemberN{Value: "0"},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var got types.AttributeValue
			switch v := tc.in.(type) {
			case string:
				got = dt.N(v)
			case int:
				got = dt.N(v)
			case int32:
				got = dt.N(v)
			case int64:
				got = dt.N(v)
			case uint:
				got = dt.N(v)
			case uint64:
				got = dt.N(v)
			case float32:
				got = dt.N(v)
			case float64:
				got = dt.N(v)
			default:
				t.Fatalf("unsupported type: %T", v)
			}
			if diff := deep.Equal(got, tc.out); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestB(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	got := dt.B(data)
	want := &types.AttributeValueMemberB{Value: data}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestBOOL(t *testing.T) {
	testcases := []struct {
		in  bool
		out types.AttributeValue
	}{
		{
			in:  true,
			out: &types.AttributeValueMemberBOOL{Value: true},
		},
		{
			in:  false,
			out: &types.AttributeValueMemberBOOL{Value: false},
		},
	}

	for _, tc := range testcases {
		got := dt.BOOL(tc.in)
		if diff := deep.Equal(got, tc.out); diff != nil {
			t.Error(diff)
		}
	}
}

func TestNULL(t *testing.T) {
	got := dt.NULL()
	want := &types.AttributeValueMemberNULL{Value: true}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestSS(t *testing.T) {
	got := dt.SS("a", "b", "c")
	want := &types.AttributeValueMemberSS{Value: []string{"a", "b", "c"}}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestNS(t *testing.T) {
	got := dt.NS("1", "2", "3")
	want := &types.AttributeValueMemberNS{Value: []string{"1", "2", "3"}}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestBS(t *testing.T) {
	data1 := []byte{0x01, 0x02}
	data2 := []byte{0x03, 0x04}
	got := dt.BS(data1, data2)
	want := &types.AttributeValueMemberBS{Value: [][]byte{data1, data2}}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

func TestL(t *testing.T) {
	got := dt.L(dt.S("hello"), dt.N("42"), dt.BOOL(true))
	want := &types.AttributeValueMemberL{
		Value: []types.AttributeValue{
			&types.AttributeValueMemberS{Value: "hello"},
			&types.AttributeValueMemberN{Value: "42"},
			&types.AttributeValueMemberBOOL{Value: true},
		},
	}
	if diff := deep.Equal(got, want); diff != nil {
		t.Error(diff)
	}
}

