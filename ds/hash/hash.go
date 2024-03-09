package hash

// hash实现

type (
	// Hash 哈希表结构定义
	Hash struct {
		record Record
	}

	// Record hash record to save
	Record map[string]map[string][]byte
)

// New new a hash ds
func New() *Hash {
	return &Hash{make(Record)}
}

// HSet 将哈希表 hash 中域 field 的值设置为 value
// 如果给定的哈希表并不存在， 那么一个新的哈希表将被创建并执行 HSet 操作
// 如果域 field 已经存在于哈希表中， 那么它的旧值将被新值 value 覆盖
func (h *Hash) HSet(key string, field string, value []byte) int {
	if !h.exist(key) {
		h.record[key] = make(map[string][]byte)
	}

	h.record[key][field] = value
	return len(h.record[key])
}

// HSetNx 当且仅当域 field 尚未存在于哈希表的情况下， 将它的值设置为 value
// 如果给定域已经存在于哈希表当中， 那么命令将放弃执行设置操作
func (h *Hash) HSetNx(key string, field string, value []byte) bool {
	if !h.exist(key) {
		h.record[key] = make(map[string][]byte)
	}

	if _, exist := h.record[key][field]; !exist {
		h.record[key][field] = value
		return true
	}

	return false
}

// HGet 返回哈希表中给定域的值
func (h *Hash) HGet(key, field string) []byte {
	if !h.exist(key) {
		return nil
	}

	return h.record[key][field]
}

// HGetAll 返回哈希表 key 中，所有的域和值
func (h *Hash) HGetAll(key string) (res [][]byte) {
	if !h.exist(key) {
		return
	}

	for k, v := range h.record[key] {
		res = append(res, []byte(k), v)
	}

	return
}

// HDel 删除哈希表 key 中的一个或多个指定域，不存在的域将被忽略
// 返回是否被成功移除
func (h *Hash) HDel(key, field string) bool {
	if !h.exist(key) {
		return false
	}

	if _, exist := h.record[key][field]; exist {
		delete(h.record[key], field)
		return true
	}

	return false
}

// HExists 检查给定域 field 是否存在于key对应的哈希表中
func (h *Hash) HExists(key, field string) bool {
	if !h.exist(key) {
		return false
	}

	_, exist := h.record[key][field]
	return exist
}

// HLen 返回哈希表 key 中域的数量
func (h *Hash) HLen(key string) int {
	if !h.exist(key) {
		return 0
	}

	return len(h.record[key])
}

// HKeys 返回哈希表 key 中的所有域
func (h *Hash) HKeys(key string) (val []string) {
	if !h.exist(key) {
		return
	}

	for k := range h.record[key] {
		val = append(val, k)
	}

	return
}

// HValues 返回哈希表 key 中的所有域对应的值
func (h *Hash) HValues(key string) (val [][]byte) {

	if !h.exist(key) {
		return
	}

	for _, v := range h.record[key] {
		val = append(val, v)
	}

	return
}

func (h *Hash) exist(key string) bool {
	_, exist := h.record[key]
	return exist
}
