import {
	getToken
} from "./index.js";
import {
	nav
} from "../nav";

// 页面白名单，不受拦截
const whiteList = [
	'/pages/user/login/index',
	'/pages/user/login/email',
	'/pages/user/login/password',
	'/pages/index/index',
	'/pages/user/index',
	'/'
]

const funcList = []
// ["navigateTo", "redirectTo", "reLaunch", "switchTab"];


function hasPermission(url) {
	const token = getToken()
	return whiteList.indexOf(url) !== -1 || token
}

funcList.forEach(item => {
	uni.addInterceptor(item, {
		invoke(e) {
			const arr = e.url.split("?")
			const path = arr[0]
			if (!hasPermission(path)) {
				nav.login();
				return false
			}
			return true
		},
		success(e) {}
	})
})