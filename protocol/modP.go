package protocol

const me string = "/mP Protocol" // common module for others

func Say() string {
	return me[3:] +
		" said: tinker MORE here."
}
