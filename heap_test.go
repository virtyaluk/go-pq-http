package main

import "testing"

func TestMinHeap(t *testing.T) {
	mh := NewMinHeap()

	cases := []struct {
		key int
		val string
	}{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
	}

	for _, item := range cases {
		mh.Push(item.key, item.val)
	}

	for _, item := range cases {
		topKey, topVal := mh.Pop()

		if topKey != item.key || topVal != item.val {
			t.Errorf("expected min heap top to be (%d, %s), got (%d, %s).", item.key, item.val, topKey, topVal)
		}
	}

	if !mh.Empty() {
		t.Errorf("expected min heap to be empty")
	}
}
