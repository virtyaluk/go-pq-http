package main

import "sync"

type node struct {
	key int
	val interface{}
}

type minHeap struct {
	heapList []*node
	curSize  int
	mutex    sync.Mutex
}

func NewMinHeap() *minHeap {
	emptyNode := &node{}

	return &minHeap{
		heapList: []*node{emptyNode},
		curSize:  0,
	}
}

func (mh *minHeap) siftUp(i int) {
	for i/2 > 0 {
		if mh.heapList[i].key < mh.heapList[i/2].key {
			mh.heapList[i], mh.heapList[i/2] = mh.heapList[i/2], mh.heapList[i]
		}

		i /= 2
	}
}

func (mh *minHeap) Size() int {
	mh.mutex.Lock()
	defer mh.mutex.Unlock()

	return mh.curSize
}

func (mh *minHeap) Empty() bool {
	mh.mutex.Lock()
	defer mh.mutex.Unlock()

	return mh.Size() == 0
}

func (mh *minHeap) Push(k int, v interface{}) {
	mh.mutex.Lock()
	defer mh.mutex.Unlock()

	mh.heapList = append(mh.heapList, &node{k, v})
	mh.curSize++

	mh.siftUp(mh.curSize)
}

func (mh *minHeap) minChild(i int) int {
	if (i*2)+1 > mh.curSize {
		return i * 2
	}

	if mh.heapList[i*2].key < mh.heapList[i*2+1].key {
		return i * 2
	}

	return i*2 + 1
}

func (mh *minHeap) siftDown(i int) {
	for (i * 2) <= mh.curSize {
		mc := mh.minChild(i)

		if mh.heapList[i].key > mh.heapList[mc].key {
			mh.heapList[i], mh.heapList[mc] = mh.heapList[mc], mh.heapList[i]
		}

		i = mc
	}
}

func (mh *minHeap) Top() (int, interface{}) {
	mh.mutex.Lock()
	defer mh.mutex.Unlock()

	top := mh.heapList[1]

	return top.key, top.val
}

func (mh *minHeap) Pop() (int, interface{}) {
	mh.mutex.Lock()
	defer mh.mutex.Unlock()

	root := mh.heapList[1]

	mh.heapList[1] = mh.heapList[mh.curSize]
	mh.heapList = mh.heapList[:len(mh.heapList)-1]
	mh.curSize--

	mh.siftDown(1)

	return root.key, root.val
}
