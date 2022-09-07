package shell

import (
	"bytes"
	"chains-gotest-backend/api/plt/config"
	"chains-gotest-backend/api/plt/pkg/log"
	"io"
	"os"
	"os/exec"

)

func Exec(filepath string, args ...string) {
	var errStdout, errStderr error

	shellpath := config.ShellPath(filepath)
	cmd := exec.Command(shellpath, args...)
	//cmd.Env = addEnv()

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	stdout := NewCapturingPassThroughWriter(os.Stdout)
	stderr := NewCapturingPassThroughWriter(os.Stderr)

	if err := cmd.Start(); err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	// waiting for execution end
	{
		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
		}()
		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
		}()
		if err := cmd.Wait(); err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
	}

	// capture and print result
	if isError(errStderr) || isError(errStdout) {
		log.Fatalf("failed to capture stdout or stderr\n")
	}

	//outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	//log.Infof("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

//const (
//	EnvWorkspace      = "PaletteWorkspace"
//	EnvNodeIndexStart = "PaletteNodeIndexStart"
//	EnvNodeNumber     = "PaletteNodeNumber"
//	EnvNodeIndexEnd   = "PaletteNodeIndexEnd"
//	EnvNetworkID      = "PaletteNetworkID"
//	EnvStartRPCPort   = "PaletteStartRPCPort"
//	EnvStartP2PPort   = "PaletteStartP2PPort"
//	EnvLogLevel       = "PaletteLogLevel"
//)
//
//func addEnv() []string {
//	env := config.Conf.Environment
//	list := make([]string, 0)
//	var add = func(env string, value interface{}) {
//		list = append(list, fmt.Sprintf("%s=%v", env, value))
//	}
//
//	add(EnvWorkspace, env.NodeDir)
//	add(EnvNodeIndexStart, env.NodeIdxStart)
//	if env.NodeNum == 0 {
//		env.NodeNum = config.GenesisNodeNumber()
//	}
//	add(EnvNodeNumber, env.NodeNum)
//	add(EnvNodeIndexEnd, int(env.NodeIdxStart+env.NodeNum-1))
//	add(EnvNetworkID, env.NetworkID)
//	add(EnvStartRPCPort, env.StartRPCPort)
//	add(EnvStartP2PPort, env.StartP2PPort)
//	add(EnvLogLevel, env.LogLevel)
//
//	log.Infof("start env: workspace %s, nodeIndexStart %d, nodeNum %d, networkID %d, startRPCPort %d, startP2PPort %d, logLevel %d",
//		env.NodeDir, env.NodeIdxStart, env.NodeNum, env.NetworkID, env.StartRPCPort, env.StartP2PPort, env.LogLevel)
//
//	return append(os.Environ(), list...)
//}

// CapturingPassThroughWriter is a writer that remembers
// data written to it and passes it to w
type CapturingPassThroughWriter struct {
	buf bytes.Buffer
	w   io.Writer
}

// NewCapturingPassThroughWriter creates new CapturingPassThroughWriter
func NewCapturingPassThroughWriter(w io.Writer) *CapturingPassThroughWriter {
	return &CapturingPassThroughWriter{
		w: w,
	}
}

func (w *CapturingPassThroughWriter) Write(d []byte) (int, error) {
	w.buf.Write(d)
	return w.w.Write(d)
}

// Bytes returns bytes written to the writer
func (w *CapturingPassThroughWriter) Bytes() []byte {
	return w.buf.Bytes()
}

func isError(err error) bool {
	if err != nil && err == os.ErrClosed {
		return true
	}
	return false
}
