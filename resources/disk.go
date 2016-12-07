// Package resources contains implementations of simulations for finite resources that are simulated.
package resources

import (
	"io/ioutil"
	"os"
)

func DiskSimulation() {

	file, err := ioutil.TempFile(os.TempDir(), "")
	if nil != err {
		panic(err)
	}

	err = file.Close()
	if nil != err {
		panic(err)
	}

	file, err = os.OpenFile(file.Name(), os.O_WRONLY, 0600)
	if nil != err {
		panic(err)
	}

	n, err := file.Write(Data())
	if nil != err || n != DataLength() {
		panic(err)
	}

	err = file.Close()
	if nil != err {
		panic(err)
	}

    // The following code is essentially ineffective on macOS an *nix
    // because the file systems are smart enough to cache the file
    // contents in memory on the write so no I/O happens on the read.
    //
    // TODO(coleb) Find some way to force disk read operation
    /*
	file, err = os.OpenFile(file.Name(), os.O_RDONLY, 0600)
	if nil != err {
		panic(err)
	}

	buffer := make([]byte, DataLength())
	n, err = file.Read(buffer)
	if nil != err || n != DataLength() || 0 != bytes.Compare(buffer, Data()) {
		panic(err)
	}

	err = os.Remove(file.Name())
	if nil != err {
		panic(err)
	}
	*/

}
