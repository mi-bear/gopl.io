package join

func Join(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}
