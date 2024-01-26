// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/stakeibc/trade_route.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Stores pool information needed to execute the swap along a trade route
type TradeConfig struct {
	// Currently Osmosis is the only trade chain so this is an osmosis pool id
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// Spot price in the pool to convert the reward denom to the host denom
	// output_tokens = swap_price * input tokens
	// This value may be slightly stale as it is updated by an ICQ
	SwapPrice github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=swap_price,json=swapPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"swap_price"`
	// unix time in seconds that the price was last updated
	PriceUpdateTimestamp uint64 `protobuf:"varint,3,opt,name=price_update_timestamp,json=priceUpdateTimestamp,proto3" json:"price_update_timestamp,omitempty"`
	// Threshold defining the percentage of tokens that could be lost in the trade
	// This captures both the loss from slippage and from a stale price on stride
	// 0.05 means the output from the trade can be no less than a 5% deviation
	// from the current value
	MaxAllowedSwapLossRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=max_allowed_swap_loss_rate,json=maxAllowedSwapLossRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"max_allowed_swap_loss_rate"`
	// min and max set boundaries of reward denom on trade chain we will swap
	// min also decides when reward token transfers are worth it (transfer fees)
	MinSwapAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=min_swap_amount,json=minSwapAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_swap_amount"`
	MaxSwapAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=max_swap_amount,json=maxSwapAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_swap_amount"`
}

func (m *TradeConfig) Reset()         { *m = TradeConfig{} }
func (m *TradeConfig) String() string { return proto.CompactTextString(m) }
func (*TradeConfig) ProtoMessage()    {}
func (*TradeConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c252b142ecf88017, []int{0}
}
func (m *TradeConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradeConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradeConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradeConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeConfig.Merge(m, src)
}
func (m *TradeConfig) XXX_Size() int {
	return m.Size()
}
func (m *TradeConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TradeConfig proto.InternalMessageInfo

func (m *TradeConfig) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *TradeConfig) GetPriceUpdateTimestamp() uint64 {
	if m != nil {
		return m.PriceUpdateTimestamp
	}
	return 0
}

// TradeRoute represents a round trip including info on transfer and how to do
// the swap. It makes the assumption that the reward token is always foreign to
// the host so therefore the first two hops are to unwind the ibc denom enroute
// to the trade chain and the last hop is the return so funds start/end in the
// withdrawl ICA on hostZone
// The structure is key'd on reward denom and host denom in their native forms
// (i.e. reward_denom_on_reward_zone and host_denom_on_host_zone)
type TradeRoute struct {
	// ibc denom for the reward on the host zone
	RewardDenomOnHostZone string `protobuf:"bytes,1,opt,name=reward_denom_on_host_zone,json=rewardDenomOnHostZone,proto3" json:"reward_denom_on_host_zone,omitempty"`
	// should be the native denom for the reward chain
	RewardDenomOnRewardZone string `protobuf:"bytes,2,opt,name=reward_denom_on_reward_zone,json=rewardDenomOnRewardZone,proto3" json:"reward_denom_on_reward_zone,omitempty"`
	// ibc denom of the reward on the trade chain, input to the swap
	RewardDenomOnTradeZone string `protobuf:"bytes,3,opt,name=reward_denom_on_trade_zone,json=rewardDenomOnTradeZone,proto3" json:"reward_denom_on_trade_zone,omitempty"`
	// ibc of the host denom on the trade chain, output from the swap
	HostDenomOnTradeZone string `protobuf:"bytes,4,opt,name=host_denom_on_trade_zone,json=hostDenomOnTradeZone,proto3" json:"host_denom_on_trade_zone,omitempty"`
	// should be the same as the native host denom on the host chain
	HostDenomOnHostZone string `protobuf:"bytes,5,opt,name=host_denom_on_host_zone,json=hostDenomOnHostZone,proto3" json:"host_denom_on_host_zone,omitempty"`
	// ICAAccount on the host zone with the reward tokens
	// This is the same as the host zone withdrawal ICA account
	HostAccount ICAAccount `protobuf:"bytes,6,opt,name=host_account,json=hostAccount,proto3" json:"host_account"`
	// ICAAccount on the reward zone that is acts as the intermediate
	// receiver of the transfer from host zone to trade zone
	RewardAccount ICAAccount `protobuf:"bytes,7,opt,name=reward_account,json=rewardAccount,proto3" json:"reward_account"`
	// ICAAccount responsible for executing the swap of reward
	// tokens for host tokens
	TradeAccount ICAAccount `protobuf:"bytes,8,opt,name=trade_account,json=tradeAccount,proto3" json:"trade_account"`
	// Channel responsible for the transfer of reward tokens from the host
	// zone to the reward zone. This is the channel ID on the host zone side
	HostToRewardChannelId string `protobuf:"bytes,9,opt,name=host_to_reward_channel_id,json=hostToRewardChannelId,proto3" json:"host_to_reward_channel_id,omitempty"`
	// Channel responsible for the transfer of reward tokens from the reward
	// zone to the trade zone. This is the channel ID on the reward zone side
	RewardToTradeChannelId string `protobuf:"bytes,10,opt,name=reward_to_trade_channel_id,json=rewardToTradeChannelId,proto3" json:"reward_to_trade_channel_id,omitempty"`
	// Channel responsible for the transfer of host tokens from the trade
	// zone, back to the host zone. This is the channel ID on the trade zone side
	TradeToHostChannelId string `protobuf:"bytes,11,opt,name=trade_to_host_channel_id,json=tradeToHostChannelId,proto3" json:"trade_to_host_channel_id,omitempty"`
	// specifies the configuration needed to execute the swap
	// such as pool_id, slippage, min trade amount, etc.
	TradeConfig TradeConfig `protobuf:"bytes,12,opt,name=trade_config,json=tradeConfig,proto3" json:"trade_config"`
}

