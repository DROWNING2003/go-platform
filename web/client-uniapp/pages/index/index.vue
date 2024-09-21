<template>
	<view class="hello">
		<uv-image radius="8rpx" width="250rpx" height="250rpx" :src="$logo"></uv-image>
		<div class="text">Hello Go-Platform</div>
	</view>
</template>

<script></script>

<style lang="less" scoped>
.hello {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	height: 100vh;
	.text {
		font-size: 30rpx;
		font-weight: 800;
	}
}
</style>
