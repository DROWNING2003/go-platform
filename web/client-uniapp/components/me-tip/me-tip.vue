<template>
	<view class="tip">
		<view class="error" :style="{ width: size + 'rpx', height: size + 'rpx' }">
			<uv-no-network></uv-no-network>
			<image src="@/static/error/error.png"></image>
		</view>
		<div class="text">{{ text }}</div>
	</view>
</template>

<script setup>
defineProps({
	text: {
		type: String,
		default: '暂无数据'
	},
	size: {
		type: Number,
		default: () => 280
	}
});
</script>

<style lang="less" scoped>
.tip {
	display: flex;
	flex-direction: column;
	align-items: center;
	width: 100%;
	image {
		width: 100%;
		height: 100%;
	}
	.text {
		color: #333;
	}
}
</style>
