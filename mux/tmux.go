package mux

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type Tmux struct{}

func NewTmux() *Tmux {
	return &Tmux{}
}

func (t *Tmux) Name() string {
	return "tmux"
}

// Returns the equivilant of TERM, TERM_PROGRAM hiding behind the MUX
// TERM, TERM_PROGRAM, error
func (t *Tmux) Probe() (string, string, error) {
	tmux, err := exec.LookPath("tmux")
	if err != nil {
		return "", "", fmt.Errorf("failed to locate tmux in path: %v", err)
	}
	cmd := exec.Command(tmux, "show-environment")

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	out, errout := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		return "", "", fmt.Errorf("failed to probe tmux: %v, %s", err, errout)
	}

	var term, program string = "", ""
	scanner := bufio.NewScanner(strings.NewReader(out))
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "=") {
			continue
		}
		kv := strings.Split(scanner.Text(), "=")
		switch kv[0] {
		case "TERM":
			term = kv[1]
		case "TERM_PROGRAM":
			program = kv[1]
		}
		if term != "" && program != "" {
			break
		}
	}
	if err = scanner.Err(); err != nil {
		return term, program, fmt.Errorf("failed to probe tmux while scanning tmux env: %v", err)
	}
	return term, program, nil
}
