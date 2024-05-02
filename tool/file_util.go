package tool

import "os"

// 帮助我实现根据文件路径读取文件内容，返回[]byte
func ReadFile(filePath string) ([]byte, error) {
	//帮助我实现根据文件路径读取文件内容，返回[]byte
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	//获取文件信息
	fInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}
	//获取文件大小
	size := fInfo.Size()
	//创建一个切片
	buf := make([]byte, size)
	//读取文件内容
	_, err = f.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
