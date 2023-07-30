package server

import "miopkg/governor"

func (eng *Engine) NewServerGovernor() error {
	server := governor.StdConfig("governor").Build()
	return eng.Serve(server)
}
