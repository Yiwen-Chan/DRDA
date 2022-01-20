package drda

import (
	"fmt"
)

func main() {
	conn, err := NewConnect("127.0.0.1:50000")
	if err != nil {
		panic(err)
	}
	conn.Write(&DRDA{
		DDM: &DDM{Magic: 0xd0, Format: 0x41, CorrelId: 1, CodePoint: EXCSAT},
		Parameters: []*Parameter{
			{CodePoint: EXTNAM, Payload: ToEBCDIC([]byte("QDB2/NT64"))},
			{CodePoint: MGRLVLLS, Payload: []byte{0x14, 0x03, 0x00, 0x0a, 0x24, 0x07, 0x00, 0x0b, 0x14, 0x74, 0x00, 0x05, 0x24, 0x0f,
				0x00, 0x0c, 0x14, 0x40, 0x00, 0x0a, 0x1c, 0x08, 0x04, 0xb8}},
			{CodePoint: SRVCLSNM, Payload: ToEBCDIC([]byte("QDB2/NT64"))},
			{CodePoint: SRVNAM, Payload: ToEBCDIC([]byte("DESKTOP-TS0HAVA"))},
			{CodePoint: SRVRLSLV, Payload: ToEBCDIC([]byte("SQL11055"))},
		},
	})
	conn.Write(&DRDA{
		DDM: &DDM{Magic: 0xd0, Format: 0x01, CorrelId: 2, CodePoint: ACCSEC},
		Parameters: []*Parameter{
			{CodePoint: SECMEC, Payload: []byte{0x00, 0x09}},
			{CodePoint: RDBNAM, Payload: []byte{0xd6, 0xd5, 0xc5, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40, 0x40}},
			{CodePoint: SECTKN, Payload: []byte{0xa2, 0x04, 0x1b, 0x29, 0x03, 0x46, 0xfc, 0x6c, 0xa3, 0x0e, 0xc9, 0x4b, 0xa2, 0xd3, 0x28, 0x5f, 0x5a, 0x99,
				0xa8, 0x1c, 0xd2, 0x89, 0x80, 0xbf, 0x80, 0x09, 0xc8, 0x92, 0x32, 0x79, 0xdd, 0x02}},
		},
	})
	drda, _ := conn.Read()
	fmt.Printf("%x\n", drda.DDM.CodePoint)
	drda, _ = conn.Read()
	fmt.Printf("%x\n", drda.DDM.CodePoint)
	conn.Close()
}
