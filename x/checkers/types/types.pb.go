// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tablegamechain/checkers/types.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// All possible checkers game states
type GameState int32

const (
	// The game state is not set
	GameState_GAME_STATE_UNSPECIFIED GameState = 0
	// The game board is set up and ready to play
	GameState_GAME_STATE_READY GameState = 1
	// The game is in progress
	GameState_GAME_STATE_PLAYING GameState = 2
	// The game is complete - no more moves can be made
	GameState_GAME_STATE_COMPLETE GameState = 3
	// The game has expired - no further moves will be accepted
	GameState_GAME_STATE_EXPIRED GameState = 4
	// The game has been forfeited by one of the players
	GameState_GAME_STATE_FORFEITED GameState = 5
)

var GameState_name = map[int32]string{
	0: "GAME_STATE_UNSPECIFIED",
	1: "GAME_STATE_READY",
	2: "GAME_STATE_PLAYING",
	3: "GAME_STATE_COMPLETE",
	4: "GAME_STATE_EXPIRED",
	5: "GAME_STATE_FORFEITED",
}

var GameState_value = map[string]int32{
	"GAME_STATE_UNSPECIFIED": 0,
	"GAME_STATE_READY":       1,
	"GAME_STATE_PLAYING":     2,
	"GAME_STATE_COMPLETE":    3,
	"GAME_STATE_EXPIRED":     4,
	"GAME_STATE_FORFEITED":   5,
}

func (x GameState) String() string {
	return proto.EnumName(GameState_name, int32(x))
}

func (GameState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f223ddd4c8f4063, []int{0}
}

