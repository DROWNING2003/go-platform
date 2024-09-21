<template>
	<view class="page">
		<view class="header-bg bg-top">
			<uv-image :fade="false" :src="UserBg" width="100%" height="100%"></uv-image>
		</view>

		<view class="header">
			<view class="user" @click="handleNavUser">
				<view class="logo">
					<uv-image
						:src="$user.isLogin ? $rurl($user.avatarUrl, 100, 100) : $logo"
						radius="8rpx"
						width="120rpx"
						height="120rpx"
					></uv-image>
				</view>
				<template v-if="$user.isLogin">
					<view class="info">
						<view class="title">{{ $user.nickName }}</view>
						<view class="sub-title">uuid：{{ $user.id }}</view>
					</view>
					<view class="setting">
						<uv-icon name="arrow-right" color="#ccc" :size="20"></uv-icon>
					</view>
				</template>
				<template v-else>
					<view class="setting">
						<uv-icon
							label="立即登陆"
							labelPos="left"
							name="arrow-right"
							color="#84868c"
							:size="16"
						></uv-icon>
					</view>
				</template>
			</view>
		</view>

		<view class="content">
			<!-- 	<view class="card" style="margin-top: -30rpx">
				<uv-grid :border="false" :col="4">
					<uv-grid-item v-for="(item, index) in baseList" :key="index">
						<view class="grid-item">
							<view class="grid-icon"><uv-icon :name="item.name" :size="24"></uv-icon></view>
							<text class="grid-text">{{ item.title }}</text>
						</view>
					</uv-grid-item>
				</uv-grid>
			</view> -->

			<view class="card" style="margin-top: -30rpx">
        <uv-cell-group :border="false">
          <uv-cell
              icon="account-fill"
              title="个人信息"
              :border="false"
              clickable
              isLink
              @click="$to('/pages/user/info/index')"
          ></uv-cell>
          <uv-cell
              icon="warning-fill"
              title="问题反馈"
              clickable
              :border="false"
              isLink
              @click="$to('/pages/user/feedback/index')"
          ></uv-cell>
          <uv-cell
              icon="setting-fill"
              title="系统设置"
              clickable
              :border="false"
              isLink
              @click="$to('/pages/user/setting/index')"
          ></uv-cell>
        </uv-cell-group>
			</view>
		</view>
	</view>
</template>

<script setup>
import { useUserStore } from '@/library/store/user';
import UserBg from '@/static/images/user-bg.png';
import { nav } from '@/library/nav';

const userStore = useUserStore();
const handleNavUser = () => {
	if (!userStore.isLogin) {
		nav.login();
	} else {
		uni.navigateTo({
			url: './info/index'
		});
	}
};
const baseList = [
	// {
	// 	name: 'photo',
	// 	title: '李白李白'
	// },
];
</script>

<style lang="scss" scoped>
.page {
	position: relative;
}
.bg-top {
	z-index: 1;
}
.header-bg {
	position: absolute;
	width: 100%;
	height: 320rpx;
	&::after {
		z-index: 2;
		position: absolute;
		bottom: 0px;
		left: 0;
		display: block;
		width: 100%;
		height: 240rpx;
		background: url('@/static/images/wave.gif') no-repeat;
		background-size: 100% 100%;
		mix-blend-mode: screen;
		content: '';
	}
}

.header {
	position: relative;
	z-index: 3;
	width: 100%;
	height: 380rpx;
	padding: 80rpx 30rpx 0px 30rpx;
	box-sizing: border-box;
	.login-btn {
		width: 40%;
	}
	.user {
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: space-between;
		width: 100%;
		.logo {
			background-color: #fff;
			padding: 10rpx;
			margin-right: 30rpx;
		}
		.info {
			flex: 1;
			.title {
				font-size: $font-lgx;
				font-weight: bold;
				margin-bottom: 20rpx;
			}
			.sub-title {
				color: #84868c;
			}
		}
		.setting {
			// width: 100rpx;
		}
	}
}

.card {
	padding: 30rpx 0rpx;
	box-shadow: 0px 0px 25px 0px rgba(0, 0, 0, 0.07);
	border-radius: 8rpx;
	.grid-item {
		text-align: center;
		.grid-icon {
			padding: 20rpx;
			display: flex;
			justify-content: center;
		}
	}
}
</style>
