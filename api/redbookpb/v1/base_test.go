package basev1

import (
	"fmt"
	"testing"
)

type Acursor struct {
	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func TestCursor(t *testing.T) {
	offset := int32(0)
	limit := int32(10)
	cursor := &Cursor{
		Offset: offset,
		Limit:  limit,
	}
	fmt.Printf("cursor:%+v\n", cursor)
	acursor := &Acursor{
		Offset: offset,
		Limit:  limit,
	}
	fmt.Printf("acursor:%+v\n", acursor)
}
