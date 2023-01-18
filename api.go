// package squidr your squid infosec companion
package squidr

import (
	"bufio"
	"os"
)

// Listen starts the squidr instance
func Listen() {
	go debugLog()
	go servReports()
	go servPreview()
	go func() {
		for line := range outChan {
			// debugLogChan <- _ioDebugOUT + line
			os.Stdout.Write([]byte(line + _linefeed))
		}
	}()
	r := bufio.NewScanner(os.Stdin)
	for r.Scan() {
		go action(r.Text())
	}
}
