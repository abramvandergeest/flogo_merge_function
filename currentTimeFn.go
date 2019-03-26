package currentTimeFn

import (
	"time"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnCurrentTime{})
}

type fnCurrentTime struct {
}

func (fnCurrentTime) Name() string {
	return "currentTimeUnix"
}

func (fnCurrentTime) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt64}, false
}

func (fnCurrentTime) Eval(params ...interface{}) (interface{}, error) {

	return time.Now().Unix(), nil
}
