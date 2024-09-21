import {
	ref
} from "vue"
const isRedirectToLogin = ref(false);

export const nav = {
	home: () => {
		uni.switchTab({
			url: '/pages/index/index'
		});
	},
	login: () => {
		if (!isRedirectToLogin.value) {
			isRedirectToLogin.value = true
			uni.redirectTo({
				url: "/pages/user/login/index?type=login"
			})
			setTimeout(() => {
				isRedirectToLogin.value = false
			}, 500)
		}
	},
	bind: (id) => {
		uni.redirectTo({
			url: "/pages/user/login/index?type=bind&oAuthUid=" + id
		})
	},
	register: () => {
		if (!isRedirectToLogin.value) {
			isRedirectToLogin.value = true
			uni.redirectTo({
				url: "/pages/user/login/index?type=register"
			})
			setTimeout(() => {
				isRedirectToLogin.value = false
			}, 500)
		}
	},
	error: (text) => {
		uni.redirectTo({
			url: "/pages/error/index?text=" + text
		})
	},
	open: (path) => {
		if (path.indexOf("http") === 0) {
			uni.navigateTo({
				url: "/pages/app/webview/index?src=" + path
			})
		} else {
			uni.navigateTo({
				url: path
			})
		}
	}
}