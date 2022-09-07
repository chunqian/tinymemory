package memory

import unsafe "unsafe"

func _cgos_PX_isBigEndianCPU_PX_Typedef() int32 {
	type _cgoa_14_PX_Typedef struct {
		i uint32
	}
	var c _cgoa_14_PX_Typedef
	c.i = uint32(305419896)
	return func() int32 {
		if int32(18) == int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&*(*[4]uint8)(unsafe.Pointer(&c)))))) + uintptr(int32(0))))) {
			return 1
		} else {
			return 0
		}
	}()
}
func PX_htoi(hex_str *int8) uint32 {
	var ch int8
	var iret uint32 = uint32(0)
	for int32(func() (_cgo_ret int8) {
		_cgo_addr := &ch
		*_cgo_addr = *func() (_cgo_ret *int8) {
			_cgo_addr := &hex_str
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
		return *_cgo_addr
	}()) != int32(0) {
		iret = iret<<int32(4) | uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&_cgos_px_hex_to_dex_table_PX_Typedef)))) + uintptr(ch)*4)))
	}
	return iret
}
func PX_atoi(s *int8) int32 {
	var i int32
	var n int32
	var sign int32 = int32(1)
	for i = int32(0); int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == ' '; i++ {
	}
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '-' {
		sign = func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) == '-' {
				return -1
			} else {
				return int32(1)
			}
		}()
		i++
	}
	for n = int32(0); int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) >= '0' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) <= '9'; i++ {
		n = int32(10)*n + (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + uintptr(i)))) - '0')
	}
	return sign * n
}
func PX_atof(fstr *int8) float32 {
	var temp float64 = float64(int32(10))
	var ispnum int32 = int32(1)
	var ans float64 = float64(int32(0))
	var be int32 = int32(0)
	if int32(*fstr) == '-' {
		ispnum = int32(0)
		*(*uintptr)(unsafe.Pointer(&fstr))++
	} else if int32(*fstr) == '+' {
		*(*uintptr)(unsafe.Pointer(&fstr))++
	}
	for int32(*fstr) != '\x00' {
		if int32(*fstr) == '.' {
			*(*uintptr)(unsafe.Pointer(&fstr))++
			break
		}
		if int32(*fstr) == 'e' || int32(*fstr) == 'E' {
			*(*uintptr)(unsafe.Pointer(&fstr))++
			be = int32(1)
			break
		}
		ans = ans*float64(int32(10)) + float64(int32(*fstr)-'0')
		*(*uintptr)(unsafe.Pointer(&fstr))++
	}
	if be != 0 {
		var e int32 = PX_atoi(fstr)
		var e10 float32 = float32(int32(1))
		if e > int32(0) {
			for e != 0 {
				e10 *= float32(int32(10))
				e--
			}
		} else {
			for e != 0 {
				e10 /= float32(int32(10))
				e++
			}
		}
		ans *= float64(e10)
	} else {
		for int32(*fstr) != '\x00' {
			ans = ans + float64(int32(*fstr)-'0')/temp
			temp *= float64(int32(10))
			*(*uintptr)(unsafe.Pointer(&fstr))++
		}
	}
	if ispnum != 0 {
		return float32(ans)
	} else {
		return float32(ans) * float32(-1)
	}
	return 0
}
func PX_ftos(f float32, precision int32) _cgoa_1_PX_Typedef {
	var str _cgoa_1_PX_Typedef
	PX_ftoa(f, (*int8)(unsafe.Pointer(&str.data)), int32(64), precision)
	return str
}
func PX_itos(num int32, radix int32) _cgoa_1_PX_Typedef {
	var str _cgoa_1_PX_Typedef
	PX_itoa(num, (*int8)(unsafe.Pointer(&str.data)), int32(64), radix)
	return str
}
func PX_AscToWord(asc *int8, u16 *uint16) {
	for *asc != 0 {
		*u16 = uint16(*asc)
		*(*uintptr)(unsafe.Pointer(&u16)) += 2
		*(*uintptr)(unsafe.Pointer(&asc))++
	}
	*u16 = uint16(0)
}
func PX_ftoa(f float32, outbuf *int8, maxlen int32, precision int32) int32 {
	var i_value int32
	var f_value int32
	var shl int32 = int32(10)
	var len int32
	var zero_oft int32 = int32(0)
	if maxlen == int32(0) {
		return int32(0)
	}
	shl = PX_pow_ii(shl, precision)
	i_value = int32(f)
	f_value = int32(func() float32 {
		if float32(shl)*(f-float32(int32(f))) > float32(int32(0)) {
			return float32(shl) * (f - float32(int32(f)))
		} else {
			return -(float32(shl) * (f - float32(int32(f))))
		}
	}())
	if f_value < int32(0) {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(0)))) = int8('N')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(1)))) = int8('a')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(2)))) = int8('N')
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(3)))) = int8('\x00')
		return int32(0)
	}
	if i_value == int32(0) && f < float32(int32(0)) {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(int32(0)))) = int8('-')
		PX_itoa(i_value, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(int32(1)))), maxlen-int32(1), int32(10))
	} else {
		PX_itoa(i_value, outbuf, maxlen, int32(10))
	}
	len = PX_strlen(outbuf)
	if precision == int32(0) {
		return len
	}
	if len > maxlen-int32(3) {
		return len
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len))) = int8('.')
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len+int32(1)))) = int8('\x00')
	for f_value < shl {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(int32(1)))))) + uintptr(zero_oft))) = int8('0')
		shl /= int32(10)
		zero_oft++
	}
	PX_itoa(f_value, (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(zero_oft))), maxlen-len-int32(1), int32(10))
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf))+uintptr(len)))))+uintptr(int32(1)))))) + uintptr(precision))) = int8('\x00')
	len += PX_strlen((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(outbuf)) + uintptr(len))))
	return len
}
func PX_itoa(num int32, str *int8, MaxStrSize int32, radix int32) int32 {
	var index [37]int8 = [37]int8{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '\x00'}
	var unum uint32
	var temp int8
	var i int32 = int32(0)
	var j int32
	var l int32
	if MaxStrSize == int32(0) {
		return int32(0)
	}
	if radix == int32(10) && num < int32(0) {
		unum = uint32(-num)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()))) = int8('-')
	} else {
		unum = uint32(num)
	}
	for {
		if MaxStrSize <= i+int32(1) {
			return int32(0)
		}
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &i
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}()))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&index)))) + uintptr(unum%uint32(radix))))
		unum /= uint32(radix)
		if !(unum != 0) {
			break
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = int8('\x00')
	l = i
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-' {
		for func() int32 {
			j = int32(1)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				*_cgo_addr = i - int32(1)
				return *_cgo_addr
			}()
		}(); j < i; func() int32 {
			j++
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
			temp = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = temp
		}
	} else {
		for func() int32 {
			j = int32(0)
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}(); j < i; func() int32 {
			j++
			return func() (_cgo_ret int32) {
				_cgo_addr := &i
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}()
		}() {
			temp = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(j))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i)))
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(i))) = temp
		}
	}
	return l
}
func _cgos_PX_i32SwapEndian_PX_Typedef(val int32) int32 {
	val = int32(uint32(val<<int32(8))&uint32(4278255360) | uint32(val>>int32(8)&int32(16711935)))
	return val<<int32(16) | val>>int32(16)
}
func PX_SwapEndian(val uint32) uint32 {
	val = val<<int32(8)&uint32(4278255360) | val>>int32(8)&uint32(16711935)
	return val<<int32(16) | val>>int32(16)
}
func _cgos_PX_i64SwapEndian_PX_Typedef(val int64) int64 {
	var u32_host_h int32
	var u32_host_l int32
	var u64_net int64
	u32_host_l = int32(val & int64(4294967295))
	u32_host_h = int32(val >> int32(32) & int64(4294967295))
	u64_net = int64(_cgos_PX_i32SwapEndian_PX_Typedef(u32_host_l))
	u64_net = u64_net<<int32(32) | int64(_cgos_PX_i32SwapEndian_PX_Typedef(u32_host_h))
	return u64_net
}
func PX_sqrt(number float32) float32 {
	var i int32
	var x2 float32
	var y float32
	var threehalfs float32 = 1.5
	x2 = number * 0.5
	y = number
	i = *(*int32)(unsafe.Pointer(&y))
	if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
		i = _cgos_PX_i32SwapEndian_PX_Typedef(i)
	}
	i = int32(1597463174) - i>>int32(1)
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	return number * y
}
func PX_sqrtd(number float64) float64 {
	var i int64
	var x2 float64
	var y float64
	var threehalfs float64 = 1.5
	x2 = number * 0.5
	y = number
	i = *(*int64)(unsafe.Pointer(&y))
	if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
		i = _cgos_PX_i64SwapEndian_PX_Typedef(i)
	}
	i = int64(6910470738111508698) - i/int64(2)
	y = *(*float64)(unsafe.Pointer(&i))
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	return number * y
}
func PX_SqrtRec(number float32) float32 {
	var i int32
	var x2 float32
	var y float32
	var threehalfs float32 = 1.5
	x2 = number * 0.5
	y = number
	i = *(*int32)(unsafe.Pointer(&y))
	if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
		i = _cgos_PX_i32SwapEndian_PX_Typedef(i)
	}
	i = int32(1597463174) - i>>int32(1)
	y = *(*float32)(unsafe.Pointer(&i))
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	y = y * (threehalfs - x2*y*y)
	return y
}
func PX_fast_exp(x float64) float64 {
	x = 1 + x/float64(int32(65536))
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	x *= x
	return x
}
func PX_copysign(x float64, y float64) float64 {
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4)) = int32(uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))&int32(2147483647)) | uint32(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&y)))) + uintptr(int32(1))*4)))&uint32(2147483648))
	return x
}
func PX_scalbn(x float64, n float64) float64 {
	var k int32
	var hx int32
	var lx int32
	var two54 float64 = 18014398509481984
	var twom54 float64 = 5.5511151231257827e-17
	var huge float64 = 1.0000000000000001e+300
	var tiny float64 = 1.0e-300
	hx = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))
	lx = *(*int32)(unsafe.Pointer(&x))
	k = hx & int32(2146435072) >> int32(20)
	if k == int32(0) {
		if lx|hx&int32(2147483647) == int32(0) {
			return x
		}
		x *= two54
		hx = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))
		k = hx&int32(2146435072)>>int32(20) - int32(54)
		if n < float64(-50000) {
			return tiny * x
		}
	}
	if k == int32(2047) {
		return x + x
	}
	k = int32(float64(k) + n)
	if k > int32(2046) {
		return huge * PX_copysign(huge, x)
	}
	if k > int32(0) {
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4)) = int32(uint32(hx)&uint32(2148532223) | uint32(k<<int32(20)))
		return x
	}
	if k <= -54 {
		if n > float64(int32(50000)) {
			return huge * PX_copysign(huge, x)
		} else {
			return tiny * PX_copysign(tiny, x)
		}
	}
	k += int32(54)
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4)) = int32(uint32(hx)&uint32(2148532223) | uint32(k<<int32(20)))
	return x * twom54
}
func __px_pow_dd(x float64, y float64) float64 {
	var z float64
	var ax float64
	var z_h float64
	var z_l float64
	var p_h float64
	var p_l float64
	var y1 float64
	var t1 float64
	var t2 float64
	var r float64
	var s float64
	var t float64
	var u float64
	var v float64
	var w float64
	var i0 int32
	var i1 int32
	var i int32
	var j int32
	var k int32
	var yisint int32
	var n int32
	var hx int32
	var hy int32
	var ix int32
	var iy int32
	var lx uint32
	var ly uint32
	i0 = *(*int32)(unsafe.Pointer(&_cgos_one_PX_Typedef))>>int32(29) ^ int32(1)
	i1 = int32(1) - i0
	hx = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))
	lx = uint32(*(*int32)(unsafe.Pointer(&x)))
	hy = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&y)))) + uintptr(int32(1))*4))
	ly = uint32(*(*int32)(unsafe.Pointer(&y)))
	ix = hx & int32(2147483647)
	iy = hy & int32(2147483647)
	if uint32(iy)|ly == uint32(0) {
		return _cgos_one_PX_Typedef
	}
	if ix > int32(2146435072) || ix == int32(2146435072) && lx != uint32(0) || iy > int32(2146435072) || iy == int32(2146435072) && ly != uint32(0) {
		return x + y
	}
	yisint = int32(0)
	if hx < int32(0) {
		if iy >= int32(1128267776) {
			yisint = int32(2)
		} else if iy >= int32(1072693248) {
			k = iy>>int32(20) - int32(1023)
			if k > int32(20) {
				j = int32(ly >> (int32(52) - k))
				if uint32(j<<(int32(52)-k)) == ly {
					yisint = int32(2) - j&int32(1)
				}
			} else if ly == uint32(0) {
				j = iy >> (int32(20) - k)
				if j<<(int32(20)-k) == iy {
					yisint = int32(2) - j&int32(1)
				}
			}
		}
	}
	if ly == uint32(0) {
		if iy == int32(2146435072) {
			if uint32(ix-int32(1072693248))|lx == uint32(0) {
				return y - y
			} else if ix >= int32(1072693248) {
				return func() float64 {
					if hy >= int32(0) {
						return y
					} else {
						return _cgos_zero_PX_Typedef
					}
				}()
			} else {
				return func() float64 {
					if hy < int32(0) {
						return -y
					} else {
						return _cgos_zero_PX_Typedef
					}
				}()
			}
		}
		if iy == int32(1072693248) {
			if hy < int32(0) {
				return _cgos_one_PX_Typedef / x
			} else {
				return x
			}
		}
		if hy == int32(1073741824) {
			return x * x
		}
		if hy == int32(1071644672) {
			if hx >= int32(0) {
				return PX_sqrtd(x)
			}
		}
	}
	ax = func() float64 {
		if x > float64(int32(0)) {
			return x
		} else {
			return -x
		}
	}()
	if lx == uint32(0) {
		if ix == int32(2146435072) || ix == int32(0) || ix == int32(1072693248) {
			z = ax
			if hy < int32(0) {
				z = _cgos_one_PX_Typedef / z
			}
			if hx < int32(0) {
				if ix-int32(1072693248)|yisint == int32(0) {
					z = (z - z) / (z - z)
				} else if yisint == int32(1) {
					z = -z
				}
			}
			return z
		}
	}
	n = hx>>int32(31) + int32(1)
	if n|yisint == int32(0) {
		return (x - x) / (x - x)
	}
	s = _cgos_one_PX_Typedef
	if n|(yisint-int32(1)) == int32(0) {
		s = -_cgos_one_PX_Typedef
	}
	if iy > int32(1105199104) {
		if iy > int32(1139802112) {
			if ix <= int32(1072693247) {
				return func() float64 {
					if hy < int32(0) {
						return _cgos_huge_PX_Typedef
					} else {
						return _cgos_tiny_PX_Typedef
					}
				}()
			}
			if ix >= int32(1072693248) {
				return func() float64 {
					if hy > int32(0) {
						return _cgos_huge_PX_Typedef
					} else {
						return _cgos_tiny_PX_Typedef
					}
				}()
			}
		}
		if ix < int32(1072693247) {
			return func() float64 {
				if hy < int32(0) {
					return s * _cgos_huge_PX_Typedef * _cgos_huge_PX_Typedef
				} else {
					return s * _cgos_tiny_PX_Typedef * _cgos_tiny_PX_Typedef
				}
			}()
		}
		if ix > int32(1072693248) {
			return func() float64 {
				if hy > int32(0) {
					return s * _cgos_huge_PX_Typedef * _cgos_huge_PX_Typedef
				} else {
					return s * _cgos_tiny_PX_Typedef * _cgos_tiny_PX_Typedef
				}
			}()
		}
		t = ax - _cgos_one_PX_Typedef
		w = t * t * (0.5 - t*(0.33333333333333331-t*0.25))
		u = _cgos_ivln2_h_PX_Typedef * t
		v = t*_cgos_ivln2_l_PX_Typedef - w*_cgos_ivln2_PX_Typedef
		t1 = u + v
		*(*int32)(unsafe.Pointer(&t1)) = int32(0)
		t2 = v - (t1 - u)
	} else {
		var ss float64
		var s2 float64
		var s_h float64
		var s_l float64
		var t_h float64
		var t_l float64
		n = int32(0)
		if ix < int32(1048576) {
			ax *= _cgos_two53_PX_Typedef
			n -= int32(53)
			ix = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&ax)))) + uintptr(int32(1))*4))
		}
		n += ix>>int32(20) - int32(1023)
		j = ix & int32(1048575)
		ix = j | int32(1072693248)
		if j <= int32(235662) {
			k = int32(0)
		} else if j < int32(767610) {
			k = int32(1)
		} else {
			k = int32(0)
			n += int32(1)
			ix -= int32(1048576)
		}
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&ax)))) + uintptr(int32(1))*4)) = ix
		u = ax - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_bp_PX_Typedef)))) + uintptr(k)*8))
		v = _cgos_one_PX_Typedef / (ax + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_bp_PX_Typedef)))) + uintptr(k)*8)))
		ss = u * v
		s_h = ss
		*(*int32)(unsafe.Pointer(&s_h)) = int32(0)
		t_h = _cgos_zero_PX_Typedef
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&t_h)))) + uintptr(int32(1))*4)) = ix>>int32(1) | int32(536870912) + int32(524288) + k<<int32(18)
		t_l = ax - (t_h - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_bp_PX_Typedef)))) + uintptr(k)*8)))
		s_l = v * (u - s_h*t_h - s_h*t_l)
		s2 = ss * ss
		r = s2 * s2 * (_cgos_L1_PX_Typedef + s2*(_cgos_L2_PX_Typedef+s2*(_cgos_L3_PX_Typedef+s2*(_cgos_L4_PX_Typedef+s2*(_cgos_L5_PX_Typedef+s2*_cgos_L6_PX_Typedef)))))
		r += s_l * (s_h + ss)
		s2 = s_h * s_h
		t_h = 3 + s2 + r
		*(*int32)(unsafe.Pointer(&t_h)) = int32(0)
		t_l = r - (t_h - 3 - s2)
		u = s_h * t_h
		v = s_l*t_h + t_l*ss
		p_h = u + v
		*(*int32)(unsafe.Pointer(&p_h)) = int32(0)
		p_l = v - (p_h - u)
		z_h = _cgos_cp_h_PX_Typedef * p_h
		z_l = _cgos_cp_l_PX_Typedef*p_h + p_l*_cgos_cp_PX_Typedef + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_dp_l_PX_Typedef)))) + uintptr(k)*8))
		t = float64(n)
		t1 = z_h + z_l + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_dp_h_PX_Typedef)))) + uintptr(k)*8)) + t
		*(*int32)(unsafe.Pointer(&t1)) = int32(0)
		t2 = z_l - (t1 - t - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_dp_h_PX_Typedef)))) + uintptr(k)*8)) - z_h)
	}
	y1 = y
	*(*int32)(unsafe.Pointer(&y1)) = int32(0)
	p_l = (y-y1)*t1 + y*t2
	p_h = y1 * t1
	z = p_l + p_h
	j = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&z)))) + uintptr(int32(1))*4))
	i = *(*int32)(unsafe.Pointer(&z))
	if j >= int32(1083179008) {
		if j-int32(1083179008)|i != int32(0) {
			return s * _cgos_huge_PX_Typedef * _cgos_huge_PX_Typedef
		} else if p_l+_cgos_ovt_PX_Typedef > z-p_h {
			return s * _cgos_huge_PX_Typedef * _cgos_huge_PX_Typedef
		}
	} else if j&int32(2147483647) >= int32(1083231232) {
		if uint32(j)-uint32(3230714880)|uint32(i) != uint32(0) {
			return s * _cgos_tiny_PX_Typedef * _cgos_tiny_PX_Typedef
		} else if p_l <= z-p_h {
			return s * _cgos_tiny_PX_Typedef * _cgos_tiny_PX_Typedef
		}
	}
	i = j & int32(2147483647)
	k = i>>int32(20) - int32(1023)
	n = int32(0)
	if i > int32(1071644672) {
		n = j + int32(1048576)>>(k+int32(1))
		k = n&int32(2147483647)>>int32(20) - int32(1023)
		t = _cgos_zero_PX_Typedef
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&t)))) + uintptr(int32(1))*4)) = n & ^(int32(1048575) >> k)
		n = (n&int32(1048575) | int32(1048576)) >> (int32(20) - k)
		if j < int32(0) {
			n = -n
		}
		p_h -= t
	}
	t = p_l + p_h
	*(*int32)(unsafe.Pointer(&t)) = int32(0)
	u = t * _cgos_lg2_h_PX_Typedef
	v = (p_l-(t-p_h))*_cgos_lg2_PX_Typedef + t*_cgos_lg2_l_PX_Typedef
	z = u + v
	w = v - (z - u)
	t = z * z
	t1 = z - t*(_cgos_P1_PX_Typedef+t*(_cgos_P2_PX_Typedef+t*(_cgos_P3_PX_Typedef+t*(_cgos_P4_PX_Typedef+t*_cgos_P5_PX_Typedef))))
	r = z*t1/(t1-_cgos_two_PX_Typedef) - (w + z*w)
	z = _cgos_one_PX_Typedef - (r - z)
	j = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&z)))) + uintptr(int32(1))*4))
	j += n << int32(20)
	if j>>int32(20) <= int32(0) {
		z = PX_scalbn(z, float64(n))
	} else {
		*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&z)))) + uintptr(int32(1))*4)) += n << int32(20)
	}
	func() int {
		func() int {
			_ = i0
			return 0
		}()
		return func() int {
			_ = i1
			return 0
		}()
	}()
	return s * z
}

