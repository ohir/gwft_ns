package server

import (
	proto "example.com/mP"
)

const me string = "/mS Server"

func Say() string {
	return me[3:] + " (" + Vstamp + ") " +
		" said: tinker here.\n\tThen: " +
		proto.Say()
}
