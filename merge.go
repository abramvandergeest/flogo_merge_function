package merge

import (

	// "time"

	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnMerge{})
}

type fnMerge struct {
}

func (fnMerge) Name() string {
	return "merge"
}

func (fnMerge) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray, data.TypeString}, false
}

func (fnMerge) Eval(params ...interface{}) (interface{}, error) {
	maps := params[0].([]interface{})
	mergeKey := params[1].(string)
	output := make(map[string]interface{})
	// fmt.Println(maps)
	// fmt.Println(mergeKey)

	// indval := maps[0].(map[string]interface{})[mergeKey]
	// fmt.Println(indval)

	var indval interface{}

	for i, m := range maps {
		if m == nil {
			continue
		}
		if indval == nil {
			indval = m.(map[string]interface{})[mergeKey]
		}
		// fmt.Println(m)
		// fmt.Println(m.(map[string]interface{})[mergeKey])
		if m.(map[string]interface{})[mergeKey] != indval {
			return nil, fmt.Errorf("there exists multiple indexs, currently we are assuming one index")
		}
		for key, val := range m.(map[string]interface{}) {
			if key != mergeKey || i == 0 {
				output[key] = val
			}
		}
	}

	return output, nil
}
