/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package log

import (
	"github.com/shieldworks/aegis/core/env"
	"log"
	"sync"
)

type Level int

const Off Level = 1
const Error Level = 2
const Warn Level = 3
const Info Level = 4
const Debug Level = 5
const Trace Level = 6

var currentLevel = Level(env.LogLevel())
var mux sync.Mutex

func SetLevel(l Level) {
	mux.Lock()
	defer mux.Unlock()
	if l < Off || l > Trace {
		return
	}
	currentLevel = l
}

func GetLevel() Level {
	mux.Lock()
	defer mux.Unlock()
	return currentLevel
}

func FatalLn(correlationId *string, v ...any) {
	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Fatalln(args...)
}

func ErrorLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Error {
		return
	}

	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

func WarnLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Warn {
		return
	}

	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

func InfoLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Info {
		return
	}

	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

func DebugLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Debug {
		return
	}

	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}

func TraceLn(correlationId *string, v ...any) {
	l := GetLevel()
	if l < Trace {
		return
	}

	var args []any
	args = append(args, *correlationId)
	args = append(args, v...)
	log.Println(args...)
}
