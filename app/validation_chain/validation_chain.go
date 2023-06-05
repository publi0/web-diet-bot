package validation_chain

import "web-diet-bot/dto"

type Action interface {
	Execute(event *dto.Event)
	SetNext(action Action)
}
