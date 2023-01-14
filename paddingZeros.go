package andre

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnandreZeros{})
}

type fnandreZeros struct {
}

func (fnandreZeros) Name() string {
	return "andre"
}

func (fnandreZeros) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnandreZeros) Eval(params ...interface{}) (interface{}, error) {

	//andre with zeroes
	fmt.Printf("'%04dkm'", params)

	return params, fmt.Errorf("andreWithZeroes")
}