func (m *TradeRoute) Reset()         { *m = TradeRoute{} }
func (m *TradeRoute) String() string { return proto.CompactTextString(m) }
func (*TradeRoute) ProtoMessage()    {}
func (*TradeRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_c252b142ecf88017, []int{1}
}
func (m *TradeRoute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TradeRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TradeRoute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TradeRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeRoute.Merge(m, src)
}
func (m *TradeRoute) XXX_Size() int {
	return m.Size()
}
func (m *TradeRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TradeRoute proto.InternalMessageInfo

func (m *TradeRoute) GetRewardDenomOnHostZone() string {
	if m != nil {
		return m.RewardDenomOnHostZone
	}
	return ""
}

func (m *TradeRoute) GetRewardDenomOnRewardZone() string {
	if m != nil {
		return m.RewardDenomOnRewardZone
	}
	return ""
}

func (m *TradeRoute) GetRewardDenomOnTradeZone() string {
	if m != nil {
		return m.RewardDenomOnTradeZone
	}
	return ""
}

func (m *TradeRoute) GetHostDenomOnTradeZone() string {
	if m != nil {
		return m.HostDenomOnTradeZone
	}
	return ""
}

func (m *TradeRoute) GetHostDenomOnHostZone() string {
	if m != nil {
		return m.HostDenomOnHostZone
	}
	return ""
}

func (m *TradeRoute) GetHostAccount() ICAAccount {
	if m != nil {
		return m.HostAccount
	}
	return ICAAccount{}
}

func (m *TradeRoute) GetRewardAccount() ICAAccount {
	if m != nil {
		return m.RewardAccount
	}
	return ICAAccount{}
}

func (m *TradeRoute) GetTradeAccount() ICAAccount {
	if m != nil {
		return m.TradeAccount
	}
	return ICAAccount{}
}

func (m *TradeRoute) GetHostToRewardChannelId() string {
	if m != nil {
		return m.HostToRewardChannelId
	}
	return ""
}

func (m *TradeRoute) GetRewardToTradeChannelId() string {
	if m != nil {
		return m.RewardToTradeChannelId
	}
	return ""
}

func (m *TradeRoute) GetTradeToHostChannelId() string {
	if m != nil {
		return m.TradeToHostChannelId
	}
	return ""
}

func (m *TradeRoute) GetTradeConfig() TradeConfig {
	if m != nil {
		return m.TradeConfig
	}
	return TradeConfig{}
}

func init() {
	proto.RegisterType((*TradeConfig)(nil), "stride.stakeibc.TradeConfig")
	proto.RegisterType((*TradeRoute)(nil), "stride.stakeibc.TradeRoute")
}

func init() { proto.RegisterFile("stride/stakeibc/trade_route.proto", fileDescriptor_c252b142ecf88017) }

var fileDescriptor_c252b142ecf88017 = []byte{
	// 661 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcf, 0x4f, 0x13, 0x41,
	0x14, 0xc7, 0x5b, 0xf9, 0x65, 0xa7, 0x45, 0x92, 0x15, 0x61, 0x29, 0xa6, 0x20, 0x07, 0xc3, 0x85,
	0xdd, 0x88, 0x84, 0x10, 0xc3, 0xa5, 0x50, 0x0d, 0x4d, 0x48, 0x34, 0x4b, 0xf5, 0x80, 0x87, 0xc9,
	0x74, 0x77, 0x2c, 0x1b, 0xba, 0xf3, 0x36, 0x3b, 0x53, 0xbb, 0xfa, 0x57, 0xf8, 0xc7, 0xf8, 0x47,
	0x70, 0x24, 0x9e, 0x8c, 0x07, 0x62, 0xe8, 0x9f, 0xe1, 0xc5, 0xcc, 0x9b, 0x5d, 0xd8, 0x82, 0x07,
	0x34, 0x9e, 0xda, 0xd9, 0xf7, 0x3e, 0xdf, 0x37, 0x6f, 0xbe, 0x33, 0x8f, 0x3c, 0x91, 0x2a, 0x09,
	0x03, 0xee, 0x4a, 0xc5, 0x4e, 0x79, 0xd8, 0xf5, 0x5d, 0x95, 0xb0, 0x80, 0xd3, 0x04, 0x06, 0x8a,
	0x3b, 0x71, 0x02, 0x0a, 0xac, 0x39, 0x93, 0xe2, 0xe4, 0x29, 0xf5, 0xf9, 0x1e, 0xf4, 0x00, 0x63,
	0xae, 0xfe, 0x67, 0xd2, 0xea, 0xb7, 0x94, 0x42, 0x9f, 0x51, 0xe6, 0xfb, 0x30, 0x10, 0x2a, 0x4b,
	0x59, 0xf2, 0x41, 0x46, 0x20, 0xa9, 0x61, 0xcd, 0xc2, 0x84, 0xd6, 0x46, 0x13, 0xa4, 0xda, 0xd1,
	0xa5, 0xf7, 0x41, 0x7c, 0x08, 0x7b, 0xd6, 0x22, 0x99, 0x89, 0x01, 0xfa, 0x34, 0x0c, 0xec, 0xf2,
	0x6a, 0x79, 0x7d, 0xd2, 0x9b, 0xd6, 0xcb, 0x76, 0x60, 0xbd, 0x27, 0x44, 0x0e, 0x59, 0x4c, 0xe3,
	0x24, 0xf4, 0xb9, 0x7d, 0x6f, 0xb5, 0xbc, 0x5e, 0xd9, 0xdb, 0x3d, 0xbb, 0x58, 0x29, 0xfd, 0xb8,
	0x58, 0x79, 0xda, 0x0b, 0xd5, 0xc9, 0xa0, 0xeb, 0xf8, 0x10, 0x65, 0xea, 0xd9, 0xcf, 0x86, 0x0c,
	0x4e, 0x5d, 0xf5, 0x29, 0xe6, 0xd2, 0x69, 0x71, 0xff, 0xdb, 0xd7, 0x0d, 0x92, 0x15, 0x6f, 0x71,
	0xdf, 0xab, 0x68, 0xbd, 0x37, 0x5a, 0xce, 0xda, 0x22, 0x0b, 0xa8, 0x4b, 0x07, 0x71, 0xc0, 0x14,
	0xa7, 0x2a, 0x8c, 0xb8, 0x54, 0x2c, 0x8a, 0xed, 0x09, 0xdc, 0xc4, 0x3c, 0x46, 0xdf, 0x62, 0xb0,
	0x93, 0xc7, 0xac, 0x94, 0xd4, 0x23, 0x96, 0x52, 0xd6, 0xef, 0xc3, 0x90, 0x07, 0x14, 0xb7, 0xd7,
	0x07, 0x29, 0x69, 0xc2, 0x14, 0xb7, 0x27, 0xff, 0xc3, 0x16, 0x17, 0x22, 0x96, 0x36, 0x8d, 0xfc,
	0xd1, 0x90, 0xc5, 0x87, 0x20, 0xa5, 0xc7, 0x14, 0xb7, 0xde, 0x91, 0xb9, 0x28, 0x14, 0xa6, 0x22,
	0x8b, 0xf4, 0x49, 0xdb, 0x53, 0x58, 0xce, 0xf9, 0x8b, 0x72, 0x6d, 0xa1, 0xbc, 0xd9, 0x28, 0x14,
	0x5a, 0xb9, 0x89, 0x22, 0xa8, 0xcb, 0xd2, 0x31, 0xdd, 0xe9, 0x7f, 0xd4, 0x65, 0xe9, 0xb5, 0xee,
	0xda, 0xaf, 0x29, 0x42, 0xd0, 0x65, 0x4f, 0xdf, 0x2f, 0x6b, 0x87, 0x2c, 0x25, 0x7c, 0xc8, 0x92,
	0x80, 0x06, 0x5c, 0x40, 0x44, 0x41, 0xd0, 0x13, 0x90, 0x8a, 0x7e, 0x06, 0xc1, 0xd1, 0xf6, 0x8a,
	0xf7, 0xc8, 0x24, 0xb4, 0x74, 0xfc, 0xb5, 0x38, 0x00, 0xa9, 0x8e, 0x41, 0x70, 0x6b, 0x97, 0x2c,
	0xdf, 0x24, 0xb3, 0x35, 0xb2, 0x78, 0x2d, 0xbc, 0xc5, 0x31, 0xd6, 0xc3, 0x05, 0xd2, 0x2f, 0x48,
	0xfd, 0x26, 0x6d, 0xae, 0x3d, 0xc2, 0x13, 0x08, 0x2f, 0x8c, 0xc1, 0xb8, 0x69, 0x64, 0xb7, 0x89,
	0x8d, 0x7b, 0xfc, 0x13, 0x89, 0x56, 0x7b, 0xf3, 0x3a, 0x7e, 0x8b, 0xdb, 0x22, 0x8b, 0xe3, 0xdc,
	0x75, 0xa7, 0x68, 0x99, 0xf7, 0xb0, 0x80, 0x5d, 0xf5, 0xd9, 0x22, 0x35, 0xcc, 0xcb, 0xde, 0x11,
	0xba, 0x50, 0xdd, 0x5c, 0x76, 0x6e, 0x3c, 0x49, 0xa7, 0xbd, 0xdf, 0x6c, 0x9a, 0x94, 0xbd, 0x49,
	0x6d, 0x91, 0x57, 0xd5, 0x58, 0xf6, 0xc9, 0x3a, 0x20, 0x0f, 0xb2, 0x7e, 0x73, 0x9d, 0x99, 0xbb,
	0xea, 0xcc, 0x1a, 0x30, 0x57, 0x7a, 0x45, 0x66, 0x4d, 0xbf, 0xb9, 0xd0, 0xfd, 0xbb, 0x0a, 0xd5,
	0x90, 0xcb, 0x75, 0x76, 0xc8, 0x12, 0xf6, 0xa5, 0x20, 0xf7, 0xcd, 0x3f, 0x61, 0x42, 0x70, 0x7c,
	0xf0, 0x15, 0xe3, 0xbc, 0x4e, 0xe8, 0x80, 0xb1, 0x6d, 0xdf, 0x44, 0xdb, 0x41, 0xc1, 0x3b, 0x05,
	0xd9, 0xd9, 0x17, 0x50, 0x52, 0xf4, 0xae, 0x03, 0x66, 0xa2, 0x5c, 0xb1, 0xdb, 0xc4, 0x36, 0x84,
	0x02, 0x73, 0xfc, 0x05, 0xb2, 0x6a, 0xbc, 0xc3, 0x78, 0x07, 0xb4, 0x01, 0xd7, 0xdc, 0x4b, 0x52,
	0xcb, 0x2a, 0xe1, 0x70, 0xb2, 0x6b, 0xd8, 0xf4, 0xe3, 0x5b, 0x4d, 0x17, 0x06, 0x58, 0x6e, 0x83,
	0x2a, 0x7c, 0x3a, 0x3c, 0xbb, 0x6c, 0x94, 0xcf, 0x2f, 0x1b, 0xe5, 0x9f, 0x97, 0x8d, 0xf2, 0x97,
	0x51, 0xa3, 0x74, 0x3e, 0x6a, 0x94, 0xbe, 0x8f, 0x1a, 0xa5, 0xe3, 0xcd, 0xc2, 0x73, 0x3a, 0x42,
	0xd1, 0x8d, 0x43, 0xd6, 0x95, 0x6e, 0x36, 0x52, 0x3f, 0x3e, 0xdb, 0x71, 0xd3, 0xc2, 0x88, 0xd6,
	0xcf, 0xab, 0x3b, 0x8d, 0x83, 0xf3, 0xf9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x8e, 0x30,
	0xfc, 0xc2, 0x05, 0x00, 0x00,
}

func (m *TradeConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradeConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradeConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxSwapAmount.Size()
		i -= size
		if _, err := m.MaxSwapAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.MinSwapAmount.Size()
		i -= size
		if _, err := m.MinSwapAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.MaxAllowedSwapLossRate.Size()
		i -= size
		if _, err := m.MaxAllowedSwapLossRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PriceUpdateTimestamp != 0 {
		i = encodeVarintTradeRoute(dAtA, i, uint64(m.PriceUpdateTimestamp))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.SwapPrice.Size()
		i -= size
		if _, err := m.SwapPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PoolId != 0 {
		i = encodeVarintTradeRoute(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TradeRoute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TradeRoute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TradeRoute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TradeConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.TradeToHostChannelId) > 0 {
		i -= len(m.TradeToHostChannelId)
		copy(dAtA[i:], m.TradeToHostChannelId)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.TradeToHostChannelId)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.RewardToTradeChannelId) > 0 {
		i -= len(m.RewardToTradeChannelId)
		copy(dAtA[i:], m.RewardToTradeChannelId)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.RewardToTradeChannelId)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.HostToRewardChannelId) > 0 {
		i -= len(m.HostToRewardChannelId)
		copy(dAtA[i:], m.HostToRewardChannelId)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.HostToRewardChannelId)))
		i--
		dAtA[i] = 0x4a
	}
	{
		size, err := m.TradeAccount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size, err := m.RewardAccount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.HostAccount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTradeRoute(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.HostDenomOnHostZone) > 0 {
		i -= len(m.HostDenomOnHostZone)
		copy(dAtA[i:], m.HostDenomOnHostZone)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.HostDenomOnHostZone)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.HostDenomOnTradeZone) > 0 {
		i -= len(m.HostDenomOnTradeZone)
		copy(dAtA[i:], m.HostDenomOnTradeZone)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.HostDenomOnTradeZone)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.RewardDenomOnTradeZone) > 0 {
		i -= len(m.RewardDenomOnTradeZone)
		copy(dAtA[i:], m.RewardDenomOnTradeZone)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.RewardDenomOnTradeZone)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.RewardDenomOnRewardZone) > 0 {
		i -= len(m.RewardDenomOnRewardZone)
		copy(dAtA[i:], m.RewardDenomOnRewardZone)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.RewardDenomOnRewardZone)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.RewardDenomOnHostZone) > 0 {
		i -= len(m.RewardDenomOnHostZone)
		copy(dAtA[i:], m.RewardDenomOnHostZone)
		i = encodeVarintTradeRoute(dAtA, i, uint64(len(m.RewardDenomOnHostZone)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTradeRoute(dAtA []byte, offset int, v uint64) int {
	offset -= sovTradeRoute(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TradeConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovTradeRoute(uint64(m.PoolId))
	}
	l = m.SwapPrice.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	if m.PriceUpdateTimestamp != 0 {
		n += 1 + sovTradeRoute(uint64(m.PriceUpdateTimestamp))
	}
	l = m.MaxAllowedSwapLossRate.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	l = m.MinSwapAmount.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	l = m.MaxSwapAmount.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	return n
}

func (m *TradeRoute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RewardDenomOnHostZone)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.RewardDenomOnRewardZone)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.RewardDenomOnTradeZone)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.HostDenomOnTradeZone)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.HostDenomOnHostZone)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = m.HostAccount.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	l = m.RewardAccount.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	l = m.TradeAccount.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	l = len(m.HostToRewardChannelId)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.RewardToTradeChannelId)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = len(m.TradeToHostChannelId)
	if l > 0 {
		n += 1 + l + sovTradeRoute(uint64(l))
	}
	l = m.TradeConfig.Size()
	n += 1 + l + sovTradeRoute(uint64(l))
	return n
}

