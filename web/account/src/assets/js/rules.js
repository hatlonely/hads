export default {
    required: v => !!v || "必要字段",
    atleast8characters: v => v.length >= 8 || "至少8个字符",
    atmost64characters: v => v.length <= 64 || "至多64个字符",
    atmost32characters: v => v.length <= 32 || "至多32个字符",
    validemail: v => /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/.test(v) || "无效的 email",
    validphone: v => /^1[345789][0-9]{9}$/.test(v) || "无效的电话号码哦",
    validcode: v => /^[0-9]{6}$/.test(v) || "无效的验证码",
    validPhoneOrEmail: v => /^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$/.test(v) || /^1[345789][0-9]{9}$/.test(v) || "无效的电话或邮箱"
}