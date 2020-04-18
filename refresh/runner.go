package refresh

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func (m *Manager) runner() {
	var cmd *exec.Cmd
	for {
		select {
		case <-m.Restart:
		case <-m.context.Done():
			return
		}
		if cmd != nil {
			// kill the previous command
			pid := cmd.Process.Pid
			m.Logger.Success("Stopping: PID %d", pid)
			cmd.Process.Kill()
		}
		if m.Debug {
			bp := m.FullBuildPath()
			args := []string{"exec", bp}
			args = append(args, m.CommandFlags...)
			cmd = exec.Command("dlv", args...)
		} else {
			cmd = exec.Command(m.FullBuildPath(), m.CommandFlags...)
		}
		go func() {
			err := m.runAndListen(cmd)
			if err != nil {
				m.Logger.Error(err)
			}
		}()
	}
}

func (m *Manager) runAndListen(cmd *exec.Cmd) error {
	cmd.Stderr = m.Stderr
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}

	cmd.Stdin = m.Stdin
	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	cmd.Stdout = m.Stdout
	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}

	var stderrBuf bytes.Buffer

	errR, errW := io.Pipe()
	go func() {
		for {
			buf := make([]byte, 32*1024)
			n, err := errR.Read(buf)
			if err != nil {
				if err != io.EOF {
					m.Logger.Error("Error reading stderr: %s.", err)
				}
				return
			}
			buf = buf[:n]
			_, _ = stderrBuf.Write(buf) // Best effort.
			m.Logger.Error("%s\n", buf)
		}
	}()

	defer func() { _ = errW.Close() }() // Best effort.
	cmd.Stderr = io.MultiWriter(errW, cmd.Stderr)

	// Set the environment variables from config
	if len(m.CommandEnv) != 0 {
		cmd.Env = append(m.CommandEnv, os.Environ()...)
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderrBuf.String())
	}

	m.Logger.Print("Running: %s (PID: %d)", strings.Join(cmd.Args, " "), cmd.Process.Pid)
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("%s\n%s", err, stderrBuf.String())
	}
	return nil
}
