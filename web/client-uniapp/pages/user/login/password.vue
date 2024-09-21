<template>
	<view>
		<uv-navbar placeholder :title="'密码' + typeText" auto-back></uv-navbar>
		<uv-toast ref="toast"></uv-toast>
		<uv-no-network></uv-no-network>

		<view class="content">
			<view class="header">
				<view class="logo">
					<uv-image radius="8rpx" width="130rpx" height="130rpx" :src="$rurl(appStore.logoUrl)"></uv-image>
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
			</view>
			<view class="login form">
				<uv-form :model="form" :rules="rules" ref="formRef" labelWidth="0">
					<uv-form-item prop="username">
						<uv-input
							v-model="form.username"
							prefixIcon="account"
							placeholder="请输入账号"
							:customStyle="{ padding: '16rpx 10rpx' }"
							:prefixIconStyle="{ fontSize: '38rpx' }"
							clearable
						></uv-input>
					</uv-form-item>
					<uv-form-item prop="password">
						<uv-input
							v-model="form.password"
							password
							prefixIcon="lock"
							placeholder="请输入密码"
							:customStyle="{ padding: '16rpx 10rpx' }"
							:prefixIconStyle="{ fontSize: '38rpx' }"
							clearable
						></uv-input>
					</uv-form-item>
					<uv-form-item v-if="type === 'register'" prop="dupPassword">
						<uv-input
							v-model="form.dupPassword"
							password
							prefixIcon="lock"
							placeholder="请确认密码"
							:customStyle="{ padding: '16rpx 10rpx' }"
							:prefixIconStyle="{ fontSize: '38rpx' }"
							clearable
						></uv-input>
					</uv-form-item>
					<uv-form-item prop="captcha">
						<view class="captcha">
							<view class="captcha-input">
								<uv-input
									v-model="form.captcha"
									prefixIcon="empty-permission"
									placeholder="请输入验证码"
									:customStyle="{ padding: '16rpx 10rpx' }"
									:prefixIconStyle="{ fontSize: '38rpx' }"
									clearable
								>
									<template v-slot:suffix></template>
								</uv-input>
							</view>
							<view class="captcha-image" @click="fetchCaptcha">
								<image v-if="captchaBase64" class="image" mode="widthFix" :src="captchaBase64"></image>
							</view>
						</view>
					</uv-form-item>

					<uv-form-item>
						<view class="submit">
							<uv-button type="primary" :text="typeText" @click="submit"></uv-button>
						</view>
					</uv-form-item>
					<uv-form-item>
						<view class="submit">
							<uv-button
								type="primary"
								:plain="true"
								@click="$back"
								:text="'更换' + typeText + '方式'"
							></uv-button>
						</view>
					</uv-form-item>
				</uv-form>
			</view>
		</view>
	</view>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted, watch } from 'vue';
import { setToken } from '@/library/auth';
import { useAppStore } from '@/library/store/app';
import { useUserStore } from '@/library/store/user';
import { nav } from '@/library/nav';
import { PasswordBind, PasswordLogin, PasswordRegister, GenAuthCaptch } from '@/api/system/usercenter';

const captchaBase64 = ref('');
const timeInter = ref(null);
const toast = ref();
const appStore = useAppStore();
const uCode = ref();
const uCodeSending = ref(false);
const uCodeTips = ref('');
const formRef = ref(null);
const captchaSecond = ref(180);
const userStore = useUserStore();
const props = defineProps({ type: String, oAuthUid: String });
const typeText = appStore.loginTypes[props.type];

const form = reactive({
	username: '',
	password: '',
	dupPassword: '',
	oAuthUid: props.oAuthUid,
	captcha: '',
	captchaId: '',
	app: appStore.keyword
});

const rules = ref({
	username: [
		{
			required: true,
			message: '请输入账号',
			trigger: ['blur', 'change']
		},
		{
			validator: (rule, value, callback) => {
				if (!value) return true;
				const regex = /^[a-zA-Z][a-zA-Z0-9_]{5,11}$/;
				return regex.test(value);
			},
			message: '账号格式错误',
			trigger: ['blur', 'change']
		}
	],
	password: [
		{
			required: true,
			message: '请输入密码',
			trigger: ['blur', 'change']
		},
		{
			validator: (rule, value, callback) => {
				if (!value) return true;
				const regex = /^(?=.*[a-zA-Z])(?=.*[0-9]).{8,}$/;
				return regex.test(value);
			},
			message: '密码格式错误（包含英文、数字、最少8位）',
			trigger: ['blur', 'change']
		}
	],
	captcha: [
		{
			required: true,
			message: '请输入验证码',
			trigger: ['blur', 'change']
		}
	]
});

const registerRule = {
	dupPassword: [
		{
			validator: (rule, value, callback) => {
				if (!value) return true;
				return value === form.password;
			},
			message: '两次密码不一致',
			trigger: ['blur', 'change']
		}
	]
};

const fetchCaptcha = async () => {
	const type = getCaptchaType();
	if (!type) {
		toast.value.error('场景类型错误');
		return;
	}

	// 清除定时器
	clearInterval(timeInter.value);

	// 请求验证码
	const data = await GenAuthCaptch({ type: type });
	captchaSecond.value = data.expire;
	form.captchaId = data.id;
	captchaBase64.value = data.base64;
	// 定时刷新
	timeInter.value = setInterval(() => {
		fetchCaptcha();
	}, data.expire * 1000);
};

// 登录
const submit = async () => {
	await formRef.value.validate();
	if (!form.captchaId) {
		toast.value.error('请先获取验证码');
		return;
	}
	let res = {};
	switch (props.type) {
		case 'login':
			res = await PasswordLogin(form).catch((e) => {
				fetchCaptcha();
				throw e;
			});
			toast.value.success('登陆成功');
			break;
		case 'bind':
			res = await PasswordBind(form).catch((e) => {
				fetchCaptcha();
				throw e;
			});
			toast.value.success('绑定成功');
			break;
		case 'register':
			res = await PasswordRegister(form).catch((e) => {
				fetchCaptcha();
				throw e;
			});
			toast.value.success('注册成功');
			break;
	}
	if (!res.token) {
		toast.value.error('登陆失败！');
		return;
	}

	setToken(res.token);
	await userStore.userinfo();
	nav.home();
};

const getCaptchaType = () => {
	switch (props.type) {
		case 'login':
			return 'loginImage';
		case 'bind':
			return 'bindImage';
		case 'register':
			return 'registerImage';
		default:
			return '';
	}
};

onMounted(async () => {
	if (props.type === 'register') {
		rules.value = {
			...rules.value,
			...registerRule
		};
	}
	fetchCaptcha();
});

onUnmounted(() => {
	clearInterval(timeInter.value);
	timeInter.value = null;
});
</script>

<style lang="scss" scoped>
.content {
	padding: 0px 30rpx;

	.app {
		margin-top: 30rpx;
		.title {
			font-size: 38rpx;
		}
		.desc {
			color: #999;
		}
	}
	.login {
		padding: 30rpx 0rpx;
		.submit {
			width: 100%;
		}
		.captcha {
			display: flex;
			width: 100%;
			.captcha-input {
				flex: 1;
			}
			.captcha-image {
				min-width: 120px;
				max-width: 120px;
				max-height: 40px;
				min-height: 40px;
				border: 1px solid #f2f3f5;
				overflow: hidden;
				.image {
					height: 100%;
					width: 100%;
				}
			}
		}
	}
	.privacy-policy {
		margin-bottom: 60rpx;
	}
}
</style>
