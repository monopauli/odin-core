package proof

// import (
// 	"encoding/hex"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// 	"github.com/tendermint/tendermint/crypto/tmhash"

// 	"github.com/GeoDB-Limited/odin-core/x/oracle/types"
// )

// func hexToBytes(hexstr string) []byte {
// 	b, err := hex.DecodeString(hexstr)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return b
// }

// func leafHash(item []byte) []byte {
// 	// leaf prefix is 0
// 	return tmhash.Sum(append([]byte{0}, item...))
// }

// func branchHash(left, right []byte) []byte {
// 	// branch prefix is 1
// 	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
// }

// func TestEncodeRelay(t *testing.T) {
// 	block := BlockRelayProof{
// 		MultiStoreProof: MultiStoreProof{
// 			AccToGovStoresMerkleHash:          hexToBytes("0D4AF3F5FFA02A56B1DEED7BC8C16732AEB8FD003C67EEF26048B314C351FAE8"),
// 			MainAndMintStoresMerkleHash:       hexToBytes("68BF06D17DBF1F5870D3092E1433A99FDAF6E263EFD5F8C82C533691D87592B7"),
// 			OracleIAVLStateHash:               hexToBytes("4F900B8B425CF85AB2A1ED2907D4830BF674703C64954847DAAEE9B81A09BA31"),
// 			ParamsStoresMerkleHash:            hexToBytes("2B6A7E0F44ED9C179A47A40F93D5824189A5426D6C3F77692DE28E50E20A33DD"),
// 			SlashingToUpgradeStoresMerkleHash: hexToBytes("F5E26E9E91F18051F41453B5A5FC82279C364351155CEA5564B9D61BC12BE58A"),
// 		},
// 		BlockHeaderMerkleParts: BlockHeaderMerkleParts{
// 			VersionAndChainIdHash:             hexToBytes("3561783E9C3BDF932A16580FC0C9CEFFEC4C509073FFF65A42871BFAB64408AE"),
// 			Height:                            3021518,
// 			TimeSecond:                        1605721438,
// 			TimeNanoSecond:                    605059026,
// 			LastBlockIDAndOther:               hexToBytes("21114E3076A55C6853B4730FB8678B5BF2314C1DF6DCE169ACEE9AECE893C60F"),
// 			NextValidatorHashAndConsensusHash: hexToBytes("EA01CD62E714B603323A21A4A7382F8AB04788C867A0C99ADE687D00E7D5FE62"),
// 			LastResultsHash:                   hexToBytes("AA3C7CBEFF135291E6415ECA2528FC98D263B120C67BCECD8D8CCD3253EFECC1"),
// 			EvidenceAndProposerHash:           hexToBytes("68D9EF5EB2AFAF2E36940299C8CDA2F43ACB015FC2D6CAFD2C577CA48F1B2C26"),
// 		},
// 		Signatures: []TMSignature{
// 			{
// 				R:                hexToBytes("D090EF654A5C8B59EB97346EB46E601A6D57DDA9174C1E535C40485FD5A8D414"),
// 				S:                hexToBytes("47A7B3E582103C2B28B8D3D3C8D9B7D715D852E449711C7B91BDF9EDCB3F7ADA"),
// 				V:                28,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510DAD8C69B02320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("26ED90EE89D4F6B5D172904DDB82A18419E6833BFC74522416800E3A7D2E3AE0"),
// 				S:                hexToBytes("0C2DF3F307106861195AA41643CF6DA1F5BE195B11C74B9C90329F99B1E24CA7"),
// 				V:                27,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A9F1DCC402320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("5CA760BDA037610B623B83CBF8E0F55FECA37A9A621EBAD17360E5725F8E3425"),
// 				S:                hexToBytes("3DF55E237CB9D1EAE87243CD4D3F5F1091B9FF2E2F16A2B083928F2B0F09C7BB"),
// 				V:                28,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A196B3C502320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("37466D99BB8ED9EA462B6A7E988189224A628FFA973A05A29BB401C4C4FC9B2D"),
// 				S:                hexToBytes("7FD076B0E28B7E6A35CC80AD6FA688958E9F22093CE00C22D4FE67124C633B04"),
// 				V:                28,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD05108BFF92C502320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("99CFB8E6848F039863D98C48AC985F4AD6A27FE0B602A8735D44654DF1B06DE2"),
// 				S:                hexToBytes("75C7162F5ABADA35FB48460F2196E5348334E98BA05DE1236D59DB329EC2E4AA"),
// 				V:                27,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD05108E8BADC302320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("780733F9013D10F88A2FA936BC2789A61146749B779FB847C0F46279CA31543E"),
// 				S:                hexToBytes("28EB691288B02C06D29CECC3DDEA3387CB10F69AEC9F6C0A51ED39784D6675D0"),
// 				V:                27,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510F0CDC5D402320F62616E642D6775616E79752D706F61"),
// 			},
// 			{
// 				R:                hexToBytes("2914567728CAE2ABE2B707933AE63660C7C79786C7B897CB867B68F67AAC07A8"),
// 				S:                hexToBytes("04F683B20F6B043A0CA79C39DD68167869262BF262CEF1B1CE8E30B4A06B0D5D"),
// 				V:                27,
// 				SignedDataPrefix: hexToBytes("74080211CE1A2E000000000022480A20"),
// 				SignedDataSuffix: hexToBytes("12240A202DC401225B681224CB8F597D157A5DE78EF4F04FE1C884595F2B18D941EBCA2010012A0C08E1BAD5FD0510A7CEC19A02320F62616E642D6775616E79752D706F61"),
// 			},
// 		},
// 	}
// 	result, err := block.encodeToEthData()
// 	require.Nil(t, err)
// 	expect := hexToBytes("0d4af3f5ffa02a56b1deed7bc8c16732aeb8fd003c67eef26048b314c351fae868bf06d17dbf1f5870d3092e1433a99fdaf6e263efd5f8c82c533691d87592b74f900b8b425cf85ab2a1ed2907d4830bf674703c64954847daaee9b81a09ba312b6a7e0f44ed9c179a47a40f93d5824189a5426d6c3f77692de28e50e20a33ddf5e26e9e91f18051f41453b5a5fc82279c364351155cea5564b9d61bc12be58a3561783e9c3bdf932a16580fc0c9ceffec4c509073fff65a42871bfab64408ae00000000000000000000000000000000000000000000000000000000002e1ace000000000000000000000000000000000000000000000000000000005fb55d5e00000000000000000000000000000000000000000000000000000000241077d221114e3076a55c6853b4730fb8678b5bf2314c1df6dce169acee9aece893c60fea01cd62e714b603323a21a4a7382f8ab04788c867a0c99ade687d00e7d5fe62aa3c7cbeff135291e6415eca2528fc98d263b120c67bcecd8d8ccd3253efecc168d9ef5eb2afaf2e36940299c8cda2f43acb015fc2d6cafd2c577ca48f1b2c2600000000000000000000000000000000000000000000000000000000000001c0000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000003a00000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000066000000000000000000000000000000000000000000000000000000000000007c00000000000000000000000000000000000000000000000000000000000000920d090ef654a5c8b59eb97346eb46e601a6d57dda9174c1e535c40485fd5a8d41447a7b3e582103c2b28b8d3d3c8d9b7d715d852e449711c7b91bdf9edcb3f7ada000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510dad8c69b02320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000026ed90ee89d4f6b5d172904ddb82a18419e6833bfc74522416800e3a7d2e3ae00c2df3f307106861195aa41643cf6da1f5be195b11c74b9c90329f99b1e24ca7000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a9f1dcc402320f62616e642d6775616e79752d706f610000000000000000000000000000000000000000000000000000005ca760bda037610b623b83cbf8e0f55feca37a9a621ebad17360e5725f8e34253df55e237cb9d1eae87243cd4d3f5f1091b9ff2e2f16a2b083928f2b0f09c7bb000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a196b3c502320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000037466d99bb8ed9ea462b6a7e988189224a628ffa973a05a29bb401c4c4fc9b2d7fd076b0e28b7e6a35cc80ad6fa688958e9f22093ce00c22d4fe67124c633b04000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd05108bff92c502320f62616e642d6775616e79752d706f6100000000000000000000000000000000000000000000000000000099cfb8e6848f039863d98c48ac985f4ad6a27fe0b602a8735d44654df1b06de275c7162f5abada35fb48460f2196e5348334e98ba05de1236d59db329ec2e4aa000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd05108e8badc302320f62616e642d6775616e79752d706f61000000000000000000000000000000000000000000000000000000780733f9013d10f88a2fa936bc2789a61146749b779fb847c0f46279ca31543e28eb691288b02c06d29cecc3ddea3387cb10f69aec9f6c0a51ed39784d6675d0000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510f0cdc5d402320f62616e642d6775616e79752d706f610000000000000000000000000000000000000000000000000000002914567728cae2abe2b707933ae63660c7c79786c7b897cb867b68f67aac07a804f683b20f6b043a0ca79c39dd68167869262bf262cef1b1ce8e30b4a06b0d5d000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000001074080211ce1a2e000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004512240a202dc401225b681224cb8f597d157a5de78ef4f04fe1c884595f2b18d941ebca2010012a0c08e1bad5fd0510a7cec19a02320f62616e642d6775616e79752d706f61000000000000000000000000000000000000000000000000000000")

