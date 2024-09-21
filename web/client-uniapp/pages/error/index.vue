<template>
	<view class="error">
		<uv-no-network></uv-no-network>
		<image src="@/static/error/error.png"></image>
		<div class="text">{{ text }}</div>
	</view>
</template>

<script>
export default {
	data() {
		return {
			text: '系统异常，请联系负责人员'
		};
	},
	onLoad(params) {
		if (params.text) {
			this.text = params.text;
		}
	}
};
</script>

<style lang="less" scoped>
.error {
	height: 100vh;
	overflow: hidden;
	display: flex;
	flex-direction: column;
	align-items: center;
	width: 100%;
	image {
		margin-top: 200rpx;
		width: 600rpx;
		height: 600rpx;
	}
	.text {
		font-size: 32rpx;
		color: #333;
		font-weight: 700;
	}
}
</style>
