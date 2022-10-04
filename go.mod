module github.com/subdotnet/avalancheseqlogger

go 1.18

// Try to use the same versions of go.uber.org packages as 
// https://github.com/ava-labs/avalanchego/blob/master/go.mod

require go.uber.org/zap v1.21.0

require (
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
)
