package utils

import (
	"fmt"
	"path/filepath"
	"time"
)

// GenFilename 用时间戳生成文件名
func GenFilename(filename string) string {
	ext := filepath.Ext(filename)
	return fmt.Sprintf("%d%s", time.Now().UnixMilli(), ext)
}
