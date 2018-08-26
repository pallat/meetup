package cast

// Func Type
//
// ยกตัวอย่างเช่น
//
type Signature interface {
	//
	// type ที่มี signature
	//
	// โดยปกติถ้าจะ implement ก็มักจะสร้าง struct เช่น
	//	type something struct {}
	//
	//	func (s something) Do(a,b int) (int, int) {
	//		return b, a
	//	}
	//
	Do(int, int) (int, int)
}
