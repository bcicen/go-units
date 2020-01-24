package units

const (
	// byte ratio constants
	_          = iota
	kb float64 = 1 << (10 * iota)
	mb
	gb
	tb
	pb
	eb
	zb
	yb
)

var (
	Bi   = UnitOptionQuantity("bits")
	Data = UnitOptionQuantity("bytes")

	Byte      = NewUnit("byte", "B", Data)
	KiloByte  = NewUnit("kilobyte", "KB", Data)
	MegaByte  = NewUnit("megabyte", "MB", Data)
	GigaByte  = NewUnit("gigabyte", "GB", Data)
	TeraByte  = NewUnit("terabyte", "TB", Data)
	PetaByte  = NewUnit("petabyte", "PB", Data)
	ExaByte   = NewUnit("exabyte", "", Data)
	ZettaByte = NewUnit("zettabyte", "", Data)
	YottaByte = NewUnit("yottabyte", "", Data)

	Bit     = NewUnit("bit", "b", Bi)
	ExaBit  = Exa(Bit)
	PetaBit = Peta(Bit)
	TeraBit = Tera(Bit)
	GigaBit = Giga(Bit)
	MegaBit = Mega(Bit)
	KiloBit = Kilo(Bit)

	Nibble = NewUnit("nibble", "", Data)
)

func init() {
	NewRatioConversion(Nibble, Bit, 4.0)
	NewRatioConversion(Byte, Bit, 8.0)
	NewRatioConversion(KiloByte, Byte, kb)
	NewRatioConversion(MegaByte, Byte, mb)
	NewRatioConversion(GigaByte, Byte, gb)
	NewRatioConversion(TeraByte, Byte, tb)
	NewRatioConversion(PetaByte, Byte, pb)
	NewRatioConversion(ExaByte, Byte, eb)
	NewRatioConversion(ZettaByte, Byte, zb)
	NewRatioConversion(YottaByte, Byte, yb)
}