func sovTradeRoute(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTradeRoute(x uint64) (n int) {
	return sovTradeRoute(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TradeConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTradeRoute
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
			return fmt.Errorf("proto: TradeConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradeConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwapPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwapPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PriceUpdateTimestamp", wireType)
			}
			m.PriceUpdateTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PriceUpdateTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllowedSwapLossRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxAllowedSwapLossRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSwapAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinSwapAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSwapAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxSwapAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTradeRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTradeRoute
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
func (m *TradeRoute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTradeRoute
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
			return fmt.Errorf("proto: TradeRoute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TradeRoute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDenomOnHostZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardDenomOnHostZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDenomOnRewardZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardDenomOnRewardZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardDenomOnTradeZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardDenomOnTradeZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDenomOnTradeZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostDenomOnTradeZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostDenomOnHostZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostDenomOnHostZone = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HostAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RewardAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostToRewardChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostToRewardChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardToTradeChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardToTradeChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeToHostChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TradeToHostChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TradeConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTradeRoute
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
				return ErrInvalidLengthTradeRoute
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTradeRoute
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TradeConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTradeRoute(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTradeRoute
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
func skipTradeRoute(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTradeRoute
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
					return 0, ErrIntOverflowTradeRoute
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
					return 0, ErrIntOverflowTradeRoute
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
				return 0, ErrInvalidLengthTradeRoute
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTradeRoute
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTradeRoute
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTradeRoute        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTradeRoute          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTradeRoute = fmt.Errorf("proto: unexpected end of group")
)
