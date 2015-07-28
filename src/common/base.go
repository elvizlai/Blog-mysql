package common

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/base64"
    "encoding/pem"
    "errors"
)

// 公钥和私钥可以从文件中读取
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDo0bTPdgn/qzY9XSQiPeG7dSVAr/Zr1i7gAoVcSc/qzFffqvEq
fBdHr4cupN5+eurdlYNw4WF63BM8q2TJgAxrRgPwWzdeJx1rpJuP3JWJgLuIJ47G
PIFVIdB4EzDi9bN8MKMhkMfdx2/6Afo/McOjB6dvd7RZ2kiVKd1zKReuewIDAQAB
AoGBALKF4vxVydL3KQ8itYtgIhBJAni4tN75jFYO+M3Md5bWe+cxP93Q61T3nlPA
7i7T9/ZTKEfNOp8n08RotE7iViIRGPvwe/rb4ACJRTsX+McKR2wrHsFgL7dL+wfU
I3HtrhKKILR7lbhdKWgV5B2vCTIHlO9U7JsFZC8KU1Sop595AkEA/PTivZV465Eo
+mVLM0h1a1z+BC5/7Vbzm/FOIgmXo+fPkP2t0rJq9VPVk19JkLzy7x5X+5Kae4DP
329kBjNubwJBAOuey/Q0PEXzdQPKPP5Vov3OYDr3wlAokKAr6hR4aMDYNceLAo69
/G59DGf7Vtn6632TE6zVOjzh4accwWIcBrUCQH4m8t1xqfhxUGpwEezlegmtOtGD
DzGiZ6Oh2EGJXyLS/OVmXkXxzP3EbYMtxlZ0pQMzstU36+sj9oeL2eptw+kCQQCN
jvQeHZvwstokksaeTzkDn4/1HZFis1xgvsF91vGomr2EyyGYPNCCWSKw/jIp+DSv
V0PE3L6GgXce/Ym5tfjZAkBmPYmp5uF6SpAuO501K4Z4vQhxRpFOd3VmtVn7Jgok
ClkgNcIZQDNzsFHDqwadtOvSbVibpYqihwPv2w6bc/O7
-----END RSA PRIVATE KEY-----
`)

var PublicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDo0bTPdgn/qzY9XSQiPeG7dSVA
r/Zr1i7gAoVcSc/qzFffqvEqfBdHr4cupN5+eurdlYNw4WF63BM8q2TJgAxrRgPw
WzdeJx1rpJuP3JWJgLuIJ47GPIFVIdB4EzDi9bN8MKMhkMfdx2/6Afo/McOjB6dv
d7RZ2kiVKd1zKReuewIDAQAB
-----END PUBLIC KEY-----
`)

func Base64Encode(bytes []byte) string {
    return base64.StdEncoding.EncodeToString(bytes)
}

func Base64Decode(str string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(str)
}

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
    block, _ := pem.Decode(PublicKey)
    if block == nil {
        return nil, errors.New("public key error")
    }
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    pub := pubInterface.(*rsa.PublicKey)
    rsa, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
    return []byte(Base64Encode(rsa)), err
}

// 解密
func RsaDecrypt(cipherbase64 []byte) ([]byte, error) {
    cipher, err := Base64Decode(string(cipherbase64))
    if err != nil {
        return nil, err
    }
    block, _ := pem.Decode(privateKey)
    if block == nil {
        return nil, errors.New("private key error!")
    }
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return nil, err
    }
    return rsa.DecryptPKCS1v15(rand.Reader, priv, cipher)
}
