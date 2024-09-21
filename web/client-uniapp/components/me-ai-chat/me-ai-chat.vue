<template>
	<view class="chat">
		<uv-navbar placeholder leftIcon="" :title="title"></uv-navbar>
		<uv-toast ref="toast"></uv-toast>
		<uv-no-network></uv-no-network>
		<view
			class="chat-box"
			:style="{
				height: 'calc(100vh - ' + $uv.addUnit(44 + $uv.sys().statusBarHeight, 'px') + ')'
			}"
		>
			<view class="pre-content" v-if="!sessionId">
				<uv-image
					:src="logo"
					mode="widthFix"
					:lazy-load="true"
					width="80"
					height="80"
					radius="8"
					:custom-style="{ background: '#fff' }"
				></uv-image>
				<view class="app">{{ app ? app : title }}</view>
				<view class="welcome">
					<view class="welcome-text">{{ desc }}</view>
				</view>
				<view class="guiding" v-if="guiding.length">
					<view class="guiding-header">猜你想问</view>
					<view class="guiding-content">
						<view class="guiding-item" v-for="(item, index) in guiding" :key="index">
							<view class="guiding-text" @click="handleGuiding(item)">
								<template v-if="item.type === 'question'">
									<uv-icon color="#3c9cff" name="question-circle"></uv-icon>
									<text>{{ item.text }}</text>
								</template>
								<!-- 								
								<template v-if="item.type === 'webview'">
									<uv-icon color="#3c9cff" name="question-circle"></uv-icon>
									<text>{{ item.text }}</text>
								</template> -->
							</view>
						</view>
					</view>
				</view>
			</view>
			<view class="chat-content" v-else>
				<view class="header-msg">
					<view class="user">
						<uv-image
							:src="logo"
							mode="widthFix"
							:lazy-load="true"
							width="40"
							height="40"
							radius="4"
							:custom-style="{ background: '#fff' }"
						></uv-image>
						<view class="app">{{ app ? app : title }}</view>
					</view>
					<view class="welcome">
						<view class="welcome-text">{{ desc }}</view>
					</view>
				</view>

				<view class="chat-body">
					<view
						v-for="(item, index) in getMsgList()"
						:key="index"
						:class="item.type == 'reply' ? 'left-msg-box' : 'right-msg-box'"
						:id="'last_' + index"
					>
						<view v-if="item.type == 'reply'" class="left-msg">
							<template v-if="isProcess && index == 0">
								<view v-if="!item.text" class="innter-col-flex">
									<uv-image :src="loading" height="12px" width="24px"></uv-image>
									<text class="wait">正在思考中</text>
								</view>
								<view v-else>
									<text>{{ item.text }}</text>
									<uv-image
										class="innter-col-flex"
										:src="loading"
										height="12px"
										width="24px"
									></uv-image>
								</view>
							</template>
							<template v-else>
								<text>{{ item.text }}</text>
							</template>
						</view>
						<view v-else class="right-msg">{{ item.text }}</view>
					</view>
				</view>
			</view>
			<view class="chat-footer form">
				<view class="clear">
					<uv-icon name="clear" color="#ddd" size="20" custom-prefix="custom-icon"></uv-icon>
				</view>
				<uv-input
					placeholder="请输入内容"
					border="surround"
					v-model="inputValue"
					@change="handleInput"
					:maxlength="maxlength"
				>
					<template v-slot:suffix>
						<uv-icon
							v-if="isProcess"
							name="blue-pause"
							color="#3e8cfa"
							size="20"
							custom-prefix="custom-icon"
						></uv-icon>
						<uv-icon
							v-else
							@click="handleSend"
							size="20"
							:class="!isSend ? 'gray' : ''"
							:name="sendLogo"
						></uv-icon>
					</template>
				</uv-input>
			</view>
		</view>
	</view>
</template>

<script setup>
import { ref, defineProps, defineEmits, defineExpose, nextTick, onMounted } from 'vue';
import sendLogo from './send.png';
import loading from './loading.gif';
import { email } from '../../uni_modules/uv-ui-tools/libs/function/test';

const emit = defineEmits(['send']);
defineProps({
	title: String,
	logo: String,
	app: String,
	desc: String,
	guiding: {
		type: Array,
		default: () => {
			return [];
		}
	},
	maxlength: {
		type: Number,
		default: 120
	}
});

