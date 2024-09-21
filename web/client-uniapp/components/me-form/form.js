import {
	getBySha
} from "@/common/api/system/resource";
import formatUrl from "@/library/global/resource";

const rules = {
	require: {
		required: true,
		message: '不能为空',
		trigger: ['change', 'blur']
	},
	string: {
		type: "string",
		message: '必须为字符',
		trigger: ['change', 'blur']
	},
	number: {
		type: "number",
		message: '必须为数字',
		trigger: ['change', 'blur']
	},
	integer: {
		type: "integer",
		message: '必须为整数',
		trigger: ['change', 'blur']
	},
	float: {
		type: "float",
		message: '必须为浮点数',
		trigger: ['change', 'blur']
	},
	date: {
		type: "date",
		message: '必须为日期',
		trigger: ['change', 'blur']
	},
	url: {
		type: "url",
		message: '必须为链接',
		trigger: ['change', 'blur']
	},
	email: {
		type: "email",
		message: '必须为邮箱',
		trigger: ['change', 'blur']
	},
	phone: {
		validator: (rule, value, callback) => {
			return uni.$uv.test.mobile(value);
		},
		message: '必须为手机号',
		trigger: ['change', 'blur'],
	},
	idcard: {
		validator: (rule, value, callback) => {
			return uni.$uv.test.idCard(value);
		},
		message: '必须为身份证号',
		trigger: ['change', 'blur'],
	},
}

export const handleModel = (components) => {
	const model = {};
	components.forEach((item) => {
		model[item.field] = item.config.value;
	});
	return model;
};

export const handleRules = (components) => {
	const formRule = {};
	components.forEach((item) => {
		formRule[item.field] = [];
		if (item.config.required) {
			const rule = {
				required: true,
				message: item.config.label + '为必填项',
				trigger: ['blur', 'change']
			}
			if (item.type === "checkbox") {
				rule.type = "array"
			}
			formRule[item.field].push(rule);
		}

		if (item.config.rule && rules[item.config.rule.value]) {
			const rule = uni.$uv.deepClone(rules[item.config.rule.value])
			rule.message = item.config.label + rule.message
			formRule[item.field].push(rule);
		}
	});
	return formRule;
}