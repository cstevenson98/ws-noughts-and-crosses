package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func toForce(inputEventMessage UserInputEventMessage) [2]float64 {
	var force [2]float64
	if inputEventMessage.W {
		force[1] -= 1
	}
	if inputEventMessage.A {
		force[0] -= 1
	}
	if inputEventMessage.S {
		force[1] += 1
	}
	if inputEventMessage.D {
		force[0] += 1
	}
	return force
}

type Player struct {
	IPlayer
	InputStack InputStack

	Pos [2]float64
	Vel [2]float64

	Hub    *Hub
	Game   *Game
	Conn   *websocket.Conn
	Stream chan []byte
}

func (p *Player) Update(force [2]float64, dt float64) {
	p.Pos[0] += p.Vel[0] * dt
	p.Pos[1] += p.Vel[1] * dt
	p.Vel[0] += force[0] * PlayerAcceleration * dt
	p.Vel[1] += force[1] * PlayerAcceleration * dt
}

func (p *Player) Evolve(tf float64) {
	t := p.Game.t0

	for _, input := range p.InputStack.Inputs {
		dt := input.Timestamp - t
		if dt < 0 {
			// Disregard inputs that are in the past
			continue
		}

		// Split dt into sub-dt increments
		subTime := t
		for subTime+float64(subDt) < input.Timestamp {
			subTime += subDt
			p.Update(toForce(input), subDt/float64(time.Millisecond*1000))
		}

	}

}

func (p *Player) ReadPump() {
	defer func() {
		p.Hub.Unregister <- p
		p.Conn.Close()
	}()

	for {
		_, msg, err := p.Conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var input UserInputEventMessage
		err = json.Unmarshal(msg, &input)
		if err != nil {
			log.Println("could not unmarshal user input")
		}

		p.InputStack.Push(input)
	}
}

func (p *Player) WritePump() {
	defer func() {
		p.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-p.Stream:
			if !ok {
				// The hub closed the channel.
				p.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := p.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func (p *Player) ProcessTurn(turn []byte) error {

	return nil
}
