// Package cast to present type conversions
// 	https://tour.golang.org มีภาษาไทยแล้วนะ รู้ยัง!!
//
// Cast Type 3 แบบ
//
// 1. primitive type
//
// 2. struct type
//
// 3. signature type
//
//
// Primitive Type
//
// primitie type
//	[]byte <-> string
//	string <-> []rune
//	string <-> int
//	int <-> float64
//
// Struct Type
//
// มีเงื่อนไขดังนี้
//	1. มีจำนวน fields เท่ากัน
//	2. ทุก fields มีชื่อตรงกัน
//	3. ทุก fields มี type ตรงกัน
//
// Signature Type
//
// ยกตัวอย่างเช่น
//	type Stringer interface {
//		String() string
//	}
//
// โดยปกติถ้าจะ implement ก็มักจะสร้าง struct เช่น
//	type something struct {
//		asset string
//	}
//
//	func (s something) String() string {
//		return s.asset
//	}
//
//
package cast
