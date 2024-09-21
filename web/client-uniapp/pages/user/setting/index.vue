<template>
	<view class="page-content">
		<uv-toast ref="toast"></uv-toast>
		<uv-no-network></uv-no-network>
		<uv-navbar placeholder title="系统设置" leftText="返回" auto-back></uv-navbar>
		<view class="title">系统设置</view>
		<uv-cell-group :customStyle="{ color: '#84868c', background: '#fff' }" :border="false">
			<uv-cell
				@click="handleDeveloper"
				icon="setting-fill"
				title="开发者模式"
				:cell-style="{ padding: '30rpx' }"
				isLink
			></uv-cell>
			<!-- <uv-cell icon="integral-fill" title="版本更新" :cell-style="{ padding: '30rpx' }"></uv-cell> -->
		</uv-cell-group>

		<!-- 		<view class="title">更新记录</view>
		<uv-cell-group :customStyle="{ color: '#84868c', background: '#fff' }" :border="false">
			<uv-cell icon="setting-fill" title="更新记录" :cell-style="{ padding: '30rpx' }"></uv-cell>
			<uv-cell icon="integral-fill" title="版本更新" :cell-style="{ padding: '30rpx' }"></uv-cell>
		</uv-cell-group> -->

		<view class="title gap"></view>
		<view class="cell" @click="logout">切换账号</view>
		<view class="title gap"></view>
		<view class="cell danger" @click="logout">退出登陆</view>
	</view>
</template>

<script setup>
import { ref } from 'vue';
import { useUserStore } from '@/library/store/user';
import { nav } from '@/library/nav';
import Vconsole from 'vconsole';

const toast = ref();
const handleDeveloper = () => {
	new Vconsole();
	toast.value.success('已开启开发者模式');
};

const logout = () => {
	useUserStore().logout();
};
</script>

<style lang="scss" scoped>
.page-content {
	background-color: #f2f2f2;
	height: 100vh;
}
.title {
	padding: 20rpx 30rpx;
	color: #84868c;
}
.gap {
	padding: 15rpx !important;
}
.cell {
	padding: 30rpx;
	text-align: center;
	background-color: #fff;
}
.danger {
	color: $uv-error;
}
</style>
