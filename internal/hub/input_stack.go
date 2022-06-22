package hub

import (
	"sync"
	"time"
)

// InputStack is a stack of user input events, received since
// the last global tick
type InputStack struct {
	Lock sync.Mutex

	Inputs []UserInputEventMessage
}

func (i *InputStack) Push(input UserInputEventMessage) {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	input.Timestamp = float64(time.Now().UnixMilli())
	i.Inputs = append(i.Inputs, input)
}

func (i *InputStack) Reset() {
	i.Lock.Lock()
	defer i.Lock.Unlock()

	i.Inputs = nil
}
