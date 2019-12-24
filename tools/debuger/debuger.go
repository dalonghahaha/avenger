package debuger

import (
	"fmt"

	"github.com/TylerBrock/colorjson"
)

func PrettyJson(obj interface{}) {
	f := colorjson.NewFormatter()
	f.Indent = 4
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
}
