<template>
	<uv-upload :file-list="files" @delete="deleteFile" @afterRead="uploadFile">
		<slot></slot>
	</uv-upload>
</template>

<script setup>
import { ref, reactive } from 'vue';
import config from '@/config';
import crypto from '@/library/crypto';
import { PrepareUpload, Upload } from '@/api/system/resource.js';

const instance = uni.$global;

const props = defineProps({
	fileList: {
		type: Array,
		default: () => []
	},
	listType: {
		type: String,
		default: 'picture-card'
	},
	autoUpload: {
		type: Boolean,
		default: true
	},
	app: {
		type: String,
		require: true,
		default: () => {
			return config.server;
		}
	},
	path: {
		type: String,
		require: true
	}
});

const emit = defineEmits(['change', 'success']);

const files = reactive(props.fileList);

const deleteFile = (event) => {
	files.splice(event.index, 1);
	emit('change', files);
};

const uploadFile = async (event) => {
	const list = [].concat(event.file);
	const flen = files.length;
	list.map((item) => {
		const file = { ...item };
		uploadFactory(file);
	});
};

const binaryToSha = (binary) => {
	return crypto.md5(binary);
};

const toBinary = async (url) => {
	const response = await uni.request({
		url: url,
		responseType: 'arraybuffer'
	});
	return response.data;
};

const onSuccess = (file, data) => {
	files.forEach((item, index) => {
		if (item.url !== file.url) {
			return;
		}
		files[index].status = 'success';
		files[index].message = '';
		files[index].sha = data.sha;
		files[index].source = uni.$global.$rurl(data.src);
		emit('change', files);
		emit('success', files[index], index);
	});
};

const onFailed = (file) => {
	files.forEach((item, index) => {
		if (item.url !== file.url) {
			return;
		}
		files[index].status = 'failed';
		files[index].message = '上传失败';
	});
};

const onUploading = (file) => {
	file.status = 'uploading';
	file.message = '上传中';
	files.push(file);
};

const uploadFactory = async (file) => {
	// 设置为上传中
	onUploading(file);
	// 获取文件二进制数据
	const binary = await toBinary(file.url);
	// 获取文件hash
	const hash = binaryToSha(binary);
	// 进行预传
	const data = await PrepareUpload({
		directoryPath: props.path,
		name: file.name,
		sha: hash,
		size: Math.ceil(file.size / 1024)
	}).catch(() => {
		onFailed(file);
	});
	// 上传失败直接返回
	if (!data) {
		return;
	}
	// 触发秒传
	if (data.uploaded) {
		onSuccess(file, data);
		return;
	}

	console.log(data);
	const count = data.chunkCount;
	const size = data.chunkSize;
	const uploaed = data.uploadChunks;
	const taskArr = [];

	if (count <= 1) {
		taskArr.push(Upload(new Blob([binary]), { uploadId: data.uploadId, index: 1 }));
	} else {
		for (let i = 0; i < count; i += 1) {
			if (uploaed.includes(i + 1)) continue;
			let sliceBinary = new ArrayBuffer(0);
			if ((i + 1) * size > binary.byteLength) {
				sliceBinary = binary.slice(i * size, binary.byteLength);
			} else {
				sliceBinary = binary.slice(i * size, (i + 1) * size);
			}
			taskArr.push(Upload(new Blob([sliceBinary]), { uploadId: data.uploadId, index: i + 1 }));
		}
	}
	taskArr.forEach((fn, index) => {
		fn.then((res) => {
			if (index + 1 == taskArr.length) {
				onSuccess(file, res);
			}
		}).catch((res) => {
			onFailed(file);
		});
	});
};
</script>

<style lang="scss" scoped></style>
