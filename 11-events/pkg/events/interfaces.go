package events

import (
	"sync"
	"time"
)

// Evento
type EventInterface interface {
	GetName() string         // nome do evento
	GetDateTime() time.Time  // momento que o evento foi disparado
	GetPayload() interface{} // conteúdo do evento
}

// Operações que serão executadas quando um evento é chamado
type EventHandlerInterface interface {
	Handle(event EventInterface, wg *sync.WaitGroup) // método que executa a operação a partir do evento
}

// Gerenciador de eventos/operações
type EventDespatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error // Registra o evento
	Dispatch(event EventInterface) error                            //Cara que vai fazer com que os handlers sejam executados, ou seja as operações a partir dos eventos sejam executadas
	Remove(eventName string, handler EventHandlerInterface) error
	Has(eventName string, handler EventHandlerInterface) bool
	Clear() error // exclui todos os eventos que foram registrados
}
