package engine

import (
	"DoctorStrange/errorcode"
	"sync"
)

type EngineManager struct {
	engines map[string]*engine
}

var defaultManager *EngineManager
var once sync.Once

func DefaultManager()(*EngineManager){
	once.Do(func() {
		defaultManager =  &EngineManager{
			engines:make(map[string]*engine,100),
		}
	})

	return defaultManager
}

func (m *EngineManager)Start ()  {
	//默认情况下，启动 btc/usdt 引擎
	m.AddEngine("btc/usdt", 10001)
}

//释放所有engine
func (m *EngineManager)Stop ()  {
	for _, e := range m.engines{
		e.Stop()
	}
}

//engine 可插拔
func (m *EngineManager)AddEngine (symbol string, price float32) errorcode.ErrorCode  {

	if m.EngineExist(symbol) == errorcode.ErrorCodeEngineExist {
		return errorcode.ErrorCodeEngineExist
	}

	e := New(symbol,price)
	m.engines[symbol] = e
	e.Start()

	return errorcode.OK
}

func (m *EngineManager)RemoveEngine (symbol string) errorcode.ErrorCode {

	if m.EngineExist(symbol) == errorcode.ErrorCodeEngineNotExist {
		return errorcode.ErrorCodeEngineNotExist
	}

	e := m.engines[symbol]
	e.Stop()

	return errorcode.OK
}

func (m *EngineManager)EngineExist(symbol string) errorcode.ErrorCode {
	if len(symbol) == 0 {
		return errorcode.ErrorCodeInvalid
	}

	e := m.engines[symbol]
	if e == nil {
		return errorcode.ErrorCodeEngineNotExist
	}

	return errorcode.ErrorCodeEngineExist
}

func (m *EngineManager)Engine (symbol string) (*engine, errorcode.ErrorCode) {
	if len(symbol) == 0 {
		return nil, errorcode.ErrorCodeInvalid
	}

	e := m.engines[symbol]
	if e == nil {
		return nil, errorcode.ErrorCodeEngineNotExist
	}

	return e, errorcode.OK
}