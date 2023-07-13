package rlp_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/INFURA/go-ethlibs/rlp"
)

func TestFrom(t *testing.T) {
	{
		// This is a list (actually a raw transaction)
		input := "0xf86d820144843b9aca0082520894b78777860637d56543da23312c7865024833f7d188016345785d8a0000802ba0e2539a5d9f056d7095bd19d6b77b850910eeafb71534ebd45159915fab202e91a007484420f3968697974413fc55d1142dc76285d30b1b9231ccb71ed1e720faae"

		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, 9, len(decoded.List))

		payload := decoded.List
		require.Equal(t, "0x0144", payload[0].String, "Nonce")
		require.Equal(t, "0x3b9aca00", payload[1].String, "Price")
		require.Equal(t, "0x5208", payload[2].String, "GasLimit")
		require.Equal(t, "0xb78777860637d56543da23312c7865024833f7d1", payload[3].String, "To")
		require.Equal(t, "0x016345785d8a0000", payload[4].String, "Amount")
		require.Equal(t, "0x", payload[5].String, "Data")
		require.Equal(t, "0x2b", payload[6].String, "V")
		require.Equal(t, "0xe2539a5d9f056d7095bd19d6b77b850910eeafb71534ebd45159915fab202e91", payload[7].String, "R")
		require.Equal(t, "0x07484420f3968697974413fc55d1142dc76285d30b1b9231ccb71ed1e720faae", payload[8].String, "S")

		encoded, err := decoded.Encode()
		require.NoError(t, err)
		require.Equal(t, input, encoded)
	}

	// Some examples from https://github.com/ethereum/tests/blob/develop/RLPTests/rlptest.json
	//  Copyright 2014 Ethereum Foundation - MIT Licensed
	{
		// multilist: [ "zw", [ 4 ], 1 ]
		input := "0xc6827a77c10401"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, 3, len(decoded.List))
		require.Equal(t, "0x7a77", decoded.List[0].String)
		require.Equal(t, "0x04", decoded.List[1].List[0].String)
		require.Equal(t, "0x01", decoded.List[2].String)
	}

	{
		// listsoflists: [ [ [], [] ], [] ]
		input := "0xc4c2c0c0c0"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, 2, len(decoded.List))
		require.Equal(t, 2, len(decoded.List[0].List))
		require.Equal(t, 0, len(decoded.List[0].List[0].List))
		require.Equal(t, 0, len(decoded.List[0].List[1].List))
		require.Equal(t, 0, len(decoded.List[1].List))
	}

	{
		// shortListMax1: [ "asdf", "qwer", "zxcv", "asdf","qwer", "zxcv", "asdf", "qwer", "zxcv", "asdf", "qwer"]
		input := "0xf784617364668471776572847a78637684617364668471776572847a78637684617364668471776572847a78637684617364668471776572"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, 11, len(decoded.List))
		require.Equal(t, "0x61736466", decoded.List[0].String)
		require.Equal(t, "0x71776572", decoded.List[1].String)
		require.Equal(t, "0x7a786376", decoded.List[2].String)
		require.Equal(t, decoded.List[0].String, decoded.List[3].String)
		require.Equal(t, decoded.List[1].String, decoded.List[4].String)
		require.Equal(t, decoded.List[2].String, decoded.List[5].String)
		require.Equal(t, decoded.List[3].String, decoded.List[6].String)
		require.Equal(t, decoded.List[4].String, decoded.List[7].String)
		require.Equal(t, decoded.List[5].String, decoded.List[8].String)
		require.Equal(t, decoded.List[6].String, decoded.List[9].String)
		require.Equal(t, decoded.List[7].String, decoded.List[10].String)
	}

	{
		// bigint: #115792089237316195423570985008687907853269984665640564039457584007913129639936
		input := "0xa1010000000000000000000000000000000000000000000000000000000000000000"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, "0x010000000000000000000000000000000000000000000000000000000000000000", decoded.String)
	}

	{
		// block
		input := "0xf913cff90215a012150305917c903204f0ea6bcde99c163782711243118cb5fadfeec77f5ce2fba01dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347945a0b54d5dc17e0aadc383d2db43b0a0d3e029c4ca0cbd1d7bf01807988d8a5ad9df3dfdb5f2fc3af72c268df5c6d4325ac4c825888a0c1ba37938ce5887b5b70efc371d993207b2bc05183f1a20dcff016b48ede53c0a0a3e28fd769b01c8d7e3bb1bfcf9386832c369b9bc06889541f887149cd25dcd5b90100800000300820148400021200280001a0458008083580801c00000688ca21002012028a1084008080c0300084210801001100000008081c0300a84800800e18201600000484ac12004800c02880005800406020490414102022002008508002001180c0491e0204483004008004100c8864830001c00109080002041c150400210000644000800001020100512a8080200010400c018003829520202000180080040e20444022240601091490400142013e041400a008c0010800220c00200043048c09024040002024000400000040001040624202081003412000101004602000c044008044004c0c3000000680020000450500022060888010001240180af68707df60d8014acb8393c55283989680839883af845e6fc9eb94737061726b706f6f6c2d6574682d636e2d687a32a006902382c3573b497e8705b2802b8fabfdd1a2cfd8691f4148f44454515bdf0288e3695db8087a93b8f911b3f8ad83073f9f851bf08eb000830222e09468d57c9a1c35f63e2c83ee8e49a64e9d70528d2580b844a9059cbb0000000000000000000000005310850866bbf6637223e222cf27db17cc0d7881000000000000000000000000000000000000000000000a968163f0a57b40000026a0ae6f192ba953388a12626b683770fcdf9aa414e63e3349db548fecad8351e2eea02b6b7a08fdb91def21b477d1e473109f15b9ff629665bb427004ad8a9d03fb47f8ad8308f40b851bf08eb000830222e094df574c24545e5ffecb9a659c229253d4111d87e180b844a9059cbb000000000000000000000000d72ae3c3a7c9819d1d6b411ebf561e660ebbb10800000000000000000000000000000000000000000000000000000b5e5c19670025a05bb59ef692702b02421eb05f92f74f1298eebd2e115fc3bb4489f2f4a38c7a19a06b5c19f4f5bd34a1e0efd0da80dec632151248aaa303bb2f7f8551b690cf53a9f86b0a85123f6c944482520894f8b67009f8d7bab795b15f08ccbf824ca7c12eb58734f3b61cf31bcc8025a046d86ac0d625690c62aa795a2373a3e41d398162df23adcd6a86bfdad3365871a03532229160981ebc80495061664a12960f9058988e65151b133436a039a3f634f86b348506942b576682520894777f415324d56e1d54fa832902d8797db7a4c57c871910540d9760008025a00b0df9605ec31ced5588825f8c2c33aeab9417553ce6dc2a38ed546a5f6256faa053bb9c79555e01c529f8283eb8f7dc9799ff72b687b3e80282e515629a5733eff8ad830173738503b9aca000830129b194dac17f958d2ee523a2206206994597c13d831ec780b844a9059cbb0000000000000000000000001d6cc8eefae3b3043e7d1503d6f0889df7496b2f0000000000000000000000000000000000000000000000000000000020bbacb126a0f15037623d1e09dddcc861594f56051805e1614e8350f82cf3b1027262fa254ea03e98a6c41e099649f7ee5470b9ee1104dbba833a4afbb15625623705bdf12f92f8ad830b59a685037e11d6008301d4c094543ff227f64aa17ea132bf9886cab5db55dcaddf80b844a9059cbb0000000000000000000000002db8c89cabe1735d41cdd4ed95ab3cb5afd8a60d0000000000000000000000000000000000000000000000196a79c202bf84900026a013329cc5e204895ed79ab0e74c86c26b723c223c88090a30899158cdc17bb1aaa046c70cc40b33e70e9aa899f4dc455dc57f0d39d49e5804f82c22f0355097f7a0f89080850218711a00830434b994bcf935d206ca32929e1b887a07ed240f0d8ccd22876a94d74f430000a48853b53e000000000000000000000000000000000000000000000000000000000003336125a0e73e5fa1991c0a34c219b9329c077cdb43d45dc159ccebd38e32ca4ccd33eddfa05dd531d5860ada18db7902378e3af68a27354bcdcbf934ca8fa1ca5a278aa7a8f901ae8301ce838501dce285008302035494fb80bfa19cae9e00f28b0f7e1023109deeb1048380b9014464887334000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000000093c55100000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000001496160eb6061f9c30a06fae400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c1a5a519a6a610ebf2fa001600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000025a036ba1645f9e307b233e522c7190b44cf9992e7a4818d22c8170fb7b84586bdf8a057e8897d8e79d69be730f4fc8113746b549453ce10fa455441a5164b8e542e53f87083090e8c8501dcd6500083015f9094de1b41d19898f26c100528b7e8ce506e3f06f0c488016345785d8a00008026a0d8c8322fb155cd6ccc8a1bca91f17991b0a6e473aa71923da29b6144d4ad7b5aa00e2869970daec601aebf07561bf42e3a0aa3d76af3c2210d56b2349a63d0fa5bf86c0a8501d0f5d70682520894be4e7e8e143fce9e60efca5fc901c3414894f9818841f6bbd999df90008025a0fb8b3e6dc9cfec01c0ccc209c08d0edc7e13a7e51760e88b27bcecff622869d7a00ba1a0f741fd8e4f57441fcf7322e48c98d8e83c3ff65b62a77f2b27126c3aa2f86d820cfb8501a13b860082520894e6c344e675aa9b5cbf6e2e3e164ad89ed90aed95870aa87bee5380008025a04c86b1f17aee531145cfde8514ffc8b1b9bd3d48ad835a66fecbbe990e1f0637a024179a13afef376cd9b86025f80f5b1211d8ad7cef0ab5d9c9f34b94cf10b093f86d8306093484b2d05e0082c35094e67b5a0fb559215b842f9e43195dfba4f56bb3a9878e1bc9bf0400008026a0f3359dce552ee26550d860c13a4f493c8b86f162dc88caa1a71af782e78de06aa0481b94e6e0eb7c762193d20c34d731a99140420bb2188e4b7697f8068ce82e22f902ea0984b2d05e008377fb3d94d737632cac4d039c9b0eecc94c12267407a271b580b90284885b48e70000000000000000000000000000000000000000000000000000000000000140000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000001c00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000006f05b59d3b2000000000000000000000000000000000000000000000000000001f161421c8e00000000000000000000000000000000000000000000000000000000000000015180000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000005434f5631390000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005434f5631390000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005636f76313900000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000006fc46b64e3de0a4962d15a75129f39a9068c6f10000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000004be4e7267b6ae0000025a060227541f9adbb11ab3c226bc63a5eaad3b948399eaa4ed7ad1474f9cd11720ca0549d933d2093d283bb43be424bb78b0522a1bf82add9fbc90f0b1de65fe7c0c2f9010960847735940083022fc69497dec872013f6b5fb443861090ad93154287812680b8a4ddf7e1a700000000000000000000000000000000000000000000000000000000095b7a30000000000000000000000000000000000000000000000007f678338c1cede5d00000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000005e6fcd5e0000000000000000000000006b175474e89094c44da98b954eedeac495271d0f26a0a712bbfdbe3a5267d27fd9c875c57c7ca15605da03f13bd6b6cbaa1241f8da2ca06bc984efbafe82df17cf7183e2a0a6bb7acd5ca5d67144cb0722b318e464ce4af88826844cef998c831ab7d7940000000000b3f879cb30fe243b4dfee438691c0480a4a0712d68000000000000000000000000000000000000000000000000000000000000002f25a0b327eacdf13fbe51b1d9447c0a5f439d2d1798428a92d3b744aeb2cb35429c78a060ef0261174356ef2bff737ffe3d27e05a83b2bf2063b90dd9a89739dce45878f9028c82179784495d9bb3830409c1945aee1922f69ba5d51c2fe0b5d68f184069f7f85b80b90224ec83d3ba000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000038f78f6865696d64616c6c2d4f316859424402830fafd580a01c4fc52e92ead3b1c07de84b2ca79f7e3965f7d6b6dc07b7600531448b03e2e3000000000000000000000000000000000000000000000000000000000000000000000000000000415e248b667407ab6a3251ef2697ea849dd124294ac87f456f3047d04a30ee6a6f2dda5d339f2c6dae5c42ebb24001f5393820d2b38ed06d6b7f879316c8af3498010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000acf8aaf8649461ceb3e10ee836cd3c431951d57e54580bf58e00831dc200831dc2ffa0f66d140fb2d69b485abab176a5068fa7020055a9e12c459c946dffc505e16627a05abb9ed277ccc1d045e4afd8b8e97ed7547d619b590c8cedcf8a21a9cdbb3bd1845e6fc86eb8415b31a56dce912cf62682215c8d66c7f61141f79b8692daa0a6fb4bbe8ec6b92e4ac069cc956f6604a048edda01c6315531d1b444c8e866716a35fe936eba3066008000000000000000000000000000000000000000001ca05f28dbf648cd989f109b1f335c2f4f627583dc3915790c6247e73b89961d0c03a00a1315ab63c9cad83ea8ec8906f2119542b9a57b922ae7788305932fe0ddd3b0f8700c84495d9bb383033450947a78e3d47c9e20836aca681956d8a972c0c793898802c68af0bb140000846254a0ef25a0d7624b1428abd024ce47d34a01bb81a9c3993003b11d6ae4af22bbcc32900d17a01bda2fd01aeb14f95107b1a702a861839f2efa78b11fcb8414133261998ea139f8700d84495d9bb383033450947a78e3d47c9e20836aca681956d8a972c0c7938988016345785d8a00008401073bf525a05951a37c5d9568ab6767268ee294f866006a6797a2e025625760e784a25c535fa0090b3303b65d27ed417ee14ee933658f579493777c2fe64cb8f6f1a124a24e25f88a820d9d84495d9bb38303c84f9483320bc47f06d73ac63e3c3809703529547c1db080a418c9bd5b000000000000000000000000000000000000000000000000000000000000000025a08d7816a7df785c47fef6a7ad57aea1cd6ef06381f96dc55a22e6a88c21d31946a00312ca92365d903f962cb8b47b6e3b1413560e6b5b1a44176475da5c9c71e5a6f9012a82021984495d9bb382c787945e07b6f1b98a11f7e04e7ffa8707b63f1c17775380b8c471b773440000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000005e6fc899000000000000000000000000000000000000000000000000000000000000004106869708744d97f3268409b2d6b445632c497243dcdd621858601666f089ecd2248d3c06f5c38b107e81b0f208bdcb493802b0dd176183e53df4ee2242ecdfb1010000000000000000000000000000000000000000000000000000000000000026a0e40a038329ad3009f6b7814a41bb72d51ff2fe2d5fb302533a247c32d562c117a03b13569110c5469b4baa92b1e5360c346dfb2cb9d1debe04cbb85a50278d8ef2f86b3384481f228082520894883f96c3ebb7cbc548770250477f5dce9e96d1fe8802c68af0bb1400008026a0020a0cd74e55ca6f2ce7fdd79f446d4b1d1fcbc542ee119a3e416b4ddfafb913a071ce34e500a605539fe47628b77b26cf1dd57bc5d5efe62ee4a7799ce233c8e0c0"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Len(t, decoded.List, 3)
	}

	{
		// one zero byte
		// NOTE: This is valid RLP but is not a valid representation of an Ethereum "scalar" because of this
		// rule in the yellow paper
		//
		// > When interpreting RLP data, if an expected fragment is decoded as a scalar and leading zeroes are found
		// > in the byte sequence, clients are required to consider it non-canonical and treat it in the same manner as
		// > otherwise invalid RLP data, dismissing it completely.
		input := "0x8100"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, "0x00", decoded.String)
	}

	{
		// zero.
		// NOTE: Like the above this is valid RLP but not a valid representation of an Ethereum scalar, which should
		// use 0x80 instead.
		input := "0x00"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, "0x00", decoded.String)
	}

	{
		// empty string
		input := "0x80"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, "0x", decoded.String)
	}

	{
		// "empty" RLP
		input := "0x"
		decoded, err := rlp.From(input)
		require.NoError(t, err)
		require.Equal(t, "0x", decoded.String)
	}
}