package lib

import "Phoenix-Chain-Core/libs/common"

type Function struct {
	Type        FunctionType
	InputParams []interface{}
}

func NewFunction(typ FunctionType, params []interface{}) *Function {
	return &Function{typ, params}
}

func (f *Function) encodeType() BytesSlice {
	typ := f.Type.GetType()
	encoded := MinimalByteArray(typ)
	//fmt.Println("ftype0: " + hexutil.Encode(encoded))
	b := EncodeBytes(encoded, OFFSET_SHORT_STRING)
	//fmt.Println("ftype1: " + hexutil.Encode(b))
	return b
}

func (f *Function) encodeParams() []BytesSlice {
	var encodeParamItem func(p interface{}) []byte
	encodeParamItem = func(p interface{}) []byte {
		switch p.(type) {
		case []byte:
			return EncodeBytes(p.([]byte), OFFSET_SHORT_STRING)
		case BytesSlice:
			return EncodeBytes(p.(BytesSlice), OFFSET_SHORT_STRING)
		case common.Address:
			return EncodeBytes(p.(common.Address).Bytes(), OFFSET_SHORT_STRING)
		case NodeId:
			return p.(NodeId).GetEncodeData()
		case HexStringParam:
			return p.(HexStringParam).GetEncodeData()
		case Utf8String:
			return p.(Utf8String).GetEncodeData()
		case UInt16, UInt32, UInt64, UInt256:
			return p.(ParamEncoder).GetEncodeData()
		case RlpList:
			encoder := RlpEncoder{}
			return encoder.Encode(p.(RlpList))
		case []NodeId:
			var r []byte
			for _, it := range p.([]NodeId) {
				r = append(r, encodeParamItem(it)...)
			}
			r = EncodeBytes(r, OFFSET_SHORT_LIST)
			return r
		case []BytesSlice:
			var r []byte
			for _, it := range p.([]BytesSlice) {
				r = append(r, encodeParamItem(it)...)
			}
			r = EncodeBytes(r, OFFSET_SHORT_LIST)
			return r
		default:
			panic("function parameters with not support type")
		}
	}

	params := f.InputParams
	var result []BytesSlice
	for _, p := range params {
		result = append(result, encodeParamItem(p))
	}
	return result
}

func (f *Function) ToBytes() []byte {
	ftype := f.encodeType()
	params := f.encodeParams()

	argsList := append([]BytesSlice{ftype}, params...)
	//fmt.Println("Function Parameters List:")
	//for _, p := range argsList {
	//	fmt.Println(hexutil.Encode(p)[2:])
	//}

	argsBytes := EncodeBytesSlice(argsList)

	//data := hexutil.Encode(argsBytes)
	//fmt.Println("Function Data is: " + data[2:])

	return argsBytes
}
