package jsonpainter

func String(s string, opts ...Option) string {
	out := []byte(nil)
	fsm := newFiniteStateMachine(opts...)
	for _, c := range []byte(s) {
		out = append(out, fsm.Next(c)...)
	}
	out = append(out, fsm.Finish()...)
	return string(out)
}
