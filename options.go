package jsonpainter

type (
	Option func(fsm *finiteStateMachine)
	Color  []byte
)

var (
	Yellow    = Color("\033[33;1m")
	Brown     = Color("\033[33m")
	Red       = Color("\033[31;1m")
	Darkred   = Color("\033[31m")
	Pink      = Color("\033[35;1m")
	Darkpink  = Color("\033[35m")
	Blue      = Color("\033[34;1m")
	Darkblue  = Color("\033[34m")
	Green     = Color("\033[32;1m")
	Darkgreen = Color("\033[32m")
	Cyan      = Color("\033[36;1m")
	Darkcyan  = Color("\033[36m")
	White     = Color("\033[37;1m")
	Black     = Color("\033[30m")
	Lightgray = Color("\033[37m")
	Darkgray  = Color("\033[30;1m")
	None      = Color(nil)
	Off       = Color("\033[0m")
)

func ClrKey(clr Color) Option {
	return func(fsm *finiteStateMachine) {
		fsm.clrKey = clr
	}
}

func ClrSpecStr(clr Color) Option {
	return func(fsm *finiteStateMachine) {
		fsm.clrSpecStr = clr
	}
}

func ClrStr(clr Color) Option {
	return func(fsm *finiteStateMachine) {
		fsm.clrStr = clr
	}
}

func ClrCtl(clr Color) Option {
	return func(fsm *finiteStateMachine) {
		fsm.clrCtl = clr
	}
}

func ClrOff(clr Color) Option {
	return func(fsm *finiteStateMachine) {
		fsm.clrOff = clr
	}
}
