import repuest from "@/library/request"
import config from "@/config"
import crypto from "@/library/crypto"

export const GetCurrentUser = (params) => {
	return repuest({
		url: "/usercenter/client/v1/user",
		method: "get",
		params
	})
}

export const UpdateCurrentUser = (params) => {
	return repuest({
		url: "/usercenter/client/v1/user",
		method: "put",
		params
	})
}

export const GetApp = () => {
	return repuest({
		url: "/usercenter/client/v1/app?keyword=" + config.server,
		method: "get"
	})
}

export const GenAuthCaptch = (params) => {
	return repuest({
		url: "/usercenter/client/v1/auth/captcha",
		method: "post",
		params
	})
}


export const OAuthLogin = (params) => {
	return repuest({
		url: "/usercenter/client/v1/login/oauth",
		method: "post",
		params
	})
}

export const EmailLogin = (params) => {
	return repuest({
		url: "/usercenter/client/v1/login/email",
		method: "post",
		params
	})
}

export const PasswordLogin = (ps) => {
	const params = {
		...ps
	}
	params.password = crypto.rsa().encode({
		password: params.password,
		time: new Date().getTime()
	})
	return repuest({
		url: "/usercenter/client/v1/login/password",
		method: "post",
		params
	})
}

export const EmailRegister = (params) => {
	return repuest({
		url: "/usercenter/client/v1/register/email",
		method: "post",
		params
	})
}

export const PasswordRegister = (ps) => {
	const params = {
		...ps
	}
	params.password = crypto.rsa().encode({
		password: params.password,
		time: new Date().getTime()
	})
	return repuest({
		url: "/usercenter/client/v1/register/password",
		method: "post",
		params
	})
}

export const EmailBind = (params) => {
	return repuest({
		url: "/usercenter/client/v1/bind/email",
		method: "post",
		params
	})
}

export const PasswordBind = (ps) => {
	const params = {
		...ps
	}
	params.password = crypto.rsa().encode({
		password: params.password,
		time: new Date().getTime()
	})
	return repuest({
		url: "/usercenter/client/v1/bind/password",
		method: "post",
		params
	})
}

export const RefreshToken = (params) => {
	return repuest({
		url: "/usercenter/client/v1/token/refresh",
		method: "post",
		params
	})
}

export const ListFeedbackCategory = (params) => {
	return repuest({
		url: "/usercenter/client/v1/feedback_categories",
		method: "get",
		params
	})
}

export const CreateFeedback = (params) => {
	return repuest({
		url: "/usercenter/client/v1/feedback",
		method: "post",
		params
	})
}

export const Logout = () => {
	return repuest({
		url: "/usercenter/client/v1/logout",
		method: "post"
	})
}