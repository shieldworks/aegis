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

func GetLevel(v ...any) Level {
	mux.Lock()
	defer mux.Unlock()
	return currentLevel
}

func FatalLn(v ...any) {
	log.Fatalln(v...)
}

func ErrorLn(v ...any) {
	l := GetLevel()
	if l < Error {
		return
	}
	log.Println(v...)
}

func WarnLn(v ...any) {
	l := GetLevel()
	if l < Warn {
		return
	}
	log.Println(v...)
}

func InfoLn(v ...any) {
	l := GetLevel()
	if l < Info {
		return
	}
	log.Println(v...)
}

func DebugLn(v ...any) {
	l := GetLevel()
	if l < Debug {
		return
	}
	log.Println(v...)
}

func TraceLn(v ...any) {
	l := GetLevel()
	if l < Trace {
		return
	}
	log.Println(v...)
}
