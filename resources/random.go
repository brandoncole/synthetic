package resources

import (
    "crypto/rand"
)

const (
    SIZE = 1024 * 512
)

var (
    data []byte
)

func init() {

    data = make([]byte, SIZE)
    _, err := rand.Read(data)
    if nil != err {
        panic(401)
    }

}

func Data() []byte {
    return data
}

func DataLength() int {
    return SIZE
}