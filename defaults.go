package progressbar

// spinners built-in spinners for progress bar
var spinners = map[int][]string{
	0: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	1: {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	2: {"▁", "▂", "▃", "▄", "▅", "▆", "▇"},
	3: {"*--", "-*-", "--*"},
	4: {".  ", ".. ", "...", " ..", "  ."},
	5: {"*  ", "** ", "***", " **", "  *"},
	6: {"*  ", " * ", "  *"},
}

func GetSpinners() map[int][]string {
	return spinners
}

// fillers built-in fillers for progress bar
var fillers = map[int][2]string{
	0: {"█", "░"},
	1: {"|", " "},
	2: {"#", " "},
	3: {"▰", "▱"},
	5: {"■", " "},
	6: {"￭", "･"},
	7: {"▮", "▯"},
	8: {"=", "·"},
}

func GetFillers() map[int][2]string {
	return fillers
}

// edges built-in edges for progress bar
var edges = map[int][2]string{
	0: {"[", "]"},
	1: {"{", "}"},
	2: {"(", ")"},
	3: {"[", "]"},
	4: {"⌈", "⌉"},
	5: {"⌊", "⌋"},
	6: {"｢", "｣"},
	7: {"❬", "❭"},
	8: {"❰", "❱"},
}

func GetEdges() map[int][2]string {
	return edges
}