var _cgos_bp_PX_Typedef [2]float64 = [2]float64{1, 1.5}
var _cgos_dp_h_PX_Typedef [2]float64 = [2]float64{0, 0.58496248722076416}
var _cgos_dp_l_PX_Typedef [2]float64 = [2]float64{0, 1.350039202129749e-8}
var _cgos_zero_PX_Typedef float64 = 0
var _cgos_one_PX_Typedef float64 = 1
var _cgos_two_PX_Typedef float64 = 2
var _cgos_two53_PX_Typedef float64 = 9007199254740992
var _cgos_huge_PX_Typedef float64 = 1.0000000000000001e+300
var _cgos_tiny_PX_Typedef float64 = 1.0e-300
var _cgos_L1_PX_Typedef float64 = 0.59999999999999465
var _cgos_L2_PX_Typedef float64 = 0.42857142857855018
var _cgos_L3_PX_Typedef float64 = 0.33333332981837743
var _cgos_L4_PX_Typedef float64 = 0.27272812380853401
var _cgos_L5_PX_Typedef float64 = 0.23066074577556175
var _cgos_L6_PX_Typedef float64 = 0.20697501780033842
var _cgos_P1_PX_Typedef float64 = 0.16666666666666602
var _cgos_P2_PX_Typedef float64 = -0.0027777777777015593
var _cgos_P3_PX_Typedef float64 = 6.6137563214379344e-5
var _cgos_P4_PX_Typedef float64 = -1.6533902205465252e-6
var _cgos_P5_PX_Typedef float64 = 4.1381367970572385e-8
var _cgos_lg2_PX_Typedef float64 = 0.69314718055994529
var _cgos_lg2_h_PX_Typedef float64 = 0.69314718246459961
var _cgos_lg2_l_PX_Typedef float64 = -1.904654299957768e-9
var _cgos_ovt_PX_Typedef float64 = 8.0085662595372941e-17
var _cgos_cp_PX_Typedef float64 = 0.96179669392597555
var _cgos_cp_h_PX_Typedef float64 = 0.96179670095443726
var _cgos_cp_l_PX_Typedef float64 = -7.0284616509527583e-9
var _cgos_ivln2_PX_Typedef float64 = 1.4426950408889634
var _cgos_ivln2_h_PX_Typedef float64 = 1.4426950216293335
var _cgos_ivln2_l_PX_Typedef float64 = 1.9259629911266175e-8

func PX_exp(x float64) float64 {
	return __px_pow_dd(2.7182818284590451, x)
}
func PX_tanh(x float64) float64 {
	var ex float64
	var eix float64
	ex = PX_exp(x)
	eix = float64(int32(1)) / ex
	return (ex - eix) / (ex + eix)
}
func PX_sigmoid(x float64) float64 {
	var eix float64
	eix = PX_exp(-x)
	return float64(int32(1)) / (float64(int32(1)) + eix)
}
func PX_ReLU(x float64) float64 {
	return func() float64 {
		if x <= float64(int32(0)) {
			return float64(int32(0))
		} else {
			return x
		}
	}()
}
func PX_sind(x float64) float64 {
	var it float64
	var term float64
	var result float64
	it = x / (float64(int32(2)) * 3.1415926535897931)
	x = x - float64(int32(2))*3.1415926535897931*float64(int32(it))
	term = x
	result = float64(int32(0))
	result += term
	term *= -x * x / float64(6)
	result += term
	term *= -x * x / float64(20)
	result += term
	term *= -x * x / float64(42)
	result += term
	term *= -x * x / float64(72)
	result += term
	term *= -x * x / float64(110)
	result += term
	term *= -x * x / float64(156)
	result += term
	term *= -x * x / float64(210)
	result += term
	term *= -x * x / float64(272)
	result += term
	term *= -x * x / float64(342)
	result += term
	term *= -x * x / float64(420)
	result += term
	term *= -x * x / float64(506)
	result += term
	term *= -x * x / float64(600)
	result += term
	term *= -x * x / float64(702)
	result += term
	term *= -x * x / float64(812)
	result += term
	term *= -x * x / float64(930)
	result += term
	term *= -x * x / float64(1056)
	return result
}
func PX_sinc(i float64) float64 {
	return func() float64 {
		if i != 0 {
			return PX_sind(i*3.1415926535897931) / (i * 3.1415926535897931)
		} else {
			return float64(int32(1))
		}
	}()
}
func PX_sinc_interpolate(x *float64, size int32, d float64) float64 {
	var i int32
	var sum float64 = float64(int32(0))
	for i = int32(0); i < size; i++ {
		sum += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) * PX_sinc(d-float64(i))
	}
	return sum
}
func PX_cosd(radian float64) float64 {
	return PX_sind(3.1415926535897931/float64(int32(2)) - radian)
}
func PX_sin_radian(radian float32) float32 {
	return float32(PX_sind(float64(radian)))
}
func PX_cos_radian(radian float32) float32 {
	return PX_sin_radian(float32(3.1415926535897931/float64(int32(2)) - float64(radian)))
}
func PX_tan_radian(radian float32) float32 {
	return PX_sin_radian(radian) / PX_cos_radian(radian)
}
func PX_sin_angle(angle float32) float32 {
	angle -= float32(int32(angle) / int32(360) * int32(360))
	return float32(PX_sin_radian(angle * 0.0174532924))
}
func PX_cos_angle(angle float32) float32 {
	angle -= float32(int32(angle) / int32(360) * int32(360))
	return PX_cos_radian(angle * 0.0174532924)
}
func PX_tan_angle(angle float32) float32 {
	return PX_sin_angle(angle) / PX_cos_angle(angle)
}
func PX_Point_sin(v Struct__px_point) float32 {
	return v.y / PX_sqrt(v.x*v.x+v.y*v.y)
}
func PX_Point_cos(v Struct__px_point) float32 {
	return v.x / PX_sqrt(v.x*v.x+v.y*v.y)
}
func PX_BufferToHexString(data *uint8, size int32, hex_str *int8) {
	var i int32
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(int32(0)))) = int8(0)
	for i = int32(0); i < size; i++ {
		if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))) == int32(0) {
			PX_strcat((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2)))), (*int8)(unsafe.Pointer(&[3]int8{'0', '0', '\x00'})))
		} else if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))) < int32(16) {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(i*int32(2)))) = int8('0')
			PX_itoa(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2))))))+uintptr(int32(1)))), int32(3), int32(16))
		} else {
			PX_itoa(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str))+uintptr(i*int32(2)))), int32(3), int32(16))
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(hex_str)) + uintptr(i*int32(2)))) = int8(0)
}
func PX_HexStringToBuffer(hex_str *int8, data *uint8) int32 {
	var i int32 = int32(0)
	var hex [3]int8 = [3]int8{int8(0)}
	if PX_strlen(hex_str)&int32(1) != 0 {
		return int32(0)
	}
	for *hex_str != 0 {
		if int32(*hex_str) == ' ' {
			*(*uintptr)(unsafe.Pointer(&hex_str))++
			continue
		}
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&hex)))) + uintptr(int32(0)))) = *hex_str
		*(*uintptr)(unsafe.Pointer(&hex_str))++
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&hex)))) + uintptr(int32(1)))) = *hex_str
		*(*uintptr)(unsafe.Pointer(&hex_str))++
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i))) = uint8(PX_htoi((*int8)(unsafe.Pointer(&hex))))
		i++
	}
	return i
}

var _cgos_Tiny_PX_Typedef float64 = float64(2.2250738585072014e-308)

