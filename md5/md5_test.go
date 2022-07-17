package md5

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMd5Sum(t *testing.T) {
	tmp, _ := os.CreateTemp("", "")
	tmp.WriteString("sdfsdfsf\n")
	defer os.Remove(tmp.Name())
	md5 := MD5File(tmp.Name())
	assert.Equal(t, "06baf2ce6d63a21002e1d017c13f850b", md5)
}
