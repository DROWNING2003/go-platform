<template>
	<view>
		<uv-toast ref="toast"></uv-toast>
		<uv-no-network></uv-no-network>
		<uv-navbar placeholder title="个人中心" leftText="返回" auto-back></uv-navbar>
		<uv-cell-group :border="false" style="margin-top: 30rpx">
			<uv-cell
				:title="$user.nickName"
				:label="'uid：' + $user.id"
				:border="false"
				@click="handleChangeAvatar"
				:custom-style="{ padding: '10rpx 0rpx' }"
			>
				<template v-slot:title>
					<view class="title">{{ $user.nickName }}</view>
				</template>
				<template v-slot:label>
					<view class="label margin-top-10">uid：{{ $user.id }}</view>
				</template>
				<template v-slot:value>
					<uv-avatar :size="50" :src="$rurl($user.avatarUrl, 100, 100)"></uv-avatar>
				</template>
			</uv-cell>
		</uv-cell-group>
		<uv-gap height="15" bgColor="#f2f2f2"></uv-gap>
		<uv-cell-group :border="false">
			<uv-cell
				icon="account"
				title="姓名"
				:value="$user.realName ? $user.realName : '未实名'"
				:cell-style="{ padding: '30rpx' }"
				:border="false"
			></uv-cell>
			<uv-cell
				icon="woman"
				title="性别"
				:value="gender[$user.gender]"
				:cell-style="{ padding: '30rpx' }"
				:border="false"
				isLink
				@click="genderPickerRef.open()"
			></uv-cell>
			<uv-cell
				icon="integral"
				title="账号"
				:value="$user.username ? $user.username : '未绑定'"
				:cell-style="{ padding: '30rpx' }"
				:border="false"
			></uv-cell>
			<uv-cell
				icon="phone"
				title="手机"
				:value="$user.phone ? $user.phone : '未绑定'"
				:cell-style="{ padding: '30rpx' }"
				:border="false"
			></uv-cell>
			<uv-cell
				icon="email"
				title="邮箱"
				:value="$user.email ? $user.email : '未绑定'"
				:cell-style="{ padding: '30rpx' }"
				:border="false"
			></uv-cell>
		</uv-cell-group>
		<view class="content"></view>

		<!-- <uv-modal ref="modalRef" title="标题" content="不知天上宫阙，今夕是何年" @confirm="confirm"></uv-modal> -->
		<me-info
			:value="showChangeUserinfo"
			:avatar="$rurl($user.avatarUrl, 120, 120)"
			:nickname="$user.nickName"
			@close="showChangeUserinfo = false"
			@comfirm="chooseAvatar"
		></me-info>
		<uv-picker ref="genderPickerRef" :columns="[genderList]" key-name="label" @confirm="chooseGender"></uv-picker>
	</view>
</template>

<script setup>
import { reactive, ref } from 'vue';
import { useUserStore } from '@/library/store/user';
import { UpdateCurrentUser } from '@/api/system/usercenter';

const gender = {
	U: '未知',
	M: '男',
	F: '女'
};

const genderList = [
	{ value: 'M', label: '男' },
	{ value: 'F', label: '女' },
	{ value: 'U', label: '未知' }
];

const modalRef = ref();
const showChangeUserinfo = ref(false);
const toast = ref();
const formRef = ref();
const rules = ref([]);
const genderPickerRef = ref();
const userStore = useUserStore();
const form = reactive({
	realName: userStore.realName,
	nickName: userStore.nickName,
	gender: userStore.gender,
	genderName: gender[userStore.gender],
	avatar: userStore.avatar
});

const handleChangeAvatar = () => {
	showChangeUserinfo.value = true;
};

const submit = async () => {
	await UpdateCurrentUser(form);
	toast.value.success('保存成功');
	await userStore.userinfo();
	showChangeUserinfo.value = false;
};

const showGenderPicker = () => {
	genderPickerRef.value.open();
};

const chooseGender = (item) => {
	const res = item.value[0];
	form.gender = res.value;
	form.genderName = res.label;
	submit();
};

const chooseAvatar = (item) => {
	form.avatar = item.avatar;
	form.nickName = item.nickname;
	submit();
};

const showRealName = () => {
	modalRef.value.open();
};
</script>

<style scoped lang="scss">
.title {
	font-size: $font-lg;
	font-weight: bold;
}
.label {
	font-size: $font-sm;
	color: #84868c;
}
</style>
