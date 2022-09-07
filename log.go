// Copyright (C) 2023 CHUNQIAN SHEN. All rights reserved.

package tinymemory

func PX_ASSERT() {
}

func PX_ERROR(err error) {
	if err != nil {
		panic(err)
	}
}
