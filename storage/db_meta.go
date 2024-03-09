package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// DBMeta 保存数据库的一些额外信息
type DBMeta struct {
	ActiveWriteOff int64 `json:"active_write_off"` //当前数据文件的写偏移
}

// LoadMeta 加载数据信息
func LoadMeta(path string) (m *DBMeta) {
	m = &DBMeta{}

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return
	}
	defer file.Close()
	//io/ioutil.ReadAll() 从文件中读取数据
	b, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(b, m)
	return
}

// Store 存储数据信息
func (m *DBMeta) Store(path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	b, _ := json.Marshal(m)
	_, err = file.Write(b)
	return err
}