// 	require.Equal(t, expect, result)
// }

// func TestEncodeVerify(t *testing.T) {
// 	data := OracleDataProof{
// 		Version: 180,
// 		RequestPacket: types.OracleRequestPacketData{
// 			ClientID:       "test",
// 			OracleScriptID: 1,
// 			Calldata:       hexToBytes("000000034254430000000000000064"),
// 			AskCount:       4,
// 			MinCount:       4,
// 		},
// 		ResponsePacket: types.OracleResponsePacketData{
// 			ClientID:      "test",
// 			RequestID:     1,
// 			AnsCount:      4,
// 			RequestTime:   1592898765,
// 			ResolveTime:   1592898773,
// 			ResolveStatus: 1,
// 			Result:        hexToBytes("00000000000eb9e6"),
// 		},
// 		MerklePaths: []IAVLMerklePath{
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  1,
// 				SubtreeSize:    2,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("108873CE6589ADD98E40996E7EA55D6592CFEE6EBF140CFBE7F0C7EAF554B8BD"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  2,
// 				SubtreeSize:    3,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("796AEB2094F52E848A9C9ACC57C380F92010FA19956CCF9E2C6E8E7D1E490F0C"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  3,
// 				SubtreeSize:    5,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("423F98BE29BD72613A6815E696ECF229E2B489198E9385D3BF1F743B99CA1573"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  4,
// 				SubtreeSize:    8,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("386E60DB569E57EAF6F2DBFA0E87AC5DAFB849EBF1DB7C6B29C0231D644AD92E"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  5,
// 				SubtreeSize:    21,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("465252AF4FE160AD23C446CE54DFC3EECAD65B96BDA8C68871A244033424616A"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  6,
// 				SubtreeSize:    41,
// 				SubtreeVersion: 180,
// 				SiblingHash:    hexToBytes("203800627751EE71822B76CF8C24F806D2528DC82247F5A49E42AF9DF04ADC65"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  7,
// 				SubtreeSize:    77,
// 				SubtreeVersion: 190,
// 				SiblingHash:    hexToBytes("68B4E7B02288DCE3C92C590CF09C96B64D18FF48C92D57D1F14D877C6A2720F9"),
// 			},
// 		},
// 	}

