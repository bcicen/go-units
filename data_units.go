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

	Bit     = NewUnit("bit", "b", Data)
	ExaBit  = Exa(Bit)
	PetaBit = Peta(Bit)
	TeraBit = Tera(Bit)
	GigaBit = Giga(Bit)
	MegaBit = Mega(Bit)
	KiloBit = Kilo(Bit)

	Nibble = NewUnit("nibble", "", Data)
)

func init() {
	NewRatioConv(Nibble, Bit, 4.0)
	NewRatioConv(Byte, Bit, 8.0)
	NewRatioConv(KiloByte, Byte, kb)
	NewRatioConv(MegaByte, Byte, mb)
	NewRatioConv(GigaByte, Byte, gb)
	NewRatioConv(TeraByte, Byte, tb)
	NewRatioConv(PetaByte, Byte, pb)
	NewRatioConv(ExaByte, Byte, eb)
	NewRatioConv(ZettaByte, Byte, zb)
	NewRatioConv(YottaByte, Byte, yb)
}