func _cgos_Tail_PX_Typedef(x float64) float64 {
	var HalfPi float64 = 3.1415926535897931 / float64(int32(2))
	var p03 float64 = float64(-0.33333333333317677)
	var p05 float64 = float64(0.19999999997148041)
	var p07 float64 = float64(-0.1428571408980793)
	var p09 float64 = float64(0.11111104059171241)
	var p11 float64 = float64(-0.090907570635999899)
	var p13 float64 = float64(0.076901890547487486)
	var p15 float64 = float64(-0.066468502185766831)
	var p17 float64 = float64(0.057557260576963837)
	var p19 float64 = float64(-0.047097798903229605)
	var p21 float64 = float64(0.031306840249638944)
	var p23 float64 = float64(-0.011836408476918686)
	var p000 float64 = float64(1.5707963267948966)
	var p001 float64 = float64(6.123233995736766e-17)
	var y float64
	var y2 float64
	if 9007199254740992 <= x {
		return HalfPi + _cgos_Tiny_PX_Typedef
	}
	y = float64(int32(1)) / x
	y2 = y * y
	return p001 - (((((((((((p23*y2+p21)*y2+p19)*y2+p17)*y2+p15)*y2+p13)*y2+p11)*y2+p09)*y2+p07)*y2+p05)*y2+p03)*y2*y + y) + p000
}
func _cgos_atani0_PX_Typedef(x float64) float64 {
	var p03 float64 = float64(-0.3333333333333301)
	var p05 float64 = float64(0.19999999999908513)
	var p07 float64 = float64(-0.1428571427670465)
	var p09 float64 = float64(0.11111111066822131)
	var p11 float64 = float64(-0.090908963576435756)
	var p13 float64 = float64(0.076920744390691756)
	var p15 float64 = float64(-0.066638061688310171)
	var p17 float64 = float64(0.058582174411949621)
	var p19 float64 = float64(-0.05121467553745504)
	var p21 float64 = float64(0.041846137206663775)
	var p23 float64 = float64(-0.027399903124251841)
	var p25 float64 = float64(0.010089798028531199)
	var x2 float64 = x * x
	return (((((((((((p25*x2+p23)*x2+p21)*x2+p19)*x2+p17)*x2+p15)*x2+p13)*x2+p11)*x2+p09)*x2+p07)*x2+p05)*x2+p03)*x2*x + x
}
func _cgos_atani1_PX_Typedef(x float64) float64 {
	var p00 float64 = float64(0.55859931534356555)
	var p01 float64 = float64(0.7191011235955026)
	var p02 float64 = float64(-0.3231915162226941)
	var p03 float64 = float64(0.021304010058557601)
	var p04 float64 = float64(0.1018414372498045)
	var p05 float64 = float64(-0.082426134720760047)
	var p06 float64 = float64(0.012919562983764597)
	var p07 float64 = float64(0.032383633298517352)
	var p08 float64 = float64(-0.032438298163273233)
	var p09 float64 = float64(0.0078007868315617766)
	var p10 float64 = float64(0.012361337366866022)
	var p11 float64 = float64(-0.014421128514667323)
	var p12 float64 = float64(0.004164927187425534)
	var y float64 = x - float64(0.62500000000000433)
	return (((((((((((p12*y+p11)*y+p10)*y+p09)*y+p08)*y+p07)*y+p06)*y+p05)*y+p04)*y+p03)*y+p02)*y+p01)*y + p00
}
func _cgos_atani2_PX_Typedef(x float64) float64 {
	var p00 float64 = float64(0.71882999968450156)
	var p01 float64 = float64(0.56637168135360871)
	var p02 float64 = float64(-0.28067977129496502)
	var p03 float64 = float64(0.078538292532232434)
	var p04 float64 = float64(0.021102089406107673)
	var p05 float64 = float64(-0.043421395006427825)
	var p06 float64 = float64(0.027896540867196881)
	var p07 float64 = float64(-0.006133549533512985)
	var p08 float64 = float64(-0.0065320891914041248)
	var p09 float64 = float64(0.0084575497048781853)
	var p10 float64 = float64(-0.0044917438891326966)
	var p11 float64 = float64(1.0931200219364778e-4)
	var y float64 = x - float64(0.87500000011101731)
	return ((((((((((p11*y+p10)*y+p09)*y+p08)*y+p07)*y+p06)*y+p05)*y+p04)*y+p03)*y+p02)*y+p01)*y + p00
}
func _cgos_atani3_PX_Typedef(x float64) float64 {
	var p00 float64 = float64(0.86217005466722441)
	var p01 float64 = float64(0.42352941176470776)
	var p02 float64 = float64(-0.20927335640138578)
	var p03 float64 = float64(0.078081823733023734)
	var p04 float64 = float64(-0.013555699763751467)
	var p05 float64 = float64(-0.0091249925913829862)
	var p06 float64 = float64(0.011342191383891984)
	var p07 float64 = float64(-0.0068469987314645911)
	var p08 float64 = float64(0.0023178361431887792)
	var p09 float64 = float64(2.1926428334741529e-4)
	var p10 float64 = float64(-9.8086742285925571e-4)
	var p11 float64 = float64(8.098046818543229e-4)
	var p12 float64 = float64(-3.740695734151908e-4)
	var y float64 = x - float64(1.1666666666666208)
	return (((((((((((p12*y+p11)*y+p10)*y+p09)*y+p08)*y+p07)*y+p06)*y+p05)*y+p04)*y+p03)*y+p02)*y+p01)*y + p00
}
func _cgos_atani4_PX_Typedef(x float64) float64 {
	var p00 float64 = float64(0.98279372324732927)
	var p01 float64 = float64(0.30769230769230727)
	var p02 float64 = float64(-0.1420118343195424)
	var p03 float64 = float64(0.055833712638714519)
	var p04 float64 = float64(-0.01680613422985678)
	var p05 float64 = float64(0.0021029213489315839)
	var p06 float64 = float64(0.0018297786863776509)
	var p07 float64 = float64(-0.0019099127535115346)
	var p08 float64 = float64(0.00112049030402938)
	var p09 float64 = float64(-4.6246854128182552e-4)
	var p10 float64 = float64(1.0416195839921295e-4)
	var p11 float64 = float64(3.0815753227650312e-5)
	var y float64 = x - 1.5000000000000067
	return ((((((((((p11*y+p10)*y+p09)*y+p08)*y+p07)*y+p06)*y+p05)*y+p04)*y+p03)*y+p02)*y+p01)*y + p00
}
func _cgos_atani5_PX_Typedef(x float64) float64 {
	var p00 float64 = float64(1.0714496050638533)
	var p01 float64 = float64(0.22929936310012725)
	var p02 float64 = float64(-0.096393362838219906)
	var p03 float64 = float64(0.036503334859833564)
	var p04 float64 = float64(-0.011966558088390809)
	var p05 float64 = float64(0.0030267288042871818)
	var p06 float64 = float64(-2.9135544344454343e-4)
	var p07 float64 = float64(-2.8572536016898764e-4)
	var p08 float64 = float64(2.6032522025351611e-4)
	var p09 float64 = float64(-1.4508835052352259e-4)
	var p10 float64 = float64(6.1205254960895611e-5)
	var y float64 = x - float64(1.8333333331112962)
	return (((((((((p10*y+p09)*y+p08)*y+p07)*y+p06)*y+p05)*y+p04)*y+p03)*y+p02)*y+p01)*y + p00
}
func PX_atan(x float64) float64 {
	if x < float64(int32(0)) {
		if x < float64(-1) {
			if x < float64(-5)/3 {
				if x < float64(-2) {
					return -_cgos_Tail_PX_Typedef(-x)
				} else {
					return -_cgos_atani5_PX_Typedef(-x)
				}
			} else if x < float64(-4)/3 {
				return -_cgos_atani4_PX_Typedef(-x)
			} else {
				return -_cgos_atani3_PX_Typedef(-x)
			}
		} else if x < -0.5 {
			if x < -0.75 {
				return -_cgos_atani2_PX_Typedef(-x)
			} else {
				return -_cgos_atani1_PX_Typedef(-x)
			}
		} else if x < float64(1.3538603431225864e-8) {
			return _cgos_atani0_PX_Typedef(x)
		} else if x <= float64(2.2250738585072014e-308) {
			return (_cgos_Tiny_PX_Typedef + float64(int32(1))) * x
		} else if x == float64(int32(0)) {
			return x
		} else {
			return x*_cgos_Tiny_PX_Typedef + x
		}
	} else if x <= float64(int32(1)) {
		if x <= 0.5 {
			if x <= float64(1.3538603431225864e-8) {
				if x < float64(2.2250738585072014e-308) {
					if x == float64(int32(0)) {
						return x
					} else {
						return x*_cgos_Tiny_PX_Typedef + x
					}
				} else {
					return (_cgos_Tiny_PX_Typedef + float64(int32(1))) * x
				}
			} else {
				return _cgos_atani0_PX_Typedef(x)
			}
		} else if x <= 0.75 {
			return _cgos_atani1_PX_Typedef(x)
		} else {
			return _cgos_atani2_PX_Typedef(x)
		}
	} else if x <= float64(int32(5))/3 {
		if x <= float64(int32(4))/3 {
			return _cgos_atani3_PX_Typedef(x)
		} else {
			return _cgos_atani4_PX_Typedef(x)
		}
	} else if x <= float64(int32(2)) {
		return _cgos_atani5_PX_Typedef(x)
	} else {
		return _cgos_Tail_PX_Typedef(x)
	}
	return 0
}
func PX_atan2(y float64, x float64) float64 {
	if x > float64(int32(0)) {
		return PX_atan(y / x)
	} else if y >= float64(int32(0)) && x < float64(int32(0)) {
		return PX_atan(y/x) + 3.1415926535897931
	} else if y < float64(int32(0)) && x < float64(int32(0)) {
		return PX_atan(y/x) - 3.1415926535897931
	} else if y > float64(int32(0)) && x == float64(int32(0)) {
		return float64(3.1415926535897931 / float64(int32(2)))
	} else if y < float64(int32(0)) && x == float64(int32(0)) {
		return float64(-1) * 3.1415926535897931 / float64(int32(2))
	} else {
		return float64(int32(0))
	}
	return 0
}
func PX_asin(x float64) float64 {
	return PX_atan2(x, PX_sqrtd(1-x*x))
}
func PX_acos(x float64) float64 {
	return PX_atan2(PX_sqrtd(1-x*x), x)
}
func PX_STRINGFORMAT_INT(_i int32) _cgoa_8_PX_Typedef {
	var fmt _cgoa_8_PX_Typedef
	fmt.type_ = int32(0)
	*(*int32)(unsafe.Pointer(&fmt._cgoa_9_PX_Typedef)) = _i
	return fmt
}
func PX_STRINGFORMAT_FLOAT(_f float32) _cgoa_8_PX_Typedef {
	var fmt _cgoa_8_PX_Typedef
	fmt.type_ = int32(1)
	*(*float32)(unsafe.Pointer(&fmt._cgoa_9_PX_Typedef)) = _f
	return fmt
}
func PX_STRINGFORMAT_STRING(_s *int8) _cgoa_8_PX_Typedef {
	var fmt _cgoa_8_PX_Typedef
	fmt.type_ = int32(2)
	fmt._pstring = _s
	return fmt
}
func PX_sprintf8(_out_str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef, _4 _cgoa_8_PX_Typedef, _5 _cgoa_8_PX_Typedef, _6 _cgoa_8_PX_Typedef, _7 _cgoa_8_PX_Typedef, _8 _cgoa_8_PX_Typedef) int32 {
	var length int32 = int32(0)
	var p *int8 = nil
	var tret _cgoa_1_PX_Typedef
	var pstringfmt _cgoa_8_PX_Typedef
	var precision int32 = int32(3)
	if !(_out_str != nil) || !(str_size != 0) {
		for p = fmt; *p != 0; *(*uintptr)(unsafe.Pointer(&p))++ {
			if int32(*p) != '%' {
				length++
				continue
			}
			switch int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) {
			case '1':
				pstringfmt = _1
				break
			case '2':
				pstringfmt = _2
				break
			case '3':
				pstringfmt = _3
				break
			case '4':
				pstringfmt = _4
				break
			case '5':
				pstringfmt = _5
				break
			case '6':
				pstringfmt = _6
				break
			case '7':
				pstringfmt = _7
				break
			case '8':
				pstringfmt = _8
				break
			default:
				length++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))))) == '.' && PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) != 0 {
				precision = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) - '0'
				*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(3))
			} else {
				*(*uintptr)(unsafe.Pointer(&p))++
			}
			switch uint32(pstringfmt.type_) {
			case uint32(0):
				tret = PX_itos(*(*int32)(unsafe.Pointer(&pstringfmt._cgoa_9_PX_Typedef)), int32(10))
				length += PX_strlen((*int8)(unsafe.Pointer(&(&tret).data)))
				break
			case uint32(1):
				tret = PX_ftos(*(*float32)(unsafe.Pointer(&pstringfmt._cgoa_9_PX_Typedef)), precision)
				length += PX_strlen((*int8)(unsafe.Pointer(&(&tret).data)))
				break
			case uint32(2):
				length += PX_strlen(pstringfmt._pstring)
				break
			default:
				return int32(0)
			}
		}
		return length
	}
	PX_memset(unsafe.Pointer(_out_str), uint8(0), str_size)
	for p = fmt; *p != 0; *(*uintptr)(unsafe.Pointer(&p))++ {
		if int32(*p) != '%' {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out_str)) + uintptr(length))) = *p
			length++
			if length >= str_size {
				PX_ASSERT()
			}
			continue
		}
		switch int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(1))))) {
		case '1':
			pstringfmt = _1
			break
		case '2':
			pstringfmt = _2
			break
		case '3':
			pstringfmt = _3
			break
		case '4':
			pstringfmt = _4
			break
		case '5':
			pstringfmt = _5
			break
		case '6':
			pstringfmt = _6
			break
		case '7':
			pstringfmt = _7
			break
		case '8':
			pstringfmt = _8
			break
		default:
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out_str)) + uintptr(length))) = *p
			length++
			if length >= str_size {
				PX_ASSERT()
			}
			continue
		}
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(2))))) == '.' && PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) != 0 {
			precision = int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(int32(3))))) - '0'
			*(*uintptr)(unsafe.Pointer(&p)) += uintptr(int32(3))
		} else {
			*(*uintptr)(unsafe.Pointer(&p))++
		}
		switch uint32(pstringfmt.type_) {
		case uint32(0):
			tret = PX_itos(*(*int32)(unsafe.Pointer(&pstringfmt._cgoa_9_PX_Typedef)), int32(10))
			if length+PX_strlen((*int8)(unsafe.Pointer(&tret.data))) < str_size {
				PX_strcat(_out_str, (*int8)(unsafe.Pointer(&tret.data)))
				length += PX_strlen((*int8)(unsafe.Pointer(&tret.data)))
			} else {
				return length
			}
			break
		case uint32(1):
			tret = PX_ftos(*(*float32)(unsafe.Pointer(&pstringfmt._cgoa_9_PX_Typedef)), precision)
			if length+PX_strlen((*int8)(unsafe.Pointer(&tret.data))) < str_size {
				PX_strcat(_out_str, (*int8)(unsafe.Pointer(&tret.data)))
				length += PX_strlen((*int8)(unsafe.Pointer(&tret.data)))
			} else {
				return length
			}
			break
		case uint32(2):
			if length+PX_strlen(pstringfmt._pstring) < str_size {
				PX_strcat(_out_str, pstringfmt._pstring)
				length += PX_strlen(pstringfmt._pstring)
			} else {
				return length
			}
			break
		default:
			return int32(0)
		}
	}
	if length >= str_size {
		PX_ASSERT()
	}
	return length
}
func PX_sprintf7(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef, _4 _cgoa_8_PX_Typedef, _5 _cgoa_8_PX_Typedef, _6 _cgoa_8_PX_Typedef, _7 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, _6, _7, PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf6(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef, _4 _cgoa_8_PX_Typedef, _5 _cgoa_8_PX_Typedef, _6 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, _6, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf5(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef, _4 _cgoa_8_PX_Typedef, _5 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, _5, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf4(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef, _4 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, _4, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf3(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef, _3 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, _3, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf2(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef, _2 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, _2, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf1(str *int8, str_size int32, fmt *int8, _1 _cgoa_8_PX_Typedef) int32 {
	return PX_sprintf8(str, str_size, fmt, _1, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_sprintf0(str *int8, str_size int32, fmt *int8) int32 {
	return PX_sprintf8(str, str_size, fmt, PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)), PX_STRINGFORMAT_INT(int32(0)))
}
func PX_MatrixZero(Mat *Struct__px_matrix) {
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4)) = float32(int32(0))
}
func PX_MatrixIdentity(Mat *Struct__px_matrix) {
	*(*float32)(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 4)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 20)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 40)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixMultiply(Mat1 Struct__px_matrix, Mat2 Struct__px_matrix) Struct__px_matrix {
	var ptmat Struct__px_matrix
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(0))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(1))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(2))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(0))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(0))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(1))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(1))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(2))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(2))*16)))))) + uintptr(int32(3))*4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(int32(3))*16)))))) + uintptr(int32(3))*4))
	return ptmat
}
func PX_MatrixAdd(Mat1 Struct__px_matrix, Mat2 Struct__px_matrix) Struct__px_matrix {
	var ptmat Struct__px_matrix
	*(*float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) = *(*float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + *(*float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 4)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 8)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 8)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 8))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 12)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 12)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 12))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 16)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 16)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 16))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 20)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 20)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 20))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 24)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 24)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 24))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 28)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 28)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 28))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 32)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 32)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 32))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 36)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 36)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 36))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 40)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 40)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 40))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 44)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 44)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 44))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 48)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 48)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 48))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 52)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 52)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 52))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 56)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 56)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 56))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 60)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 60)) + *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 60))
	return ptmat
}
func PX_MatrixSub(Mat1 Struct__px_matrix, Mat2 Struct__px_matrix) Struct__px_matrix {
	var ptmat Struct__px_matrix
	*(*float32)(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) = *(*float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) - *(*float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 4)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 8)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 8)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 8))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 12)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 12)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 12))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 16)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 16)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 16))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 20)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 20)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 20))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 24)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 24)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 24))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 28)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 28)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 28))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 32)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 32)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 32))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 36)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 36)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 36))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 40)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 40)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 40))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 44)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 44)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 44))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 48)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 48)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 48))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 52)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 52)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 52))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 56)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 56)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 56))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&ptmat._cgoa_2_PX_Typedef)) + 60)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)) + 60)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)) + 60))
	return ptmat
}
func PX_MatrixEqual(Mat1 Struct__px_matrix, Mat2 Struct__px_matrix) int32 {
	var i int32
	var j int32
	for i = int32(0); i < int32(4); i++ {
		for j = int32(0); j < int32(4); j++ {
			if *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat1._cgoa_2_PX_Typedef)))))) + uintptr(i)*16)))))) + uintptr(j)*4)) != *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&Mat2._cgoa_2_PX_Typedef)))))) + uintptr(i)*16)))))) + uintptr(j)*4)) {
				return int32(0)
			}
		}
	}
	return int32(1)
}
func PX_MatrixRotateVector(mat *Struct__px_matrix, v_base Struct__px_point, v Struct__px_point) {
	var cross Struct__px_point = PX_PointCross(v_base, v)
	var cosv float32 = PX_PointDot(v_base, v) / PX_PointMod(v_base) / PX_PointMod(v)
	var sinv float32 = float32(PX_sqrtd(float64(float32(int32(1)) - cosv*cosv)))
	if cross.z < float32(int32(0)) {
		sinv = -sinv
	}
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = cosv
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = sinv
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = -sinv
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = cosv
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixTranslation(mat *Struct__px_matrix, x float32, y float32, z float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = x
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = y
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = z
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateX(mat *Struct__px_matrix, Angle float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = -PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateY(mat *Struct__px_matrix, Angle float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = -PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateZ(mat *Struct__px_matrix, Angle float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = -PX_sin_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = PX_cos_angle(Angle)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateXRadian(mat *Struct__px_matrix, rad float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(int32(0))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = -PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateYRadian(mat *Struct__px_matrix, rad float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = -PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixRotateZRadian(mat *Struct__px_matrix, rad float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = -PX_sin_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = PX_cos_radian(rad)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = float32(1)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func PX_MatrixScale(mat *Struct__px_matrix, x float32, y float32, z float32) {
	*(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) = x
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) = y
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) = z
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56)) = float32(0)
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60)) = float32(1)
}
func _cgos_ptabs_PX_Typedef(x float32) float32 {
	return func() float32 {
		if x > float32(int32(0)) {
			return x
		} else {
			return -x
		}
	}()
}
func _cgos_Gauss_PX_Typedef(A *[4]float32, B *[4]float32) int32 {
	var i int32
	var j int32
	var k int32
	var max float32
	var temp float32
	var t [4][4]float32
	for i = int32(0); i < int32(4); i++ {
		for j = int32(0); j < int32(4); j++ {
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(j)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*16)))))) + uintptr(j)*4))
		}
	}
	for i = int32(0); i < int32(4); i++ {
		for j = int32(0); j < int32(4); j++ {
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(j)*4)) = func() float32 {
				if i == j {
					return float32(int32(1))
				} else {
					return float32(int32(0))
				}
			}()
		}
	}
	for i = int32(0); i < int32(4); i++ {
		max = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(i)*4))
		k = i
		for j = i + int32(1); j < int32(4); j++ {
			if _cgos_ptabs_PX_Typedef(*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(j)*16)))))) + uintptr(i)*4))) > _cgos_ptabs_PX_Typedef(max) {
				max = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(j)*16)))))) + uintptr(i)*4))
				k = j
			}
		}
		if k != i {
			for j = int32(0); j < int32(4); j++ {
				temp = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(j)*4))
				*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(j)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(k)*16)))))) + uintptr(j)*4))
				*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(k)*16)))))) + uintptr(j)*4)) = temp
				temp = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(j)*4))
				*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(j)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(k)*16)))))) + uintptr(j)*4))
				*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(k)*16)))))) + uintptr(j)*4)) = temp
			}
		}
		if *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(i)*4)) == float32(int32(0)) {
			return int32(0)
		}
		temp = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(i)*4))
		for j = int32(0); j < int32(4); j++ {
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(j)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(j)*4)) / temp
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(j)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(j)*4)) / temp
		}
		for j = int32(0); j < int32(4); j++ {
			if j != i {
				temp = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(j)*16)))))) + uintptr(i)*4))
				for k = int32(0); k < int32(4); k++ {
					*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(j)*16)))))) + uintptr(k)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(j)*16)))))) + uintptr(k)*4)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*[4]float32)(unsafe.Pointer(&t)))) + uintptr(i)*16)))))) + uintptr(k)*4))*temp
					*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(j)*16)))))) + uintptr(k)*4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(j)*16)))))) + uintptr(k)*4)) - *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer((*float32)(unsafe.Pointer(&*(*[4]float32)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*16)))))) + uintptr(k)*4))*temp
				}
			}
		}
	}
	return int32(1)
}
func PX_MatrixInverse(mat *Struct__px_matrix) int32 {
	return _cgos_Gauss_PX_Typedef((*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)))), (*[4]float32)(unsafe.Pointer(&*(*[4][4]float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)))))
}
func PX_MatrixTranspose(matrix *Struct__px_matrix) {
	var mat Struct__px_matrix = *matrix
	*(*float32)(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) = *(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 4)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 8)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 12)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 16)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 20)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 24)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 28)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 32)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 36)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 40)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 44)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 48)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 52)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 56)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44))
	*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&matrix._cgoa_2_PX_Typedef)) + 60)) = *(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60))
}
func PX_COLOR(a uint8, r uint8, g uint8, b uint8) Struct__px_color {
	var color Struct__px_color
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 3)) = a
	*(*uint8)(unsafe.Pointer(&color._argb)) = r
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 1)) = g
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 2)) = b
	return color
}
func PX_ColorInverse(clr Struct__px_color) Struct__px_color {
	var color Struct__px_color
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 3)) = *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&clr._argb)) + 3))
	*(*uint8)(unsafe.Pointer(&color._argb)) = uint8(int32(255) - int32(*(*uint8)(unsafe.Pointer(&clr._argb))))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 1)) = uint8(int32(255) - int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&clr._argb)) + 1))))
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 2)) = uint8(int32(255) - int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&clr._argb)) + 2))))
	return color
}
func PX_ColorIncrease(color *Struct__px_color, inc uint8) {
	if int32(*(*uint8)(unsafe.Pointer(&color._argb)))+int32(inc) >= int32(255) {
		*(*uint8)(unsafe.Pointer(&color._argb)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(&color._argb)) += uint8(int32(inc))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 1)))+int32(inc) >= int32(255) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 1)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 1)) += uint8(int32(inc))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 2)))+int32(inc) >= int32(255) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 2)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color._argb)) + 2)) += uint8(int32(inc))
	}
}
func PX_ColorAdd(color1 Struct__px_color, color2 Struct__px_color) Struct__px_color {
	if int32(*(*uint8)(unsafe.Pointer(&color1._argb)))+int32(*(*uint8)(unsafe.Pointer(&color2._argb))) >= int32(255) {
		*(*uint8)(unsafe.Pointer(&color1._argb)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(&color1._argb)) += uint8(int32(*(*uint8)(unsafe.Pointer(&color2._argb))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)))+int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 1))) >= int32(255) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)) += uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 1))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)))+int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 2))) >= int32(255) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)) += uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 2))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)))+int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 3))) >= int32(255) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)) = uint8(255)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)) += uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 3))))
	}
	return color1
}
func PX_ColorSub(color1 Struct__px_color, color2 Struct__px_color) Struct__px_color {
	if int32(*(*uint8)(unsafe.Pointer(&color1._argb)))-int32(*(*uint8)(unsafe.Pointer(&color2._argb))) < int32(0) {
		*(*uint8)(unsafe.Pointer(&color1._argb)) = uint8(0)
	} else {
		*(*uint8)(unsafe.Pointer(&color1._argb)) -= uint8(int32(*(*uint8)(unsafe.Pointer(&color2._argb))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)))-int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 1))) < int32(0) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)) = uint8(0)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)) -= uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 1))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)))-int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 2))) < int32(0) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)) = uint8(0)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)) -= uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 2))))
	}
	if int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)))-int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 3))) < int32(0) {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)) = uint8(0)
	} else {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)) -= uint8(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color2._argb)) + 3))))
	}
	return color1
}
func PX_ColorEqual(color1 Struct__px_color, color2 Struct__px_color) int32 {
	return func() int32 {
		if *(*uint32)(unsafe.Pointer(&color1._argb)) == *(*uint32)(unsafe.Pointer(&color2._argb)) {
			return 1
		} else {
			return 0
		}
	}()
}
func PX_ColorRGBToHSL(color_rgb Struct__px_color) Struct__px_color_hsl {
	var hsl Struct__px_color_hsl
	var max float32 = float32(int32(0))
	var min float32 = float32(int32(0))
	var r float32
	var g float32
	var b float32
	r = float32(int32(*(*uint8)(unsafe.Pointer(&color_rgb._argb)))) / 255
	g = float32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color_rgb._argb)) + 1)))) / 255
	b = float32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color_rgb._argb)) + 2)))) / 255
	max = r
	max = func() float32 {
		if max < g {
			return g
		} else {
			return max
		}
	}()
	max = func() float32 {
		if max < b {
			return b
		} else {
			return max
		}
	}()
	min = r
	min = func() float32 {
		if min > g {
			return g
		} else {
			return min
		}
	}()
	min = func() float32 {
		if min > b {
			return b
		} else {
			return min
		}
	}()
	hsl.L = (max + min) / 2
	if max == min {
		hsl.S = float32(int32(0))
	} else {
		hsl.S = func() float32 {
			if hsl.L < 0.5 {
				return (max - min) / (max + min)
			} else {
				return (max - min) / (2 - max - min)
			}
		}()
	}
	if r == max {
		hsl.H = (g - b) * 1 / (max - min)
	}
	if g == max {
		hsl.H = 2 + (b-r)*1/(max-min)
	}
	if b == max {
		hsl.H = 4 + (r-g)*1/(max-min)
	}
	hsl.H = hsl.H * float32(int32(60))
	if hsl.H < float32(int32(0)) {
		hsl.H += float32(int32(360))
	}
	hsl.a = float32(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color_rgb._argb)) + 3)))) / 255
	return hsl
}
func PX_ColorHSLToRGB(color_hsl Struct__px_color_hsl) Struct__px_color {
	var r float32
	var g float32
	var b float32
	var temp1 float32
	var temp2 float32
	var temp3 float32
	var H float32
	var S float32
	var L float32
	H = color_hsl.H
	S = color_hsl.S
	L = color_hsl.L
	if color_hsl.S == float32(int32(0)) {
		r = func() (_cgo_ret float32) {
			_cgo_addr := &g
			*_cgo_addr = func() (_cgo_ret float32) {
				_cgo_addr := &b
				*_cgo_addr = color_hsl.L
				return *_cgo_addr
			}()
			return *_cgo_addr
		}()
	} else {
		if float64(L) < 0.5 {
			temp2 = L * (1 + S)
		} else {
			temp2 = L + S - L*S
		}
		temp1 = 2*L - temp2
		H /= float32(int32(360))
		r = func() (_cgo_ret float32) {
			_cgo_addr := &temp3
			*_cgo_addr = H + 1/3
			return *_cgo_addr
		}()
		if temp3 < float32(int32(0)) {
			temp3 = temp3 + 1
		}
		if temp3 > float32(int32(1)) {
			temp3 = temp3 - 1
		}
		if 6*temp3 < float32(int32(1)) {
			r = temp1 + (temp2-temp1)*6*temp3
		} else if 2*temp3 < float32(int32(1)) {
			r = temp2
		} else if 3*temp3 < float32(int32(2)) {
			r = temp1 + (temp2-temp1)*(2/3-temp3)*6
		} else {
			r = temp1
		}
		temp3 = H
		if temp3 < float32(int32(0)) {
			temp3 = temp3 + 1
		}
		if temp3 > float32(int32(1)) {
			temp3 = temp3 - 1
		}
		if 6*temp3 < float32(int32(1)) {
			g = temp1 + (temp2-temp1)*6*temp3
		} else if 2*temp3 < float32(int32(1)) {
			g = temp2
		} else if 3*temp3 < float32(int32(2)) {
			g = temp1 + (temp2-temp1)*(2/3-temp3)*6
		} else {
			g = temp1
		}
		temp3 = H - 1/3
		if temp3 < float32(int32(0)) {
			temp3 = temp3 + 1
		}
		if temp3 > float32(int32(1)) {
			temp3 = temp3 - 1
		}
		if 6*temp3 < float32(int32(1)) {
			b = temp1 + (temp2-temp1)*6*temp3
		} else if 2*temp3 < float32(int32(1)) {
			b = temp2
		} else if 3*temp3 < float32(int32(2)) {
			b = temp1 + (temp2-temp1)*(2/3-temp3)*6
		} else {
			b = temp1
		}
	}
	return PX_COLOR(uint8(color_hsl.a*float32(int32(255))), uint8(r*float32(int32(255))), uint8(g*float32(int32(255))), uint8(b*float32(int32(255))))
}
func PX_ColorMul(color1 Struct__px_color, muls float64) Struct__px_color {
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)) = uint8(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 3)))) * muls)
	*(*uint8)(unsafe.Pointer(&color1._argb)) = uint8(float64(int32(*(*uint8)(unsafe.Pointer(&color1._argb)))) * muls)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)) = uint8(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 1)))) * muls)
	*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)) = uint8(float64(int32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&color1._argb)) + 2)))) * muls)
	return color1
}
func PX_PointAdd(p1 Struct__px_point, p2 Struct__px_point) Struct__px_point {
	p1.x += p2.x
	p1.y += p2.y
	p1.z += p2.z
	return p1
}
func PX_Point2DAdd(p1 _cgoa_6_PX_Typedef, p2 _cgoa_6_PX_Typedef) _cgoa_6_PX_Typedef {
	p1.x += p2.x
	p1.y += p2.y
	return p1
}
func PX_PointSub(p1 Struct__px_point, p2 Struct__px_point) Struct__px_point {
	p1.x -= p2.x
	p1.y -= p2.y
	p1.z -= p2.z
	return p1
}
func PX_Point2DSub(p1 _cgoa_6_PX_Typedef, p2 _cgoa_6_PX_Typedef) _cgoa_6_PX_Typedef {
	p1.x -= p2.x
	p1.y -= p2.y
	return p1
}
func PX_Point4DSub(p1 Struct__px_point4, p2 Struct__px_point4) Struct__px_point4 {
	var v Struct__px_point4
	v.x = p1.x - p2.x
	v.y = p1.y - p2.y
	v.z = p1.z - p2.z
	v.w = float32(int32(1))
	return v
}
func PX_PointMul(p1 Struct__px_point, m float32) Struct__px_point {
	p1.x *= m
	p1.y *= m
	p1.z *= m
	return p1
}
func PX_Point2DMul(p1 _cgoa_6_PX_Typedef, m float32) _cgoa_6_PX_Typedef {
	p1.x *= m
	p1.y *= m
	return p1
}
func PX_PointDiv(p1 Struct__px_point, m float32) Struct__px_point {
	p1.x /= m
	p1.y /= m
	p1.z /= m
	return p1
}
func PX_Point2DRrthonormal(v _cgoa_6_PX_Typedef) _cgoa_6_PX_Typedef {
	return PX_Point2DNormalization(PX_POINT2D(v.y, -v.x))
}
func PX_Point2DBase(base1 _cgoa_6_PX_Typedef, base2 _cgoa_6_PX_Typedef, target _cgoa_6_PX_Typedef) _cgoa_6_PX_Typedef {
	base1 = PX_Point2DNormalization(base1)
	base2 = PX_Point2DNormalization(base2)
	return PX_POINT2D((target.x*base2.y-base2.x*target.y)/(base1.x*base2.y-base2.x*base1.y), (target.x*base1.y-base1.x*target.y)/(base2.x*base1.y-base1.x*base2.y))
}
func PX_Point2DDiv(p1 _cgoa_6_PX_Typedef, m float32) _cgoa_6_PX_Typedef {
	p1.x /= m
	p1.y /= m
	return p1
}
func PX_PointDot(p1 Struct__px_point, p2 Struct__px_point) float32 {
	return p1.x*p2.x + p1.y*p2.y + p1.z*p2.z
}
func PX_Point2DDot(p1 _cgoa_6_PX_Typedef, p2 _cgoa_6_PX_Typedef) float32 {
	return p1.x*p2.x + p1.y*p2.y
}
func PX_Point4DDot(p1 Struct__px_point4, p2 Struct__px_point4) float32 {
	return p1.x*p2.x + p1.y*p2.y + p1.z*p2.z
}
func PX_PointCross(p1 Struct__px_point, p2 Struct__px_point) Struct__px_point {
	var pt Struct__px_point
	pt.x = p1.y*p2.z - p2.y*p1.z
	pt.y = p1.z*p2.x - p2.z*p1.x
	pt.z = p1.x*p2.y - p2.x*p1.y
	return pt
}
func PX_Point4DCross(p1 Struct__px_point4, p2 Struct__px_point4) Struct__px_point4 {
	var pt Struct__px_point4
	pt.x = p1.y*p2.z - p2.y*p1.z
	pt.y = p1.z*p2.x - p2.z*p1.x
	pt.z = p1.x*p2.y - p2.x*p1.y
	pt.w = float32(int32(1))
	return pt
}
func PX_PointInverse(p1 Struct__px_point) Struct__px_point {
	return PX_POINT(-p1.x, -p1.y, -p1.z)
}
func PX_PointMod(p Struct__px_point) float32 {
	return PX_sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}
