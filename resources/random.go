package resources

import (
    "crypto/rand"
)

var (
    megabyte []byte
)

func init() {

    megabyte = make([]byte, 1024 * 1024)
    _, err := rand.Read(megabyte)
    if nil != err {
        panic(401)
    }

}

func Data() []byte {
    return megabyte
}