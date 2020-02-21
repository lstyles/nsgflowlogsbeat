// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package compact

// parents maps a compact index of a tag to the compact index of the parent of
// this tag.
var parents = []ID{ // 775 elements
	// Entry 0 - 3F
	0x0000, 0x0000, 0x0001, 0x0001, 0x0000, 0x0004, 0x0000, 0x0006,
	0x0000, 0x0008, 0x0000, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a,
	0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a,
	0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a,
	0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x000a, 0x0000,
	0x0000, 0x0028, 0x0000, 0x002a, 0x0000, 0x002c, 0x0000, 0x0000,
	0x002f, 0x002e, 0x002e, 0x0000, 0x0033, 0x0000, 0x0035, 0x0000,
	0x0037, 0x0000, 0x0039, 0x0000, 0x003b, 0x0000, 0x0000, 0x003e,
	// Entry 40 - 7F
	0x0000, 0x0040, 0x0040, 0x0000, 0x0043, 0x0043, 0x0000, 0x0046,
	0x0000, 0x0048, 0x0000, 0x0000, 0x004b, 0x004a, 0x004a, 0x0000,
	0x004f, 0x004f, 0x004f, 0x004f, 0x0000, 0x0054, 0x0054, 0x0000,
	0x0057, 0x0000, 0x0059, 0x0000, 0x005b, 0x0000, 0x005d, 0x005d,
	0x0000, 0x0060, 0x0000, 0x0062, 0x0000, 0x0064, 0x0000, 0x0066,
	0x0066, 0x0000, 0x0069, 0x0000, 0x006b, 0x006b, 0x006b, 0x006b,
	0x006b, 0x006b, 0x006b, 0x0000, 0x0073, 0x0000, 0x0075, 0x0000,
	0x0077, 0x0000, 0x0000, 0x007a, 0x0000, 0x007c, 0x0000, 0x007e,
	// Entry 80 - BF
	0x0000, 0x0080, 0x0080, 0x0000, 0x0083, 0x0083, 0x0000, 0x0086,
	0x0087, 0x0087, 0x0087, 0x0086, 0x0088, 0x0087, 0x0087, 0x0087,
	0x0086, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0088,
	0x0087, 0x0087, 0x0087, 0x0087, 0x0088, 0x0087, 0x0088, 0x0087,
	0x0087, 0x0088, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0087, 0x0087, 0x0087, 0x0086, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0086, 0x0087, 0x0086,
	// Entry C0 - FF
	0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0088, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0086, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0088, 0x0087,
	0x0087, 0x0088, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087,
	0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0086, 0x0086, 0x0087,
	0x0087, 0x0086, 0x0087, 0x0087, 0x0087, 0x0087, 0x0087, 0x0000,
	0x00ef, 0x0000, 0x00f1, 0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x00f2,
	0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x00f1, 0x00f2, 0x00f1, 0x00f1,
	// Entry 100 - 13F
	0x00f2, 0x00f2, 0x00f1, 0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x00f1,
	0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x00f2, 0x0000, 0x010e,
	0x0000, 0x0110, 0x0000, 0x0112, 0x0000, 0x0114, 0x0114, 0x0000,
	0x0117, 0x0117, 0x0117, 0x0117, 0x0000, 0x011c, 0x0000, 0x011e,
	0x0000, 0x0120, 0x0120, 0x0000, 0x0123, 0x0123, 0x0123, 0x0123,
	0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123,
	0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123,
	0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123,
	// Entry 140 - 17F
	0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123,
	0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123, 0x0123,
	0x0123, 0x0123, 0x0000, 0x0152, 0x0000, 0x0154, 0x0000, 0x0156,
	0x0000, 0x0158, 0x0000, 0x015a, 0x0000, 0x015c, 0x015c, 0x015c,
	0x0000, 0x0160, 0x0000, 0x0000, 0x0163, 0x0000, 0x0165, 0x0000,
	0x0167, 0x0167, 0x0167, 0x0000, 0x016b, 0x0000, 0x016d, 0x0000,
	0x016f, 0x0000, 0x0171, 0x0171, 0x0000, 0x0174, 0x0000, 0x0176,
	0x0000, 0x0178, 0x0000, 0x017a, 0x0000, 0x017c, 0x0000, 0x017e,
	// Entry 180 - 1BF
	0x0000, 0x0000, 0x0000, 0x0182, 0x0000, 0x0184, 0x0184, 0x0184,
	0x0184, 0x0000, 0x0000, 0x0000, 0x018b, 0x0000, 0x0000, 0x018e,
	0x0000, 0x0000, 0x0191, 0x0000, 0x0000, 0x0000, 0x0195, 0x0000,
	0x0197, 0x0000, 0x0000, 0x019a, 0x0000, 0x0000, 0x019d, 0x0000,
	0x019f, 0x0000, 0x01a1, 0x0000, 0x01a3, 0x0000, 0x01a5, 0x0000,
	0x01a7, 0x0000, 0x01a9, 0x0000, 0x01ab, 0x0000, 0x01ad, 0x0000,
	0x01af, 0x0000, 0x01b1, 0x01b1, 0x0000, 0x01b4, 0x0000, 0x01b6,
	0x0000, 0x01b8, 0x0000, 0x01ba, 0x0000, 0x01bc, 0x0000, 0x0000,
	// Entry 1C0 - 1FF
	0x01bf, 0x0000, 0x01c1, 0x0000, 0x01c3, 0x0000, 0x01c5, 0x0000,
	0x01c7, 0x0000, 0x01c9, 0x0000, 0x01cb, 0x01cb, 0x01cb, 0x01cb,
	0x0000, 0x01d0, 0x0000, 0x01d2, 0x01d2, 0x0000, 0x01d5, 0x0000,
	0x01d7, 0x0000, 0x01d9, 0x0000, 0x01db, 0x0000, 0x01dd, 0x0000,
	0x01df, 0x01df, 0x0000, 0x01e2, 0x0000, 0x01e4, 0x0000, 0x01e6,
	0x0000, 0x01e8, 0x0000, 0x01ea, 0x0000, 0x01ec, 0x0000, 0x01ee,
	0x0000, 0x01f0, 0x0000, 0x0000, 0x01f3, 0x0000, 0x01f5, 0x01f5,
	0x01f5, 0x0000, 0x01f9, 0x0000, 0x01fb, 0x0000, 0x01fd, 0x0000,
	// Entry 200 - 23F
	0x01ff, 0x0000, 0x0000, 0x0202, 0x0000, 0x0204, 0x0204, 0x0000,
	0x0207, 0x0000, 0x0209, 0x0209, 0x0000, 0x020c, 0x020c, 0x0000,
	0x020f, 0x020f, 0x020f, 0x020f, 0x020f, 0x020f, 0x020f, 0x0000,
	0x0217, 0x0000, 0x0219, 0x0000, 0x021b, 0x0000, 0x0000, 0x0000,
	0x0000, 0x0000, 0x0221, 0x0000, 0x0000, 0x0224, 0x0000, 0x0226,
	0x0226, 0x0000, 0x0229, 0x0000, 0x022b, 0x022b, 0x0000, 0x0000,
	0x022f, 0x022e, 0x022e, 0x0000, 0x0000, 0x0234, 0x0000, 0x0236,
	0x0000, 0x0238, 0x0000, 0x0244, 0x023a, 0x0244, 0x0244, 0x0244,
	// Entry 240 - 27F
	0x0244, 0x0244, 0x0244, 0x0244, 0x023a, 0x0244, 0x0244, 0x0000,
	0x0247, 0x0247, 0x0247, 0x0000, 0x024b, 0x0000, 0x024d, 0x0000,
	0x024f, 0x024f, 0x0000, 0x0252, 0x0000, 0x0254, 0x0254, 0x0254,
	0x0254, 0x0254, 0x0254, 0x0000, 0x025b, 0x0000, 0x025d, 0x0000,
	0x025f, 0x0000, 0x0261, 0x0000, 0x0263, 0x0000, 0x0265, 0x0000,
	0x0000, 0x0268, 0x0268, 0x0268, 0x0000, 0x026c, 0x0000, 0x026e,
	0x0000, 0x0270, 0x0000, 0x0000, 0x0000, 0x0274, 0x0273, 0x0273,
	0x0000, 0x0278, 0x0000, 0x027a, 0x0000, 0x027c, 0x0000, 0x0000,
	// Entry 280 - 2BF
	0x0000, 0x0000, 0x0281, 0x0000, 0x0000, 0x0284, 0x0000, 0x0286,
	0x0286, 0x0286, 0x0286, 0x0000, 0x028b, 0x028b, 0x028b, 0x0000,
	0x028f, 0x028f, 0x028f, 0x028f, 0x028f, 0x0000, 0x0295, 0x0295,
	0x0295, 0x0295, 0x0000, 0x0000, 0x0000, 0x0000, 0x029d, 0x029d,
	0x029d, 0x0000, 0x02a1, 0x02a1, 0x02a1, 0x02a1, 0x0000, 0x0000,
	0x02a7, 0x02a7, 0x02a7, 0x02a7, 0x0000, 0x02ac, 0x0000, 0x02ae,
	0x02ae, 0x0000, 0x02b1, 0x0000, 0x02b3, 0x0000, 0x02b5, 0x02b5,
	0x0000, 0x0000, 0x02b9, 0x0000, 0x0000, 0x0000, 0x02bd, 0x0000,
	// Entry 2C0 - 2FF
	0x02bf, 0x02bf, 0x0000, 0x0000, 0x02c3, 0x0000, 0x02c5, 0x0000,
	0x02c7, 0x0000, 0x02c9, 0x0000, 0x02cb, 0x0000, 0x02cd, 0x02cd,
	0x0000, 0x0000, 0x02d1, 0x0000, 0x02d3, 0x02d0, 0x02d0, 0x0000,
	0x0000, 0x02d8, 0x02d7, 0x02d7, 0x0000, 0x0000, 0x02dd, 0x0000,
	0x02df, 0x0000, 0x02e1, 0x0000, 0x0000, 0x02e4, 0x0000, 0x02e6,
	0x0000, 0x0000, 0x02e9, 0x0000, 0x02eb, 0x0000, 0x02ed, 0x0000,
	0x02ef, 0x02ef, 0x0000, 0x0000, 0x02f3, 0x02f2, 0x02f2, 0x0000,
	0x02f7, 0x0000, 0x02f9, 0x02f9, 0x02f9, 0x02f9, 0x02f9, 0x0000,
	// Entry 300 - 33F
	0x02ff, 0x0300, 0x02ff, 0x0000, 0x0303, 0x0051, 0x00e6,
} // Size: 1574 bytes

// Total table size 1574 bytes (1KiB); checksum: 895AAF0B
