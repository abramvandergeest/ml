// Code generated by protoc-gen-go. DO NOT EDIT.
// source: input.proto

/*
Package tensorflow_serving is a generated protocol buffer package.

It is generated from these files:
	input.proto

It has these top-level messages:
	ExampleList
	ExampleListWithContext
	Input
*/
package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow1 "github.com/project-flogo/contrib/activity/inference/tensorflow/tensorflow/core/example"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Specifies one or more fully independent input Examples.
// See examples at:
//     https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/example/example.proto
type ExampleList struct {
	Examples []*tensorflow1.Example `protobuf:"bytes,1,rep,name=examples" json:"examples,omitempty"`
}

func (m *ExampleList) Reset()                    { *m = ExampleList{} }
func (m *ExampleList) String() string            { return proto.CompactTextString(m) }
func (*ExampleList) ProtoMessage()               {}
func (*ExampleList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExampleList) GetExamples() []*tensorflow1.Example {
	if m != nil {
		return m.Examples
	}
	return nil
}

// Specifies one or more independent input Examples, with a common context
// Example.
//
// The common use case for context is to cleanly and optimally specify some
// features that are common across multiple examples.
//
// See example below with a search query as the context and multiple restaurants
// to perform some inference on.
//
// context: {
//   feature: {
//     key  : "query"
//     value: {
//       bytes_list: {
//         value: [ "pizza" ]
//       }
//     }
//   }
// }
// examples: {
//   feature: {
//     key  : "cuisine"
//     value: {
//       bytes_list: {
//         value: [ "Pizzeria" ]
//       }
//     }
//   }
// }
// examples: {
//   feature: {
//     key  : "cuisine"
//     value: {
//       bytes_list: {
//         value: [ "Taqueria" ]
//       }
//     }
//   }
// }
//
// Implementations of ExampleListWithContext merge the context Example into each
// of the Examples. Note that feature keys must not be duplicated between the
// Examples and context Example, or the behavior is undefined.
//
// See also:
//     tensorflow/core/example/example.proto
//     https://developers.google.com/protocol-buffers/docs/proto3#maps
type ExampleListWithContext struct {
	Examples []*tensorflow1.Example `protobuf:"bytes,1,rep,name=examples" json:"examples,omitempty"`
	Context  *tensorflow1.Example   `protobuf:"bytes,2,opt,name=context" json:"context,omitempty"`
}

func (m *ExampleListWithContext) Reset()                    { *m = ExampleListWithContext{} }
func (m *ExampleListWithContext) String() string            { return proto.CompactTextString(m) }
func (*ExampleListWithContext) ProtoMessage()               {}
func (*ExampleListWithContext) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExampleListWithContext) GetExamples() []*tensorflow1.Example {
	if m != nil {
		return m.Examples
	}
	return nil
}

func (m *ExampleListWithContext) GetContext() *tensorflow1.Example {
	if m != nil {
		return m.Context
	}
	return nil
}

