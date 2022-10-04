// Sends log messages to seq by posting log messages as Json - https://docs.datalust.co/docs/posting-raw-events
package avalancheseqlogger

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"go.uber.org/zap/zapcore"
)

const (
	SeqDefaultUrl = "http://localhost:18080/api/events/raw/?clef"
)

// Creates the encoder and enriches the message with :
// - commandLine : this program's command line and arguments
// - nodeNumber : when using avalanche-network-runner - extracted from command line arguments
// - app : "avalancheGo"
func NewSeqEncoder(levelEncoder zapcore.LevelEncoder) zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "@t",
		NameKey:        "_nameKey",
		LevelKey:       "@l",
		CallerKey:      "_callerKey",
		MessageKey:     "@m",
		StacktraceKey:  "@x",
		FunctionKey:    "_functionKey",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeName:     zapcore.FullNameEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    levelEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)

	commandLine := fmt.Sprint(os.Args)
	encoder.AddString("commandLine", commandLine)

	compiledRegexp, err := regexp.Compile("node[0-9]+")
	var nodeNumber string
	if err != nil {
		nodeNumber = "Error in regexp.Compile : " + err.Error()
	} else {
		nodeNumber = compiledRegexp.FindString(commandLine)
	}
	encoder.AddString("@nodeNumber", nodeNumber)
	encoder.AddString("@app", "avalancheGo")

	return encoder
}

// Creates the writer to post logs to the Seq api
func NewSeqWriter(url string) (_ io.WriteCloser, err error) {
	var w = &SeqWriter{
		url: url,
	}
	return w, nil
}

type SeqWriter struct {
	io.Writer
	url string
}

// Nothing to close
func (*SeqWriter) Close() error {
	return nil
}

// Posts the message to the Seq api.
// Authentication is not implemented.
func (w *SeqWriter) Write(buf []byte) (int, error) {
	n := len(buf)
	req, err := http.NewRequest("POST", w.url, bytes.NewReader(buf))
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		txtResp, err := io.ReadAll(resp.Body)
		if err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("error creating seq event: %v - Input : %v", string(txtResp), string(buf))
	}
	return n, nil
}
