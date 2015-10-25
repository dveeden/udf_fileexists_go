package main

// #include <stdio.h>
// #include <sys/types.h>
// #include <sys/stat.h>
// #include <stdlib.h>
// #include <string.h>
// #include <mysql.h>
// #cgo CFLAGS: -I/opt/mysql/5.7.9-2/include -fabi-version=2 -fno-omit-frame-pointer
import "C"
import "os"

//export udf_fileexists_go_init
func udf_fileexists_go_init(initid *C.UDF_INIT, args *C.UDF_ARGS, message *C.char) C.my_bool {
	if args.arg_count != 1 {
		msg := "udf_fileexists_go() requires one string argument\n"
		C.strcpy(message, C.CString(msg))
		return 1;
	}

	return 0;
}

//export udf_fileexists_go
func udf_fileexists_go(initid *C.UDF_INIT, args *C.UDF_ARGS, result *C.char,
                    res_length uint64, null_value *C.char, error *C.char) int64 {
	filename := C.GoString(*args.args)
	_, err := os.Stat(filename)
	if err != nil {
		return 0
	}

	return 1; // file exists
}

func main() {}
