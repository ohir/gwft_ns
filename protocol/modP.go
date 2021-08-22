package protocol

const me string = "/mP Protocol" // common module for others

func Say() string {
	return me[3:] + " (" + Vstamp + ") " +
		" said: tinker here."
}
