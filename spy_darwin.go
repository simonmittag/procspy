package procspy

import (
	"os/exec"
)

const (
	netstatBinary = "netstat"
	lsofBinary    = "lsof"
)

// Connections returns all established (TCP) connections. No need to be root
// to run this. If processes is true it also tries to fill in the process
// fields of the connection. You need to be root to find all processes.
var cbConnections = func(processes bool) (ConnIter, error) {
	out, err := exec.Command(
		netstatBinary,
		"-n", // no number resolving
		"-W", // Wide output
		// "-l", // full IPv6 addresses // What does this do?
		"-p", "tcp", // only TCP
		"-v", // with pid
	).CombinedOutput()
	if err != nil {
		// log.Printf("lsof error: %s", err)
		return nil, err
	}
	connections := parseDarwinNetstat(string(out))

	f := fixedConnIter(connections)
	return &f, nil
}
