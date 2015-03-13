// Crypto
// utils

package main

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "fmt"
)

var coder = base64.StdEncoding

func base64Encode(src []byte) []byte {
    return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
    return coder.DecodeString(string(src))
}

func aesEncode(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)

    return base64Encode(crypted), nil
}

func aesDecode(src, key []byte) ([]byte, error) {
    // decode
    crypted, err := base64Decode(src)
    if err != nil {
        fmt.Println(err.Error())
    }

    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    // origData := crypted
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    // origData = ZeroUnPadding(origData)
    return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func testAes() {
    key := []byte("sfe023f_9fd&fwfl")
    result, err := aesEncode([]byte("polaris@studygolang"), key)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(result))
    origData, err := aesDecode(result, key)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(origData))
}

func main() {
    testAes()
}
