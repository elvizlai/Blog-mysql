package common

import (
    "testing"
)

var orignList = [...]string{
    "p@ssw0rd",
    "www.huagai.com",
    "你好，世界！",
    "huagai\n%s",
    "!23456789",
}

func TestBase64Decode(t *testing.T) {
    for i := 0; i < len(orignList); i++ {
        encode := Base64Encode([]byte(orignList[i]))
        decode, err := Base64Decode(encode)
        if err != nil {
            t.Error("Test case Base64Decode Failed!", err)
        }
        if string(decode) != orignList[i] {
            t.Error("Test case Base64Decode Failed!")
        }
    }
}

func TestRsaDecrypt(t *testing.T) {
    for i := 0; i < len(orignList); i++ {
        encrypt, err := RsaEncrypt([]byte(orignList[i]))
        if err != nil {
            t.Error("Test case RsaDecrypt Failed!", err)
        }

        decrypt, err := RsaDecrypt(encrypt)
        if err != nil {
            t.Error("Test case RsaDecrypt Failed!", err)
        }

        if string(decrypt) != orignList[i] {
            t.Error("Test case RsaDecrypt Failed!")
        }
    }
}
