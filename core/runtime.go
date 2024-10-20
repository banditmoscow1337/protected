package core

import (
	"os/exec"
	"strconv"
)

func (s *Server) run() error {
	cmd := exec.Command("java", "-Xmx"+strconv.FormatInt(s.memory.max, 10)+"M", "-Xms"+strconv.FormatInt(s.memory.min, 10)+"M", "-jar", "server.jar", "nogui")
	return cmd.Run()
}
