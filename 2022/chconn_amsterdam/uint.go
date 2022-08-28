package main

// UINT START OMIT

type Uint128 struct {
	Lo uint64 // lower 64-bit half
	Hi uint64 // upper 64-bit half
}

type Uint256 struct {
	Lo Uint128 // lower 128-bit half
	Hi Uint128 // upper 128-bit half
}

// UINT END OMIT
