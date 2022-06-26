/*
@Time       : 2022/3/30
@Author     : wuqiusheng
@File       : token.go
@Description:
*/
package easyutil

//token生成器
func GenToken(length int) string {
	token := make([]byte, length)
	for i := 0; i < length; i++ {
		v := RandNumBetween(1, 4)
		if v == 3 {
			b := RandNumBetween(48, 57)
			token[i] = byte(b)
		} else if v == 1 {
			b := RandNumBetween(65, 90)
			token[i] = byte(b)
		} else {
			b := RandNumBetween(97, 122)
			token[i] = byte(b)
		}
	}
	return string(token)
}

func PackLoginToken(msg interface{}) (string, error) {
	data, err := PBPack(msg)
	if err != nil {
		return "", err
	}
	token := Base64Encode(data)
	return string(token), nil
}

func UnPackLoginToken(token string, msg interface{}) error {
	data, err := Base64Decode([]byte(token))
	if err != nil {
		return err
	}
	return PBUnPack(data, msg)
}
