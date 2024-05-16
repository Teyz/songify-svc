package constants

type RoundStatus string

const (
	RoundStatusNotStarted RoundStatus = "not_started"
	RoundStatusStarted    RoundStatus = "started"
	RoundStatusFinished   RoundStatus = "finished"
)
