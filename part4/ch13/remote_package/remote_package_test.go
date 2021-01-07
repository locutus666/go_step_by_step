package remote

// 远程包导入
import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

// 远程包调用
func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
