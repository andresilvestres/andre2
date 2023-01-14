package andre

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnandreRight{})
}

type fnandreRight struct {
}

func (fnandreRight) Name() string {
	return "andre"
}

func (fnandreRight) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnandreRight) Eval(params ...interface{}) (interface{}, error) {

	//andreRight
	fmt.Printf("'%-4s'", params)

	return params, fmt.Errorf("andreRight")

}
