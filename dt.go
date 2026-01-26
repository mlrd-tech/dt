// Copyright 2026 mlrd.tech, Inc.
// http://www.apache.org/licenses/LICENSE-2.0

package dt

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func Map(v ...any) map[string]types.AttributeValue {
	if len(v)%2 != 0 {
		panic("dt.Map: odd number of inputs, must be even set of key-value pairs")
	}
	m := map[string]types.AttributeValue{}
	for i := 0; i < len(v); i += 2 {
		key, ok := v[i].(string)
		if !ok {
			panic(fmt.Sprintf("dt.Map: v[%d] not a string: %v", i, v[i]))
		}
		val, ok := v[i+1].(types.AttributeValue)
		if !ok {
			panic(fmt.Sprintf("dt.Map: v[%d] not a types.AttributeValue: %v", i, v[i]))
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

func L(l ...types.AttributeValue) types.AttributeValue {
	return &types.AttributeValueMemberL{Value: l}
}

func M(m map[string]types.AttributeValue) types.AttributeValue {
	return &types.AttributeValueMemberM{Value: m}
}
