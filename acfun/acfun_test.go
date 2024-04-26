package acfun

import (
	"github.com/starskim/DDBOT-Plugins/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAcfun(t *testing.T) {
	assert.NotEmpty(t, APath(PathApiChannelList))
	assert.NotEmpty(t, APath("api/channel/list"))
	assert.NotEmpty(t, LiveUrl(test.UID1))
}
