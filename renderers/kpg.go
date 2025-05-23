package renderers

import "image"

var diacritics = []string{
	"\u0305",
	"\u030D",
	"\u030E",
	"\u0310",
	"\u0312",
	"\u033D",
	"\u033E",
	"\u033F",
	"\u0346",
	"\u034A",
	"\u034B",
	"\u034C",
	"\u0350",
	"\u0351",
	"\u0352",
	"\u0357",
	"\u035B",
	"\u0363",
	"\u0364",
	"\u0365",
	"\u0366",
	"\u0367",
	"\u0368",
	"\u0369",
	"\u036A",
	"\u036B",
	"\u036C",
	"\u036D",
	"\u036E",
	"\u036F",
	"\u0483",
	"\u0484",
	"\u0485",
	"\u0486",
	"\u0487",
	"\u0592",
	"\u0593",
	"\u0594",
	"\u0595",
	"\u0597",
	"\u0598",
	"\u0599",
	"\u059C",
	"\u059D",
	"\u059E",
	"\u059F",
	"\u05A0",
	"\u05A1",
	"\u05A8",
	"\u05A9",
	"\u05AB",
	"\u05AC",
	"\u05AF",
	"\u05C4",
	"\u0610",
	"\u0611",
	"\u0612",
	"\u0613",
	"\u0614",
	"\u0615",
	"\u0616",
	"\u0617",
	"\u0657",
	"\u0658",
	"\u0659",
	"\u065A",
	"\u065B",
	"\u065D",
	"\u065E",
	"\u06D6",
	"\u06D7",
	"\u06D8",
	"\u06D9",
	"\u06DA",
	"\u06DB",
	"\u06DC",
	"\u06DF",
	"\u06E0",
	"\u06E1",
	"\u06E2",
	"\u06E4",
	"\u06E7",
	"\u06E8",
	"\u06EB",
	"\u06EC",
	"\u0730",
	"\u0732",
	"\u0733",
	"\u0735",
	"\u0736",
	"\u073A",
	"\u073D",
	"\u073F",
	"\u0740",
	"\u0741",
	"\u0743",
	"\u0745",
	"\u0747",
	"\u0749",
	"\u074A",
	"\u07EB",
	"\u07EC",
	"\u07ED",
	"\u07EE",
	"\u07EF",
	"\u07F0",
	"\u07F1",
	"\u07F3",
	"\u0816",
	"\u0817",
	"\u0818",
	"\u0819",
	"\u081B",
	"\u081C",
	"\u081D",
	"\u081E",
	"\u081F",
	"\u0820",
	"\u0821",
	"\u0822",
	"\u0823",
	"\u0825",
	"\u0826",
	"\u0827",
	"\u0829",
	"\u082A",
	"\u082B",
	"\u082C",
	"\u082D",
	"\u0951",
	"\u0953",
	"\u0954",
	"\u0F82",
	"\u0F83",
	"\u0F86",
	"\u0F87",
	"\u135D",
	"\u135E",
	"\u135F",
	"\u17DD",
	"\u193A",
	"\u1A17",
	"\u1A75",
	"\u1A76",
	"\u1A77",
	"\u1A78",
	"\u1A79",
	"\u1A7A",
	"\u1A7B",
	"\u1A7C",
	"\u1B6B",
	"\u1B6D",
	"\u1B6E",
	"\u1B6F",
	"\u1B70",
	"\u1B71",
	"\u1B72",
	"\u1B73",
	"\u1CD0",
	"\u1CD1",
	"\u1CD2",
	"\u1CDA",
	"\u1CDB",
	"\u1CE0",
	"\u1DC0",
	"\u1DC1",
	"\u1DC3",
	"\u1DC4",
	"\u1DC5",
	"\u1DC6",
	"\u1DC7",
	"\u1DC8",
	"\u1DC9",
	"\u1DCB",
	"\u1DCC",
	"\u1DD1",
	"\u1DD2",
	"\u1DD3",
	"\u1DD4",
	"\u1DD5",
	"\u1DD6",
	"\u1DD7",
	"\u1DD8",
	"\u1DD9",
	"\u1DDA",
	"\u1DDB",
	"\u1DDC",
	"\u1DDD",
	"\u1DDE",
	"\u1DDF",
	"\u1DE0",
	"\u1DE1",
	"\u1DE2",
	"\u1DE3",
	"\u1DE4",
	"\u1DE5",
	"\u1DE6",
	"\u1DFE",
	"\u20D0",
	"\u20D1",
	"\u20D4",
	"\u20D5",
	"\u20D6",
	"\u20D7",
	"\u20DB",
	"\u20DC",
	"\u20E1",
	"\u20E7",
	"\u20E9",
	"\u20F0",
	"\u2CEF",
	"\u2CF0",
	"\u2CF1",
	"\u2DE0",
	"\u2DE1",
	"\u2DE2",
	"\u2DE3",
	"\u2DE4",
	"\u2DE5",
	"\u2DE6",
	"\u2DE7",
	"\u2DE8",
	"\u2DE9",
	"\u2DEA",
	"\u2DEB",
	"\u2DEC",
	"\u2DED",
	"\u2DEE",
	"\u2DEF",
	"\u2DF0",
	"\u2DF1",
	"\u2DF2",
	"\u2DF3",
	"\u2DF4",
	"\u2DF5",
	"\u2DF6",
	"\u2DF7",
	"\u2DF8",
	"\u2DF9",
	"\u2DFA",
	"\u2DFB",
	"\u2DFC",
	"\u2DFD",
	"\u2DFE",
	"\u2DFF",
	"\uA66F",
	"\uA67C",
	"\uA67D",
	"\uA6F0",
	"\uA6F1",
	"\uA8E0",
	"\uA8E1",
	"\uA8E2",
	"\uA8E3",
	"\uA8E4",
	"\uA8E5",
	"\uA8E6",
	"\uA8E7",
	"\uA8E8",
	"\uA8E9",
	"\uA8EA",
	"\uA8EB",
	"\uA8EC",
	"\uA8ED",
	"\uA8EE",
	"\uA8EF",
	"\uA8F0",
	"\uA8F1",
	"\uAAB0",
	"\uAAB2",
	"\uAAB3",
	"\uAAB7",
	"\uAAB8",
	"\uAABE",
	"\uAABF",
	"\uAAC1",
	"\uFE20",
	"\uFE21",
	"\uFE22",
	"\uFE23",
	"\uFE24",
	"\uFE25",
	"\uFE26",
	"\u10A0F",
	"\u10A38",
	"\u1D185",
	"\u1D186",
	"\u1D187",
	"\u1D188",
	"\u1D189",
	"\u1D1AA",
	"\u1D1AB",
	"\u1D1AC",
	"\u1D1AD",
	"\u1D242",
	"\u1D243",
	"\u1D244",
}

type Kgp struct{}

func (k *Kgp) Name() string {
	return "kgp"
}

func (k *Kgp) ImageShow(image *Image) (image.Rectangle, error) {
}

func (k *Kgp) ImageErase(image *Image) error {
}