// Game defines the Game type.
type Game struct {
	Board     string `protobuf:"bytes,2,opt,name=board,proto3" json:"board,omitempty"`
	Red       string `protobuf:"bytes,3,opt,name=red,proto3" json:"red,omitempty"`
	Black     string `protobuf:"bytes,4,opt,name=black,proto3" json:"black,omitempty"`
	Turn      string `protobuf:"bytes,5,opt,name=turn,proto3" json:"turn,omitempty"`
	State     string `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	StartTime string `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Creator   string `protobuf:"bytes,9,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f223ddd4c8f4063, []int{0}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Game.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return m.Size()
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func (m *Game) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *Game) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *Game) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

func (m *Game) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Game) GetStartTime() string {
	if m != nil {
		return m.StartTime
	}
	return ""
}

func (m *Game) GetEndTime() string {
	if m != nil {
		return m.EndTime
	}
	return ""
}

func (m *Game) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// A game identified by an index.
type IndexedGame struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Game  *Game  `protobuf:"bytes,2,opt,name=game,proto3" json:"game,omitempty"`
}

func (m *IndexedGame) Reset()         { *m = IndexedGame{} }
func (m *IndexedGame) String() string { return proto.CompactTextString(m) }
func (*IndexedGame) ProtoMessage()    {}
func (*IndexedGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f223ddd4c8f4063, []int{1}
}
func (m *IndexedGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IndexedGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IndexedGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IndexedGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexedGame.Merge(m, src)
}
func (m *IndexedGame) XXX_Size() int {
	return m.Size()
}
func (m *IndexedGame) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexedGame.DiscardUnknown(m)
}

var xxx_messageInfo_IndexedGame proto.InternalMessageInfo

func (m *IndexedGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *IndexedGame) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

func init() {
	proto.RegisterEnum("tablegamechain.checkers.GameState", GameState_name, GameState_value)
	proto.RegisterType((*Game)(nil), "tablegamechain.checkers.Game")
	proto.RegisterType((*IndexedGame)(nil), "tablegamechain.checkers.IndexedGame")
}

func init() {
	proto.RegisterFile("tablegamechain/checkers/types.proto", fileDescriptor_6f223ddd4c8f4063)
}

var fileDescriptor_6f223ddd4c8f4063 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0xdb, 0x74, 0xbb, 0x79, 0xf7, 0x32, 0x8c, 0x65, 0x77, 0x14, 0x36, 0xe8, 0x7a,
	0x11, 0x61, 0xb3, 0xa8, 0x47, 0x4f, 0x71, 0x33, 0x2d, 0x81, 0xfd, 0x13, 0xd2, 0x28, 0xd6, 0x4b,
	0x99, 0x26, 0x83, 0x0d, 0x6d, 0x92, 0x32, 0x19, 0xa1, 0x7e, 0x0b, 0x3f, 0x82, 0x9f, 0xc5, 0x93,
	0xc7, 0x1e, 0x3d, 0x4a, 0xfb, 0x45, 0x64, 0x26, 0x2d, 0xc4, 0xc2, 0xde, 0xde, 0xe7, 0x79, 0x7e,
	0xc9, 0xcc, 0x33, 0xbc, 0xf0, 0x52, 0xf1, 0xe9, 0x42, 0x7c, 0xe5, 0x85, 0x48, 0x67, 0x3c, 0x2f,
	0xaf, 0xd3, 0x99, 0x48, 0xe7, 0x42, 0xd6, 0xd7, 0xea, 0xfb, 0x52, 0xd4, 0xde, 0x52, 0x56, 0xaa,
	0x22, 0xe7, 0xff, 0x43, 0xde, 0x1e, 0xba, 0xfc, 0x85, 0xc0, 0x1e, 0xf2, 0x42, 0x90, 0x3e, 0x74,
	0xa7, 0x15, 0x97, 0x19, 0x3d, 0x7a, 0x8e, 0x5e, 0x39, 0x71, 0x23, 0x08, 0x86, 0x8e, 0x14, 0x19,
	0xed, 0x18, 0x4f, 0x8f, 0x86, 0x5b, 0xf0, 0x74, 0x4e, 0xed, 0x1d, 0xa7, 0x05, 0x21, 0x60, 0xab,
	0x6f, 0xb2, 0xa4, 0x5d, 0x63, 0x9a, 0x59, 0x93, 0xb5, 0xe2, 0x4a, 0xd0, 0xe3, 0x86, 0x34, 0x82,
	0x5c, 0x00, 0xd4, 0x8a, 0x4b, 0x35, 0x51, 0x79, 0x21, 0x68, 0xcf, 0x44, 0x8e, 0x71, 0x92, 0xbc,
	0x10, 0xe4, 0x29, 0x9c, 0x88, 0x32, 0x6b, 0xc2, 0x13, 0x13, 0xf6, 0x44, 0x99, 0x99, 0x88, 0x42,
	0x2f, 0x95, 0x82, 0xab, 0x4a, 0x52, 0xa7, 0x49, 0x76, 0xf2, 0xf2, 0x13, 0x9c, 0x86, 0x65, 0x26,
	0x56, 0x22, 0xdb, 0x57, 0xc9, 0xb5, 0xa4, 0xa8, 0x39, 0xd8, 0x08, 0xf2, 0x06, 0x6c, 0xdd, 0xdf,
	0xf4, 0x3b, 0x7d, 0x7b, 0xe1, 0x3d, 0xf2, 0x22, 0x9e, 0xfe, 0x45, 0x6c, 0xd0, 0xd7, 0x3f, 0x11,
	0x38, 0x5a, 0x8e, 0xcc, 0xcd, 0x9f, 0xc1, 0xd9, 0xd0, 0xbf, 0x63, 0x93, 0x51, 0xe2, 0x27, 0x6c,
	0xf2, 0xf1, 0x7e, 0x14, 0xb1, 0x9b, 0x70, 0x10, 0xb2, 0x00, 0x5b, 0xa4, 0x0f, 0xb8, 0x95, 0xc5,
	0xcc, 0x0f, 0xc6, 0x18, 0x91, 0x33, 0x20, 0x2d, 0x37, 0xba, 0xf5, 0xc7, 0xe1, 0xfd, 0x10, 0x1f,
	0x91, 0x73, 0x78, 0xd2, 0xf2, 0x6f, 0x1e, 0xee, 0xa2, 0x5b, 0x96, 0x30, 0xdc, 0x39, 0xf8, 0x80,
	0x7d, 0x8e, 0xc2, 0x98, 0x05, 0xd8, 0x26, 0x14, 0xfa, 0x2d, 0x7f, 0xf0, 0x10, 0x0f, 0x58, 0x98,
	0xb0, 0x00, 0x77, 0x3f, 0xbc, 0xff, 0xbd, 0x71, 0xd1, 0x7a, 0xe3, 0xa2, 0xbf, 0x1b, 0x17, 0xfd,
	0xd8, 0xba, 0xd6, 0x7a, 0xeb, 0x5a, 0x7f, 0xb6, 0xae, 0xf5, 0xe5, 0x85, 0x29, 0x78, 0xa5, 0x8b,
	0x5c, 0x35, 0x9b, 0xb1, 0x3a, 0xd8, 0x8d, 0xe9, 0xb1, 0x59, 0x8e, 0x77, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xde, 0x5a, 0x1e, 0xeb, 0x43, 0x02, 0x00, 0x00,
}

func (m *Game) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Game) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Game) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.EndTime) > 0 {
		i -= len(m.EndTime)
		copy(dAtA[i:], m.EndTime)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EndTime)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.StartTime) > 0 {
		i -= len(m.StartTime)
		copy(dAtA[i:], m.StartTime)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.StartTime)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Turn) > 0 {
		i -= len(m.Turn)
		copy(dAtA[i:], m.Turn)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Turn)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Board) > 0 {
		i -= len(m.Board)
		copy(dAtA[i:], m.Board)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Board)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *IndexedGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IndexedGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IndexedGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Game != nil {
		{
			size, err := m.Game.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Game) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Board)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Turn)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.StartTime)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EndTime)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *IndexedGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Game != nil {
		l = m.Game.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Game) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Game: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Game: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Board = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Turn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndTime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IndexedGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IndexedGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IndexedGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Game", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Game == nil {
				m.Game = &Game{}
			}
			if err := m.Game.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
