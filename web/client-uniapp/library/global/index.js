import resource from "./resource.js"
import logo from '@/static/logo.png';

import {
	formatTime,
	parseTime
} from "./time.js"
import {
	useUserStore
} from "@/library/store/user";



export default (app) => {
	uni.$global = {
		$rurl: resource,
		$ft: formatTime,
		$pt: parseTime,
		$back: () => {
			uni.navigateBack()
		},
		$to: (path) => {
			uni.navigateTo({
				url: path
			});
		},
		$logo: logo,
		$user: useUserStore()
	}

	const keys = Object.keys(uni.$global)
	keys.forEach(key => {
		app.config.globalProperties[key] = uni.$global[key]
	})
}