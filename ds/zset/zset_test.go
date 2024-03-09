package zset

import (
	"math/rand"
	"testing"
	"time"
)

// TestZSet_ZAdd 向有序集合中添加元素
func InitZSet() *SortedSet {
	zSet := New()
	zSet.ZAdd("myzset", 19, "ced")
	zSet.ZAdd("myzset", 12, "acd")
	zSet.ZAdd("myzset", 17, "bcd")
	zSet.ZAdd("myzset", 32, "acc")
	zSet.ZAdd("myzset", 17, "mcd")
	zSet.ZAdd("myzset", 21, "ccd")
	zSet.ZAdd("myzset", 17, "ecd")

	return zSet
}

func TestZSet_ZAdd(t *testing.T) {

	t.Run("normal data", func(t *testing.T) {
		zSet := InitZSet()
		zSet.ZAdd("myzset", 39, "mmd")

		t.Log(zSet.ZCard("myzset"))
		t.Log(zSet.ZScore("myzset", "ced"))
		t.Log(zSet.ZScore("myzset", "mmd"))
	})

	t.Run("large data", func(t *testing.T) {
		zset := New()
		rand.Seed(time.Now().Unix())

		s := "abcdefghijklmnopqrstuvwxyz"
		randomVal := func() (val string) {
			for i := 0; i < 12; i++ {
				val += string(s[rand.Intn(26)])
			}
			return
		}

		for i := 0; i < 100000; i++ {
			zset.ZAdd("myzset", float64(rand.Intn(100000)), randomVal())
		}

		t.Log(zset.ZCard("myzset"))
		dummy := zset.record["myzset"].skl.head
		p := dummy.level[0].forward
		for i := 0; i < 100; i++ {
			t.Log(p.member, p.score)
			p = p.level[0].forward
		}
	})
}

func TestZSet_ZScore(t *testing.T) {
	zSet := InitZSet()
	t.Log(zSet.ZScore("myzset", "acd"))
	t.Log(zSet.ZScore("myzset", "ccd"))
	t.Log(zSet.ZScore("myzset", "ccd"))
	t.Log(zSet.ZScore("myzset", "accsssss"))
}

func TestZSet_ZRank(t *testing.T) {

	key := "myzset"
	zset := InitZSet()
	rank := zset.ZRank(key, "acd")
	t.Log(rank)

	t.Log(zset.ZRank(key, "acc"))
	t.Log(zset.ZRank(key, "mcd"))
	t.Log(zset.ZRank(key, "ecd"))
	t.Log(zset.ZRank(key, "bcd"))
}

func TestZSet_ZRevRank(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	rank := zset.ZRevRank(key, "acc")
	t.Log(rank)

	t.Log(zset.ZRevRank(key, "ccd"))
	t.Log(zset.ZRevRank(key, "acd"))

	t.Log(zset.ZRevRank(key, "bcd"))
	t.Log(zset.ZRevRank(key, "mcd"))
	t.Log(zset.ZRevRank(key, "ecd"))
}

func TestZSet_ZIncrBy(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	incrBy := zset.ZIncrBy(key, 300, "acd")
	t.Log(incrBy)

	t.Log(zset.ZIncrBy(key, 100, "acc"))

	t.Log(zset.ZRank(key, "acd"))
	t.Log(zset.ZRank(key, "acc"))
}

func TestZSet_ZRange(t *testing.T) {
	zSet := InitZSet()
	key := "myzset"

	ran := zSet.ZRange(key, 0, -1)
	t.Log(len(ran))

	for _, v := range ran {
		t.Logf("%+v", v)
	}
}

func TestZSet_ZRevRange(t *testing.T) {
	zSet := InitZSet()
	key := "myzset"

	ran := zSet.ZRevRange(key, 0, -1)
	t.Log(len(ran))

	for _, v := range ran {
		t.Logf("%+v", v)
	}
}

func TestZSet_ZRem(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	ok := zset.ZRem(key, "acd")

	t.Log(ok)
	t.Log(zset.ZRem(key, "aaaaaaa"))
	t.Log(zset.ZCard(key))
}

func TestSortedSet_ZGetByRank(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	val := zset.ZGetByRank(key, 2)
	if val != nil {
		for _, v := range val {
			t.Logf("%+v ", v)
		}
	}
}

func TestSortedSet_ZRevGetByRank(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	rand.Seed(time.Now().Unix())
	s := "abcdefghijklmnopqrstuvwxyz"
	randomVal := func() (val string) {
		for i := 0; i < 12; i++ {
			val += string(s[rand.Intn(26)])
		}
		return
	}

	for i := 0; i < 100000; i++ {
		zset.ZAdd("myzset", float64(rand.Intn(100000)), randomVal())
	}

	val := zset.ZGetByRank(key, 0)
	if val != nil {
		for _, v := range val {
			t.Logf("%+v ", v)
		}
	}

	dummy := zset.record["myzset"].skl.head
	p := dummy.level[0].forward
	for i := 0; i < 100; i++ {
		t.Log(p.member, p.score)
		p = p.level[0].forward
	}
}

func TestSortedSet_ZScoreRange(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	zset.ZAdd(key, 13, "aa")

	val := zset.ZScoreRange(key, -12, 500)
	for _, v := range val {
		t.Logf("%+v", v)
	}
}

func TestSortedSet_ZRevScoreRange(t *testing.T) {
	zset := InitZSet()
	key := "myzset"

	t.Run("normal", func(t *testing.T) {
		zset.ZAdd(key, 45, "aa")

		val := zset.ZRevScoreRange(key, 17, 17)
		for _, v := range val {
			t.Logf("%+v", v)
		}
	})

	t.Run("large data", func(t *testing.T) {
		rand.Seed(time.Now().Unix())
		s := "abcdefghijklmnopqrstuvwxyz"
		randomVal := func() (val string) {
			for i := 0; i < 12; i++ {
				val += string(s[rand.Intn(26)])
			}
			return
		}

		start := time.Now()
		for i := 0; i < 600000; i++ {
			zset.ZAdd("myzset", float64(rand.Intn(600000)), randomVal())
		}
		t.Log("add time spend ", time.Since(start))

		start = time.Now()
		val := zset.ZRevScoreRange(key, 359980, 359090)
		t.Log("query time spend ", time.Since(start))

		t.Log(len(val))
		for _, v := range val {
			t.Logf("%+v", v)
		}
	})
}
