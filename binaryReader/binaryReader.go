package BinaryReader

type Br struct {
	Data    []byte
	Pointer int
}

// Memory Functions
func GetByte(br *Br) int {
	(*br).Pointer += 1
	return int((*br).Data[(*br).Pointer - 1] & 255)
}

func GetBytes(br *Br) []byte {
	return (*br).Data
}


func GetByteIndex(br *Br, index int) int {
	return int((*br).Data[index] & 255)
}

func GetInt(br *Br) int {
	return GetByte(br) | (GetByte(br) << 8) | (GetByte(br) << 16) | (GetByte(br) << 24)
}

func GetShort(br *Br) int16 {
	return int16(((GetByte(br)) | GetByte(br)<<8) & 65535)
}

func GetPointer(br *Br) int {
	return (*br).Pointer
}

func SetPointer(br *Br, pointer int) {
	(*br).Pointer = pointer
}

func HasNext(br *Br) bool {
	return (*br).Pointer < len((*br).Data)
}

func GetString(br *Br, length int) string {
	slice := (*br).Data[(*br).Pointer : (*br).Pointer+length]
	var result []byte
	for i := 0; i < length; i++ {
		if slice[i] == 0 {
			break
		}
		result = append(result, slice[i])
	}
	(*br).Pointer += length
	return string(result)
}

func GetVariableLengthString(br *Br) string {
	length := GetByte(br)
	return GetString(br, length)
}
