<template>
	<view>
		<uv-datetime-picker
			ref="datetimePicker"
			v-model="datetimeValue"
			:mode="datetimeMode"
			@confirm="datetimeConfirm"
		></uv-datetime-picker>
		<uv-picker ref="picker" :columns="pickerColumns" @confirm="pickerConfirm"></uv-picker>
		<view class="form">
			<uv-form ref="formRef" label-position="top" label-width="auto" :model="formModel" :rules="rules">
				<template v-for="(component, index) in components" :key="index">
					<template v-if="component.type === 'input'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								clearable
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
								:maxlength="component.config.max_length"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'textarea'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-textarea
								v-model="formModel[component.field]"
								clearable
								count
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
								:maxlength="component.config.max_length"
							></uv-textarea>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'number'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-number-box
								class="number-box"
								inputWidth="100%"
								buttonSize="37"
								v-model="formModel[component.field]"
								:min="component.config.min"
								:max="component.config.max"
								:step="component.config.step"
								border="bottom"
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
							></uv-number-box>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'password'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								type="password"
								clearable
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
								:maxlength="component.config.max_length"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'radio'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-radio-group v-model="formModel[component.field]" placement="cloumn">
								<template v-for="(option, optInd) in component.config.options" :key="optInd">
									<uv-radio
										:customStyle="{ padding: '8px 0px' }"
										:label="option"
										:name="option"
									></uv-radio>
								</template>
							</uv-radio-group>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'checkbox'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-checkbox-group v-model="formModel[component.field]" placement="cloumn">
								<template v-for="(option, optInd) in component.config.options" :key="optInd">
									<uv-checkbox
										:customStyle="{ padding: '8px 0px' }"
										:label="option"
										:name="option"
									></uv-checkbox>
								</template>
							</uv-checkbox-group>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'select'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								@click="handleSelect(component)"
								suffixIcon="arrow-right"
								readonly
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
								:maxlength="component.config.max_length"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'date'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								@click="handleDatetime(component, 'date')"
								suffixIcon="arrow-right"
								readonly
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'time'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								@click="handleDatetime(component, 'time')"
								suffixIcon="arrow-right"
								readonly
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'datetime'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<uv-input
								v-model="formModel[component.field]"
								@click="handleDatetime(component, 'datetime')"
								suffixIcon="arrow-right"
								readonly
								:disabled="component.config.disabled"
								:placeholder="component.config.placeholder"
							></uv-input>
						</uv-form-item>
					</template>

					<template v-if="component.type === 'upload'">
						<uv-form-item :label="component.config.label" :prop="component.field">
							<me-upload
								:fileList="files[component.field]"
								width="80"
								height="80"
								:maxCount="1"
								:accept="component.config.filetype"
								@change="uploadSuccess(component, $event)"
								:path="'task/' + id"
								app="PartyAffairs"
							></me-upload>
						</uv-form-item>
					</template>
				</template>
				<uv-button type="primary" text="提交" customStyle="margin-top: 10px" @click="submit"></uv-button>
			</uv-form>
		</view>
	</view>
</template>

<script setup>
import { computed, ref } from 'vue';
import { handleModel, handleRules } from './form.js';
import { parseTime } from '@/library/global/time.js';
import { getBySha } from '@/common/api/system/resource.js';

const datetimePicker = ref();
const datetimeValue = ref();
const datetimeMode = ref('date');
const emit = defineEmits(['submit']);
const props = defineProps({
	components: Array,
	value: Object,
	id: String,
	files: {
		type: Object,
		default: () => {
			return {};
		}
	}
});
const currentComponent = ref({});
const picker = ref(null);
const pickerColumns = ref([]);

// 组件列表
const modelValue = handleModel(props.components);
console.log(modelValue);
// 填写model
const formRef = ref(null);
const formModel = ref({});

if (props.value) {
	formModel.value = Object.assign({}, modelValue, props.value);
}

// 填写规则
const rules = handleRules(props.components);
// 处理提交
const submit = () => {
	formRef.value.validate().then((res) => {
		emit('submit', formModel.value);
	});
};

const pickerConfirm = (val) => {
	formModel.value[currentComponent.value.field] = val.value[0];
};

const datetimeConfirm = (val) => {
	let value = '';
	if (datetimeMode.value == 'date') {
		value = parseTime(val.value, '{y}年{m}月{d}日');
	}
	if (datetimeMode.value == 'datetime') {
		value = parseTime(val.value, '{y}年{m}月{d}日 {h}时{i}分');
	}
	if (datetimeMode.value.trim() == 'time') {
		const arr = val.value.split(':');
		value = arr[0] + '时' + arr[1] + '分';
	}
	formModel.value[currentComponent.value.field] = value;
};

const handleSelect = (component) => {
	if (component.config.disable) {
		return;
	}
	currentComponent.value = component;
	pickerColumns.value = [component.config.options];
	picker.value.open();
};

const dateFormatter = (type, value) => {
	if (type === 'year') {
		return `${value}年`;
	}
	if (type === 'month') {
		return `${value}月`;
	}
	if (type === 'day') {
		return `${value}日`;
	}
	return value;
};

const timeFormatter = (type, value) => {
	if (type === 'hour') {
		return `${value}时`;
	}
	if (type === 'minute') {
		return `${value}分`;
	}
	return value;
};

const datetimeFormatter = (type, value) => {
	if (type === 'year') {
		return `${value}年`;
	}
	if (type === 'month') {
		return `${value}月`;
	}
	if (type === 'day') {
		return `${value}日`;
	}
	if (type === 'hour') {
		return `${value}时`;
	}
	if (type === 'minute') {
		return `${value}分`;
	}
	return value;
};

const handleDatetime = (component, mode) => {
	if (component.config.disable) {
		return;
	}
	currentComponent.value = component;
	datetimeMode.value = mode;
	if (mode == 'date') {
		datetimePicker.value.setFormatter(dateFormatter);
	}
	if (mode == 'datetime') {
		datetimePicker.value.setFormatter(datetimeFormatter);
	}
	if (mode == 'time') {
		datetimePicker.value.setFormatter(timeFormatter);
	}
	datetimePicker.value.open();
};

const uploadSuccess = (component, files) => {
	if (!(files && files.length)) {
		return;
	}
	formModel.value[component.field] = files[0].sha;
};
</script>

<style lang="scss">
.number-box {
	width: 100%;
}
</style>
