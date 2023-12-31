package funcs

const (
	TIOCGWINSZ = 0x5413
)

// type winsize struct {
// 	row    uint16
// 	col    uint16
// 	xpixel uint16
// 	ypixel uint16
// }

// func bytesToUint16(b []byte) uint16 {
// 	return uint16(b[1])<<8 | uint16(b[0])
// }

// func getTerminalWidth() (int, error) {
//     ws := &winsize{}
//     wsBytes := make([]byte, 8)
//     _, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdout), uintptr(TIOCGWINSZ), uintptr(&wsBytes[0]))
//     if errno != 0 {
//         return 0, errno
//     }
//     ws.row = bytesToUint16(wsBytes[0:2])
//     ws.col = bytesToUint16(wsBytes[2:4])
//     return int(ws.col), nil
// }
