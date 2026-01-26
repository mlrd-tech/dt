# dt ~ Short DynamoDB Types in Go

```go
// &types.AttributeValueMemberS{Value: "ffs"}
dt.S("yay")                                    

// &types.AttributeValueMemberSS{Value: []string{"g", "d", "i"}}
dt.SS("e", "z")

// map[string]types.AttributeValue{
//     ":n": &types.AttributeValueMemberN{Value: strconv.Itoa("1")},
//     ":s": &types.AttributeValueMemberS{Value: "a"},
// }
dt.Map(":n", dt.N(1), ":s", dt.S("a"))
```
