package configs

import (
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/crypto/bls"
)

var (
	initialChosenMainNetDPosNodes = []initNode{
		{
			"enode://f4d3989cfa7c30f09ebd24ab733ff97d5e10006fdba93af7c64e02e4cc9becec28d888badfeb9515a23a9c17497161e6bec03ea28113ecaa7adab8cb1b5990a0@www.phoenix.global:16789",
			"ca0c71e76b2ba6147467258238714ffe55ce5e57f3b50adb80f8d58aadc9f4eb7e338bb46d4ac8cdd4801625257d4d05c3aa0f446048b00ebb36eb94c3bc3dd18baa0e164dd150cc11bd708baef1e9b08c0496d6b498944b7eaa8cd685238c93",
		},
		{
			"enode://0116873bb440fb2aeadbb402e6093f5b97f0574cec02c53db6ff99cc5c2dc6dc4ee578961640ed80f9f5b69605493c3a88c8137eb06497596fe960b4c4a9053e@www.phoenix.global:16789",
			"e4f3774f6e043bc106503c291935047eed3f6b109bd3c29762fab6d14a00c775a315dac8ee2a82832476232efed4ff126e006ae4bccc9c708a615da105a1b3a0998dc2ea8a5f910e3280225b1bbc19c1132dfc661e4227fbf29bbb391706a589",
		},
		{
			"enode://8c816c9441f3b6d64efe2382110d17cd1ccd10a4a60db742383c36e6cabaeef2a2f87815ecd0f285e74ad6f9911a2db1dde7fe33a249e9c625ce225b0dbe25f0@www.phoenix.global:16789",
			"5c8e0280a5b72cb756437db1d62cf0803d4789a8cc3ff1989b26175cb63880d953c43352f6aa3e2c16c7cc933a669c1862a8011506262625a2a1ba7e79b141ce9ce54fa64c2d98073ecb8644accf63129229581eb318cc14d4f5e529171c2f95",
		},
		{
			"enode://33761aca567c3ce253635bfea65cb48d5518eaabfefd92e3d443414ea31a1abfff5dfda1f0c47f6b4cab0efd55a535a268acd82b3a9d078b40e328b890f49291@www.phoenix.global:16789",
			"660067d3c46fc7e41cd9b38e81937dbcb5a124524daefa3c9ade3330d84f92f19146e53c2c60abeb4a2044351077571866947234b4516e5200f925fc6f3d69b732bfc6a103f86cc2c25a9925b170adba04dd194864b90e9eb22c21196907830f",
		},
		{
			"enode://be5409650380b505bbb754663b0cc04f105fd792dcab394a78cd62be641f8ba976ff11bc814becfa23174b0852cc5f6db6d631a89d411c140d104c0927fb3822@www.phoenix.global:16789",
			"305a676d0b03d3f2c2018791bb4a7a855d4b2d226a09b9ed029dedec1a7774364dc8a8f55ba4fb8d39a40093d6bacb17732465002baf959913737c5f6e062704bb29b815ab5d8684ffa7de79a3e518fde7c7df34a18b92660a64f37eddfddb0d",
		},
		{
			"enode://39f455da5a67144c6f453bd3c375db85071c51c400c6b408a1361493f4ca1b169ac3e1ab76cfad8ba6612118c484e374516f0747facd9b3c22acd4fdd4c347b5@www.phoenix.global:16789",
			"ed62192c500c0fc4b1d7ad204b4eb88edb602a19e6767722e2e8aed44bd0eee4e8998e2407c654c2684cb975ab7c25019c09d5a519e32232aa532433af09753f8c3c3ee6d8ade73eca72fe5194f917a990bef42d11e9f417989e0c4d9db1a412",
		},
		{
			"enode://fb334dd1f0803291aebc242eeac3deb4f0132cbfa7c1d9c6acd1fb8277f9c83d3ed12cf9cca0a0f7d224d7761f0f93dc1c82c5540db8640bf0f21316795e0ccd@www.phoenix.global:16789",
			"4375213672319c43b731639e4ad70ed7b4ec8ae46289fdba02a5cc5b2813731f7f71fcb7e07263e58fc9e2badb9b780e01c90e10918620c01734429ef30a70df138edddb253f0efcde885339c7e16f9115bff7d24e33a7e283f45d7b4227d997",
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
