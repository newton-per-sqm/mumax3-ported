package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/v3/cuda/cu"
	"github.com/mumax/3/v3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for addcubicanisotropy2 kernel
var addcubicanisotropy2_code cu.Function

// Stores the arguments for addcubicanisotropy2 kernel invocation
type addcubicanisotropy2_args_t struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_Ms_     unsafe.Pointer
	arg_Ms_mul  float32
	arg_k1_     unsafe.Pointer
	arg_k1_mul  float32
	arg_k2_     unsafe.Pointer
	arg_k2_mul  float32
	arg_k3_     unsafe.Pointer
	arg_k3_mul  float32
	arg_c1x_    unsafe.Pointer
	arg_c1x_mul float32
	arg_c1y_    unsafe.Pointer
	arg_c1y_mul float32
	arg_c1z_    unsafe.Pointer
	arg_c1z_mul float32
	arg_c2x_    unsafe.Pointer
	arg_c2x_mul float32
	arg_c2y_    unsafe.Pointer
	arg_c2y_mul float32
	arg_c2z_    unsafe.Pointer
	arg_c2z_mul float32
	arg_N       int
	argptr      [27]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for addcubicanisotropy2 kernel invocation
var addcubicanisotropy2_args addcubicanisotropy2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	addcubicanisotropy2_args.argptr[0] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Bx)
	addcubicanisotropy2_args.argptr[1] = unsafe.Pointer(&addcubicanisotropy2_args.arg_By)
	addcubicanisotropy2_args.argptr[2] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Bz)
	addcubicanisotropy2_args.argptr[3] = unsafe.Pointer(&addcubicanisotropy2_args.arg_mx)
	addcubicanisotropy2_args.argptr[4] = unsafe.Pointer(&addcubicanisotropy2_args.arg_my)
	addcubicanisotropy2_args.argptr[5] = unsafe.Pointer(&addcubicanisotropy2_args.arg_mz)
	addcubicanisotropy2_args.argptr[6] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Ms_)
	addcubicanisotropy2_args.argptr[7] = unsafe.Pointer(&addcubicanisotropy2_args.arg_Ms_mul)
	addcubicanisotropy2_args.argptr[8] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k1_)
	addcubicanisotropy2_args.argptr[9] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k1_mul)
	addcubicanisotropy2_args.argptr[10] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k2_)
	addcubicanisotropy2_args.argptr[11] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k2_mul)
	addcubicanisotropy2_args.argptr[12] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k3_)
	addcubicanisotropy2_args.argptr[13] = unsafe.Pointer(&addcubicanisotropy2_args.arg_k3_mul)
	addcubicanisotropy2_args.argptr[14] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1x_)
	addcubicanisotropy2_args.argptr[15] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1x_mul)
	addcubicanisotropy2_args.argptr[16] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1y_)
	addcubicanisotropy2_args.argptr[17] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1y_mul)
	addcubicanisotropy2_args.argptr[18] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1z_)
	addcubicanisotropy2_args.argptr[19] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c1z_mul)
	addcubicanisotropy2_args.argptr[20] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2x_)
	addcubicanisotropy2_args.argptr[21] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2x_mul)
	addcubicanisotropy2_args.argptr[22] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2y_)
	addcubicanisotropy2_args.argptr[23] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2y_mul)
	addcubicanisotropy2_args.argptr[24] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2z_)
	addcubicanisotropy2_args.argptr[25] = unsafe.Pointer(&addcubicanisotropy2_args.arg_c2z_mul)
	addcubicanisotropy2_args.argptr[26] = unsafe.Pointer(&addcubicanisotropy2_args.arg_N)
}

