<template>
	<view>
		<uv-toast ref="toastRef"></uv-toast>
		<uv-no-network></uv-no-network>
		<uv-navbar placeholder auto-back leftText="返回" title="问题反馈"></uv-navbar>
		<view class="content">
			<uv-form ref="formRef" :model="form" :rules="rules" label-position="top" labelWidth="120" class="form">
				<uv-form-item label="分类" prop="categoryId" @click="showCategorySelect">
					<uv-input v-model="form.categoryName" disabled disabledColor="#ffffff" placeholder="请选择问题分类">
						<template v-slot:suffix>
							<uv-icon name="arrow-right"></uv-icon>
						</template>
					</uv-input>
				</uv-form-item>
				<uv-form-item label="标题" prop="title">
					<uv-input v-model="form.title" placeholder="请输入问题标题"></uv-input>
				</uv-form-item>
				<uv-form-item label="内容" prop="content">
					<uv-textarea v-model="form.content" placeholder="请输入问题内容"></uv-textarea>
				</uv-form-item>
				<uv-form-item :label="'截图(' + imageCount + '/3)'" prop="images">
					<me-upload
						width="100"
						height="100"
						:maxCount="3"
						accept="image"
						@change="uploadFeedbackImages"
						path="user/feedback"
					></me-upload>
				</uv-form-item>
				<uv-form-item label="联系方式" prop="contact">
					<uv-input v-model="form.contact" placeholder="请输入联系方式"></uv-input>
				</uv-form-item>
				<uv-form-item>
					<uv-button
						style="width: 100%"
						type="primary"
						text="提交"
						customStyle="margin-top: 10px"
						@click="submit"
					></uv-button>
				</uv-form-item>
			</uv-form>
		</view>
		<uv-picker ref="pickerRef" :columns="[categories]" key-name="name" @confirm="confirmCategory"></uv-picker>
	</view>
</template>

<script setup>
import { ref, getCurrentInstance } from 'vue';
import { useUserStore } from '@/library/store/user';
import { useAppStore } from '@/library/store/app';
import { CreateFeedback, ListFeedbackCategory } from '@/api/system/usercenter';

const userStore = useUserStore();
const appStore = useAppStore();
const sys = uni.$uv.sys();
const toastRef = ref();
const imageCount = ref(0);
const form = ref({
	appId: appStore.id,
	categoryId: undefined,
	categoryName: '',
	contact: userStore.phone,
	platform: appStore.platform,
	version: appStore.version,
	device: JSON.stringify({
		browserName: sys.browserName,
		browserVersion: sys.browserVersion,
		deviceType: sys.deviceType,
		platform: sys.platform,
		system: sys.system,
		screenHeight: sys.screenHeight,
		screenWidth: sys.screenWidth,
		orientation: sys.deviceOrientation,
		userAgent: sys.ua
	})
});
const formRef = ref();
const pickerRef = ref();
const categories = ref([]);
const rules = {
	categoryId: [
		{
			type: 'number',
			required: true,
			message: '请选择问题分类',
			trigger: ['blur', 'change']
		}
	],
	title: [
		{
			required: true,
			message: '请输入问题标题',
			trigger: ['blur', 'change']
		}
	],
	content: [
		{
			required: true,
			message: '请输入问题内容',
			trigger: ['blur', 'change']
		}
	]
};

const fetchFeedbackCategories = async () => {
	const params = { page: 1, pageSize: 50 };
	const data = await ListFeedbackCategory(params);
	categories.value = data.list;
};

fetchFeedbackCategories();

const confirmCategory = (obj) => {
	const value = obj.value[0];
	form.value.categoryId = value.id;
	form.value.categoryName = value.name;
};

const showCategorySelect = () => {
	pickerRef.value.open();
};

const uploadFeedbackImages = (files) => {
	if (!(files && files.length)) {
		return;
	}
	imageCount.value = files.length;
	const list = [];
	files.forEach((item) => {
		list.push(item.sha);
	});
	form.value.images = JSON.stringify(list);
};

const submit = async () => {
	await formRef.value.validate();
	uni.$uv.debounce(async () => {
		await CreateFeedback(form.value);
		toastRef.value.show({
			type: 'success',
			message: '反馈成功',
			complete() {
				uni.navigateBack();
			}
		});
	});
};
</script>

<style scoped lang="scss"></style>
