package Parser

import (
	"strings"

	"github.com/moritzkalwa/go-parse/BinaryReader"
	"github.com/moritzkalwa/go-parse/Types"
)

func ParseBytesFormat5(bArr *[]byte) []Types.Info {
	br := BinaryReader.Br{Data: *bArr, Pointer: 0}
	var arrList []Types.Info

	for BinaryReader.HasNext(&br) {
		count := BinaryReader.GetInt(&br)
		for i2:= 0; i2 < count; i2++ {
			parentId := BinaryReader.GetVariableLengthString(&br)
			parentName := BinaryReader.GetVariableLengthString(&br)
			omc := BinaryReader.GetInt(&br)
			coord := Types.New(float64(BinaryReader.GetInt(&br))/100.0, float64(BinaryReader.GetInt(&br))/100.0)
			netLength := BinaryReader.GetByte(&br)
			nets := make([]Types.Net, netLength)
			for netIndex := 1; netIndex <= netLength; netIndex++ {
				netName := strings.ToLower(BinaryReader.GetString(&br, 3))
				zoneLength := BinaryReader.GetByte(&br)
				zones := make([]int, zoneLength)
				for zoneIndex := 1; zoneIndex <= zoneLength; zoneIndex++ {
					zones = append(zones, BinaryReader.GetInt(&br))
				}
				nets = append(nets, Types.Net{Name: netName, Zones: zones})
			}
			length := int(BinaryReader.GetShort(&br))
			for index := 1; index <= length; index++ {
				id := BinaryReader.GetVariableLengthString(&br)
				name := BinaryReader.GetVariableLengthString(&br)
				level := BinaryReader.GetShort(&br)
				elevation := BinaryReader.GetShort(&br)
				isTransferStation := BinaryReader.GetByte(&br) == 1
				length2 := BinaryReader.GetByte(&br)
				for index2 := 1; index2 <= int(length2); index2++ {
					BinaryReader.GetInt(&br)
					BinaryReader.GetInt(&br)
				}
				arrList = append(arrList, Types.Info{
					Coord:             coord,
					Elevation:         elevation,
					Id:                id,
					IsTransferStation: isTransferStation,
					Level:             level,
					Name:              name,
					Nets:              nets,
					Omc:               omc,
					ParentId:          parentId,
					ParentName:        parentName})
			}
		}
	}
	return arrList
}
