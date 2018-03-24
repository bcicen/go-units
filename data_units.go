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
	Data = NewQuantity("bytes")

	Byte      = Data.NewUnit("byte", "B")
	KiloByte  = Data.NewUnit("kilobyte", "KB")
	MegaByte  = Data.NewUnit("megabyte", "MB")
	GigaByte  = Data.NewUnit("gigabyte", "GB")
	TeraByte  = Data.NewUnit("terabyte", "TB")
	PetaByte  = Data.NewUnit("petabyte", "PB")
	ExaByte   = Data.NewUnit("exabyte", "")
	ZettaByte = Data.NewUnit("zettabyte", "")
	YottaByte = Data.NewUnit("yottabyte", "")

	Bit     = Data.NewUnit("bit", "b")
	ExaBit  = Exa(Bit)
	PetaBit = Peta(Bit)
	TeraBit = Tera(Bit)
	GigaBit = Giga(Bit)
	MegaBit = Mega(Bit)
	KiloBit = Kilo(Bit)

	Nibble = Data.NewUnit("nibble", "")
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
