# dt ~ Short DynamoDB Types in Go

This:

```go
dt.L(dt.S("yes"), dt.N(42), dt.BOOL(true))   
```

Rather than:

```go
&types.AttributeValueMemberL{      
    Value: []types.AttributeValue{    
        &types.AttributeValueMemberS{Value: "wtf"},    
        &types.AttributeValueMemberN{Value: strconv.Itoa(42)},    
        &types.AttributeValueMemberBOOL{Value: true},    
    }
}
```

The `dt` API:

| `dt` Function | Rather Than |                               
|---|---|                                                   
| `dt.S("yes")` | `&types.AttributeValueMemberS{Value: "ugh"}` |           
| `dt.N(100)` | `&types.AttributeValueMemberN{Value: "0"}` |        
| `dt.B([]byte("yay"))` | `&types.AttributeValueMemberB{Value: []byte{...}}` |
| `dt.BOOL(true)` | `&types.AttributeValueMemberBOOL{Value: true}` |              
| `dt.NULL()` | `&types.AttributeValueMemberNULL{Value: true}` |                  
| `dt.SS("o", "m", "g")` | `&types.AttributeValueMemberSS{Value: []string{...}}` |           
| `dt.NS("1", "2", "3")` | `&types.AttributeValueMemberNS{Value: []string{...}}` |                     
| `dt.BS([]byte("x"), []byte("y"))` | `&types.AttributeValueMemberBS{Value: [][]byte{...}}` |                 
| `dt.L(dt.S("wheee"), dt.N(7))` | `&types.AttributeValueMemberL{Value: []types.AttributeValue{...}}` |                                
| `dt.M(":k", dt.S("easy"))` | `map[string]types.AttributeValue{":k": &types.AttributeValueMemberS{...}}` |
