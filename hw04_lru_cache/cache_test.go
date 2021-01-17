package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

type testNumber struct {
	key           Key
	expectedValue interface{}
	expectedOk    bool
}

func itoKey(i int) Key {
	return Key(strconv.Itoa(i))
}

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("LRU wiki example", func(t *testing.T) {
		//Пример из https://ru.bmstu.wiki/LRU_(Least_Recently_Used)
		c := NewCache(3)
		numberList := []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5}
		for _, number := range numberList {
			c.Set(itoKey(number), number)
		}

		for _, tst := range [...]testNumber{
			{
				key:           "1",
				expectedValue: nil,
				expectedOk:    false,
			},
			{
				key:           "2",
				expectedValue: nil,
				expectedOk:    false,
			},
			{
				key:           "3",
				expectedValue: 3,
				expectedOk:    true,
			},
			{
				key:           "4",
				expectedValue: 4,
				expectedOk:    true,
			},
			{
				key:           "5",
				expectedValue: 5,
				expectedOk:    true,
			},
		} {
			val, ok := c.Get(tst.key)
			require.Equal(t, tst.expectedValue, val)
			require.Equal(t, tst.expectedOk, ok)
		}
	})

	/*
		в задании сказано тест "на логику выталкивания редкоиспользуемых элементов"
		но LRU это не про частоту, а про последовательность использования
		тест ниже показывает это
	*/
	t.Run("NONE often/rarely logic", func(t *testing.T) {
		c := NewCache(3)
		numberList := []int{
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 10
			2, 2, 2, 2, 2, 2, 2, 2, 2, 2, // 20
			1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 30
			3, 4, 5,
		}
		for _, number := range numberList {
			c.Set(itoKey(number), number)
		}

		for _, tst := range [...]testNumber{
			{
				key:           "1",
				expectedValue: nil,
				expectedOk:    false, // хоть это и самый частоиспользуемый элемент
			},
			{
				key:           "2",
				expectedValue: nil,
				expectedOk:    false,
			},
			{
				key:           "3",
				expectedValue: 3,
				expectedOk:    true,
			},
			{
				key:           "4",
				expectedValue: 4,
				expectedOk:    true,
			},
			{
				key:           "5",
				expectedValue: 5,
				expectedOk:    true,
			},
		} {
			val, ok := c.Get(tst.key)
			require.Equal(t, tst.expectedValue, val)
			require.Equal(t, tst.expectedOk, ok)
		}
	})

	//тест на выталкивание самого давно использованного элемента
	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)

		numberListSet := []int{1, 2, 3}
		for _, number := range numberListSet {
			c.Set(itoKey(number), number)
		}

		numberListGet := []int{3, 1, 1, 1, 3, 3, 1}
		for _, number := range numberListGet {
			c.Get(itoKey(number))
		}

		wasInCache := c.Set("new Key", "new Value")
		require.False(t, wasInCache)

		val, ok := c.Get("2")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