func PX_Point2DMod(p _cgoa_6_PX_Typedef) float32 {
	return PX_sqrt(p.x*p.x + p.y*p.y)
}
func PX_PointSquare(p Struct__px_point) float32 {
	return p.x*p.x + p.y*p.y + p.z*p.z
}
func PX_PointNormalization(p Struct__px_point) Struct__px_point {
	if p.x != 0 || p.y != 0 || p.z != 0 {
		return PX_PointDiv(p, PX_PointMod(p))
	}
	return p
}
func PX_Point2DNormalization(p _cgoa_6_PX_Typedef) _cgoa_6_PX_Typedef {
	if p.x != 0 || p.y != 0 {
		return PX_Point2DDiv(p, PX_Point2DMod(p))
	}
	return p
}
func PX_Point4DUnit(p Struct__px_point4) Struct__px_point4 {
	var pt Struct__px_point4
	if p.x != 0 || p.y != 0 || p.z != 0 {
		var sqare float32 = PX_sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
		pt.x = p.x / sqare
		pt.y = p.y / sqare
		pt.z = p.z / sqare
		pt.w = float32(int32(1))
		return pt
	}
	return p
}
func PX_PointReflectX(vector_refer Struct__px_point, respoint Struct__px_point) Struct__px_point {
	var ret Struct__px_point
	var cosx float32
	var sinx float32
	var mod float32 = PX_PointMod(vector_refer)
	cosx = vector_refer.x / mod
	sinx = -vector_refer.y / mod
	ret.x = respoint.x*cosx + respoint.y*sinx
	ret.y = respoint.y*cosx - respoint.x*sinx
	ret.z = respoint.z
	return ret
}
func PX_isRectCrossRect(rect1 Struct__px_rect, rect2 Struct__px_rect) int32 {
	var disx float32
	var disy float32
	disx = func() float32 {
		if rect1.x+rect1.width/float32(int32(2))-rect2.x-rect2.width/float32(int32(2)) > float32(int32(0)) {
			return rect1.x + rect1.width/float32(int32(2)) - rect2.x - rect2.width/float32(int32(2))
		} else {
			return -(rect1.x + rect1.width/float32(int32(2)) - rect2.x - rect2.width/float32(int32(2)))
		}
	}()
	disy = func() float32 {
		if rect1.y+rect1.height/float32(int32(2))-rect2.y-rect2.height/float32(int32(2)) > float32(int32(0)) {
			return rect1.y + rect1.height/float32(int32(2)) - rect2.y - rect2.height/float32(int32(2))
		} else {
			return -(rect1.y + rect1.height/float32(int32(2)) - rect2.y - rect2.height/float32(int32(2)))
		}
	}()
	if disx*float32(int32(2)) < rect1.width+rect2.width && disy*float32(int32(2)) < rect1.height+rect2.height {
		return int32(1)
	}
	return int32(0)
}
func PX_isRectCrossCircle(rect1 Struct__px_rect, center Struct__px_point, radius float32) int32 {
	var dx float32 = func() float32 {
		if rect1.x+rect1.width/float32(int32(2))-center.x > float32(int32(0)) {
			return rect1.x + rect1.width/float32(int32(2)) - center.x
		} else {
			return -(rect1.x + rect1.width/float32(int32(2)) - center.x)
		}
	}()
	var dy float32 = func() float32 {
		if rect1.y+rect1.height/float32(int32(2))-center.y > float32(int32(0)) {
			return rect1.y + rect1.height/float32(int32(2)) - center.y
		} else {
			return -(rect1.y + rect1.height/float32(int32(2)) - center.y)
		}
	}()
	return func() int32 {
		if dx < radius+rect1.width/float32(int32(2)) && dy < radius+rect1.height/float32(int32(2)) {
			return 1
		} else {
			return 0
		}
	}()
}
func PX_isCircleCrossCircle(center1 Struct__px_point, radius1 float32, center2 Struct__px_point, radius2 float32) int32 {
	var dis float32 = PX_PointDistance(center1, center2)
	return func() int32 {
		if dis < radius1+radius2 {
			return 1
		} else {
			return 0
		}
	}()
}
func PX_Covariance(x *float64, y *float64, n int32) float64 {
	var i int32
	var x_average float64
	var y_average float64
	var sum float64
	sum = float64(int32(0))
	if n-int32(1) <= int32(0) {
		return float64(int32(0))
	}
	for i = int32(0); i < n; i++ {
		sum += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8))
	}
	x_average = sum / float64(n)
	sum = float64(int32(0))
	for i = int32(0); i < n; i++ {
		sum += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(i)*8))
	}
	y_average = sum / float64(n)
	sum = float64(int32(0))
	for i = int32(0); i < n; i++ {
		sum += (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) - x_average) * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(i)*8)) - y_average)
	}
	return float64(sum) / float64(n-int32(1))
}
func PX_Variance(x *float64, n int32) float64 {
	var i int32
	var average float64
	var sum float64
	if n-int32(1) <= int32(0) {
		return float64(int32(0))
	}
	sum = float64(int32(0))
	for i = int32(0); i < n; i++ {
		sum += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8))
	}
	average = sum / float64(n)
	for i = int32(0); i < n; i++ {
		sum += (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) - average) * (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8)) - average)
	}
	return float64(sum) / float64(n-int32(1))
}
func PX_isPointInRect(p Struct__px_point, rect Struct__px_rect) int32 {
	if p.x < rect.x {
		return int32(0)
	}
	if p.x > rect.x+rect.width {
		return int32(0)
	}
	if p.y < rect.y {
		return int32(0)
	}
	if p.y > rect.y+rect.height {
		return int32(0)
	}
	return int32(1)
}
func PX_isXYInRegion(x float32, y float32, rectx float32, recty float32, width float32, height float32) int32 {
	if x < rectx {
		return int32(0)
	}
	if x > rectx+width {
		return int32(0)
	}
	if y < recty {
		return int32(0)
	}
	if y > recty+height {
		return int32(0)
	}
	return int32(1)
}
func PX_isLineCrossRect(p1 Struct__px_point, p2 Struct__px_point, rect Struct__px_rect, cp1 *Struct__px_point, cp2 *Struct__px_point) int32 {
	var k float32
	var b float32
	var calx float32
	var caly float32
	var bcross int32 = int32(0)
	if !(cp1 != nil) || !(cp2 != nil) {
		return int32(0)
	}
	if cp1 != nil && cp2 != nil {
		*cp1 = p1
		*cp2 = p2
	}
	if PX_isPointInRect(p1, rect) != 0 && PX_isPointInRect(p2, rect) != 0 {
		return int32(1)
	}
	if func() float32 {
		if p1.x-p2.x > float32(int32(0)) {
			return p1.x - p2.x
		} else {
			return -(p1.x - p2.x)
		}
	}() < 9.99999997e-7 {
		if p2.y < p1.y {
			var p Struct__px_point
			p = p2
			p2 = p1
			p1 = p
		}
		*cp1 = p1
		*cp2 = p2
		if p1.x > rect.x && p1.x < rect.x+rect.width {
			if p1.y > rect.y+rect.height {
				return int32(0)
			}
			if p2.y < rect.y {
				return int32(0)
			}
			if p1.y < rect.y {
				cp1.y = rect.y
			}
			if p2.y > rect.y+rect.height {
				cp2.y = rect.y + rect.height
			}
			return int32(1)
		}
		return int32(0)
	}
	if func() float32 {
		if p1.y-p2.y > float32(int32(0)) {
			return p1.y - p2.y
		} else {
			return -(p1.y - p2.y)
		}
	}() < 9.99999997e-7 {
		if p2.x < p1.x {
			var p Struct__px_point
			p = p2
			p2 = p1
			p1 = p
		}
		*cp1 = p1
		*cp2 = p2
		if p1.y > rect.y && p1.y < rect.y+rect.height {
			if p1.x > rect.x+rect.width {
				return int32(0)
			}
			if p2.x < rect.x {
				return int32(0)
			}
			if p1.x < rect.x {
				cp1.x = rect.x
			}
			if p2.x > rect.x+rect.width {
				cp2.x = rect.x + rect.width
			}
			return int32(1)
		}
		return int32(0)
	}
	if p2.x < p1.x {
		var p Struct__px_point
		p = p2
		p2 = p1
		p1 = p
	}
	k = (p2.y - p1.y) / (p2.x - p1.x)
	b = p2.y - k*p2.x
	calx = rect.x
	caly = calx*k + b
	if PX_isPointInRect(PX_POINT(calx, caly, float32(int32(0))), rect) != 0 {
		if calx > p1.x && calx < p2.x {
			cp1.x = calx
			cp1.y = caly
			bcross = int32(1)
		}
	}
	calx = rect.x + rect.width
	caly = calx*k + b
	if PX_isPointInRect(PX_POINT(calx, caly, float32(int32(0))), rect) != 0 {
		if calx > p1.x && calx < p2.x {
			if cp2 != nil {
				cp2.x = calx
				cp2.y = caly
			}
			bcross = int32(1)
		}
	}
	caly = rect.y
	calx = (caly - b) / k
	if PX_isPointInRect(PX_POINT(calx, caly, float32(int32(0))), rect) != 0 {
		if calx > p1.x && calx < p2.x {
			if k > float32(int32(0)) {
				if cp1 != nil {
					cp1.x = calx
					cp1.y = caly
				}
			} else if cp2 != nil {
				cp2.x = calx
				cp2.y = caly
			}
			bcross = int32(1)
		}
	}
	caly = rect.y + rect.height
	calx = (caly - b) / k
	if PX_isPointInRect(PX_POINT(calx, caly, float32(int32(0))), rect) != 0 {
		if calx > p1.x && calx < p2.x {
			if k > float32(int32(0)) {
				if cp2 != nil {
					cp2.x = calx
					cp2.y = caly
				}
			} else if cp1 != nil {
				cp1.x = calx
				cp1.y = caly
			}
			bcross = int32(1)
		}
	}
	return bcross
}
func PX_memset(dst unsafe.Pointer, byte uint8, size int32) {
	var dw uint32 = uint32(func() int32 {
		if int32(byte) != 0 {
			return int32(byte)<<int32(24) | int32(byte)<<int32(16) | int32(byte)<<int32(8) | int32(byte)
		} else {
			return int32(0)
		}
	}())
	var _4byteMovDst *uint32 = (*uint32)(dst)
	var pdst *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst))) + uintptr(size&-4)))
	var _movTs uint32 = uint32(size >> int32(2))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovDst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = dw
	}
	_movTs = uint32(size & int32(3))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint8) {
			_cgo_addr := &pdst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = byte
	}
}
func PX_memdwordset(dst unsafe.Pointer, dw uint32, count int32) {
	var p *uint32 = (*uint32)(dst)
	for func() (_cgo_ret int32) {
		_cgo_addr := &count
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint32) {
			_cgo_addr := &p
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() = dw
	}
}
func PX_memequ(dst unsafe.Pointer, src unsafe.Pointer, size int32) int32 {
	var _4byteMovSrc *uint32 = (*uint32)(src)
	var _4byteMovDst *uint32 = (*uint32)(dst)
	var psrc *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(src))) + uintptr(size&-4)))
	var pdst *uint8 = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst))) + uintptr(size&-4)))
	var _movTs uint32 = uint32(size >> int32(2))
	if uintptr(unsafe.Pointer(dst)) == uintptr(unsafe.Pointer(nil)) || uintptr(unsafe.Pointer(src)) == uintptr(unsafe.Pointer(nil)) {
		PX_ASSERT()
		return int32(0)
	}
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovDst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() != *func() (_cgo_ret *uint32) {
			_cgo_addr := &_4byteMovSrc
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
			return
		}() {
			return int32(0)
		}
	}
	_movTs = uint32(size & int32(3))
	for func() (_cgo_ret uint32) {
		_cgo_addr := &_movTs
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &pdst
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()) != int32(*func() (_cgo_ret *uint8) {
			_cgo_addr := &psrc
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()) {
			return int32(0)
		}
	}
	return int32(1)
}
func PX_memcpy(dst unsafe.Pointer, src unsafe.Pointer, size int32) {
	type _cgoa_15_PX_Typedef struct {
		m [16]uint8
	}
	type PX_MEMCPY_16 = _cgoa_15_PX_Typedef
	type _cgoa_16_PX_Typedef struct {
		m [32]uint8
	}
	type PX_MEMCPY_32 = _cgoa_16_PX_Typedef
	type _cgoa_17_PX_Typedef struct {
		m [64]uint8
	}
	type PX_MEMCPY_64 = _cgoa_17_PX_Typedef
	type _cgoa_18_PX_Typedef struct {
		m [128]uint8
	}
	type PX_MEMCPY_128 = _cgoa_18_PX_Typedef
	type _cgoa_19_PX_Typedef struct {
		m [256]uint8
	}
	type PX_MEMCPY_256 = _cgoa_19_PX_Typedef
	type _cgoa_20_PX_Typedef struct {
		m [512]uint8
	}
	type PX_MEMCPY_512 = _cgoa_20_PX_Typedef
	type _cgoa_21_PX_Typedef struct {
		m [1024]uint8
	}
	type PX_MEMCPY_1024 = _cgoa_21_PX_Typedef
	type _cgoa_22_PX_Typedef struct {
		m [2048]uint8
	}
	type PX_MEMCPY_2048 = _cgoa_22_PX_Typedef
	type _cgoa_23_PX_Typedef struct {
		m [4096]uint8
	}
	type PX_MEMCPY_4096 = _cgoa_23_PX_Typedef
	var _4byteMovSrc *uint32
	var _4byteMovDst *uint32
	var psrc *uint8
	var pdst *uint8
	var _4kbyteMovSrc *_cgoa_23_PX_Typedef
	var _4kbyteMovDst *_cgoa_23_PX_Typedef
	var _2kbyteMovSrc *_cgoa_22_PX_Typedef
	var _2kbyteMovDst *_cgoa_22_PX_Typedef
	var _1kbyteMovSrc *_cgoa_21_PX_Typedef
	var _1kbyteMovDst *_cgoa_21_PX_Typedef
	var _512byteMovSrc *_cgoa_20_PX_Typedef
	var _512byteMovDst *_cgoa_20_PX_Typedef
	var _256byteMovSrc *_cgoa_19_PX_Typedef
	var _256byteMovDst *_cgoa_19_PX_Typedef
	var _128byteMovSrc *_cgoa_18_PX_Typedef
	var _128byteMovDst *_cgoa_18_PX_Typedef
	var _64byteMovSrc *_cgoa_17_PX_Typedef
	var _64byteMovDst *_cgoa_17_PX_Typedef
	var _32byteMovSrc *_cgoa_16_PX_Typedef
	var _32byteMovDst *_cgoa_16_PX_Typedef
	var _16byteMovSrc *_cgoa_15_PX_Typedef
	var _16byteMovDst *_cgoa_15_PX_Typedef
	var _movTs uint32
	if size <= int32(0) {
		return
	}
	if uintptr(unsafe.Pointer(dst)) > uintptr(unsafe.Pointer(unsafe.Pointer(src))) && uintptr(unsafe.Pointer((*int8)(dst))) < uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(src)))+uintptr(size))))) {
		var _4byteMov uint32 = uint32(0)
		psrc = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(src)))+uintptr(size))))) - uintptr(int32(1))))
		pdst = (*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(dst)))+uintptr(size))))) - uintptr(int32(1))))
		_movTs = uint32(size & int32(3))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &pdst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &psrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))--
				return
			}()
		}
		*(*uintptr)(unsafe.Pointer(&pdst)) -= uintptr(int32(3))
		*(*uintptr)(unsafe.Pointer(&psrc)) -= uintptr(int32(3))
		_4byteMovDst = (*uint32)(unsafe.Pointer(pdst))
		_4byteMovSrc = (*uint32)(unsafe.Pointer(psrc))
		_movTs = uint32(size >> int32(2))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			_4byteMov = *_4byteMovSrc
			*(*uintptr)(unsafe.Pointer(&_4byteMovSrc)) -= 4
			*_4byteMovDst = _4byteMov
			*(*uintptr)(unsafe.Pointer(&_4byteMovDst)) -= 4
		}
	} else if uintptr(unsafe.Pointer(src)) > uintptr(unsafe.Pointer(unsafe.Pointer(dst))) && uintptr(unsafe.Pointer((*int8)(src))) < uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(dst)))+uintptr(size))))) {
		var _4byteMov uint32 = uint32(0)
		psrc = (*uint8)(src)
		pdst = (*uint8)(dst)
		_4byteMovDst = (*uint32)(unsafe.Pointer(pdst))
		_4byteMovSrc = (*uint32)(unsafe.Pointer(psrc))
		_movTs = uint32(size >> int32(2))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			_4byteMov = *_4byteMovSrc
			*(*uintptr)(unsafe.Pointer(&_4byteMovSrc)) += 4
			*_4byteMovDst = _4byteMov
			*(*uintptr)(unsafe.Pointer(&_4byteMovDst)) += 4
		}
		_movTs = uint32(size & int32(3))
		psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
		pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
		for func() (_cgo_ret uint32) {
			_cgo_addr := &_movTs
			_cgo_ret = *_cgo_addr
			*_cgo_addr--
			return
		}() != 0 {
			*func() (_cgo_ret *uint8) {
				_cgo_addr := &pdst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *uint8) {
				_cgo_addr := &psrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		}
		*(*uintptr)(unsafe.Pointer(&pdst)) += uintptr(int32(3))
		*(*uintptr)(unsafe.Pointer(&psrc)) += uintptr(int32(3))
	} else {
		_movTs = uint32(size >> int32(12))
		if _movTs != 0 {
			_4kbyteMovSrc = (*_cgoa_23_PX_Typedef)(src)
			_4kbyteMovDst = (*_cgoa_23_PX_Typedef)(dst)
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *_cgoa_23_PX_Typedef) {
					_cgo_addr := &_4kbyteMovDst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4096
					return
				}() = *func() (_cgo_ret *_cgoa_23_PX_Typedef) {
					_cgo_addr := &_4kbyteMovSrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4096
					return
				}()
			}
			src = unsafe.Pointer(_4kbyteMovSrc)
			dst = unsafe.Pointer(_4kbyteMovDst)
		}
		_movTs = uint32(size & 2048)
		if _movTs != 0 {
			_2kbyteMovSrc = (*_cgoa_22_PX_Typedef)(src)
			_2kbyteMovDst = (*_cgoa_22_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_22_PX_Typedef) {
				_cgo_addr := &_2kbyteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2048
				return
			}() = *func() (_cgo_ret *_cgoa_22_PX_Typedef) {
				_cgo_addr := &_2kbyteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2048
				return
			}()
			src = unsafe.Pointer(_2kbyteMovSrc)
			dst = unsafe.Pointer(_2kbyteMovDst)
		}
		_movTs = uint32(size & 1024)
		if _movTs != 0 {
			_1kbyteMovSrc = (*_cgoa_21_PX_Typedef)(src)
			_1kbyteMovDst = (*_cgoa_21_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_21_PX_Typedef) {
				_cgo_addr := &_1kbyteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 1024
				return
			}() = *func() (_cgo_ret *_cgoa_21_PX_Typedef) {
				_cgo_addr := &_1kbyteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 1024
				return
			}()
			src = unsafe.Pointer(_1kbyteMovSrc)
			dst = unsafe.Pointer(_1kbyteMovDst)
		}
		_movTs = uint32(size & 512)
		if _movTs != 0 {
			_512byteMovSrc = (*_cgoa_20_PX_Typedef)(src)
			_512byteMovDst = (*_cgoa_20_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_20_PX_Typedef) {
				_cgo_addr := &_512byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 512
				return
			}() = *func() (_cgo_ret *_cgoa_20_PX_Typedef) {
				_cgo_addr := &_512byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 512
				return
			}()
			src = unsafe.Pointer(_512byteMovSrc)
			dst = unsafe.Pointer(_512byteMovDst)
		}
		_movTs = uint32(size & 256)
		if _movTs != 0 {
			_256byteMovSrc = (*_cgoa_19_PX_Typedef)(src)
			_256byteMovDst = (*_cgoa_19_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_19_PX_Typedef) {
				_cgo_addr := &_256byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 256
				return
			}() = *func() (_cgo_ret *_cgoa_19_PX_Typedef) {
				_cgo_addr := &_256byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 256
				return
			}()
			src = unsafe.Pointer(_256byteMovSrc)
			dst = unsafe.Pointer(_256byteMovDst)
		}
		_movTs = uint32(size & 128)
		if _movTs != 0 {
			_128byteMovSrc = (*_cgoa_18_PX_Typedef)(src)
			_128byteMovDst = (*_cgoa_18_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_18_PX_Typedef) {
				_cgo_addr := &_128byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 128
				return
			}() = *func() (_cgo_ret *_cgoa_18_PX_Typedef) {
				_cgo_addr := &_128byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 128
				return
			}()
			src = unsafe.Pointer(_128byteMovSrc)
			dst = unsafe.Pointer(_128byteMovDst)
		}
		_movTs = uint32(size & 64)
		if _movTs != 0 {
			_64byteMovSrc = (*_cgoa_17_PX_Typedef)(src)
			_64byteMovDst = (*_cgoa_17_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_17_PX_Typedef) {
				_cgo_addr := &_64byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 64
				return
			}() = *func() (_cgo_ret *_cgoa_17_PX_Typedef) {
				_cgo_addr := &_64byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 64
				return
			}()
			src = unsafe.Pointer(_64byteMovSrc)
			dst = unsafe.Pointer(_64byteMovDst)
		}
		_movTs = uint32(size & 32)
		if _movTs != 0 {
			_32byteMovSrc = (*_cgoa_16_PX_Typedef)(src)
			_32byteMovDst = (*_cgoa_16_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_16_PX_Typedef) {
				_cgo_addr := &_32byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 32
				return
			}() = *func() (_cgo_ret *_cgoa_16_PX_Typedef) {
				_cgo_addr := &_32byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 32
				return
			}()
			src = unsafe.Pointer(_32byteMovSrc)
			dst = unsafe.Pointer(_32byteMovDst)
		}
		_movTs = uint32(size & 16)
		if _movTs != 0 {
			_16byteMovSrc = (*_cgoa_15_PX_Typedef)(src)
			_16byteMovDst = (*_cgoa_15_PX_Typedef)(dst)
			*func() (_cgo_ret *_cgoa_15_PX_Typedef) {
				_cgo_addr := &_16byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 16
				return
			}() = *func() (_cgo_ret *_cgoa_15_PX_Typedef) {
				_cgo_addr := &_16byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 16
				return
			}()
			src = unsafe.Pointer(_16byteMovSrc)
			dst = unsafe.Pointer(_16byteMovDst)
		}
		_movTs = uint32(size & int32(15))
		if _movTs >= uint32(12) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(12)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else if _movTs >= uint32(8) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(8)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else if _movTs >= uint32(4) {
			_4byteMovSrc = (*uint32)(src)
			_4byteMovDst = (*uint32)(dst)
			*func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovDst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}() = *func() (_cgo_ret *uint32) {
				_cgo_addr := &_4byteMovSrc
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 4
				return
			}()
			_movTs -= uint32(4)
			psrc = (*uint8)(unsafe.Pointer(_4byteMovSrc))
			pdst = (*uint8)(unsafe.Pointer(_4byteMovDst))
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		} else {
			psrc = (*uint8)(src)
			pdst = (*uint8)(dst)
			for func() (_cgo_ret uint32) {
				_cgo_addr := &_movTs
				_cgo_ret = *_cgo_addr
				*_cgo_addr--
				return
			}() != 0 {
				*func() (_cgo_ret *uint8) {
					_cgo_addr := &pdst
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}() = *func() (_cgo_ret *uint8) {
					_cgo_addr := &psrc
					_cgo_ret = *_cgo_addr
					*(*uintptr)(unsafe.Pointer(_cgo_addr))++
					return
				}()
			}
		}
	}
}
func PX_strcpy(dst *int8, src *int8, size int32) {
	for func() (_cgo_ret int32) {
		_cgo_addr := &size
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *src != 0 {
			*func() (_cgo_ret *int8) {
				_cgo_addr := &dst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}() = *func() (_cgo_ret *int8) {
				_cgo_addr := &src
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr))++
				return
			}()
		} else {
			*dst = int8('\x00')
			return
		}
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) - uintptr(int32(1)))) = int8('\x00')
}
func PX_wstrcpy(dst *uint16, src *uint16, size int32) {
	for func() (_cgo_ret int32) {
		_cgo_addr := &size
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		if *src != 0 {
			*func() (_cgo_ret *uint16) {
				_cgo_addr := &dst
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
				return
			}() = *func() (_cgo_ret *uint16) {
				_cgo_addr := &src
				_cgo_ret = *_cgo_addr
				*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
				return
			}()
		} else {
			*dst = uint16('\x00')
			return
		}
	}
	*(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) - uintptr(int32(1))*2)) = uint16('\x00')
}
func PX_strcat(src *int8, cat *int8) {
	var len int32 = PX_strlen(cat)
	for *src != 0 {
		*(*uintptr)(unsafe.Pointer(&src))++
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *int8) {
			_cgo_addr := &src
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}() = *func() (_cgo_ret *int8) {
			_cgo_addr := &cat
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}
	*src = int8('\x00')
}
func PX_strcat_s(src *int8, size int32, cat *int8) {
	if PX_strlen(src)+PX_strlen(cat) < size {
		PX_strcat(src, cat)
	}
}
func PX_wstrcat(src *uint16, cat *uint16) {
	var len int32 = PX_wstrlen(cat)
	for *src != 0 {
		*(*uintptr)(unsafe.Pointer(&src)) += 2
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		*func() (_cgo_ret *uint16) {
			_cgo_addr := &src
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
			return
		}() = *func() (_cgo_ret *uint16) {
			_cgo_addr := &cat
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr)) += 2
			return
		}()
	}
	*src = uint16('\x00')
}
func PX_strset(dst *int8, src *int8) {
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(int32(0)))) = int8(0)
	PX_strcat(dst, src)
}
func PX_strlen(dst *int8) int32 {
	var len int32 = int32(0)
	if !(dst != nil) {
		PX_ERROR((*int8)(unsafe.Pointer(&[21]int8{'N', 'U', 'L', 'L', ' ', 'p', 'o', 'i', 'n', 't', 'e', 'r', ' ', 'a', 's', 's', 'e', 'r', 't', '!', '\x00'})))
		return int32(0)
	}
	for *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}()))) != 0 {
	}
	return len - int32(1)
}
func PX_wstrlen(dst *uint16) int32 {
	var len int32 = int32(0)
	if !(dst != nil) {
		PX_ERROR((*int8)(unsafe.Pointer(&[21]int8{'N', 'U', 'L', 'L', ' ', 'p', 'o', 'i', 'n', 't', 'e', 'r', ' ', 'a', 's', 's', 'e', 'r', 't', '!', '\x00'})))
		return int32(0)
	}
	for *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(dst)) + uintptr(func() (_cgo_ret int32) {
		_cgo_addr := &len
		_cgo_ret = *_cgo_addr
		*_cgo_addr++
		return
	}())*2)) != 0 {
	}
	return len - int32(1)
}
func PX_strequ2(src *int8, dst *int8) int32 {
	var _l int8
	var _r int8
	for int32(1) != 0 {
		_l = *src
		_r = *dst
		if int32(_l) >= 'a' && int32(_l) <= 'z' {
			_l += int8(-32)
		}
		if int32(_r) >= 'a' && int32(_r) <= 'z' {
			_r += int8(-32)
		}
		if int32(_l) == int32(_r) {
			if int32(_l) == int32(0) {
				return int32(1)
			}
		} else {
			return int32(0)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
		*(*uintptr)(unsafe.Pointer(&dst))++
	}
	return int32(0)
}
func PX_strupr(src *int8) {
	for int32(*src) != '\x00' {
		if int32(*src) >= 'a' && int32(*src) <= 'z' {
			*src -= int8(32)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
	}
}
func PX_strlwr(src *int8) {
	for int32(*src) != '\x00' {
		if int32(*src) > 'A' && int32(*src) <= 'Z' {
			*src += int8(32)
		}
		*(*uintptr)(unsafe.Pointer(&src))++
	}
}
func PX_Point2DMulMatrix(p _cgoa_6_PX_Typedef, mat Struct__px_matrix) _cgoa_6_PX_Typedef {
	var point _cgoa_6_PX_Typedef
	point.x = p.x**(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48))
	point.y = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52))
	return point
}
func PX_PointMulMatrix(p Struct__px_point, mat Struct__px_matrix) Struct__px_point {
	var point Struct__px_point
	point.x = p.x**(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48))
	point.y = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52))
	point.z = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56))
	return point
}
func PX_Point4DMulMatrix(p Struct__px_point4, mat Struct__px_matrix) Struct__px_point4 {
	var point Struct__px_point4
	point.x = p.x**(*float32)(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 16)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 32)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 48))
	point.y = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 4)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 20)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 36)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 52))
	point.z = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 8)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 24)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 40)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 56))
	point.w = p.x**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 12)) + p.y**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 28)) + p.z**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 44)) + float32(int32(1))**(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(&mat._cgoa_2_PX_Typedef)) + 60))
	return point
}
func PX_PLANE(p Struct__px_point, n Struct__px_point) _cgoa_7_PX_Typedef {
	var pl _cgoa_7_PX_Typedef
	pl.n = n
	pl.p0 = p
	return pl
}
func PX_RECT(x float32, y float32, width float32, height float32) Struct__px_rect {
	var rect Struct__px_rect
	rect.x = x
	rect.y = y
	rect.width = width
	rect.height = height
	if width < float32(int32(0)) {
		rect.width = -width
		rect.x += width
	}
	if height < float32(int32(0)) {
		rect.height = -height
		rect.y += height
	}
	return rect
}
func PX_RECTPOINT2(p1 Struct__px_point, p2 Struct__px_point) Struct__px_rect {
	var rect Struct__px_rect
	rect.width = func() float32 {
		if p2.x-p1.x > float32(int32(0)) {
			return p2.x - p1.x
		} else {
			return -(p2.x - p1.x)
		}
	}()
	rect.height = func() float32 {
		if p2.y-p1.y > float32(int32(0)) {
			return p2.y - p1.y
		} else {
			return -(p2.y - p1.y)
		}
	}()
	if p1.x < p2.x {
		rect.x = p1.x
	} else {
		rect.x = p2.x
	}
	if p1.y < p2.y {
		rect.y = p1.y
	} else {
		rect.y = p2.y
	}
	return rect
}
func PX_complexBuild(re float32, im float32) Struct___px_complex {
	var cx Struct___px_complex
	cx.re = float64(re)
	cx.im = float64(im)
	return cx
}
func PX_complexAdd(a Struct___px_complex, b Struct___px_complex) Struct___px_complex {
	var ret Struct___px_complex
	ret.re = a.re + b.re
	ret.im = a.im + b.im
	return ret
}
func PX_complexMult(a Struct___px_complex, b Struct___px_complex) Struct___px_complex {
	var ret Struct___px_complex
	ret.re = a.re*b.re - a.im*b.im
	ret.im = a.im*b.re + a.re*b.im
	return ret
}
func PX_complexMod(a Struct___px_complex) float64 {
	return PX_sqrtd(a.re*a.re + a.im*a.im)
}
func PX_complexLog(a Struct___px_complex) Struct___px_complex {
	var ret Struct___px_complex
	ret.re = PX_sqrtd(a.re*a.re + a.im*a.im)
	ret.im = PX_atan2(a.im, a.re)
	return ret
}
func PX_complexExp(a Struct___px_complex) Struct___px_complex {
	var p float64
	var ret Struct___px_complex
	p = PX_exp(a.re)
	ret.re = p * PX_cosd(a.im)
	ret.im = p * PX_cosd(a.re)
	return ret
}
func PX_complexSin(a Struct___px_complex) Struct___px_complex {
	var q float64
	var p float64
	var ret Struct___px_complex
	p = PX_exp(a.im)
	q = float64(int32(1)) / p
	ret.re = PX_sind(a.re) * (q + p) / 2
	ret.im = PX_cosd(a.re) * (p - q) / 2
	return ret
}
func PX_DFT(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	var k int32
	var n int32
	var Wnk Struct___px_complex
	if uintptr(unsafe.Pointer(x)) == uintptr(unsafe.Pointer(X)) {
		PX_ASSERT()
	}
	for k = int32(0); k < N; k++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(k)*16))).re = float64(int32(0))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(k)*16))).im = float64(int32(0))
		for n = int32(0); n < N; n++ {
			Wnk.re = float64(float32(PX_cosd(float64(int32(2)) * 3.1415926535897931 * float64(k) * float64(n) / float64(N))))
			Wnk.im = float64(float32(-PX_sind(float64(int32(2)) * 3.1415926535897931 * float64(k) * float64(n) / float64(N))))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(k)*16)) = PX_complexAdd(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(k)*16)), PX_complexMult(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(n)*16)), Wnk))
		}
	}
}
func PX_DCT(x *float64, X *float64, N int32) {
	var n int32
	var m int32
	var v float64
	if uintptr(unsafe.Pointer(x)) == uintptr(unsafe.Pointer(X)) {
		PX_ASSERT()
	}
	PX_memset(unsafe.Pointer(X), uint8(0), int32(8*uint64(N)))
	for n = int32(0); n < N; n++ {
		for m = int32(0); m < N; m++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*8)) += *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(m)*8)) * PX_cosd(float64(int32(2)*m+int32(1))*3.1415926535897931*float64(n)/float64(int32(2)*N))
		}
	}
	v = PX_sqrtd(2 / float64(N))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(0))*8)) *= PX_sqrtd(1 / float64(N))
	for n = int32(1); n < N; n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*8)) *= v
	}
}
func PX_IDFT(X *Struct___px_complex, x *Struct___px_complex, N int32) {
	var k int32
	var n int32
	var im float32 = float32(int32(0))
	var ejw Struct___px_complex
	if uintptr(unsafe.Pointer(x)) == uintptr(unsafe.Pointer(X)) {
		PX_ASSERT()
	}
	for k = int32(0); k < N; k++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16))).re = float64(int32(0))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16))).im = float64(int32(0))
		for n = int32(0); n < N; n++ {
			ejw.re = float64(float32(PX_cosd(float64(int32(2)) * 3.1415926535897931 * float64(k) * float64(n) / float64(N))))
			ejw.im = float64(float32(PX_sind(float64(int32(2)) * 3.1415926535897931 * float64(k) * float64(n) / float64(N))))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16)) = PX_complexAdd(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16)), PX_complexMult(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*16)), ejw))
		}
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16))).re /= float64(N)
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16))).im /= float64(N)
	}
	func() int {
		_ = im
		return 0
	}()
}
func PX_IDCT(X *float64, x *float64, N int32) {
	var n int32
	var m int32
	var v0 float64
	var v float64
	PX_memset(unsafe.Pointer(x), uint8(0), int32(8*uint64(N)))
	v = PX_sqrtd(2 / float64(N))
	v0 = PX_sqrtd(1 / float64(N))
	for m = int32(0); m < N; m++ {
		for n = int32(0); n < N; n++ {
			if n == int32(0) {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(m)*8)) += v0 * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*8)) * PX_cosd(3.1415926535897931*float64(n)*float64(int32(2)*m+int32(1))/float64(int32(2)*N))
			} else {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(m)*8)) += v * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*8)) * PX_cosd(3.1415926535897931*float64(n)*float64(int32(2)*m+int32(1))/float64(int32(2)*N))
			}
		}
	}
}
func _cgos_PX_FDCT_SortArray_PX_Typedef(x *Struct___px_complex, N int32, sN int32) {
	var i int32
	var j int32
	for i = int32(0); i < N/sN; i++ {
		if sN == int32(4) {
			var temp Struct___px_complex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*int32(4)+int32(1))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*int32(4)+int32(1))*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*int32(4)+int32(2))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*int32(4)+int32(2))*16)) = temp
		} else {
			for j = int32(0); j < sN/int32(4); j++ {
				var temp Struct___px_complex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*sN+sN/int32(4)+j)*16))
				*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*sN+sN/int32(4)+j)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*sN+sN/int32(2)+j)*16))
				*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*sN+sN/int32(2)+j)*16)) = temp
			}
		}
	}
}
func PX_FDCT(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	var i int32
	var sLen int32 = int32(4)
	var v float64
	PX_memcpy(unsafe.Pointer(X), unsafe.Pointer(x), int32(16*uint64(N)))
	if N&(N-int32(1)) != 0 {
		return
	}
	if N < int32(4) {
		return
	}
	for sLen <= N {
		_cgos_PX_FDCT_SortArray_PX_Typedef(X, N, sLen)
		sLen *= int32(2)
	}
	for i = int32(0); i < N/int32(4); i++ {
		var t Struct___px_complex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N/int32(2)+i)*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N/int32(2)+i)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N-i-int32(1))*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N-i-int32(1))*16)) = t
	}
	PX_FFT(X, X, N)
	for i = int32(0); i < N; i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re*PX_cosd(float64(i)*3.1415926535897931/float64(N*int32(2))) + (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im*PX_sind(float64(i)*3.1415926535897931/float64(N*int32(2)))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
	}
	v = PX_sqrtd(2 / float64(N))
	(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(0))*16))).re *= PX_sqrtd(1 / float64(N))
	for i = int32(1); i < N; i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re *= v
	}
}
func PX_FIDCT(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	var i int32
	var sLen int32 = N
	var v float64
	PX_memcpy(unsafe.Pointer(X), unsafe.Pointer(x), int32(16*uint64(N)))
	v = PX_sqrtd(2 / float64(N))
	for i = int32(1); i < N; i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re *= v
	}
	(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(0))*16))).re *= PX_sqrtd(1 / float64(N))
	for i = int32(0); i < N; i++ {
		var t float64 = (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = t * PX_cosd(float64(i)*3.1415926535897931/float64(N*int32(2)))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = -t * PX_sind(float64(i)*3.1415926535897931/float64(N*int32(2)))
	}
	PX_FFT(X, X, N)
	for i = int32(0); i < N/int32(4); i++ {
		var t Struct___px_complex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N/int32(2)+i)*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N/int32(2)+i)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N-i-int32(1))*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(N-i-int32(1))*16)) = t
	}
	for sLen >= int32(4) {
		_cgos_PX_FDCT_SortArray_PX_Typedef(X, N, sLen)
		sLen /= int32(2)
	}
	for i = int32(1); i < N; i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
	}
}
func FFT_Base2(x *Struct___px_complex, N int32) {
	var exbase int32
	var exrang int32
	var i int32
	var j int32
	var k int32
	var excomplex Struct___px_complex
	var Wnk Struct___px_complex
	var cx0 Struct___px_complex
	var cx1 Struct___px_complex
	if N>>int32(2) != 0 {
		exbase = int32(1)
		exrang = int32(0)
		for exrang < N {
			exrang = exbase << int32(2)
			for i = int32(0); i < N/exrang; i++ {
				for j = int32(0); j < exbase; j++ {
					excomplex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(exrang*i+exbase+j)*16))
					*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(exrang*i+exbase+j)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(exrang*i+exbase*int32(2)+j)*16))
					*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(exrang*i+exbase*int32(2)+j)*16)) = excomplex
				}
			}
			exbase <<= int32(1)
		}
		FFT_Base2(x, N>>int32(1))
		FFT_Base2((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x))+uintptr(N>>int32(1))*16)), N>>int32(1))
		for k = int32(0); k < N>>int32(1); k++ {
			Wnk.re = float64(float32(PX_cosd(float64(-2) * 3.1415926535897931 * float64(k) / float64(N))))
			Wnk.im = float64(float32(PX_sind(float64(-2) * 3.1415926535897931 * float64(k) / float64(N))))
			cx0 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16))
			cx1 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k+N>>int32(1))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*16)) = PX_complexAdd(cx0, PX_complexMult(Wnk, cx1))
			Wnk.re = -Wnk.re
			Wnk.im = -Wnk.im
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k+N>>int32(1))*16)) = PX_complexAdd(cx0, PX_complexMult(Wnk, cx1))
		}
	} else {
		cx0 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(int32(0))*16))
		cx1 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(int32(1))*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(int32(0))*16)) = PX_complexAdd(cx0, cx1)
		cx1.im = -cx1.im
		cx1.re = -cx1.re
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(int32(1))*16)) = PX_complexAdd(cx0, cx1)
	}
}
func PX_FFT(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	PX_memcpy(unsafe.Pointer(X), unsafe.Pointer(x), int32(16*uint64(N)))
	FFT_Base2(X, N)
}
func IFFT_Base2(X *Struct___px_complex, N int32) {
	var exbase int32
	var exrang int32
	var i int32
	var j int32
	var n int32
	var excomplex Struct___px_complex
	var Wnnk Struct___px_complex
	var cx0 Struct___px_complex
	var cx1 Struct___px_complex
	if N>>int32(2) != 0 {
		exbase = int32(1)
		exrang = int32(0)
		for exrang < N {
			exrang = exbase << int32(2)
			for i = int32(0); i < N/exrang; i++ {
				for j = int32(0); j < exbase; j++ {
					excomplex = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(exrang*i+exbase+j)*16))
					*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(exrang*i+exbase+j)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(exrang*i+exbase*int32(2)+j)*16))
					*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(exrang*i+exbase*int32(2)+j)*16)) = excomplex
				}
			}
			exbase <<= int32(1)
		}
		IFFT_Base2(X, N>>int32(1))
		IFFT_Base2((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N>>int32(1))*16)), N>>int32(1))
		for n = int32(0); n < N>>int32(1); n++ {
			Wnnk.re = float64(float32(PX_cosd(float64(int32(2)) * 3.1415926535897931 * float64(n) / float64(N))))
			Wnnk.im = float64(float32(PX_sind(float64(int32(2)) * 3.1415926535897931 * float64(n) / float64(N))))
			cx0 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*16))
			cx1 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n+N>>int32(1))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n)*16)) = PX_complexAdd(cx0, PX_complexMult(Wnnk, cx1))
			Wnnk.re = -Wnnk.re
			Wnnk.im = -Wnnk.im
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(n+N>>int32(1))*16)) = PX_complexAdd(cx0, PX_complexMult(Wnnk, cx1))
		}
	} else {
		cx0 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(0))*16))
		cx1 = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(1))*16))
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(0))*16)) = PX_complexAdd(cx0, cx1)
		cx1.im = -cx1.im
		cx1.re = -cx1.re
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(int32(1))*16)) = PX_complexAdd(cx0, cx1)
	}
}
func PX_IFFT(X *Struct___px_complex, x *Struct___px_complex, N int32) {
	var i int32
	PX_memcpy(unsafe.Pointer(X), unsafe.Pointer(x), int32(16*uint64(N)))
	IFFT_Base2(X, N)
	for i = int32(0); i < N; i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re /= float64(N)
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im /= float64(N)
	}
}
func PX_FFT_2(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	var i int32
	var cx int32
	var cy int32
	var _t Struct___px_complex
	for i = int32(0); i < N; i++ {
		PX_FFT(&*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*N)*16)), &*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), N)
	}
	for cy = int32(0); cy < N; cy++ {
		for cx = cy + int32(1); cx < N; cx++ {
			_t = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16)) = _t
		}
	}
	for i = int32(0); i < N; i++ {
		PX_FFT(&*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), &*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), N)
	}
	for cy = int32(0); cy < N; cy++ {
		for cx = cy + int32(1); cx < N; cx++ {
			_t = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16)) = _t
		}
	}
}
func PX_IFFT_2(X *Struct___px_complex, x *Struct___px_complex, N int32) {
	var cx int32
	var cy int32
	var i int32
	var _t Struct___px_complex
	for cy = int32(0); cy < N; cy++ {
		for cx = cy + int32(1); cx < N; cx++ {
			_t = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16)) = _t
		}
	}
	for i = int32(0); i < N; i++ {
		PX_IFFT(&*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i*N)*16)), &*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), N)
	}
	for cy = int32(0); cy < N; cy++ {
		for cx = cy + int32(1); cx < N; cx++ {
			_t = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cy*N+cx)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(cx*N+cy)*16)) = _t
		}
	}
	for i = int32(0); i < N; i++ {
		PX_IFFT(&*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), &*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*N)*16)), N)
	}
}
func PX_FFT_2_Shift(_in *Struct___px_complex, _out *Struct___px_complex, N int32) {
	var x int32
	var y int32
	var _t *Struct___px_complex = _out
	var _ext Struct___px_complex
	PX_memcpy(unsafe.Pointer(_t), unsafe.Pointer(_in), int32(16*uint64(N)*uint64(N)))
	for y = int32(0); y < N/int32(2); y++ {
		for x = int32(0); x < N/int32(2); x++ {
			_ext = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x+N/int32(2))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x+N/int32(2))*16)) = _ext
		}
	}
	for y = int32(0); y < N/int32(2); y++ {
		for x = N / int32(2); x < N; x++ {
			_ext = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x-N/int32(2))*16))
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x-N/int32(2))*16)) = _ext
		}
	}
}
func PX_DCT_2_Shift(_in *float64, _out *float64, N int32) {
	var x int32
	var y int32
	var _t *float64 = _out
	var _ext float64
	PX_memcpy(unsafe.Pointer(_t), unsafe.Pointer(_in), int32(8*uint64(N)*uint64(N)))
	for y = int32(0); y < N/int32(2); y++ {
		for x = int32(0); x < N/int32(2); x++ {
			_ext = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x+N/int32(2))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x+N/int32(2))*8)) = _ext
		}
	}
	for y = int32(0); y < N/int32(2); y++ {
		for x = N / int32(2); x < N; x++ {
			_ext = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+x)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x-N/int32(2))*8))
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(_t)) + uintptr(y*N+N*N/int32(2)+x-N/int32(2))*8)) = _ext
		}
	}
}
func PX_FT_Symmetry(x *Struct___px_complex, X *Struct___px_complex, N int32) {
	var l int32 = int32(1)
	var r int32 = N - int32(1)
	PX_memcpy(unsafe.Pointer(X), unsafe.Pointer(x), int32(uint64(N/int32(2))*16+uint64(1)))
	for l < r {
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(r)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(l)*16))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(r)*16))).im = -(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(r)*16))).im
		l++
		r--
	}
}
func PX_Cepstrum(x *Struct___px_complex, X *Struct___px_complex, N int32, type_ int32) {
	var i int32
	PX_FFT(x, X, N)
	if uint32(type_) == uint32(1) {
		for i = int32(0); i < N; i++ {
			*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16)) = PX_complexLog(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16)))
		}
		PX_IFFT(X, X, N)
	} else {
		for i = int32(0); i < N; i++ {
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = PX_log((*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re + (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im + float64(int32(1)))
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
		}
		PX_IFFT(X, X, N)
		for i = int32(0); i < N; i++ {
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re + (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
		}
	}
}
func PX_ZeroCrossingRate(x *float64, N int32) float64 {
	var i int32
	var zeroCross int32
	zeroCross = int32(0)
	for i = int32(1); i < N; i++ {
		if *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i-int32(1))*8)) < float64(int32(0)) {
			zeroCross++
		}
	}
	return float64(zeroCross) * 1 / float64(N)
}
func PX_ZeroCrossingRateComplex(x *Struct___px_complex, N int32) float64 {
	var i int32
	var zeroCross int32
	zeroCross = int32(0)
	for i = int32(1); i < N; i++ {
		if (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*16))).re*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i-int32(1))*16))).re < float64(int32(0)) {
			zeroCross++
		}
	}
	return float64(zeroCross) * 1 / float64(N)
}
func PX_PitchEstimation(x *Struct___px_complex, N int32, sampleRate int32, low_Hz int32, high_Hz int32) int32 {
	var low int32 = int32(0)
	var high int32 = int32(0)
	var i int32
	var idx int32 = int32(0)
	var max float64 = float64(int32(0))
	var zeroCross int32 = int32(0)
	low = sampleRate / high_Hz
	high = sampleRate / low_Hz
	for i = int32(1); i < N; i++ {
		if (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*16))).re*(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i-int32(1))*16))).re < float64(int32(0)) {
			zeroCross++
		}
	}
	if float64(float32(zeroCross)*1/float32(N)) > 0.14999999999999999 {
		return int32(0)
	}
	PX_Cepstrum(x, x, N, int32(0))
	if high >= N || high == low {
		return int32(0)
	}
	for i = low; i <= high; i++ {
		if (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*16))).re > max {
			max = (*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*16))).re
			idx = i
		}
	}
	if idx == int32(0) {
		return int32(0)
	}
	return sampleRate / idx
}
func PX_PreEmphasise(data *float64, len int32, out *float64, preF float64) {
	var i int32
	for i = len - int32(1); i >= int32(1); i-- {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) - preF**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i-int32(1))*8))
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(out)) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(int32(0))*8))
}
func PX_LinearInterpolationResample(x *float64, X *float64, N int32, M int32) {
	var k int32
	var m int32 = int32(0)
	var d1 float64 = float64(int32(0))
	var d2 float64 = float64(int32(0))
	var freqscale float64 = float64(N) * 1 / float64(M)
	for k = int32(0); k < M; k++ {
		d1 = float64(0)
		if m < N {
			d1 += (1 - d2) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(m)*8))
		}
		if m+int32(1) < N {
			d1 += d2 * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(m+int32(1))*8))
		}
		d2 += freqscale
		if d2 >= 1 {
			m += int32(d2)
			d2 -= float64(int32(d2))
		}
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(k)*8)) = d1
	}
}
func PX_SincInterpolationResample(x *float64, X *float64, N int32, M int32) {
	var step float64 = float64(int32(1) / (M - int32(1)))
	var i int32
	for i = int32(0); i < M; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*8)) = PX_sinc_interpolate(x, N, step*float64(i))
	}
}
func PX_UpSampled(x *Struct___px_complex, X *Struct___px_complex, N int32, L int32) {
	var i int32
	var j int32
	for i = int32(0); i < N; i++ {
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*L)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(i)*16))
		for j = int32(1); j < L; j++ {
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*L+j)*16))).re = float64(int32(0))
			(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*L+j)*16))).im = float64(int32(0))
		}
	}
	if N*L&(N*L-int32(1)) == int32(0) {
		PX_FFT(X, X, N*L)
	} else {
		PX_DFT(X, (*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N*L)*16)), N*L)
		PX_memcpy(unsafe.Pointer(X), unsafe.Pointer((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N*L)*16))), int32(16*uint64(N)*uint64(L)))
	}
	for i = N/int32(2) + int32(1); i < N*L/int32(2)+int32(1); i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = float64(int32(0))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
	}
	PX_FT_Symmetry(X, X, N*L)
	if N*L&(N*L-int32(1)) == int32(0) {
		PX_IFFT(X, X, N*L)
	} else {
		PX_IDFT(X, (*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N*L)*16)), N*L)
		PX_memcpy(unsafe.Pointer(X), unsafe.Pointer((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N*L)*16))), int32(16*uint64(N)*uint64(L)))
	}
}
func PX_DownSampled(x *Struct___px_complex, X *Struct___px_complex, N int32, M int32) {
	var i int32
	if N&(N-int32(1)) == int32(0) {
		PX_FFT(x, X, N)
	} else {
		PX_DFT(x, (*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N)*16)), N)
		PX_memcpy(unsafe.Pointer(X), unsafe.Pointer((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N)*16))), int32(16*uint64(N)))
	}
	for i = N/(int32(2)*M) + int32(1); i < N/int32(2)+int32(1); i++ {
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).re = float64(int32(0))
		(*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16))).im = float64(int32(0))
	}
	PX_FT_Symmetry(X, X, N)
	if N&(N-int32(1)) == int32(0) {
		PX_IFFT(X, X, N)
	} else {
		PX_IDFT(X, (*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N)*16)), N)
		PX_memcpy(unsafe.Pointer(X), unsafe.Pointer((*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X))+uintptr(N)*16))), int32(16*uint64(N)))
	}
	for i = int32(0); i < N/M; i++ {
		*(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i)*16)) = *(*Struct___px_complex)(unsafe.Pointer(uintptr(unsafe.Pointer(X)) + uintptr(i*M)*16))
	}
}
func PX_POINT(x float32, y float32, z float32) Struct__px_point {
	var p Struct__px_point
	p.x = x
	p.y = y
	p.z = z
	return p
}
func PX_POINT2D(x float32, y float32) _cgoa_6_PX_Typedef {
	var p _cgoa_6_PX_Typedef
	p.x = x
	p.y = y
	return p
}
func PX_POINT4D(x float32, y float32, z float32) Struct__px_point4 {
	var p Struct__px_point4
	p.x = x
	p.y = y
	p.z = z
	p.w = float32(int32(1))
	return p
}
func PX_strequ(src *int8, dst *int8) int32 {
	var ret int32 = int32(0)
	for !(func() (_cgo_ret int32) {
		_cgo_addr := &ret
		*_cgo_addr = int32(*(*uint8)(unsafe.Pointer(src))) - int32(*(*uint8)(unsafe.Pointer(dst)))
		return *_cgo_addr
	}() != 0) && int32(*src) != 0 {
		*(*uintptr)(unsafe.Pointer(&src))++
		*(*uintptr)(unsafe.Pointer(&dst))++
	}
	return func() int32 {
		if !(ret != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func PX_strIsNumeric(str *int8) int32 {
	var Dot int32 = int32(0)
	var CurrentCharIndex int32
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '\x00' {
		return int32(0)
	}
	if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != 0) && !(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-') {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != '.' {
			return int32(0)
		}
	}
	for CurrentCharIndex = int32(1); CurrentCharIndex < PX_strlen(str); CurrentCharIndex++ {
		if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) != 0) {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == 'e' && (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '-') {
				CurrentCharIndex++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == '.' {
				if Dot != 0 {
					return int32(0)
				} else {
					Dot = int32(1)
					continue
				}
			}
			return int32(0)
		}
	}
	return int32(1)
}
func PX_strIsFloat(str *int8) int32 {
	var Dot int32 = int32(0)
	var CurrentCharIndex int32
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '\x00' {
		return int32(0)
	}
	if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != 0) && !(int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) == '-') {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(int32(0))))) != '.' {
			return int32(0)
		}
	}
	for CurrentCharIndex = int32(1); CurrentCharIndex < PX_strlen(str); CurrentCharIndex++ {
		if !(PX_charIsNumeric(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) != 0) {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == 'e' && (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '+' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex+int32(1))))) == '-') {
				CurrentCharIndex++
				continue
			}
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(str)) + uintptr(CurrentCharIndex)))) == '.' {
				if Dot != 0 {
					return int32(0)
				} else {
					Dot = int32(1)
					continue
				}
			}
			return int32(0)
		}
	}
	return func() int32 {
		if Dot != 0 {
			return int32(1)
		} else {
			return int32(0)
		}
	}()
}
func PX_charIsNumeric(chr int8) int32 {
	if int32(chr) >= '0' && int32(chr) <= '9' {
		return int32(1)
	}
	return int32(0)
}
func PX_strIsInt(str *int8) int32 {
	if PX_strIsNumeric(str) != 0 {
		if PX_strIsFloat(str) != 0 {
			return int32(0)
		}
		return int32(1)
	}
	return int32(0)
}
func _cgos___px_pow_i_PX_Typedef(num float64, n int32) float64 {
	var powint float64 = float64(int32(1))
	var i int32
	for i = int32(1); i <= n; i++ {
		powint *= num
		if powint == float64(int32(0)) {
			break
		}
		if powint > float64(int32(0)) {
			if powint >= 1.7976931348623157e+308 || powint <= 4.9406564584124654e-324 {
				PX_ASSERT()
				break
			}
		} else if powint <= -1.7976931348623157e+308 || powint >= -4.9406564584124654e-324 {
			PX_ASSERT()
			break
		}
	}
	return powint
}
func _cgos___px_pow_f_PX_Typedef(num float64, m float64) float64 {
	return PX_exp(m * PX_log(num))
}
func PX_pow(num float64, m float64) float64 {
	if m-float64(int32(m)) == float64(int32(0)) {
		return _cgos___px_pow_i_PX_Typedef(num, int32(m))
	} else {
		return __px_pow_dd(num, m)
	}
	return 0
}
func PX_pow_ii(i int32, n int32) int32 {
	var fin int32 = int32(1)
	if n == int32(0) {
		return int32(1)
	}
	if n < int32(0) {
		return int32(0)
	}
	for func() (_cgo_ret int32) {
		_cgo_addr := &n
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0 {
		fin *= i
	}
	return fin
}
func _cgos___int_as_float_PX_Typedef(a int32) float32 {
	var r float32
	PX_memcpy(unsafe.Pointer(&r), unsafe.Pointer(&a), int32(4))
	return r
}
func _cgos___float_as_int_PX_Typedef(a float32) int32 {
	var r int32
	PX_memcpy(unsafe.Pointer(&r), unsafe.Pointer(&a), int32(4))
	return r
}
func PX_ln(x float64) float64 {
	var hfsq float64
	var f float64
	var s float64
	var z float64
	var R float64
	var w float64
	var t1 float64
	var t2 float64
	var dk float64
	var k int32
	var hx int32
	var i int32
	var j int32
	var lx uint32
	var zero float64 = 0
	hx = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))
	lx = uint32(*(*int32)(unsafe.Pointer(&x)))
	k = int32(0)
	if hx < int32(1048576) {
		if uint32(hx&int32(2147483647))|lx == uint32(0) {
			return -_cgos_two54_PX_Typedef / zero
		}
		if hx < int32(0) {
			return (x - x) / zero
		}
		k -= int32(54)
		x *= _cgos_two54_PX_Typedef
		hx = *(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4))
	}
	if hx >= int32(2146435072) {
		return x + x
	}
	k += hx>>int32(20) - int32(1023)
	hx &= int32(1048575)
	i = (hx + int32(614244)) & int32(1048576)
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&x)))) + uintptr(int32(1))*4)) = hx | (i ^ int32(1072693248))
	k += i >> int32(20)
	f = x - 1
	if int32(1048575)&(int32(2)+hx) < int32(3) {
		if f == zero {
			if k == int32(0) {
				return zero
			} else {
				dk = float64(k)
				return dk*_cgos_ln2_hi_PX_Typedef + dk*_cgos_ln2_lo_PX_Typedef
			}
		}
		R = f * f * (0.5 - 0.33333333333333331*f)
		if k == int32(0) {
			return f - R
		} else {
			dk = float64(k)
			return dk*_cgos_ln2_hi_PX_Typedef - (R - dk*_cgos_ln2_lo_PX_Typedef - f)
		}
	}
	s = f / (2 + f)
	dk = float64(k)
	z = s * s
	i = hx - int32(398458)
	w = z * z
	j = int32(440401) - hx
	t1 = w * (_cgos_Lg2_PX_Typedef + w*(_cgos_Lg4_PX_Typedef+w*_cgos_Lg6_PX_Typedef))
	t2 = z * (_cgos_Lg1_PX_Typedef + w*(_cgos_Lg3_PX_Typedef+w*(_cgos_Lg5_PX_Typedef+w*_cgos_Lg7_PX_Typedef)))
	i |= j
	R = t2 + t1
	if i > int32(0) {
		hfsq = 0.5 * f * f
		if k == int32(0) {
			return f - (hfsq - s*(hfsq+R))
		} else {
			return dk*_cgos_ln2_hi_PX_Typedef - (hfsq - (s*(hfsq+R) + dk*_cgos_ln2_lo_PX_Typedef) - f)
		}
	} else if k == int32(0) {
		return f - s*(f-R)
	} else {
		return dk*_cgos_ln2_hi_PX_Typedef - (s*(f-R) - dk*_cgos_ln2_lo_PX_Typedef - f)
	}
	return 0
}

