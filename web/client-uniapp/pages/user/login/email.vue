<template>
	<view>
		<uv-navbar placeholder :title="'邮箱' + typeText" auto-back></uv-navbar>

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
				<uv-form labelPosition="left" :model="form" :rules="rules" ref="formRef" labelWidth="0">
					<uv-form-item prop="email">
						<uv-input
							v-model="form.email"
							placeholder="请输入邮箱"
							prefix-icon="email"
							:customStyle="{ padding: '16rpx 10rpx' }"
							:prefixIconStyle="{ fontSize: '38rpx' }"
						></uv-input>
					</uv-form-item>
					<uv-form-item prop="captcha">
						<uv-input
							v-model="form.captcha"
							prefix-icon="empty-permission"
							placeholder="请输入验证码"
							:customStyle="{ padding: '16rpx 10rpx' }"
							:prefixIconStyle="{ fontSize: '38rpx' }"
						>
							<template v-slot:suffix>
								<uv-code
									ref="uCode"
									@end="uCodeSending = false"
									@start="uCodeSending = true"
									@change="codeChange"
									:seconds="captchaSecond"
									changeText="X秒重新后获取"
								></uv-code>

								<uv-tags
									@click="getCode"
									:text="uCodeTips"
									:disabled="uCodeSending"
									plain
									size="mini"
								></uv-tags>
							</template>
						</uv-input>
					</uv-form-item>

					<uv-form-item>
						<view class="submit">
							<uv-button :throttleTime="500" type="primary" :text="typeText" @click="submit"></uv-button>
						</view>
					</uv-form-item>
					<uv-form-item>
						<view class="submit">
							<uv-button
								type="primary"
								:plain="true"
								@click="$back()"
								:text="'更换' + typeText + '方式'"
							></uv-button>
						</view>
					</uv-form-item>
				</uv-form>
			</view>

			<uv-toast ref="toast"></uv-toast>
			<uv-no-network></uv-no-network>
		</view>
	</view>
</template>

<script setup>
import { ref, reactive, onMounted, watch } from 'vue';
import { setToken } from '@/library/auth';
import { useAppStore } from '@/library/store/app';
import { useUserStore } from '@/library/store/user';
import { nav } from '@/library/nav';
import { EmailBind, EmailLogin, EmailRegister, GenAuthCaptch } from '@/api/system/usercenter';

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
	email: '',
	captcha: '',
	captchaId: '',
	oAuthUid: props.oAuthUid,
	app: appStore.keyword
});

const rules = ref({
	email: [
		{
			required: true,
			message: '请输入邮箱',
			trigger: ['blur', 'change']
		},
		{
			type: 'email',
			message: '邮箱格式错误',
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

const codeChange = (text) => {
	uCodeTips.value = text;
};

const getCaptchaType = () => {
	switch (props.type) {
		case 'login':
			return 'loginEmail';
		case 'bind':
			return 'bindEmail';
		case 'register':
			return 'registerEmail';
		default:
			return '';
	}
};

const getCode = async () => {
	if (!form.email) {
		toast.value.error('请先输入邮箱');
		return;
	}

	const isEmail = uni.$uv.test.email(form.email);
	if (!isEmail) {
		toast.value.error('邮箱格式错误');
		return;
	}

	if (!uCode.value.canGetCode) {
		return;
	}

	const type = getCaptchaType();
	if (!type) {
		toast.value.error('场景类型错误');
		return;
	}

	const data = await GenAuthCaptch({ type: type, email: form.email });
	captchaSecond.value = data.expire;
	form.captchaId = data.id;
	toast.value.success('验证码已发送');
	uCode.value.start();
};

// 提交
const submit = async () => {
	await formRef.value.validate();
	if (!form.captchaId) {
		toast.value.error('请先获取验证码');
		return;
	}
	let res = {};
	switch (props.type) {
		case 'login':
			res = await EmailLogin(form);
			toast.value.success('登陆成功');
			break;
		case 'bind':
			res = await EmailBind(form);
			toast.value.success('绑定成功');
			break;
		case 'register':
			res = await EmailRegister(form);
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
</script>

<style lang="scss" scoped>
.content {
	padding: 0px 30rpx;

	.header {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		.logo {
			padding: 30rpx 0px;
			display: flex;
			flex-direction: row;
			justify-content: space-between;
			align-items: center;
			margin-right: 30rpx;
			.tip {
				.disable {
					color: #999;
				}
				.primary {
					color: #165dff;
				}
			}
		}
		.app {
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
	}

	.login {
		padding: 30rpx 0rpx;
		.submit {
			width: 100%;
		}
	}
	.privacy-policy {
		margin-bottom: 60rpx;
	}
}
</style>
