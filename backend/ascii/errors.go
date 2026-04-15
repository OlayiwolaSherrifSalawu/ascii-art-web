package ascii

type ConstantError string

func (e ConstantError) Error() string {
	return string(e)
}

const (
	INVALID_CHAR_VAl = ConstantError("Non Printable character detected")
	EMPTY_FILE       = ConstantError("font file is empty")
	NOCOLORWORD      = ConstantError("cannot print empty color word ")
	INVALID_INPUTS   = ConstantError("Invalid input arguements")
	CANT_READ_BANNER = ConstantError("cant read banner file , file corrupted ")
)
