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
	return []data.Type{data.TypeArray, data.TypeArray}, false
}

func (fnMerge) Eval(params ...interface{}) (interface{}, error) {
	maps := params[0].([]map[string]interface{})
	mergeKey := params[1].([]string)
	var output map[string]interface{}

	indval := maps[0][mergeKey[0]]

	for i, m := range maps {
		if m[mergeKey[i]] != indval {
			return nil, fmt.Errorf("there exists multiple indexs, currently we are assuming one index")
		}
		for key, val := range m {
			if key != mergeKey[i] || i == 0 {
				output[key] = val
			}
		}
	}

	return output, nil
}
