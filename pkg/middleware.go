package pkg

func Validate(from, to int) string {
	switch {
	case from == 0:
		return "The start value should be bigger than 0\n"
	case to == 0:
		return "The end value should be bigger than 0\n"
	case from < 0:
		return "The start value should be bigger than 0\n"
	case to < 0:
		return "The end value should be bigger than 0\n"
	case from > to:
		return "The end value of the interval is bigger than the start value\n"
	case from > 2147483647:
		return "The biggest value for from is 2 147 483 647\n"
	case to > 2147483647:
		return "The biggest value for to is 2 147 483 647\n"
	}
	return ""
}

//type JSONValue struct {
//	number int
//	value  *big.Int
//}
//type JSONValues []JSONValue
//
//func Prepare(response *lib.Response, from int) JSONValues {
//	var table JSONValues
//
//	for _, value := range response.GetList() {
//		k := value.GetValue()
//		fmt.Println(k)
//		jnew := JSONValue{
//			number: int(value.Id),
//			value:  new(big.Int).SetBytes(value.Value),
//		}
//		fmt.Println(jnew)
//		table = append(table, jnew)
//	}
//	return table
//}
