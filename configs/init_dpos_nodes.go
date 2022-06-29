package configs

import (
	"Phoenix-Chain-Core/ethereum/p2p/discover"
	"Phoenix-Chain-Core/libs/crypto/bls"
)

var (
	initialChosenMainNetDPosNodes = []initNode{
		{
			"enode://6c719bccb655c76748c2e9a31402fa5112de24cf6ad4de3e03ff66a324c25913e28bc68418e1e88615eb696bb9088bc3b8aa14dda41dbe6f794919e645982691@39.104.68.32:16789",
			"0c257c93b61b51f7d45b77c50ed8bac0b95736622318aeb7634a1f77be82c5eb61eb3305d396e1739d1d576ba504ba1559f02202642803d772307262d56eabed12c785902c3f2c5d0badb59e98edc8638d4cb45421ceacd3c30aa5e14b5a1790",
		},
		{
			"enode://797e7c01fb7462d0a3eaf8dd8b982d7cb6363503521627d3f31a49b6389c357839e218d154388b11d9e10011f81ae3bbcc93e07d6f2063257d8015018494f3e1@39.104.68.32:16790",
			"d02a0dcc882879d3ee2ae82af8be442b234e9f4541f0d7a96c9199d2223599af77dd3125af0caf318cd1b0d11eb10f18f5b8fff79fedb680ecbdffd20840204103d4f1905414191ed0cc414ca0bc681cd9b178de3b37c9e9f2e29cbe29d7ce18",
		},
		{
			"enode://5db732679631792e4098f0c3cc32653acc394867dc4fbed0855f478adddfe758823571ec9b647a157957af84c90254e8c8d8a9a0848bea1e24cd4006d584cd3e@39.104.62.41:16791",
			"bb18a4dd5041b94d2367ad9e532c4c4da964bfc435a539a94180c8b7d9241cdbf2dc94fbb58b536560e0a3e30713c70873bb96aeab9a085047e4d6de5987f0790c95d9bc1abf7a6452cfaa6ca19564a511c99899678fcde1b84f5b2798099688",
		},
		{
			"enode://f6acf06029e09eba5c222fcaf9cbb55b178457b93d01b2bf92065c7dfdfa844f477fd4fdf5291372cb1ac8af89095bc68b51567d5eaed5d6fd5782b2ce0e64b6@39.104.62.41:16792",
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
