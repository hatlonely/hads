<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-8>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/img/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <h2>验证您的手机号码</h2>
        </v-flex>
        <v-flex my-10>
          <p class="body-2">为了安全起见，hads 希望确定是您本人在登录。hads 将通过短信向您发送一个 6 位数的验证码</p>
        </v-flex>
        <v-flex mt-10 mb-4>
          <p class="body-2">
            <i>需要按标准费率支付费用</i>
          </p>
          <v-layout mx-0 row wrap>
            <v-flex xs12>
              <v-text-field label="电话号码" :rules="[rules.required, rules.phone]" outlined filled></v-text-field>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" to="/signup" pl-0>上一步</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed to="/signup/verifycode">下一步</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/img/signup_phone.svg')" max-width="244" inline></v-img>
            </v-flex>
            <v-flex xs12 px-4>
              <p class="body-1">您的个人信息不会对外公开，并且绝对安全</p>
            </v-flex>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-card>
</template>

<script>
export default {
  data() {
    return {
      show: false,
      rules: {
        required: v => !!v || "必要字段",
        min: v => (!!v && v.length >= 8) || "至少8个字符",
        phone: v => v.match(/^1[345789][0-9]{9}$/) || "请输入正确的电话号码哦"
      }
    };
  }
};
</script>
