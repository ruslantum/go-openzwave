package openzwave

// #cgo LDFLAGS: -lopenzwave -Lgo/src/github.com/ninjasphere/go-openzwave/openzwave
// #cgo CPPFLAGS: -Iopenzwave/cpp/src/platform -Iopenzwave/cpp/src -Iopenzwave/cpp/src/value_classes
//
// #include "api.h"
import "C"

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave/CODE"
	"github.com/ninjasphere/go-openzwave/NT"
	"github.com/ninjasphere/go-openzwave/VT"
)

// The type of notifications received via the API's Notifications() channel.
type Notification struct {
	inC *C.Notification
}

// Converts the notification into a string representation.
func (self Notification) String() string {
	return fmt.Sprintf(
		"Notification["+
			"node=0x%08x:0x%02x, "+
			"notificationType=%s/%s, "+
			"valueType=%s, "+
			"valueId=0x%08x]",
		self.inC.nodeId.homeId,
		self.inC.nodeId.nodeId,
		NT.ToEnum(int(self.inC.notificationType)),
		CODE.ToEnum(int(self.inC.notificationCode)),
		VT.ToEnum(int(self.inC.valueId.valueType)),
		self.inC.valueId.id)
}

func (apiNotification *Notification) Free() {
	C.freeNotification(apiNotification.inC)
}

func (notification *Notification) GetValueID() *ValueID {
	return &ValueID{notification.inC.valueId};
}

func (notification *Notification) GetValue() *Value {
	return &Value{notification.inC.value};
}

