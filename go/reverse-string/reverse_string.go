package reverse

func Reverse(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}
