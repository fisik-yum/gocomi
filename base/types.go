package base

type Result struct {
	Name string
	ID   string
}

type DLdata struct {
	ID    string //this can be any string that is unique to the comic that is being downloaded
	Start string //if date, should be compatible with the time package, preferably time.UnixDate?
	Stop  string //same requirements as Stop
}
