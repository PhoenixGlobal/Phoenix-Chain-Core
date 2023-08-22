package configs

import (
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/ethereum/p2p/discover"
	"github.com/PhoenixGlobal/Phoenix-Chain-Core/libs/crypto/bls"
)

var (
	initialChosenMainNetDPosNodes = []initNode{
		{
			"enode://6c719bccb655c76748c2e9a31402fa5112de24cf6ad4de3e03ff66a324c25913e28bc68418e1e88615eb696bb9088bc3b8aa14dda41dbe6f794919e645982691@13.229.133.200:16789",
			"0c257c93b61b51f7d45b77c50ed8bac0b95736622318aeb7634a1f77be82c5eb61eb3305d396e1739d1d576ba504ba1559f02202642803d772307262d56eabed12c785902c3f2c5d0badb59e98edc8638d4cb45421ceacd3c30aa5e14b5a1790",
		},
		{
			"enode://797e7c01fb7462d0a3eaf8dd8b982d7cb6363503521627d3f31a49b6389c357839e218d154388b11d9e10011f81ae3bbcc93e07d6f2063257d8015018494f3e1@13.229.133.200:16790",
			"d02a0dcc882879d3ee2ae82af8be442b234e9f4541f0d7a96c9199d2223599af77dd3125af0caf318cd1b0d11eb10f18f5b8fff79fedb680ecbdffd20840204103d4f1905414191ed0cc414ca0bc681cd9b178de3b37c9e9f2e29cbe29d7ce18",
		},
		{
			"enode://5db732679631792e4098f0c3cc32653acc394867dc4fbed0855f478adddfe758823571ec9b647a157957af84c90254e8c8d8a9a0848bea1e24cd4006d584cd3e@13.215.175.20:16791",
			"bb18a4dd5041b94d2367ad9e532c4c4da964bfc435a539a94180c8b7d9241cdbf2dc94fbb58b536560e0a3e30713c70873bb96aeab9a085047e4d6de5987f0790c95d9bc1abf7a6452cfaa6ca19564a511c99899678fcde1b84f5b2798099688",
		},
		{
			"enode://f6acf06029e09eba5c222fcaf9cbb55b178457b93d01b2bf92065c7dfdfa844f477fd4fdf5291372cb1ac8af89095bc68b51567d5eaed5d6fd5782b2ce0e64b6@13.215.175.20:16792",
			"e5182404de9160f49afb9ae7305220360e286981e572b7ce94934fb316caae0e414bfd7aff40ea419c63d40e0cb4e20da3119c77426798ec189b7e803cb4d9b6086b737ad468a4dbf33facb8ca382c28c60363923a30e2021d333ec8af8e6807",
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

	addChosenMainNetDPosNodes = []initNode{
		{
			"enode://c885c04643fb42478138634b9c7a1e558c564b05da17b7df301dda74d34083d6a7dd73643a40d48f1a51f8fc7dfec01d99f7b9166228f378da0c08884dc59db0@44.234.65.70:16789",
			"4f4fba0bfe256b74bbcb29c9cf5fd29218c8ad9ce645d70d3b8137231e3705461b87e642191df0dc2caec95583a9d8196cb5be440a5740d53a3603631fbc2e53891d438de0a5e7a9d13e6073e352295c2caae381fefe827c4c2fe7d8d29e1c89",
		},
		{
			"enode://25e3f1d44a78bd33c784d11dc36cc651bb5efa83cf3240df57ba82810e2f1741a1cb0497a0d59c1f3dfead03bded3571e7ba721fafe5063ebbab4f3f68459519@34.223.113.42:16789",
			"71a2eafbcd8cd132240b5a8c11035daaccbbc12a48271ec5032bd2fd915832b40f1c662681e723ca931052bda5ed7d0cdcb8760554b6808e2098dc6b58f22a1e8badb11aead0c992538af7368988e674e28712156c034610eb0550a9370b390f",
		},
		{
			"enode://026ffc3cf1a51a55524062bee5cca7fcde08f19d03dde5c16f4a66909fab7781e4029ab25732f2f35de949dcc1d3686c07ff6af2916fa0d961dac82ca94be3e1@44.234.51.161:16789",
			"d33fc19baa95deda6bbb85d1aec019e63eb1ac21cb1f68524d7788b7d342095f02e419c7713694b9228ea671a69aab123e46c2d05592bfcb834ea377124dc539e76dfe47963d5e2b7dee0fc1fb67f992e7f06669338467e34df6bdb7b558b096",
		},
	}


	InitialChosenMainNetValidators []Validator

	InitialChosenTestnetValidators []Validator

	AddChosenMainNetValidators []Validator
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

	for _, n := range addChosenMainNetDPosNodes {
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
		AddChosenMainNetValidators = append(AddChosenMainNetValidators, *validator)
	}
}
