package kv

import (
	"github.com/1071496910/kv-server/util"
	"github.com/cespare/xxhash"
)

type KV interface {
	Set(k string) string
	Get(k string) string
	Storage() string
}

type kvNode struct {
	key  string
	val  string
	next *kvNode
}

type kvList struct {
	head *kvNode
	tail *kvNode
	len  int
}

func newKvList() *kvList {
	tmp := &kvNode{}
	return &kvList{
		head: tmp,
		tail: tmp,
		len:  0,
	}
}

func (kl *kvList) append(node *kvNode) *kvList {
	if node == nil {
		return kl
	}

	kl.tail.next = node
	kl.tail = node
	kl.len++
	return kl
}

func (kl *kvList) Add(node *kvNode) *kvList {
	if node == nil {
		return kl
	}
	for current := kl.head.next; current != nil; current = current.next {
		if current.key == node.key {
			current.val = node.val
			return kl
		}
	}
	return kl.append(node)
}

func (kl *kvList) Get(k string) string {
	for current := kl.head.next; current != nil; current = current.next {
		if current.key == k {
			return current.val
		}
	}
	return ""
}

func (kl *kvList) Delete(k string) *kvList {
	if k == "" {
		return kl
	}
	for prev, current := kl.head, kl.head.next; current != nil; prev, current = current, current.next {
		if current.key == k {
			prev.next = current.next
			if current == kl.tail {
				kl.tail = prev
			}
		}
	}
	return kl
}

type hashListKV struct {
	hashList   []*kvList
	expectSize int
}

func newHashListKV(size int) *hashListKV {
	h := &hashListKV{}
	h.expectSize = size
	h.hashList = make([]*kvList, size)
	for i := 0; i < h.expectSize; i++ {
		h.hashList[i] = newKvList()
	}
	return h
}

func (h *hashListKV) Set(k string, v string) string {
	index := xxhash.Sum64(util.Bytes(k)) % uint64(h.expectSize)
	h.hashList[index].Add(&kvNode{
		key:  k,
		val:  v,
		next: nil,
	})
	return v
}

func (h hashListKV) Get(k string) string {
	index := xxhash.Sum64(util.Bytes(k)) % uint64(h.expectSize)
	return h.hashList[index].Get(k)
}

func (h hashListKV) Storage() string {
	panic("implement me")
}
