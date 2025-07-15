// 判断文件是否为空的函数
package PathExist

import (
	"os"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else if os.IsExist(err) {
		return true, nil
	} else {
		return false, err
	}

}
