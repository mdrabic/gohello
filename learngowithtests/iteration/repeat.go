package iteration

func Repeat(s string, n int) string {
	var o string
	for i := 0; i < n; i++ {
		o += s
	}
	return o
}
