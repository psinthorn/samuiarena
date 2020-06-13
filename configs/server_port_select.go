package configs

import "os"

type server struct {}

var Server server

func (s *server) PortRunning(port string) string {
	portEnv := os.Getenv("PORT")
	if portEnv == "" {
		portEnv = port 
	}
	return portEnv
}