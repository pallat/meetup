package cast

// Struct type
// มีเงื่อนไขดังนี้
//	1. มีจำนวน fields เท่ากัน
//	2. ทุก fields มีชื่อตรงกัน
//	3. ทุก fields มี type ตรงกัน
// ยกตัวอย่างเช่น
type Struct struct {
	Foo string
	Bar string
	Baz string
}