const sessionId = ref('1');
const isProcess = ref(false);

const inputValue = ref('');
const isSend = ref(false);
const msgList = ref([]);

const getMsgList = () => {
	return [...msgList.value].reverse();
};

const handleInput = (val) => {
	isSend.value = val.trim() != '';
};

const handleGuiding = (value) => {
	console.log(value);
};

const handleSend = () => {
	if (!(isSend.value && !isProcess.value)) {
		return;
	}
	emit('send', inputValue.value);
	msgList.value.push({ type: 'send', text: inputValue.value });
	inputValue.value = '';
	isSend.value = false;
};

const repleier = {
	send: (msg) => {
		console.log(msg);
		const index = msgList.value.length - 1;
		if (isProcess.value) {
			msgList.value[index].text += msg;
		} else {
			isProcess.value = true;
			msgList.value.push({ type: 'reply', text: inputValue.value });
		}
	},
	close: () => {
		isProcess.value = false;
	}
};

defineExpose({ repleier });
</script>
<style scoped>
.form .uv-input {
	background-color: #fff !important;
	border-radius: 8rpx !important;
}
</style>
<style lang="scss" scoped>
.chat {
	width: 100%;
	height: 100%;
	.chat-box {
		background-color: #f3f5f9;
		width: 100%;
		position: relative;
		display: flex;
		flex-direction: column;
		.chat-content {
			height: 100%;
			padding: 40rpx;
			display: flex;
			flex-direction: column;
			overflow: hidden;
			.header-msg {
				padding-bottom: 30rpx;
				margin-bottom: 30rpx;
				border-bottom: 1px solid #e8e9ed;
				.user {
					display: flex;
					align-items: center;
					.app {
						margin-left: 30rpx;
						font-size: $font-lg;
						font-weight: 700;
					}
				}
				.welcome {
					margin-top: 30rpx;
					width: 100%;
					.welcome-text {
						display: inline-block;
						border-radius: 16rpx;
						background: #fff;
						padding: 24rpx;
					}
				}
			}

			.chat-body {
				flex: 1;
				overflow: scroll;
				display: flex;
				flex-direction: column-reverse;
				.left-msg-box,
				.right-msg-box {
					display: block;
					width: 100%;
					padding-bottom: 40rpx;
				}
				.right-msg-box {
					display: flex;
					flex-direction: column;
					align-items: flex-end;
				}
				.left-msg,
				.right-msg {
					display: inline-block;
					border-radius: 16rpx;
					background: #fff;
					padding: 30rpx;
				}
				.left-msg {
					.innter-col-flex {
						display: inline-flex;
						flex-direction: row;
						align-items: center;
					}
					.wait {
						margin-left: 10rpx;
						color: #84868c;
					}
				}
				.right-msg {
					background-color: #3c9cff;
					color: #fff;
				}
			}
		}

		.pre-content {
			padding: 40rpx;
			flex: 1;
			display: flex;
			flex-direction: column;
			align-items: center;
			.app {
				padding: 30rpx;
				font-size: $font-lg;
				font-weight: 700;
			}
			.welcome {
				margin-top: 30rpx;
				width: 100%;
				.welcome-text {
					display: inline-block;
					border-radius: 16rpx;
					background: #e7ecf2;
					padding: 30rpx;
				}
			}
			.guiding {
				margin-top: 40rpx;
				width: 100%;
				box-sizing: border-box;
				.guiding-header {
					color: #84868c;
					margin-bottom: 20rpx;
				}
				.guiding-content {
					.guiding-item {
						margin-bottom: 16rpx;
						.guiding-text {
							display: inline-flex;
							padding: 20rpx 20rpx;
							background-color: #fff;
							color: #3c9cff;
							border-radius: 10rpx;

							text {
								margin-left: 10rpx;
							}
						}
					}
				}
			}
		}
		.chat-footer {
			display: flex;
			align-items: center;
			padding: 20rpx 40rpx;
			margin-bottom: 30rpx;
			.clear {
				margin-right: 30rpx;
			}
			.gray {
				filter: grayscale(80%);
				// &:hover {
				// 	filter: grayscale(0%);
				// }
			}
		}
	}
}
</style>
