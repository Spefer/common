package hash

const (
	HASH_START      = int64(17)
	HASH_PULTIPLIER = int64(37)
)

func ToHashNum(str string) int64 {
	var start = HASH_START
	for _, v := range str {
		start = start + int64(v) * HASH_PULTIPLIER
	}
	return start
}