var _cgos_ln2_hi_PX_Typedef float64 = 0.69314718036912382
var _cgos_ln2_lo_PX_Typedef float64 = 1.9082149292705877e-10
var _cgos_two54_PX_Typedef float64 = 18014398509481984
var _cgos_Lg1_PX_Typedef float64 = 0.66666666666667351
var _cgos_Lg2_PX_Typedef float64 = 0.39999999999409419
var _cgos_Lg3_PX_Typedef float64 = 0.28571428743662391
var _cgos_Lg4_PX_Typedef float64 = 0.22222198432149784
var _cgos_Lg5_PX_Typedef float64 = 0.1818357216161805
var _cgos_Lg6_PX_Typedef float64 = 0.15313837699209373
var _cgos_Lg7_PX_Typedef float64 = 0.14798198605116586

func PX_log(__x float64) float64 {
	return PX_ln(__x)
}
func PX_lg(__x float64) float64 {
	return PX_log10(__x)
}
func PX_log10(__x float64) float64 {
	return PX_ln(__x) / 2.3025850929940459
}
func PX_tand(radian float64) float64 {
	return float64(PX_tan_radian(float32(radian)))
}

var _cgos_px_srand_seed_PX_Typedef uint64 = uint64(826366246)

func PX_srand(seed uint64) {
	seed = (seed*uint64(16807) + uint64(1)) % uint64(4026531839)
	_cgos_px_srand_seed_PX_Typedef = seed
}
func PX_rand() uint32 {
	return uint32(func() (_cgo_ret uint64) {
		_cgo_addr := &_cgos_px_srand_seed_PX_Typedef
		*_cgo_addr = (_cgos_px_srand_seed_PX_Typedef*uint64(314159269) + uint64(453806245)) % uint64(18446744071562067968)
		return *_cgo_addr
	}()) & uint32(16777215)
}
func PX_randRange(min float64, max float64) float64 {
	return min + float64(PX_rand())*1/float64(int32(16777215))*(max-min)
}
func PX_randEx(seed uint64) uint32 {
	return uint32(func() (_cgo_ret uint64) {
		_cgo_addr := &seed
		*_cgo_addr = seed * uint64(764261123) % uint64(4026531839)
		return *_cgo_addr
	}())
}
func PX_GaussRand() float64 {
	var u float64
	var v float64
	var r float64
	var c float64
	for int32(1) != 0 {
		u = float64(PX_rand())/float64(int32(16777215))*float64(int32(2)) - float64(int32(1))
		v = float64(PX_rand())/float64(int32(16777215))*float64(int32(2)) - float64(int32(1))
		r = u*u + v*v
		if r == float64(int32(0)) || r > float64(int32(1)) {
			continue
		}
		c = float64(PX_sqrt(float32(float64(-2) * PX_ln(r) / r)))
		return u * c
	}
	return 0
}
func PX_Ceil(v float64) float64 {
	return func() float64 {
		if v-float64(int32(v)) != 0 {
			return float64(int32(v + float64(int32(1))))
		} else {
			return v
		}
	}()
}
func PX_FileGetName(filefullName *int8, _out *int8, outSize int32) {
	var s int32
	if outSize == int32(0) {
		return
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(0)))) = int8(0)
	s = PX_strlen(filefullName)
	if s == int32(0) {
		return
	}
	s--
	for s != 0 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) == '/' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) == '\\' {
			s++
			break
		}
		s--
	}
	for outSize > int32(1) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) != 0 && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) != '.' {
		outSize--
		*_out = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}())))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(1)))) = int8('\x00')
		*(*uintptr)(unsafe.Pointer(&_out))++
	}
}
func PX_FileGetPath(filefullName *int8, _out *int8, outSize int32) {
	var s int32
	var i int32
	if outSize == int32(0) {
		return
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(0)))) = int8(0)
	s = PX_strlen(filefullName)
	if s == int32(0) {
		return
	}
	s--
	for s != 0 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) == '/' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) == '\\' {
			break
		}
		s--
	}
	for i = int32(0); i < s; i++ {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(i))) = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(i)))
		if i >= outSize-int32(1) {
			*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(0)))) = int8(0)
			return
		}
	}
	if i > int32(0) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(i-int32(1))))) == ':' {
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(i))) = int8('/')
		i++
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(i))) = int8(0)
}
func PX_FileGetExt(filefullName *int8, _out *int8, outSize int32) {
	var s int32
	var bDot int32 = int32(0)
	if outSize == int32(0) {
		return
	}
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(0)))) = int8(0)
	s = PX_strlen(filefullName)
	if s == int32(0) {
		return
	}
	s--
	for s != 0 {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) == '.' {
			s++
			bDot = int32(1)
			break
		}
		s--
	}
	for outSize > int32(1) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(s)))) != 0 {
		outSize--
		*_out = *(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(filefullName)) + uintptr(func() (_cgo_ret int32) {
			_cgo_addr := &s
			_cgo_ret = *_cgo_addr
			*_cgo_addr++
			return
		}())))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(_out)) + uintptr(int32(1)))) = int8('\x00')
		*(*uintptr)(unsafe.Pointer(&_out))++
	}
	func() int {
		_ = bDot
		return 0
	}()
}

