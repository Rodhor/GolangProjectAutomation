package types

// State is an enumeration that represents the current view/state.
type State int

const (
	projectName State = iota
	languageSelection
	packageSelection
	packageSubDir
	SummaryAndConfirmation
)