// Wrapper for addcubicanisotropy2 CUDA kernel, asynchronous.
func k_addcubicanisotropy2_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, Ms_ unsafe.Pointer, Ms_mul float32, k1_ unsafe.Pointer, k1_mul float32, k2_ unsafe.Pointer, k2_mul float32, k3_ unsafe.Pointer, k3_mul float32, c1x_ unsafe.Pointer, c1x_mul float32, c1y_ unsafe.Pointer, c1y_mul float32, c1z_ unsafe.Pointer, c1z_mul float32, c2x_ unsafe.Pointer, c2x_mul float32, c2y_ unsafe.Pointer, c2y_mul float32, c2z_ unsafe.Pointer, c2z_mul float32, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("addcubicanisotropy2")
	}

	addcubicanisotropy2_args.Lock()
	defer addcubicanisotropy2_args.Unlock()

	if addcubicanisotropy2_code == 0 {
		addcubicanisotropy2_code = fatbinLoad(addcubicanisotropy2_map, "addcubicanisotropy2")
	}

	addcubicanisotropy2_args.arg_Bx = Bx
	addcubicanisotropy2_args.arg_By = By
	addcubicanisotropy2_args.arg_Bz = Bz
	addcubicanisotropy2_args.arg_mx = mx
	addcubicanisotropy2_args.arg_my = my
	addcubicanisotropy2_args.arg_mz = mz
	addcubicanisotropy2_args.arg_Ms_ = Ms_
	addcubicanisotropy2_args.arg_Ms_mul = Ms_mul
	addcubicanisotropy2_args.arg_k1_ = k1_
	addcubicanisotropy2_args.arg_k1_mul = k1_mul
	addcubicanisotropy2_args.arg_k2_ = k2_
	addcubicanisotropy2_args.arg_k2_mul = k2_mul
	addcubicanisotropy2_args.arg_k3_ = k3_
	addcubicanisotropy2_args.arg_k3_mul = k3_mul
	addcubicanisotropy2_args.arg_c1x_ = c1x_
	addcubicanisotropy2_args.arg_c1x_mul = c1x_mul
	addcubicanisotropy2_args.arg_c1y_ = c1y_
	addcubicanisotropy2_args.arg_c1y_mul = c1y_mul
	addcubicanisotropy2_args.arg_c1z_ = c1z_
	addcubicanisotropy2_args.arg_c1z_mul = c1z_mul
	addcubicanisotropy2_args.arg_c2x_ = c2x_
	addcubicanisotropy2_args.arg_c2x_mul = c2x_mul
	addcubicanisotropy2_args.arg_c2y_ = c2y_
	addcubicanisotropy2_args.arg_c2y_mul = c2y_mul
	addcubicanisotropy2_args.arg_c2z_ = c2z_
	addcubicanisotropy2_args.arg_c2z_mul = c2z_mul
	addcubicanisotropy2_args.arg_N = N

	args := addcubicanisotropy2_args.argptr[:]
	cu.LaunchKernel(addcubicanisotropy2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("addcubicanisotropy2")
	}
}

// maps compute capability on PTX code for addcubicanisotropy2 kernel.
var addcubicanisotropy2_map = map[int]string{0: "",
	60: addcubicanisotropy2_ptx_60}

