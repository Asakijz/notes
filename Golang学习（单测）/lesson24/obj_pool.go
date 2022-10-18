package lesson24

import (
	"errors"
	"time"
)

type ReusableObj struct {
}
type ObjPool struct {
	bufChan chan *ReusableObj
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (o *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	// 如果想要任意对象可以使用 interface{}，不过这样需要在返回时进行类型断言
	select {
	case ret := <-o.bufChan:
		return ret, nil
	case <-time.After(timeout): // 超时控制
		return nil, errors.New("time out")
	}
}

func (o *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case o.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