var _cgos_crc32tab_PX_Typedef [256]uint32 = [256]uint32{uint32(0), uint32(1996959894), uint32(3993919788), uint32(2567524794), uint32(124634137), uint32(1886057615), uint32(3915621685), uint32(2657392035), uint32(249268274), uint32(2044508324), uint32(3772115230), uint32(2547177864), uint32(162941995), uint32(2125561021), uint32(3887607047), uint32(2428444049), uint32(498536548), uint32(1789927666), uint32(4089016648), uint32(2227061214), uint32(450548861), uint32(1843258603), uint32(4107580753), uint32(2211677639), uint32(325883990), uint32(1684777152), uint32(4251122042), uint32(2321926636), uint32(335633487), uint32(1661365465), uint32(4195302755), uint32(2366115317), uint32(997073096), uint32(1281953886), uint32(3579855332), uint32(2724688242), uint32(1006888145), uint32(1258607687), uint32(3524101629), uint32(2768942443), uint32(901097722), uint32(1119000684), uint32(3686517206), uint32(2898065728), uint32(853044451), uint32(1172266101), uint32(3705015759), uint32(2882616665), uint32(651767980), uint32(1373503546), uint32(3369554304), uint32(3218104598), uint32(565507253), uint32(1454621731), uint32(3485111705), uint32(3099436303), uint32(671266974), uint32(1594198024), uint32(3322730930), uint32(2970347812), uint32(795835527), uint32(1483230225), uint32(3244367275), uint32(3060149565), uint32(1994146192), uint32(31158534), uint32(2563907772), uint32(4023717930), uint32(1907459465), uint32(112637215), uint32(2680153253), uint32(3904427059), uint32(2013776290), uint32(251722036), uint32(2517215374), uint32(3775830040), uint32(2137656763), uint32(141376813), uint32(2439277719), uint32(3865271297), uint32(1802195444), uint32(476864866), uint32(2238001368), uint32(4066508878), uint32(1812370925), uint32(453092731), uint32(2181625025), uint32(4111451223), uint32(1706088902), uint32(314042704), uint32(2344532202), uint32(4240017532), uint32(1658658271), uint32(366619977), uint32(2362670323), uint32(4224994405), uint32(1303535960), uint32(984961486), uint32(2747007092), uint32(3569037538), uint32(1256170817), uint32(1037604311), uint32(2765210733), uint32(3554079995), uint32(1131014506), uint32(879679996), uint32(2909243462), uint32(3663771856), uint32(1141124467), uint32(855842277), uint32(2852801631), uint32(3708648649), uint32(1342533948), uint32(654459306), uint32(3188396048), uint32(3373015174), uint32(1466479909), uint32(544179635), uint32(3110523913), uint32(3462522015), uint32(1591671054), uint32(702138776), uint32(2966460450), uint32(3352799412), uint32(1504918807), uint32(783551873), uint32(3082640443), uint32(3233442989), uint32(3988292384), uint32(2596254646), uint32(62317068), uint32(1957810842), uint32(3939845945), uint32(2647816111), uint32(81470997), uint32(1943803523), uint32(3814918930), uint32(2489596804), uint32(225274430), uint32(2053790376), uint32(3826175755), uint32(2466906013), uint32(167816743), uint32(2097651377), uint32(4027552580), uint32(2265490386), uint32(503444072), uint32(1762050814), uint32(4150417245), uint32(2154129355), uint32(426522225), uint32(1852507879), uint32(4275313526), uint32(2312317920), uint32(282753626), uint32(1742555852), uint32(4189708143), uint32(2394877945), uint32(397917763), uint32(1622183637), uint32(3604390888), uint32(2714866558), uint32(953729732), uint32(1340076626), uint32(3518719985), uint32(2797360999), uint32(1068828381), uint32(1219638859), uint32(3624741850), uint32(2936675148), uint32(906185462), uint32(1090812512), uint32(3747672003), uint32(2825379669), uint32(829329135), uint32(1181335161), uint32(3412177804), uint32(3160834842), uint32(628085408), uint32(1382605366), uint32(3423369109), uint32(3138078467), uint32(570562233), uint32(1426400815), uint32(3317316542), uint32(2998733608), uint32(733239954), uint32(1555261956), uint32(3268935591), uint32(3050360625), uint32(752459403), uint32(1541320221), uint32(2607071920), uint32(3965973030), uint32(1969922972), uint32(40735498), uint32(2617837225), uint32(3943577151), uint32(1913087877), uint32(83908371), uint32(2512341634), uint32(3803740692), uint32(2075208622), uint32(213261112), uint32(2463272603), uint32(3855990285), uint32(2094854071), uint32(198958881), uint32(2262029012), uint32(4057260610), uint32(1759359992), uint32(534414190), uint32(2176718541), uint32(4139329115), uint32(1873836001), uint32(414664567), uint32(2282248934), uint32(4279200368), uint32(1711684554), uint32(285281116), uint32(2405801727), uint32(4167216745), uint32(1634467795), uint32(376229701), uint32(2685067896), uint32(3608007406), uint32(1308918612), uint32(956543938), uint32(2808555105), uint32(3495958263), uint32(1231636301), uint32(1047427035), uint32(2932959818), uint32(3654703836), uint32(1088359270), uint32(936918000), uint32(2847714899), uint32(3736837829), uint32(1202900863), uint32(817233897), uint32(3183342108), uint32(3401237130), uint32(1404277552), uint32(615818150), uint32(3134207493), uint32(3453421203), uint32(1423857449), uint32(601450431), uint32(3009837614), uint32(3294710456), uint32(1567103746), uint32(711928724), uint32(3020668471), uint32(3272380065), uint32(1510334235), uint32(755167117)}

