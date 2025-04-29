package errs

import "fmt"

// NewErrIndexOutOfBounds 创建索引越界错误
func NewErrIndexOutOfBounds(len int, index int) error {
	return fmt.Errorf("gokit: 索引越界, 长度: %d, 索引: %d", len, index)
}
