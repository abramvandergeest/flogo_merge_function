package merge

import (

	// "time"

	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnRandom{})
}

type fnRandom struct {
}

func (fnRandom) Name() string {
	return "merge"
}

func (fnRandom) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt}, true
}

func (fnRandom) Eval(params ...interface{}) (interface{}, error) {
	var outmap map[string][]interface{}
	var indexset map[interface{}][]int
	maps := params[0].([]map[string]interface{})
	mergeKey := params[1].([]string)
	keymap := map[string]int{}

	for i, m := range maps {
		k := mergeKey[i]

		if len(m[k].([]interface{})) > 1 {
			for j, _ := range m[k].([]interface{}) {
				indexset[m[k]][i] = j
			}
		} else {
			indexset[m[k]][i] = 0
		}

		for key := range m {
			if _, ok := keymap[key]; !ok && key != mergeKey[i] {
				keymap[key] = i
			}
			// else{    //IF THE KEY IS A DUPLICATE
			// 	keymap[]
			// }
		}
		blah := m[k]
		fmt.Println(blah)
	}

	fmt.Println(mergeKey, maps, outmap, indexset)
	// limit := 10
	// if len(params) > 0 {
	// 	limit = params[0].(int)
	// }
	// rand.Seed(time.Now().UnixNano())
	return 0, nil
}
