package utils

type SliceType struct {
	SliceInterface []interface{}
	LoopRet        map[string]int
}

func (st *SliceType) LoopAndCount() {
	var (
		ifMap = make(map[string]int)
	)
	for _, elem := range st.SliceInterface {
		switch elem.(type) {
		//fmt.Println("int:", elemTyped)
		case int:
			//fmt.Println("int:", elemTyped)
			ifMap["int"] += 1
		case string:
			//fmt.Println("string:", elemTyped)
			ifMap["string"] += 1
		case []string:
			//fmt.Println("[]string:", elemTyped)
			ifMap["sliceStr"] += 1
		case interface{}:
			//fmt.Println("interface:", elemTyped)
			ifMap["interface"] += 1
		}
	}
	st.LoopRet = ifMap
}
