package hash

import "testing"

var key = "my_hash"

// InitHash 初始化一个hash
func InitHash() *Hash {
	hash := New()

	hash.HSet(key, "a", []byte("hash_data_001"))
	hash.HSet(key, "b", []byte("hash_data_002"))
	hash.HSet(key, "c", []byte("hash_data_003"))

	return hash
}

func TestHash_HSet(t *testing.T) {
	hash := InitHash()

	n := hash.HSet("my_hash", "d", []byte("123"))
	t.Log(n)

	n = hash.HSet("my_hash", "e", []byte("234"))
	t.Log(n)
}

func TestHash_HSetNx(t *testing.T) {
	hash := InitHash()

	t.Log(hash.HSetNx(key, "a", []byte("new one")))
	t.Log(hash.HSetNx(key, "d", []byte("d-new one")))

	t.Log(hash.HLen(key))
}

func TestHash_HGet(t *testing.T) {
	hash := InitHash()

	val := hash.HGet(key, "a")
	t.Log(string(val))

	t.Log(string(hash.HGet(key, "c")))
	t.Log(string(hash.HGet(key, "m")))
}

func TestHash_HGetAll(t *testing.T) {

	hash := InitHash()

	vals := hash.HGetAll(key)
	for _, v := range vals {
		t.Log(string(v))
	}
}

func TestHash_HDel(t *testing.T) {
	hash := InitHash()

	r := hash.HDel(key, "a")
	t.Log(r)
	r = hash.HDel(key, "c")
	t.Log(r)
}

func TestHash_HExists(t *testing.T) {
	hash := InitHash()

	t.Log(hash.HExists(key, "a"))
	t.Log(hash.HExists(key, "c"))
	t.Log(hash.HExists(key, "s"))
}

func TestHash_HKeys(t *testing.T) {
	hash := InitHash()

	keys := hash.HKeys(key)
	for _, k := range keys {
		t.Log(k)
	}

	res := hash.HKeys("no")
	t.Log(len(res))
}

func TestHash_HValues(t *testing.T) {
	hash := InitHash()

	values := hash.HValues(key)
	for _, v := range values {
		t.Log(string(v))
	}
}
