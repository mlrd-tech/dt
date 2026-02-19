// Copyright 2026 mlrd.tech, Inc.
// http://www.apache.org/licenses/LICENSE-2.0

package dt

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func M(v ...any) map[string]types.AttributeValue {
	if len(v)%2 != 0 {
		panic("dt.M: odd number of inputs, must be even set of key-value pairs")
	}
	m := map[string]types.AttributeValue{}
	for i := 0; i < len(v); i += 2 {
		var key string
		if s, ok := v[i].(string); ok {
			key = s
		} else if s, ok := v[i].(fmt.Stringer); ok {
			key = s.String()
		} else {
			panic(fmt.Sprintf("dt.M: v[%d] not a string or fmt.Stringer: %v (type: %T)", i, v[i], v[i]))
		}
		val, ok := v[i+1].(types.AttributeValue)
		if !ok {
			panic(fmt.Sprintf("dt.M: v[%d] not a types.AttributeValue: %v (type: %T)", i+1, v[i+1], v[i+1]))
		}
		m[key] = val
	}
	return m
}

func S(s string) types.AttributeValue {
	return &types.AttributeValueMemberS{Value: s}
}

// Number is a type constraint for numeric types and strings.
// Since DynamoDB stores numbers as strings, this constraint
// allows strings, too, so N("1") works as well as N(1).
type Number interface {
	~string |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func N[T Number](n T) types.AttributeValue {
	return &types.AttributeValueMemberN{Value: fmt.Sprint(n)}
}

func B(b []byte) types.AttributeValue {
	return &types.AttributeValueMemberB{Value: b}
}

func BOOL(b bool) types.AttributeValue {
	return &types.AttributeValueMemberBOOL{Value: b}
}

func NULL() types.AttributeValue {
	return &types.AttributeValueMemberNULL{Value: true}
}

func SS(ss ...string) types.AttributeValue {
	return &types.AttributeValueMemberSS{Value: ss}
}

func NS(ns ...string) types.AttributeValue {
	return &types.AttributeValueMemberNS{Value: ns}
}

func BS(bs ...[]byte) types.AttributeValue {
	return &types.AttributeValueMemberBS{Value: bs}
}

func L(l ...any) types.AttributeValue {
	vals := make([]types.AttributeValue, len(l))
	for i, v := range l {
		av, ok := v.(types.AttributeValue)
		if !ok {
			panic(fmt.Sprintf("dt.L: l[%d] not a types.AttributeValue: %v", i, v))
		}
		vals[i] = av
	}
	return &types.AttributeValueMemberL{Value: vals}
}
