package configs

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main PhoenixChain network.
var MainnetBootnodes = []string{
	"enode://be5409650380b505bbb754663b0cc04f105fd792dcab394a78cd62be641f8ba976ff11bc814becfa23174b0852cc5f6db6d631a89d411c140d104c0927fb3822@www.phoenix.global:16789",
	"enode://39f455da5a67144c6f453bd3c375db85071c51c400c6b408a1361493f4ca1b169ac3e1ab76cfad8ba6612118c484e374516f0747facd9b3c22acd4fdd4c347b5@www.phoenix.global:16789",
	"enode://fb334dd1f0803291aebc242eeac3deb4f0132cbfa7c1d9c6acd1fb8277f9c83d3ed12cf9cca0a0f7d224d7761f0f93dc1c82c5540db8640bf0f21316795e0ccd@www.phoenix.global:16789",
	"enode://1ae3a2bfe9d93a4c96d44cabdda7d1ba2b5c653ab13d4593020add6238de1074b159f7eb072ca1a0ede3252fd44023aabda61a7f4fef16001bc9f83e1877296f@www.phoenix.global:16789",
	"enode://f4d3989cfa7c30f09ebd24ab733ff97d5e10006fdba93af7c64e02e4cc9becec28d888badfeb9515a23a9c17497161e6bec03ea28113ecaa7adab8cb1b5990a0@www.phoenix.global:16789",
	"enode://0116873bb440fb2aeadbb402e6093f5b97f0574cec02c53db6ff99cc5c2dc6dc4ee578961640ed80f9f5b69605493c3a88c8137eb06497596fe960b4c4a9053e@www.phoenix.global:16789",
	"enode://8c816c9441f3b6d64efe2382110d17cd1ccd10a4a60db742383c36e6cabaeef2a2f87815ecd0f285e74ad6f9911a2db1dde7fe33a249e9c625ce225b0dbe25f0@www.phoenix.global:16789",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the test network.
var TestnetBootnodes = []string{
	"enode://33761aca567c3ce253635bfea65cb48d5518eaabfefd92e3d443414ea31a1abfff5dfda1f0c47f6b4cab0efd55a535a268acd82b3a9d078b40e328b890f49291@127.0.0.1:16789",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{}
