import {
	useAppStore
} from "@/library/store/app";
import config from "@/config"

const TokenKey = 'uni_app_token';
const PlatformKey = 'uni_app_platform';

export function hasToken() {
	return !!uni.getStorageSync(TokenKey)
}

export function getToken() {
	return uni.getStorageSync(TokenKey)
}

export function setToken(token) {
	return uni.setStorageSync(TokenKey, token)
}

export function removeToken() {
	return uni.removeStorageSync(TokenKey)
}


export function getAuthCode() {
	const appStore = useAppStore();
	return new Promise((reslove, reject) => {
		switch (appStore.platform) {
			case "yb":
				if (appStore.query.verify_request) reslove(appStore.query.verify_request)
				else reject('请重新打开之后进入')
				break
			case "mp":
			case "qq":
				uni.login({
					success(res) {
						if (res.code) reslove(res.code)
						else reject(res.errMsg)
					},
					fail(e) {
						reject(res.errMsg)
					}
				})
				break;
		}

	})

}