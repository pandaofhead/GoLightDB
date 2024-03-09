package storage

import "testing"

// TestDBMeta_Store 测试dbMeta的持久化
func TestDBMeta_Store(t *testing.T) {
	m := &DBMeta{43}
	if err := m.Store("/Users/roseduan/resources/rosedb/db.Meta"); err != nil {
		t.Error(err)
	}
}

func TestLoad(t *testing.T) {
	m := LoadMeta("/Users/roseduan/resources/rosedb/db.Meta")
	t.Logf("%+v \n", m)
	t.Log(m == nil)
}
