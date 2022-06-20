package configs

import (
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/crypto/bls"
)

var (
	initialChosenMainNetDPosNodes = []initNode{
		{
			"enode://e0fbf44d916ee125098327e378a3b5f23b421c9636b7662e8fec1d0c8ab4c7addcbeae865c005734f8451873eef11f64f361bdac3c57ad2143e892af24127769@127.0.0.1:16789",
			"18fb11d0264dca2650b6e1a99208042d5a4552477f1899f04223b601b68371723a86da2325e566dab293b1f2d11b7e0610fef25bb1a6e44a3e1cfa580c0d4420dcaa480e7d01485ea8cb0f0225b6ac0a4693f0fc7600ab995acf09b3a025e790",
		},
		{
			"enode://21178e44c436c1f0c824726870b41317d1dbb79971d10afa4690df5d16ccad7e4da2ddcf52e06fc0b7bc21946bb9795e10dc860c4389b6ba223e66172f690cdb@127.0.0.1:16790",
			"b2109c8e29d777cca4a0e073c29591e5368dab8a0be8343ff07019e63d2864dc2e1bededd50b1131ac6f1580168db51704ca77efaa8503d2728bff66455b4f82cb7d3e4961ec680717a83ec9ead66859f5d51c96fcb042e43ef847bd9caf4f89",
		},
		{
			"enode://b29e88da35aebf5571616032effe8540453c1fb3eb0fdc6b0f1d7880d853e7ca63f8835f11dd5960095027b64e01a3e446f3492bc7ec2cea709016f195a4661b@127.0.0.1:16791",
			"74575eb9e2654f917af57345592a737cf577a5db16af81cd87db87c05e9a775c8ae6bb6eb2ca882e48ee451b849cd5182c7645ae80fdb7583c6752e03a1a4be19f58f98dedd50532992b46696d2577901e51c651d018503b80b918423d80128d",
		},
		{
			"enode://1dc14f2f7a6e69b8357fa07d21f62e0cd5fb52f1704a06351401bfa543293ed3e7e23cd1432b8df33c3df2caecf083cfc69493c4003942710005387b888e36fe@127.0.0.1:16792",
			"4ec30d20e3ff50518a53536d2ec6cca8b399b2c60415e59ea263a3b885c8e64db95a7e4452e22dad5784dfbb055b8808020230d270d2bed0515dc1b4e8de049f6efca5f4c9cbe8c0a272d954d83d23c7a5035efe273384a98d5a334595ac910b",
		},
	}

	initialChosenTestnetDPosNodes = []initNode{
		{
			"enode://1ae3a2bfe9d93a4c96d44cabdda7d1ba2b5c653ab13d4593020add6238de1074b159f7eb072ca1a0ede3252fd44023aabda61a7f4fef16001bc9f83e1877296f@127.0.0.1:16789",
			"964211deb768dce2782600e0f9bbc55a1c5c6813b4ca262a0ba59fe389070950f2a7d25cb73e1e25dcc2262a12451714070e81b9fddd43ba1c0c456584338d1a2e5ec274a4cd4f8ef808174b323cadb2d2e06b64b08d590186b2efebc3c10b05",
		},
		{
			"enode://db18af9be2af9dff2347c3d06db4b1bada0598d099a210275251b68fa7b5a863d47fcdd382cc4b3ea01e5b55e9dd0bdbce654133b7f58928ce74629d5e68b974@127.0.0.1:16789",
			"5d0f8a399533b3f9b3a7198282c4b7b8b414529c66861d7958ebf908664707e5e6b353630b94ac5c1173c36e889fb403208ff73d233c12865d9e32256bbb988b931d41fda48e450b992fa5ec67790081e730965f548120b6d9fdc6156d66a614",
		},
	}

	InitialChosenMainNetValidators []Validator

	InitialChosenTestnetValidators []Validator
)

func init()  {
	for _, n := range initialChosenMainNetDPosNodes {
		validator := new(Validator)
		if node, err := discover.ParseNode(n.Enode); nil == err {
			validator.NodeId = node.ID
		}
		if n.BlsPubkey != "" {
			var blsPk bls.PublicKeyHex
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				validator.BlsPubKey = blsPk
			}
		}
		InitialChosenMainNetValidators = append(InitialChosenMainNetValidators, *validator)
	}

	for _, n := range initialChosenTestnetDPosNodes {
		validator := new(Validator)
		if node, err := discover.ParseNode(n.Enode); nil == err {
			validator.NodeId = node.ID
		}
		if n.BlsPubkey != "" {
			var blsPk bls.PublicKeyHex
			if err := blsPk.UnmarshalText([]byte(n.BlsPubkey)); nil == err {
				validator.BlsPubKey = blsPk
			}
		}
		InitialChosenTestnetValidators = append(InitialChosenTestnetValidators, *validator)
	}
}
