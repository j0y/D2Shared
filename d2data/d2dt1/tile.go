package d2dt1

type Tile struct {
	Direction          int32
	RoofHeight         int16
	MaterialFlags      int16
	Height             int32
	Width              int32
	Orientation        int32
	MainIndex          int32
	SubIndex           int32
	RarityFrameIndex   int32
	SubTileFlags       [25]SubTileFlags
	blockHeaderPointer int32
	blockHeaderSize    int32
	Blocks             []Block
}
