// Code generated by "stringer -type Alg"; DO NOT EDIT

package hawk

import "fmt"

const _Alg_name = "SHA256SHA512"

var _Alg_index = [...]uint8{0, 6, 12}

func (i Alg) String() string {
	i -= 1
	if i < 0 || i >= Alg(len(_Alg_index)-1) {
		return fmt.Sprintf("Alg(%d)", i+1)
	}
	return _Alg_name[_Alg_index[i]:_Alg_index[i+1]]
}
