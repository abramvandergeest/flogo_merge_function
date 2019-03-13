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
	maps := params[0].([]map[string]interface{})
	mergeKey := params[1].([]string)
	var output map[string]interface{}

	indval := maps[0][mergeKey[0]]

	for i, m := range maps {
		if mergeKey[i] != indval {
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
