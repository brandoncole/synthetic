package resources

import (
    "crypto/sha512"
)

func ProcessorSimulation() {

    hasher := sha512.New()
    hasher.Write(Data())
    hasher.Sum(nil)

}