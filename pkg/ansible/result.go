package ansible

type Result struct {
	Plays []Play
	Stats []map[string]Stat
}

type Play struct {
	Duration Duration
	Id    string
	Name  string
}

type Duration struct {
	Start string
	End   string

}

type Task struct {

}

type Host struct {

}

type Stat struct {
	Change      int
	Failures    int
	Ok          int
	Skipped     int
	Unreachable int
}
