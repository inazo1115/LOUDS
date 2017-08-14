package louds

import (
	"reflect"

	"github.com/inazo1115/bitarray"
	"github.com/inazo1115/encode"
	"github.com/inazo1115/fid"
	lane "gopkg.in/oleiade/lane.v1"
)

type TreeNode interface {
	Val() interface{}
	Parent() TreeNode
	FirstChild() TreeNode
	NextBrother() TreeNode
}

type LOUDS struct {
	values []interface{}
	fid_   *fid.FID
}

func BuildLOUDS(node TreeNode) *LOUDS {

	// Breadth-first search
	values := make([]interface{}, 0)
	nChildren := make([]int, 1)
	nChildren[0] = 1 // guard node
	queue := lane.NewQueue()
	queue.Enqueue(node)
	for {
		if queue.Size() == 0 {
			break
		}
		node_ := queue.Dequeue().(TreeNode)
		values = append(values, node_.Val())
		if node_.FirstChild() == nil {
			nChildren = append(nChildren, 0)
			continue
		}
		x := 0
		for n := node_.FirstChild(); !reflect.ValueOf(n).IsNil(); n = n.NextBrother() {
			queue.Enqueue(n)
			x++
		}
		nChildren = append(nChildren, x)
	}

	// Encode
	strings_ := make([]string, 0)
	for _, v := range nChildren {
		if v == 0 {
			strings_ = append(strings_, "1")
		} else {
			strings_ = append(strings_, "0"+encode.ToUnary(v))
		}
	}
	bools := make([]bool, 0)
	for _, v := range strings_ {
		for i := 0; i < len(v); i++ {
			if string(v[i]) == "0" {
				bools = append(bools, false)
			} else {
				bools = append(bools, true)
			}
		}
	}
	fid_ := fid.NewFID(bitarray.NewBitArrayWithInit(bools))

	return &LOUDS{values: values, fid_: fid_}
}

func (louds *LOUDS) Size() int {
	ret, err := louds.fid_.Rank(true, louds.fid_.Bits().Size()-1)
	if err != nil {
		panic("logic bug: " + err.Error())
	}
	return ret
}

func (louds *LOUDS) Val(idx int) interface{} {
	return louds.values[idx]
}

func (louds *LOUDS) Parent(i int) int {
	idx, err := louds.fid_.Select(false, i)
	if err != nil {
		panic(err.Error())
	}
	idx, err = louds.fid_.Rank(true, idx)
	if err != nil {
		panic(err.Error())
	}
	return idx - 1
}

func (louds *LOUDS) FirstChild(i int) int {
	idx, err := louds.fid_.Select(true, i)
	if err != nil {
		panic(err.Error())
	}
	x, err := louds.fid_.Access(idx + 1)
	if err != nil {
		panic(err.Error())
	}
	if x == true {
		return -1
	}
	idx, err = louds.fid_.Rank(false, idx+1)
	if err != nil {
		panic(err.Error())
	}
	return idx
}

func (louds *LOUDS) NextBrother(i int) int {
	idx, err := louds.fid_.Select(false, i)
	if err != nil {
		panic(err.Error())
	}
	x, err := louds.fid_.Access(idx + 1)
	if err != nil {
		panic(err.Error())
	}
	if x == true {
		return -1
	}
	idx, err = louds.fid_.Rank(false, idx+1)
	if err != nil {
		panic(err.Error())
	}
	return idx
}