func PX_crc32(buffer unsafe.Pointer, size uint32) uint32 {
	var i uint32
	var crc uint32
	crc = uint32(4294967295)
	for i = uint32(0); i < size; i++ {
		crc = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint32)(unsafe.Pointer(&_cgos_crc32tab_PX_Typedef)))) + uintptr((crc^uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(buffer))) + uintptr(i)))))&uint32(255))*4)) ^ crc>>int32(8)
	}
	return crc ^ uint32(4294967295)
}
func PX_crc16(buffer unsafe.Pointer, size uint32) uint16 {
	var pos uint32
	var i uint32
	var crc uint16 = uint16(65535)
	var p_data *uint8 = (*uint8)(buffer)
	for pos = uint32(0); pos < size; pos++ {
		crc ^= uint16(int32(uint16(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(p_data)) + uintptr(pos))))))
		for i = uint32(8); i != uint32(0); i-- {
			if int32(crc)&int32(1) != int32(0) {
				crc >>= int32(1)
				crc ^= uint16(40961)
			} else {
				crc >>= int32(1)
			}
		}
	}
	crc = uint16(int32(crc)&int32(255)<<int32(8) | int32(crc)&int32(65280)>>int32(8))
	return crc
}
func PX_strchr(s *int8, ch int32) *int8 {
	for int32(*s) != '\x00' {
		if int32(*s)-ch == int32(0) {
			return (*int8)(unsafe.Pointer(s))
		}
		*(*uintptr)(unsafe.Pointer(&s))++
	}
	return (*int8)(nil)
}
func PX_strcmp(str1 *int8, str2 *int8) int32 {
	var ret int32 = int32(0)
	for !(func() (_cgo_ret int32) {
		_cgo_addr := &ret
		*_cgo_addr = int32(*(*uint8)(unsafe.Pointer(str1))) - int32(*(*uint8)(unsafe.Pointer(str2)))
		return *_cgo_addr
	}() != 0) && int32(*str1) != 0 {
		*(*uintptr)(unsafe.Pointer(&str1))++
		*(*uintptr)(unsafe.Pointer(&str2))++
	}
	return ret
}
func PX_strstr(dest *int8, src *int8) *int8 {
	var start *int8 = (*int8)(unsafe.Pointer(dest))
	var substart *int8 = (*int8)(unsafe.Pointer(src))
	var cp *int8 = (*int8)(unsafe.Pointer(dest))
	for *cp != 0 {
		start = cp
		for int32(*start) != '\x00' && int32(*substart) != '\x00' && int32(*start) == int32(*substart) {
			*(*uintptr)(unsafe.Pointer(&start))++
			*(*uintptr)(unsafe.Pointer(&substart))++
		}
		if int32(*substart) == '\x00' {
			return cp
		}
		substart = (*int8)(unsafe.Pointer(src))
		*(*uintptr)(unsafe.Pointer(&cp))++
	}
	return (*int8)(nil)
}
func PX_strcut(dest *int8, left int32, right int32) {
	var len int32 = PX_strlen(dest)
	if len == int32(0) {
		return
	}
	if right < left {
		var t int32 = left
		left = right
		right = t
	}
	if left < int32(0) {
		left = int32(0)
	}
	if left > len-int32(1) {
		left = len - int32(1)
	}
	if right < int32(0) {
		right = int32(0)
	}
	if right > len-int32(1) {
		right = len - int32(1)
	}
	PX_memcpy(unsafe.Pointer(dest), unsafe.Pointer((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest))+uintptr(left)))), right-left+int32(1))
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(dest)) + uintptr(right-left+int32(1)))) = int8(0)
}
func PX_isPointInCircle(p Struct__px_point, circle Struct__px_point, radius float32) int32 {
	if (p.x-circle.x)*(p.x-circle.x)+(p.y-circle.y)*(p.y-circle.y) <= radius*radius {
		return int32(1)
	}
	return int32(0)
}
func PX_isPoint2DInCircle(p _cgoa_6_PX_Typedef, circle _cgoa_6_PX_Typedef, radius float32) int32 {
	if (p.x-circle.x)*(p.x-circle.x)+(p.y-circle.y)*(p.y-circle.y) <= radius*radius {
		return int32(1)
	}
	return int32(0)
}
func PX_inet_addr(cp *int8) uint32 {
	var ipBytes [4]uint8 = [4]uint8{uint8(0)}
	var i int32
	for i = int32(0); i < int32(4); func() *int8 {
		i++
		return func() (_cgo_ret *int8) {
			_cgo_addr := &cp
			_cgo_ret = *_cgo_addr
			*(*uintptr)(unsafe.Pointer(_cgo_addr))++
			return
		}()
	}() {
		*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer((*uint8)(unsafe.Pointer(&ipBytes)))) + uintptr(i))) = uint8(PX_atoi(cp))
		if !(func() (_cgo_ret *int8) {
			_cgo_addr := &cp
			*_cgo_addr = PX_strchr(cp, '.')
			return *_cgo_addr
		}() != nil) {
			break
		}
	}
	return *(*uint32)(unsafe.Pointer((*uint8)(unsafe.Pointer(&ipBytes))))
}
func PX_inet_port(cp *int8) uint32 {
	var p *int8 = cp
	for *p != 0 {
		if int32(*p) == ':' {
			*(*uintptr)(unsafe.Pointer(&p))++
			break
		}
		*(*uintptr)(unsafe.Pointer(&p))++
	}
	return uint32(PX_atoi(p))
}
func PX_inet_ntoa(ipv4 uint32) _cgoa_1_PX_Typedef {
	var b [4]int32
	var ret _cgoa_1_PX_Typedef = _cgoa_1_PX_Typedef{[64]int8{int8(0)}}
	var a *int8 = (*int8)(unsafe.Pointer(&ret.data))
	*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(int32(0)))) = int8('\x00')
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(0))*4)) = int32(ipv4 & uint32(4278190080) >> int32(24))
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(1))*4)) = int32(ipv4 & uint32(16711680) >> int32(16))
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(2))*4)) = int32(ipv4 & uint32(65280) >> int32(8))
	*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(3))*4)) = int32(ipv4 & uint32(255) >> int32(0))
	PX_itoa(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(3))*4)), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a))+uintptr(PX_strlen(a)))), int32(5), int32(10))
	PX_strcat(a, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})))
	PX_itoa(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(2))*4)), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a))+uintptr(PX_strlen(a)))), int32(5), int32(10))
	PX_strcat(a, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})))
	PX_itoa(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(1))*4)), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a))+uintptr(PX_strlen(a)))), int32(5), int32(10))
	PX_strcat(a, (*int8)(unsafe.Pointer(&[2]int8{'.', '\x00'})))
	PX_itoa(*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer((*int32)(unsafe.Pointer(&b)))) + uintptr(int32(0))*4)), (*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(a))+uintptr(PX_strlen(a)))), int32(5), int32(10))
	return ret
}
func PX_IsValidIPAddress(ip_addr *int8) int32 {
	var len int32
	var position int32 = int32(0)
	var bitsNumber int32 = int32(0)
	var bitsLen int32
	var bits int32
	var n int32
	var i int32
	len = PX_strlen(ip_addr)
	if len < int32(7) || len > int32(15) {
		return int32(0)
	}
	for i = int32(0); i < len; i++ {
		if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(ip_addr)) + uintptr(i)))) == '.' || i == len-int32(1) {
			bitsLen = i - position
			if i == len-int32(1) {
				bitsLen++
			}
			if bitsLen > int32(3) || bitsLen == int32(0) {
				return int32(0)
			}
			if bitsLen != int32(1) && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(ip_addr)) + uintptr(position)))) == '0' {
				return int32(0)
			}
			for func() int32 {
				n = int32(0)
				return func() (_cgo_ret int32) {
					_cgo_addr := &bits
					*_cgo_addr = int32(0)
					return *_cgo_addr
				}()
			}(); n < bitsLen; func() int32 {
				n++
				return func() (_cgo_ret int32) {
					_cgo_addr := &position
					_cgo_ret = *_cgo_addr
					*_cgo_addr++
					return
				}()
			}() {
				bits = int32(10)*bits + (int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(ip_addr)) + uintptr(position)))) - '0')
			}
			if bits > int32(255) {
				return int32(0)
			}
			position++
			bitsNumber++
		} else if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(ip_addr)) + uintptr(i)))) >= '0' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer(ip_addr)) + uintptr(i)))) <= '9' {
			continue
		} else {
			return int32(0)
		}
	}
	if bitsNumber != int32(4) {
		return int32(0)
	}
	return int32(1)
}
func PX_TimeFormString(t *int8) Struct___px_timestamp {
	var time Struct___px_timestamp = Struct___px_timestamp{int16(0), 0, 0, 0, 0, 0}
	var data [20]int8 = [20]int8{int8(0)}
	if PX_strlen(t) == int32(10) {
		PX_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&data))), unsafe.Pointer(t), int32(10))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(4)))) = int8(0)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(7)))) = int8(0)
		time.year = int16(PX_atoi((*int8)(unsafe.Pointer(&data))))
		time.month = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(5))))))
		time.day = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(8))))))
	} else if PX_strlen(t) == int32(19) {
		PX_memcpy(unsafe.Pointer((*int8)(unsafe.Pointer(&data))), unsafe.Pointer(t), int32(19))
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(4)))) = int8(0)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(7)))) = int8(0)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(10)))) = int8(0)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(13)))) = int8(0)
		*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(16)))) = int8(0)
		time.year = int16(PX_atoi((*int8)(unsafe.Pointer(&data))))
		time.month = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(5))))))
		time.day = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(8))))))
		time.hour = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(11))))))
		time.minute = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(14))))))
		time.second = int16(PX_atoi((*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&data)))) + uintptr(int32(17))))))
	}
	return time
}
func PX_htonl(h uint32) uint32 {
	return func() uint32 {
		if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
			return h
		} else {
			return uint32(h)&uint32(4278190080)>>int32(24) | uint32(h)&uint32(16711680)>>int32(8) | uint32(h)&uint32(65280)<<int32(8) | uint32(h)&uint32(255)<<int32(24)
		}
	}()
}
func PX_ntohl(n uint32) uint32 {
	return func() uint32 {
		if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
			return n
		} else {
			return uint32(n)&uint32(4278190080)>>int32(24) | uint32(n)&uint32(16711680)>>int32(8) | uint32(n)&uint32(65280)<<int32(8) | uint32(n)&uint32(255)<<int32(24)
		}
	}()
}
func PX_htons(h uint16) uint16 {
	return uint16(func() int32 {
		if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
			return int32(h)
		} else {
			return int32(uint16(h))&int32(65280)>>int32(8) | int32(uint16(h))&int32(255)<<int32(8)
		}
	}())
}
func PX_ntohs(n uint16) uint16 {
	return uint16(func() int32 {
		if _cgos_PX_isBigEndianCPU_PX_Typedef() != 0 {
			return int32(n)
		} else {
			return int32(uint16(n))&int32(65280)>>int32(8) | int32(uint16(n))&int32(255)<<int32(8)
		}
	}())
}
func PX_sum32(buffer unsafe.Pointer, size uint32) uint32 {
	var _sum32 uint32 = uint32(0)
	var i uint32
	var pbuffer *uint8 = (*uint8)(buffer)
	for i = uint32(0); i < size; i++ {
		_sum32 += uint32(*(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(pbuffer)) + uintptr(i))))
	}
	return _sum32
}
func PX_PointRotate(p Struct__px_point, angle float32) Struct__px_point {
	var mat Struct__px_matrix
	PX_MatrixRotateZ(&mat, angle)
	return PX_PointMulMatrix(p, mat)
}
func PX_Point2DRotate(p _cgoa_6_PX_Typedef, angle float32) _cgoa_6_PX_Typedef {
	var mat Struct__px_matrix
	PX_MatrixRotateZ(&mat, angle)
	return PX_Point2DMulMatrix(p, mat)
}
func PX_PointDistance(p1 Struct__px_point, p2 Struct__px_point) float32 {
	return PX_sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}
