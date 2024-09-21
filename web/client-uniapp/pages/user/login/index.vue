<template>
	<view>
		<uv-navbar placeholder leftIcon="" :title="'用户' + typeText"></uv-navbar>
		<uv-loading-page :loading="loading" loading-mode="semicircle"></uv-loading-page>
		<uv-toast ref="toast"></uv-toast>
		<uv-no-network></uv-no-network>
		<uv-modal
			ref="bindModal"
			title="绑定提示"
			content="当前登陆方式未绑定账号,请先进行绑定!"
			@confirm="confirmBind"
		></uv-modal>
		<view class="content">
			<view class="header">
				<view class="logo">
					<uv-image radius="8rpx" width="130rpx" height="130rpx" :src="$rurl(appStore.logoUrl)"></uv-image>
					<view class="tip" v-if="appStore.allowRegistry && type !== 'register'">
						<text class="disable">没有账号？</text>
						<text class="primary" @click="nav.register()">立即注册</text>
					</view>
				</view>

				<view class="app">
					<view class="title">
						{{ typeText }} {{ appStore.name }}
						<text class="version">{{ appStore.version }}</text>
					</view>
					<view class="desc">
						{{ appStore.description }}
					</view>
				</view>
				<view class="login">
					<uv-list :customStyle="{ gap: '20rpx' }" padding="30rpx">
						<template v-if="!loading && type === 'login' && suggest.name">
							<uv-list-item
								@click="handleLogin"
								clickable
								:title="suggest.name + typeText"
								:thumb="$rurl(suggest.logo)"
								thumb-size="lg"
								show-arrow
								:customStyle="{ backgroundColor: '#f4f5f9' }"
								rightText="推荐"
							></uv-list-item>
						</template>
						<template v-for="(item, index) in appStore.channels" :key="index">
							<template v-if="['password', 'email'].includes(item.keyword)">
								<uv-list-item
									:title="item.name + typeText"
									:thumb="$rurl(item.logo)"
									thumb-size="lg"
									show-arrow
									clickable
									@click="handleNav(item.keyword)"
									:customStyle="{ backgroundColor: '#f4f5f9' }"
								></uv-list-item>
							</template>
						</template>
					</uv-list>
				</view>
			</view>
		</view>
	</view>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue';
import { setToken } from '@/library/auth';
import { useAppStore } from '@/library/store/app';
import { useUserStore } from '@/library/store/user';
import { nav } from '@/library/nav';
import { OAuthLogin } from '@/api/system/usercenter';

const toast = ref();
const loading = ref(true);
const appStore = useAppStore();
const userStore = useUserStore();
const suggest = ref(appStore.suggestChannel);
const bindModal = ref();
const props = defineProps({ type: String, oAuthUid: String });
const typeText = appStore.loginTypes[props.type];
const curAuthUid = ref('');
const handleLogin = async () => {
	toast.value.show({ type: 'loading', message: '正在登录中', duration: 10 * 1000 });
	const code = await appStore.authCode().catch((err) => {
		nav.error('登陆失败:' + err);
		throw err;
	});
	const params = {
		channel: appStore.platform,
		code: code,
		app: appStore.keyword
	};
	const res = await OAuthLogin(params);
	toast.value.hide();
	if (res.isBind) {
		setToken(res.token);
		await userStore.userinfo();
		nav.home();
	} else {
		curAuthUid.value = res.oAuthUid;
		bindModal.value.open();
	}
};

const confirmBind = () => {
	nav.bind(curAuthUid.value);
};

const handleNav = (tp) => {
	if (tp === 'email') {
		let url = '/pages/user/login/email?type=' + props.type;
		if (props.oAuthUid) {
			url = url + '&oAuthUid=' + props.oAuthUid;
		}
		uni.navigateTo({ url: url });
	}

	if (tp === 'password') {
		let url = '/pages/user/login/password?type=' + props.type;
		if (props.oAuthUid) {
			url = url + '&oAuthUid=' + props.oAuthUid;
		}
		uni.navigateTo({ url: url });
	}
};

//自动登录回调，自动登录
onMounted(async () => {
	if (userStore.isLogin) {
		nav.home();
		return;
	}

	if (props.type !== 'login') {
		loading.value = false;
		return;
	}

	loading.value = false;
});
</script>

<style lang="scss" scoped>
.content {
	padding: 0px 30rpx;
	.header {
		.logo {
			padding: 30rpx 0px;
			display: flex;
			flex-direction: row;
			justify-content: space-between;
			align-items: center;
			.tip {
				.disable {
					color: #999;
				}
				.primary {
					color: #165dff;
				}
			}
		}
	}
	.app {
		margin-top: 20rpx;
		.title {
			font-size: 38rpx;
			margin-bottom: 10rpx;
			.version {
				font-size: 28rpx;
				color: #999;
			}
		}
		.desc {
			color: #999;
		}
	}
	.login {
		padding: 30rpx 0rpx;
	}
	.privacy-policy {
		margin-bottom: 60rpx;
	}
	.oauth {
		text-align: center;
		.desc {
			line-height: 50rpx;
		}
	}
}
</style>
