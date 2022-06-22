package hub

import "ws-noughts-and-crosses/pkg/vec"

// This code deals with the physics of the player, which is constant velocity
// motion in the direction of the user input

func toVel(inputEventMessage UserInputEventMessage) vec.Vec {
	var force vec.Vec
	if inputEventMessage.W {
		force[1] -= PlayerSpeed
	}
	if inputEventMessage.A {
		force[0] -= PlayerSpeed
	}
	if inputEventMessage.S {
		force[1] += PlayerSpeed
	}
	if inputEventMessage.D {
		force[0] += PlayerSpeed
	}
	return force
}

// Evolve evolves the player from the previous global tick to
// the time `tf`.
func (p *Player) Evolve(tf float64) {
	newPos := p.Pos
	for i, input := range p.InputStack.Inputs {
		tI := input.Timestamp
		vel := toVel(input)

		if i < len(p.InputStack.Inputs)-1 {
			tIp1 := p.InputStack.Inputs[i+1].Timestamp
			newPos = vec.Add(newPos, vec.Mul(vel, tIp1-tI))
		} else {
			newPos = vec.Add(newPos, vec.Mul(vel, tf-tI))
		}
	}

	lastInput := p.InputStack.Inputs[len(p.InputStack.Inputs)-1]
	lastInput.Timestamp = tf

	p.InputStack.Reset()
	p.InputStack.Push(lastInput)
	
	p.Pos = newPos
}