func PX_WindowFunction_triangular(data *float64, N int32) {
	var i int32
	if N%int32(2) == int32(1) {
		for i = int32(0); i < (N-int32(1))/int32(2); i++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) = float64(int32(2)) * float64(i+int32(1)) / float64(N+int32(1))
		}
		for i = (N - int32(1)) / int32(2); i < N; i++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) = float64(int32(2)) * float64(N-i) / float64(N+int32(1))
		}
	} else {
		for i = int32(0); i < N/int32(2); i++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) = float64(i+i+int32(1)) * float64(int32(1)) / float64(N)
		}
		for i = N / int32(2); i < N; i++ {
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(N-int32(1)-i)*8))
		}
	}
}
func PX_WindowFunction_Apply(data *float64, window *float64, N int32) {
	var i int32
	for i = int32(0); i < N; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)*8)) *= *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(window)) + uintptr(i)*8))
	}
}
func PX_gain(b *float64, a *float64, m int32, n int32, x *float64, y *float64, len int32, sign int32) {
	var i int32
	var k int32
	var ar float64
	var ai float64
	var br float64
	var bi float64
	var zr float64
	var zi float64
	var im float64
	var re float64
	var den float64
	var numr float64
	var numi float64
	var freq float64
	var temp float64
	for k = int32(0); k < len; k++ {
		freq = float64(k) * 0.5 / float64(len-int32(1))
		zr = PX_cosd(-8 * PX_atan(1) * freq)
		zi = PX_sind(-8 * PX_atan(1) * freq)
		br = float64(int32(0))
		bi = float64(int32(0))
		for i = m; i > int32(0); i-- {
			re = br
			im = bi
			br = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(i)*8)))*zr - im*zi
			bi = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(i)*8)))*zi + im*zr
		}
		ar = float64(0)
		ai = float64(int32(0))
		for i = n; i > int32(0); i-- {
			re = ar
			im = ai
			ar = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)))*zr - im*zi
			ai = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(i)*8)))*zi + im*zr
		}
		br = br + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(int32(0))*8))
		ar = ar + 1
		numr = ar*br + ai*bi
		numi = ar*bi - ai*br
		den = ar*ar + ai*ai
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = numr / den
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = numi / den
		switch sign {
		case int32(1):
			{
				temp = PX_sqrtd(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)))
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = PX_atan2(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)))
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = temp
			}
			break
		case int32(2):
			{
				temp = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = PX_atan2(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)))
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = float64(int32(10)) * PX_log10(temp)
			}
			break
		}
	}
}
func PX_gainc(b *float64, a *float64, n int32, ns int32, x *float64, y *float64, len int32, sign int32) {
	var i int32
	var j int32
	var k int32
	var nl int32
	var ar float64
	var ai float64
	var br float64
	var bi float64
	var zr float64
	var zi float64
	var im float64
	var re float64
	var den float64
	var numr float64
	var numi float64
	var freq float64
	var temp float64
	var hr float64
	var hi float64
	var tr float64
	var ti float64
	nl = n + int32(1)
	for k = int32(0); k < len; k++ {
		freq = float64(k) * 0.5 / float64(len-int32(1))
		zr = PX_cosd(-8 * PX_atan(1) * freq)
		zi = PX_sind(-8 * PX_atan(1) * freq)
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = float64(1)
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = float64(0)
		for j = int32(0); j < ns; j++ {
			br = float64(int32(0))
			bi = float64(int32(0))
			for i = n; i > int32(0); i-- {
				re = br
				im = bi
				br = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(j*nl+i)*8)))*zr - im*zi
				bi = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(j*nl+i)*8)))*zi + im*zr
			}
			ar = float64(int32(0))
			ai = float64(int32(0))
			for i = n; i > int32(0); i-- {
				re = ar
				im = ai
				ar = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(a)) + uintptr(j*nl+i)*8)))*zr - im*zi
				ai = (re+*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(j*nl+i)*8)))*zi + im*zr
			}
			br = br + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(b)) + uintptr(j*nl+int32(0))*8))
			ar = ar + 1
			numr = ar*br + ai*bi
			numi = ar*bi - ai*br
			den = ar*ar + ai*ai
			hr = numr / den
			hi = numi / den
			tr = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))*hr - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))*hi
			ti = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))*hi + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))*hr
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = tr
			*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = ti
		}
		switch sign {
		case int32(1):
			{
				temp = PX_sqrtd(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)))
				if temp != float64(int32(0)) {
					*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = PX_atan2(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)))
				} else {
					*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = float64(int32(0))
				}
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = temp
			}
			break
		case int32(2):
			{
				temp = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8))
				if temp != float64(int32(0)) {
					*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = PX_atan2(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)))
				} else {
					temp = float64(9.9999999999999992e-41)
					*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(y)) + uintptr(k)*8)) = float64(int32(0))
				}
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + uintptr(k)*8)) = float64(int32(10)) * PX_log10(temp)
			}
			break
		}
	}
}
func PX_WindowFunction_blackMan(data *float64, N int32) {
	var n int32
	for n = int32(0); n < N; n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = 0.41999999999999998 - 0.5*PX_cosd(float64(float32(float64(int32(2))*3.1415926535897931*float64(n)/float64(N-int32(1)))+0.0799999982*PX_cos_radian(float32(float64(int32(4))*3.1415926535897931*float64(n)/float64(N-int32(1))))))
	}
}
func PX_WindowFunction_hamming(data *float64, N int32) {
	var n int32
	for n = int32(0); n < N; n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = 0.54000000000000004 - 0.46000000000000002*PX_cosd(float64(float32(float64(int32(2))*3.1415926535897931*float64(n)/float64(N-int32(1)))))
	}
}
func PX_WindowFunction_sinc(data *float64, N int32) {
	var n int32
	for n = int32(0); n < N; n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(N-n-int32(1))*8)) = func() (_cgo_ret float64) {
			_cgo_addr := &*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8))
			*_cgo_addr = PX_sind(3.1415926535897931 * (float64(n) + 0.5) / float64(N))
			return *_cgo_addr
		}()
	}
}
func PX_Bessel(n int32, x float64) float64 {
	var i int32
	var m int32
	var t float64
	var y float64
	var p float64 = float64(int32(0))
	var b0 float64
	var b1 float64
	var q float64
	if n < int32(0) {
		n = -n
	}
	t = func() float64 {
		if x > float64(int32(0)) {
			return x
		} else {
			return -x
		}
	}()
	if n != int32(1) {
		if t < 3.75 {
			y = x / 3.75 * (x / 3.75)
			p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_a_PX_Typedef)))) + uintptr(int32(6))*8))
			for i = int32(5); i >= int32(0); i-- {
				p = p*y + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_a_PX_Typedef)))) + uintptr(i)*8))
			}
		} else {
			y = 3.75 / t
			p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_c_PX_Typedef)))) + uintptr(int32(8))*8))
			for i = int32(7); i >= int32(0); i-- {
				p = p*y + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_c_PX_Typedef)))) + uintptr(i)*8))
			}
			p = p * PX_exp(t) / PX_sqrtd(t)
		}
	}
	if n == int32(0) {
		return p
	}
	q = p
	if t < 3.75 {
		y = x / 3.75 * (x / 3.75)
		p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_b_PX_Typedef)))) + uintptr(int32(6))*8))
		for i = int32(5); i >= int32(0); i-- {
			p = p*y + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_b_PX_Typedef)))) + uintptr(i)*8))
		}
		p = p * t
	} else {
		y = 3.75 / t
		p = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_d_PX_Typedef)))) + uintptr(int32(8))*8))
		for i = int32(7); i >= int32(0); i-- {
			p = p*y + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&_cgos_d_PX_Typedef)))) + uintptr(i)*8))
		}
		p = p * PX_exp(t) / PX_sqrtd(t)
	}
	if x < 0 {
		p = -p
	}
	if n == int32(1) {
		return p
	}
	if x == 0 {
		return float64(0)
	}
	y = 2 / t
	t = float64(0)
	b1 = float64(1)
	b0 = float64(0)
	m = n + int32(PX_sqrtd(40*float64(n)))
	m = int32(2) * m
	for i = m; i > int32(0); i-- {
		p = b0 + float64(i)*y*b1
		b0 = b1
		b1 = p
		if func() float64 {
			if b1 > float64(int32(0)) {
				return b1
			} else {
				return -b1
			}
		}() > 1.0e+10 {
			t = t * 1.0e-10
			b0 = b0 * 1.0e-10
			b1 = b1 * 1.0e-10
		}
		if i == n {
			t = b0
		}
	}
	p = t * q / b1
	if x < 0 && n%int32(2) == int32(1) {
		p = -p
	}
	return p
}

var _cgos_a_PX_Typedef [7]float64 = [7]float64{1, 3.5156228999999999, 3.0899424, 1.2067492, 0.26597320000000002, 0.036076799999999999, 0.0045813}
var _cgos_b_PX_Typedef [7]float64 = [7]float64{0.5, 0.87890594, 0.51498869000000003, 0.15084934, 0.02658773, 0.0030153200000000002, 3.2411000000000001e-4}
var _cgos_c_PX_Typedef [9]float64 = [9]float64{0.39894227999999998, 0.01328592, 0.00225319, -0.0015756500000000001, 0.0091628100000000004, -0.020577060000000001, 0.026355369999999999, -0.016476330000000001, 0.00392377}
var _cgos_d_PX_Typedef [9]float64 = [9]float64{0.39894227999999998, -0.039880239999999997, -0.0036201800000000002, 0.0016380100000000001, -0.01031555, 0.02282967, -0.028953119999999999, 0.01787654, -0.0042005899999999997}

func PX_WindowFunction_tukey(data *float64, N int32) {
	var n int32
	if N < int32(16) {
		return
	}
	for n = int32(0); n <= (N-int32(2))/int32(10); n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = 0.5 * (float64(int32(1)) - PX_cosd(float64(float32(float64(int32(10))*3.1415926535897931*float64(n)/float64(N+int32(8))))))
	}
	for n = (N - int32(2)) / int32(10); n <= int32(9)*(N-int32(2))/int32(10); n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = float64(int32(1))
	}
	for n = int32(9) * (N - int32(2)) / int32(10); n <= N-int32(1); n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = 0.5 * (float64(int32(1)) - PX_cosd(float64(float32(float64(int32(10))*3.1415926535897931*float64(N-n-int32(1))/float64(N+int32(8))))))
	}
}
func PX_WindowFunction_hanning(data *float64, N int32) {
	var n int32
	for n = int32(0); n < N; n++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = 0.5 - 0.5*PX_cosd(float64(float32(float64(int32(2))*3.1415926535897931*float64(n)/float64(N-int32(1)))))
	}
}
func PX_WindowFunction_kaiser(beta float64, data *float64, N int32) {
	var n int32
	var theta float64
	for n = int32(0); n < N; n++ {
		theta = beta * PX_sqrtd(float64(int32(1))-(float64(int32(2))*float64(n)/float64(N-int32(1))-float64(int32(1)))*(float64(int32(2))*float64(n)/float64(N-int32(1))-float64(int32(1))))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(n)*8)) = float64(PX_Bessel(int32(0), theta)) / PX_Bessel(int32(0), beta)
	}
}
func _cgos_PX_FIRKaiser_PX_Typedef(i int32, n int32, beta float64) float64 {
	var a float64
	var w float64
	var a2 float64
	var b1 float64
	var b2 float64
	var beta1 float64
	b1 = PX_Bessel(int32(0), beta)
	a = 2*float64(i)/float64(float64(n)-1) - 1
	a2 = a * a
	beta1 = beta * PX_sqrtd(1-a2)
	b2 = PX_Bessel(int32(0), beta1)
	w = b2 / b1
	return w
}
func _cgos_PX_FIRWindow_PX_Typedef(type_ int32, n int32, i int32, beta float64) float64 {
	var k int32
	var pi float64
	var w float64
	pi = 4 * PX_atan(1)
	w = float64(1)
	switch uint32(type_) {
	case uint32(0):
		{
			w = float64(1)
			break
		}
	case uint32(1):
		{
			k = (n - int32(2)) / int32(10)
			if i <= k {
				w = 0.5 * (1 - float64(PX_cos_radian(float32(float64(i)*pi/float64(float32(k)+1)))))
			}
			if i > n-k-int32(2) {
				w = 0.5 * float64(float32(1-float64(PX_cos_radian(float32(float64(float32(n-i)-1)*pi/float64(float32(k)+1))))))
			}
			break
		}
	case uint32(2):
		{
			w = 1 - func() float64 {
				if 1-float64(int32(2)*i)/(float64(n)-1) > float64(int32(0)) {
					return 1 - float64(int32(2)*i)/(float64(n)-1)
				} else {
					return -(1 - float64(int32(2)*i)/(float64(n)-1))
				}
			}()
			break
		}
	case uint32(3):
		{
			w = 0.5 * (1 - float64(PX_cos_radian(float32(float64(2*float32(i))*pi/float64(float32(n)-1)))))
			break
		}
	case uint32(4):
		{
			w = 0.54000000000000004 - 0.46000000000000002*float64(PX_cos_radian(float32(float64(2*float32(i))*pi/(float64(n)-1))))
			break
		}
	case uint32(5):
		{
			w = 0.41999999999999998 - 0.5*float64(PX_cos_radian(float32(float64(2*float32(i))*pi/float64(float32(n)-1)))) + 0.080000000000000002*float64(PX_cos_radian(float32(float64(4*float32(i))*pi/float64(float32(n)-1))))
			break
		}
	case uint32(6):
		{
			w = _cgos_PX_FIRKaiser_PX_Typedef(i, n, beta)
			break
		}
	}
	return w
}
func PX_SINE(A float64, P float64, F float64) _cgoa_13_PX_Typedef {
	var s _cgoa_13_PX_Typedef
	s.A = A
	s.p = P
	s.f = F
	return s
}
func PX_InstantaneousFrequency(src _cgoa_13_PX_Typedef, p2 float64, delta_t float64) _cgoa_13_PX_Typedef {
	var delta float64 = p2 - src.p
	delta = delta - float64(int32(delta)) - float64(0)
	delta /= delta_t
	src.f += delta
	return src
}
func PX_FIRFilterBuild(bandtype int32, fln float64, fhn float64, wn int32, h *float64, n int32, beta float64) {
	var i int32
	var n2 int32
	var mid int32
	var s float64
	var pi float64
	var wc1 float64
	var wc2 float64 = float64(int32(0))
	var delay float64
	pi = 4 * PX_atan(1)
	if n%int32(2) == int32(0) {
		n2 = n/int32(2) - int32(1)
		mid = int32(1)
	} else {
		n2 = n / int32(2)
		mid = int32(0)
	}
	delay = float64(n) / 2
	wc1 = 2 * pi * fln
	if uint32(bandtype) >= uint32(3) {
		wc2 = 2 * pi * fhn
	}
	switch uint32(bandtype) {
	case uint32(0):
		{
			for i = int32(0); i <= n2; i++ {
				s = float64(i) - delay
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = PX_sind(wc1*s) / (pi * s) * _cgos_PX_FIRWindow_PX_Typedef(wn, n+int32(1), i, beta)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n-i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8))
			}
			if mid == int32(1) {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n/int32(2))*8)) = wc1 / pi
			}
		}
		break
	case uint32(1):
		{
			for i = int32(0); i <= n2; i++ {
				s = float64(i) - delay
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = (PX_sind(pi*s) - PX_sind(wc1*s)) / (pi * s)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) * _cgos_PX_FIRWindow_PX_Typedef(wn, n+int32(1), i, beta)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n-i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8))
			}
			if mid == int32(1) {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n/int32(2))*8)) = 1 - wc1/pi
			}
		}
		break
	case uint32(2):
		{
			for i = int32(0); i <= n2; i++ {
				s = float64(i) - delay
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = (PX_sind(wc2*s) - PX_sind(wc1*s)) / (pi * s)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) * _cgos_PX_FIRWindow_PX_Typedef(wn, n+int32(1), i, beta)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n-i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8))
			}
			if mid == int32(1) {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n/int32(2))*8)) = (wc2 - wc1) / pi
			}
		}
		break
	case uint32(3):
		{
			for i = int32(0); i <= n2; i++ {
				s = float64(i) - delay
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = (PX_sind(wc1*s) - PX_sind(wc2*s)) / (pi * s)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8)) * _cgos_PX_FIRWindow_PX_Typedef(wn, n+int32(1), i, beta)
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n-i)*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(i)*8))
			}
			if mid == int32(1) {
				*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(h)) + uintptr(n/int32(2))*8)) = (wc1 + pi - wc2) / pi
			}
		}
		break
	}
}
func PX_GroupDelay(f float64, B *float64, sizeB int32, A *float64, sizeA int32, FS float64) float64 {
	var i int32
	var omega float64
	var Br [2]float64 = [2]float64{float64(int32(0))}
	var Be [2]float64 = [2]float64{float64(int32(0))}
	var Ar [2]float64 = [2]float64{float64(int32(0))}
	var Ae [2]float64 = [2]float64{float64(int32(1)), float64(int32(0))}
	var BrxAe [2]float64 = [2]float64{float64(int32(0))}
	var BexAr [2]float64 = [2]float64{float64(int32(0))}
	var AxB [2]float64 = [2]float64{float64(int32(0))}
	var Num [2]float64 = [2]float64{float64(int32(0))}
	var c [2]float64 = [2]float64{float64(int32(0))}
	var magd2 float64
	omega = float64(float64(2) * 3.1415926535897931 * f / FS)
	for i = int32(1); i < sizeB; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(0))*8)) += float64(i) * float64(PX_cosd(float64(i)*1*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(1))*8)) -= float64(i) * float64(PX_sind(float64(i)*1*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
	}
	for i = int32(0); i < sizeB; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(0))*8)) += float64(PX_cosd(float64(i)*1*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(1))*8)) -= float64(PX_sind(float64(i)*1*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
	}
	for i = int32(0); i < sizeA; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(0))*8)) += float64(i+int32(1)) * float64(PX_cosd((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(1))*8)) -= float64(i+int32(1)) * float64(PX_sind((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
	}
	for i = int32(0); i < sizeA; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(0))*8)) += float64(PX_cosd((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(1))*8)) -= float64(PX_sind((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
	}
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BrxAe)))) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(1))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BrxAe)))) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(1))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Br)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(0))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BexAr)))) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(1))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BexAr)))) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(1))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ar)))) + uintptr(int32(0))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(1))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(1))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Ae)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Be)))) + uintptr(int32(0))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(0))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BrxAe)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BexAr)))) + uintptr(int32(0))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(1))*8)) = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BrxAe)))) + uintptr(int32(1))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&BexAr)))) + uintptr(int32(1))*8))
	magd2 = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(0))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(1))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&c)))) + uintptr(int32(0))*8)) = (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(0))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(1))*8))) / magd2
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&c)))) + uintptr(int32(1))*8)) = (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&Num)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&AxB)))) + uintptr(int32(1))*8))) / magd2
	return *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&c)))) + uintptr(int32(0))*8))
}
func _cgos_PX_PhaseIIR_PX_Typedef(omega float64, B *float64, sizeB int32, A *float64, sizeA int32) float64 {
	var C [2]float64
	var HB [2]float64 = [2]float64{float64(int32(0))}
	var HA [2]float64 = [2]float64{float64(int32(0))}
	var magd2 float64
	var i int32
	for i = int32(0); i < sizeB; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(0))*8)) += float64(PX_cosd(float64(i)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(1))*8)) -= float64(PX_sind(float64(i)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(B)) + uintptr(i)*8))
	}
	for i = int32(0); i < sizeA; i++ {
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(0))*8)) += float64(PX_cosd((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
		*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(1))*8)) -= float64(PX_sind((float64(i)+1)*omega)) * *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer(A)) + uintptr(i)*8))
	}
	magd2 = *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(0))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(1))*8))
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&C)))) + uintptr(int32(0))*8)) = (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(0))*8)) + *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(1))*8))) / magd2
	*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&C)))) + uintptr(int32(1))*8)) = (*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(1))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(0))*8)) - *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HB)))) + uintptr(int32(0))*8))**(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&HA)))) + uintptr(int32(1))*8))) / magd2
	return float64(PX_atan2(*(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&C)))) + uintptr(int32(1))*8)), *(*float64)(unsafe.Pointer(uintptr(unsafe.Pointer((*float64)(unsafe.Pointer(&C)))) + uintptr(int32(0))*8))))
}
func PX_PhaseDelayDerive(omega float64, B *float64, sizeB int32, A *float64, sizeA int32, delta float64) float64 {
	var omega1 float64 = omega - delta
	var omega2 float64 = omega + delta
	var p1 float64 = _cgos_PX_PhaseIIR_PX_Typedef(omega1, B, sizeB, A, sizeA)
	var p2 float64 = _cgos_PX_PhaseIIR_PX_Typedef(omega2, B, sizeB, A, sizeA)
	return (-omega1*p2 + omega2*p1) / (float64(int32(2)) * delta * omega1 * omega2)
}
func PX_PhaseDelay(f float64, B *float64, sizeB int32, A *float64, sizeA int32, FS float64) float64 {
	var grpdel float64 = PX_GroupDelay(f, B, sizeB, A, sizeA, FS)
	var omega float64 = float64(2 * 3.1415926535897931 * f / FS)
	return grpdel - omega*PX_PhaseDelayDerive(omega, B, sizeB, A, sizeA, float64(5.00000024e-4))
}