// 	result, err := data.encodeToEthData(191)
// 	require.Nil(t, err)
// 	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000000bf00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000b4000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f000000034254430000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000005ef1b4cd000000000000000000000000000000000000000000000000000000005ef1b4d50000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000eb9e6000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000b4108873ce6589add98e40996e7ea55d6592cfee6ebf140cfbe7f0c7eaf554b8bd00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000b4796aeb2094f52e848a9c9acc57c380f92010fa19956ccf9e2c6e8e7d1e490f0c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000b4423f98be29bd72613a6815e696ecf229e2b489198e9385d3bf1f743b99ca157300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000b4386e60db569e57eaf6f2dbfa0e87ac5dafb849ebf1db7c6b29c0231d644ad92e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001500000000000000000000000000000000000000000000000000000000000000b4465252af4fe160ad23c446ce54dfc3eecad65b96bda8c68871a244033424616a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002900000000000000000000000000000000000000000000000000000000000000b4203800627751ee71822b76cf8c24f806d2528dc82247f5a49e42af9df04adc6500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000004d00000000000000000000000000000000000000000000000000000000000000be68b4e7b02288dce3c92c590cf09c96b64d18ff48c92d57d1f14d877c6a2720f9")
// 	require.Equal(t, expect, result)
// }

// func TestEncodeVerifyCount(t *testing.T) {
// 	data := RequestsCountProof{
// 		Count:   1372854,
// 		Version: 3213160,
// 		MerklePaths: []IAVLMerklePath{
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  1,
// 				SubtreeSize:    2,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("8142C1671AAE5994E3F38D76950E01918C16081365DCD695B38A801391FC25CB"),
// 			},
// 			{
// 				IsDataOnRight:  true,
// 				SubtreeHeight:  2,
// 				SubtreeSize:    4,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("29AC504248FF3C4593803A107DE1F8CFCA59D92D5010F84B51173619ABE09FB8"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  3,
// 				SubtreeSize:    6,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("5DE57CC701C2910DD50D047C4ECDF451A63411C2C88910584FDE650B47807DD6"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  4,
// 				SubtreeSize:    12,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("2D10ED6A267545CADD9F15948E6F60CF56683CFABADD51621E1107D6D3615A6D"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  5,
// 				SubtreeSize:    18,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("060B4736F4CF280889E697EAEBFBA07933BE03206B122F5425914E7392B6A5F0"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  6,
// 				SubtreeSize:    36,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("00E029F6CD33AC0E43B7C02BA3CA706FF260E82790423958711A0D646F33A002"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  7,
// 				SubtreeSize:    54,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("1820D392C6F545D8B4A0F6ADDBF5CF47222715A83845EA5BA5EC0E1A8C522952"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  8,
// 				SubtreeSize:    126,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("D9303F5DD61104A6DCEC483AA4B7E29DDC9DAE4DC60ACB513892229B30282A7D"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  9,
// 				SubtreeSize:    198,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("A71FFCB1D79742229D571AC7A34CB505128DC8E1AC95D9DCE0AAF991CCCC0BB2"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  10,
// 				SubtreeSize:    414,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("7E56CABDB251D8D8C4DE99E36B21AF4BC937B3D931886D9338C084F1F6B6AFFF"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  11,
// 				SubtreeSize:    630,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("82E7252442A4C9BF471E686B52EE0176BAF70FDB90E0A4A190D059C9DD794E05"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  12,
// 				SubtreeSize:    1278,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("3AB09B4A90C9105FDF210837FD01CED1C30D37C9A0AF3696ADBBAC910DBACC69"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  13,
// 				SubtreeSize:    1926,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("9410500A015FD7513164CE5A0B3BFC0C56AB6721B4ECBA22EFF5C78FB40D78E9"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  14,
// 				SubtreeSize:    3870,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("7A5B1DD4BAE4D11AE6A7970B31961140AE9BEF572B1FF44D3C3C1A858DBF464C"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  15,
// 				SubtreeSize:    7758,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("C74AF02CCFC053002BFA7DC675B71D2398EEF1A70B2D03AECD88B49264F46E71"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  16,
// 				SubtreeSize:    15534,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("208D7325C9E5D943FCC9A0B24718437C9381F6FABC4731E3859E6688ED3F9C1F"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  17,
// 				SubtreeSize:    23310,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("6FEBF3C129E337BA04DFA019F04BEBEDA000D1AA460AA2FA2E7E90AC9892DB9F"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  18,
// 				SubtreeSize:    46638,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("E9B4F48E1A6919E12E27E4D75ACC52E41118006C261270A65EBEE75D4E70AAD9"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  19,
// 				SubtreeSize:    93294,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("F68039D9B3F9C0385FBB339CE35B36F9B63472050D0E8750ECB03ADC8C9CC232"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  20,
// 				SubtreeSize:    186606,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("68BAFB5C28F4CF18BAFBBB846FDC9431DC0530B2A966BF1BA37CCF1E301E0627"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  21,
// 				SubtreeSize:    373230,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("9CB7C561D0C8BC006C2EDBC3BA2251C8765F877DB583BCF12AFA94B8B6F5834F"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  22,
// 				SubtreeSize:    746478,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("0FA89300B33B4CE28BD319EA32D47198DE9C5110804FB9E9FA434D5972B4C35A"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  23,
// 				SubtreeSize:    1410607,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("B432978151E364DCC206D5FA0044C333D5D185BE6BA7CE2FF09B45DE90A500DC"),
// 			},
// 			{
// 				IsDataOnRight:  false,
// 				SubtreeHeight:  25,
// 				SubtreeSize:    8292144,
// 				SubtreeVersion: 3213160,
// 				SiblingHash:    hexToBytes("34999511AD8232F700A9CD267C7CE64C12BDD862A56BE47A94EC342892A6A4E5"),
// 			},
// 		},
// 	}

