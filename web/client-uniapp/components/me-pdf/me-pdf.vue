<template>
	<view>
		<web-view :src="url"></web-view>
	</view>
</template>

<script>
export default {
	data() {
		return {
			url: ''
		};
	},
	onLoad(params) {
		let { url } = params;
		this.url = `/static/pdf.html?url=${url}`;
	}
};
</script>