type Input struct {
	// Types that are valid to be assigned to Kind:
	//	*Input_ExampleList
	//	*Input_ExampleListWithContext
	Kind isInput_Kind `protobuf_oneof:"kind"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isInput_Kind interface {
	isInput_Kind()
}

type Input_ExampleList struct {
	ExampleList *ExampleList `protobuf:"bytes,1,opt,name=example_list,json=exampleList,oneof"`
}
type Input_ExampleListWithContext struct {
	ExampleListWithContext *ExampleListWithContext `protobuf:"bytes,2,opt,name=example_list_with_context,json=exampleListWithContext,oneof"`
}

func (*Input_ExampleList) isInput_Kind()            {}
func (*Input_ExampleListWithContext) isInput_Kind() {}

func (m *Input) GetKind() isInput_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Input) GetExampleList() *ExampleList {
	if x, ok := m.GetKind().(*Input_ExampleList); ok {
		return x.ExampleList
	}
	return nil
}

func (m *Input) GetExampleListWithContext() *ExampleListWithContext {
	if x, ok := m.GetKind().(*Input_ExampleListWithContext); ok {
		return x.ExampleListWithContext
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Input) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Input_OneofMarshaler, _Input_OneofUnmarshaler, _Input_OneofSizer, []interface{}{
		(*Input_ExampleList)(nil),
		(*Input_ExampleListWithContext)(nil),
	}
}

func _Input_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Input)
	// kind
	switch x := m.Kind.(type) {
	case *Input_ExampleList:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExampleList); err != nil {
			return err
		}
	case *Input_ExampleListWithContext:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExampleListWithContext); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Input.Kind has unexpected type %T", x)
	}
	return nil
}

func _Input_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Input)
	switch tag {
	case 1: // kind.example_list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExampleList)
		err := b.DecodeMessage(msg)
		m.Kind = &Input_ExampleList{msg}
		return true, err
	case 2: // kind.example_list_with_context
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ExampleListWithContext)
		err := b.DecodeMessage(msg)
		m.Kind = &Input_ExampleListWithContext{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Input_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Input)
	// kind
	switch x := m.Kind.(type) {
	case *Input_ExampleList:
		s := proto.Size(x.ExampleList)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Input_ExampleListWithContext:
		s := proto.Size(x.ExampleListWithContext)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ExampleList)(nil), "tensorflow.serving.ExampleList")
	proto.RegisterType((*ExampleListWithContext)(nil), "tensorflow.serving.ExampleListWithContext")
	proto.RegisterType((*Input)(nil), "tensorflow.serving.Input")
}

func init() { proto.RegisterFile("input.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x2b, 0x28,
	0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2a, 0x49, 0xcd, 0x2b, 0xce, 0x2f, 0x4a,
	0xcb, 0xc9, 0x2f, 0xd7, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc, 0x4b, 0x97, 0x52, 0x45, 0x88, 0xe9,
	0x27, 0xe7, 0x17, 0xa5, 0xea, 0xa7, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xc0, 0x69, 0x88, 0x56, 0x25,
	0x3b, 0x2e, 0x6e, 0x57, 0x88, 0x80, 0x4f, 0x66, 0x71, 0x89, 0x90, 0x3e, 0x17, 0x07, 0x54, 0xbe,
	0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0x48, 0x58, 0x0f, 0xc9, 0x70, 0xa8, 0xd2, 0x20, 0xb8,
	0x22, 0xa5, 0x0a, 0x2e, 0x31, 0x24, 0xfd, 0xe1, 0x99, 0x25, 0x19, 0xce, 0xf9, 0x79, 0x25, 0xa9,
	0x15, 0xa4, 0x1b, 0x25, 0xa4, 0xcb, 0xc5, 0x9e, 0x0c, 0xd1, 0x2b, 0xc1, 0xa4, 0xc0, 0x88, 0x4b,
	0x3d, 0x4c, 0x8d, 0xd2, 0x31, 0x46, 0x2e, 0x56, 0x4f, 0x50, 0x20, 0x08, 0x79, 0x70, 0xf1, 0x40,
	0x0d, 0x89, 0xcf, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x04, 0xeb, 0x96, 0xd7, 0xc3, 0x0c, 0x15, 0x3d,
	0x24, 0xb7, 0x3a, 0x31, 0x69, 0x30, 0x7a, 0x30, 0x04, 0x71, 0xa7, 0x22, 0x79, 0x3f, 0x9b, 0x4b,
	0x12, 0xd9, 0xa4, 0xf8, 0xf2, 0xcc, 0x92, 0x8c, 0x78, 0x54, 0x47, 0x69, 0x11, 0x30, 0x16, 0x29,
	0x08, 0xa0, 0x36, 0x88, 0xa5, 0x62, 0x97, 0x65, 0xe3, 0x62, 0xc9, 0xce, 0xcc, 0x4b, 0x71, 0x62,
	0xfe, 0xc1, 0xc8, 0x98, 0xc4, 0x06, 0x8e, 0x0e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd2,
	0x01, 0xe5, 0x77, 0xd8, 0x01, 0x00, 0x00,
}
