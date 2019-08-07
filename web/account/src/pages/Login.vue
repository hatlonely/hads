<template>
  <v-app>
    <v-content>
      <v-layout align-center justify-center row fill-height text-center>
        <v-flex xs12 md4>
          <v-card class="mx-auto" flat outlined rounded="10">
            <v-flex mt-12 mb-5 xs12 lg12>
              <v-layout align-center justify-center>
                <v-img :src="require('../assets/logo.png')" max-width="100" inline></v-img>
              </v-layout>
            </v-flex>

            <v-flex my-5 xs12 lg12>
              <div class="headline">登陆</div>
            </v-flex>
            <v-flex my-5 xs12 lg12>
              <div class="body-1">使用您的 hads 账号</div>
            </v-flex>
            <v-flex my-12 mx-12>
              <v-text-field label="电子邮件或者电话号码" placeholder="account@gmail.com" outlined filled></v-text-field>
              <div class="body-2 text-lg-left">
                <a>
                  <strong>忘记了电子邮件地址？</strong>
                </a>
              </div>
            </v-flex>
            <v-flex my-12 mx-12>
              <div class="body-2 text-lg-left">
                <p>不是您自己的计算机？请使用访客模式无痕登录。</p>
                <p>
                  <a>
                    <strong>了解详情</strong>
                  </a>
                </p>
                <p></p>
              </div>
            </v-flex>
            <v-flex my-12 mx-12>
              <v-layout align-left justify-center row fill-height text-left>
                <v-flex xs3>
                  <v-btn text color="primary" pl-0>创建账号</v-btn>
                </v-flex>
                <v-flex xs6></v-flex>
                <v-flex xs3>
                  <v-btn color="primary" dark>下一步</v-btn>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-card>
          <div class="caption text-lg-right login-footer mt-12">
            <a href="#">帮助</a>&nbsp;&nbsp;&nbsp;&nbsp;
            <a href="#">隐私</a>&nbsp;&nbsp;&nbsp;&nbsp;
            <a href="#">条款</a>
          </div>
        </v-flex>
      </v-layout>
    </v-content>
  </v-app>
</template>

<style>
.login-footer a {
  color: rgba(0, 0, 0, 0.54);
}
</style>

<script>
export default {};
</script>

