package models

import "Grpc-crud/app/pb"

type Binder struct {
	BlockA []string `json:"block_a,omitempty"`
	BlockB []string `json:"block_b,omitempty"`
	BlockC []string `json:"block_c,omitempty"`
	BlockD []string `json:"block_d,omitempty"`
	BlockE []string `json:"block_e,omitempty"`
	BlockF []string `json:"block_f,omitempty"`
	BlockG []string `json:"block_g,omitempty"`
	BlockH []string `json:"block_h,omitempty"`
	BlockI []string `json:"block_i,omitempty"`
	BlockJ []string `json:"block_j,omitempty"`
	BlockK []string `json:"block_k,omitempty"`
}
type BinderResponse struct {
	Message              string   `json:"message,omitempty"`
}
func (b *Binder) ToProtoBuf() *pb.Binder {
	return &pb.Binder{
		BlockA: b.BlockA,
		BlockB: b.BlockB,
		BlockC: b.BlockC,
		BlockD: b.BlockD,
		BlockE: b.BlockE,
		BlockF: b.BlockF,
		BlockG: b.BlockG,
		BlockH: b.BlockH,
		BlockI: b.BlockI,
		BlockJ: b.BlockJ,
		BlockK: b.BlockK,
	}
}
func (b * Binder)FromProtoBuf(bin * pb.Binder){
	b.BlockA = bin.BlockA
	b.BlockB = bin.BlockB
	b.BlockC = bin.BlockC
	b.BlockD = bin.BlockD
	b.BlockE = bin.BlockE
	b.BlockF = bin.BlockF
	b.BlockG = bin.BlockG
	b.BlockH = bin.BlockH
	b.BlockI = bin.BlockI
	b.BlockJ = bin.BlockJ
	b.BlockK = bin.BlockK 
}

func (b *BinderResponse) ToProtoBuf ()*pb.BinderResponse{
	return &pb.BinderResponse{
		Message: b.Message,
	}
}
func (b * BinderResponse) FromProtoBuf (br * pb.BinderResponse){
	b.Message = br.Message
}