// 	result, err := data.encodeToEthData(3213161)
// 	require.Nil(t, err)
// 	expect := hexToBytes("0000000000000000000000000000000000000000000000000000000000310769000000000000000000000000000000000000000000000000000000000014f2b600000000000000000000000000000000000000000000000000000000003107680000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000003107688142c1671aae5994e3f38d76950e01918c16081365dcd695b38a801391fc25cb000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000031076829ac504248ff3c4593803a107de1f8cfca59d92d5010f84b51173619abe09fb800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000003107685de57cc701c2910dd50d047c4ecdf451a63411c2c88910584fde650b47807dd600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000003107682d10ed6a267545cadd9f15948e6f60cf56683cfabadd51621e1107d6d3615a6d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000120000000000000000000000000000000000000000000000000000000000310768060b4736f4cf280889e697eaebfba07933be03206b122f5425914e7392b6a5f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000060000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000031076800e029f6cd33ac0e43b7c02ba3ca706ff260e82790423958711a0d646f33a00200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000003600000000000000000000000000000000000000000000000000000000003107681820d392c6f545d8b4a0f6addbf5cf47222715a83845ea5ba5ec0e1a8c52295200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000007e0000000000000000000000000000000000000000000000000000000000310768d9303f5dd61104a6dcec483aa4b7e29ddc9dae4dc60acb513892229b30282a7d0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000000000c60000000000000000000000000000000000000000000000000000000000310768a71ffcb1d79742229d571ac7a34cb505128dc8e1ac95d9dce0aaf991cccc0bb20000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000019e00000000000000000000000000000000000000000000000000000000003107687e56cabdb251d8d8c4de99e36b21af4bc937b3d931886d9338c084f1f6b6afff0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b0000000000000000000000000000000000000000000000000000000000000276000000000000000000000000000000000000000000000000000000000031076882e7252442a4c9bf471e686b52ee0176baf70fdb90e0a4a190d059c9dd794e050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000c00000000000000000000000000000000000000000000000000000000000004fe00000000000000000000000000000000000000000000000000000000003107683ab09b4a90c9105fdf210837fd01ced1c30d37c9a0af3696adbbac910dbacc690000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000d000000000000000000000000000000000000000000000000000000000000078600000000000000000000000000000000000000000000000000000000003107689410500a015fd7513164ce5a0b3bfc0c56ab6721b4ecba22eff5c78fb40d78e90000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000f1e00000000000000000000000000000000000000000000000000000000003107687a5b1dd4bae4d11ae6a7970b31961140ae9bef572b1ff44d3c3c1a858dbf464c0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f0000000000000000000000000000000000000000000000000000000000001e4e0000000000000000000000000000000000000000000000000000000000310768c74af02ccfc053002bfa7dc675b71d2398eef1a70b2d03aecd88b49264f46e71000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000003cae0000000000000000000000000000000000000000000000000000000000310768208d7325c9e5d943fcc9a0b24718437c9381f6fabc4731e3859e6688ed3f9c1f000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000110000000000000000000000000000000000000000000000000000000000005b0e00000000000000000000000000000000000000000000000000000000003107686febf3c129e337ba04dfa019f04bebeda000d1aa460aa2fa2e7e90ac9892db9f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000b62e0000000000000000000000000000000000000000000000000000000000310768e9b4f48e1a6919e12e27e4d75acc52e41118006c261270a65ebee75d4e70aad9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000130000000000000000000000000000000000000000000000000000000000016c6e0000000000000000000000000000000000000000000000000000000000310768f68039d9b3f9c0385fbb339ce35b36f9b63472050d0e8750ecb03adc8c9cc23200000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000002d8ee000000000000000000000000000000000000000000000000000000000031076868bafb5c28f4cf18bafbbb846fdc9431dc0530b2a966bf1ba37ccf1e301e062700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000015000000000000000000000000000000000000000000000000000000000005b1ee00000000000000000000000000000000000000000000000000000000003107689cb7c561d0c8bc006c2edbc3ba2251c8765f877db583bcf12afa94b8b6f5834f0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000b63ee00000000000000000000000000000000000000000000000000000000003107680fa89300b33b4ce28bd319ea32d47198de9c5110804fb9e9fa434d5972b4c35a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000017000000000000000000000000000000000000000000000000000000000015862f0000000000000000000000000000000000000000000000000000000000310768b432978151e364dcc206d5fa0044c333d5d185be6ba7ce2ff09b45de90a500dc0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001900000000000000000000000000000000000000000000000000000000007e8730000000000000000000000000000000000000000000000000000000000031076834999511ad8232f700a9cd267c7ce64c12bdd862a56be47a94ec342892a6a4e5")
// 	require.Equal(t, expect, result)
// }
