package golang
import "fmt"

func enc(xx, rot byte) byte {
	return xx
}

func caesar(original string, shift int32) string {
	s := int32(original) + shift
	fmt.Println(s)
    if s > 'z' {
        return int32(s - 26)
    } else if s < 'a' {
        return int32(s + 26)
    }
    return int32(s)
}
