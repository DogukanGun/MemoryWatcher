package gmail

import (
	"MemoryWatcher/utils"
	"testing"
)

func TestSendingEmail(t *testing.T) {
	utils.Send("Teest")
}
