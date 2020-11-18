package proof

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/bandprotocol/chain/x/oracle/types"
)

func hexToBytes(hexstr string) []byte {
	b, err := hex.DecodeString(hexstr)
	if err != nil {
		panic(err)
	}
	return b
}

func leafHash(item []byte) []byte {
	// leaf prefix is 0
	return tmhash.Sum(append([]byte{0}, item...))
}

func branchHash(left, right []byte) []byte {
	// branch prefix is 1
	return tmhash.Sum(append([]byte{1}, append(left, right...)...))
}

func TestEncodeRelay(t *testing.T) {
	block := BlockRelayProof{
		MultiStoreProof: MultiStoreProof{
			AccToGovStoresMerkleHash:          hexToBytes("685430546D23A44E6B8034EAAFBC2F4CD7FEF54B54D5B66528CB4E5225BD74FB"),
			MainAndMintStoresMerkleHash:       hexToBytes("4F8D0BB0CD3EB9DC70B4DBBEA5F0CBD5B523195F7BAE02BB401BB00A93ABA08E"),
			OracleIAVLStateHash:               hexToBytes("1912057FFF0B3E85ABF1A319F75D37B21B430F3DA7DB9E486A5041DE47C686D3"),
			ParamsStoresMerkleHash:            hexToBytes("B1F2FD852E790E735CA2D3014F96A2A53C60393E9C6BBF941B9A6DD6A05CF6F9"),
			SlashingToUpgradeStoresMerkleHash: hexToBytes("91CC906286235B676AD402FC04ED768EB2BFECA664E8D595C286571DA1433C60"),
		},
		BlockHeaderMerkleParts: BlockHeaderMerkleParts{
			VersionAndChainIdHash:             hexToBytes("32FA694879095840619F5E49380612BD296FF7E950EAFB66FF654D99CA70869E"),
			TimeHash:                          hexToBytes("D9F175396C0E2D0E77F69856ABF5D8E69283CB915EB8886262FEDA1D519B3005"),
			LastBlockIDAndOther:               hexToBytes("4F4D548668A3986DB253689234B9CAC96303A128B7081B16E53B79AB9E65B887"),
			NextValidatorHashAndConsensusHash: hexToBytes("004209A161040AB1778E2F2C00EE482F205B28EFBA439FCB04EA283F619478D9"),
			LastResultsHash:                   hexToBytes("6E340B9CFFB37A989CA544E6BB780A2C78901D3FB33738768511A30617AFA01D"),
			EvidenceAndProposerHash:           hexToBytes("D991DA4D4E69473CC75A4B819F9E07D4956671A6F4A74DF4CC16596FCBE68137"),
		},
		Signatures: []TMSignature{
			{
				R:                hexToBytes("9279257914BFEE6FAEC46DF086E9A673A4C1576FB299094D45F77E12FA3728E7"),
				S:                hexToBytes("04DBF23F5EBB07BA8FD7CA06843DFE363B5E86596930E1889D9BD5BD3E98FC53"),
				V:                28,
				SignedDataPrefix: hexToBytes("6E080211BF0000000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A2066424E8F0417945A71067A55B7121282A90524E3C0709D8F8ADDB6ADC8FD46D110012A0C08E6E9C6F70510979FBD9801320962616E64636861696E"),
			},
			{
				R:                hexToBytes("826BB17B714EBCD8199EE2A01334102F19248087CFDEEE42EBD406B3991C3895"),
				S:                hexToBytes("621281EEDF97F3A9EC121224CEF9F8C07872D2C81CCE44DD846BB3678DD50523"),
				V:                27,
				SignedDataPrefix: hexToBytes("6E080211BF0000000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A2066424E8F0417945A71067A55B7121282A90524E3C0709D8F8ADDB6ADC8FD46D110012A0C08E6E9C6F70510DD89E98A01320962616E64636861696E"),
			},
			{
				R:                hexToBytes("D775FD0E1580499EF16A4AB1998DCB4CBD47CF6F342CDC51A9552D1434552ED8"),
				S:                hexToBytes("5A1A2075AE97A6BBD07A5A40EFE287ECEAB37C7A7CAAE82E6C997C51C6233470"),
				V:                27,
				SignedDataPrefix: hexToBytes("6E080211BF0000000000000022480A20"),
				SignedDataSuffix: hexToBytes("12240A2066424E8F0417945A71067A55B7121282A90524E3C0709D8F8ADDB6ADC8FD46D110012A0C08E6E9C6F70510AB94A59201320962616E64636861696E"),
			},
		},
	}
	result, err := block.encodeToEthData(191)
	require.Nil(t, err)
	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000000bf685430546d23a44e6b8034eaafbc2f4cd7fef54b54d5b66528cb4e5225bd74fb4f8d0bb0cd3eb9dc70b4dbbea5f0cbd5b523195f7bae02bb401bb00a93aba08e1912057fff0b3e85abf1a319f75d37b21b430f3da7db9e486a5041de47c686d3b1f2fd852e790e735ca2d3014f96a2a53c60393e9c6bbf941b9a6dd6a05cf6f991cc906286235b676ad402fc04ed768eb2bfeca664e8d595c286571da1433c6032fa694879095840619f5e49380612bd296ff7e950eafb66ff654d99ca70869ed9f175396c0e2d0e77f69856abf5d8e69283cb915eb8886262feda1d519b30054f4d548668a3986db253689234b9cac96303a128b7081b16e53b79ab9e65b887004209a161040ab1778e2f2c00ee482f205b28efba439fcb04ea283f619478d96e340b9cffb37a989ca544e6bb780a2c78901d3fb33738768511a30617afa01dd991da4d4e69473cc75a4b819f9e07d4956671a6f4a74df4cc16596fcbe6813700000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001a000000000000000000000000000000000000000000000000000000000000002e09279257914bfee6faec46df086e9a673a4c1576fb299094d45f77e12fa3728e704dbf23f5ebb07ba8fd7ca06843dfe363b5e86596930e1889d9bd5bd3e98fc53000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000106e080211bf0000000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003f12240a2066424e8f0417945a71067a55b7121282a90524e3c0709d8f8addb6adc8fd46d110012a0c08e6e9c6f70510979fbd9801320962616e64636861696e00826bb17b714ebcd8199ee2a01334102f19248087cfdeee42ebd406b3991c3895621281eedf97f3a9ec121224cef9f8c07872d2c81cce44dd846bb3678dd50523000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000106e080211bf0000000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003f12240a2066424e8f0417945a71067a55b7121282a90524e3c0709d8f8addb6adc8fd46d110012a0c08e6e9c6f70510dd89e98a01320962616e64636861696e00d775fd0e1580499ef16a4ab1998dcb4cbd47cf6f342cdc51a9552d1434552ed85a1a2075ae97a6bbd07a5a40efe287eceab37c7a7caae82e6c997c51c6233470000000000000000000000000000000000000000000000000000000000000001b00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000106e080211bf0000000000000022480a2000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003f12240a2066424e8f0417945a71067a55b7121282a90524e3c0709d8f8addb6adc8fd46d110012a0c08e6e9c6f70510ab94a59201320962616e64636861696e00")
	require.Equal(t, expect, result)
}

