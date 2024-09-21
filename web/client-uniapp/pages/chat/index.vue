<template>
	<me-ai-chat
		ref="chatRef"
		title="我的Agent应用-测试"
		logo="http://p3-armor.byteimg.com/tos-cn-i-49unhts6dw/dfdba5317c0c20ce20e64fac803d52bc.svg~tplv-49unhts6dw-image.image"
		desc="你好，同学。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？"
		:guiding="[
			{ text: '你好同学', type: 'question', path: 'http://baidu.com' },
			{ text: '你好同学', type: 'question', path: 'http://baidu.com' }
		]"
		@send="handleSend"
	></me-ai-chat>
	<!-- 。我是负责贫困资助的AI助手，请问你有什么想要咨询的问题吗？ -->
</template>

<script setup>
import { ref } from 'vue';
const chatRef = ref();
const handleSend = (val) => {
	console.log(val);
	handleReply();
};

const handleReply = () => {
	setInterval(() => {
		chatRef.value.repleier.send('你好你好你好你好你好');
	}, 1000);
	setInterval(() => {
		console.log('close');
		chatRef.value.repleier.close();
	}, 3 * 1000);
};
</script>

<style></style>
