export default {
    required: v => !!v || "必要字段",
    atleast8characters: v => (!!v && v.length >= 8) || "至少8个字符",
    validemail: v => (!!v && /.+@.+\..+/.test(v)) || "请输入一个有效的 email",
    validphone: v =>
        (!!v && !!/^1[345789][0-9]{9}$/.test(v)) || "请输入正确的电话号码哦",
    validcode: v => (!!v && /^[0-9]{6}$/.test(v)) || "请输入正确的验证码",
}