func TestEncodeVerify(t *testing.T) {
	data := OracleDataProof{
		Version: 180,
		RequestPacket: types.OracleRequestPacketData{
			ClientID:       "test",
			OracleScriptID: 1,
			Calldata:       hexToBytes("000000034254430000000000000064"),
			AskCount:       4,
			MinCount:       4,
		},
		ResponsePacket: types.OracleResponsePacketData{
			ClientID:      "test",
			RequestID:     1,
			AnsCount:      4,
			RequestTime:   1592898765,
			ResolveTime:   1592898773,
			ResolveStatus: 1,
			Result:        hexToBytes("00000000000eb9e6"),
		},
		MerklePaths: []IAVLMerklePath{
			{
				IsDataOnRight:  true,
				SubtreeHeight:  1,
				SubtreeSize:    2,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("108873CE6589ADD98E40996E7EA55D6592CFEE6EBF140CFBE7F0C7EAF554B8BD"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  2,
				SubtreeSize:    3,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("796AEB2094F52E848A9C9ACC57C380F92010FA19956CCF9E2C6E8E7D1E490F0C"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  3,
				SubtreeSize:    5,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("423F98BE29BD72613A6815E696ECF229E2B489198E9385D3BF1F743B99CA1573"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  4,
				SubtreeSize:    8,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("386E60DB569E57EAF6F2DBFA0E87AC5DAFB849EBF1DB7C6B29C0231D644AD92E"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  5,
				SubtreeSize:    21,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("465252AF4FE160AD23C446CE54DFC3EECAD65B96BDA8C68871A244033424616A"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  6,
				SubtreeSize:    41,
				SubtreeVersion: 180,
				SiblingHash:    hexToBytes("203800627751EE71822B76CF8C24F806D2528DC82247F5A49E42AF9DF04ADC65"),
			},
			{
				IsDataOnRight:  true,
				SubtreeHeight:  7,
				SubtreeSize:    77,
				SubtreeVersion: 190,
				SiblingHash:    hexToBytes("68B4E7B02288DCE3C92C590CF09C96B64D18FF48C92D57D1F14D877C6A2720F9"),
			},
		},
	}

	result, err := data.encodeToEthData(191)
	require.Nil(t, err)
	expect := hexToBytes("00000000000000000000000000000000000000000000000000000000000000bf00000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000001c000000000000000000000000000000000000000000000000000000000000000b4000000000000000000000000000000000000000000000000000000000000032000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000f000000034254430000000000000064000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000005ef1b4cd000000000000000000000000000000000000000000000000000000005ef1b4d50000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000047465737400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000eb9e6000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000b4108873ce6589add98e40996e7ea55d6592cfee6ebf140cfbe7f0c7eaf554b8bd00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000b4796aeb2094f52e848a9c9acc57c380f92010fa19956ccf9e2c6e8e7d1e490f0c00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000b4423f98be29bd72613a6815e696ecf229e2b489198e9385d3bf1f743b99ca157300000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000b4386e60db569e57eaf6f2dbfa0e87ac5dafb849ebf1db7c6b29c0231d644ad92e00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000001500000000000000000000000000000000000000000000000000000000000000b4465252af4fe160ad23c446ce54dfc3eecad65b96bda8c68871a244033424616a00000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000002900000000000000000000000000000000000000000000000000000000000000b4203800627751ee71822b76cf8c24f806d2528dc82247f5a49e42af9df04adc6500000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000004d00000000000000000000000000000000000000000000000000000000000000be68b4e7b02288dce3c92c590cf09c96b64d18ff48c92d57d1f14d877c6a2720f9")
	require.Equal(t, expect, result)
}
