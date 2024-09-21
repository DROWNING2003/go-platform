import {
	defineStore
} from 'pinia';

import {
	GetCurrentUser,
	Logout
} from "@/api/system/usercenter";

import {
	nav
} from '@/library/nav';
import {
	removeToken
} from '../auth';

export const useUserStore = defineStore('user', {
	state: () => {
		return {
			id: 0,
			realName: "",
			nickName: "",
			username: "",
			phone: "",
			email: "",
			avatar: "",
			avatarUrl: "",
			gender: "",
			isLogin: false,
		};
	},
	getters: {

	},
	actions: {
		async userinfo() {
			const data = await GetCurrentUser();
			this.setInfo({
				...data,
				isLogin: true
			});
		},
		setInfo(partial) {
			this.$patch(partial);
		},
		async logout() {
			await Logout();
			this.$patch({
				isLogin: false
			});
			removeToken()
			nav.login();
		}
	},
});