// addcubicanisotropy2 PTX code for various compute capabilities.
const (
	addcubicanisotropy2_ptx_60 = `
.version 7.5
.target sm_60
.address_size 64

	// .globl	addcubicanisotropy2

.visible .entry addcubicanisotropy2(
	.param .u64 addcubicanisotropy2_param_0,
	.param .u64 addcubicanisotropy2_param_1,
	.param .u64 addcubicanisotropy2_param_2,
	.param .u64 addcubicanisotropy2_param_3,
	.param .u64 addcubicanisotropy2_param_4,
	.param .u64 addcubicanisotropy2_param_5,
	.param .u64 addcubicanisotropy2_param_6,
	.param .f32 addcubicanisotropy2_param_7,
	.param .u64 addcubicanisotropy2_param_8,
	.param .f32 addcubicanisotropy2_param_9,
	.param .u64 addcubicanisotropy2_param_10,
	.param .f32 addcubicanisotropy2_param_11,
	.param .u64 addcubicanisotropy2_param_12,
	.param .f32 addcubicanisotropy2_param_13,
	.param .u64 addcubicanisotropy2_param_14,
	.param .f32 addcubicanisotropy2_param_15,
	.param .u64 addcubicanisotropy2_param_16,
	.param .f32 addcubicanisotropy2_param_17,
	.param .u64 addcubicanisotropy2_param_18,
	.param .f32 addcubicanisotropy2_param_19,
	.param .u64 addcubicanisotropy2_param_20,
	.param .f32 addcubicanisotropy2_param_21,
	.param .u64 addcubicanisotropy2_param_22,
	.param .f32 addcubicanisotropy2_param_23,
	.param .u64 addcubicanisotropy2_param_24,
	.param .f32 addcubicanisotropy2_param_25,
	.param .u32 addcubicanisotropy2_param_26
)
{
	.reg .pred 	%p<15>;
	.reg .f32 	%f<187>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<60>;


	ld.param.u64 	%rd1, [addcubicanisotropy2_param_0];
	ld.param.u64 	%rd2, [addcubicanisotropy2_param_1];
	ld.param.u64 	%rd3, [addcubicanisotropy2_param_2];
	ld.param.u64 	%rd4, [addcubicanisotropy2_param_3];
	ld.param.u64 	%rd5, [addcubicanisotropy2_param_4];
	ld.param.u64 	%rd6, [addcubicanisotropy2_param_5];
	ld.param.u64 	%rd7, [addcubicanisotropy2_param_6];
	ld.param.f32 	%f174, [addcubicanisotropy2_param_7];
	ld.param.u64 	%rd8, [addcubicanisotropy2_param_8];
	ld.param.f32 	%f176, [addcubicanisotropy2_param_9];
	ld.param.u64 	%rd9, [addcubicanisotropy2_param_10];
	ld.param.f32 	%f177, [addcubicanisotropy2_param_11];
	ld.param.u64 	%rd10, [addcubicanisotropy2_param_12];
	ld.param.f32 	%f178, [addcubicanisotropy2_param_13];
	ld.param.u64 	%rd11, [addcubicanisotropy2_param_14];
	ld.param.f32 	%f179, [addcubicanisotropy2_param_15];
	ld.param.u64 	%rd12, [addcubicanisotropy2_param_16];
	ld.param.f32 	%f180, [addcubicanisotropy2_param_17];
	ld.param.u64 	%rd13, [addcubicanisotropy2_param_18];
	ld.param.f32 	%f181, [addcubicanisotropy2_param_19];
	ld.param.u64 	%rd14, [addcubicanisotropy2_param_20];
	ld.param.f32 	%f183, [addcubicanisotropy2_param_21];
	ld.param.u64 	%rd15, [addcubicanisotropy2_param_22];
	ld.param.f32 	%f184, [addcubicanisotropy2_param_23];
	ld.param.u64 	%rd16, [addcubicanisotropy2_param_24];
	ld.param.f32 	%f185, [addcubicanisotropy2_param_25];
	ld.param.u32 	%r2, [addcubicanisotropy2_param_26];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	$L__BB0_28;

	setp.eq.s64 	%p2, %rd7, 0;
	@%p2 bra 	$L__BB0_3;

	cvta.to.global.u64 	%rd17, %rd7;
	mul.wide.s32 	%rd18, %r1, 4;
	add.s64 	%rd19, %rd17, %rd18;
	ld.global.nc.f32 	%f45, [%rd19];
	mul.f32 	%f174, %f45, %f174;

$L__BB0_3:
	setp.eq.f32 	%p3, %f174, 0f00000000;
	mov.f32 	%f175, 0f00000000;
	@%p3 bra 	$L__BB0_5;

	rcp.rn.f32 	%f175, %f174;

$L__BB0_5:
	setp.eq.s64 	%p4, %rd8, 0;
	@%p4 bra 	$L__BB0_7;

	cvta.to.global.u64 	%rd20, %rd8;
	mul.wide.s32 	%rd21, %r1, 4;
	add.s64 	%rd22, %rd20, %rd21;
	ld.global.nc.f32 	%f47, [%rd22];
	mul.f32 	%f176, %f47, %f176;

$L__BB0_7:
	setp.eq.s64 	%p5, %rd9, 0;
	@%p5 bra 	$L__BB0_9;

	cvta.to.global.u64 	%rd23, %rd9;
	mul.wide.s32 	%rd24, %r1, 4;
	add.s64 	%rd25, %rd23, %rd24;
	ld.global.nc.f32 	%f48, [%rd25];
	mul.f32 	%f177, %f48, %f177;

$L__BB0_9:
	setp.eq.s64 	%p6, %rd10, 0;
	@%p6 bra 	$L__BB0_11;

	cvta.to.global.u64 	%rd26, %rd10;
	mul.wide.s32 	%rd27, %r1, 4;
	add.s64 	%rd28, %rd26, %rd27;
	ld.global.nc.f32 	%f49, [%rd28];
	mul.f32 	%f178, %f49, %f178;

$L__BB0_11:
	mul.f32 	%f11, %f175, %f178;
	setp.eq.s64 	%p7, %rd11, 0;
	mul.f32 	%f12, %f175, %f176;
	mul.f32 	%f13, %f175, %f177;
	@%p7 bra 	$L__BB0_13;

	cvta.to.global.u64 	%rd29, %rd11;
	mul.wide.s32 	%rd30, %r1, 4;
	add.s64 	%rd31, %rd29, %rd30;
	ld.global.nc.f32 	%f50, [%rd31];
	mul.f32 	%f179, %f50, %f179;

$L__BB0_13:
	setp.eq.s64 	%p8, %rd12, 0;
	@%p8 bra 	$L__BB0_15;

	cvta.to.global.u64 	%rd32, %rd12;
	mul.wide.s32 	%rd33, %r1, 4;
	add.s64 	%rd34, %rd32, %rd33;
	ld.global.nc.f32 	%f51, [%rd34];
	mul.f32 	%f180, %f51, %f180;

$L__BB0_15:
	setp.eq.s64 	%p9, %rd13, 0;
	@%p9 bra 	$L__BB0_17;

	cvta.to.global.u64 	%rd35, %rd13;
	mul.wide.s32 	%rd36, %r1, 4;
	add.s64 	%rd37, %rd35, %rd36;
	ld.global.nc.f32 	%f52, [%rd37];
	mul.f32 	%f181, %f52, %f181;

$L__BB0_17:
	mul.f32 	%f54, %f180, %f180;
	fma.rn.f32 	%f55, %f179, %f179, %f54;
	fma.rn.f32 	%f56, %f181, %f181, %f55;
	sqrt.rn.f32 	%f20, %f56;
	setp.eq.f32 	%p10, %f20, 0f00000000;
	mov.f32 	%f182, 0f00000000;
	@%p10 bra 	$L__BB0_19;

	rcp.rn.f32 	%f182, %f20;

$L__BB0_19:
	mul.f32 	%f23, %f179, %f182;
	mul.f32 	%f24, %f180, %f182;
	mul.f32 	%f25, %f181, %f182;
	setp.eq.s64 	%p11, %rd14, 0;
	@%p11 bra 	$L__BB0_21;

	cvta.to.global.u64 	%rd38, %rd14;
	mul.wide.s32 	%rd39, %r1, 4;
	add.s64 	%rd40, %rd38, %rd39;
	ld.global.nc.f32 	%f57, [%rd40];
	mul.f32 	%f183, %f57, %f183;

$L__BB0_21:
	setp.eq.s64 	%p12, %rd15, 0;
	@%p12 bra 	$L__BB0_23;

	cvta.to.global.u64 	%rd41, %rd15;
	mul.wide.s32 	%rd42, %r1, 4;
	add.s64 	%rd43, %rd41, %rd42;
	ld.global.nc.f32 	%f58, [%rd43];
	mul.f32 	%f184, %f58, %f184;

$L__BB0_23:
	setp.eq.s64 	%p13, %rd16, 0;
	@%p13 bra 	$L__BB0_25;

	cvta.to.global.u64 	%rd44, %rd16;
	mul.wide.s32 	%rd45, %r1, 4;
	add.s64 	%rd46, %rd44, %rd45;
	ld.global.nc.f32 	%f59, [%rd46];
	mul.f32 	%f185, %f59, %f185;

$L__BB0_25:
	mul.f32 	%f61, %f184, %f184;
	fma.rn.f32 	%f62, %f183, %f183, %f61;
	fma.rn.f32 	%f63, %f185, %f185, %f62;
	sqrt.rn.f32 	%f32, %f63;
	setp.eq.f32 	%p14, %f32, 0f00000000;
	mov.f32 	%f186, 0f00000000;
	@%p14 bra 	$L__BB0_27;

	rcp.rn.f32 	%f186, %f32;

$L__BB0_27:
	mul.f32 	%f64, %f185, %f186;
	mul.f32 	%f65, %f24, %f64;
	mul.f32 	%f66, %f184, %f186;
	mul.f32 	%f67, %f25, %f66;
	sub.f32 	%f68, %f65, %f67;
	mul.f32 	%f69, %f183, %f186;
	mul.f32 	%f70, %f25, %f69;
	mul.f32 	%f71, %f23, %f64;
	sub.f32 	%f72, %f70, %f71;
	mul.f32 	%f73, %f23, %f66;
	mul.f32 	%f74, %f24, %f69;
	sub.f32 	%f75, %f73, %f74;
	cvta.to.global.u64 	%rd47, %rd4;
	mul.wide.s32 	%rd48, %r1, 4;
	add.s64 	%rd49, %rd47, %rd48;
	ld.global.nc.f32 	%f76, [%rd49];
	cvta.to.global.u64 	%rd50, %rd5;
	add.s64 	%rd51, %rd50, %rd48;
	ld.global.nc.f32 	%f77, [%rd51];
	mul.f32 	%f78, %f24, %f77;
	fma.rn.f32 	%f79, %f23, %f76, %f78;
	cvta.to.global.u64 	%rd52, %rd6;
	add.s64 	%rd53, %rd52, %rd48;
	ld.global.nc.f32 	%f80, [%rd53];
	fma.rn.f32 	%f81, %f25, %f80, %f79;
	mul.f32 	%f82, %f66, %f77;
	fma.rn.f32 	%f83, %f69, %f76, %f82;
	fma.rn.f32 	%f84, %f64, %f80, %f83;
	mul.f32 	%f85, %f77, %f72;
	fma.rn.f32 	%f86, %f76, %f68, %f85;
	fma.rn.f32 	%f87, %f75, %f80, %f86;
	mul.f32 	%f88, %f84, %f84;
	mul.f32 	%f89, %f87, %f87;
	add.f32 	%f90, %f88, %f89;
	mul.f32 	%f91, %f23, %f81;
	mul.f32 	%f92, %f24, %f81;
	mul.f32 	%f93, %f25, %f81;
	mul.f32 	%f94, %f81, %f81;
	add.f32 	%f95, %f94, %f89;
	mul.f32 	%f96, %f69, %f84;
	mul.f32 	%f97, %f66, %f84;
	mul.f32 	%f98, %f64, %f84;
	mul.f32 	%f99, %f96, %f95;
	mul.f32 	%f100, %f97, %f95;
	mul.f32 	%f101, %f98, %f95;
	fma.rn.f32 	%f102, %f91, %f90, %f99;
	fma.rn.f32 	%f103, %f92, %f90, %f100;
	fma.rn.f32 	%f104, %f93, %f90, %f101;
	add.f32 	%f105, %f94, %f88;
	mul.f32 	%f106, %f68, %f87;
	mul.f32 	%f107, %f72, %f87;
	mul.f32 	%f108, %f75, %f87;
	fma.rn.f32 	%f109, %f106, %f105, %f102;
	fma.rn.f32 	%f110, %f107, %f105, %f103;
	fma.rn.f32 	%f111, %f108, %f105, %f104;
	mul.f32 	%f112, %f12, 0fC0000000;
	mul.f32 	%f113, %f112, %f109;
	mul.f32 	%f114, %f112, %f110;
	mul.f32 	%f115, %f112, %f111;
	mul.f32 	%f116, %f88, %f89;
	mul.f32 	%f117, %f94, %f89;
	mul.f32 	%f118, %f96, %f117;
	mul.f32 	%f119, %f97, %f117;
	mul.f32 	%f120, %f98, %f117;
	fma.rn.f32 	%f121, %f91, %f116, %f118;
	fma.rn.f32 	%f122, %f92, %f116, %f119;
	fma.rn.f32 	%f123, %f93, %f116, %f120;
	mul.f32 	%f124, %f94, %f88;
	fma.rn.f32 	%f125, %f106, %f124, %f121;
	fma.rn.f32 	%f126, %f107, %f124, %f122;
	fma.rn.f32 	%f127, %f108, %f124, %f123;
	add.f32 	%f128, %f13, %f13;
	mul.f32 	%f129, %f128, %f125;
	mul.f32 	%f130, %f128, %f126;
	mul.f32 	%f131, %f128, %f127;
	sub.f32 	%f132, %f113, %f129;
	sub.f32 	%f133, %f114, %f130;
	sub.f32 	%f134, %f115, %f131;
	mul.f32 	%f135, %f88, %f88;
	mul.f32 	%f136, %f89, %f89;
	add.f32 	%f137, %f135, %f136;
	mul.f32 	%f138, %f81, %f94;
	mul.f32 	%f139, %f23, %f138;
	mul.f32 	%f140, %f24, %f138;
	mul.f32 	%f141, %f25, %f138;
	fma.rn.f32 	%f142, %f94, %f94, %f136;
	mul.f32 	%f143, %f84, %f88;
	mul.f32 	%f144, %f69, %f143;
	mul.f32 	%f145, %f66, %f143;
	mul.f32 	%f146, %f64, %f143;
	mul.f32 	%f147, %f144, %f142;
	mul.f32 	%f148, %f145, %f142;
	mul.f32 	%f149, %f146, %f142;
	fma.rn.f32 	%f150, %f139, %f137, %f147;
	fma.rn.f32 	%f151, %f140, %f137, %f148;
	fma.rn.f32 	%f152, %f141, %f137, %f149;
	fma.rn.f32 	%f153, %f94, %f94, %f135;
	mul.f32 	%f154, %f87, %f89;
	mul.f32 	%f155, %f68, %f154;
	mul.f32 	%f156, %f72, %f154;
	mul.f32 	%f157, %f75, %f154;
	fma.rn.f32 	%f158, %f153, %f155, %f150;
	fma.rn.f32 	%f159, %f153, %f156, %f151;
	fma.rn.f32 	%f160, %f153, %f157, %f152;
	mul.f32 	%f161, %f11, 0f40800000;
	mul.f32 	%f162, %f161, %f158;
	mul.f32 	%f163, %f161, %f159;
	mul.f32 	%f164, %f161, %f160;
	sub.f32 	%f165, %f132, %f162;
	sub.f32 	%f166, %f133, %f163;
	sub.f32 	%f167, %f134, %f164;
	cvta.to.global.u64 	%rd54, %rd1;
	add.s64 	%rd55, %rd54, %rd48;
	ld.global.f32 	%f168, [%rd55];
	add.f32 	%f169, %f168, %f165;
	st.global.f32 	[%rd55], %f169;
	cvta.to.global.u64 	%rd56, %rd2;
	add.s64 	%rd57, %rd56, %rd48;
	ld.global.f32 	%f170, [%rd57];
	add.f32 	%f171, %f170, %f166;
	st.global.f32 	[%rd57], %f171;
	cvta.to.global.u64 	%rd58, %rd3;
	add.s64 	%rd59, %rd58, %rd48;
	ld.global.f32 	%f172, [%rd59];
	add.f32 	%f173, %f172, %f167;
	st.global.f32 	[%rd59], %f173;

$L__BB0_28:
	ret;

}

`
)
