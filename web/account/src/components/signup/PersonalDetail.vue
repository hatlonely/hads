<template>
  <v-card width="748" height="511" flat outlined>
    <v-layout row wrap mx-0 text-left px-8>
      <v-flex xs8 pr-8>
        <v-flex mt-8 mb-5 xs12 lg12>
          <v-layout>
            <v-img :src="require('../../assets/logo.png')" max-width="70" inline></v-img>
          </v-layout>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <div class="headline">欢迎使用 hads</div>
        </v-flex>
        <v-flex my-4 xs12 lg12>
          <!-- <v-icon>account_circle</v-icon>&nbsp;&nbsp; hatlonely1992@gmail.com -->
          <v-layout align-center justify-left>
            <v-card class="h-round-card px-2 py-1" outlined flat>
              <v-icon class="mr-2">account_circle</v-icon>
              <span class="body-1">hatlonely@gmail.com</span>
            </v-card>
          </v-layout>
        </v-flex>

        <v-flex mt-10 mb-4>
          <v-layout mx-0 row wrap>
            <v-flex xs12>
              <v-text-field label="生日" outlined filled></v-text-field>
            </v-flex>
            <v-flex xs12>
              <v-text-field label="性别" outlined filled></v-text-field>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex mt-0 mx-0>
          <v-layout align-left justify-center row fill-height text-left>
            <v-flex xs3>
              <v-btn text color="primary" to="/signup/verifyphone" pl-0>上一步</v-btn>
            </v-flex>
            <v-flex xs6></v-flex>
            <v-flex xs3>
              <v-btn color="primary" depressed to="/signup/privacy">下一步</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-flex>

      <v-flex xs4>
        <v-layout align-center justify-center row fill-height mx-0 class="text-center">
          <v-flex xs12>
            <v-flex xs12>
              <v-img :src="require('../../assets/signup_phone.svg')" max-width="244" inline></v-img>
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
        verifycode: v => (!!v && v.match(/^[0-9]{6}$/)) || "请输入正确的验证码",
        phone: v => v.match(/^1[345789][0-9]{9}$/) || "请输入正确的电话号码哦"
      }
    };
  }
};
</script>
