package setDefault

func Str(arg *string, defaultValue string) {
	if *arg == "" {
		*arg = defaultValue
	}
}
func Int(arg *int, defaultValue int) {
	if *arg == 0 {
		*arg = defaultValue
	}
}
func Ints(arg *[]int, defaultValue []int) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
func Strs(arg *[]string, defaultValue []string) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
func MapStrStr(arg *map[string]string, defaultValue map[string]string) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
func MapStrInt(arg *map[string]int, defaultValue map[string]int) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
func MapIntStr(arg *map[int]string, defaultValue map[int]string) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
func MapIntInt(arg *map[int]int, defaultValue map[int]int) {
	if len(*arg) == 0 {
		*arg = defaultValue
	}
}
