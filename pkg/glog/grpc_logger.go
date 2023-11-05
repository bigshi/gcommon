/**
 * Create Time:2023/11/5
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package glog

import (
	"google.golang.org/grpc/grpclog"
	"os"
)

type grpcLogger struct {
}

func init() {
	grpclog.SetLoggerV2(&grpcLogger{})
}

// V reports whether verbosity level l is at least the requested verbose level.
func (logger *grpcLogger) V(l int) bool {
	return bool(V(Level(l)))
}

// Info logs to the INFO log.
func (logger *grpcLogger) Info(args ...any) {
	Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func (logger *grpcLogger) Infof(format string, args ...any) {
	Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func (logger *grpcLogger) Infoln(args ...any) {
	Infoln(args...)
}

// Warning logs to the WARNING log.
func (logger *grpcLogger) Warning(args ...any) {
	Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func (logger *grpcLogger) Warningf(format string, args ...any) {
	Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func (logger *grpcLogger) Warningln(args ...any) {
	Warningln(args...)
}

// Error logs to the ERROR log.
func (logger *grpcLogger) Error(args ...any) {
	Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func (logger *grpcLogger) Errorf(format string, args ...any) {
	Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func (logger *grpcLogger) Errorln(args ...any) {
	Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func (logger *grpcLogger) Fatal(args ...any) {
	Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func (logger *grpcLogger) Fatalf(format string, args ...any) {
	Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func (logger *grpcLogger) Fatalln(args ...any) {
	Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func (logger *grpcLogger) Print(args ...any) {
	Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func (logger *grpcLogger) Printf(format string, args ...any) {
	Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func (logger *grpcLogger) Println(args ...any) {
	Infoln(args...)
}
