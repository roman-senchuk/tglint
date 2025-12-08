package rules

// Issue represents a linting issue
type Issue struct {
	File    string
	Line    int
	Column  int
	Message string
	Rule    string
}
