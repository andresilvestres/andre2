package andre

import (
	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	_ = function.Register(&fnandreLeft{})
}

type fnandreLeft struct {
}

func (fnandreLeft) Name() string {
	return "andre"
}

func (fnandreLeft) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString}, true
}

func (fnandreLeft) Eval(params ...interface{}) (interface{}, error) {

	//andre Left
	//fmt.Println("'%4dkm'", params)
	fmt.Printf("'%4dkm'", params)
	return params, fmt.Errorf("andreleft")

	//if params == nil {
		//Do nothing
		//return 0, nil
	//}
	//return "...ERROR...", fmt.Errorf("fnandreLeft function must have three arguments")
}
