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

// CUDA handle for madd5 kernel
var madd5_code cu.Function

// Stores the arguments for madd5 kernel invocation
type madd5_args_t struct {
	arg_dst  unsafe.Pointer
	arg_src1 unsafe.Pointer
	arg_fac1 float32
	arg_src2 unsafe.Pointer
	arg_fac2 float32
	arg_src3 unsafe.Pointer
	arg_fac3 float32
	arg_src4 unsafe.Pointer
	arg_fac4 float32
	arg_src5 unsafe.Pointer
	arg_fac5 float32
	arg_N    int
	argptr   [12]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for madd5 kernel invocation
var madd5_args madd5_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	madd5_args.argptr[0] = unsafe.Pointer(&madd5_args.arg_dst)
	madd5_args.argptr[1] = unsafe.Pointer(&madd5_args.arg_src1)
	madd5_args.argptr[2] = unsafe.Pointer(&madd5_args.arg_fac1)
	madd5_args.argptr[3] = unsafe.Pointer(&madd5_args.arg_src2)
	madd5_args.argptr[4] = unsafe.Pointer(&madd5_args.arg_fac2)
	madd5_args.argptr[5] = unsafe.Pointer(&madd5_args.arg_src3)
	madd5_args.argptr[6] = unsafe.Pointer(&madd5_args.arg_fac3)
	madd5_args.argptr[7] = unsafe.Pointer(&madd5_args.arg_src4)
	madd5_args.argptr[8] = unsafe.Pointer(&madd5_args.arg_fac4)
	madd5_args.argptr[9] = unsafe.Pointer(&madd5_args.arg_src5)
	madd5_args.argptr[10] = unsafe.Pointer(&madd5_args.arg_fac5)
	madd5_args.argptr[11] = unsafe.Pointer(&madd5_args.arg_N)
}

// Wrapper for madd5 CUDA kernel, asynchronous.
func k_madd5_async(dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, src3 unsafe.Pointer, fac3 float32, src4 unsafe.Pointer, fac4 float32, src5 unsafe.Pointer, fac5 float32, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("madd5")
	}

	madd5_args.Lock()
	defer madd5_args.Unlock()

	if madd5_code == 0 {
		madd5_code = fatbinLoad(madd5_map, "madd5")
	}

	madd5_args.arg_dst = dst
	madd5_args.arg_src1 = src1
	madd5_args.arg_fac1 = fac1
	madd5_args.arg_src2 = src2
	madd5_args.arg_fac2 = fac2
	madd5_args.arg_src3 = src3
	madd5_args.arg_fac3 = fac3
	madd5_args.arg_src4 = src4
	madd5_args.arg_fac4 = fac4
	madd5_args.arg_src5 = src5
	madd5_args.arg_fac5 = fac5
	madd5_args.arg_N = N

	args := madd5_args.argptr[:]
	cu.LaunchKernel(madd5_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("madd5")
	}
}

// maps compute capability on PTX code for madd5 kernel.
var madd5_map = map[int]string{0: "",
	60: madd5_ptx_60}

// madd5 PTX code for various compute capabilities.
const (
	madd5_ptx_60 = `
.version 7.5
.target sm_60
.address_size 64

	// .globl	madd5

.visible .entry madd5(
	.param .u64 madd5_param_0,
	.param .u64 madd5_param_1,
	.param .f32 madd5_param_2,
	.param .u64 madd5_param_3,
	.param .f32 madd5_param_4,
	.param .u64 madd5_param_5,
	.param .f32 madd5_param_6,
	.param .u64 madd5_param_7,
	.param .f32 madd5_param_8,
	.param .u64 madd5_param_9,
	.param .f32 madd5_param_10,
	.param .u32 madd5_param_11
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<16>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [madd5_param_0];
	ld.param.u64 	%rd2, [madd5_param_1];
	ld.param.f32 	%f1, [madd5_param_2];
	ld.param.u64 	%rd3, [madd5_param_3];
	ld.param.f32 	%f2, [madd5_param_4];
	ld.param.u64 	%rd4, [madd5_param_5];
	ld.param.f32 	%f3, [madd5_param_6];
	ld.param.u64 	%rd5, [madd5_param_7];
	ld.param.f32 	%f4, [madd5_param_8];
	ld.param.u64 	%rd6, [madd5_param_9];
	ld.param.f32 	%f5, [madd5_param_10];
	ld.param.u32 	%r2, [madd5_param_11];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32 	%p1, %r1, %r2;
	@%p1 bra 	$L__BB0_2;

	cvta.to.global.u64 	%rd7, %rd2;
	mul.wide.s32 	%rd8, %r1, 4;
	add.s64 	%rd9, %rd7, %rd8;
	ld.global.nc.f32 	%f6, [%rd9];
	cvta.to.global.u64 	%rd10, %rd3;
	add.s64 	%rd11, %rd10, %rd8;
	ld.global.nc.f32 	%f7, [%rd11];
	mul.f32 	%f8, %f7, %f2;
	fma.rn.f32 	%f9, %f6, %f1, %f8;
	cvta.to.global.u64 	%rd12, %rd4;
	add.s64 	%rd13, %rd12, %rd8;
	ld.global.nc.f32 	%f10, [%rd13];
	fma.rn.f32 	%f11, %f10, %f3, %f9;
	cvta.to.global.u64 	%rd14, %rd5;
	add.s64 	%rd15, %rd14, %rd8;
	ld.global.nc.f32 	%f12, [%rd15];
	fma.rn.f32 	%f13, %f12, %f4, %f11;
	cvta.to.global.u64 	%rd16, %rd6;
	add.s64 	%rd17, %rd16, %rd8;
	ld.global.nc.f32 	%f14, [%rd17];
	fma.rn.f32 	%f15, %f14, %f5, %f13;
	cvta.to.global.u64 	%rd18, %rd1;
	add.s64 	%rd19, %rd18, %rd8;
	st.global.f32 	[%rd19], %f15;

$L__BB0_2:
	ret;

}

`
)
