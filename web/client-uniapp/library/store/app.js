import {
	defineStore,
} from 'pinia';
import {
	nav
} from '@/library/nav';
import {
	getAuthCode
} from '../auth';

const platform = () => {
	return new Promise((resolve) => {
		// #ifdef H5

		// 特殊内嵌平台兼容
		// yb
		document.addEventListener("YibanDeviceReady", () => {
			resolve("yb")
		});
		if (window.yiban) {
			resolve("yb")
			return
		}
		// #endif

		resolve(uni.$uv.platform)
	})
}


export const useAppStore = defineStore('app', {
	state: () => {
		return {
			id: 0,
			name: "",
			logoUrl: "",
			keyword: "",
			status: false,
			channels: [],
			fields: [],
			description: "",
			allowRegistry: false,
			copyright: "",
			version: "",
			platform: "",
			query: {},
		};
	},
	getters: {
		suggestChannel: (state) => {
			let data = {}
			state.channels.forEach(item => {
				if (item.keyword === state.platform) {
					data = item
				}
			})
			return data
		},
		loginTypes: (state) => {
			return {
				bind: '绑定',
				login: '登陆',
				register: '注册'
			}
		}
	},
	actions: {
		async set(info) {
			const pt = await platform();
			// #ifdef H5
			window.document.title = info.name
			// #endif
			this.$patch({
				...info,
				platform: pt,
			});
		},
		authCode: async (state) => {
			return getAuthCode()
		}

	},
});